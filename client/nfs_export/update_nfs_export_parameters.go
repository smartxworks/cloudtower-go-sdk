// Code generated by go-swagger; DO NOT EDIT.

package nfs_export

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

// NewUpdateNfsExportParams creates a new UpdateNfsExportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNfsExportParams() *UpdateNfsExportParams {
	return &UpdateNfsExportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNfsExportParamsWithTimeout creates a new UpdateNfsExportParams object
// with the ability to set a timeout on a request.
func NewUpdateNfsExportParamsWithTimeout(timeout time.Duration) *UpdateNfsExportParams {
	return &UpdateNfsExportParams{
		timeout: timeout,
	}
}

// NewUpdateNfsExportParamsWithContext creates a new UpdateNfsExportParams object
// with the ability to set a context for a request.
func NewUpdateNfsExportParamsWithContext(ctx context.Context) *UpdateNfsExportParams {
	return &UpdateNfsExportParams{
		Context: ctx,
	}
}

// NewUpdateNfsExportParamsWithHTTPClient creates a new UpdateNfsExportParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNfsExportParamsWithHTTPClient(client *http.Client) *UpdateNfsExportParams {
	return &UpdateNfsExportParams{
		HTTPClient: client,
	}
}

/* UpdateNfsExportParams contains all the parameters to send to the API endpoint
   for the update nfs export operation.

   Typically these are written to a http.Request.
*/
type UpdateNfsExportParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.NfsExportUpdationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update nfs export params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNfsExportParams) WithDefaults() *UpdateNfsExportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update nfs export params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNfsExportParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := UpdateNfsExportParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update nfs export params
func (o *UpdateNfsExportParams) WithTimeout(timeout time.Duration) *UpdateNfsExportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update nfs export params
func (o *UpdateNfsExportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update nfs export params
func (o *UpdateNfsExportParams) WithContext(ctx context.Context) *UpdateNfsExportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update nfs export params
func (o *UpdateNfsExportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update nfs export params
func (o *UpdateNfsExportParams) WithHTTPClient(client *http.Client) *UpdateNfsExportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update nfs export params
func (o *UpdateNfsExportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the update nfs export params
func (o *UpdateNfsExportParams) WithContentLanguage(contentLanguage *string) *UpdateNfsExportParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the update nfs export params
func (o *UpdateNfsExportParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the update nfs export params
func (o *UpdateNfsExportParams) WithRequestBody(requestBody *models.NfsExportUpdationParams) *UpdateNfsExportParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update nfs export params
func (o *UpdateNfsExportParams) SetRequestBody(requestBody *models.NfsExportUpdationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNfsExportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
