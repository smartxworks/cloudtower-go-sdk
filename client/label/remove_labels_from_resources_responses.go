// Code generated by go-swagger; DO NOT EDIT.

package label

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// RemoveLabelsFromResourcesReader is a Reader for the RemoveLabelsFromResources structure.
type RemoveLabelsFromResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveLabelsFromResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveLabelsFromResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 304:
		result := NewRemoveLabelsFromResourcesNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 400:
		result := NewRemoveLabelsFromResourcesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewRemoveLabelsFromResourcesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewRemoveLabelsFromResourcesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRemoveLabelsFromResourcesOK creates a RemoveLabelsFromResourcesOK with default headers values
func NewRemoveLabelsFromResourcesOK() *RemoveLabelsFromResourcesOK {
	return &RemoveLabelsFromResourcesOK{}
}

/* RemoveLabelsFromResourcesOK describes a response with status code 200, with default header values.

RemoveLabelsFromResourcesOK remove labels from resources o k
*/
type RemoveLabelsFromResourcesOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskLabel
}

func (o *RemoveLabelsFromResourcesOK) Error() string {
	return fmt.Sprintf("[POST /remove-labels-from-resources][%d] removeLabelsFromResourcesOK  %+v", 200, o.Payload)
}
func (o *RemoveLabelsFromResourcesOK) GetPayload() []*models.WithTaskLabel {
	return o.Payload
}

func (o *RemoveLabelsFromResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRemoveLabelsFromResourcesNotModified creates a RemoveLabelsFromResourcesNotModified with default headers values
func NewRemoveLabelsFromResourcesNotModified() *RemoveLabelsFromResourcesNotModified {
	return &RemoveLabelsFromResourcesNotModified{}
}

/* RemoveLabelsFromResourcesNotModified describes a response with status code 304, with default header values.

Not modified
*/
type RemoveLabelsFromResourcesNotModified struct {
}

func (o *RemoveLabelsFromResourcesNotModified) Error() string {
	return fmt.Sprintf("[POST /remove-labels-from-resources][%d] removeLabelsFromResourcesNotModified ", 304)
}

func (o *RemoveLabelsFromResourcesNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveLabelsFromResourcesBadRequest creates a RemoveLabelsFromResourcesBadRequest with default headers values
func NewRemoveLabelsFromResourcesBadRequest() *RemoveLabelsFromResourcesBadRequest {
	return &RemoveLabelsFromResourcesBadRequest{}
}

/* RemoveLabelsFromResourcesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type RemoveLabelsFromResourcesBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *RemoveLabelsFromResourcesBadRequest) Error() string {
	return fmt.Sprintf("[POST /remove-labels-from-resources][%d] removeLabelsFromResourcesBadRequest  %+v", 400, o.Payload)
}
func (o *RemoveLabelsFromResourcesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *RemoveLabelsFromResourcesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRemoveLabelsFromResourcesNotFound creates a RemoveLabelsFromResourcesNotFound with default headers values
func NewRemoveLabelsFromResourcesNotFound() *RemoveLabelsFromResourcesNotFound {
	return &RemoveLabelsFromResourcesNotFound{}
}

/* RemoveLabelsFromResourcesNotFound describes a response with status code 404, with default header values.

Not found
*/
type RemoveLabelsFromResourcesNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *RemoveLabelsFromResourcesNotFound) Error() string {
	return fmt.Sprintf("[POST /remove-labels-from-resources][%d] removeLabelsFromResourcesNotFound  %+v", 404, o.Payload)
}
func (o *RemoveLabelsFromResourcesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *RemoveLabelsFromResourcesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRemoveLabelsFromResourcesInternalServerError creates a RemoveLabelsFromResourcesInternalServerError with default headers values
func NewRemoveLabelsFromResourcesInternalServerError() *RemoveLabelsFromResourcesInternalServerError {
	return &RemoveLabelsFromResourcesInternalServerError{}
}

/* RemoveLabelsFromResourcesInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type RemoveLabelsFromResourcesInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *RemoveLabelsFromResourcesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /remove-labels-from-resources][%d] removeLabelsFromResourcesInternalServerError  %+v", 500, o.Payload)
}
func (o *RemoveLabelsFromResourcesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *RemoveLabelsFromResourcesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
