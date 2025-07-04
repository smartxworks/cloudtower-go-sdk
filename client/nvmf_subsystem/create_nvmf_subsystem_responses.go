// Code generated by go-swagger; DO NOT EDIT.

package nvmf_subsystem

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// CreateNvmfSubsystemReader is a Reader for the CreateNvmfSubsystem structure.
type CreateNvmfSubsystemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNvmfSubsystemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateNvmfSubsystemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewCreateNvmfSubsystemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewCreateNvmfSubsystemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewCreateNvmfSubsystemInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateNvmfSubsystemOK creates a CreateNvmfSubsystemOK with default headers values
func NewCreateNvmfSubsystemOK() *CreateNvmfSubsystemOK {
	return &CreateNvmfSubsystemOK{}
}

/* CreateNvmfSubsystemOK describes a response with status code 200, with default header values.

CreateNvmfSubsystemOK create nvmf subsystem o k
*/
type CreateNvmfSubsystemOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskNvmfSubsystem
}

func (o *CreateNvmfSubsystemOK) Error() string {
	return fmt.Sprintf("[POST /create-nvmf-subsystem][%d] createNvmfSubsystemOK  %+v", 200, o.Payload)
}
func (o *CreateNvmfSubsystemOK) GetPayload() []*models.WithTaskNvmfSubsystem {
	return o.Payload
}

func (o *CreateNvmfSubsystemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateNvmfSubsystemBadRequest creates a CreateNvmfSubsystemBadRequest with default headers values
func NewCreateNvmfSubsystemBadRequest() *CreateNvmfSubsystemBadRequest {
	return &CreateNvmfSubsystemBadRequest{}
}

/* CreateNvmfSubsystemBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateNvmfSubsystemBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateNvmfSubsystemBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-nvmf-subsystem][%d] createNvmfSubsystemBadRequest  %+v", 400, o.Payload)
}
func (o *CreateNvmfSubsystemBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateNvmfSubsystemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateNvmfSubsystemNotFound creates a CreateNvmfSubsystemNotFound with default headers values
func NewCreateNvmfSubsystemNotFound() *CreateNvmfSubsystemNotFound {
	return &CreateNvmfSubsystemNotFound{}
}

/* CreateNvmfSubsystemNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateNvmfSubsystemNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateNvmfSubsystemNotFound) Error() string {
	return fmt.Sprintf("[POST /create-nvmf-subsystem][%d] createNvmfSubsystemNotFound  %+v", 404, o.Payload)
}
func (o *CreateNvmfSubsystemNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateNvmfSubsystemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateNvmfSubsystemInternalServerError creates a CreateNvmfSubsystemInternalServerError with default headers values
func NewCreateNvmfSubsystemInternalServerError() *CreateNvmfSubsystemInternalServerError {
	return &CreateNvmfSubsystemInternalServerError{}
}

/* CreateNvmfSubsystemInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CreateNvmfSubsystemInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateNvmfSubsystemInternalServerError) Error() string {
	return fmt.Sprintf("[POST /create-nvmf-subsystem][%d] createNvmfSubsystemInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateNvmfSubsystemInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateNvmfSubsystemInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
