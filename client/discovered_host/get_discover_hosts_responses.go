// Code generated by go-swagger; DO NOT EDIT.

package discovered_host

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetDiscoverHostsReader is a Reader for the GetDiscoverHosts structure.
type GetDiscoverHostsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDiscoverHostsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDiscoverHostsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewGetDiscoverHostsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewGetDiscoverHostsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewGetDiscoverHostsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDiscoverHostsOK creates a GetDiscoverHostsOK with default headers values
func NewGetDiscoverHostsOK() *GetDiscoverHostsOK {
	return &GetDiscoverHostsOK{}
}

/* GetDiscoverHostsOK describes a response with status code 200, with default header values.

GetDiscoverHostsOK get discover hosts o k
*/
type GetDiscoverHostsOK struct {
	XTowerRequestID string

	Payload []*models.DiscoveredHost
}

func (o *GetDiscoverHostsOK) Error() string {
	return fmt.Sprintf("[POST /get-discover-hosts][%d] getDiscoverHostsOK  %+v", 200, o.Payload)
}
func (o *GetDiscoverHostsOK) GetPayload() []*models.DiscoveredHost {
	return o.Payload
}

func (o *GetDiscoverHostsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDiscoverHostsBadRequest creates a GetDiscoverHostsBadRequest with default headers values
func NewGetDiscoverHostsBadRequest() *GetDiscoverHostsBadRequest {
	return &GetDiscoverHostsBadRequest{}
}

/* GetDiscoverHostsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetDiscoverHostsBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetDiscoverHostsBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-discover-hosts][%d] getDiscoverHostsBadRequest  %+v", 400, o.Payload)
}
func (o *GetDiscoverHostsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDiscoverHostsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDiscoverHostsNotFound creates a GetDiscoverHostsNotFound with default headers values
func NewGetDiscoverHostsNotFound() *GetDiscoverHostsNotFound {
	return &GetDiscoverHostsNotFound{}
}

/* GetDiscoverHostsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetDiscoverHostsNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetDiscoverHostsNotFound) Error() string {
	return fmt.Sprintf("[POST /get-discover-hosts][%d] getDiscoverHostsNotFound  %+v", 404, o.Payload)
}
func (o *GetDiscoverHostsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDiscoverHostsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDiscoverHostsInternalServerError creates a GetDiscoverHostsInternalServerError with default headers values
func NewGetDiscoverHostsInternalServerError() *GetDiscoverHostsInternalServerError {
	return &GetDiscoverHostsInternalServerError{}
}

/* GetDiscoverHostsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetDiscoverHostsInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetDiscoverHostsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-discover-hosts][%d] getDiscoverHostsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDiscoverHostsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDiscoverHostsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
