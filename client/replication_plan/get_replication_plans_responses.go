// Code generated by go-swagger; DO NOT EDIT.

package replication_plan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetReplicationPlansReader is a Reader for the GetReplicationPlans structure.
type GetReplicationPlansReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReplicationPlansReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReplicationPlansOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewGetReplicationPlansBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewGetReplicationPlansNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewGetReplicationPlansInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetReplicationPlansOK creates a GetReplicationPlansOK with default headers values
func NewGetReplicationPlansOK() *GetReplicationPlansOK {
	return &GetReplicationPlansOK{}
}

/* GetReplicationPlansOK describes a response with status code 200, with default header values.

GetReplicationPlansOK get replication plans o k
*/
type GetReplicationPlansOK struct {
	XTowerRequestID string

	Payload []*models.ReplicationPlan
}

func (o *GetReplicationPlansOK) Error() string {
	return fmt.Sprintf("[POST /get-replication-plans][%d] getReplicationPlansOK  %+v", 200, o.Payload)
}
func (o *GetReplicationPlansOK) GetPayload() []*models.ReplicationPlan {
	return o.Payload
}

func (o *GetReplicationPlansOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetReplicationPlansBadRequest creates a GetReplicationPlansBadRequest with default headers values
func NewGetReplicationPlansBadRequest() *GetReplicationPlansBadRequest {
	return &GetReplicationPlansBadRequest{}
}

/* GetReplicationPlansBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetReplicationPlansBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetReplicationPlansBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-replication-plans][%d] getReplicationPlansBadRequest  %+v", 400, o.Payload)
}
func (o *GetReplicationPlansBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetReplicationPlansBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetReplicationPlansNotFound creates a GetReplicationPlansNotFound with default headers values
func NewGetReplicationPlansNotFound() *GetReplicationPlansNotFound {
	return &GetReplicationPlansNotFound{}
}

/* GetReplicationPlansNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetReplicationPlansNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetReplicationPlansNotFound) Error() string {
	return fmt.Sprintf("[POST /get-replication-plans][%d] getReplicationPlansNotFound  %+v", 404, o.Payload)
}
func (o *GetReplicationPlansNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetReplicationPlansNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetReplicationPlansInternalServerError creates a GetReplicationPlansInternalServerError with default headers values
func NewGetReplicationPlansInternalServerError() *GetReplicationPlansInternalServerError {
	return &GetReplicationPlansInternalServerError{}
}

/* GetReplicationPlansInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetReplicationPlansInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetReplicationPlansInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-replication-plans][%d] getReplicationPlansInternalServerError  %+v", 500, o.Payload)
}
func (o *GetReplicationPlansInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetReplicationPlansInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
