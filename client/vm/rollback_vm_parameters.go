// Code generated by go-swagger; DO NOT EDIT.

package vm

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

// NewRollbackVMParams creates a new RollbackVMParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRollbackVMParams() *RollbackVMParams {
	return &RollbackVMParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRollbackVMParamsWithTimeout creates a new RollbackVMParams object
// with the ability to set a timeout on a request.
func NewRollbackVMParamsWithTimeout(timeout time.Duration) *RollbackVMParams {
	return &RollbackVMParams{
		timeout: timeout,
	}
}

// NewRollbackVMParamsWithContext creates a new RollbackVMParams object
// with the ability to set a context for a request.
func NewRollbackVMParamsWithContext(ctx context.Context) *RollbackVMParams {
	return &RollbackVMParams{
		Context: ctx,
	}
}

// NewRollbackVMParamsWithHTTPClient creates a new RollbackVMParams object
// with the ability to set a custom HTTPClient for a request.
func NewRollbackVMParamsWithHTTPClient(client *http.Client) *RollbackVMParams {
	return &RollbackVMParams{
		HTTPClient: client,
	}
}

/* RollbackVMParams contains all the parameters to send to the API endpoint
   for the rollback Vm operation.

   Typically these are written to a http.Request.
*/
type RollbackVMParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.VMRollbackParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the rollback Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RollbackVMParams) WithDefaults() *RollbackVMParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the rollback Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RollbackVMParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := RollbackVMParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the rollback Vm params
func (o *RollbackVMParams) WithTimeout(timeout time.Duration) *RollbackVMParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rollback Vm params
func (o *RollbackVMParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rollback Vm params
func (o *RollbackVMParams) WithContext(ctx context.Context) *RollbackVMParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rollback Vm params
func (o *RollbackVMParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rollback Vm params
func (o *RollbackVMParams) WithHTTPClient(client *http.Client) *RollbackVMParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rollback Vm params
func (o *RollbackVMParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the rollback Vm params
func (o *RollbackVMParams) WithContentLanguage(contentLanguage *string) *RollbackVMParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the rollback Vm params
func (o *RollbackVMParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the rollback Vm params
func (o *RollbackVMParams) WithRequestBody(requestBody *models.VMRollbackParams) *RollbackVMParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the rollback Vm params
func (o *RollbackVMParams) SetRequestBody(requestBody *models.VMRollbackParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *RollbackVMParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
