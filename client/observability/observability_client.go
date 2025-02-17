// Code generated by go-swagger; DO NOT EDIT.

package observability

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new observability API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for observability API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ClearSystemServiceAlertNotificationConfig(params *ClearSystemServiceAlertNotificationConfigParams, opts ...ClientOption) (*ClearSystemServiceAlertNotificationConfigOK, error)

	DisassociateSystemServiceFromObsService(params *DisassociateSystemServiceFromObsServiceParams, opts ...ClientOption) (*DisassociateSystemServiceFromObsServiceOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ClearSystemServiceAlertNotificationConfig clear system service alert notification config API
*/
func (a *Client) ClearSystemServiceAlertNotificationConfig(params *ClearSystemServiceAlertNotificationConfigParams, opts ...ClientOption) (*ClearSystemServiceAlertNotificationConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewClearSystemServiceAlertNotificationConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ClearSystemServiceAlertNotificationConfig",
		Method:             "POST",
		PathPattern:        "/clear-system-service-alert-notification-config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ClearSystemServiceAlertNotificationConfigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ClearSystemServiceAlertNotificationConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ClearSystemServiceAlertNotificationConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DisassociateSystemServiceFromObsService disassociate system service from obs service API
*/
func (a *Client) DisassociateSystemServiceFromObsService(params *DisassociateSystemServiceFromObsServiceParams, opts ...ClientOption) (*DisassociateSystemServiceFromObsServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDisassociateSystemServiceFromObsServiceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DisassociateSystemServiceFromObsService",
		Method:             "POST",
		PathPattern:        "/disassociate-system-service-from-obs-service",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DisassociateSystemServiceFromObsServiceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DisassociateSystemServiceFromObsServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DisassociateSystemServiceFromObsService: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
