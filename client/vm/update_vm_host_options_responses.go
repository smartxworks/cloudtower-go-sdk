// Code generated by go-swagger; DO NOT EDIT.

package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// UpdateVMHostOptionsReader is a Reader for the UpdateVMHostOptions structure.
type UpdateVMHostOptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVMHostOptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateVMHostOptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 304:
		result := NewUpdateVMHostOptionsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 400:
		result := NewUpdateVMHostOptionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewUpdateVMHostOptionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewUpdateVMHostOptionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateVMHostOptionsOK creates a UpdateVMHostOptionsOK with default headers values
func NewUpdateVMHostOptionsOK() *UpdateVMHostOptionsOK {
	return &UpdateVMHostOptionsOK{}
}

/* UpdateVMHostOptionsOK describes a response with status code 200, with default header values.

UpdateVMHostOptionsOK update Vm host options o k
*/
type UpdateVMHostOptionsOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskVM
}

func (o *UpdateVMHostOptionsOK) Error() string {
	return fmt.Sprintf("[POST /update-vm-host-options][%d] updateVmHostOptionsOK  %+v", 200, o.Payload)
}
func (o *UpdateVMHostOptionsOK) GetPayload() []*models.WithTaskVM {
	return o.Payload
}

func (o *UpdateVMHostOptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateVMHostOptionsNotModified creates a UpdateVMHostOptionsNotModified with default headers values
func NewUpdateVMHostOptionsNotModified() *UpdateVMHostOptionsNotModified {
	return &UpdateVMHostOptionsNotModified{}
}

/* UpdateVMHostOptionsNotModified describes a response with status code 304, with default header values.

Not modified
*/
type UpdateVMHostOptionsNotModified struct {
}

func (o *UpdateVMHostOptionsNotModified) Error() string {
	return fmt.Sprintf("[POST /update-vm-host-options][%d] updateVmHostOptionsNotModified ", 304)
}

func (o *UpdateVMHostOptionsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateVMHostOptionsBadRequest creates a UpdateVMHostOptionsBadRequest with default headers values
func NewUpdateVMHostOptionsBadRequest() *UpdateVMHostOptionsBadRequest {
	return &UpdateVMHostOptionsBadRequest{}
}

/* UpdateVMHostOptionsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateVMHostOptionsBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateVMHostOptionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-vm-host-options][%d] updateVmHostOptionsBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateVMHostOptionsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateVMHostOptionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateVMHostOptionsNotFound creates a UpdateVMHostOptionsNotFound with default headers values
func NewUpdateVMHostOptionsNotFound() *UpdateVMHostOptionsNotFound {
	return &UpdateVMHostOptionsNotFound{}
}

/* UpdateVMHostOptionsNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateVMHostOptionsNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateVMHostOptionsNotFound) Error() string {
	return fmt.Sprintf("[POST /update-vm-host-options][%d] updateVmHostOptionsNotFound  %+v", 404, o.Payload)
}
func (o *UpdateVMHostOptionsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateVMHostOptionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateVMHostOptionsInternalServerError creates a UpdateVMHostOptionsInternalServerError with default headers values
func NewUpdateVMHostOptionsInternalServerError() *UpdateVMHostOptionsInternalServerError {
	return &UpdateVMHostOptionsInternalServerError{}
}

/* UpdateVMHostOptionsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UpdateVMHostOptionsInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateVMHostOptionsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /update-vm-host-options][%d] updateVmHostOptionsInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateVMHostOptionsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateVMHostOptionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
