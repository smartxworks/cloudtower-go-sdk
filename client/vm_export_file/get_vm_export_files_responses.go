// Code generated by go-swagger; DO NOT EDIT.

package vm_export_file

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetVMExportFilesReader is a Reader for the GetVMExportFiles structure.
type GetVMExportFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVMExportFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVMExportFilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVMExportFilesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVMExportFilesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVMExportFilesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVMExportFilesOK creates a GetVMExportFilesOK with default headers values
func NewGetVMExportFilesOK() *GetVMExportFilesOK {
	return &GetVMExportFilesOK{}
}

/* GetVMExportFilesOK describes a response with status code 200, with default header values.

GetVMExportFilesOK get Vm export files o k
*/
type GetVMExportFilesOK struct {
	XTowerRequestID string

	Payload []*models.VMExportFile
}

func (o *GetVMExportFilesOK) Error() string {
	return fmt.Sprintf("[POST /get-vm-export-files][%d] getVmExportFilesOK  %+v", 200, o.Payload)
}
func (o *GetVMExportFilesOK) GetPayload() []*models.VMExportFile {
	return o.Payload
}

func (o *GetVMExportFilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetVMExportFilesBadRequest creates a GetVMExportFilesBadRequest with default headers values
func NewGetVMExportFilesBadRequest() *GetVMExportFilesBadRequest {
	return &GetVMExportFilesBadRequest{}
}

/* GetVMExportFilesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetVMExportFilesBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetVMExportFilesBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-vm-export-files][%d] getVmExportFilesBadRequest  %+v", 400, o.Payload)
}
func (o *GetVMExportFilesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVMExportFilesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetVMExportFilesNotFound creates a GetVMExportFilesNotFound with default headers values
func NewGetVMExportFilesNotFound() *GetVMExportFilesNotFound {
	return &GetVMExportFilesNotFound{}
}

/* GetVMExportFilesNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetVMExportFilesNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetVMExportFilesNotFound) Error() string {
	return fmt.Sprintf("[POST /get-vm-export-files][%d] getVmExportFilesNotFound  %+v", 404, o.Payload)
}
func (o *GetVMExportFilesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVMExportFilesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetVMExportFilesInternalServerError creates a GetVMExportFilesInternalServerError with default headers values
func NewGetVMExportFilesInternalServerError() *GetVMExportFilesInternalServerError {
	return &GetVMExportFilesInternalServerError{}
}

/* GetVMExportFilesInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetVMExportFilesInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetVMExportFilesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-vm-export-files][%d] getVmExportFilesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetVMExportFilesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVMExportFilesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
