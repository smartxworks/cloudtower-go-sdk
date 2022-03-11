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

// CreateNamespaceGroupReader is a Reader for the CreateNamespaceGroup structure.
type CreateNamespaceGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNamespaceGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateNamespaceGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateNamespaceGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateNamespaceGroupOK creates a CreateNamespaceGroupOK with default headers values
func NewCreateNamespaceGroupOK() *CreateNamespaceGroupOK {
	return &CreateNamespaceGroupOK{}
}

/* CreateNamespaceGroupOK describes a response with status code 200, with default header values.

Ok
*/
type CreateNamespaceGroupOK struct {
	Payload []*models.WithTaskNamespaceGroup
}

func (o *CreateNamespaceGroupOK) Error() string {
	return fmt.Sprintf("[POST /create-namespace-group][%d] createNamespaceGroupOK  %+v", 200, o.Payload)
}
func (o *CreateNamespaceGroupOK) GetPayload() []*models.WithTaskNamespaceGroup {
	return o.Payload
}

func (o *CreateNamespaceGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNamespaceGroupBadRequest creates a CreateNamespaceGroupBadRequest with default headers values
func NewCreateNamespaceGroupBadRequest() *CreateNamespaceGroupBadRequest {
	return &CreateNamespaceGroupBadRequest{}
}

/* CreateNamespaceGroupBadRequest describes a response with status code 400, with default header values.

CreateNamespaceGroupBadRequest create namespace group bad request
*/
type CreateNamespaceGroupBadRequest struct {
	Payload string
}

func (o *CreateNamespaceGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-namespace-group][%d] createNamespaceGroupBadRequest  %+v", 400, o.Payload)
}
func (o *CreateNamespaceGroupBadRequest) GetPayload() string {
	return o.Payload
}

func (o *CreateNamespaceGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
