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

// UpdateNvmfSubsystemReader is a Reader for the UpdateNvmfSubsystem structure.
type UpdateNvmfSubsystemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNvmfSubsystemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNvmfSubsystemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewUpdateNvmfSubsystemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewUpdateNvmfSubsystemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewUpdateNvmfSubsystemInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNvmfSubsystemOK creates a UpdateNvmfSubsystemOK with default headers values
func NewUpdateNvmfSubsystemOK() *UpdateNvmfSubsystemOK {
	return &UpdateNvmfSubsystemOK{}
}

/* UpdateNvmfSubsystemOK describes a response with status code 200, with default header values.

UpdateNvmfSubsystemOK update nvmf subsystem o k
*/
type UpdateNvmfSubsystemOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskNvmfSubsystem
}

func (o *UpdateNvmfSubsystemOK) Error() string {
	return fmt.Sprintf("[POST /update-nvmf-subsystem][%d] updateNvmfSubsystemOK  %+v", 200, o.Payload)
}
func (o *UpdateNvmfSubsystemOK) GetPayload() []*models.WithTaskNvmfSubsystem {
	return o.Payload
}

func (o *UpdateNvmfSubsystemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateNvmfSubsystemBadRequest creates a UpdateNvmfSubsystemBadRequest with default headers values
func NewUpdateNvmfSubsystemBadRequest() *UpdateNvmfSubsystemBadRequest {
	return &UpdateNvmfSubsystemBadRequest{}
}

/* UpdateNvmfSubsystemBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateNvmfSubsystemBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateNvmfSubsystemBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-nvmf-subsystem][%d] updateNvmfSubsystemBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateNvmfSubsystemBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateNvmfSubsystemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateNvmfSubsystemNotFound creates a UpdateNvmfSubsystemNotFound with default headers values
func NewUpdateNvmfSubsystemNotFound() *UpdateNvmfSubsystemNotFound {
	return &UpdateNvmfSubsystemNotFound{}
}

/* UpdateNvmfSubsystemNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateNvmfSubsystemNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateNvmfSubsystemNotFound) Error() string {
	return fmt.Sprintf("[POST /update-nvmf-subsystem][%d] updateNvmfSubsystemNotFound  %+v", 404, o.Payload)
}
func (o *UpdateNvmfSubsystemNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateNvmfSubsystemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateNvmfSubsystemInternalServerError creates a UpdateNvmfSubsystemInternalServerError with default headers values
func NewUpdateNvmfSubsystemInternalServerError() *UpdateNvmfSubsystemInternalServerError {
	return &UpdateNvmfSubsystemInternalServerError{}
}

/* UpdateNvmfSubsystemInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UpdateNvmfSubsystemInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateNvmfSubsystemInternalServerError) Error() string {
	return fmt.Sprintf("[POST /update-nvmf-subsystem][%d] updateNvmfSubsystemInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateNvmfSubsystemInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateNvmfSubsystemInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
