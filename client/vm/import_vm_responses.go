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

// ImportVMReader is a Reader for the ImportVM structure.
type ImportVMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportVMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImportVMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewImportVMBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewImportVMNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewImportVMInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImportVMOK creates a ImportVMOK with default headers values
func NewImportVMOK() *ImportVMOK {
	return &ImportVMOK{}
}

/* ImportVMOK describes a response with status code 200, with default header values.

ImportVMOK import Vm o k
*/
type ImportVMOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskVM
}

func (o *ImportVMOK) Error() string {
	return fmt.Sprintf("[POST /import-vm][%d] importVmOK  %+v", 200, o.Payload)
}
func (o *ImportVMOK) GetPayload() []*models.WithTaskVM {
	return o.Payload
}

func (o *ImportVMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewImportVMBadRequest creates a ImportVMBadRequest with default headers values
func NewImportVMBadRequest() *ImportVMBadRequest {
	return &ImportVMBadRequest{}
}

/* ImportVMBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ImportVMBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *ImportVMBadRequest) Error() string {
	return fmt.Sprintf("[POST /import-vm][%d] importVmBadRequest  %+v", 400, o.Payload)
}
func (o *ImportVMBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ImportVMBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewImportVMNotFound creates a ImportVMNotFound with default headers values
func NewImportVMNotFound() *ImportVMNotFound {
	return &ImportVMNotFound{}
}

/* ImportVMNotFound describes a response with status code 404, with default header values.

Not found
*/
type ImportVMNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *ImportVMNotFound) Error() string {
	return fmt.Sprintf("[POST /import-vm][%d] importVmNotFound  %+v", 404, o.Payload)
}
func (o *ImportVMNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ImportVMNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewImportVMInternalServerError creates a ImportVMInternalServerError with default headers values
func NewImportVMInternalServerError() *ImportVMInternalServerError {
	return &ImportVMInternalServerError{}
}

/* ImportVMInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type ImportVMInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *ImportVMInternalServerError) Error() string {
	return fmt.Sprintf("[POST /import-vm][%d] importVmInternalServerError  %+v", 500, o.Payload)
}
func (o *ImportVMInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ImportVMInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
