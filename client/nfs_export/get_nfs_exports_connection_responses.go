// Code generated by go-swagger; DO NOT EDIT.

package nfs_export

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetNfsExportsConnectionReader is a Reader for the GetNfsExportsConnection structure.
type GetNfsExportsConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNfsExportsConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNfsExportsConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewGetNfsExportsConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewGetNfsExportsConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewGetNfsExportsConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNfsExportsConnectionOK creates a GetNfsExportsConnectionOK with default headers values
func NewGetNfsExportsConnectionOK() *GetNfsExportsConnectionOK {
	return &GetNfsExportsConnectionOK{}
}

/* GetNfsExportsConnectionOK describes a response with status code 200, with default header values.

GetNfsExportsConnectionOK get nfs exports connection o k
*/
type GetNfsExportsConnectionOK struct {
	XTowerRequestID string

	Payload *models.NfsExportConnection
}

func (o *GetNfsExportsConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-nfs-exports-connection][%d] getNfsExportsConnectionOK  %+v", 200, o.Payload)
}
func (o *GetNfsExportsConnectionOK) GetPayload() *models.NfsExportConnection {
	return o.Payload
}

func (o *GetNfsExportsConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.NfsExportConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNfsExportsConnectionBadRequest creates a GetNfsExportsConnectionBadRequest with default headers values
func NewGetNfsExportsConnectionBadRequest() *GetNfsExportsConnectionBadRequest {
	return &GetNfsExportsConnectionBadRequest{}
}

/* GetNfsExportsConnectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetNfsExportsConnectionBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetNfsExportsConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-nfs-exports-connection][%d] getNfsExportsConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetNfsExportsConnectionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNfsExportsConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetNfsExportsConnectionNotFound creates a GetNfsExportsConnectionNotFound with default headers values
func NewGetNfsExportsConnectionNotFound() *GetNfsExportsConnectionNotFound {
	return &GetNfsExportsConnectionNotFound{}
}

/* GetNfsExportsConnectionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetNfsExportsConnectionNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetNfsExportsConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /get-nfs-exports-connection][%d] getNfsExportsConnectionNotFound  %+v", 404, o.Payload)
}
func (o *GetNfsExportsConnectionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNfsExportsConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetNfsExportsConnectionInternalServerError creates a GetNfsExportsConnectionInternalServerError with default headers values
func NewGetNfsExportsConnectionInternalServerError() *GetNfsExportsConnectionInternalServerError {
	return &GetNfsExportsConnectionInternalServerError{}
}

/* GetNfsExportsConnectionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetNfsExportsConnectionInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetNfsExportsConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-nfs-exports-connection][%d] getNfsExportsConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetNfsExportsConnectionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNfsExportsConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
