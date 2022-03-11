// Code generated by go-swagger; DO NOT EDIT.

package everoute_package

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// GetEveroutePackagesConnectionReader is a Reader for the GetEveroutePackagesConnection structure.
type GetEveroutePackagesConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEveroutePackagesConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEveroutePackagesConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetEveroutePackagesConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEveroutePackagesConnectionOK creates a GetEveroutePackagesConnectionOK with default headers values
func NewGetEveroutePackagesConnectionOK() *GetEveroutePackagesConnectionOK {
	return &GetEveroutePackagesConnectionOK{}
}

/* GetEveroutePackagesConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetEveroutePackagesConnectionOK struct {
	Payload *models.EveroutePackageConnection
}

func (o *GetEveroutePackagesConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-everoute-packages-connection][%d] getEveroutePackagesConnectionOK  %+v", 200, o.Payload)
}
func (o *GetEveroutePackagesConnectionOK) GetPayload() *models.EveroutePackageConnection {
	return o.Payload
}

func (o *GetEveroutePackagesConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EveroutePackageConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEveroutePackagesConnectionBadRequest creates a GetEveroutePackagesConnectionBadRequest with default headers values
func NewGetEveroutePackagesConnectionBadRequest() *GetEveroutePackagesConnectionBadRequest {
	return &GetEveroutePackagesConnectionBadRequest{}
}

/* GetEveroutePackagesConnectionBadRequest describes a response with status code 400, with default header values.

GetEveroutePackagesConnectionBadRequest get everoute packages connection bad request
*/
type GetEveroutePackagesConnectionBadRequest struct {
	Payload string
}

func (o *GetEveroutePackagesConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-everoute-packages-connection][%d] getEveroutePackagesConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetEveroutePackagesConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetEveroutePackagesConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
