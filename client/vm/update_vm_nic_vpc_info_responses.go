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

// UpdateVMNicVpcInfoReader is a Reader for the UpdateVMNicVpcInfo structure.
type UpdateVMNicVpcInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVMNicVpcInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateVMNicVpcInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewUpdateVMNicVpcInfoNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewUpdateVMNicVpcInfoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateVMNicVpcInfoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateVMNicVpcInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateVMNicVpcInfoOK creates a UpdateVMNicVpcInfoOK with default headers values
func NewUpdateVMNicVpcInfoOK() *UpdateVMNicVpcInfoOK {
	return &UpdateVMNicVpcInfoOK{}
}

/* UpdateVMNicVpcInfoOK describes a response with status code 200, with default header values.

UpdateVMNicVpcInfoOK update Vm nic vpc info o k
*/
type UpdateVMNicVpcInfoOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskVM
}

func (o *UpdateVMNicVpcInfoOK) Error() string {
	return fmt.Sprintf("[POST /update-vm-vpc-nic][%d] updateVmNicVpcInfoOK  %+v", 200, o.Payload)
}
func (o *UpdateVMNicVpcInfoOK) GetPayload() []*models.WithTaskVM {
	return o.Payload
}

func (o *UpdateVMNicVpcInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateVMNicVpcInfoNotModified creates a UpdateVMNicVpcInfoNotModified with default headers values
func NewUpdateVMNicVpcInfoNotModified() *UpdateVMNicVpcInfoNotModified {
	return &UpdateVMNicVpcInfoNotModified{}
}

/* UpdateVMNicVpcInfoNotModified describes a response with status code 304, with default header values.

Not modified
*/
type UpdateVMNicVpcInfoNotModified struct {
}

func (o *UpdateVMNicVpcInfoNotModified) Error() string {
	return fmt.Sprintf("[POST /update-vm-vpc-nic][%d] updateVmNicVpcInfoNotModified ", 304)
}

func (o *UpdateVMNicVpcInfoNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateVMNicVpcInfoBadRequest creates a UpdateVMNicVpcInfoBadRequest with default headers values
func NewUpdateVMNicVpcInfoBadRequest() *UpdateVMNicVpcInfoBadRequest {
	return &UpdateVMNicVpcInfoBadRequest{}
}

/* UpdateVMNicVpcInfoBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateVMNicVpcInfoBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateVMNicVpcInfoBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-vm-vpc-nic][%d] updateVmNicVpcInfoBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateVMNicVpcInfoBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateVMNicVpcInfoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateVMNicVpcInfoNotFound creates a UpdateVMNicVpcInfoNotFound with default headers values
func NewUpdateVMNicVpcInfoNotFound() *UpdateVMNicVpcInfoNotFound {
	return &UpdateVMNicVpcInfoNotFound{}
}

/* UpdateVMNicVpcInfoNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateVMNicVpcInfoNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateVMNicVpcInfoNotFound) Error() string {
	return fmt.Sprintf("[POST /update-vm-vpc-nic][%d] updateVmNicVpcInfoNotFound  %+v", 404, o.Payload)
}
func (o *UpdateVMNicVpcInfoNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateVMNicVpcInfoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateVMNicVpcInfoInternalServerError creates a UpdateVMNicVpcInfoInternalServerError with default headers values
func NewUpdateVMNicVpcInfoInternalServerError() *UpdateVMNicVpcInfoInternalServerError {
	return &UpdateVMNicVpcInfoInternalServerError{}
}

/* UpdateVMNicVpcInfoInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UpdateVMNicVpcInfoInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateVMNicVpcInfoInternalServerError) Error() string {
	return fmt.Sprintf("[POST /update-vm-vpc-nic][%d] updateVmNicVpcInfoInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateVMNicVpcInfoInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateVMNicVpcInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
