// Code generated by go-swagger; DO NOT EDIT.

package observability

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// ClearSystemServiceAlertNotificationConfigReader is a Reader for the ClearSystemServiceAlertNotificationConfig structure.
type ClearSystemServiceAlertNotificationConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClearSystemServiceAlertNotificationConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClearSystemServiceAlertNotificationConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewClearSystemServiceAlertNotificationConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewClearSystemServiceAlertNotificationConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewClearSystemServiceAlertNotificationConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewClearSystemServiceAlertNotificationConfigOK creates a ClearSystemServiceAlertNotificationConfigOK with default headers values
func NewClearSystemServiceAlertNotificationConfigOK() *ClearSystemServiceAlertNotificationConfigOK {
	return &ClearSystemServiceAlertNotificationConfigOK{}
}

/* ClearSystemServiceAlertNotificationConfigOK describes a response with status code 200, with default header values.

ClearSystemServiceAlertNotificationConfigOK clear system service alert notification config o k
*/
type ClearSystemServiceAlertNotificationConfigOK struct {
	XTowerRequestID string

	Payload *models.WithTaskDisassociateSystemServiceFromObsServiceResult
}

func (o *ClearSystemServiceAlertNotificationConfigOK) Error() string {
	return fmt.Sprintf("[POST /clear-system-service-alert-notification-config][%d] clearSystemServiceAlertNotificationConfigOK  %+v", 200, o.Payload)
}
func (o *ClearSystemServiceAlertNotificationConfigOK) GetPayload() *models.WithTaskDisassociateSystemServiceFromObsServiceResult {
	return o.Payload
}

func (o *ClearSystemServiceAlertNotificationConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.WithTaskDisassociateSystemServiceFromObsServiceResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClearSystemServiceAlertNotificationConfigBadRequest creates a ClearSystemServiceAlertNotificationConfigBadRequest with default headers values
func NewClearSystemServiceAlertNotificationConfigBadRequest() *ClearSystemServiceAlertNotificationConfigBadRequest {
	return &ClearSystemServiceAlertNotificationConfigBadRequest{}
}

/* ClearSystemServiceAlertNotificationConfigBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ClearSystemServiceAlertNotificationConfigBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *ClearSystemServiceAlertNotificationConfigBadRequest) Error() string {
	return fmt.Sprintf("[POST /clear-system-service-alert-notification-config][%d] clearSystemServiceAlertNotificationConfigBadRequest  %+v", 400, o.Payload)
}
func (o *ClearSystemServiceAlertNotificationConfigBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ClearSystemServiceAlertNotificationConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewClearSystemServiceAlertNotificationConfigNotFound creates a ClearSystemServiceAlertNotificationConfigNotFound with default headers values
func NewClearSystemServiceAlertNotificationConfigNotFound() *ClearSystemServiceAlertNotificationConfigNotFound {
	return &ClearSystemServiceAlertNotificationConfigNotFound{}
}

/* ClearSystemServiceAlertNotificationConfigNotFound describes a response with status code 404, with default header values.

Not found
*/
type ClearSystemServiceAlertNotificationConfigNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *ClearSystemServiceAlertNotificationConfigNotFound) Error() string {
	return fmt.Sprintf("[POST /clear-system-service-alert-notification-config][%d] clearSystemServiceAlertNotificationConfigNotFound  %+v", 404, o.Payload)
}
func (o *ClearSystemServiceAlertNotificationConfigNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ClearSystemServiceAlertNotificationConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewClearSystemServiceAlertNotificationConfigInternalServerError creates a ClearSystemServiceAlertNotificationConfigInternalServerError with default headers values
func NewClearSystemServiceAlertNotificationConfigInternalServerError() *ClearSystemServiceAlertNotificationConfigInternalServerError {
	return &ClearSystemServiceAlertNotificationConfigInternalServerError{}
}

/* ClearSystemServiceAlertNotificationConfigInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type ClearSystemServiceAlertNotificationConfigInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *ClearSystemServiceAlertNotificationConfigInternalServerError) Error() string {
	return fmt.Sprintf("[POST /clear-system-service-alert-notification-config][%d] clearSystemServiceAlertNotificationConfigInternalServerError  %+v", 500, o.Payload)
}
func (o *ClearSystemServiceAlertNotificationConfigInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *ClearSystemServiceAlertNotificationConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
