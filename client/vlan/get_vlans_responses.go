// Code generated by go-swagger; DO NOT EDIT.

package vlan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetVlansReader is a Reader for the GetVlans structure.
type GetVlansReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVlansReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVlansOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewGetVlansBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewGetVlansNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewGetVlansInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVlansOK creates a GetVlansOK with default headers values
func NewGetVlansOK() *GetVlansOK {
	return &GetVlansOK{}
}

/* GetVlansOK describes a response with status code 200, with default header values.

GetVlansOK get vlans o k
*/
type GetVlansOK struct {
	XTowerRequestID string

	Payload []*models.Vlan
}

func (o *GetVlansOK) Error() string {
	return fmt.Sprintf("[POST /get-vlans][%d] getVlansOK  %+v", 200, o.Payload)
}
func (o *GetVlansOK) GetPayload() []*models.Vlan {
	return o.Payload
}

func (o *GetVlansOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetVlansBadRequest creates a GetVlansBadRequest with default headers values
func NewGetVlansBadRequest() *GetVlansBadRequest {
	return &GetVlansBadRequest{}
}

/* GetVlansBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetVlansBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetVlansBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-vlans][%d] getVlansBadRequest  %+v", 400, o.Payload)
}
func (o *GetVlansBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVlansBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetVlansNotFound creates a GetVlansNotFound with default headers values
func NewGetVlansNotFound() *GetVlansNotFound {
	return &GetVlansNotFound{}
}

/* GetVlansNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetVlansNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetVlansNotFound) Error() string {
	return fmt.Sprintf("[POST /get-vlans][%d] getVlansNotFound  %+v", 404, o.Payload)
}
func (o *GetVlansNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVlansNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetVlansInternalServerError creates a GetVlansInternalServerError with default headers values
func NewGetVlansInternalServerError() *GetVlansInternalServerError {
	return &GetVlansInternalServerError{}
}

/* GetVlansInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetVlansInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetVlansInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-vlans][%d] getVlansInternalServerError  %+v", 500, o.Payload)
}
func (o *GetVlansInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVlansInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
