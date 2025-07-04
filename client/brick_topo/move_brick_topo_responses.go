// Code generated by go-swagger; DO NOT EDIT.

package brick_topo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// MoveBrickTopoReader is a Reader for the MoveBrickTopo structure.
type MoveBrickTopoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MoveBrickTopoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMoveBrickTopoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewMoveBrickTopoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewMoveBrickTopoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewMoveBrickTopoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMoveBrickTopoOK creates a MoveBrickTopoOK with default headers values
func NewMoveBrickTopoOK() *MoveBrickTopoOK {
	return &MoveBrickTopoOK{}
}

/* MoveBrickTopoOK describes a response with status code 200, with default header values.

MoveBrickTopoOK move brick topo o k
*/
type MoveBrickTopoOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskBrickTopo
}

func (o *MoveBrickTopoOK) Error() string {
	return fmt.Sprintf("[POST /move-brick-topo][%d] moveBrickTopoOK  %+v", 200, o.Payload)
}
func (o *MoveBrickTopoOK) GetPayload() []*models.WithTaskBrickTopo {
	return o.Payload
}

func (o *MoveBrickTopoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewMoveBrickTopoBadRequest creates a MoveBrickTopoBadRequest with default headers values
func NewMoveBrickTopoBadRequest() *MoveBrickTopoBadRequest {
	return &MoveBrickTopoBadRequest{}
}

/* MoveBrickTopoBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type MoveBrickTopoBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *MoveBrickTopoBadRequest) Error() string {
	return fmt.Sprintf("[POST /move-brick-topo][%d] moveBrickTopoBadRequest  %+v", 400, o.Payload)
}
func (o *MoveBrickTopoBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *MoveBrickTopoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewMoveBrickTopoNotFound creates a MoveBrickTopoNotFound with default headers values
func NewMoveBrickTopoNotFound() *MoveBrickTopoNotFound {
	return &MoveBrickTopoNotFound{}
}

/* MoveBrickTopoNotFound describes a response with status code 404, with default header values.

Not found
*/
type MoveBrickTopoNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *MoveBrickTopoNotFound) Error() string {
	return fmt.Sprintf("[POST /move-brick-topo][%d] moveBrickTopoNotFound  %+v", 404, o.Payload)
}
func (o *MoveBrickTopoNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *MoveBrickTopoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewMoveBrickTopoInternalServerError creates a MoveBrickTopoInternalServerError with default headers values
func NewMoveBrickTopoInternalServerError() *MoveBrickTopoInternalServerError {
	return &MoveBrickTopoInternalServerError{}
}

/* MoveBrickTopoInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type MoveBrickTopoInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *MoveBrickTopoInternalServerError) Error() string {
	return fmt.Sprintf("[POST /move-brick-topo][%d] moveBrickTopoInternalServerError  %+v", 500, o.Payload)
}
func (o *MoveBrickTopoInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *MoveBrickTopoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
