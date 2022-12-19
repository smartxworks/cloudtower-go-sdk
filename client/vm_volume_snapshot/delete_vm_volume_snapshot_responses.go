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

// DeleteVMVolumeSnapshotReader is a Reader for the DeleteVMVolumeSnapshot structure.
type DeleteVMVolumeSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVMVolumeSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVMVolumeSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteVMVolumeSnapshotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteVMVolumeSnapshotNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteVMVolumeSnapshotInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteVMVolumeSnapshotOK creates a DeleteVMVolumeSnapshotOK with default headers values
func NewDeleteVMVolumeSnapshotOK() *DeleteVMVolumeSnapshotOK {
	return &DeleteVMVolumeSnapshotOK{}
}

/* DeleteVMVolumeSnapshotOK describes a response with status code 200, with default header values.

DeleteVMVolumeSnapshotOK delete Vm volume snapshot o k
*/
type DeleteVMVolumeSnapshotOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskDeleteVMVolumeSnapshot
}

func (o *DeleteVMVolumeSnapshotOK) Error() string {
	return fmt.Sprintf("[POST /delete-vm-volume-snapshot][%d] deleteVmVolumeSnapshotOK  %+v", 200, o.Payload)
}
func (o *DeleteVMVolumeSnapshotOK) GetPayload() []*models.WithTaskDeleteVMVolumeSnapshot {
	return o.Payload
}

func (o *DeleteVMVolumeSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteVMVolumeSnapshotBadRequest creates a DeleteVMVolumeSnapshotBadRequest with default headers values
func NewDeleteVMVolumeSnapshotBadRequest() *DeleteVMVolumeSnapshotBadRequest {
	return &DeleteVMVolumeSnapshotBadRequest{}
}

/* DeleteVMVolumeSnapshotBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteVMVolumeSnapshotBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteVMVolumeSnapshotBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-vm-volume-snapshot][%d] deleteVmVolumeSnapshotBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteVMVolumeSnapshotBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVMVolumeSnapshotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteVMVolumeSnapshotNotFound creates a DeleteVMVolumeSnapshotNotFound with default headers values
func NewDeleteVMVolumeSnapshotNotFound() *DeleteVMVolumeSnapshotNotFound {
	return &DeleteVMVolumeSnapshotNotFound{}
}

/* DeleteVMVolumeSnapshotNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteVMVolumeSnapshotNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteVMVolumeSnapshotNotFound) Error() string {
	return fmt.Sprintf("[POST /delete-vm-volume-snapshot][%d] deleteVmVolumeSnapshotNotFound  %+v", 404, o.Payload)
}
func (o *DeleteVMVolumeSnapshotNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVMVolumeSnapshotNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteVMVolumeSnapshotInternalServerError creates a DeleteVMVolumeSnapshotInternalServerError with default headers values
func NewDeleteVMVolumeSnapshotInternalServerError() *DeleteVMVolumeSnapshotInternalServerError {
	return &DeleteVMVolumeSnapshotInternalServerError{}
}

/* DeleteVMVolumeSnapshotInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type DeleteVMVolumeSnapshotInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteVMVolumeSnapshotInternalServerError) Error() string {
	return fmt.Sprintf("[POST /delete-vm-volume-snapshot][%d] deleteVmVolumeSnapshotInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteVMVolumeSnapshotInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteVMVolumeSnapshotInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
