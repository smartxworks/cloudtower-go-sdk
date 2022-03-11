// Code generated by go-swagger; DO NOT EDIT.

package user_role_next

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// GetUserRoleNextsReader is a Reader for the GetUserRoleNexts structure.
type GetUserRoleNextsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserRoleNextsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserRoleNextsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserRoleNextsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserRoleNextsOK creates a GetUserRoleNextsOK with default headers values
func NewGetUserRoleNextsOK() *GetUserRoleNextsOK {
	return &GetUserRoleNextsOK{}
}

/* GetUserRoleNextsOK describes a response with status code 200, with default header values.

Ok
*/
type GetUserRoleNextsOK struct {
	Payload []*models.UserRoleNext
}

func (o *GetUserRoleNextsOK) Error() string {
	return fmt.Sprintf("[POST /get-user-role-nexts][%d] getUserRoleNextsOK  %+v", 200, o.Payload)
}
func (o *GetUserRoleNextsOK) GetPayload() []*models.UserRoleNext {
	return o.Payload
}

func (o *GetUserRoleNextsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRoleNextsBadRequest creates a GetUserRoleNextsBadRequest with default headers values
func NewGetUserRoleNextsBadRequest() *GetUserRoleNextsBadRequest {
	return &GetUserRoleNextsBadRequest{}
}

/* GetUserRoleNextsBadRequest describes a response with status code 400, with default header values.

GetUserRoleNextsBadRequest get user role nexts bad request
*/
type GetUserRoleNextsBadRequest struct {
	Payload string
}

func (o *GetUserRoleNextsBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-user-role-nexts][%d] getUserRoleNextsBadRequest  %+v", 400, o.Payload)
}
func (o *GetUserRoleNextsBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetUserRoleNextsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
