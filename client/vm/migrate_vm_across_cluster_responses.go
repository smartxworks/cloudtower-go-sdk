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

// MigrateVMAcrossClusterReader is a Reader for the MigrateVMAcrossCluster structure.
type MigrateVMAcrossClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MigrateVMAcrossClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMigrateVMAcrossClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewMigrateVMAcrossClusterNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewMigrateVMAcrossClusterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMigrateVMAcrossClusterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewMigrateVMAcrossClusterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMigrateVMAcrossClusterOK creates a MigrateVMAcrossClusterOK with default headers values
func NewMigrateVMAcrossClusterOK() *MigrateVMAcrossClusterOK {
	return &MigrateVMAcrossClusterOK{}
}

/* MigrateVMAcrossClusterOK describes a response with status code 200, with default header values.

Ok
*/
type MigrateVMAcrossClusterOK struct {
	Payload []*models.WithTaskVM
}

func (o *MigrateVMAcrossClusterOK) Error() string {
	return fmt.Sprintf("[POST /migrate-vm-across-cluster][%d] migrateVmAcrossClusterOK  %+v", 200, o.Payload)
}
func (o *MigrateVMAcrossClusterOK) GetPayload() []*models.WithTaskVM {
	return o.Payload
}

func (o *MigrateVMAcrossClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMigrateVMAcrossClusterNotModified creates a MigrateVMAcrossClusterNotModified with default headers values
func NewMigrateVMAcrossClusterNotModified() *MigrateVMAcrossClusterNotModified {
	return &MigrateVMAcrossClusterNotModified{}
}

/* MigrateVMAcrossClusterNotModified describes a response with status code 304, with default header values.

Not modified
*/
type MigrateVMAcrossClusterNotModified struct {
}

func (o *MigrateVMAcrossClusterNotModified) Error() string {
	return fmt.Sprintf("[POST /migrate-vm-across-cluster][%d] migrateVmAcrossClusterNotModified ", 304)
}

func (o *MigrateVMAcrossClusterNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMigrateVMAcrossClusterBadRequest creates a MigrateVMAcrossClusterBadRequest with default headers values
func NewMigrateVMAcrossClusterBadRequest() *MigrateVMAcrossClusterBadRequest {
	return &MigrateVMAcrossClusterBadRequest{}
}

/* MigrateVMAcrossClusterBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type MigrateVMAcrossClusterBadRequest struct {
	Payload *models.ErrorBody
}

func (o *MigrateVMAcrossClusterBadRequest) Error() string {
	return fmt.Sprintf("[POST /migrate-vm-across-cluster][%d] migrateVmAcrossClusterBadRequest  %+v", 400, o.Payload)
}
func (o *MigrateVMAcrossClusterBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *MigrateVMAcrossClusterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMigrateVMAcrossClusterNotFound creates a MigrateVMAcrossClusterNotFound with default headers values
func NewMigrateVMAcrossClusterNotFound() *MigrateVMAcrossClusterNotFound {
	return &MigrateVMAcrossClusterNotFound{}
}

/* MigrateVMAcrossClusterNotFound describes a response with status code 404, with default header values.

Not found
*/
type MigrateVMAcrossClusterNotFound struct {
	Payload *models.ErrorBody
}

func (o *MigrateVMAcrossClusterNotFound) Error() string {
	return fmt.Sprintf("[POST /migrate-vm-across-cluster][%d] migrateVmAcrossClusterNotFound  %+v", 404, o.Payload)
}
func (o *MigrateVMAcrossClusterNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *MigrateVMAcrossClusterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMigrateVMAcrossClusterInternalServerError creates a MigrateVMAcrossClusterInternalServerError with default headers values
func NewMigrateVMAcrossClusterInternalServerError() *MigrateVMAcrossClusterInternalServerError {
	return &MigrateVMAcrossClusterInternalServerError{}
}

/* MigrateVMAcrossClusterInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type MigrateVMAcrossClusterInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *MigrateVMAcrossClusterInternalServerError) Error() string {
	return fmt.Sprintf("[POST /migrate-vm-across-cluster][%d] migrateVmAcrossClusterInternalServerError  %+v", 500, o.Payload)
}
func (o *MigrateVMAcrossClusterInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *MigrateVMAcrossClusterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}