// Code generated by go-swagger; DO NOT EDIT.

package virtual_private_cloud_router_gateway

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// UpdateVirtualPrivateCloudRouterGatewayReader is a Reader for the UpdateVirtualPrivateCloudRouterGateway structure.
type UpdateVirtualPrivateCloudRouterGatewayReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVirtualPrivateCloudRouterGatewayReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateVirtualPrivateCloudRouterGatewayOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateVirtualPrivateCloudRouterGatewayBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateVirtualPrivateCloudRouterGatewayNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateVirtualPrivateCloudRouterGatewayInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateVirtualPrivateCloudRouterGatewayOK creates a UpdateVirtualPrivateCloudRouterGatewayOK with default headers values
func NewUpdateVirtualPrivateCloudRouterGatewayOK() *UpdateVirtualPrivateCloudRouterGatewayOK {
	return &UpdateVirtualPrivateCloudRouterGatewayOK{}
}

/* UpdateVirtualPrivateCloudRouterGatewayOK describes a response with status code 200, with default header values.

UpdateVirtualPrivateCloudRouterGatewayOK update virtual private cloud router gateway o k
*/
type UpdateVirtualPrivateCloudRouterGatewayOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskVirtualPrivateCloudRouterGateway
}

func (o *UpdateVirtualPrivateCloudRouterGatewayOK) Error() string {
	return fmt.Sprintf("[POST /update-virtual-private-cloud-router-gateway][%d] updateVirtualPrivateCloudRouterGatewayOK  %+v", 200, o.Payload)
}
func (o *UpdateVirtualPrivateCloudRouterGatewayOK) GetPayload() []*models.WithTaskVirtualPrivateCloudRouterGateway {
	return o.Payload
}

func (o *UpdateVirtualPrivateCloudRouterGatewayOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateVirtualPrivateCloudRouterGatewayBadRequest creates a UpdateVirtualPrivateCloudRouterGatewayBadRequest with default headers values
func NewUpdateVirtualPrivateCloudRouterGatewayBadRequest() *UpdateVirtualPrivateCloudRouterGatewayBadRequest {
	return &UpdateVirtualPrivateCloudRouterGatewayBadRequest{}
}

/* UpdateVirtualPrivateCloudRouterGatewayBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateVirtualPrivateCloudRouterGatewayBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateVirtualPrivateCloudRouterGatewayBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-virtual-private-cloud-router-gateway][%d] updateVirtualPrivateCloudRouterGatewayBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateVirtualPrivateCloudRouterGatewayBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateVirtualPrivateCloudRouterGatewayBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateVirtualPrivateCloudRouterGatewayNotFound creates a UpdateVirtualPrivateCloudRouterGatewayNotFound with default headers values
func NewUpdateVirtualPrivateCloudRouterGatewayNotFound() *UpdateVirtualPrivateCloudRouterGatewayNotFound {
	return &UpdateVirtualPrivateCloudRouterGatewayNotFound{}
}

/* UpdateVirtualPrivateCloudRouterGatewayNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateVirtualPrivateCloudRouterGatewayNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateVirtualPrivateCloudRouterGatewayNotFound) Error() string {
	return fmt.Sprintf("[POST /update-virtual-private-cloud-router-gateway][%d] updateVirtualPrivateCloudRouterGatewayNotFound  %+v", 404, o.Payload)
}
func (o *UpdateVirtualPrivateCloudRouterGatewayNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateVirtualPrivateCloudRouterGatewayNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateVirtualPrivateCloudRouterGatewayInternalServerError creates a UpdateVirtualPrivateCloudRouterGatewayInternalServerError with default headers values
func NewUpdateVirtualPrivateCloudRouterGatewayInternalServerError() *UpdateVirtualPrivateCloudRouterGatewayInternalServerError {
	return &UpdateVirtualPrivateCloudRouterGatewayInternalServerError{}
}

/* UpdateVirtualPrivateCloudRouterGatewayInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UpdateVirtualPrivateCloudRouterGatewayInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateVirtualPrivateCloudRouterGatewayInternalServerError) Error() string {
	return fmt.Sprintf("[POST /update-virtual-private-cloud-router-gateway][%d] updateVirtualPrivateCloudRouterGatewayInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateVirtualPrivateCloudRouterGatewayInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateVirtualPrivateCloudRouterGatewayInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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