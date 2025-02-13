// Code generated by go-swagger; DO NOT EDIT.

package virtual_private_cloud_floating_ip

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// BatchCreateVirtualPrivateCloudFloatingIpsReader is a Reader for the BatchCreateVirtualPrivateCloudFloatingIps structure.
type BatchCreateVirtualPrivateCloudFloatingIpsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BatchCreateVirtualPrivateCloudFloatingIpsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBatchCreateVirtualPrivateCloudFloatingIpsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBatchCreateVirtualPrivateCloudFloatingIpsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBatchCreateVirtualPrivateCloudFloatingIpsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBatchCreateVirtualPrivateCloudFloatingIpsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBatchCreateVirtualPrivateCloudFloatingIpsOK creates a BatchCreateVirtualPrivateCloudFloatingIpsOK with default headers values
func NewBatchCreateVirtualPrivateCloudFloatingIpsOK() *BatchCreateVirtualPrivateCloudFloatingIpsOK {
	return &BatchCreateVirtualPrivateCloudFloatingIpsOK{}
}

/* BatchCreateVirtualPrivateCloudFloatingIpsOK describes a response with status code 200, with default header values.

BatchCreateVirtualPrivateCloudFloatingIpsOK batch create virtual private cloud floating ips o k
*/
type BatchCreateVirtualPrivateCloudFloatingIpsOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskVirtualPrivateCloudFloatingIP
}

func (o *BatchCreateVirtualPrivateCloudFloatingIpsOK) Error() string {
	return fmt.Sprintf("[POST /batch-create-virtual-private-cloud-floating-ips][%d] batchCreateVirtualPrivateCloudFloatingIpsOK  %+v", 200, o.Payload)
}
func (o *BatchCreateVirtualPrivateCloudFloatingIpsOK) GetPayload() []*models.WithTaskVirtualPrivateCloudFloatingIP {
	return o.Payload
}

func (o *BatchCreateVirtualPrivateCloudFloatingIpsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewBatchCreateVirtualPrivateCloudFloatingIpsBadRequest creates a BatchCreateVirtualPrivateCloudFloatingIpsBadRequest with default headers values
func NewBatchCreateVirtualPrivateCloudFloatingIpsBadRequest() *BatchCreateVirtualPrivateCloudFloatingIpsBadRequest {
	return &BatchCreateVirtualPrivateCloudFloatingIpsBadRequest{}
}

/* BatchCreateVirtualPrivateCloudFloatingIpsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type BatchCreateVirtualPrivateCloudFloatingIpsBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *BatchCreateVirtualPrivateCloudFloatingIpsBadRequest) Error() string {
	return fmt.Sprintf("[POST /batch-create-virtual-private-cloud-floating-ips][%d] batchCreateVirtualPrivateCloudFloatingIpsBadRequest  %+v", 400, o.Payload)
}
func (o *BatchCreateVirtualPrivateCloudFloatingIpsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *BatchCreateVirtualPrivateCloudFloatingIpsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewBatchCreateVirtualPrivateCloudFloatingIpsNotFound creates a BatchCreateVirtualPrivateCloudFloatingIpsNotFound with default headers values
func NewBatchCreateVirtualPrivateCloudFloatingIpsNotFound() *BatchCreateVirtualPrivateCloudFloatingIpsNotFound {
	return &BatchCreateVirtualPrivateCloudFloatingIpsNotFound{}
}

/* BatchCreateVirtualPrivateCloudFloatingIpsNotFound describes a response with status code 404, with default header values.

Not found
*/
type BatchCreateVirtualPrivateCloudFloatingIpsNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *BatchCreateVirtualPrivateCloudFloatingIpsNotFound) Error() string {
	return fmt.Sprintf("[POST /batch-create-virtual-private-cloud-floating-ips][%d] batchCreateVirtualPrivateCloudFloatingIpsNotFound  %+v", 404, o.Payload)
}
func (o *BatchCreateVirtualPrivateCloudFloatingIpsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *BatchCreateVirtualPrivateCloudFloatingIpsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewBatchCreateVirtualPrivateCloudFloatingIpsInternalServerError creates a BatchCreateVirtualPrivateCloudFloatingIpsInternalServerError with default headers values
func NewBatchCreateVirtualPrivateCloudFloatingIpsInternalServerError() *BatchCreateVirtualPrivateCloudFloatingIpsInternalServerError {
	return &BatchCreateVirtualPrivateCloudFloatingIpsInternalServerError{}
}

/* BatchCreateVirtualPrivateCloudFloatingIpsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type BatchCreateVirtualPrivateCloudFloatingIpsInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *BatchCreateVirtualPrivateCloudFloatingIpsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /batch-create-virtual-private-cloud-floating-ips][%d] batchCreateVirtualPrivateCloudFloatingIpsInternalServerError  %+v", 500, o.Payload)
}
func (o *BatchCreateVirtualPrivateCloudFloatingIpsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *BatchCreateVirtualPrivateCloudFloatingIpsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
