// Code generated by go-swagger; DO NOT EDIT.

package snapshot_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetSnapshotGroupsConnectionReader is a Reader for the GetSnapshotGroupsConnection structure.
type GetSnapshotGroupsConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSnapshotGroupsConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSnapshotGroupsConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewGetSnapshotGroupsConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewGetSnapshotGroupsConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewGetSnapshotGroupsConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSnapshotGroupsConnectionOK creates a GetSnapshotGroupsConnectionOK with default headers values
func NewGetSnapshotGroupsConnectionOK() *GetSnapshotGroupsConnectionOK {
	return &GetSnapshotGroupsConnectionOK{}
}

/* GetSnapshotGroupsConnectionOK describes a response with status code 200, with default header values.

GetSnapshotGroupsConnectionOK get snapshot groups connection o k
*/
type GetSnapshotGroupsConnectionOK struct {
	XTowerRequestID string

	Payload *models.SnapshotGroupConnection
}

func (o *GetSnapshotGroupsConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-snapshot-groups-connection][%d] getSnapshotGroupsConnectionOK  %+v", 200, o.Payload)
}
func (o *GetSnapshotGroupsConnectionOK) GetPayload() *models.SnapshotGroupConnection {
	return o.Payload
}

func (o *GetSnapshotGroupsConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.SnapshotGroupConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSnapshotGroupsConnectionBadRequest creates a GetSnapshotGroupsConnectionBadRequest with default headers values
func NewGetSnapshotGroupsConnectionBadRequest() *GetSnapshotGroupsConnectionBadRequest {
	return &GetSnapshotGroupsConnectionBadRequest{}
}

/* GetSnapshotGroupsConnectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetSnapshotGroupsConnectionBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetSnapshotGroupsConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-snapshot-groups-connection][%d] getSnapshotGroupsConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetSnapshotGroupsConnectionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSnapshotGroupsConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSnapshotGroupsConnectionNotFound creates a GetSnapshotGroupsConnectionNotFound with default headers values
func NewGetSnapshotGroupsConnectionNotFound() *GetSnapshotGroupsConnectionNotFound {
	return &GetSnapshotGroupsConnectionNotFound{}
}

/* GetSnapshotGroupsConnectionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetSnapshotGroupsConnectionNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetSnapshotGroupsConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /get-snapshot-groups-connection][%d] getSnapshotGroupsConnectionNotFound  %+v", 404, o.Payload)
}
func (o *GetSnapshotGroupsConnectionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSnapshotGroupsConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSnapshotGroupsConnectionInternalServerError creates a GetSnapshotGroupsConnectionInternalServerError with default headers values
func NewGetSnapshotGroupsConnectionInternalServerError() *GetSnapshotGroupsConnectionInternalServerError {
	return &GetSnapshotGroupsConnectionInternalServerError{}
}

/* GetSnapshotGroupsConnectionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetSnapshotGroupsConnectionInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetSnapshotGroupsConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-snapshot-groups-connection][%d] getSnapshotGroupsConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetSnapshotGroupsConnectionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSnapshotGroupsConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
