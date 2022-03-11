// Code generated by go-swagger; DO NOT EDIT.

package disk

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

// NewUnmountDiskParams creates a new UnmountDiskParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUnmountDiskParams() *UnmountDiskParams {
	return &UnmountDiskParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUnmountDiskParamsWithTimeout creates a new UnmountDiskParams object
// with the ability to set a timeout on a request.
func NewUnmountDiskParamsWithTimeout(timeout time.Duration) *UnmountDiskParams {
	return &UnmountDiskParams{
		timeout: timeout,
	}
}

// NewUnmountDiskParamsWithContext creates a new UnmountDiskParams object
// with the ability to set a context for a request.
func NewUnmountDiskParamsWithContext(ctx context.Context) *UnmountDiskParams {
	return &UnmountDiskParams{
		Context: ctx,
	}
}

// NewUnmountDiskParamsWithHTTPClient creates a new UnmountDiskParams object
// with the ability to set a custom HTTPClient for a request.
func NewUnmountDiskParamsWithHTTPClient(client *http.Client) *UnmountDiskParams {
	return &UnmountDiskParams{
		HTTPClient: client,
	}
}

/* UnmountDiskParams contains all the parameters to send to the API endpoint
   for the unmount disk operation.

   Typically these are written to a http.Request.
*/
type UnmountDiskParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.DiskUnmountParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the unmount disk params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnmountDiskParams) WithDefaults() *UnmountDiskParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the unmount disk params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnmountDiskParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := UnmountDiskParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the unmount disk params
func (o *UnmountDiskParams) WithTimeout(timeout time.Duration) *UnmountDiskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unmount disk params
func (o *UnmountDiskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unmount disk params
func (o *UnmountDiskParams) WithContext(ctx context.Context) *UnmountDiskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unmount disk params
func (o *UnmountDiskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unmount disk params
func (o *UnmountDiskParams) WithHTTPClient(client *http.Client) *UnmountDiskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unmount disk params
func (o *UnmountDiskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the unmount disk params
func (o *UnmountDiskParams) WithContentLanguage(contentLanguage *string) *UnmountDiskParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the unmount disk params
func (o *UnmountDiskParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the unmount disk params
func (o *UnmountDiskParams) WithRequestBody(requestBody *models.DiskUnmountParams) *UnmountDiskParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the unmount disk params
func (o *UnmountDiskParams) SetRequestBody(requestBody *models.DiskUnmountParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UnmountDiskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
