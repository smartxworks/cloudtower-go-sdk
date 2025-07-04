// Code generated by go-swagger; DO NOT EDIT.

package zone_topo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetZoneTopoesConnectionReader is a Reader for the GetZoneTopoesConnection structure.
type GetZoneTopoesConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetZoneTopoesConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetZoneTopoesConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewGetZoneTopoesConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewGetZoneTopoesConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewGetZoneTopoesConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetZoneTopoesConnectionOK creates a GetZoneTopoesConnectionOK with default headers values
func NewGetZoneTopoesConnectionOK() *GetZoneTopoesConnectionOK {
	return &GetZoneTopoesConnectionOK{}
}

/* GetZoneTopoesConnectionOK describes a response with status code 200, with default header values.

GetZoneTopoesConnectionOK get zone topoes connection o k
*/
type GetZoneTopoesConnectionOK struct {
	XTowerRequestID string

	Payload *models.ZoneTopoConnection
}

func (o *GetZoneTopoesConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-zone-topoes-connection][%d] getZoneTopoesConnectionOK  %+v", 200, o.Payload)
}
func (o *GetZoneTopoesConnectionOK) GetPayload() *models.ZoneTopoConnection {
	return o.Payload
}

func (o *GetZoneTopoesConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.ZoneTopoConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetZoneTopoesConnectionBadRequest creates a GetZoneTopoesConnectionBadRequest with default headers values
func NewGetZoneTopoesConnectionBadRequest() *GetZoneTopoesConnectionBadRequest {
	return &GetZoneTopoesConnectionBadRequest{}
}

/* GetZoneTopoesConnectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetZoneTopoesConnectionBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetZoneTopoesConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-zone-topoes-connection][%d] getZoneTopoesConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetZoneTopoesConnectionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetZoneTopoesConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetZoneTopoesConnectionNotFound creates a GetZoneTopoesConnectionNotFound with default headers values
func NewGetZoneTopoesConnectionNotFound() *GetZoneTopoesConnectionNotFound {
	return &GetZoneTopoesConnectionNotFound{}
}

/* GetZoneTopoesConnectionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetZoneTopoesConnectionNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetZoneTopoesConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /get-zone-topoes-connection][%d] getZoneTopoesConnectionNotFound  %+v", 404, o.Payload)
}
func (o *GetZoneTopoesConnectionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetZoneTopoesConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetZoneTopoesConnectionInternalServerError creates a GetZoneTopoesConnectionInternalServerError with default headers values
func NewGetZoneTopoesConnectionInternalServerError() *GetZoneTopoesConnectionInternalServerError {
	return &GetZoneTopoesConnectionInternalServerError{}
}

/* GetZoneTopoesConnectionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetZoneTopoesConnectionInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetZoneTopoesConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-zone-topoes-connection][%d] getZoneTopoesConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetZoneTopoesConnectionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetZoneTopoesConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
