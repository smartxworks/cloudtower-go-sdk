// Code generated by go-swagger; DO NOT EDIT.

package view

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// DeleteViewReader is a Reader for the DeleteView structure.
type DeleteViewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteViewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteViewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewDeleteViewBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewDeleteViewNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewDeleteViewInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteViewOK creates a DeleteViewOK with default headers values
func NewDeleteViewOK() *DeleteViewOK {
	return &DeleteViewOK{}
}

/* DeleteViewOK describes a response with status code 200, with default header values.

DeleteViewOK delete view o k
*/
type DeleteViewOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskDeleteView
}

func (o *DeleteViewOK) Error() string {
	return fmt.Sprintf("[POST /delete-view][%d] deleteViewOK  %+v", 200, o.Payload)
}
func (o *DeleteViewOK) GetPayload() []*models.WithTaskDeleteView {
	return o.Payload
}

func (o *DeleteViewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteViewBadRequest creates a DeleteViewBadRequest with default headers values
func NewDeleteViewBadRequest() *DeleteViewBadRequest {
	return &DeleteViewBadRequest{}
}

/* DeleteViewBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteViewBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteViewBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-view][%d] deleteViewBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteViewBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteViewBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteViewNotFound creates a DeleteViewNotFound with default headers values
func NewDeleteViewNotFound() *DeleteViewNotFound {
	return &DeleteViewNotFound{}
}

/* DeleteViewNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteViewNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteViewNotFound) Error() string {
	return fmt.Sprintf("[POST /delete-view][%d] deleteViewNotFound  %+v", 404, o.Payload)
}
func (o *DeleteViewNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteViewNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteViewInternalServerError creates a DeleteViewInternalServerError with default headers values
func NewDeleteViewInternalServerError() *DeleteViewInternalServerError {
	return &DeleteViewInternalServerError{}
}

/* DeleteViewInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type DeleteViewInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteViewInternalServerError) Error() string {
	return fmt.Sprintf("[POST /delete-view][%d] deleteViewInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteViewInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteViewInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
