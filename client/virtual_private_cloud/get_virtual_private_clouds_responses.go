// Code generated by go-swagger; DO NOT EDIT.

package virtual_private_cloud

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetVirtualPrivateCloudsReader is a Reader for the GetVirtualPrivateClouds structure.
type GetVirtualPrivateCloudsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVirtualPrivateCloudsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVirtualPrivateCloudsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewGetVirtualPrivateCloudsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewGetVirtualPrivateCloudsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewGetVirtualPrivateCloudsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVirtualPrivateCloudsOK creates a GetVirtualPrivateCloudsOK with default headers values
func NewGetVirtualPrivateCloudsOK() *GetVirtualPrivateCloudsOK {
	return &GetVirtualPrivateCloudsOK{}
}

/* GetVirtualPrivateCloudsOK describes a response with status code 200, with default header values.

GetVirtualPrivateCloudsOK get virtual private clouds o k
*/
type GetVirtualPrivateCloudsOK struct {
	XTowerRequestID string

	Payload []*models.VirtualPrivateCloud
}

func (o *GetVirtualPrivateCloudsOK) Error() string {
	return fmt.Sprintf("[POST /get-virtual-private-clouds][%d] getVirtualPrivateCloudsOK  %+v", 200, o.Payload)
}
func (o *GetVirtualPrivateCloudsOK) GetPayload() []*models.VirtualPrivateCloud {
	return o.Payload
}

func (o *GetVirtualPrivateCloudsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVirtualPrivateCloudsBadRequest creates a GetVirtualPrivateCloudsBadRequest with default headers values
func NewGetVirtualPrivateCloudsBadRequest() *GetVirtualPrivateCloudsBadRequest {
	return &GetVirtualPrivateCloudsBadRequest{}
}

/* GetVirtualPrivateCloudsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetVirtualPrivateCloudsBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetVirtualPrivateCloudsBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-virtual-private-clouds][%d] getVirtualPrivateCloudsBadRequest  %+v", 400, o.Payload)
}
func (o *GetVirtualPrivateCloudsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVirtualPrivateCloudsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetVirtualPrivateCloudsNotFound creates a GetVirtualPrivateCloudsNotFound with default headers values
func NewGetVirtualPrivateCloudsNotFound() *GetVirtualPrivateCloudsNotFound {
	return &GetVirtualPrivateCloudsNotFound{}
}

/* GetVirtualPrivateCloudsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetVirtualPrivateCloudsNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetVirtualPrivateCloudsNotFound) Error() string {
	return fmt.Sprintf("[POST /get-virtual-private-clouds][%d] getVirtualPrivateCloudsNotFound  %+v", 404, o.Payload)
}
func (o *GetVirtualPrivateCloudsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVirtualPrivateCloudsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetVirtualPrivateCloudsInternalServerError creates a GetVirtualPrivateCloudsInternalServerError with default headers values
func NewGetVirtualPrivateCloudsInternalServerError() *GetVirtualPrivateCloudsInternalServerError {
	return &GetVirtualPrivateCloudsInternalServerError{}
}

/* GetVirtualPrivateCloudsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetVirtualPrivateCloudsInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetVirtualPrivateCloudsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-virtual-private-clouds][%d] getVirtualPrivateCloudsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetVirtualPrivateCloudsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVirtualPrivateCloudsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
