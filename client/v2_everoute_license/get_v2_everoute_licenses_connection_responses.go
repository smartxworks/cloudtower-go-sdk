// Code generated by go-swagger; DO NOT EDIT.

package v2_everoute_license

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetV2EverouteLicensesConnectionReader is a Reader for the GetV2EverouteLicensesConnection structure.
type GetV2EverouteLicensesConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV2EverouteLicensesConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV2EverouteLicensesConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewGetV2EverouteLicensesConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewGetV2EverouteLicensesConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewGetV2EverouteLicensesConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV2EverouteLicensesConnectionOK creates a GetV2EverouteLicensesConnectionOK with default headers values
func NewGetV2EverouteLicensesConnectionOK() *GetV2EverouteLicensesConnectionOK {
	return &GetV2EverouteLicensesConnectionOK{}
}

/* GetV2EverouteLicensesConnectionOK describes a response with status code 200, with default header values.

GetV2EverouteLicensesConnectionOK get v2 everoute licenses connection o k
*/
type GetV2EverouteLicensesConnectionOK struct {
	XTowerRequestID string

	Payload *models.V2EverouteLicenseConnection
}

func (o *GetV2EverouteLicensesConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-v-2-everoute-licenses-connection][%d] getV2EverouteLicensesConnectionOK  %+v", 200, o.Payload)
}
func (o *GetV2EverouteLicensesConnectionOK) GetPayload() *models.V2EverouteLicenseConnection {
	return o.Payload
}

func (o *GetV2EverouteLicensesConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.V2EverouteLicenseConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV2EverouteLicensesConnectionBadRequest creates a GetV2EverouteLicensesConnectionBadRequest with default headers values
func NewGetV2EverouteLicensesConnectionBadRequest() *GetV2EverouteLicensesConnectionBadRequest {
	return &GetV2EverouteLicensesConnectionBadRequest{}
}

/* GetV2EverouteLicensesConnectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetV2EverouteLicensesConnectionBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetV2EverouteLicensesConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-v-2-everoute-licenses-connection][%d] getV2EverouteLicensesConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetV2EverouteLicensesConnectionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetV2EverouteLicensesConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetV2EverouteLicensesConnectionNotFound creates a GetV2EverouteLicensesConnectionNotFound with default headers values
func NewGetV2EverouteLicensesConnectionNotFound() *GetV2EverouteLicensesConnectionNotFound {
	return &GetV2EverouteLicensesConnectionNotFound{}
}

/* GetV2EverouteLicensesConnectionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetV2EverouteLicensesConnectionNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetV2EverouteLicensesConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /get-v-2-everoute-licenses-connection][%d] getV2EverouteLicensesConnectionNotFound  %+v", 404, o.Payload)
}
func (o *GetV2EverouteLicensesConnectionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetV2EverouteLicensesConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetV2EverouteLicensesConnectionInternalServerError creates a GetV2EverouteLicensesConnectionInternalServerError with default headers values
func NewGetV2EverouteLicensesConnectionInternalServerError() *GetV2EverouteLicensesConnectionInternalServerError {
	return &GetV2EverouteLicensesConnectionInternalServerError{}
}

/* GetV2EverouteLicensesConnectionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetV2EverouteLicensesConnectionInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetV2EverouteLicensesConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-v-2-everoute-licenses-connection][%d] getV2EverouteLicensesConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetV2EverouteLicensesConnectionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetV2EverouteLicensesConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
