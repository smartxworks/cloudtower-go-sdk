// Code generated by go-swagger; DO NOT EDIT.

package cloud_tower_application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// DeleteCloudTowerApplicationPackageReader is a Reader for the DeleteCloudTowerApplicationPackage structure.
type DeleteCloudTowerApplicationPackageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCloudTowerApplicationPackageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCloudTowerApplicationPackageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteCloudTowerApplicationPackageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteCloudTowerApplicationPackageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteCloudTowerApplicationPackageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteCloudTowerApplicationPackageOK creates a DeleteCloudTowerApplicationPackageOK with default headers values
func NewDeleteCloudTowerApplicationPackageOK() *DeleteCloudTowerApplicationPackageOK {
	return &DeleteCloudTowerApplicationPackageOK{}
}

/* DeleteCloudTowerApplicationPackageOK describes a response with status code 200, with default header values.

DeleteCloudTowerApplicationPackageOK delete cloud tower application package o k
*/
type DeleteCloudTowerApplicationPackageOK struct {
	XTowerRequestID string

	Payload []*models.DeleteCloudTowerApplicationPackage
}

func (o *DeleteCloudTowerApplicationPackageOK) Error() string {
	return fmt.Sprintf("[POST /delete-cloudtower-application-package][%d] deleteCloudTowerApplicationPackageOK  %+v", 200, o.Payload)
}
func (o *DeleteCloudTowerApplicationPackageOK) GetPayload() []*models.DeleteCloudTowerApplicationPackage {
	return o.Payload
}

func (o *DeleteCloudTowerApplicationPackageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteCloudTowerApplicationPackageBadRequest creates a DeleteCloudTowerApplicationPackageBadRequest with default headers values
func NewDeleteCloudTowerApplicationPackageBadRequest() *DeleteCloudTowerApplicationPackageBadRequest {
	return &DeleteCloudTowerApplicationPackageBadRequest{}
}

/* DeleteCloudTowerApplicationPackageBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteCloudTowerApplicationPackageBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteCloudTowerApplicationPackageBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-cloudtower-application-package][%d] deleteCloudTowerApplicationPackageBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteCloudTowerApplicationPackageBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteCloudTowerApplicationPackageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteCloudTowerApplicationPackageNotFound creates a DeleteCloudTowerApplicationPackageNotFound with default headers values
func NewDeleteCloudTowerApplicationPackageNotFound() *DeleteCloudTowerApplicationPackageNotFound {
	return &DeleteCloudTowerApplicationPackageNotFound{}
}

/* DeleteCloudTowerApplicationPackageNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteCloudTowerApplicationPackageNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteCloudTowerApplicationPackageNotFound) Error() string {
	return fmt.Sprintf("[POST /delete-cloudtower-application-package][%d] deleteCloudTowerApplicationPackageNotFound  %+v", 404, o.Payload)
}
func (o *DeleteCloudTowerApplicationPackageNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteCloudTowerApplicationPackageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteCloudTowerApplicationPackageInternalServerError creates a DeleteCloudTowerApplicationPackageInternalServerError with default headers values
func NewDeleteCloudTowerApplicationPackageInternalServerError() *DeleteCloudTowerApplicationPackageInternalServerError {
	return &DeleteCloudTowerApplicationPackageInternalServerError{}
}

/* DeleteCloudTowerApplicationPackageInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type DeleteCloudTowerApplicationPackageInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteCloudTowerApplicationPackageInternalServerError) Error() string {
	return fmt.Sprintf("[POST /delete-cloudtower-application-package][%d] deleteCloudTowerApplicationPackageInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteCloudTowerApplicationPackageInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteCloudTowerApplicationPackageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
