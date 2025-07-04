// Code generated by go-swagger; DO NOT EDIT.

package security_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// CreateSecurityPolicyReader is a Reader for the CreateSecurityPolicy structure.
type CreateSecurityPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSecurityPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateSecurityPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewCreateSecurityPolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewCreateSecurityPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewCreateSecurityPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateSecurityPolicyOK creates a CreateSecurityPolicyOK with default headers values
func NewCreateSecurityPolicyOK() *CreateSecurityPolicyOK {
	return &CreateSecurityPolicyOK{}
}

/* CreateSecurityPolicyOK describes a response with status code 200, with default header values.

CreateSecurityPolicyOK create security policy o k
*/
type CreateSecurityPolicyOK struct {
	XTowerRequestID string

	Payload *models.WithTaskSecurityPolicy
}

func (o *CreateSecurityPolicyOK) Error() string {
	return fmt.Sprintf("[POST /create-security-policy][%d] createSecurityPolicyOK  %+v", 200, o.Payload)
}
func (o *CreateSecurityPolicyOK) GetPayload() *models.WithTaskSecurityPolicy {
	return o.Payload
}

func (o *CreateSecurityPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.WithTaskSecurityPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSecurityPolicyBadRequest creates a CreateSecurityPolicyBadRequest with default headers values
func NewCreateSecurityPolicyBadRequest() *CreateSecurityPolicyBadRequest {
	return &CreateSecurityPolicyBadRequest{}
}

/* CreateSecurityPolicyBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateSecurityPolicyBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateSecurityPolicyBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-security-policy][%d] createSecurityPolicyBadRequest  %+v", 400, o.Payload)
}
func (o *CreateSecurityPolicyBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateSecurityPolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateSecurityPolicyNotFound creates a CreateSecurityPolicyNotFound with default headers values
func NewCreateSecurityPolicyNotFound() *CreateSecurityPolicyNotFound {
	return &CreateSecurityPolicyNotFound{}
}

/* CreateSecurityPolicyNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateSecurityPolicyNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateSecurityPolicyNotFound) Error() string {
	return fmt.Sprintf("[POST /create-security-policy][%d] createSecurityPolicyNotFound  %+v", 404, o.Payload)
}
func (o *CreateSecurityPolicyNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateSecurityPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateSecurityPolicyInternalServerError creates a CreateSecurityPolicyInternalServerError with default headers values
func NewCreateSecurityPolicyInternalServerError() *CreateSecurityPolicyInternalServerError {
	return &CreateSecurityPolicyInternalServerError{}
}

/* CreateSecurityPolicyInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CreateSecurityPolicyInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateSecurityPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /create-security-policy][%d] createSecurityPolicyInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateSecurityPolicyInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateSecurityPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
