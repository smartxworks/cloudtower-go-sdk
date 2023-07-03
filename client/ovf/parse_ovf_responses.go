// Code generated by go-swagger; DO NOT EDIT.

package ovf

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// ParseOvfReader is a Reader for the ParseOvf structure.
type ParseOvfReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ParseOvfReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewParseOvfOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewParseOvfBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewParseOvfNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewParseOvfInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewParseOvfOK creates a ParseOvfOK with default headers values
func NewParseOvfOK() *ParseOvfOK {
	return &ParseOvfOK{}
}

/* ParseOvfOK describes a response with status code 200, with default header values.

ParseOvfOK parse ovf o k
*/
type ParseOvfOK struct {
	XTowerRequestID string

	Payload *models.ParsedOVF
}

func (o *ParseOvfOK) Error() string {
	return fmt.Sprintf("[POST /parse-ovf][%d] parseOvfOK  %+v", 200, o.Payload)
}
func (o *ParseOvfOK) GetPayload() *models.ParsedOVF {
	return o.Payload
}

func (o *ParseOvfOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.ParsedOVF)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewParseOvfBadRequest creates a ParseOvfBadRequest with default headers values
func NewParseOvfBadRequest() *ParseOvfBadRequest {
	return &ParseOvfBadRequest{}
}

/* ParseOvfBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ParseOvfBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *ParseOvfBadRequest) Error() string {
	return fmt.Sprintf("[POST /parse-ovf][%d] parseOvfBadRequest  %+v", 400, o.Payload)
}
func (o *ParseOvfBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ParseOvfBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewParseOvfNotFound creates a ParseOvfNotFound with default headers values
func NewParseOvfNotFound() *ParseOvfNotFound {
	return &ParseOvfNotFound{}
}

/* ParseOvfNotFound describes a response with status code 404, with default header values.

Not found
*/
type ParseOvfNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *ParseOvfNotFound) Error() string {
	return fmt.Sprintf("[POST /parse-ovf][%d] parseOvfNotFound  %+v", 404, o.Payload)
}
func (o *ParseOvfNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ParseOvfNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewParseOvfInternalServerError creates a ParseOvfInternalServerError with default headers values
func NewParseOvfInternalServerError() *ParseOvfInternalServerError {
	return &ParseOvfInternalServerError{}
}

/* ParseOvfInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type ParseOvfInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *ParseOvfInternalServerError) Error() string {
	return fmt.Sprintf("[POST /parse-ovf][%d] parseOvfInternalServerError  %+v", 500, o.Payload)
}
func (o *ParseOvfInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ParseOvfInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
