// Code generated by go-swagger; DO NOT EDIT.

package virtual_private_cloud_edge_gateway

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetVirtualPrivateCloudEdgeGatewaysConnectionReader is a Reader for the GetVirtualPrivateCloudEdgeGatewaysConnection structure.
type GetVirtualPrivateCloudEdgeGatewaysConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVirtualPrivateCloudEdgeGatewaysConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVirtualPrivateCloudEdgeGatewaysConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewGetVirtualPrivateCloudEdgeGatewaysConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewGetVirtualPrivateCloudEdgeGatewaysConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewGetVirtualPrivateCloudEdgeGatewaysConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVirtualPrivateCloudEdgeGatewaysConnectionOK creates a GetVirtualPrivateCloudEdgeGatewaysConnectionOK with default headers values
func NewGetVirtualPrivateCloudEdgeGatewaysConnectionOK() *GetVirtualPrivateCloudEdgeGatewaysConnectionOK {
	return &GetVirtualPrivateCloudEdgeGatewaysConnectionOK{}
}

/* GetVirtualPrivateCloudEdgeGatewaysConnectionOK describes a response with status code 200, with default header values.

GetVirtualPrivateCloudEdgeGatewaysConnectionOK get virtual private cloud edge gateways connection o k
*/
type GetVirtualPrivateCloudEdgeGatewaysConnectionOK struct {
	XTowerRequestID string

	Payload *models.VirtualPrivateCloudEdgeGatewayConnection
}

func (o *GetVirtualPrivateCloudEdgeGatewaysConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-virtual-private-cloud-edge-gateways-connection][%d] getVirtualPrivateCloudEdgeGatewaysConnectionOK  %+v", 200, o.Payload)
}
func (o *GetVirtualPrivateCloudEdgeGatewaysConnectionOK) GetPayload() *models.VirtualPrivateCloudEdgeGatewayConnection {
	return o.Payload
}

func (o *GetVirtualPrivateCloudEdgeGatewaysConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.VirtualPrivateCloudEdgeGatewayConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVirtualPrivateCloudEdgeGatewaysConnectionBadRequest creates a GetVirtualPrivateCloudEdgeGatewaysConnectionBadRequest with default headers values
func NewGetVirtualPrivateCloudEdgeGatewaysConnectionBadRequest() *GetVirtualPrivateCloudEdgeGatewaysConnectionBadRequest {
	return &GetVirtualPrivateCloudEdgeGatewaysConnectionBadRequest{}
}

/* GetVirtualPrivateCloudEdgeGatewaysConnectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetVirtualPrivateCloudEdgeGatewaysConnectionBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetVirtualPrivateCloudEdgeGatewaysConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-virtual-private-cloud-edge-gateways-connection][%d] getVirtualPrivateCloudEdgeGatewaysConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetVirtualPrivateCloudEdgeGatewaysConnectionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVirtualPrivateCloudEdgeGatewaysConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetVirtualPrivateCloudEdgeGatewaysConnectionNotFound creates a GetVirtualPrivateCloudEdgeGatewaysConnectionNotFound with default headers values
func NewGetVirtualPrivateCloudEdgeGatewaysConnectionNotFound() *GetVirtualPrivateCloudEdgeGatewaysConnectionNotFound {
	return &GetVirtualPrivateCloudEdgeGatewaysConnectionNotFound{}
}

/* GetVirtualPrivateCloudEdgeGatewaysConnectionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetVirtualPrivateCloudEdgeGatewaysConnectionNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetVirtualPrivateCloudEdgeGatewaysConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /get-virtual-private-cloud-edge-gateways-connection][%d] getVirtualPrivateCloudEdgeGatewaysConnectionNotFound  %+v", 404, o.Payload)
}
func (o *GetVirtualPrivateCloudEdgeGatewaysConnectionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVirtualPrivateCloudEdgeGatewaysConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetVirtualPrivateCloudEdgeGatewaysConnectionInternalServerError creates a GetVirtualPrivateCloudEdgeGatewaysConnectionInternalServerError with default headers values
func NewGetVirtualPrivateCloudEdgeGatewaysConnectionInternalServerError() *GetVirtualPrivateCloudEdgeGatewaysConnectionInternalServerError {
	return &GetVirtualPrivateCloudEdgeGatewaysConnectionInternalServerError{}
}

/* GetVirtualPrivateCloudEdgeGatewaysConnectionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetVirtualPrivateCloudEdgeGatewaysConnectionInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetVirtualPrivateCloudEdgeGatewaysConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-virtual-private-cloud-edge-gateways-connection][%d] getVirtualPrivateCloudEdgeGatewaysConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetVirtualPrivateCloudEdgeGatewaysConnectionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVirtualPrivateCloudEdgeGatewaysConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
