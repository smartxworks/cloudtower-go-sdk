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

// GetV2EverouteLicensesReader is a Reader for the GetV2EverouteLicenses structure.
type GetV2EverouteLicensesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV2EverouteLicensesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV2EverouteLicensesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetV2EverouteLicensesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetV2EverouteLicensesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetV2EverouteLicensesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV2EverouteLicensesOK creates a GetV2EverouteLicensesOK with default headers values
func NewGetV2EverouteLicensesOK() *GetV2EverouteLicensesOK {
	return &GetV2EverouteLicensesOK{}
}

/* GetV2EverouteLicensesOK describes a response with status code 200, with default header values.

GetV2EverouteLicensesOK get v2 everoute licenses o k
*/
type GetV2EverouteLicensesOK struct {
	XTowerRequestID string

	Payload []*models.V2EverouteLicense
}

func (o *GetV2EverouteLicensesOK) Error() string {
	return fmt.Sprintf("[POST /get-v2-everoute-licenses][%d] getV2EverouteLicensesOK  %+v", 200, o.Payload)
}
func (o *GetV2EverouteLicensesOK) GetPayload() []*models.V2EverouteLicense {
	return o.Payload
}

func (o *GetV2EverouteLicensesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetV2EverouteLicensesBadRequest creates a GetV2EverouteLicensesBadRequest with default headers values
func NewGetV2EverouteLicensesBadRequest() *GetV2EverouteLicensesBadRequest {
	return &GetV2EverouteLicensesBadRequest{}
}

/* GetV2EverouteLicensesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetV2EverouteLicensesBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetV2EverouteLicensesBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-v2-everoute-licenses][%d] getV2EverouteLicensesBadRequest  %+v", 400, o.Payload)
}
func (o *GetV2EverouteLicensesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetV2EverouteLicensesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetV2EverouteLicensesNotFound creates a GetV2EverouteLicensesNotFound with default headers values
func NewGetV2EverouteLicensesNotFound() *GetV2EverouteLicensesNotFound {
	return &GetV2EverouteLicensesNotFound{}
}

/* GetV2EverouteLicensesNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetV2EverouteLicensesNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetV2EverouteLicensesNotFound) Error() string {
	return fmt.Sprintf("[POST /get-v2-everoute-licenses][%d] getV2EverouteLicensesNotFound  %+v", 404, o.Payload)
}
func (o *GetV2EverouteLicensesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetV2EverouteLicensesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetV2EverouteLicensesInternalServerError creates a GetV2EverouteLicensesInternalServerError with default headers values
func NewGetV2EverouteLicensesInternalServerError() *GetV2EverouteLicensesInternalServerError {
	return &GetV2EverouteLicensesInternalServerError{}
}

/* GetV2EverouteLicensesInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetV2EverouteLicensesInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetV2EverouteLicensesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-v2-everoute-licenses][%d] getV2EverouteLicensesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetV2EverouteLicensesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetV2EverouteLicensesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
