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

// CreateVdsWithAccessVlanReader is a Reader for the CreateVdsWithAccessVlan structure.
type CreateVdsWithAccessVlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVdsWithAccessVlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateVdsWithAccessVlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewCreateVdsWithAccessVlanBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewCreateVdsWithAccessVlanNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewCreateVdsWithAccessVlanInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateVdsWithAccessVlanOK creates a CreateVdsWithAccessVlanOK with default headers values
func NewCreateVdsWithAccessVlanOK() *CreateVdsWithAccessVlanOK {
	return &CreateVdsWithAccessVlanOK{}
}

/* CreateVdsWithAccessVlanOK describes a response with status code 200, with default header values.

CreateVdsWithAccessVlanOK create vds with access vlan o k
*/
type CreateVdsWithAccessVlanOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskVds
}

func (o *CreateVdsWithAccessVlanOK) Error() string {
	return fmt.Sprintf("[POST /create-vds-with-access-vlan][%d] createVdsWithAccessVlanOK  %+v", 200, o.Payload)
}
func (o *CreateVdsWithAccessVlanOK) GetPayload() []*models.WithTaskVds {
	return o.Payload
}

func (o *CreateVdsWithAccessVlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateVdsWithAccessVlanBadRequest creates a CreateVdsWithAccessVlanBadRequest with default headers values
func NewCreateVdsWithAccessVlanBadRequest() *CreateVdsWithAccessVlanBadRequest {
	return &CreateVdsWithAccessVlanBadRequest{}
}

/* CreateVdsWithAccessVlanBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateVdsWithAccessVlanBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateVdsWithAccessVlanBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-vds-with-access-vlan][%d] createVdsWithAccessVlanBadRequest  %+v", 400, o.Payload)
}
func (o *CreateVdsWithAccessVlanBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateVdsWithAccessVlanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateVdsWithAccessVlanNotFound creates a CreateVdsWithAccessVlanNotFound with default headers values
func NewCreateVdsWithAccessVlanNotFound() *CreateVdsWithAccessVlanNotFound {
	return &CreateVdsWithAccessVlanNotFound{}
}

/* CreateVdsWithAccessVlanNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateVdsWithAccessVlanNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateVdsWithAccessVlanNotFound) Error() string {
	return fmt.Sprintf("[POST /create-vds-with-access-vlan][%d] createVdsWithAccessVlanNotFound  %+v", 404, o.Payload)
}
func (o *CreateVdsWithAccessVlanNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateVdsWithAccessVlanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateVdsWithAccessVlanInternalServerError creates a CreateVdsWithAccessVlanInternalServerError with default headers values
func NewCreateVdsWithAccessVlanInternalServerError() *CreateVdsWithAccessVlanInternalServerError {
	return &CreateVdsWithAccessVlanInternalServerError{}
}

/* CreateVdsWithAccessVlanInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CreateVdsWithAccessVlanInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateVdsWithAccessVlanInternalServerError) Error() string {
	return fmt.Sprintf("[POST /create-vds-with-access-vlan][%d] createVdsWithAccessVlanInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateVdsWithAccessVlanInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateVdsWithAccessVlanInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
