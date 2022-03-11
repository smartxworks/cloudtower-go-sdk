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

// GetDatacentersReader is a Reader for the GetDatacenters structure.
type GetDatacentersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDatacentersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDatacentersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDatacentersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDatacentersOK creates a GetDatacentersOK with default headers values
func NewGetDatacentersOK() *GetDatacentersOK {
	return &GetDatacentersOK{}
}

/* GetDatacentersOK describes a response with status code 200, with default header values.

Ok
*/
type GetDatacentersOK struct {
	Payload []*models.Datacenter
}

func (o *GetDatacentersOK) Error() string {
	return fmt.Sprintf("[POST /get-datacenters][%d] getDatacentersOK  %+v", 200, o.Payload)
}
func (o *GetDatacentersOK) GetPayload() []*models.Datacenter {
	return o.Payload
}

func (o *GetDatacentersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDatacentersBadRequest creates a GetDatacentersBadRequest with default headers values
func NewGetDatacentersBadRequest() *GetDatacentersBadRequest {
	return &GetDatacentersBadRequest{}
}

/* GetDatacentersBadRequest describes a response with status code 400, with default header values.

GetDatacentersBadRequest get datacenters bad request
*/
type GetDatacentersBadRequest struct {
	Payload string
}

func (o *GetDatacentersBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-datacenters][%d] getDatacentersBadRequest  %+v", 400, o.Payload)
}
func (o *GetDatacentersBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetDatacentersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
