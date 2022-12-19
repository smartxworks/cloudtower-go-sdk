// Code generated by go-swagger; DO NOT EDIT.

package snapshot_plan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetSnapshotPlansReader is a Reader for the GetSnapshotPlans structure.
type GetSnapshotPlansReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSnapshotPlansReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSnapshotPlansOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSnapshotPlansBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSnapshotPlansNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSnapshotPlansInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSnapshotPlansOK creates a GetSnapshotPlansOK with default headers values
func NewGetSnapshotPlansOK() *GetSnapshotPlansOK {
	return &GetSnapshotPlansOK{}
}

/* GetSnapshotPlansOK describes a response with status code 200, with default header values.

GetSnapshotPlansOK get snapshot plans o k
*/
type GetSnapshotPlansOK struct {
	XTowerRequestID string

	Payload []*models.SnapshotPlan
}

func (o *GetSnapshotPlansOK) Error() string {
	return fmt.Sprintf("[POST /get-snapshot-plans][%d] getSnapshotPlansOK  %+v", 200, o.Payload)
}
func (o *GetSnapshotPlansOK) GetPayload() []*models.SnapshotPlan {
	return o.Payload
}

func (o *GetSnapshotPlansOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSnapshotPlansBadRequest creates a GetSnapshotPlansBadRequest with default headers values
func NewGetSnapshotPlansBadRequest() *GetSnapshotPlansBadRequest {
	return &GetSnapshotPlansBadRequest{}
}

/* GetSnapshotPlansBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetSnapshotPlansBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetSnapshotPlansBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-snapshot-plans][%d] getSnapshotPlansBadRequest  %+v", 400, o.Payload)
}
func (o *GetSnapshotPlansBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSnapshotPlansBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSnapshotPlansNotFound creates a GetSnapshotPlansNotFound with default headers values
func NewGetSnapshotPlansNotFound() *GetSnapshotPlansNotFound {
	return &GetSnapshotPlansNotFound{}
}

/* GetSnapshotPlansNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetSnapshotPlansNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetSnapshotPlansNotFound) Error() string {
	return fmt.Sprintf("[POST /get-snapshot-plans][%d] getSnapshotPlansNotFound  %+v", 404, o.Payload)
}
func (o *GetSnapshotPlansNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSnapshotPlansNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSnapshotPlansInternalServerError creates a GetSnapshotPlansInternalServerError with default headers values
func NewGetSnapshotPlansInternalServerError() *GetSnapshotPlansInternalServerError {
	return &GetSnapshotPlansInternalServerError{}
}

/* GetSnapshotPlansInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetSnapshotPlansInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetSnapshotPlansInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-snapshot-plans][%d] getSnapshotPlansInternalServerError  %+v", 500, o.Payload)
}
func (o *GetSnapshotPlansInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSnapshotPlansInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
