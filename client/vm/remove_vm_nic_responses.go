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

// RemoveVMNicReader is a Reader for the RemoveVMNic structure.
type RemoveVMNicReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveVMNicReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveVMNicOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveVMNicBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRemoveVMNicOK creates a RemoveVMNicOK with default headers values
func NewRemoveVMNicOK() *RemoveVMNicOK {
	return &RemoveVMNicOK{}
}

/* RemoveVMNicOK describes a response with status code 200, with default header values.

Ok
*/
type RemoveVMNicOK struct {
	Payload []*models.WithTaskVM
}

func (o *RemoveVMNicOK) Error() string {
	return fmt.Sprintf("[POST /remove-vm-nic][%d] removeVmNicOK  %+v", 200, o.Payload)
}
func (o *RemoveVMNicOK) GetPayload() []*models.WithTaskVM {
	return o.Payload
}

func (o *RemoveVMNicOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveVMNicBadRequest creates a RemoveVMNicBadRequest with default headers values
func NewRemoveVMNicBadRequest() *RemoveVMNicBadRequest {
	return &RemoveVMNicBadRequest{}
}

/* RemoveVMNicBadRequest describes a response with status code 400, with default header values.

RemoveVMNicBadRequest remove Vm nic bad request
*/
type RemoveVMNicBadRequest struct {
	Payload string
}

func (o *RemoveVMNicBadRequest) Error() string {
	return fmt.Sprintf("[POST /remove-vm-nic][%d] removeVmNicBadRequest  %+v", 400, o.Payload)
}
func (o *RemoveVMNicBadRequest) GetPayload() string {
	return o.Payload
}

func (o *RemoveVMNicBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
