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

// NewEjectIsoFromVMCdRomParams creates a new EjectIsoFromVMCdRomParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEjectIsoFromVMCdRomParams() *EjectIsoFromVMCdRomParams {
	return &EjectIsoFromVMCdRomParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEjectIsoFromVMCdRomParamsWithTimeout creates a new EjectIsoFromVMCdRomParams object
// with the ability to set a timeout on a request.
func NewEjectIsoFromVMCdRomParamsWithTimeout(timeout time.Duration) *EjectIsoFromVMCdRomParams {
	return &EjectIsoFromVMCdRomParams{
		timeout: timeout,
	}
}

// NewEjectIsoFromVMCdRomParamsWithContext creates a new EjectIsoFromVMCdRomParams object
// with the ability to set a context for a request.
func NewEjectIsoFromVMCdRomParamsWithContext(ctx context.Context) *EjectIsoFromVMCdRomParams {
	return &EjectIsoFromVMCdRomParams{
		Context: ctx,
	}
}

// NewEjectIsoFromVMCdRomParamsWithHTTPClient creates a new EjectIsoFromVMCdRomParams object
// with the ability to set a custom HTTPClient for a request.
func NewEjectIsoFromVMCdRomParamsWithHTTPClient(client *http.Client) *EjectIsoFromVMCdRomParams {
	return &EjectIsoFromVMCdRomParams{
		HTTPClient: client,
	}
}

/* EjectIsoFromVMCdRomParams contains all the parameters to send to the API endpoint
   for the eject iso from Vm cd rom operation.

   Typically these are written to a http.Request.
*/
type EjectIsoFromVMCdRomParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.VMEjectCdRomParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the eject iso from Vm cd rom params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EjectIsoFromVMCdRomParams) WithDefaults() *EjectIsoFromVMCdRomParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the eject iso from Vm cd rom params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EjectIsoFromVMCdRomParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := EjectIsoFromVMCdRomParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the eject iso from Vm cd rom params
func (o *EjectIsoFromVMCdRomParams) WithTimeout(timeout time.Duration) *EjectIsoFromVMCdRomParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the eject iso from Vm cd rom params
func (o *EjectIsoFromVMCdRomParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the eject iso from Vm cd rom params
func (o *EjectIsoFromVMCdRomParams) WithContext(ctx context.Context) *EjectIsoFromVMCdRomParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the eject iso from Vm cd rom params
func (o *EjectIsoFromVMCdRomParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the eject iso from Vm cd rom params
func (o *EjectIsoFromVMCdRomParams) WithHTTPClient(client *http.Client) *EjectIsoFromVMCdRomParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the eject iso from Vm cd rom params
func (o *EjectIsoFromVMCdRomParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the eject iso from Vm cd rom params
func (o *EjectIsoFromVMCdRomParams) WithContentLanguage(contentLanguage *string) *EjectIsoFromVMCdRomParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the eject iso from Vm cd rom params
func (o *EjectIsoFromVMCdRomParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the eject iso from Vm cd rom params
func (o *EjectIsoFromVMCdRomParams) WithRequestBody(requestBody *models.VMEjectCdRomParams) *EjectIsoFromVMCdRomParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the eject iso from Vm cd rom params
func (o *EjectIsoFromVMCdRomParams) SetRequestBody(requestBody *models.VMEjectCdRomParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *EjectIsoFromVMCdRomParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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