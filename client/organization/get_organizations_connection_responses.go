// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// GetOrganizationsConnectionReader is a Reader for the GetOrganizationsConnection structure.
type GetOrganizationsConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationsConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationsConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrganizationsConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationsConnectionOK creates a GetOrganizationsConnectionOK with default headers values
func NewGetOrganizationsConnectionOK() *GetOrganizationsConnectionOK {
	return &GetOrganizationsConnectionOK{}
}

/* GetOrganizationsConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetOrganizationsConnectionOK struct {
	Payload *models.OrganizationConnection
}

func (o *GetOrganizationsConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-organizations-connection][%d] getOrganizationsConnectionOK  %+v", 200, o.Payload)
}
func (o *GetOrganizationsConnectionOK) GetPayload() *models.OrganizationConnection {
	return o.Payload
}

func (o *GetOrganizationsConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrganizationConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsConnectionBadRequest creates a GetOrganizationsConnectionBadRequest with default headers values
func NewGetOrganizationsConnectionBadRequest() *GetOrganizationsConnectionBadRequest {
	return &GetOrganizationsConnectionBadRequest{}
}

/* GetOrganizationsConnectionBadRequest describes a response with status code 400, with default header values.

GetOrganizationsConnectionBadRequest get organizations connection bad request
*/
type GetOrganizationsConnectionBadRequest struct {
	Payload string
}

func (o *GetOrganizationsConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-organizations-connection][%d] getOrganizationsConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetOrganizationsConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetOrganizationsConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
