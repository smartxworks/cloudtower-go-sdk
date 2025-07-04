// Code generated by go-swagger; DO NOT EDIT.

package vm_volume_snapshot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// CreateVMVolumeSnapshotReader is a Reader for the CreateVMVolumeSnapshot structure.
type CreateVMVolumeSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVMVolumeSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateVMVolumeSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewCreateVMVolumeSnapshotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewCreateVMVolumeSnapshotNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewCreateVMVolumeSnapshotInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateVMVolumeSnapshotOK creates a CreateVMVolumeSnapshotOK with default headers values
func NewCreateVMVolumeSnapshotOK() *CreateVMVolumeSnapshotOK {
	return &CreateVMVolumeSnapshotOK{}
}

/* CreateVMVolumeSnapshotOK describes a response with status code 200, with default header values.

CreateVMVolumeSnapshotOK create Vm volume snapshot o k
*/
type CreateVMVolumeSnapshotOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskVMVolumeSnapshot
}

func (o *CreateVMVolumeSnapshotOK) Error() string {
	return fmt.Sprintf("[POST /create-vm-volume-snapshot][%d] createVmVolumeSnapshotOK  %+v", 200, o.Payload)
}
func (o *CreateVMVolumeSnapshotOK) GetPayload() []*models.WithTaskVMVolumeSnapshot {
	return o.Payload
}

func (o *CreateVMVolumeSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateVMVolumeSnapshotBadRequest creates a CreateVMVolumeSnapshotBadRequest with default headers values
func NewCreateVMVolumeSnapshotBadRequest() *CreateVMVolumeSnapshotBadRequest {
	return &CreateVMVolumeSnapshotBadRequest{}
}

/* CreateVMVolumeSnapshotBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateVMVolumeSnapshotBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateVMVolumeSnapshotBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-vm-volume-snapshot][%d] createVmVolumeSnapshotBadRequest  %+v", 400, o.Payload)
}
func (o *CreateVMVolumeSnapshotBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateVMVolumeSnapshotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateVMVolumeSnapshotNotFound creates a CreateVMVolumeSnapshotNotFound with default headers values
func NewCreateVMVolumeSnapshotNotFound() *CreateVMVolumeSnapshotNotFound {
	return &CreateVMVolumeSnapshotNotFound{}
}

/* CreateVMVolumeSnapshotNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateVMVolumeSnapshotNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateVMVolumeSnapshotNotFound) Error() string {
	return fmt.Sprintf("[POST /create-vm-volume-snapshot][%d] createVmVolumeSnapshotNotFound  %+v", 404, o.Payload)
}
func (o *CreateVMVolumeSnapshotNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateVMVolumeSnapshotNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateVMVolumeSnapshotInternalServerError creates a CreateVMVolumeSnapshotInternalServerError with default headers values
func NewCreateVMVolumeSnapshotInternalServerError() *CreateVMVolumeSnapshotInternalServerError {
	return &CreateVMVolumeSnapshotInternalServerError{}
}

/* CreateVMVolumeSnapshotInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CreateVMVolumeSnapshotInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateVMVolumeSnapshotInternalServerError) Error() string {
	return fmt.Sprintf("[POST /create-vm-volume-snapshot][%d] createVmVolumeSnapshotInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateVMVolumeSnapshotInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateVMVolumeSnapshotInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
