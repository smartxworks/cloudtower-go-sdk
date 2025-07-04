// Code generated by go-swagger; DO NOT EDIT.

package vm_volume

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// RebuildVMVolumeReader is a Reader for the RebuildVMVolume structure.
type RebuildVMVolumeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RebuildVMVolumeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRebuildVMVolumeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 304:
		result := NewRebuildVMVolumeNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 400:
		result := NewRebuildVMVolumeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewRebuildVMVolumeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewRebuildVMVolumeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRebuildVMVolumeOK creates a RebuildVMVolumeOK with default headers values
func NewRebuildVMVolumeOK() *RebuildVMVolumeOK {
	return &RebuildVMVolumeOK{}
}

/* RebuildVMVolumeOK describes a response with status code 200, with default header values.

RebuildVMVolumeOK rebuild Vm volume o k
*/
type RebuildVMVolumeOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskVMVolume
}

func (o *RebuildVMVolumeOK) Error() string {
	return fmt.Sprintf("[POST /rebuild-vm-volume][%d] rebuildVmVolumeOK  %+v", 200, o.Payload)
}
func (o *RebuildVMVolumeOK) GetPayload() []*models.WithTaskVMVolume {
	return o.Payload
}

func (o *RebuildVMVolumeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRebuildVMVolumeNotModified creates a RebuildVMVolumeNotModified with default headers values
func NewRebuildVMVolumeNotModified() *RebuildVMVolumeNotModified {
	return &RebuildVMVolumeNotModified{}
}

/* RebuildVMVolumeNotModified describes a response with status code 304, with default header values.

Not modified
*/
type RebuildVMVolumeNotModified struct {
}

func (o *RebuildVMVolumeNotModified) Error() string {
	return fmt.Sprintf("[POST /rebuild-vm-volume][%d] rebuildVmVolumeNotModified ", 304)
}

func (o *RebuildVMVolumeNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRebuildVMVolumeBadRequest creates a RebuildVMVolumeBadRequest with default headers values
func NewRebuildVMVolumeBadRequest() *RebuildVMVolumeBadRequest {
	return &RebuildVMVolumeBadRequest{}
}

/* RebuildVMVolumeBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type RebuildVMVolumeBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *RebuildVMVolumeBadRequest) Error() string {
	return fmt.Sprintf("[POST /rebuild-vm-volume][%d] rebuildVmVolumeBadRequest  %+v", 400, o.Payload)
}
func (o *RebuildVMVolumeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *RebuildVMVolumeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRebuildVMVolumeNotFound creates a RebuildVMVolumeNotFound with default headers values
func NewRebuildVMVolumeNotFound() *RebuildVMVolumeNotFound {
	return &RebuildVMVolumeNotFound{}
}

/* RebuildVMVolumeNotFound describes a response with status code 404, with default header values.

Not found
*/
type RebuildVMVolumeNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *RebuildVMVolumeNotFound) Error() string {
	return fmt.Sprintf("[POST /rebuild-vm-volume][%d] rebuildVmVolumeNotFound  %+v", 404, o.Payload)
}
func (o *RebuildVMVolumeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *RebuildVMVolumeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRebuildVMVolumeInternalServerError creates a RebuildVMVolumeInternalServerError with default headers values
func NewRebuildVMVolumeInternalServerError() *RebuildVMVolumeInternalServerError {
	return &RebuildVMVolumeInternalServerError{}
}

/* RebuildVMVolumeInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type RebuildVMVolumeInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *RebuildVMVolumeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /rebuild-vm-volume][%d] rebuildVmVolumeInternalServerError  %+v", 500, o.Payload)
}
func (o *RebuildVMVolumeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *RebuildVMVolumeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
