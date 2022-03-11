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

// UpdateVMDiskReader is a Reader for the UpdateVMDisk structure.
type UpdateVMDiskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVMDiskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateVMDiskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateVMDiskBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateVMDiskOK creates a UpdateVMDiskOK with default headers values
func NewUpdateVMDiskOK() *UpdateVMDiskOK {
	return &UpdateVMDiskOK{}
}

/* UpdateVMDiskOK describes a response with status code 200, with default header values.

Ok
*/
type UpdateVMDiskOK struct {
	Payload []*models.WithTaskVM
}

func (o *UpdateVMDiskOK) Error() string {
	return fmt.Sprintf("[POST /update-vm-disk][%d] updateVmDiskOK  %+v", 200, o.Payload)
}
func (o *UpdateVMDiskOK) GetPayload() []*models.WithTaskVM {
	return o.Payload
}

func (o *UpdateVMDiskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVMDiskBadRequest creates a UpdateVMDiskBadRequest with default headers values
func NewUpdateVMDiskBadRequest() *UpdateVMDiskBadRequest {
	return &UpdateVMDiskBadRequest{}
}

/* UpdateVMDiskBadRequest describes a response with status code 400, with default header values.

UpdateVMDiskBadRequest update Vm disk bad request
*/
type UpdateVMDiskBadRequest struct {
	Payload string
}

func (o *UpdateVMDiskBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-vm-disk][%d] updateVmDiskBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateVMDiskBadRequest) GetPayload() string {
	return o.Payload
}

func (o *UpdateVMDiskBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
