package watchor

import (
	"errors"
	"math"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/smartxworks/cloudtower-go-sdk/v2/client"
	resource_change_client "github.com/smartxworks/cloudtower-go-sdk/v2/client/resource_change"
	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
	utils "github.com/smartxworks/cloudtower-go-sdk/v2/utils"
)

var (
	// ErrAlreadyStarted is returned when attempting to start an already started watch client
	ErrAlreadyStarted = errors.New("watch client already started")
	// ErrNotStarted is returned when attempting to operate on a watch client that hasn't been started
	ErrNotStarted = errors.New("watch client not started")
)

// WarningEvent represents a non-fatal error during resource change watching
type WarningEvent struct {
	Err error
}

// ErrorEventType defines the type of error that occurred during resource change watching
type ErrorEventType string

const (
	// ErrorEventTypeCompacted indicates the requested revision has been compacted
	ErrorEventTypeCompacted ErrorEventType = "compacted"
	// ErrorEventTypeRequest indicates a request error occurred
	ErrorEventTypeRequest ErrorEventType = "request"
	// ErrorEventTypeUnsupported indicates an unsupported error occurred
	ErrorEventTypeUnsupported ErrorEventType = "unsupported"
)

// ErrorEvent represents a fatal error during resource change watching
type ErrorEvent struct {
	Type            ErrorEventType
	Err             error
	CompactRevision *string
}

// ResourceChangeWatchClient watches for resource changes in CloudTower
type ResourceChangeWatchClient struct {
	client        *client.Cloudtower
	clientOptions resource_change_client.ClientOption

	// channels
	channel        chan *models.ResourceChangeEvent
	warningChannel chan *WarningEvent
	errorChannel   chan *ErrorEvent

	// shared context
	mu              sync.RWMutex
	currentRevision *string
	started         atomic.Bool
	catchedUp       atomic.Bool

	// initial data, should readonly
	resourceID      *string
	resourceTypes   []string
	maxRetries      int
	pollingInterval time.Duration
}

// NewResourceChangeWatchClientParams contains parameters for creating a new ResourceChangeWatchClient
type NewResourceChangeWatchClientParams struct {
	// Client is the CloudTower client instance
	Client *client.Cloudtower
	// ClientOptions is the options for the client
	ClientOptions resource_change_client.ClientOption
	// MaxRetries is the maximum number of retry attempts
	MaxRetries int

	// ResourceID is the ID of the resource to watch
	ResourceID *string
	// ResourceTypes are the types of resources to watch
	ResourceTypes []string
	// PollingInterval is the interval between polling requests
	PollingInterval time.Duration
}

// NewResourceChangeWatchClient creates a new ResourceChangeWatchClient with the given parameters
func NewResourceChangeWatchClient(params *NewResourceChangeWatchClientParams) (*ResourceChangeWatchClient, error) {
	var error_list []error
	if params.Client == nil {
		error_list = append(error_list, errors.New("client is required"))
	}
	if params.ClientOptions == nil {
		error_list = append(error_list, errors.New("client options is required"))
	}
	if params.MaxRetries < 0 {
		error_list = append(error_list, errors.New("max retries must be greater than 0"))
	} else if params.MaxRetries == 0 {
		params.MaxRetries = 3
	}
	if params.PollingInterval < 0 {
		error_list = append(error_list, errors.New("polling interval must be greater than 0"))
	} else if params.PollingInterval == 0 {
		params.PollingInterval = 1 * time.Second
	}

	if len(error_list) > 0 {
		return nil, errors.Join(error_list...)
	}

	return &ResourceChangeWatchClient{
		client:          params.Client,
		clientOptions:   params.ClientOptions,
		mu:              sync.RWMutex{},
		started:         atomic.Bool{},
		catchedUp:       atomic.Bool{},
		resourceID:      params.ResourceID,
		resourceTypes:   params.ResourceTypes,
		maxRetries:      params.MaxRetries,
		pollingInterval: params.PollingInterval,
	}, nil
}

// ResourceChangeWatchStartParams contains parameters for starting a ResourceChangeWatchClient
type ResourceChangeWatchStartParams struct {
	// StartRevision is the revision to start watching from (can be nil)
	StartRevision *string
}

// Start begins watching for resource change events
func (c *ResourceChangeWatchClient) Start(params *ResourceChangeWatchStartParams) error {
	if c.started.Load() {
		return ErrAlreadyStarted
	}

	if err := c.initialized(params); err != nil {
		return err
	}

	go c.poll()
	return nil
}

func (c *ResourceChangeWatchClient) initialized(params *ResourceChangeWatchStartParams) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.started.Load() {
		return ErrAlreadyStarted
	}

	c.currentRevision = params.StartRevision

	// mark started flag to avoid multiple start
	c.started.Store(true)
	// reset catched up flag on start
	c.catchedUp.Store(false)

	c.channel = make(chan *models.ResourceChangeEvent, 500)
	c.warningChannel = make(chan *WarningEvent)
	c.errorChannel = make(chan *ErrorEvent)
	return nil
}

