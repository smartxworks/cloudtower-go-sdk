// Code generated by go-swagger; DO NOT EDIT.

package consistency_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetConsistencyGroupsConnectionReader is a Reader for the GetConsistencyGroupsConnection structure.
type GetConsistencyGroupsConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConsistencyGroupsConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConsistencyGroupsConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewGetConsistencyGroupsConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewGetConsistencyGroupsConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewGetConsistencyGroupsConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConsistencyGroupsConnectionOK creates a GetConsistencyGroupsConnectionOK with default headers values
func NewGetConsistencyGroupsConnectionOK() *GetConsistencyGroupsConnectionOK {
	return &GetConsistencyGroupsConnectionOK{}
}

/* GetConsistencyGroupsConnectionOK describes a response with status code 200, with default header values.

GetConsistencyGroupsConnectionOK get consistency groups connection o k
*/
type GetConsistencyGroupsConnectionOK struct {
	XTowerRequestID string

	Payload *models.ConsistencyGroupConnection
}

func (o *GetConsistencyGroupsConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-consistency-groups-connection][%d] getConsistencyGroupsConnectionOK  %+v", 200, o.Payload)
}
func (o *GetConsistencyGroupsConnectionOK) GetPayload() *models.ConsistencyGroupConnection {
	return o.Payload
}

func (o *GetConsistencyGroupsConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.ConsistencyGroupConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConsistencyGroupsConnectionBadRequest creates a GetConsistencyGroupsConnectionBadRequest with default headers values
func NewGetConsistencyGroupsConnectionBadRequest() *GetConsistencyGroupsConnectionBadRequest {
	return &GetConsistencyGroupsConnectionBadRequest{}
}

/* GetConsistencyGroupsConnectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetConsistencyGroupsConnectionBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetConsistencyGroupsConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-consistency-groups-connection][%d] getConsistencyGroupsConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetConsistencyGroupsConnectionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConsistencyGroupsConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetConsistencyGroupsConnectionNotFound creates a GetConsistencyGroupsConnectionNotFound with default headers values
func NewGetConsistencyGroupsConnectionNotFound() *GetConsistencyGroupsConnectionNotFound {
	return &GetConsistencyGroupsConnectionNotFound{}
}

/* GetConsistencyGroupsConnectionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetConsistencyGroupsConnectionNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetConsistencyGroupsConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /get-consistency-groups-connection][%d] getConsistencyGroupsConnectionNotFound  %+v", 404, o.Payload)
}
func (o *GetConsistencyGroupsConnectionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConsistencyGroupsConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetConsistencyGroupsConnectionInternalServerError creates a GetConsistencyGroupsConnectionInternalServerError with default headers values
func NewGetConsistencyGroupsConnectionInternalServerError() *GetConsistencyGroupsConnectionInternalServerError {
	return &GetConsistencyGroupsConnectionInternalServerError{}
}

/* GetConsistencyGroupsConnectionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetConsistencyGroupsConnectionInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetConsistencyGroupsConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-consistency-groups-connection][%d] getConsistencyGroupsConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetConsistencyGroupsConnectionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConsistencyGroupsConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
