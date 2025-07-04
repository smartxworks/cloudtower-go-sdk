// Code generated by go-swagger; DO NOT EDIT.

package gpu_device

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// UpdateGpuDeviceDescriptionReader is a Reader for the UpdateGpuDeviceDescription structure.
type UpdateGpuDeviceDescriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateGpuDeviceDescriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateGpuDeviceDescriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 304:
		result := NewUpdateGpuDeviceDescriptionNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 400:
		result := NewUpdateGpuDeviceDescriptionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewUpdateGpuDeviceDescriptionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewUpdateGpuDeviceDescriptionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateGpuDeviceDescriptionOK creates a UpdateGpuDeviceDescriptionOK with default headers values
func NewUpdateGpuDeviceDescriptionOK() *UpdateGpuDeviceDescriptionOK {
	return &UpdateGpuDeviceDescriptionOK{}
}

/* UpdateGpuDeviceDescriptionOK describes a response with status code 200, with default header values.

UpdateGpuDeviceDescriptionOK update gpu device description o k
*/
type UpdateGpuDeviceDescriptionOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskGpuDevice
}

func (o *UpdateGpuDeviceDescriptionOK) Error() string {
	return fmt.Sprintf("[POST /update-gpu-device-description][%d] updateGpuDeviceDescriptionOK  %+v", 200, o.Payload)
}
func (o *UpdateGpuDeviceDescriptionOK) GetPayload() []*models.WithTaskGpuDevice {
	return o.Payload
}

func (o *UpdateGpuDeviceDescriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateGpuDeviceDescriptionNotModified creates a UpdateGpuDeviceDescriptionNotModified with default headers values
func NewUpdateGpuDeviceDescriptionNotModified() *UpdateGpuDeviceDescriptionNotModified {
	return &UpdateGpuDeviceDescriptionNotModified{}
}

/* UpdateGpuDeviceDescriptionNotModified describes a response with status code 304, with default header values.

Not modified
*/
type UpdateGpuDeviceDescriptionNotModified struct {
}

func (o *UpdateGpuDeviceDescriptionNotModified) Error() string {
	return fmt.Sprintf("[POST /update-gpu-device-description][%d] updateGpuDeviceDescriptionNotModified ", 304)
}

func (o *UpdateGpuDeviceDescriptionNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateGpuDeviceDescriptionBadRequest creates a UpdateGpuDeviceDescriptionBadRequest with default headers values
func NewUpdateGpuDeviceDescriptionBadRequest() *UpdateGpuDeviceDescriptionBadRequest {
	return &UpdateGpuDeviceDescriptionBadRequest{}
}

/* UpdateGpuDeviceDescriptionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateGpuDeviceDescriptionBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateGpuDeviceDescriptionBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-gpu-device-description][%d] updateGpuDeviceDescriptionBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateGpuDeviceDescriptionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateGpuDeviceDescriptionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateGpuDeviceDescriptionNotFound creates a UpdateGpuDeviceDescriptionNotFound with default headers values
func NewUpdateGpuDeviceDescriptionNotFound() *UpdateGpuDeviceDescriptionNotFound {
	return &UpdateGpuDeviceDescriptionNotFound{}
}

/* UpdateGpuDeviceDescriptionNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateGpuDeviceDescriptionNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateGpuDeviceDescriptionNotFound) Error() string {
	return fmt.Sprintf("[POST /update-gpu-device-description][%d] updateGpuDeviceDescriptionNotFound  %+v", 404, o.Payload)
}
func (o *UpdateGpuDeviceDescriptionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateGpuDeviceDescriptionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateGpuDeviceDescriptionInternalServerError creates a UpdateGpuDeviceDescriptionInternalServerError with default headers values
func NewUpdateGpuDeviceDescriptionInternalServerError() *UpdateGpuDeviceDescriptionInternalServerError {
	return &UpdateGpuDeviceDescriptionInternalServerError{}
}

/* UpdateGpuDeviceDescriptionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UpdateGpuDeviceDescriptionInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateGpuDeviceDescriptionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /update-gpu-device-description][%d] updateGpuDeviceDescriptionInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateGpuDeviceDescriptionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateGpuDeviceDescriptionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
