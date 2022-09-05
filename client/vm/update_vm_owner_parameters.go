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

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// NewUpdateVMOwnerParams creates a new UpdateVMOwnerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateVMOwnerParams() *UpdateVMOwnerParams {
	return &UpdateVMOwnerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVMOwnerParamsWithTimeout creates a new UpdateVMOwnerParams object
// with the ability to set a timeout on a request.
func NewUpdateVMOwnerParamsWithTimeout(timeout time.Duration) *UpdateVMOwnerParams {
	return &UpdateVMOwnerParams{
		timeout: timeout,
	}
}

// NewUpdateVMOwnerParamsWithContext creates a new UpdateVMOwnerParams object
// with the ability to set a context for a request.
func NewUpdateVMOwnerParamsWithContext(ctx context.Context) *UpdateVMOwnerParams {
	return &UpdateVMOwnerParams{
		Context: ctx,
	}
}

// NewUpdateVMOwnerParamsWithHTTPClient creates a new UpdateVMOwnerParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateVMOwnerParamsWithHTTPClient(client *http.Client) *UpdateVMOwnerParams {
	return &UpdateVMOwnerParams{
		HTTPClient: client,
	}
}

/* UpdateVMOwnerParams contains all the parameters to send to the API endpoint
   for the update Vm owner operation.

   Typically these are written to a http.Request.
*/
type UpdateVMOwnerParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.VMUpdateOwnerParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update Vm owner params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVMOwnerParams) WithDefaults() *UpdateVMOwnerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update Vm owner params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVMOwnerParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := UpdateVMOwnerParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update Vm owner params
func (o *UpdateVMOwnerParams) WithTimeout(timeout time.Duration) *UpdateVMOwnerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update Vm owner params
func (o *UpdateVMOwnerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update Vm owner params
func (o *UpdateVMOwnerParams) WithContext(ctx context.Context) *UpdateVMOwnerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update Vm owner params
func (o *UpdateVMOwnerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update Vm owner params
func (o *UpdateVMOwnerParams) WithHTTPClient(client *http.Client) *UpdateVMOwnerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update Vm owner params
func (o *UpdateVMOwnerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the update Vm owner params
func (o *UpdateVMOwnerParams) WithContentLanguage(contentLanguage *string) *UpdateVMOwnerParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the update Vm owner params
func (o *UpdateVMOwnerParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the update Vm owner params
func (o *UpdateVMOwnerParams) WithRequestBody(requestBody *models.VMUpdateOwnerParams) *UpdateVMOwnerParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update Vm owner params
func (o *UpdateVMOwnerParams) SetRequestBody(requestBody *models.VMUpdateOwnerParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVMOwnerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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