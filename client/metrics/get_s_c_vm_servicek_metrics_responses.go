// Code generated by go-swagger; DO NOT EDIT.

package metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetSCVMServicekMetricsReader is a Reader for the GetSCVMServicekMetrics structure.
type GetSCVMServicekMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSCVMServicekMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSCVMServicekMetricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewGetSCVMServicekMetricsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewGetSCVMServicekMetricsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewGetSCVMServicekMetricsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSCVMServicekMetricsOK creates a GetSCVMServicekMetricsOK with default headers values
func NewGetSCVMServicekMetricsOK() *GetSCVMServicekMetricsOK {
	return &GetSCVMServicekMetricsOK{}
}

/* GetSCVMServicekMetricsOK describes a response with status code 200, with default header values.

GetSCVMServicekMetricsOK get s c Vm servicek metrics o k
*/
type GetSCVMServicekMetricsOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskMetric
}

func (o *GetSCVMServicekMetricsOK) Error() string {
	return fmt.Sprintf("[POST /get-scvm-service-metrics][%d] getSCVmServicekMetricsOK  %+v", 200, o.Payload)
}
func (o *GetSCVMServicekMetricsOK) GetPayload() []*models.WithTaskMetric {
	return o.Payload
}

func (o *GetSCVMServicekMetricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSCVMServicekMetricsBadRequest creates a GetSCVMServicekMetricsBadRequest with default headers values
func NewGetSCVMServicekMetricsBadRequest() *GetSCVMServicekMetricsBadRequest {
	return &GetSCVMServicekMetricsBadRequest{}
}

/* GetSCVMServicekMetricsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetSCVMServicekMetricsBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetSCVMServicekMetricsBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-scvm-service-metrics][%d] getSCVmServicekMetricsBadRequest  %+v", 400, o.Payload)
}
func (o *GetSCVMServicekMetricsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSCVMServicekMetricsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSCVMServicekMetricsNotFound creates a GetSCVMServicekMetricsNotFound with default headers values
func NewGetSCVMServicekMetricsNotFound() *GetSCVMServicekMetricsNotFound {
	return &GetSCVMServicekMetricsNotFound{}
}

/* GetSCVMServicekMetricsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetSCVMServicekMetricsNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetSCVMServicekMetricsNotFound) Error() string {
	return fmt.Sprintf("[POST /get-scvm-service-metrics][%d] getSCVmServicekMetricsNotFound  %+v", 404, o.Payload)
}
func (o *GetSCVMServicekMetricsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSCVMServicekMetricsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSCVMServicekMetricsInternalServerError creates a GetSCVMServicekMetricsInternalServerError with default headers values
func NewGetSCVMServicekMetricsInternalServerError() *GetSCVMServicekMetricsInternalServerError {
	return &GetSCVMServicekMetricsInternalServerError{}
}

/* GetSCVMServicekMetricsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetSCVMServicekMetricsInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetSCVMServicekMetricsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-scvm-service-metrics][%d] getSCVmServicekMetricsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetSCVMServicekMetricsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSCVMServicekMetricsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
