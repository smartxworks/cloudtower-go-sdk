// Code generated by go-swagger; DO NOT EDIT.

package brick_topo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetBrickTopoesConnectionReader is a Reader for the GetBrickTopoesConnection structure.
type GetBrickTopoesConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBrickTopoesConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBrickTopoesConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetBrickTopoesConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBrickTopoesConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetBrickTopoesConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBrickTopoesConnectionOK creates a GetBrickTopoesConnectionOK with default headers values
func NewGetBrickTopoesConnectionOK() *GetBrickTopoesConnectionOK {
	return &GetBrickTopoesConnectionOK{}
}

/* GetBrickTopoesConnectionOK describes a response with status code 200, with default header values.

GetBrickTopoesConnectionOK get brick topoes connection o k
*/
type GetBrickTopoesConnectionOK struct {
	XTowerRequestID string

	Payload *models.BrickTopoConnection
}

func (o *GetBrickTopoesConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-brick-topoes-connection][%d] getBrickTopoesConnectionOK  %+v", 200, o.Payload)
}
func (o *GetBrickTopoesConnectionOK) GetPayload() *models.BrickTopoConnection {
	return o.Payload
}

func (o *GetBrickTopoesConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.BrickTopoConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBrickTopoesConnectionBadRequest creates a GetBrickTopoesConnectionBadRequest with default headers values
func NewGetBrickTopoesConnectionBadRequest() *GetBrickTopoesConnectionBadRequest {
	return &GetBrickTopoesConnectionBadRequest{}
}

/* GetBrickTopoesConnectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetBrickTopoesConnectionBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetBrickTopoesConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-brick-topoes-connection][%d] getBrickTopoesConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetBrickTopoesConnectionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetBrickTopoesConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetBrickTopoesConnectionNotFound creates a GetBrickTopoesConnectionNotFound with default headers values
func NewGetBrickTopoesConnectionNotFound() *GetBrickTopoesConnectionNotFound {
	return &GetBrickTopoesConnectionNotFound{}
}

/* GetBrickTopoesConnectionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetBrickTopoesConnectionNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetBrickTopoesConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /get-brick-topoes-connection][%d] getBrickTopoesConnectionNotFound  %+v", 404, o.Payload)
}
func (o *GetBrickTopoesConnectionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetBrickTopoesConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetBrickTopoesConnectionInternalServerError creates a GetBrickTopoesConnectionInternalServerError with default headers values
func NewGetBrickTopoesConnectionInternalServerError() *GetBrickTopoesConnectionInternalServerError {
	return &GetBrickTopoesConnectionInternalServerError{}
}

/* GetBrickTopoesConnectionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetBrickTopoesConnectionInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetBrickTopoesConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-brick-topoes-connection][%d] getBrickTopoesConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetBrickTopoesConnectionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetBrickTopoesConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
