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

// CreateAlertNotifierReader is a Reader for the CreateAlertNotifier structure.
type CreateAlertNotifierReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAlertNotifierReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateAlertNotifierOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewCreateAlertNotifierBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewCreateAlertNotifierNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewCreateAlertNotifierInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateAlertNotifierOK creates a CreateAlertNotifierOK with default headers values
func NewCreateAlertNotifierOK() *CreateAlertNotifierOK {
	return &CreateAlertNotifierOK{}
}

/* CreateAlertNotifierOK describes a response with status code 200, with default header values.

CreateAlertNotifierOK create alert notifier o k
*/
type CreateAlertNotifierOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskAlertNotifier
}

func (o *CreateAlertNotifierOK) Error() string {
	return fmt.Sprintf("[POST /create-alert-notifier][%d] createAlertNotifierOK  %+v", 200, o.Payload)
}
func (o *CreateAlertNotifierOK) GetPayload() []*models.WithTaskAlertNotifier {
	return o.Payload
}

func (o *CreateAlertNotifierOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateAlertNotifierBadRequest creates a CreateAlertNotifierBadRequest with default headers values
func NewCreateAlertNotifierBadRequest() *CreateAlertNotifierBadRequest {
	return &CreateAlertNotifierBadRequest{}
}

/* CreateAlertNotifierBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateAlertNotifierBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateAlertNotifierBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-alert-notifier][%d] createAlertNotifierBadRequest  %+v", 400, o.Payload)
}
func (o *CreateAlertNotifierBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateAlertNotifierBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateAlertNotifierNotFound creates a CreateAlertNotifierNotFound with default headers values
func NewCreateAlertNotifierNotFound() *CreateAlertNotifierNotFound {
	return &CreateAlertNotifierNotFound{}
}

/* CreateAlertNotifierNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateAlertNotifierNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateAlertNotifierNotFound) Error() string {
	return fmt.Sprintf("[POST /create-alert-notifier][%d] createAlertNotifierNotFound  %+v", 404, o.Payload)
}
func (o *CreateAlertNotifierNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateAlertNotifierNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateAlertNotifierInternalServerError creates a CreateAlertNotifierInternalServerError with default headers values
func NewCreateAlertNotifierInternalServerError() *CreateAlertNotifierInternalServerError {
	return &CreateAlertNotifierInternalServerError{}
}

/* CreateAlertNotifierInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type CreateAlertNotifierInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *CreateAlertNotifierInternalServerError) Error() string {
	return fmt.Sprintf("[POST /create-alert-notifier][%d] createAlertNotifierInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateAlertNotifierInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *CreateAlertNotifierInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