// Stop marks the client as stopped and stops polling
func (c *ResourceChangeWatchClient) Stop() error {
	if !c.started.Load() {
		return ErrNotStarted
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.started.Load() {
		return ErrNotStarted
	}

	c.started.Store(false)
	return nil
}

// Channel returns the channel for normal resource change events
func (c *ResourceChangeWatchClient) Channel() <-chan *models.ResourceChangeEvent {
	return c.channel
}

// WarningChannel returns the channel for warning events
func (c *ResourceChangeWatchClient) WarningChannel() <-chan *WarningEvent {
	return c.warningChannel
}

// ErrorChannel returns the channel for error events
func (c *ResourceChangeWatchClient) ErrorChannel() <-chan *ErrorEvent {
	return c.errorChannel
}

// poll continuously polls for resource changes until stopped or an error occurs
func (c *ResourceChangeWatchClient) poll() {
	defer func() {
		c.mu.Lock()
		// close channel, but allow to read from it
		if c.channel != nil {
			close(c.channel)
		}
		if c.errorChannel != nil {
			close(c.errorChannel)
		}
		if c.warningChannel != nil {
			close(c.warningChannel)
		}
		c.mu.Unlock()
	}()

	requestErrCount := 0

	for {
		if !c.started.Load() {
			return
		}

		if err := c.pollOnce(); err != nil {
			if err, ok := err.(*models.UnexpectedError); ok {
				if err.Code() == 404 {
					c.writeToErrorChannel(&ErrorEvent{
						Type: ErrorEventTypeUnsupported,
						Err:  err,
					})
					return
				}
			}
			// exponential backoff = interval * 2^i seconds
			time.Sleep(time.Duration(math.Pow(2, float64(requestErrCount))) * c.pollingInterval)
			requestErrCount++
			if requestErrCount > c.maxRetries {
				c.writeToErrorChannel(&ErrorEvent{
					Type: ErrorEventTypeRequest,
					Err:  err,
				})
				return
			} else {
				c.writeToWarningChannel(&WarningEvent{
					Err: err,
				})
			}
		} else {
			requestErrCount = 0
			time.Sleep(c.pollingInterval)
		}
	}
}

// createParams creates the parameters for the GetResourceChanges API call
func (c *ResourceChangeWatchClient) createParams() (*resource_change_client.GetResourceChangesParams, resource_change_client.ClientOption) {
	params := resource_change_client.NewGetResourceChangesParams()
	params.ResourceID = c.resourceID
	resourceTypes := strings.Join(c.resourceTypes, ",")
	if len(resourceTypes) > 0 {
		params.ResourceType = &resourceTypes
	}
	params.StartRevision = c.currentRevision
	var limit int32 = 100
	params.Limit = &limit
	return params, c.clientOptions
}

// pollOnce performs a single poll for resource changes
func (c *ResourceChangeWatchClient) pollOnce() error {
	if !c.started.Load() {
		return ErrNotStarted
	}

	params, clientOptions := c.createParams()

	events, err := c.client.ResourceChange.GetResourceChanges(params, clientOptions)
	if err != nil {
		return err
	}

	if c.currentRevision != nil && c.catchedUp.Load() {
		// first time, or already catched up, handling compact error
		if compare, err := utils.CompareBigIntStrings(c.currentRevision, events.Payload.CompactRevision); err != nil {
			return err
		} else if compare < 0 {
			compactEvent := &ErrorEvent{
				Type:            ErrorEventTypeCompacted,
				CompactRevision: events.Payload.CompactRevision,
			}
			c.writeToErrorChannel(compactEvent)
		}
	}

	for _, event := range events.Payload.Data {
		if !c.started.Load() {
			return nil
		}

		c.currentRevision = event.Revision
		c.writeToChannel(event)
	}

	if !c.catchedUp.Load() {
		if compare, err := utils.CompareBigIntStrings(c.currentRevision, events.Payload.CurrentRevision); err != nil {
			return err
		} else if compare >= 0 {
			// if latest event's revision is greater or equal to current revision
			// mean we already catched up newest data
			c.catchedUp.Store(true)
		}
	}

	return nil
}

// writeToChannel sends an event to the resource change event channel
func (c *ResourceChangeWatchClient) writeToChannel(event *models.ResourceChangeEvent) {
	if c.channel != nil {
		c.channel <- event
	}
}

// writeToWarningChannel sends an event to the warning channel
func (c *ResourceChangeWatchClient) writeToWarningChannel(event *WarningEvent) {
	if c.warningChannel != nil {
		select {
		case <-c.warningChannel:
		default:
		}
		c.warningChannel <- event
	}
}

// writeToErrorChannel sends an event to the error channel and stops the client
func (c *ResourceChangeWatchClient) writeToErrorChannel(event *ErrorEvent) {
	if c.errorChannel != nil {
		c.errorChannel <- event
	}
	c.Stop()
}
