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

// DeleteAlertNotifierReader is a Reader for the DeleteAlertNotifier structure.
type DeleteAlertNotifierReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAlertNotifierReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAlertNotifierOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewDeleteAlertNotifierBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewDeleteAlertNotifierNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewDeleteAlertNotifierInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAlertNotifierOK creates a DeleteAlertNotifierOK with default headers values
func NewDeleteAlertNotifierOK() *DeleteAlertNotifierOK {
	return &DeleteAlertNotifierOK{}
}

/* DeleteAlertNotifierOK describes a response with status code 200, with default header values.

DeleteAlertNotifierOK delete alert notifier o k
*/
type DeleteAlertNotifierOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskDeleteAlertNotifier
}

func (o *DeleteAlertNotifierOK) Error() string {
	return fmt.Sprintf("[POST /delete-alert-notifier][%d] deleteAlertNotifierOK  %+v", 200, o.Payload)
}
func (o *DeleteAlertNotifierOK) GetPayload() []*models.WithTaskDeleteAlertNotifier {
	return o.Payload
}

func (o *DeleteAlertNotifierOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteAlertNotifierBadRequest creates a DeleteAlertNotifierBadRequest with default headers values
func NewDeleteAlertNotifierBadRequest() *DeleteAlertNotifierBadRequest {
	return &DeleteAlertNotifierBadRequest{}
}

/* DeleteAlertNotifierBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteAlertNotifierBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteAlertNotifierBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-alert-notifier][%d] deleteAlertNotifierBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteAlertNotifierBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAlertNotifierBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteAlertNotifierNotFound creates a DeleteAlertNotifierNotFound with default headers values
func NewDeleteAlertNotifierNotFound() *DeleteAlertNotifierNotFound {
	return &DeleteAlertNotifierNotFound{}
}

/* DeleteAlertNotifierNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteAlertNotifierNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteAlertNotifierNotFound) Error() string {
	return fmt.Sprintf("[POST /delete-alert-notifier][%d] deleteAlertNotifierNotFound  %+v", 404, o.Payload)
}
func (o *DeleteAlertNotifierNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAlertNotifierNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteAlertNotifierInternalServerError creates a DeleteAlertNotifierInternalServerError with default headers values
func NewDeleteAlertNotifierInternalServerError() *DeleteAlertNotifierInternalServerError {
	return &DeleteAlertNotifierInternalServerError{}
}

/* DeleteAlertNotifierInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type DeleteAlertNotifierInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteAlertNotifierInternalServerError) Error() string {
	return fmt.Sprintf("[POST /delete-alert-notifier][%d] deleteAlertNotifierInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteAlertNotifierInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAlertNotifierInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
