// Code generated by go-swagger; DO NOT EDIT.

package vm_snapshot

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

// NewGetVMSnapshotsParams creates a new GetVMSnapshotsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVMSnapshotsParams() *GetVMSnapshotsParams {
	return &GetVMSnapshotsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVMSnapshotsParamsWithTimeout creates a new GetVMSnapshotsParams object
// with the ability to set a timeout on a request.
func NewGetVMSnapshotsParamsWithTimeout(timeout time.Duration) *GetVMSnapshotsParams {
	return &GetVMSnapshotsParams{
		timeout: timeout,
	}
}

// NewGetVMSnapshotsParamsWithContext creates a new GetVMSnapshotsParams object
// with the ability to set a context for a request.
func NewGetVMSnapshotsParamsWithContext(ctx context.Context) *GetVMSnapshotsParams {
	return &GetVMSnapshotsParams{
		Context: ctx,
	}
}

// NewGetVMSnapshotsParamsWithHTTPClient creates a new GetVMSnapshotsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVMSnapshotsParamsWithHTTPClient(client *http.Client) *GetVMSnapshotsParams {
	return &GetVMSnapshotsParams{
		HTTPClient: client,
	}
}

/* GetVMSnapshotsParams contains all the parameters to send to the API endpoint
   for the get Vm snapshots operation.

   Typically these are written to a http.Request.
*/
type GetVMSnapshotsParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetVMSnapshotsRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Vm snapshots params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVMSnapshotsParams) WithDefaults() *GetVMSnapshotsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Vm snapshots params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVMSnapshotsParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetVMSnapshotsParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get Vm snapshots params
func (o *GetVMSnapshotsParams) WithTimeout(timeout time.Duration) *GetVMSnapshotsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Vm snapshots params
func (o *GetVMSnapshotsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Vm snapshots params
func (o *GetVMSnapshotsParams) WithContext(ctx context.Context) *GetVMSnapshotsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Vm snapshots params
func (o *GetVMSnapshotsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Vm snapshots params
func (o *GetVMSnapshotsParams) WithHTTPClient(client *http.Client) *GetVMSnapshotsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Vm snapshots params
func (o *GetVMSnapshotsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get Vm snapshots params
func (o *GetVMSnapshotsParams) WithContentLanguage(contentLanguage *string) *GetVMSnapshotsParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get Vm snapshots params
func (o *GetVMSnapshotsParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get Vm snapshots params
func (o *GetVMSnapshotsParams) WithRequestBody(requestBody *models.GetVMSnapshotsRequestBody) *GetVMSnapshotsParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get Vm snapshots params
func (o *GetVMSnapshotsParams) SetRequestBody(requestBody *models.GetVMSnapshotsRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetVMSnapshotsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
