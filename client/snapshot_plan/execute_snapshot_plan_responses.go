// Code generated by go-swagger; DO NOT EDIT.

package snapshot_plan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// ExecuteSnapshotPlanReader is a Reader for the ExecuteSnapshotPlan structure.
type ExecuteSnapshotPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExecuteSnapshotPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExecuteSnapshotPlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExecuteSnapshotPlanBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExecuteSnapshotPlanOK creates a ExecuteSnapshotPlanOK with default headers values
func NewExecuteSnapshotPlanOK() *ExecuteSnapshotPlanOK {
	return &ExecuteSnapshotPlanOK{}
}

/* ExecuteSnapshotPlanOK describes a response with status code 200, with default header values.

Ok
*/
type ExecuteSnapshotPlanOK struct {
	Payload []*models.WithTaskSnapshotPlan
}

func (o *ExecuteSnapshotPlanOK) Error() string {
	return fmt.Sprintf("[POST /execute-snapshot-plan][%d] executeSnapshotPlanOK  %+v", 200, o.Payload)
}
func (o *ExecuteSnapshotPlanOK) GetPayload() []*models.WithTaskSnapshotPlan {
	return o.Payload
}

func (o *ExecuteSnapshotPlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecuteSnapshotPlanBadRequest creates a ExecuteSnapshotPlanBadRequest with default headers values
func NewExecuteSnapshotPlanBadRequest() *ExecuteSnapshotPlanBadRequest {
	return &ExecuteSnapshotPlanBadRequest{}
}

/* ExecuteSnapshotPlanBadRequest describes a response with status code 400, with default header values.

ExecuteSnapshotPlanBadRequest execute snapshot plan bad request
*/
type ExecuteSnapshotPlanBadRequest struct {
	Payload string
}

func (o *ExecuteSnapshotPlanBadRequest) Error() string {
	return fmt.Sprintf("[POST /execute-snapshot-plan][%d] executeSnapshotPlanBadRequest  %+v", 400, o.Payload)
}
func (o *ExecuteSnapshotPlanBadRequest) GetPayload() string {
	return o.Payload
}

func (o *ExecuteSnapshotPlanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
