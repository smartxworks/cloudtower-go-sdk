// Code generated by go-swagger; DO NOT EDIT.

package everoute_cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetEverouteClustersReader is a Reader for the GetEverouteClusters structure.
type GetEverouteClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEverouteClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEverouteClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewGetEverouteClustersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewGetEverouteClustersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewGetEverouteClustersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEverouteClustersOK creates a GetEverouteClustersOK with default headers values
func NewGetEverouteClustersOK() *GetEverouteClustersOK {
	return &GetEverouteClustersOK{}
}

/* GetEverouteClustersOK describes a response with status code 200, with default header values.

GetEverouteClustersOK get everoute clusters o k
*/
type GetEverouteClustersOK struct {
	XTowerRequestID string

	Payload []*models.EverouteCluster
}

func (o *GetEverouteClustersOK) Error() string {
	return fmt.Sprintf("[POST /get-everoute-clusters][%d] getEverouteClustersOK  %+v", 200, o.Payload)
}
func (o *GetEverouteClustersOK) GetPayload() []*models.EverouteCluster {
	return o.Payload
}

func (o *GetEverouteClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetEverouteClustersBadRequest creates a GetEverouteClustersBadRequest with default headers values
func NewGetEverouteClustersBadRequest() *GetEverouteClustersBadRequest {
	return &GetEverouteClustersBadRequest{}
}

/* GetEverouteClustersBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetEverouteClustersBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetEverouteClustersBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-everoute-clusters][%d] getEverouteClustersBadRequest  %+v", 400, o.Payload)
}
func (o *GetEverouteClustersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetEverouteClustersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetEverouteClustersNotFound creates a GetEverouteClustersNotFound with default headers values
func NewGetEverouteClustersNotFound() *GetEverouteClustersNotFound {
	return &GetEverouteClustersNotFound{}
}

/* GetEverouteClustersNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetEverouteClustersNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetEverouteClustersNotFound) Error() string {
	return fmt.Sprintf("[POST /get-everoute-clusters][%d] getEverouteClustersNotFound  %+v", 404, o.Payload)
}
func (o *GetEverouteClustersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetEverouteClustersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetEverouteClustersInternalServerError creates a GetEverouteClustersInternalServerError with default headers values
func NewGetEverouteClustersInternalServerError() *GetEverouteClustersInternalServerError {
	return &GetEverouteClustersInternalServerError{}
}

/* GetEverouteClustersInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetEverouteClustersInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetEverouteClustersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-everoute-clusters][%d] getEverouteClustersInternalServerError  %+v", 500, o.Payload)
}
func (o *GetEverouteClustersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetEverouteClustersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
