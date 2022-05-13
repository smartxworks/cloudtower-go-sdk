// Code generated by go-swagger; DO NOT EDIT.

package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// StartVMReader is a Reader for the StartVM structure.
type StartVMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartVMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartVMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStartVMBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStartVMNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStartVMInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStartVMOK creates a StartVMOK with default headers values
func NewStartVMOK() *StartVMOK {
	return &StartVMOK{}
}

/* StartVMOK describes a response with status code 200, with default header values.

Ok
*/
type StartVMOK struct {
	Payload []*models.WithTaskVM
}

func (o *StartVMOK) Error() string {
	return fmt.Sprintf("[POST /start-vm][%d] startVmOK  %+v", 200, o.Payload)
}
func (o *StartVMOK) GetPayload() []*models.WithTaskVM {
	return o.Payload
}

func (o *StartVMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartVMBadRequest creates a StartVMBadRequest with default headers values
func NewStartVMBadRequest() *StartVMBadRequest {
	return &StartVMBadRequest{}
}

/* StartVMBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type StartVMBadRequest struct {
	Payload *models.ErrorBody
}

func (o *StartVMBadRequest) Error() string {
	return fmt.Sprintf("[POST /start-vm][%d] startVmBadRequest  %+v", 400, o.Payload)
}
func (o *StartVMBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *StartVMBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartVMNotFound creates a StartVMNotFound with default headers values
func NewStartVMNotFound() *StartVMNotFound {
	return &StartVMNotFound{}
}

/* StartVMNotFound describes a response with status code 404, with default header values.

Not found
*/
type StartVMNotFound struct {
	Payload *models.ErrorBody
}

func (o *StartVMNotFound) Error() string {
	return fmt.Sprintf("[POST /start-vm][%d] startVmNotFound  %+v", 404, o.Payload)
}
func (o *StartVMNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *StartVMNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartVMInternalServerError creates a StartVMInternalServerError with default headers values
func NewStartVMInternalServerError() *StartVMInternalServerError {
	return &StartVMInternalServerError{}
}

/* StartVMInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type StartVMInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *StartVMInternalServerError) Error() string {
	return fmt.Sprintf("[POST /start-vm][%d] startVmInternalServerError  %+v", 500, o.Payload)
}
func (o *StartVMInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *StartVMInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
