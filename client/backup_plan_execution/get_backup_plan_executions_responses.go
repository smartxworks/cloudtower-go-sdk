// Code generated by go-swagger; DO NOT EDIT.

package backup_plan_execution

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetBackupPlanExecutionsReader is a Reader for the GetBackupPlanExecutions structure.
type GetBackupPlanExecutionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBackupPlanExecutionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBackupPlanExecutionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewGetBackupPlanExecutionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewGetBackupPlanExecutionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewGetBackupPlanExecutionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBackupPlanExecutionsOK creates a GetBackupPlanExecutionsOK with default headers values
func NewGetBackupPlanExecutionsOK() *GetBackupPlanExecutionsOK {
	return &GetBackupPlanExecutionsOK{}
}

/* GetBackupPlanExecutionsOK describes a response with status code 200, with default header values.

GetBackupPlanExecutionsOK get backup plan executions o k
*/
type GetBackupPlanExecutionsOK struct {
	XTowerRequestID string

	Payload []*models.BackupPlanExecution
}

func (o *GetBackupPlanExecutionsOK) Error() string {
	return fmt.Sprintf("[POST /get-backup-plan-executions][%d] getBackupPlanExecutionsOK  %+v", 200, o.Payload)
}
func (o *GetBackupPlanExecutionsOK) GetPayload() []*models.BackupPlanExecution {
	return o.Payload
}

func (o *GetBackupPlanExecutionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetBackupPlanExecutionsBadRequest creates a GetBackupPlanExecutionsBadRequest with default headers values
func NewGetBackupPlanExecutionsBadRequest() *GetBackupPlanExecutionsBadRequest {
	return &GetBackupPlanExecutionsBadRequest{}
}

/* GetBackupPlanExecutionsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetBackupPlanExecutionsBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetBackupPlanExecutionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-backup-plan-executions][%d] getBackupPlanExecutionsBadRequest  %+v", 400, o.Payload)
}
func (o *GetBackupPlanExecutionsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetBackupPlanExecutionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetBackupPlanExecutionsNotFound creates a GetBackupPlanExecutionsNotFound with default headers values
func NewGetBackupPlanExecutionsNotFound() *GetBackupPlanExecutionsNotFound {
	return &GetBackupPlanExecutionsNotFound{}
}

/* GetBackupPlanExecutionsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetBackupPlanExecutionsNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetBackupPlanExecutionsNotFound) Error() string {
	return fmt.Sprintf("[POST /get-backup-plan-executions][%d] getBackupPlanExecutionsNotFound  %+v", 404, o.Payload)
}
func (o *GetBackupPlanExecutionsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetBackupPlanExecutionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetBackupPlanExecutionsInternalServerError creates a GetBackupPlanExecutionsInternalServerError with default headers values
func NewGetBackupPlanExecutionsInternalServerError() *GetBackupPlanExecutionsInternalServerError {
	return &GetBackupPlanExecutionsInternalServerError{}
}

/* GetBackupPlanExecutionsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetBackupPlanExecutionsInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetBackupPlanExecutionsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-backup-plan-executions][%d] getBackupPlanExecutionsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetBackupPlanExecutionsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetBackupPlanExecutionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
