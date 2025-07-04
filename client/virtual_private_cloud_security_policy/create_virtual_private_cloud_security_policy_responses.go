// Code generated by go-swagger; DO NOT EDIT.

package virtual_private_cloud_security_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// CreateVirtualPrivateCloudSecurityPolicyReader is a Reader for the CreateVirtualPrivateCloudSecurityPolicy structure.
type CreateVirtualPrivateCloudSecurityPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVirtualPrivateCloudSecurityPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateVirtualPrivateCloudSecurityPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewCreateVirtualPrivateCloudSecurityPolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewCreateVirtualPrivateCloudSecurityPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewCreateVirtualPrivateCloudSecurityPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateVirtualPrivateCloudSecurityPolicyOK creates a CreateVirtualPrivateCloudSecurityPolicyOK with default headers values
func NewCreateVirtualPrivateCloudSecurityPolicyOK() *CreateVirtualPrivateCloudSecurityPolicyOK {
	return &CreateVirtualPrivateCloudSecurityPolicyOK{}
}

/* CreateVirtualPrivateCloudSecurityPolicyOK describes a response with status code 200, with default header values.

CreateVirtualPrivateCloudSecurityPolicyOK create virtual private cloud security policy o k
*/
type CreateVirtualPrivateCloudSecurityPolicyOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskVirtualPrivateCloudSecurityPolicy
}

func (o *CreateVirtualPrivateCloudSecurityPolicyOK) Error() string {
	return fmt.Sprintf("[POST /create-virtual-private-cloud-security-policy][%d] createVirtualPrivateCloudSecurityPolicyOK  %+v", 200, o.Payload)
}
func (o *CreateVirtualPrivateCloudSecurityPolicyOK) GetPayload() []*models.WithTaskVirtualPrivateCloudSecurityPolicy {
	return o.Payload
}

func (o *CreateVirtualPrivateCloudSecurityPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateVirtualPrivateCloudSecurityPolicyBadRequest creates a CreateVirtualPrivateCloudSecurityPolicyBadRequest with default headers values
func NewCreateVirtualPrivateCloudSecurityPolicyBadRequest() *CreateVirtualPrivateCloudSecurityPolicyBadRequest {
	return &CreateVirtualPrivateCloudSecurityPolicyBadRequest{}
}

/* CreateVirtualPrivateCloudSecurityPolicyBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateVirtualPrivateCloudSecurityPolicyBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateVirtualPrivateCloudSecurityPolicyBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-virtual-private-cloud-security-policy][%d] createVirtualPrivateCloudSecurityPolicyBadRequest  %+v", 400, o.Payload)
}
func (o *CreateVirtualPrivateCloudSecurityPolicyBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateVirtualPrivateCloudSecurityPolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateVirtualPrivateCloudSecurityPolicyNotFound creates a CreateVirtualPrivateCloudSecurityPolicyNotFound with default headers values
func NewCreateVirtualPrivateCloudSecurityPolicyNotFound() *CreateVirtualPrivateCloudSecurityPolicyNotFound {
	return &CreateVirtualPrivateCloudSecurityPolicyNotFound{}
}

/* CreateVirtualPrivateCloudSecurityPolicyNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateVirtualPrivateCloudSecurityPolicyNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateVirtualPrivateCloudSecurityPolicyNotFound) Error() string {
	return fmt.Sprintf("[POST /create-virtual-private-cloud-security-policy][%d] createVirtualPrivateCloudSecurityPolicyNotFound  %+v", 404, o.Payload)
}
func (o *CreateVirtualPrivateCloudSecurityPolicyNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateVirtualPrivateCloudSecurityPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateVirtualPrivateCloudSecurityPolicyInternalServerError creates a CreateVirtualPrivateCloudSecurityPolicyInternalServerError with default headers values
func NewCreateVirtualPrivateCloudSecurityPolicyInternalServerError() *CreateVirtualPrivateCloudSecurityPolicyInternalServerError {
	return &CreateVirtualPrivateCloudSecurityPolicyInternalServerError{}
}

/* CreateVirtualPrivateCloudSecurityPolicyInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CreateVirtualPrivateCloudSecurityPolicyInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateVirtualPrivateCloudSecurityPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /create-virtual-private-cloud-security-policy][%d] createVirtualPrivateCloudSecurityPolicyInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateVirtualPrivateCloudSecurityPolicyInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateVirtualPrivateCloudSecurityPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
