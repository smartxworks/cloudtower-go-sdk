// Code generated by go-swagger; DO NOT EDIT.

package log_collection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// DeleteLogCollectionReader is a Reader for the DeleteLogCollection structure.
type DeleteLogCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLogCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteLogCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewDeleteLogCollectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewDeleteLogCollectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewDeleteLogCollectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteLogCollectionOK creates a DeleteLogCollectionOK with default headers values
func NewDeleteLogCollectionOK() *DeleteLogCollectionOK {
	return &DeleteLogCollectionOK{}
}

/* DeleteLogCollectionOK describes a response with status code 200, with default header values.

DeleteLogCollectionOK delete log collection o k
*/
type DeleteLogCollectionOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskDeleteLogCollection
}

func (o *DeleteLogCollectionOK) Error() string {
	return fmt.Sprintf("[POST /delete-log-collection][%d] deleteLogCollectionOK  %+v", 200, o.Payload)
}
func (o *DeleteLogCollectionOK) GetPayload() []*models.WithTaskDeleteLogCollection {
	return o.Payload
}

func (o *DeleteLogCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteLogCollectionBadRequest creates a DeleteLogCollectionBadRequest with default headers values
func NewDeleteLogCollectionBadRequest() *DeleteLogCollectionBadRequest {
	return &DeleteLogCollectionBadRequest{}
}

/* DeleteLogCollectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteLogCollectionBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteLogCollectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-log-collection][%d] deleteLogCollectionBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteLogCollectionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLogCollectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteLogCollectionNotFound creates a DeleteLogCollectionNotFound with default headers values
func NewDeleteLogCollectionNotFound() *DeleteLogCollectionNotFound {
	return &DeleteLogCollectionNotFound{}
}

/* DeleteLogCollectionNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteLogCollectionNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteLogCollectionNotFound) Error() string {
	return fmt.Sprintf("[POST /delete-log-collection][%d] deleteLogCollectionNotFound  %+v", 404, o.Payload)
}
func (o *DeleteLogCollectionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLogCollectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteLogCollectionInternalServerError creates a DeleteLogCollectionInternalServerError with default headers values
func NewDeleteLogCollectionInternalServerError() *DeleteLogCollectionInternalServerError {
	return &DeleteLogCollectionInternalServerError{}
}

/* DeleteLogCollectionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type DeleteLogCollectionInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteLogCollectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /delete-log-collection][%d] deleteLogCollectionInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteLogCollectionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLogCollectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
