// Code generated by go-swagger; DO NOT EDIT.

package vm_volume

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// GetVMVolumesConnectionReader is a Reader for the GetVMVolumesConnection structure.
type GetVMVolumesConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVMVolumesConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVMVolumesConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVMVolumesConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVMVolumesConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVMVolumesConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVMVolumesConnectionOK creates a GetVMVolumesConnectionOK with default headers values
func NewGetVMVolumesConnectionOK() *GetVMVolumesConnectionOK {
	return &GetVMVolumesConnectionOK{}
}

/* GetVMVolumesConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetVMVolumesConnectionOK struct {
	Payload *models.VMVolumeConnection
}

func (o *GetVMVolumesConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-vm-volumes-connection][%d] getVmVolumesConnectionOK  %+v", 200, o.Payload)
}
func (o *GetVMVolumesConnectionOK) GetPayload() *models.VMVolumeConnection {
	return o.Payload
}

func (o *GetVMVolumesConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMVolumeConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVMVolumesConnectionBadRequest creates a GetVMVolumesConnectionBadRequest with default headers values
func NewGetVMVolumesConnectionBadRequest() *GetVMVolumesConnectionBadRequest {
	return &GetVMVolumesConnectionBadRequest{}
}

/* GetVMVolumesConnectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetVMVolumesConnectionBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetVMVolumesConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-vm-volumes-connection][%d] getVmVolumesConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetVMVolumesConnectionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVMVolumesConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVMVolumesConnectionNotFound creates a GetVMVolumesConnectionNotFound with default headers values
func NewGetVMVolumesConnectionNotFound() *GetVMVolumesConnectionNotFound {
	return &GetVMVolumesConnectionNotFound{}
}

/* GetVMVolumesConnectionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetVMVolumesConnectionNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetVMVolumesConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /get-vm-volumes-connection][%d] getVmVolumesConnectionNotFound  %+v", 404, o.Payload)
}
func (o *GetVMVolumesConnectionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVMVolumesConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVMVolumesConnectionInternalServerError creates a GetVMVolumesConnectionInternalServerError with default headers values
func NewGetVMVolumesConnectionInternalServerError() *GetVMVolumesConnectionInternalServerError {
	return &GetVMVolumesConnectionInternalServerError{}
}

/* GetVMVolumesConnectionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetVMVolumesConnectionInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetVMVolumesConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-vm-volumes-connection][%d] getVmVolumesConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetVMVolumesConnectionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVMVolumesConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
