// Code generated by go-swagger; DO NOT EDIT.

package vm_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// ConvertVMTemplateFromVMReader is a Reader for the ConvertVMTemplateFromVM structure.
type ConvertVMTemplateFromVMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConvertVMTemplateFromVMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConvertVMTemplateFromVMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewConvertVMTemplateFromVMBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewConvertVMTemplateFromVMNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewConvertVMTemplateFromVMInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewConvertVMTemplateFromVMOK creates a ConvertVMTemplateFromVMOK with default headers values
func NewConvertVMTemplateFromVMOK() *ConvertVMTemplateFromVMOK {
	return &ConvertVMTemplateFromVMOK{}
}

/* ConvertVMTemplateFromVMOK describes a response with status code 200, with default header values.

ConvertVMTemplateFromVMOK convert Vm template from Vm o k
*/
type ConvertVMTemplateFromVMOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskVMTemplate
}

func (o *ConvertVMTemplateFromVMOK) Error() string {
	return fmt.Sprintf("[POST /convert-vm-template-from-vm][%d] convertVmTemplateFromVmOK  %+v", 200, o.Payload)
}
func (o *ConvertVMTemplateFromVMOK) GetPayload() []*models.WithTaskVMTemplate {
	return o.Payload
}

func (o *ConvertVMTemplateFromVMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewConvertVMTemplateFromVMBadRequest creates a ConvertVMTemplateFromVMBadRequest with default headers values
func NewConvertVMTemplateFromVMBadRequest() *ConvertVMTemplateFromVMBadRequest {
	return &ConvertVMTemplateFromVMBadRequest{}
}

/* ConvertVMTemplateFromVMBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ConvertVMTemplateFromVMBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *ConvertVMTemplateFromVMBadRequest) Error() string {
	return fmt.Sprintf("[POST /convert-vm-template-from-vm][%d] convertVmTemplateFromVmBadRequest  %+v", 400, o.Payload)
}
func (o *ConvertVMTemplateFromVMBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ConvertVMTemplateFromVMBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewConvertVMTemplateFromVMNotFound creates a ConvertVMTemplateFromVMNotFound with default headers values
func NewConvertVMTemplateFromVMNotFound() *ConvertVMTemplateFromVMNotFound {
	return &ConvertVMTemplateFromVMNotFound{}
}

/* ConvertVMTemplateFromVMNotFound describes a response with status code 404, with default header values.

Not found
*/
type ConvertVMTemplateFromVMNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *ConvertVMTemplateFromVMNotFound) Error() string {
	return fmt.Sprintf("[POST /convert-vm-template-from-vm][%d] convertVmTemplateFromVmNotFound  %+v", 404, o.Payload)
}
func (o *ConvertVMTemplateFromVMNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ConvertVMTemplateFromVMNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewConvertVMTemplateFromVMInternalServerError creates a ConvertVMTemplateFromVMInternalServerError with default headers values
func NewConvertVMTemplateFromVMInternalServerError() *ConvertVMTemplateFromVMInternalServerError {
	return &ConvertVMTemplateFromVMInternalServerError{}
}

/* ConvertVMTemplateFromVMInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type ConvertVMTemplateFromVMInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *ConvertVMTemplateFromVMInternalServerError) Error() string {
	return fmt.Sprintf("[POST /convert-vm-template-from-vm][%d] convertVmTemplateFromVmInternalServerError  %+v", 500, o.Payload)
}
func (o *ConvertVMTemplateFromVMInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ConvertVMTemplateFromVMInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
