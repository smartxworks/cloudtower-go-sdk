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

// NewGetVMVncInfoParams creates a new GetVMVncInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVMVncInfoParams() *GetVMVncInfoParams {
	return &GetVMVncInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVMVncInfoParamsWithTimeout creates a new GetVMVncInfoParams object
// with the ability to set a timeout on a request.
func NewGetVMVncInfoParamsWithTimeout(timeout time.Duration) *GetVMVncInfoParams {
	return &GetVMVncInfoParams{
		timeout: timeout,
	}
}

// NewGetVMVncInfoParamsWithContext creates a new GetVMVncInfoParams object
// with the ability to set a context for a request.
func NewGetVMVncInfoParamsWithContext(ctx context.Context) *GetVMVncInfoParams {
	return &GetVMVncInfoParams{
		Context: ctx,
	}
}

// NewGetVMVncInfoParamsWithHTTPClient creates a new GetVMVncInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVMVncInfoParamsWithHTTPClient(client *http.Client) *GetVMVncInfoParams {
	return &GetVMVncInfoParams{
		HTTPClient: client,
	}
}

/* GetVMVncInfoParams contains all the parameters to send to the API endpoint
   for the get Vm vnc info operation.

   Typically these are written to a http.Request.
*/
type GetVMVncInfoParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetVMVncInfoParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Vm vnc info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVMVncInfoParams) WithDefaults() *GetVMVncInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Vm vnc info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVMVncInfoParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetVMVncInfoParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get Vm vnc info params
func (o *GetVMVncInfoParams) WithTimeout(timeout time.Duration) *GetVMVncInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Vm vnc info params
func (o *GetVMVncInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Vm vnc info params
func (o *GetVMVncInfoParams) WithContext(ctx context.Context) *GetVMVncInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Vm vnc info params
func (o *GetVMVncInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Vm vnc info params
func (o *GetVMVncInfoParams) WithHTTPClient(client *http.Client) *GetVMVncInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Vm vnc info params
func (o *GetVMVncInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get Vm vnc info params
func (o *GetVMVncInfoParams) WithContentLanguage(contentLanguage *string) *GetVMVncInfoParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get Vm vnc info params
func (o *GetVMVncInfoParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get Vm vnc info params
func (o *GetVMVncInfoParams) WithRequestBody(requestBody *models.GetVMVncInfoParams) *GetVMVncInfoParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get Vm vnc info params
func (o *GetVMVncInfoParams) SetRequestBody(requestBody *models.GetVMVncInfoParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetVMVncInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
