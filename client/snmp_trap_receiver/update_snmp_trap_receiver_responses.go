// Code generated by go-swagger; DO NOT EDIT.

package snmp_trap_receiver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// UpdateSnmpTrapReceiverReader is a Reader for the UpdateSnmpTrapReceiver structure.
type UpdateSnmpTrapReceiverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSnmpTrapReceiverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSnmpTrapReceiverOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewUpdateSnmpTrapReceiverBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewUpdateSnmpTrapReceiverNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewUpdateSnmpTrapReceiverInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateSnmpTrapReceiverOK creates a UpdateSnmpTrapReceiverOK with default headers values
func NewUpdateSnmpTrapReceiverOK() *UpdateSnmpTrapReceiverOK {
	return &UpdateSnmpTrapReceiverOK{}
}

/* UpdateSnmpTrapReceiverOK describes a response with status code 200, with default header values.

UpdateSnmpTrapReceiverOK update snmp trap receiver o k
*/
type UpdateSnmpTrapReceiverOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskSnmpTrapReceiver
}

func (o *UpdateSnmpTrapReceiverOK) Error() string {
	return fmt.Sprintf("[POST /update-snmp-trap-receiver][%d] updateSnmpTrapReceiverOK  %+v", 200, o.Payload)
}
func (o *UpdateSnmpTrapReceiverOK) GetPayload() []*models.WithTaskSnmpTrapReceiver {
	return o.Payload
}

func (o *UpdateSnmpTrapReceiverOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateSnmpTrapReceiverBadRequest creates a UpdateSnmpTrapReceiverBadRequest with default headers values
func NewUpdateSnmpTrapReceiverBadRequest() *UpdateSnmpTrapReceiverBadRequest {
	return &UpdateSnmpTrapReceiverBadRequest{}
}

/* UpdateSnmpTrapReceiverBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateSnmpTrapReceiverBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateSnmpTrapReceiverBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-snmp-trap-receiver][%d] updateSnmpTrapReceiverBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateSnmpTrapReceiverBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateSnmpTrapReceiverBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateSnmpTrapReceiverNotFound creates a UpdateSnmpTrapReceiverNotFound with default headers values
func NewUpdateSnmpTrapReceiverNotFound() *UpdateSnmpTrapReceiverNotFound {
	return &UpdateSnmpTrapReceiverNotFound{}
}

/* UpdateSnmpTrapReceiverNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateSnmpTrapReceiverNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateSnmpTrapReceiverNotFound) Error() string {
	return fmt.Sprintf("[POST /update-snmp-trap-receiver][%d] updateSnmpTrapReceiverNotFound  %+v", 404, o.Payload)
}
func (o *UpdateSnmpTrapReceiverNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateSnmpTrapReceiverNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateSnmpTrapReceiverInternalServerError creates a UpdateSnmpTrapReceiverInternalServerError with default headers values
func NewUpdateSnmpTrapReceiverInternalServerError() *UpdateSnmpTrapReceiverInternalServerError {
	return &UpdateSnmpTrapReceiverInternalServerError{}
}

/* UpdateSnmpTrapReceiverInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UpdateSnmpTrapReceiverInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateSnmpTrapReceiverInternalServerError) Error() string {
	return fmt.Sprintf("[POST /update-snmp-trap-receiver][%d] updateSnmpTrapReceiverInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateSnmpTrapReceiverInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateSnmpTrapReceiverInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
