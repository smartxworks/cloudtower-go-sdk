// Code generated by go-swagger; DO NOT EDIT.

package zone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetZonesConnectionReader is a Reader for the GetZonesConnection structure.
type GetZonesConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetZonesConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetZonesConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewGetZonesConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewGetZonesConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewGetZonesConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetZonesConnectionOK creates a GetZonesConnectionOK with default headers values
func NewGetZonesConnectionOK() *GetZonesConnectionOK {
	return &GetZonesConnectionOK{}
}

/* GetZonesConnectionOK describes a response with status code 200, with default header values.

GetZonesConnectionOK get zones connection o k
*/
type GetZonesConnectionOK struct {
	XTowerRequestID string

	Payload *models.ZoneConnection
}

func (o *GetZonesConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-zones-connection][%d] getZonesConnectionOK  %+v", 200, o.Payload)
}
func (o *GetZonesConnectionOK) GetPayload() *models.ZoneConnection {
	return o.Payload
}

func (o *GetZonesConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.ZoneConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetZonesConnectionBadRequest creates a GetZonesConnectionBadRequest with default headers values
func NewGetZonesConnectionBadRequest() *GetZonesConnectionBadRequest {
	return &GetZonesConnectionBadRequest{}
}

/* GetZonesConnectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetZonesConnectionBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetZonesConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-zones-connection][%d] getZonesConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetZonesConnectionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetZonesConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetZonesConnectionNotFound creates a GetZonesConnectionNotFound with default headers values
func NewGetZonesConnectionNotFound() *GetZonesConnectionNotFound {
	return &GetZonesConnectionNotFound{}
}

/* GetZonesConnectionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetZonesConnectionNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetZonesConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /get-zones-connection][%d] getZonesConnectionNotFound  %+v", 404, o.Payload)
}
func (o *GetZonesConnectionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetZonesConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetZonesConnectionInternalServerError creates a GetZonesConnectionInternalServerError with default headers values
func NewGetZonesConnectionInternalServerError() *GetZonesConnectionInternalServerError {
	return &GetZonesConnectionInternalServerError{}
}

/* GetZonesConnectionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetZonesConnectionInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetZonesConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-zones-connection][%d] getZonesConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetZonesConnectionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetZonesConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
