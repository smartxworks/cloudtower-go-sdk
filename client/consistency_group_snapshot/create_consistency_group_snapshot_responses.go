// Code generated by go-swagger; DO NOT EDIT.

package consistency_group_snapshot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// CreateConsistencyGroupSnapshotReader is a Reader for the CreateConsistencyGroupSnapshot structure.
type CreateConsistencyGroupSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateConsistencyGroupSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateConsistencyGroupSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewCreateConsistencyGroupSnapshotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewCreateConsistencyGroupSnapshotNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewCreateConsistencyGroupSnapshotInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateConsistencyGroupSnapshotOK creates a CreateConsistencyGroupSnapshotOK with default headers values
func NewCreateConsistencyGroupSnapshotOK() *CreateConsistencyGroupSnapshotOK {
	return &CreateConsistencyGroupSnapshotOK{}
}

/* CreateConsistencyGroupSnapshotOK describes a response with status code 200, with default header values.

CreateConsistencyGroupSnapshotOK create consistency group snapshot o k
*/
type CreateConsistencyGroupSnapshotOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskConsistencyGroupSnapshot
}

func (o *CreateConsistencyGroupSnapshotOK) Error() string {
	return fmt.Sprintf("[POST /create-consistency-snapshot-group][%d] createConsistencyGroupSnapshotOK  %+v", 200, o.Payload)
}
func (o *CreateConsistencyGroupSnapshotOK) GetPayload() []*models.WithTaskConsistencyGroupSnapshot {
	return o.Payload
}

func (o *CreateConsistencyGroupSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateConsistencyGroupSnapshotBadRequest creates a CreateConsistencyGroupSnapshotBadRequest with default headers values
func NewCreateConsistencyGroupSnapshotBadRequest() *CreateConsistencyGroupSnapshotBadRequest {
	return &CreateConsistencyGroupSnapshotBadRequest{}
}

/* CreateConsistencyGroupSnapshotBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateConsistencyGroupSnapshotBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateConsistencyGroupSnapshotBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-consistency-snapshot-group][%d] createConsistencyGroupSnapshotBadRequest  %+v", 400, o.Payload)
}
func (o *CreateConsistencyGroupSnapshotBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateConsistencyGroupSnapshotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateConsistencyGroupSnapshotNotFound creates a CreateConsistencyGroupSnapshotNotFound with default headers values
func NewCreateConsistencyGroupSnapshotNotFound() *CreateConsistencyGroupSnapshotNotFound {
	return &CreateConsistencyGroupSnapshotNotFound{}
}

/* CreateConsistencyGroupSnapshotNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateConsistencyGroupSnapshotNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateConsistencyGroupSnapshotNotFound) Error() string {
	return fmt.Sprintf("[POST /create-consistency-snapshot-group][%d] createConsistencyGroupSnapshotNotFound  %+v", 404, o.Payload)
}
func (o *CreateConsistencyGroupSnapshotNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateConsistencyGroupSnapshotNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateConsistencyGroupSnapshotInternalServerError creates a CreateConsistencyGroupSnapshotInternalServerError with default headers values
func NewCreateConsistencyGroupSnapshotInternalServerError() *CreateConsistencyGroupSnapshotInternalServerError {
	return &CreateConsistencyGroupSnapshotInternalServerError{}
}

/* CreateConsistencyGroupSnapshotInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CreateConsistencyGroupSnapshotInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateConsistencyGroupSnapshotInternalServerError) Error() string {
	return fmt.Sprintf("[POST /create-consistency-snapshot-group][%d] createConsistencyGroupSnapshotInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateConsistencyGroupSnapshotInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateConsistencyGroupSnapshotInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
