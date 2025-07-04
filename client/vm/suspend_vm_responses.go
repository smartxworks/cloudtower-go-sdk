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

// SuspendVMReader is a Reader for the SuspendVM structure.
type SuspendVMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SuspendVMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSuspendVMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewSuspendVMBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewSuspendVMNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewSuspendVMInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSuspendVMOK creates a SuspendVMOK with default headers values
func NewSuspendVMOK() *SuspendVMOK {
	return &SuspendVMOK{}
}

/* SuspendVMOK describes a response with status code 200, with default header values.

SuspendVMOK suspend Vm o k
*/
type SuspendVMOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskVM
}

func (o *SuspendVMOK) Error() string {
	return fmt.Sprintf("[POST /suspend-vm][%d] suspendVmOK  %+v", 200, o.Payload)
}
func (o *SuspendVMOK) GetPayload() []*models.WithTaskVM {
	return o.Payload
}

func (o *SuspendVMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSuspendVMBadRequest creates a SuspendVMBadRequest with default headers values
func NewSuspendVMBadRequest() *SuspendVMBadRequest {
	return &SuspendVMBadRequest{}
}

/* SuspendVMBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type SuspendVMBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *SuspendVMBadRequest) Error() string {
	return fmt.Sprintf("[POST /suspend-vm][%d] suspendVmBadRequest  %+v", 400, o.Payload)
}
func (o *SuspendVMBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *SuspendVMBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSuspendVMNotFound creates a SuspendVMNotFound with default headers values
func NewSuspendVMNotFound() *SuspendVMNotFound {
	return &SuspendVMNotFound{}
}

/* SuspendVMNotFound describes a response with status code 404, with default header values.

Not found
*/
type SuspendVMNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *SuspendVMNotFound) Error() string {
	return fmt.Sprintf("[POST /suspend-vm][%d] suspendVmNotFound  %+v", 404, o.Payload)
}
func (o *SuspendVMNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *SuspendVMNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSuspendVMInternalServerError creates a SuspendVMInternalServerError with default headers values
func NewSuspendVMInternalServerError() *SuspendVMInternalServerError {
	return &SuspendVMInternalServerError{}
}

/* SuspendVMInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type SuspendVMInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *SuspendVMInternalServerError) Error() string {
	return fmt.Sprintf("[POST /suspend-vm][%d] suspendVmInternalServerError  %+v", 500, o.Payload)
}
func (o *SuspendVMInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *SuspendVMInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
