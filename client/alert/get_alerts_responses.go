// Code generated by go-swagger; DO NOT EDIT.

package alert

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetAlertsReader is a Reader for the GetAlerts structure.
type GetAlertsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewGetAlertsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewGetAlertsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewGetAlertsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAlertsOK creates a GetAlertsOK with default headers values
func NewGetAlertsOK() *GetAlertsOK {
	return &GetAlertsOK{}
}

/* GetAlertsOK describes a response with status code 200, with default header values.

GetAlertsOK get alerts o k
*/
type GetAlertsOK struct {
	XTowerRequestID string

	Payload []*models.Alert
}

func (o *GetAlertsOK) Error() string {
	return fmt.Sprintf("[POST /get-alerts][%d] getAlertsOK  %+v", 200, o.Payload)
}
func (o *GetAlertsOK) GetPayload() []*models.Alert {
	return o.Payload
}

func (o *GetAlertsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAlertsBadRequest creates a GetAlertsBadRequest with default headers values
func NewGetAlertsBadRequest() *GetAlertsBadRequest {
	return &GetAlertsBadRequest{}
}

/* GetAlertsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetAlertsBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetAlertsBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-alerts][%d] getAlertsBadRequest  %+v", 400, o.Payload)
}
func (o *GetAlertsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAlertsNotFound creates a GetAlertsNotFound with default headers values
func NewGetAlertsNotFound() *GetAlertsNotFound {
	return &GetAlertsNotFound{}
}

/* GetAlertsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetAlertsNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetAlertsNotFound) Error() string {
	return fmt.Sprintf("[POST /get-alerts][%d] getAlertsNotFound  %+v", 404, o.Payload)
}
func (o *GetAlertsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAlertsInternalServerError creates a GetAlertsInternalServerError with default headers values
func NewGetAlertsInternalServerError() *GetAlertsInternalServerError {
	return &GetAlertsInternalServerError{}
}

/* GetAlertsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetAlertsInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetAlertsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-alerts][%d] getAlertsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAlertsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
