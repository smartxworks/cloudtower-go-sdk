// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// ConnectClusterReader is a Reader for the ConnectCluster structure.
type ConnectClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConnectClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConnectClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewConnectClusterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewConnectClusterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewConnectClusterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewConnectClusterOK creates a ConnectClusterOK with default headers values
func NewConnectClusterOK() *ConnectClusterOK {
	return &ConnectClusterOK{}
}

/* ConnectClusterOK describes a response with status code 200, with default header values.

ConnectClusterOK connect cluster o k
*/
type ConnectClusterOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskCluster
}

func (o *ConnectClusterOK) Error() string {
	return fmt.Sprintf("[POST /connect-cluster][%d] connectClusterOK  %+v", 200, o.Payload)
}
func (o *ConnectClusterOK) GetPayload() []*models.WithTaskCluster {
	return o.Payload
}

func (o *ConnectClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewConnectClusterBadRequest creates a ConnectClusterBadRequest with default headers values
func NewConnectClusterBadRequest() *ConnectClusterBadRequest {
	return &ConnectClusterBadRequest{}
}

/* ConnectClusterBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ConnectClusterBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *ConnectClusterBadRequest) Error() string {
	return fmt.Sprintf("[POST /connect-cluster][%d] connectClusterBadRequest  %+v", 400, o.Payload)
}
func (o *ConnectClusterBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ConnectClusterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewConnectClusterNotFound creates a ConnectClusterNotFound with default headers values
func NewConnectClusterNotFound() *ConnectClusterNotFound {
	return &ConnectClusterNotFound{}
}

/* ConnectClusterNotFound describes a response with status code 404, with default header values.

Not found
*/
type ConnectClusterNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *ConnectClusterNotFound) Error() string {
	return fmt.Sprintf("[POST /connect-cluster][%d] connectClusterNotFound  %+v", 404, o.Payload)
}
func (o *ConnectClusterNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ConnectClusterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewConnectClusterInternalServerError creates a ConnectClusterInternalServerError with default headers values
func NewConnectClusterInternalServerError() *ConnectClusterInternalServerError {
	return &ConnectClusterInternalServerError{}
}

/* ConnectClusterInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type ConnectClusterInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *ConnectClusterInternalServerError) Error() string {
	return fmt.Sprintf("[POST /connect-cluster][%d] connectClusterInternalServerError  %+v", 500, o.Payload)
}
func (o *ConnectClusterInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ConnectClusterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
