// Code generated by go-swagger; DO NOT EDIT.

package vds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetVdsesReader is a Reader for the GetVdses structure.
type GetVdsesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVdsesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVdsesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewGetVdsesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewGetVdsesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewGetVdsesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVdsesOK creates a GetVdsesOK with default headers values
func NewGetVdsesOK() *GetVdsesOK {
	return &GetVdsesOK{}
}

/* GetVdsesOK describes a response with status code 200, with default header values.

GetVdsesOK get vdses o k
*/
type GetVdsesOK struct {
	XTowerRequestID string

	Payload []*models.Vds
}

func (o *GetVdsesOK) Error() string {
	return fmt.Sprintf("[POST /get-vdses][%d] getVdsesOK  %+v", 200, o.Payload)
}
func (o *GetVdsesOK) GetPayload() []*models.Vds {
	return o.Payload
}

func (o *GetVdsesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVdsesBadRequest creates a GetVdsesBadRequest with default headers values
func NewGetVdsesBadRequest() *GetVdsesBadRequest {
	return &GetVdsesBadRequest{}
}

/* GetVdsesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetVdsesBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetVdsesBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-vdses][%d] getVdsesBadRequest  %+v", 400, o.Payload)
}
func (o *GetVdsesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVdsesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetVdsesNotFound creates a GetVdsesNotFound with default headers values
func NewGetVdsesNotFound() *GetVdsesNotFound {
	return &GetVdsesNotFound{}
}

/* GetVdsesNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetVdsesNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetVdsesNotFound) Error() string {
	return fmt.Sprintf("[POST /get-vdses][%d] getVdsesNotFound  %+v", 404, o.Payload)
}
func (o *GetVdsesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVdsesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetVdsesInternalServerError creates a GetVdsesInternalServerError with default headers values
func NewGetVdsesInternalServerError() *GetVdsesInternalServerError {
	return &GetVdsesInternalServerError{}
}

/* GetVdsesInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetVdsesInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetVdsesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-vdses][%d] getVdsesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetVdsesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVdsesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
