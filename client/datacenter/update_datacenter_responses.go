// Code generated by go-swagger; DO NOT EDIT.

package datacenter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// UpdateDatacenterReader is a Reader for the UpdateDatacenter structure.
type UpdateDatacenterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDatacenterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDatacenterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewUpdateDatacenterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewUpdateDatacenterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewUpdateDatacenterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateDatacenterOK creates a UpdateDatacenterOK with default headers values
func NewUpdateDatacenterOK() *UpdateDatacenterOK {
	return &UpdateDatacenterOK{}
}

/* UpdateDatacenterOK describes a response with status code 200, with default header values.

UpdateDatacenterOK update datacenter o k
*/
type UpdateDatacenterOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskDatacenter
}

func (o *UpdateDatacenterOK) Error() string {
	return fmt.Sprintf("[POST /update-datacenter][%d] updateDatacenterOK  %+v", 200, o.Payload)
}
func (o *UpdateDatacenterOK) GetPayload() []*models.WithTaskDatacenter {
	return o.Payload
}

func (o *UpdateDatacenterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateDatacenterBadRequest creates a UpdateDatacenterBadRequest with default headers values
func NewUpdateDatacenterBadRequest() *UpdateDatacenterBadRequest {
	return &UpdateDatacenterBadRequest{}
}

/* UpdateDatacenterBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateDatacenterBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateDatacenterBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-datacenter][%d] updateDatacenterBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateDatacenterBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateDatacenterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateDatacenterNotFound creates a UpdateDatacenterNotFound with default headers values
func NewUpdateDatacenterNotFound() *UpdateDatacenterNotFound {
	return &UpdateDatacenterNotFound{}
}

/* UpdateDatacenterNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateDatacenterNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateDatacenterNotFound) Error() string {
	return fmt.Sprintf("[POST /update-datacenter][%d] updateDatacenterNotFound  %+v", 404, o.Payload)
}
func (o *UpdateDatacenterNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateDatacenterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateDatacenterInternalServerError creates a UpdateDatacenterInternalServerError with default headers values
func NewUpdateDatacenterInternalServerError() *UpdateDatacenterInternalServerError {
	return &UpdateDatacenterInternalServerError{}
}

/* UpdateDatacenterInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UpdateDatacenterInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateDatacenterInternalServerError) Error() string {
	return fmt.Sprintf("[POST /update-datacenter][%d] updateDatacenterInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateDatacenterInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateDatacenterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
