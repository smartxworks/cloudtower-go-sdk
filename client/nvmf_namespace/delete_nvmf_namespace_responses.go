// Code generated by go-swagger; DO NOT EDIT.

package nvmf_namespace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// DeleteNvmfNamespaceReader is a Reader for the DeleteNvmfNamespace structure.
type DeleteNvmfNamespaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNvmfNamespaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteNvmfNamespaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewDeleteNvmfNamespaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewDeleteNvmfNamespaceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewDeleteNvmfNamespaceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteNvmfNamespaceOK creates a DeleteNvmfNamespaceOK with default headers values
func NewDeleteNvmfNamespaceOK() *DeleteNvmfNamespaceOK {
	return &DeleteNvmfNamespaceOK{}
}

/* DeleteNvmfNamespaceOK describes a response with status code 200, with default header values.

DeleteNvmfNamespaceOK delete nvmf namespace o k
*/
type DeleteNvmfNamespaceOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskDeleteNvmfNamespace
}

func (o *DeleteNvmfNamespaceOK) Error() string {
	return fmt.Sprintf("[POST /delete-nvmf-namespace][%d] deleteNvmfNamespaceOK  %+v", 200, o.Payload)
}
func (o *DeleteNvmfNamespaceOK) GetPayload() []*models.WithTaskDeleteNvmfNamespace {
	return o.Payload
}

func (o *DeleteNvmfNamespaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteNvmfNamespaceBadRequest creates a DeleteNvmfNamespaceBadRequest with default headers values
func NewDeleteNvmfNamespaceBadRequest() *DeleteNvmfNamespaceBadRequest {
	return &DeleteNvmfNamespaceBadRequest{}
}

/* DeleteNvmfNamespaceBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteNvmfNamespaceBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteNvmfNamespaceBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-nvmf-namespace][%d] deleteNvmfNamespaceBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteNvmfNamespaceBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteNvmfNamespaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteNvmfNamespaceNotFound creates a DeleteNvmfNamespaceNotFound with default headers values
func NewDeleteNvmfNamespaceNotFound() *DeleteNvmfNamespaceNotFound {
	return &DeleteNvmfNamespaceNotFound{}
}

/* DeleteNvmfNamespaceNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteNvmfNamespaceNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteNvmfNamespaceNotFound) Error() string {
	return fmt.Sprintf("[POST /delete-nvmf-namespace][%d] deleteNvmfNamespaceNotFound  %+v", 404, o.Payload)
}
func (o *DeleteNvmfNamespaceNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteNvmfNamespaceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteNvmfNamespaceInternalServerError creates a DeleteNvmfNamespaceInternalServerError with default headers values
func NewDeleteNvmfNamespaceInternalServerError() *DeleteNvmfNamespaceInternalServerError {
	return &DeleteNvmfNamespaceInternalServerError{}
}

/* DeleteNvmfNamespaceInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type DeleteNvmfNamespaceInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *DeleteNvmfNamespaceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /delete-nvmf-namespace][%d] deleteNvmfNamespaceInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteNvmfNamespaceInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteNvmfNamespaceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
