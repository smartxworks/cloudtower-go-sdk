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

// UninstallCloudTowerApplicationReader is a Reader for the UninstallCloudTowerApplication structure.
type UninstallCloudTowerApplicationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UninstallCloudTowerApplicationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUninstallCloudTowerApplicationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUninstallCloudTowerApplicationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUninstallCloudTowerApplicationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUninstallCloudTowerApplicationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUninstallCloudTowerApplicationOK creates a UninstallCloudTowerApplicationOK with default headers values
func NewUninstallCloudTowerApplicationOK() *UninstallCloudTowerApplicationOK {
	return &UninstallCloudTowerApplicationOK{}
}

/* UninstallCloudTowerApplicationOK describes a response with status code 200, with default header values.

Ok
*/
type UninstallCloudTowerApplicationOK struct {
	Payload []*models.CloudTowerApplication
}

func (o *UninstallCloudTowerApplicationOK) Error() string {
	return fmt.Sprintf("[POST /uninstall-cloudtower-application][%d] uninstallCloudTowerApplicationOK  %+v", 200, o.Payload)
}
func (o *UninstallCloudTowerApplicationOK) GetPayload() []*models.CloudTowerApplication {
	return o.Payload
}

func (o *UninstallCloudTowerApplicationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUninstallCloudTowerApplicationBadRequest creates a UninstallCloudTowerApplicationBadRequest with default headers values
func NewUninstallCloudTowerApplicationBadRequest() *UninstallCloudTowerApplicationBadRequest {
	return &UninstallCloudTowerApplicationBadRequest{}
}

/* UninstallCloudTowerApplicationBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UninstallCloudTowerApplicationBadRequest struct {
	Payload *models.ErrorBody
}

func (o *UninstallCloudTowerApplicationBadRequest) Error() string {
	return fmt.Sprintf("[POST /uninstall-cloudtower-application][%d] uninstallCloudTowerApplicationBadRequest  %+v", 400, o.Payload)
}
func (o *UninstallCloudTowerApplicationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UninstallCloudTowerApplicationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUninstallCloudTowerApplicationNotFound creates a UninstallCloudTowerApplicationNotFound with default headers values
func NewUninstallCloudTowerApplicationNotFound() *UninstallCloudTowerApplicationNotFound {
	return &UninstallCloudTowerApplicationNotFound{}
}

/* UninstallCloudTowerApplicationNotFound describes a response with status code 404, with default header values.

Not found
*/
type UninstallCloudTowerApplicationNotFound struct {
	Payload *models.ErrorBody
}

func (o *UninstallCloudTowerApplicationNotFound) Error() string {
	return fmt.Sprintf("[POST /uninstall-cloudtower-application][%d] uninstallCloudTowerApplicationNotFound  %+v", 404, o.Payload)
}
func (o *UninstallCloudTowerApplicationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UninstallCloudTowerApplicationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUninstallCloudTowerApplicationInternalServerError creates a UninstallCloudTowerApplicationInternalServerError with default headers values
func NewUninstallCloudTowerApplicationInternalServerError() *UninstallCloudTowerApplicationInternalServerError {
	return &UninstallCloudTowerApplicationInternalServerError{}
}

/* UninstallCloudTowerApplicationInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type UninstallCloudTowerApplicationInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *UninstallCloudTowerApplicationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /uninstall-cloudtower-application][%d] uninstallCloudTowerApplicationInternalServerError  %+v", 500, o.Payload)
}
func (o *UninstallCloudTowerApplicationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *UninstallCloudTowerApplicationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
