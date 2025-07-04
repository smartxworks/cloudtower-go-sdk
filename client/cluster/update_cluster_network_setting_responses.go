// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// UpdateClusterNetworkSettingReader is a Reader for the UpdateClusterNetworkSetting structure.
type UpdateClusterNetworkSettingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateClusterNetworkSettingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateClusterNetworkSettingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewUpdateClusterNetworkSettingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewUpdateClusterNetworkSettingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewUpdateClusterNetworkSettingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateClusterNetworkSettingOK creates a UpdateClusterNetworkSettingOK with default headers values
func NewUpdateClusterNetworkSettingOK() *UpdateClusterNetworkSettingOK {
	return &UpdateClusterNetworkSettingOK{}
}

/* UpdateClusterNetworkSettingOK describes a response with status code 200, with default header values.

UpdateClusterNetworkSettingOK update cluster network setting o k
*/
type UpdateClusterNetworkSettingOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskCluster
}

func (o *UpdateClusterNetworkSettingOK) Error() string {
	return fmt.Sprintf("[POST /update-cluster-network-setting][%d] updateClusterNetworkSettingOK  %+v", 200, o.Payload)
}
func (o *UpdateClusterNetworkSettingOK) GetPayload() []*models.WithTaskCluster {
	return o.Payload
}

func (o *UpdateClusterNetworkSettingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateClusterNetworkSettingBadRequest creates a UpdateClusterNetworkSettingBadRequest with default headers values
func NewUpdateClusterNetworkSettingBadRequest() *UpdateClusterNetworkSettingBadRequest {
	return &UpdateClusterNetworkSettingBadRequest{}
}

/* UpdateClusterNetworkSettingBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateClusterNetworkSettingBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateClusterNetworkSettingBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-cluster-network-setting][%d] updateClusterNetworkSettingBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateClusterNetworkSettingBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateClusterNetworkSettingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateClusterNetworkSettingNotFound creates a UpdateClusterNetworkSettingNotFound with default headers values
func NewUpdateClusterNetworkSettingNotFound() *UpdateClusterNetworkSettingNotFound {
	return &UpdateClusterNetworkSettingNotFound{}
}

/* UpdateClusterNetworkSettingNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateClusterNetworkSettingNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateClusterNetworkSettingNotFound) Error() string {
	return fmt.Sprintf("[POST /update-cluster-network-setting][%d] updateClusterNetworkSettingNotFound  %+v", 404, o.Payload)
}
func (o *UpdateClusterNetworkSettingNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateClusterNetworkSettingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateClusterNetworkSettingInternalServerError creates a UpdateClusterNetworkSettingInternalServerError with default headers values
func NewUpdateClusterNetworkSettingInternalServerError() *UpdateClusterNetworkSettingInternalServerError {
	return &UpdateClusterNetworkSettingInternalServerError{}
}

/* UpdateClusterNetworkSettingInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UpdateClusterNetworkSettingInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *UpdateClusterNetworkSettingInternalServerError) Error() string {
	return fmt.Sprintf("[POST /update-cluster-network-setting][%d] updateClusterNetworkSettingInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateClusterNetworkSettingInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateClusterNetworkSettingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
