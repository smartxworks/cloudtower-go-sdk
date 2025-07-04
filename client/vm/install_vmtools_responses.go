// Code generated by go-swagger; DO NOT EDIT.

package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// InstallVmtoolsReader is a Reader for the InstallVmtools structure.
type InstallVmtoolsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InstallVmtoolsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInstallVmtoolsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewInstallVmtoolsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewInstallVmtoolsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewInstallVmtoolsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInstallVmtoolsOK creates a InstallVmtoolsOK with default headers values
func NewInstallVmtoolsOK() *InstallVmtoolsOK {
	return &InstallVmtoolsOK{}
}

/* InstallVmtoolsOK describes a response with status code 200, with default header values.

InstallVmtoolsOK install vmtools o k
*/
type InstallVmtoolsOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskVM
}

func (o *InstallVmtoolsOK) Error() string {
	return fmt.Sprintf("[POST /install-vmtools][%d] installVmtoolsOK  %+v", 200, o.Payload)
}
func (o *InstallVmtoolsOK) GetPayload() []*models.WithTaskVM {
	return o.Payload
}

func (o *InstallVmtoolsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewInstallVmtoolsBadRequest creates a InstallVmtoolsBadRequest with default headers values
func NewInstallVmtoolsBadRequest() *InstallVmtoolsBadRequest {
	return &InstallVmtoolsBadRequest{}
}

/* InstallVmtoolsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type InstallVmtoolsBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *InstallVmtoolsBadRequest) Error() string {
	return fmt.Sprintf("[POST /install-vmtools][%d] installVmtoolsBadRequest  %+v", 400, o.Payload)
}
func (o *InstallVmtoolsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *InstallVmtoolsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewInstallVmtoolsNotFound creates a InstallVmtoolsNotFound with default headers values
func NewInstallVmtoolsNotFound() *InstallVmtoolsNotFound {
	return &InstallVmtoolsNotFound{}
}

/* InstallVmtoolsNotFound describes a response with status code 404, with default header values.

Not found
*/
type InstallVmtoolsNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *InstallVmtoolsNotFound) Error() string {
	return fmt.Sprintf("[POST /install-vmtools][%d] installVmtoolsNotFound  %+v", 404, o.Payload)
}
func (o *InstallVmtoolsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *InstallVmtoolsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewInstallVmtoolsInternalServerError creates a InstallVmtoolsInternalServerError with default headers values
func NewInstallVmtoolsInternalServerError() *InstallVmtoolsInternalServerError {
	return &InstallVmtoolsInternalServerError{}
}

/* InstallVmtoolsInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type InstallVmtoolsInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *InstallVmtoolsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /install-vmtools][%d] installVmtoolsInternalServerError  %+v", 500, o.Payload)
}
func (o *InstallVmtoolsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *InstallVmtoolsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
