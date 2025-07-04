// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// DeleteOrganizationReader is a Reader for the DeleteOrganization structure.
type DeleteOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewDeleteOrganizationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewDeleteOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewDeleteOrganizationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOrganizationOK creates a DeleteOrganizationOK with default headers values
func NewDeleteOrganizationOK() *DeleteOrganizationOK {
	return &DeleteOrganizationOK{}
}

/* DeleteOrganizationOK describes a response with status code 200, with default header values.

DeleteOrganizationOK delete organization o k
*/
type DeleteOrganizationOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskDeleteOrganization
}

func (o *DeleteOrganizationOK) Error() string {
	return fmt.Sprintf("[POST /delete-organization][%d] deleteOrganizationOK  %+v", 200, o.Payload)
}
func (o *DeleteOrganizationOK) GetPayload() []*models.WithTaskDeleteOrganization {
	return o.Payload
}

func (o *DeleteOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteOrganizationBadRequest creates a DeleteOrganizationBadRequest with default headers values
func NewDeleteOrganizationBadRequest() *DeleteOrganizationBadRequest {
	return &DeleteOrganizationBadRequest{}
}

/* DeleteOrganizationBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteOrganizationBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteOrganizationBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-organization][%d] deleteOrganizationBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteOrganizationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrganizationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteOrganizationNotFound creates a DeleteOrganizationNotFound with default headers values
func NewDeleteOrganizationNotFound() *DeleteOrganizationNotFound {
	return &DeleteOrganizationNotFound{}
}

/* DeleteOrganizationNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteOrganizationNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteOrganizationNotFound) Error() string {
	return fmt.Sprintf("[POST /delete-organization][%d] deleteOrganizationNotFound  %+v", 404, o.Payload)
}
func (o *DeleteOrganizationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteOrganizationInternalServerError creates a DeleteOrganizationInternalServerError with default headers values
func NewDeleteOrganizationInternalServerError() *DeleteOrganizationInternalServerError {
	return &DeleteOrganizationInternalServerError{}
}

/* DeleteOrganizationInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type DeleteOrganizationInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteOrganizationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /delete-organization][%d] deleteOrganizationInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteOrganizationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrganizationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
