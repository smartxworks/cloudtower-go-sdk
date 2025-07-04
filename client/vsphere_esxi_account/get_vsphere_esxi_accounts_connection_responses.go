// Code generated by go-swagger; DO NOT EDIT.

package vsphere_esxi_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetVsphereEsxiAccountsConnectionReader is a Reader for the GetVsphereEsxiAccountsConnection structure.
type GetVsphereEsxiAccountsConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVsphereEsxiAccountsConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVsphereEsxiAccountsConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewGetVsphereEsxiAccountsConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewGetVsphereEsxiAccountsConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewGetVsphereEsxiAccountsConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVsphereEsxiAccountsConnectionOK creates a GetVsphereEsxiAccountsConnectionOK with default headers values
func NewGetVsphereEsxiAccountsConnectionOK() *GetVsphereEsxiAccountsConnectionOK {
	return &GetVsphereEsxiAccountsConnectionOK{}
}

/* GetVsphereEsxiAccountsConnectionOK describes a response with status code 200, with default header values.

GetVsphereEsxiAccountsConnectionOK get vsphere esxi accounts connection o k
*/
type GetVsphereEsxiAccountsConnectionOK struct {
	XTowerRequestID string

	Payload *models.VsphereEsxiAccountConnection
}

func (o *GetVsphereEsxiAccountsConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-vsphere-esxi-accounts-connection][%d] getVsphereEsxiAccountsConnectionOK  %+v", 200, o.Payload)
}
func (o *GetVsphereEsxiAccountsConnectionOK) GetPayload() *models.VsphereEsxiAccountConnection {
	return o.Payload
}

func (o *GetVsphereEsxiAccountsConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.VsphereEsxiAccountConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVsphereEsxiAccountsConnectionBadRequest creates a GetVsphereEsxiAccountsConnectionBadRequest with default headers values
func NewGetVsphereEsxiAccountsConnectionBadRequest() *GetVsphereEsxiAccountsConnectionBadRequest {
	return &GetVsphereEsxiAccountsConnectionBadRequest{}
}

/* GetVsphereEsxiAccountsConnectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetVsphereEsxiAccountsConnectionBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetVsphereEsxiAccountsConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-vsphere-esxi-accounts-connection][%d] getVsphereEsxiAccountsConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetVsphereEsxiAccountsConnectionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVsphereEsxiAccountsConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetVsphereEsxiAccountsConnectionNotFound creates a GetVsphereEsxiAccountsConnectionNotFound with default headers values
func NewGetVsphereEsxiAccountsConnectionNotFound() *GetVsphereEsxiAccountsConnectionNotFound {
	return &GetVsphereEsxiAccountsConnectionNotFound{}
}

/* GetVsphereEsxiAccountsConnectionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetVsphereEsxiAccountsConnectionNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetVsphereEsxiAccountsConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /get-vsphere-esxi-accounts-connection][%d] getVsphereEsxiAccountsConnectionNotFound  %+v", 404, o.Payload)
}
func (o *GetVsphereEsxiAccountsConnectionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVsphereEsxiAccountsConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetVsphereEsxiAccountsConnectionInternalServerError creates a GetVsphereEsxiAccountsConnectionInternalServerError with default headers values
func NewGetVsphereEsxiAccountsConnectionInternalServerError() *GetVsphereEsxiAccountsConnectionInternalServerError {
	return &GetVsphereEsxiAccountsConnectionInternalServerError{}
}

/* GetVsphereEsxiAccountsConnectionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetVsphereEsxiAccountsConnectionInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetVsphereEsxiAccountsConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-vsphere-esxi-accounts-connection][%d] getVsphereEsxiAccountsConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetVsphereEsxiAccountsConnectionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVsphereEsxiAccountsConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
