// Code generated by go-swagger; DO NOT EDIT.

package report_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// NewGenerateFromReportTemplateParams creates a new GenerateFromReportTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGenerateFromReportTemplateParams() *GenerateFromReportTemplateParams {
	return &GenerateFromReportTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGenerateFromReportTemplateParamsWithTimeout creates a new GenerateFromReportTemplateParams object
// with the ability to set a timeout on a request.
func NewGenerateFromReportTemplateParamsWithTimeout(timeout time.Duration) *GenerateFromReportTemplateParams {
	return &GenerateFromReportTemplateParams{
		timeout: timeout,
	}
}

// NewGenerateFromReportTemplateParamsWithContext creates a new GenerateFromReportTemplateParams object
// with the ability to set a context for a request.
func NewGenerateFromReportTemplateParamsWithContext(ctx context.Context) *GenerateFromReportTemplateParams {
	return &GenerateFromReportTemplateParams{
		Context: ctx,
	}
}

// NewGenerateFromReportTemplateParamsWithHTTPClient creates a new GenerateFromReportTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewGenerateFromReportTemplateParamsWithHTTPClient(client *http.Client) *GenerateFromReportTemplateParams {
	return &GenerateFromReportTemplateParams{
		HTTPClient: client,
	}
}

/* GenerateFromReportTemplateParams contains all the parameters to send to the API endpoint
   for the generate from report template operation.

   Typically these are written to a http.Request.
*/
type GenerateFromReportTemplateParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.ReporteTemplateGenerationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the generate from report template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GenerateFromReportTemplateParams) WithDefaults() *GenerateFromReportTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the generate from report template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GenerateFromReportTemplateParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GenerateFromReportTemplateParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the generate from report template params
func (o *GenerateFromReportTemplateParams) WithTimeout(timeout time.Duration) *GenerateFromReportTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the generate from report template params
func (o *GenerateFromReportTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the generate from report template params
func (o *GenerateFromReportTemplateParams) WithContext(ctx context.Context) *GenerateFromReportTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the generate from report template params
func (o *GenerateFromReportTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the generate from report template params
func (o *GenerateFromReportTemplateParams) WithHTTPClient(client *http.Client) *GenerateFromReportTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the generate from report template params
func (o *GenerateFromReportTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the generate from report template params
func (o *GenerateFromReportTemplateParams) WithContentLanguage(contentLanguage *string) *GenerateFromReportTemplateParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the generate from report template params
func (o *GenerateFromReportTemplateParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the generate from report template params
func (o *GenerateFromReportTemplateParams) WithRequestBody(requestBody *models.ReporteTemplateGenerationParams) *GenerateFromReportTemplateParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the generate from report template params
func (o *GenerateFromReportTemplateParams) SetRequestBody(requestBody *models.ReporteTemplateGenerationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GenerateFromReportTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContentLanguage != nil {

		// header param content-language
		if err := r.SetHeaderParam("content-language", *o.ContentLanguage); err != nil {
			return err
		}
	}
	if o.RequestBody != nil {
		if err := r.SetBodyParam(o.RequestBody); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
