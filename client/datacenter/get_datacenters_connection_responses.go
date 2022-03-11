// Code generated by go-swagger; DO NOT EDIT.

package datacenter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// GetDatacentersConnectionReader is a Reader for the GetDatacentersConnection structure.
type GetDatacentersConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDatacentersConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDatacentersConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDatacentersConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDatacentersConnectionOK creates a GetDatacentersConnectionOK with default headers values
func NewGetDatacentersConnectionOK() *GetDatacentersConnectionOK {
	return &GetDatacentersConnectionOK{}
}

/* GetDatacentersConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetDatacentersConnectionOK struct {
	Payload *models.DatacenterConnection
}

func (o *GetDatacentersConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-datacenters-connection][%d] getDatacentersConnectionOK  %+v", 200, o.Payload)
}
func (o *GetDatacentersConnectionOK) GetPayload() *models.DatacenterConnection {
	return o.Payload
}

func (o *GetDatacentersConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DatacenterConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatacentersConnectionBadRequest creates a GetDatacentersConnectionBadRequest with default headers values
func NewGetDatacentersConnectionBadRequest() *GetDatacentersConnectionBadRequest {
	return &GetDatacentersConnectionBadRequest{}
}

/* GetDatacentersConnectionBadRequest describes a response with status code 400, with default header values.

GetDatacentersConnectionBadRequest get datacenters connection bad request
*/
type GetDatacentersConnectionBadRequest struct {
	Payload string
}

func (o *GetDatacentersConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-datacenters-connection][%d] getDatacentersConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetDatacentersConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetDatacentersConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
