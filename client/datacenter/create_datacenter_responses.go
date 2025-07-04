// Code generated by go-swagger; DO NOT EDIT.

package datacenter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// CreateDatacenterReader is a Reader for the CreateDatacenter structure.
type CreateDatacenterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDatacenterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateDatacenterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewCreateDatacenterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewCreateDatacenterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewCreateDatacenterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateDatacenterOK creates a CreateDatacenterOK with default headers values
func NewCreateDatacenterOK() *CreateDatacenterOK {
	return &CreateDatacenterOK{}
}

/* CreateDatacenterOK describes a response with status code 200, with default header values.

CreateDatacenterOK create datacenter o k
*/
type CreateDatacenterOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskDatacenter
}

func (o *CreateDatacenterOK) Error() string {
	return fmt.Sprintf("[POST /create-datacenter][%d] createDatacenterOK  %+v", 200, o.Payload)
}
func (o *CreateDatacenterOK) GetPayload() []*models.WithTaskDatacenter {
	return o.Payload
}

func (o *CreateDatacenterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateDatacenterBadRequest creates a CreateDatacenterBadRequest with default headers values
func NewCreateDatacenterBadRequest() *CreateDatacenterBadRequest {
	return &CreateDatacenterBadRequest{}
}

/* CreateDatacenterBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateDatacenterBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateDatacenterBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-datacenter][%d] createDatacenterBadRequest  %+v", 400, o.Payload)
}
func (o *CreateDatacenterBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateDatacenterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateDatacenterNotFound creates a CreateDatacenterNotFound with default headers values
func NewCreateDatacenterNotFound() *CreateDatacenterNotFound {
	return &CreateDatacenterNotFound{}
}

/* CreateDatacenterNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateDatacenterNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateDatacenterNotFound) Error() string {
	return fmt.Sprintf("[POST /create-datacenter][%d] createDatacenterNotFound  %+v", 404, o.Payload)
}
func (o *CreateDatacenterNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateDatacenterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateDatacenterInternalServerError creates a CreateDatacenterInternalServerError with default headers values
func NewCreateDatacenterInternalServerError() *CreateDatacenterInternalServerError {
	return &CreateDatacenterInternalServerError{}
}

/* CreateDatacenterInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CreateDatacenterInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateDatacenterInternalServerError) Error() string {
	return fmt.Sprintf("[POST /create-datacenter][%d] createDatacenterInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateDatacenterInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateDatacenterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
