// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// UpdateClusterLicenseReader is a Reader for the UpdateClusterLicense structure.
type UpdateClusterLicenseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateClusterLicenseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateClusterLicenseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateClusterLicenseBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateClusterLicenseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateClusterLicenseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateClusterLicenseOK creates a UpdateClusterLicenseOK with default headers values
func NewUpdateClusterLicenseOK() *UpdateClusterLicenseOK {
	return &UpdateClusterLicenseOK{}
}

/* UpdateClusterLicenseOK describes a response with status code 200, with default header values.

Ok
*/
type UpdateClusterLicenseOK struct {
	Payload []*models.WithTaskCluster
}

func (o *UpdateClusterLicenseOK) Error() string {
	return fmt.Sprintf("[POST /update-cluster-license][%d] updateClusterLicenseOK  %+v", 200, o.Payload)
}
func (o *UpdateClusterLicenseOK) GetPayload() []*models.WithTaskCluster {
	return o.Payload
}

func (o *UpdateClusterLicenseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClusterLicenseBadRequest creates a UpdateClusterLicenseBadRequest with default headers values
func NewUpdateClusterLicenseBadRequest() *UpdateClusterLicenseBadRequest {
	return &UpdateClusterLicenseBadRequest{}
}

/* UpdateClusterLicenseBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateClusterLicenseBadRequest struct {
	Payload *models.ErrorBody
}

func (o *UpdateClusterLicenseBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-cluster-license][%d] updateClusterLicenseBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateClusterLicenseBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateClusterLicenseBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClusterLicenseNotFound creates a UpdateClusterLicenseNotFound with default headers values
func NewUpdateClusterLicenseNotFound() *UpdateClusterLicenseNotFound {
	return &UpdateClusterLicenseNotFound{}
}

/* UpdateClusterLicenseNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateClusterLicenseNotFound struct {
	Payload *models.ErrorBody
}

func (o *UpdateClusterLicenseNotFound) Error() string {
	return fmt.Sprintf("[POST /update-cluster-license][%d] updateClusterLicenseNotFound  %+v", 404, o.Payload)
}
func (o *UpdateClusterLicenseNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateClusterLicenseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClusterLicenseInternalServerError creates a UpdateClusterLicenseInternalServerError with default headers values
func NewUpdateClusterLicenseInternalServerError() *UpdateClusterLicenseInternalServerError {
	return &UpdateClusterLicenseInternalServerError{}
}

/* UpdateClusterLicenseInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UpdateClusterLicenseInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *UpdateClusterLicenseInternalServerError) Error() string {
	return fmt.Sprintf("[POST /update-cluster-license][%d] updateClusterLicenseInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateClusterLicenseInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UpdateClusterLicenseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
