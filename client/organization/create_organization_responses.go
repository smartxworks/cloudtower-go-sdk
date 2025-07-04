// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// CreateOrganizationReader is a Reader for the CreateOrganization structure.
type CreateOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewCreateOrganizationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewCreateOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewCreateOrganizationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateOrganizationOK creates a CreateOrganizationOK with default headers values
func NewCreateOrganizationOK() *CreateOrganizationOK {
	return &CreateOrganizationOK{}
}

/* CreateOrganizationOK describes a response with status code 200, with default header values.

CreateOrganizationOK create organization o k
*/
type CreateOrganizationOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskOrganization
}

func (o *CreateOrganizationOK) Error() string {
	return fmt.Sprintf("[POST /create-organization][%d] createOrganizationOK  %+v", 200, o.Payload)
}
func (o *CreateOrganizationOK) GetPayload() []*models.WithTaskOrganization {
	return o.Payload
}

func (o *CreateOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateOrganizationBadRequest creates a CreateOrganizationBadRequest with default headers values
func NewCreateOrganizationBadRequest() *CreateOrganizationBadRequest {
	return &CreateOrganizationBadRequest{}
}

/* CreateOrganizationBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateOrganizationBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateOrganizationBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-organization][%d] createOrganizationBadRequest  %+v", 400, o.Payload)
}
func (o *CreateOrganizationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateOrganizationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateOrganizationNotFound creates a CreateOrganizationNotFound with default headers values
func NewCreateOrganizationNotFound() *CreateOrganizationNotFound {
	return &CreateOrganizationNotFound{}
}

/* CreateOrganizationNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateOrganizationNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateOrganizationNotFound) Error() string {
	return fmt.Sprintf("[POST /create-organization][%d] createOrganizationNotFound  %+v", 404, o.Payload)
}
func (o *CreateOrganizationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateOrganizationInternalServerError creates a CreateOrganizationInternalServerError with default headers values
func NewCreateOrganizationInternalServerError() *CreateOrganizationInternalServerError {
	return &CreateOrganizationInternalServerError{}
}

/* CreateOrganizationInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CreateOrganizationInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateOrganizationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /create-organization][%d] createOrganizationInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateOrganizationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateOrganizationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
