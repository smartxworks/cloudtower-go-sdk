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

// UpdateManyAlertNotifiersReader is a Reader for the UpdateManyAlertNotifiers structure.
type UpdateManyAlertNotifiersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateManyAlertNotifiersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateManyAlertNotifiersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateManyAlertNotifiersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateManyAlertNotifiersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateManyAlertNotifiersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateManyAlertNotifiersOK creates a UpdateManyAlertNotifiersOK with default headers values
func NewUpdateManyAlertNotifiersOK() *UpdateManyAlertNotifiersOK {
	return &UpdateManyAlertNotifiersOK{}
}

/* UpdateManyAlertNotifiersOK describes a response with status code 200, with default header values.

UpdateManyAlertNotifiersOK update many alert notifiers o k
*/
type UpdateManyAlertNotifiersOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskAlertNotifier
}

func (o *UpdateManyAlertNotifiersOK) Error() string {
	return fmt.Sprintf("[POST /update-many-alert-notifiers][%d] updateManyAlertNotifiersOK  %+v", 200, o.Payload)
}
func (o *UpdateManyAlertNotifiersOK) GetPayload() []*models.WithTaskAlertNotifier {
	return o.Payload
}

func (o *UpdateManyAlertNotifiersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateManyAlertNotifiersBadRequest creates a UpdateManyAlertNotifiersBadRequest with default headers values
func NewUpdateManyAlertNotifiersBadRequest() *UpdateManyAlertNotifiersBadRequest {
	return &UpdateManyAlertNotifiersBadRequest{}
}

/* UpdateManyAlertNotifiersBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateManyAlertNotifiersBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateManyAlertNotifiersBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-many-alert-notifiers][%d] updateManyAlertNotifiersBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateManyAlertNotifiersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateManyAlertNotifiersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateManyAlertNotifiersNotFound creates a UpdateManyAlertNotifiersNotFound with default headers values
func NewUpdateManyAlertNotifiersNotFound() *UpdateManyAlertNotifiersNotFound {
	return &UpdateManyAlertNotifiersNotFound{}
}

/* UpdateManyAlertNotifiersNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateManyAlertNotifiersNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateManyAlertNotifiersNotFound) Error() string {
	return fmt.Sprintf("[POST /update-many-alert-notifiers][%d] updateManyAlertNotifiersNotFound  %+v", 404, o.Payload)
}
func (o *UpdateManyAlertNotifiersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateManyAlertNotifiersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateManyAlertNotifiersInternalServerError creates a UpdateManyAlertNotifiersInternalServerError with default headers values
func NewUpdateManyAlertNotifiersInternalServerError() *UpdateManyAlertNotifiersInternalServerError {
	return &UpdateManyAlertNotifiersInternalServerError{}
}

/* UpdateManyAlertNotifiersInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UpdateManyAlertNotifiersInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateManyAlertNotifiersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /update-many-alert-notifiers][%d] updateManyAlertNotifiersInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateManyAlertNotifiersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateManyAlertNotifiersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
