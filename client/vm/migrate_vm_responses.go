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

// MigrateVMReader is a Reader for the MigrateVM structure.
type MigrateVMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MigrateVMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMigrateVMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewMigrateVMBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewMigrateVMNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewMigrateVMInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMigrateVMOK creates a MigrateVMOK with default headers values
func NewMigrateVMOK() *MigrateVMOK {
	return &MigrateVMOK{}
}

/* MigrateVMOK describes a response with status code 200, with default header values.

MigrateVMOK migrate Vm o k
*/
type MigrateVMOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskVM
}

func (o *MigrateVMOK) Error() string {
	return fmt.Sprintf("[POST /migrate-vm][%d] migrateVmOK  %+v", 200, o.Payload)
}
func (o *MigrateVMOK) GetPayload() []*models.WithTaskVM {
	return o.Payload
}

func (o *MigrateVMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewMigrateVMBadRequest creates a MigrateVMBadRequest with default headers values
func NewMigrateVMBadRequest() *MigrateVMBadRequest {
	return &MigrateVMBadRequest{}
}

/* MigrateVMBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type MigrateVMBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *MigrateVMBadRequest) Error() string {
	return fmt.Sprintf("[POST /migrate-vm][%d] migrateVmBadRequest  %+v", 400, o.Payload)
}
func (o *MigrateVMBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *MigrateVMBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewMigrateVMNotFound creates a MigrateVMNotFound with default headers values
func NewMigrateVMNotFound() *MigrateVMNotFound {
	return &MigrateVMNotFound{}
}

/* MigrateVMNotFound describes a response with status code 404, with default header values.

Not found
*/
type MigrateVMNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *MigrateVMNotFound) Error() string {
	return fmt.Sprintf("[POST /migrate-vm][%d] migrateVmNotFound  %+v", 404, o.Payload)
}
func (o *MigrateVMNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *MigrateVMNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewMigrateVMInternalServerError creates a MigrateVMInternalServerError with default headers values
func NewMigrateVMInternalServerError() *MigrateVMInternalServerError {
	return &MigrateVMInternalServerError{}
}

/* MigrateVMInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type MigrateVMInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *MigrateVMInternalServerError) Error() string {
	return fmt.Sprintf("[POST /migrate-vm][%d] migrateVmInternalServerError  %+v", 500, o.Payload)
}
func (o *MigrateVMInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *MigrateVMInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
