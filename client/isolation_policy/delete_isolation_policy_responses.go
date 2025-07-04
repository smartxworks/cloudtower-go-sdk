// Code generated by go-swagger; DO NOT EDIT.

package isolation_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// DeleteIsolationPolicyReader is a Reader for the DeleteIsolationPolicy structure.
type DeleteIsolationPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIsolationPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteIsolationPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewDeleteIsolationPolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewDeleteIsolationPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewDeleteIsolationPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteIsolationPolicyOK creates a DeleteIsolationPolicyOK with default headers values
func NewDeleteIsolationPolicyOK() *DeleteIsolationPolicyOK {
	return &DeleteIsolationPolicyOK{}
}

/* DeleteIsolationPolicyOK describes a response with status code 200, with default header values.

DeleteIsolationPolicyOK delete isolation policy o k
*/
type DeleteIsolationPolicyOK struct {
	XTowerRequestID string

	Payload []*models.DeleteIsolationPolicy
}

func (o *DeleteIsolationPolicyOK) Error() string {
	return fmt.Sprintf("[POST /delete-isolation-policy][%d] deleteIsolationPolicyOK  %+v", 200, o.Payload)
}
func (o *DeleteIsolationPolicyOK) GetPayload() []*models.DeleteIsolationPolicy {
	return o.Payload
}

func (o *DeleteIsolationPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteIsolationPolicyBadRequest creates a DeleteIsolationPolicyBadRequest with default headers values
func NewDeleteIsolationPolicyBadRequest() *DeleteIsolationPolicyBadRequest {
	return &DeleteIsolationPolicyBadRequest{}
}

/* DeleteIsolationPolicyBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteIsolationPolicyBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteIsolationPolicyBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-isolation-policy][%d] deleteIsolationPolicyBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteIsolationPolicyBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIsolationPolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteIsolationPolicyNotFound creates a DeleteIsolationPolicyNotFound with default headers values
func NewDeleteIsolationPolicyNotFound() *DeleteIsolationPolicyNotFound {
	return &DeleteIsolationPolicyNotFound{}
}

/* DeleteIsolationPolicyNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteIsolationPolicyNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteIsolationPolicyNotFound) Error() string {
	return fmt.Sprintf("[POST /delete-isolation-policy][%d] deleteIsolationPolicyNotFound  %+v", 404, o.Payload)
}
func (o *DeleteIsolationPolicyNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIsolationPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteIsolationPolicyInternalServerError creates a DeleteIsolationPolicyInternalServerError with default headers values
func NewDeleteIsolationPolicyInternalServerError() *DeleteIsolationPolicyInternalServerError {
	return &DeleteIsolationPolicyInternalServerError{}
}

/* DeleteIsolationPolicyInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type DeleteIsolationPolicyInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteIsolationPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /delete-isolation-policy][%d] deleteIsolationPolicyInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteIsolationPolicyInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIsolationPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
