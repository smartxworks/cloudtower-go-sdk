// Code generated by go-swagger; DO NOT EDIT.

package cloud_tower_application_package

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetCloudTowerApplicationPackagesReader is a Reader for the GetCloudTowerApplicationPackages structure.
type GetCloudTowerApplicationPackagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCloudTowerApplicationPackagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCloudTowerApplicationPackagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCloudTowerApplicationPackagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCloudTowerApplicationPackagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCloudTowerApplicationPackagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCloudTowerApplicationPackagesOK creates a GetCloudTowerApplicationPackagesOK with default headers values
func NewGetCloudTowerApplicationPackagesOK() *GetCloudTowerApplicationPackagesOK {
	return &GetCloudTowerApplicationPackagesOK{}
}

/* GetCloudTowerApplicationPackagesOK describes a response with status code 200, with default header values.

GetCloudTowerApplicationPackagesOK get cloud tower application packages o k
*/
type GetCloudTowerApplicationPackagesOK struct {
	XTowerRequestID string

	Payload []*models.CloudTowerApplicationPackage
}

func (o *GetCloudTowerApplicationPackagesOK) Error() string {
	return fmt.Sprintf("[POST /get-cloudtower-application-packages][%d] getCloudTowerApplicationPackagesOK  %+v", 200, o.Payload)
}
func (o *GetCloudTowerApplicationPackagesOK) GetPayload() []*models.CloudTowerApplicationPackage {
	return o.Payload
}

func (o *GetCloudTowerApplicationPackagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCloudTowerApplicationPackagesBadRequest creates a GetCloudTowerApplicationPackagesBadRequest with default headers values
func NewGetCloudTowerApplicationPackagesBadRequest() *GetCloudTowerApplicationPackagesBadRequest {
	return &GetCloudTowerApplicationPackagesBadRequest{}
}

/* GetCloudTowerApplicationPackagesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCloudTowerApplicationPackagesBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetCloudTowerApplicationPackagesBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-cloudtower-application-packages][%d] getCloudTowerApplicationPackagesBadRequest  %+v", 400, o.Payload)
}
func (o *GetCloudTowerApplicationPackagesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCloudTowerApplicationPackagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCloudTowerApplicationPackagesNotFound creates a GetCloudTowerApplicationPackagesNotFound with default headers values
func NewGetCloudTowerApplicationPackagesNotFound() *GetCloudTowerApplicationPackagesNotFound {
	return &GetCloudTowerApplicationPackagesNotFound{}
}

/* GetCloudTowerApplicationPackagesNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetCloudTowerApplicationPackagesNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetCloudTowerApplicationPackagesNotFound) Error() string {
	return fmt.Sprintf("[POST /get-cloudtower-application-packages][%d] getCloudTowerApplicationPackagesNotFound  %+v", 404, o.Payload)
}
func (o *GetCloudTowerApplicationPackagesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCloudTowerApplicationPackagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCloudTowerApplicationPackagesInternalServerError creates a GetCloudTowerApplicationPackagesInternalServerError with default headers values
func NewGetCloudTowerApplicationPackagesInternalServerError() *GetCloudTowerApplicationPackagesInternalServerError {
	return &GetCloudTowerApplicationPackagesInternalServerError{}
}

/* GetCloudTowerApplicationPackagesInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetCloudTowerApplicationPackagesInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetCloudTowerApplicationPackagesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-cloudtower-application-packages][%d] getCloudTowerApplicationPackagesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCloudTowerApplicationPackagesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCloudTowerApplicationPackagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
