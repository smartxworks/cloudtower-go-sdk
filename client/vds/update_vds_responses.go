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

// UpdateVdsReader is a Reader for the UpdateVds structure.
type UpdateVdsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVdsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateVdsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewUpdateVdsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewUpdateVdsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewUpdateVdsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateVdsOK creates a UpdateVdsOK with default headers values
func NewUpdateVdsOK() *UpdateVdsOK {
	return &UpdateVdsOK{}
}

/* UpdateVdsOK describes a response with status code 200, with default header values.

UpdateVdsOK update vds o k
*/
type UpdateVdsOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskVds
}

func (o *UpdateVdsOK) Error() string {
	return fmt.Sprintf("[POST /update-vds][%d] updateVdsOK  %+v", 200, o.Payload)
}
func (o *UpdateVdsOK) GetPayload() []*models.WithTaskVds {
	return o.Payload
}

func (o *UpdateVdsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateVdsBadRequest creates a UpdateVdsBadRequest with default headers values
func NewUpdateVdsBadRequest() *UpdateVdsBadRequest {
	return &UpdateVdsBadRequest{}
}

/* UpdateVdsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateVdsBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateVdsBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-vds][%d] updateVdsBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateVdsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateVdsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateVdsNotFound creates a UpdateVdsNotFound with default headers values
func NewUpdateVdsNotFound() *UpdateVdsNotFound {
	return &UpdateVdsNotFound{}
}

/* UpdateVdsNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateVdsNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateVdsNotFound) Error() string {
	return fmt.Sprintf("[POST /update-vds][%d] updateVdsNotFound  %+v", 404, o.Payload)
}
func (o *UpdateVdsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateVdsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateVdsInternalServerError creates a UpdateVdsInternalServerError with default headers values
func NewUpdateVdsInternalServerError() *UpdateVdsInternalServerError {
	return &UpdateVdsInternalServerError{}
}

/* UpdateVdsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UpdateVdsInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateVdsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /update-vds][%d] updateVdsInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateVdsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateVdsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
