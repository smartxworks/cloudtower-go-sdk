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

// ExpandVMDiskReader is a Reader for the ExpandVMDisk structure.
type ExpandVMDiskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExpandVMDiskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExpandVMDiskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewExpandVMDiskBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewExpandVMDiskNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewExpandVMDiskInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExpandVMDiskOK creates a ExpandVMDiskOK with default headers values
func NewExpandVMDiskOK() *ExpandVMDiskOK {
	return &ExpandVMDiskOK{}
}

/* ExpandVMDiskOK describes a response with status code 200, with default header values.

ExpandVMDiskOK expand Vm disk o k
*/
type ExpandVMDiskOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskVM
}

func (o *ExpandVMDiskOK) Error() string {
	return fmt.Sprintf("[POST /expand-vm-disk][%d] expandVmDiskOK  %+v", 200, o.Payload)
}
func (o *ExpandVMDiskOK) GetPayload() []*models.WithTaskVM {
	return o.Payload
}

func (o *ExpandVMDiskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewExpandVMDiskBadRequest creates a ExpandVMDiskBadRequest with default headers values
func NewExpandVMDiskBadRequest() *ExpandVMDiskBadRequest {
	return &ExpandVMDiskBadRequest{}
}

/* ExpandVMDiskBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ExpandVMDiskBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *ExpandVMDiskBadRequest) Error() string {
	return fmt.Sprintf("[POST /expand-vm-disk][%d] expandVmDiskBadRequest  %+v", 400, o.Payload)
}
func (o *ExpandVMDiskBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ExpandVMDiskBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewExpandVMDiskNotFound creates a ExpandVMDiskNotFound with default headers values
func NewExpandVMDiskNotFound() *ExpandVMDiskNotFound {
	return &ExpandVMDiskNotFound{}
}

/* ExpandVMDiskNotFound describes a response with status code 404, with default header values.

Not found
*/
type ExpandVMDiskNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *ExpandVMDiskNotFound) Error() string {
	return fmt.Sprintf("[POST /expand-vm-disk][%d] expandVmDiskNotFound  %+v", 404, o.Payload)
}
func (o *ExpandVMDiskNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ExpandVMDiskNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewExpandVMDiskInternalServerError creates a ExpandVMDiskInternalServerError with default headers values
func NewExpandVMDiskInternalServerError() *ExpandVMDiskInternalServerError {
	return &ExpandVMDiskInternalServerError{}
}

/* ExpandVMDiskInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type ExpandVMDiskInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *ExpandVMDiskInternalServerError) Error() string {
	return fmt.Sprintf("[POST /expand-vm-disk][%d] expandVmDiskInternalServerError  %+v", 500, o.Payload)
}
func (o *ExpandVMDiskInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ExpandVMDiskInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
