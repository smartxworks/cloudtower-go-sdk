// Code generated by go-swagger; DO NOT EDIT.

package cluster_upgrade_history

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetClusterUpgradeHistoriesConnectionReader is a Reader for the GetClusterUpgradeHistoriesConnection structure.
type GetClusterUpgradeHistoriesConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterUpgradeHistoriesConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterUpgradeHistoriesConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewGetClusterUpgradeHistoriesConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewGetClusterUpgradeHistoriesConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewGetClusterUpgradeHistoriesConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetClusterUpgradeHistoriesConnectionOK creates a GetClusterUpgradeHistoriesConnectionOK with default headers values
func NewGetClusterUpgradeHistoriesConnectionOK() *GetClusterUpgradeHistoriesConnectionOK {
	return &GetClusterUpgradeHistoriesConnectionOK{}
}

/* GetClusterUpgradeHistoriesConnectionOK describes a response with status code 200, with default header values.

GetClusterUpgradeHistoriesConnectionOK get cluster upgrade histories connection o k
*/
type GetClusterUpgradeHistoriesConnectionOK struct {
	XTowerRequestID string

	Payload *models.ClusterUpgradeHistoryConnection
}

func (o *GetClusterUpgradeHistoriesConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-cluster-upgrade-histories-connection][%d] getClusterUpgradeHistoriesConnectionOK  %+v", 200, o.Payload)
}
func (o *GetClusterUpgradeHistoriesConnectionOK) GetPayload() *models.ClusterUpgradeHistoryConnection {
	return o.Payload
}

func (o *GetClusterUpgradeHistoriesConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.ClusterUpgradeHistoryConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterUpgradeHistoriesConnectionBadRequest creates a GetClusterUpgradeHistoriesConnectionBadRequest with default headers values
func NewGetClusterUpgradeHistoriesConnectionBadRequest() *GetClusterUpgradeHistoriesConnectionBadRequest {
	return &GetClusterUpgradeHistoriesConnectionBadRequest{}
}

/* GetClusterUpgradeHistoriesConnectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetClusterUpgradeHistoriesConnectionBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetClusterUpgradeHistoriesConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-cluster-upgrade-histories-connection][%d] getClusterUpgradeHistoriesConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetClusterUpgradeHistoriesConnectionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetClusterUpgradeHistoriesConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterUpgradeHistoriesConnectionNotFound creates a GetClusterUpgradeHistoriesConnectionNotFound with default headers values
func NewGetClusterUpgradeHistoriesConnectionNotFound() *GetClusterUpgradeHistoriesConnectionNotFound {
	return &GetClusterUpgradeHistoriesConnectionNotFound{}
}

/* GetClusterUpgradeHistoriesConnectionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetClusterUpgradeHistoriesConnectionNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetClusterUpgradeHistoriesConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /get-cluster-upgrade-histories-connection][%d] getClusterUpgradeHistoriesConnectionNotFound  %+v", 404, o.Payload)
}
func (o *GetClusterUpgradeHistoriesConnectionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetClusterUpgradeHistoriesConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterUpgradeHistoriesConnectionInternalServerError creates a GetClusterUpgradeHistoriesConnectionInternalServerError with default headers values
func NewGetClusterUpgradeHistoriesConnectionInternalServerError() *GetClusterUpgradeHistoriesConnectionInternalServerError {
	return &GetClusterUpgradeHistoriesConnectionInternalServerError{}
}

/* GetClusterUpgradeHistoriesConnectionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetClusterUpgradeHistoriesConnectionInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetClusterUpgradeHistoriesConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-cluster-upgrade-histories-connection][%d] getClusterUpgradeHistoriesConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetClusterUpgradeHistoriesConnectionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetClusterUpgradeHistoriesConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
