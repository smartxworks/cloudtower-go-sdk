// Code generated by go-swagger; DO NOT EDIT.

package label

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// UpdateLabelReader is a Reader for the UpdateLabel structure.
type UpdateLabelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateLabelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateLabelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewUpdateLabelNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewUpdateLabelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateLabelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateLabelInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateLabelOK creates a UpdateLabelOK with default headers values
func NewUpdateLabelOK() *UpdateLabelOK {
	return &UpdateLabelOK{}
}

/* UpdateLabelOK describes a response with status code 200, with default header values.

UpdateLabelOK update label o k
*/
type UpdateLabelOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskLabel
}

func (o *UpdateLabelOK) Error() string {
	return fmt.Sprintf("[POST /update-label][%d] updateLabelOK  %+v", 200, o.Payload)
}
func (o *UpdateLabelOK) GetPayload() []*models.WithTaskLabel {
	return o.Payload
}

func (o *UpdateLabelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateLabelNotModified creates a UpdateLabelNotModified with default headers values
func NewUpdateLabelNotModified() *UpdateLabelNotModified {
	return &UpdateLabelNotModified{}
}

/* UpdateLabelNotModified describes a response with status code 304, with default header values.

Not modified
*/
type UpdateLabelNotModified struct {
}

func (o *UpdateLabelNotModified) Error() string {
	return fmt.Sprintf("[POST /update-label][%d] updateLabelNotModified ", 304)
}

func (o *UpdateLabelNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateLabelBadRequest creates a UpdateLabelBadRequest with default headers values
func NewUpdateLabelBadRequest() *UpdateLabelBadRequest {
	return &UpdateLabelBadRequest{}
}

/* UpdateLabelBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateLabelBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateLabelBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-label][%d] updateLabelBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateLabelBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateLabelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateLabelNotFound creates a UpdateLabelNotFound with default headers values
func NewUpdateLabelNotFound() *UpdateLabelNotFound {
	return &UpdateLabelNotFound{}
}

/* UpdateLabelNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateLabelNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateLabelNotFound) Error() string {
	return fmt.Sprintf("[POST /update-label][%d] updateLabelNotFound  %+v", 404, o.Payload)
}
func (o *UpdateLabelNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateLabelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateLabelInternalServerError creates a UpdateLabelInternalServerError with default headers values
func NewUpdateLabelInternalServerError() *UpdateLabelInternalServerError {
	return &UpdateLabelInternalServerError{}
}

/* UpdateLabelInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UpdateLabelInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateLabelInternalServerError) Error() string {
	return fmt.Sprintf("[POST /update-label][%d] updateLabelInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateLabelInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateLabelInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
