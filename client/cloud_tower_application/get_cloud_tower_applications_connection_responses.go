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

// GetCloudTowerApplicationsConnectionReader is a Reader for the GetCloudTowerApplicationsConnection structure.
type GetCloudTowerApplicationsConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCloudTowerApplicationsConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCloudTowerApplicationsConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCloudTowerApplicationsConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCloudTowerApplicationsConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCloudTowerApplicationsConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCloudTowerApplicationsConnectionOK creates a GetCloudTowerApplicationsConnectionOK with default headers values
func NewGetCloudTowerApplicationsConnectionOK() *GetCloudTowerApplicationsConnectionOK {
	return &GetCloudTowerApplicationsConnectionOK{}
}

/* GetCloudTowerApplicationsConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetCloudTowerApplicationsConnectionOK struct {
	Payload *models.CloudTowerApplicationConnection
}

func (o *GetCloudTowerApplicationsConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-cloud-tower-applications-connection][%d] getCloudTowerApplicationsConnectionOK  %+v", 200, o.Payload)
}
func (o *GetCloudTowerApplicationsConnectionOK) GetPayload() *models.CloudTowerApplicationConnection {
	return o.Payload
}

func (o *GetCloudTowerApplicationsConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudTowerApplicationConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudTowerApplicationsConnectionBadRequest creates a GetCloudTowerApplicationsConnectionBadRequest with default headers values
func NewGetCloudTowerApplicationsConnectionBadRequest() *GetCloudTowerApplicationsConnectionBadRequest {
	return &GetCloudTowerApplicationsConnectionBadRequest{}
}

/* GetCloudTowerApplicationsConnectionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetCloudTowerApplicationsConnectionBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetCloudTowerApplicationsConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-cloud-tower-applications-connection][%d] getCloudTowerApplicationsConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetCloudTowerApplicationsConnectionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCloudTowerApplicationsConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudTowerApplicationsConnectionNotFound creates a GetCloudTowerApplicationsConnectionNotFound with default headers values
func NewGetCloudTowerApplicationsConnectionNotFound() *GetCloudTowerApplicationsConnectionNotFound {
	return &GetCloudTowerApplicationsConnectionNotFound{}
}

/* GetCloudTowerApplicationsConnectionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetCloudTowerApplicationsConnectionNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetCloudTowerApplicationsConnectionNotFound) Error() string {
	return fmt.Sprintf("[POST /get-cloud-tower-applications-connection][%d] getCloudTowerApplicationsConnectionNotFound  %+v", 404, o.Payload)
}
func (o *GetCloudTowerApplicationsConnectionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCloudTowerApplicationsConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudTowerApplicationsConnectionInternalServerError creates a GetCloudTowerApplicationsConnectionInternalServerError with default headers values
func NewGetCloudTowerApplicationsConnectionInternalServerError() *GetCloudTowerApplicationsConnectionInternalServerError {
	return &GetCloudTowerApplicationsConnectionInternalServerError{}
}

/* GetCloudTowerApplicationsConnectionInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetCloudTowerApplicationsConnectionInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetCloudTowerApplicationsConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-cloud-tower-applications-connection][%d] getCloudTowerApplicationsConnectionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCloudTowerApplicationsConnectionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetCloudTowerApplicationsConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}