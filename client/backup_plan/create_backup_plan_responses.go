// Code generated by go-swagger; DO NOT EDIT.

package backup_plan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// CreateBackupPlanReader is a Reader for the CreateBackupPlan structure.
type CreateBackupPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateBackupPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateBackupPlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateBackupPlanBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateBackupPlanNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateBackupPlanInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateBackupPlanOK creates a CreateBackupPlanOK with default headers values
func NewCreateBackupPlanOK() *CreateBackupPlanOK {
	return &CreateBackupPlanOK{}
}

/* CreateBackupPlanOK describes a response with status code 200, with default header values.

CreateBackupPlanOK create backup plan o k
*/
type CreateBackupPlanOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskBackupPlan
}

func (o *CreateBackupPlanOK) Error() string {
	return fmt.Sprintf("[POST /create-backup-plan][%d] createBackupPlanOK  %+v", 200, o.Payload)
}
func (o *CreateBackupPlanOK) GetPayload() []*models.WithTaskBackupPlan {
	return o.Payload
}

func (o *CreateBackupPlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateBackupPlanBadRequest creates a CreateBackupPlanBadRequest with default headers values
func NewCreateBackupPlanBadRequest() *CreateBackupPlanBadRequest {
	return &CreateBackupPlanBadRequest{}
}

/* CreateBackupPlanBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateBackupPlanBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateBackupPlanBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-backup-plan][%d] createBackupPlanBadRequest  %+v", 400, o.Payload)
}
func (o *CreateBackupPlanBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateBackupPlanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateBackupPlanNotFound creates a CreateBackupPlanNotFound with default headers values
func NewCreateBackupPlanNotFound() *CreateBackupPlanNotFound {
	return &CreateBackupPlanNotFound{}
}

/* CreateBackupPlanNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateBackupPlanNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateBackupPlanNotFound) Error() string {
	return fmt.Sprintf("[POST /create-backup-plan][%d] createBackupPlanNotFound  %+v", 404, o.Payload)
}
func (o *CreateBackupPlanNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateBackupPlanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateBackupPlanInternalServerError creates a CreateBackupPlanInternalServerError with default headers values
func NewCreateBackupPlanInternalServerError() *CreateBackupPlanInternalServerError {
	return &CreateBackupPlanInternalServerError{}
}

/* CreateBackupPlanInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CreateBackupPlanInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateBackupPlanInternalServerError) Error() string {
	return fmt.Sprintf("[POST /create-backup-plan][%d] createBackupPlanInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateBackupPlanInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateBackupPlanInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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