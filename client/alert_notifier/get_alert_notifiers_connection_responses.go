// Code generated by go-swagger; DO NOT EDIT.

package alert_notifier

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetAlertNotifiersConnectionReader is a Reader for the GetAlertNotifiersConnection structure.
type GetAlertNotifiersConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertNotifiersConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertNotifiersConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAlertNotifiersConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAlertNotifiersConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAlertNotifiersConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAlertNotifiersConnectionOK creates a GetAlertNotifiersConnectionOK with default headers values
func NewGetAlertNotifiersConnectionOK() *GetAlertNotifiersConnectionOK {
	return &GetAlertNotifiersConnectionOK{}
}

/* GetAlertNotifiersConnectionOK describes a response with status code 200, with default header values.

GetAlertNotifiersConnectionOK get alert notifiers connection o k
*/
type GetAlertNotifiersConnectionOK struct {
	XTowerRequestID string

	Payload *models.AlertNotifierConnection
}

func (o *GetAlertNotifiersConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-alert-notifiers-connection][%d] getAlertNotifiersConnectionOK  %+v", 200, o.Payload)
}
func (o *GetAlertNotifiersConnectionOK) GetPayload() *models.AlertNotifierConnection {
	return o.Payload
}

func (o *GetAlertNotifiersConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.AlertNotifierConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertNotifiersConnectionBadRequest creates a GetAlertNotifiersConnectionBadRequest with default headers values
func NewGetAlertNotifiersConnectionBadRequest() *GetAlertNotifiersConnectionBadRequest {
	return &GetAlertNotifiersConnectionBadRequest{}
}

/* GetAlertNotifiersConnectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetAlertNotifiersConnectionBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetAlertNotifiersConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-alert-notifiers-connection][%d] getAlertNotifiersConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetAlertNotifiersConnectionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertNotifiersConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAlertNotifiersConnectionNotFound creates a GetAlertNotifiersConnectionNotFound with default headers values
func NewGetAlertNotifiersConnectionNotFound() *GetAlertNotifiersConnectionNotFound {
	return &GetAlertNotifiersConnectionNotFound{}
}

/* GetAlertNotifiersConnectionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetAlertNotifiersConnectionNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetAlertNotifiersConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /get-alert-notifiers-connection][%d] getAlertNotifiersConnectionNotFound  %+v", 404, o.Payload)
}
func (o *GetAlertNotifiersConnectionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertNotifiersConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAlertNotifiersConnectionInternalServerError creates a GetAlertNotifiersConnectionInternalServerError with default headers values
func NewGetAlertNotifiersConnectionInternalServerError() *GetAlertNotifiersConnectionInternalServerError {
	return &GetAlertNotifiersConnectionInternalServerError{}
}

/* GetAlertNotifiersConnectionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetAlertNotifiersConnectionInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetAlertNotifiersConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-alert-notifiers-connection][%d] getAlertNotifiersConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAlertNotifiersConnectionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertNotifiersConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
