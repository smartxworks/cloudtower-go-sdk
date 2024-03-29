// Code generated by go-swagger; DO NOT EDIT.

package host

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

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// NewEnterMaintenanceModePreCheckParams creates a new EnterMaintenanceModePreCheckParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEnterMaintenanceModePreCheckParams() *EnterMaintenanceModePreCheckParams {
	return &EnterMaintenanceModePreCheckParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEnterMaintenanceModePreCheckParamsWithTimeout creates a new EnterMaintenanceModePreCheckParams object
// with the ability to set a timeout on a request.
func NewEnterMaintenanceModePreCheckParamsWithTimeout(timeout time.Duration) *EnterMaintenanceModePreCheckParams {
	return &EnterMaintenanceModePreCheckParams{
		timeout: timeout,
	}
}

// NewEnterMaintenanceModePreCheckParamsWithContext creates a new EnterMaintenanceModePreCheckParams object
// with the ability to set a context for a request.
func NewEnterMaintenanceModePreCheckParamsWithContext(ctx context.Context) *EnterMaintenanceModePreCheckParams {
	return &EnterMaintenanceModePreCheckParams{
		Context: ctx,
	}
}

// NewEnterMaintenanceModePreCheckParamsWithHTTPClient creates a new EnterMaintenanceModePreCheckParams object
// with the ability to set a custom HTTPClient for a request.
func NewEnterMaintenanceModePreCheckParamsWithHTTPClient(client *http.Client) *EnterMaintenanceModePreCheckParams {
	return &EnterMaintenanceModePreCheckParams{
		HTTPClient: client,
	}
}

/* EnterMaintenanceModePreCheckParams contains all the parameters to send to the API endpoint
   for the enter maintenance mode pre check operation.

   Typically these are written to a http.Request.
*/
type EnterMaintenanceModePreCheckParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.EnterMaintenanceModeCheckParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the enter maintenance mode pre check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnterMaintenanceModePreCheckParams) WithDefaults() *EnterMaintenanceModePreCheckParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the enter maintenance mode pre check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnterMaintenanceModePreCheckParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := EnterMaintenanceModePreCheckParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the enter maintenance mode pre check params
func (o *EnterMaintenanceModePreCheckParams) WithTimeout(timeout time.Duration) *EnterMaintenanceModePreCheckParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enter maintenance mode pre check params
func (o *EnterMaintenanceModePreCheckParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enter maintenance mode pre check params
func (o *EnterMaintenanceModePreCheckParams) WithContext(ctx context.Context) *EnterMaintenanceModePreCheckParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enter maintenance mode pre check params
func (o *EnterMaintenanceModePreCheckParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enter maintenance mode pre check params
func (o *EnterMaintenanceModePreCheckParams) WithHTTPClient(client *http.Client) *EnterMaintenanceModePreCheckParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enter maintenance mode pre check params
func (o *EnterMaintenanceModePreCheckParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the enter maintenance mode pre check params
func (o *EnterMaintenanceModePreCheckParams) WithContentLanguage(contentLanguage *string) *EnterMaintenanceModePreCheckParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the enter maintenance mode pre check params
func (o *EnterMaintenanceModePreCheckParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the enter maintenance mode pre check params
func (o *EnterMaintenanceModePreCheckParams) WithRequestBody(requestBody *models.EnterMaintenanceModeCheckParams) *EnterMaintenanceModePreCheckParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the enter maintenance mode pre check params
func (o *EnterMaintenanceModePreCheckParams) SetRequestBody(requestBody *models.EnterMaintenanceModeCheckParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *EnterMaintenanceModePreCheckParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
