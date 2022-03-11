// Code generated by go-swagger; DO NOT EDIT.

package namespace_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// GetNamespaceGroupsConnectionReader is a Reader for the GetNamespaceGroupsConnection structure.
type GetNamespaceGroupsConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNamespaceGroupsConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNamespaceGroupsConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetNamespaceGroupsConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNamespaceGroupsConnectionOK creates a GetNamespaceGroupsConnectionOK with default headers values
func NewGetNamespaceGroupsConnectionOK() *GetNamespaceGroupsConnectionOK {
	return &GetNamespaceGroupsConnectionOK{}
}

/* GetNamespaceGroupsConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetNamespaceGroupsConnectionOK struct {
	Payload *models.NamespaceGroupConnection
}

func (o *GetNamespaceGroupsConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-namespace-groups-connection][%d] getNamespaceGroupsConnectionOK  %+v", 200, o.Payload)
}
func (o *GetNamespaceGroupsConnectionOK) GetPayload() *models.NamespaceGroupConnection {
	return o.Payload
}

func (o *GetNamespaceGroupsConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NamespaceGroupConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNamespaceGroupsConnectionBadRequest creates a GetNamespaceGroupsConnectionBadRequest with default headers values
func NewGetNamespaceGroupsConnectionBadRequest() *GetNamespaceGroupsConnectionBadRequest {
	return &GetNamespaceGroupsConnectionBadRequest{}
}

/* GetNamespaceGroupsConnectionBadRequest describes a response with status code 400, with default header values.

GetNamespaceGroupsConnectionBadRequest get namespace groups connection bad request
*/
type GetNamespaceGroupsConnectionBadRequest struct {
	Payload string
}

func (o *GetNamespaceGroupsConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-namespace-groups-connection][%d] getNamespaceGroupsConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetNamespaceGroupsConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetNamespaceGroupsConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
