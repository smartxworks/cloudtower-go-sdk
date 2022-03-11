// Code generated by go-swagger; DO NOT EDIT.

package migrate_transmitter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// GetMigrateTransmittersReader is a Reader for the GetMigrateTransmitters structure.
type GetMigrateTransmittersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMigrateTransmittersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMigrateTransmittersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetMigrateTransmittersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetMigrateTransmittersOK creates a GetMigrateTransmittersOK with default headers values
func NewGetMigrateTransmittersOK() *GetMigrateTransmittersOK {
	return &GetMigrateTransmittersOK{}
}

/* GetMigrateTransmittersOK describes a response with status code 200, with default header values.

Ok
*/
type GetMigrateTransmittersOK struct {
	Payload []*models.MigrateTransmitter
}

func (o *GetMigrateTransmittersOK) Error() string {
	return fmt.Sprintf("[POST /get-migrate-transmitters][%d] getMigrateTransmittersOK  %+v", 200, o.Payload)
}
func (o *GetMigrateTransmittersOK) GetPayload() []*models.MigrateTransmitter {
	return o.Payload
}

func (o *GetMigrateTransmittersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMigrateTransmittersBadRequest creates a GetMigrateTransmittersBadRequest with default headers values
func NewGetMigrateTransmittersBadRequest() *GetMigrateTransmittersBadRequest {
	return &GetMigrateTransmittersBadRequest{}
}

/* GetMigrateTransmittersBadRequest describes a response with status code 400, with default header values.

GetMigrateTransmittersBadRequest get migrate transmitters bad request
*/
type GetMigrateTransmittersBadRequest struct {
	Payload string
}

func (o *GetMigrateTransmittersBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-migrate-transmitters][%d] getMigrateTransmittersBadRequest  %+v", 400, o.Payload)
}
func (o *GetMigrateTransmittersBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetMigrateTransmittersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
