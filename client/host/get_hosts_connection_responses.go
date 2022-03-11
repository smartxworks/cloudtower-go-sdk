// Code generated by go-swagger; DO NOT EDIT.

package host

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// GetHostsConnectionReader is a Reader for the GetHostsConnection structure.
type GetHostsConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHostsConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHostsConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetHostsConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetHostsConnectionOK creates a GetHostsConnectionOK with default headers values
func NewGetHostsConnectionOK() *GetHostsConnectionOK {
	return &GetHostsConnectionOK{}
}

/* GetHostsConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetHostsConnectionOK struct {
	Payload *models.HostConnection
}

func (o *GetHostsConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-hosts-connection][%d] getHostsConnectionOK  %+v", 200, o.Payload)
}
func (o *GetHostsConnectionOK) GetPayload() *models.HostConnection {
	return o.Payload
}

func (o *GetHostsConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HostConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHostsConnectionBadRequest creates a GetHostsConnectionBadRequest with default headers values
func NewGetHostsConnectionBadRequest() *GetHostsConnectionBadRequest {
	return &GetHostsConnectionBadRequest{}
}

/* GetHostsConnectionBadRequest describes a response with status code 400, with default header values.

GetHostsConnectionBadRequest get hosts connection bad request
*/
type GetHostsConnectionBadRequest struct {
	Payload string
}

func (o *GetHostsConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-hosts-connection][%d] getHostsConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetHostsConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetHostsConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
