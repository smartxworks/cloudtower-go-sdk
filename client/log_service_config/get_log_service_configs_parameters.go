// Code generated by go-swagger; DO NOT EDIT.

package log_service_config

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

// NewGetLogServiceConfigsParams creates a new GetLogServiceConfigsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLogServiceConfigsParams() *GetLogServiceConfigsParams {
	return &GetLogServiceConfigsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLogServiceConfigsParamsWithTimeout creates a new GetLogServiceConfigsParams object
// with the ability to set a timeout on a request.
func NewGetLogServiceConfigsParamsWithTimeout(timeout time.Duration) *GetLogServiceConfigsParams {
	return &GetLogServiceConfigsParams{
		timeout: timeout,
	}
}

// NewGetLogServiceConfigsParamsWithContext creates a new GetLogServiceConfigsParams object
// with the ability to set a context for a request.
func NewGetLogServiceConfigsParamsWithContext(ctx context.Context) *GetLogServiceConfigsParams {
	return &GetLogServiceConfigsParams{
		Context: ctx,
	}
}

// NewGetLogServiceConfigsParamsWithHTTPClient creates a new GetLogServiceConfigsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLogServiceConfigsParamsWithHTTPClient(client *http.Client) *GetLogServiceConfigsParams {
	return &GetLogServiceConfigsParams{
		HTTPClient: client,
	}
}

/* GetLogServiceConfigsParams contains all the parameters to send to the API endpoint
   for the get log service configs operation.

   Typically these are written to a http.Request.
*/
type GetLogServiceConfigsParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetLogServiceConfigsRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get log service configs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLogServiceConfigsParams) WithDefaults() *GetLogServiceConfigsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get log service configs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLogServiceConfigsParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetLogServiceConfigsParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get log service configs params
func (o *GetLogServiceConfigsParams) WithTimeout(timeout time.Duration) *GetLogServiceConfigsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get log service configs params
func (o *GetLogServiceConfigsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get log service configs params
func (o *GetLogServiceConfigsParams) WithContext(ctx context.Context) *GetLogServiceConfigsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get log service configs params
func (o *GetLogServiceConfigsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get log service configs params
func (o *GetLogServiceConfigsParams) WithHTTPClient(client *http.Client) *GetLogServiceConfigsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get log service configs params
func (o *GetLogServiceConfigsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get log service configs params
func (o *GetLogServiceConfigsParams) WithContentLanguage(contentLanguage *string) *GetLogServiceConfigsParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get log service configs params
func (o *GetLogServiceConfigsParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get log service configs params
func (o *GetLogServiceConfigsParams) WithRequestBody(requestBody *models.GetLogServiceConfigsRequestBody) *GetLogServiceConfigsParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get log service configs params
func (o *GetLogServiceConfigsParams) SetRequestBody(requestBody *models.GetLogServiceConfigsRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetLogServiceConfigsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
