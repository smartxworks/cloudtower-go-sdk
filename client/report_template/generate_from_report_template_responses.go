// Code generated by go-swagger; DO NOT EDIT.

package report_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GenerateFromReportTemplateReader is a Reader for the GenerateFromReportTemplate structure.
type GenerateFromReportTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenerateFromReportTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGenerateFromReportTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return result, nil
	case 400:
		result := NewGenerateFromReportTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 404:
		result := NewGenerateFromReportTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	case 500:
		result := NewGenerateFromReportTemplateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, models.NewUnexpectedError(response, err)
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGenerateFromReportTemplateOK creates a GenerateFromReportTemplateOK with default headers values
func NewGenerateFromReportTemplateOK() *GenerateFromReportTemplateOK {
	return &GenerateFromReportTemplateOK{}
}

/* GenerateFromReportTemplateOK describes a response with status code 200, with default header values.

GenerateFromReportTemplateOK generate from report template o k
*/
type GenerateFromReportTemplateOK struct {
	XTowerRequestID string

	Payload []*models.WithTaskReportTask
}

func (o *GenerateFromReportTemplateOK) Error() string {
	return fmt.Sprintf("[POST /generate-from-report-template][%d] generateFromReportTemplateOK  %+v", 200, o.Payload)
}
func (o *GenerateFromReportTemplateOK) GetPayload() []*models.WithTaskReportTask {
	return o.Payload
}

func (o *GenerateFromReportTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGenerateFromReportTemplateBadRequest creates a GenerateFromReportTemplateBadRequest with default headers values
func NewGenerateFromReportTemplateBadRequest() *GenerateFromReportTemplateBadRequest {
	return &GenerateFromReportTemplateBadRequest{}
}

/* GenerateFromReportTemplateBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GenerateFromReportTemplateBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GenerateFromReportTemplateBadRequest) Error() string {
	return fmt.Sprintf("[POST /generate-from-report-template][%d] generateFromReportTemplateBadRequest  %+v", 400, o.Payload)
}
func (o *GenerateFromReportTemplateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GenerateFromReportTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGenerateFromReportTemplateNotFound creates a GenerateFromReportTemplateNotFound with default headers values
func NewGenerateFromReportTemplateNotFound() *GenerateFromReportTemplateNotFound {
	return &GenerateFromReportTemplateNotFound{}
}

/* GenerateFromReportTemplateNotFound describes a response with status code 404, with default header values.

Not found
*/
type GenerateFromReportTemplateNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GenerateFromReportTemplateNotFound) Error() string {
	return fmt.Sprintf("[POST /generate-from-report-template][%d] generateFromReportTemplateNotFound  %+v", 404, o.Payload)
}
func (o *GenerateFromReportTemplateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GenerateFromReportTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGenerateFromReportTemplateInternalServerError creates a GenerateFromReportTemplateInternalServerError with default headers values
func NewGenerateFromReportTemplateInternalServerError() *GenerateFromReportTemplateInternalServerError {
	return &GenerateFromReportTemplateInternalServerError{}
}

/* GenerateFromReportTemplateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GenerateFromReportTemplateInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GenerateFromReportTemplateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /generate-from-report-template][%d] generateFromReportTemplateInternalServerError  %+v", 500, o.Payload)
}
func (o *GenerateFromReportTemplateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GenerateFromReportTemplateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
