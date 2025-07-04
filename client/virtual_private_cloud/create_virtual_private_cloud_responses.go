// Code generated by go-swagger; DO NOT EDIT.

package virtual_private_cloud

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// CreateVirtualPrivateCloudReader is a Reader for the CreateVirtualPrivateCloud structure.
type CreateVirtualPrivateCloudReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVirtualPrivateCloudReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateVirtualPrivateCloudOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewCreateVirtualPrivateCloudBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewCreateVirtualPrivateCloudNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewCreateVirtualPrivateCloudInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateVirtualPrivateCloudOK creates a CreateVirtualPrivateCloudOK with default headers values
func NewCreateVirtualPrivateCloudOK() *CreateVirtualPrivateCloudOK {
	return &CreateVirtualPrivateCloudOK{}
}

/* CreateVirtualPrivateCloudOK describes a response with status code 200, with default header values.

CreateVirtualPrivateCloudOK create virtual private cloud o k
*/
type CreateVirtualPrivateCloudOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskVirtualPrivateCloud
}

func (o *CreateVirtualPrivateCloudOK) Error() string {
	return fmt.Sprintf("[POST /create-virtual-private-cloud][%d] createVirtualPrivateCloudOK  %+v", 200, o.Payload)
}
func (o *CreateVirtualPrivateCloudOK) GetPayload() []*models.WithTaskVirtualPrivateCloud {
	return o.Payload
}

func (o *CreateVirtualPrivateCloudOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateVirtualPrivateCloudBadRequest creates a CreateVirtualPrivateCloudBadRequest with default headers values
func NewCreateVirtualPrivateCloudBadRequest() *CreateVirtualPrivateCloudBadRequest {
	return &CreateVirtualPrivateCloudBadRequest{}
}

/* CreateVirtualPrivateCloudBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateVirtualPrivateCloudBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateVirtualPrivateCloudBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-virtual-private-cloud][%d] createVirtualPrivateCloudBadRequest  %+v", 400, o.Payload)
}
func (o *CreateVirtualPrivateCloudBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateVirtualPrivateCloudBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateVirtualPrivateCloudNotFound creates a CreateVirtualPrivateCloudNotFound with default headers values
func NewCreateVirtualPrivateCloudNotFound() *CreateVirtualPrivateCloudNotFound {
	return &CreateVirtualPrivateCloudNotFound{}
}

/* CreateVirtualPrivateCloudNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateVirtualPrivateCloudNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateVirtualPrivateCloudNotFound) Error() string {
	return fmt.Sprintf("[POST /create-virtual-private-cloud][%d] createVirtualPrivateCloudNotFound  %+v", 404, o.Payload)
}
func (o *CreateVirtualPrivateCloudNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateVirtualPrivateCloudNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateVirtualPrivateCloudInternalServerError creates a CreateVirtualPrivateCloudInternalServerError with default headers values
func NewCreateVirtualPrivateCloudInternalServerError() *CreateVirtualPrivateCloudInternalServerError {
	return &CreateVirtualPrivateCloudInternalServerError{}
}

/* CreateVirtualPrivateCloudInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CreateVirtualPrivateCloudInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateVirtualPrivateCloudInternalServerError) Error() string {
	return fmt.Sprintf("[POST /create-virtual-private-cloud][%d] createVirtualPrivateCloudInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateVirtualPrivateCloudInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateVirtualPrivateCloudInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
