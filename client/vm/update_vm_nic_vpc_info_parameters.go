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

// NewUpdateVMNicVpcInfoParams creates a new UpdateVMNicVpcInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateVMNicVpcInfoParams() *UpdateVMNicVpcInfoParams {
	return &UpdateVMNicVpcInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVMNicVpcInfoParamsWithTimeout creates a new UpdateVMNicVpcInfoParams object
// with the ability to set a timeout on a request.
func NewUpdateVMNicVpcInfoParamsWithTimeout(timeout time.Duration) *UpdateVMNicVpcInfoParams {
	return &UpdateVMNicVpcInfoParams{
		timeout: timeout,
	}
}

// NewUpdateVMNicVpcInfoParamsWithContext creates a new UpdateVMNicVpcInfoParams object
// with the ability to set a context for a request.
func NewUpdateVMNicVpcInfoParamsWithContext(ctx context.Context) *UpdateVMNicVpcInfoParams {
	return &UpdateVMNicVpcInfoParams{
		Context: ctx,
	}
}

// NewUpdateVMNicVpcInfoParamsWithHTTPClient creates a new UpdateVMNicVpcInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateVMNicVpcInfoParamsWithHTTPClient(client *http.Client) *UpdateVMNicVpcInfoParams {
	return &UpdateVMNicVpcInfoParams{
		HTTPClient: client,
	}
}

/* UpdateVMNicVpcInfoParams contains all the parameters to send to the API endpoint
   for the update Vm nic vpc info operation.

   Typically these are written to a http.Request.
*/
type UpdateVMNicVpcInfoParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.VMUpdateVpcNicParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update Vm nic vpc info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVMNicVpcInfoParams) WithDefaults() *UpdateVMNicVpcInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update Vm nic vpc info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVMNicVpcInfoParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := UpdateVMNicVpcInfoParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update Vm nic vpc info params
func (o *UpdateVMNicVpcInfoParams) WithTimeout(timeout time.Duration) *UpdateVMNicVpcInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update Vm nic vpc info params
func (o *UpdateVMNicVpcInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update Vm nic vpc info params
func (o *UpdateVMNicVpcInfoParams) WithContext(ctx context.Context) *UpdateVMNicVpcInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update Vm nic vpc info params
func (o *UpdateVMNicVpcInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update Vm nic vpc info params
func (o *UpdateVMNicVpcInfoParams) WithHTTPClient(client *http.Client) *UpdateVMNicVpcInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update Vm nic vpc info params
func (o *UpdateVMNicVpcInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the update Vm nic vpc info params
func (o *UpdateVMNicVpcInfoParams) WithContentLanguage(contentLanguage *string) *UpdateVMNicVpcInfoParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the update Vm nic vpc info params
func (o *UpdateVMNicVpcInfoParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the update Vm nic vpc info params
func (o *UpdateVMNicVpcInfoParams) WithRequestBody(requestBody *models.VMUpdateVpcNicParams) *UpdateVMNicVpcInfoParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update Vm nic vpc info params
func (o *UpdateVMNicVpcInfoParams) SetRequestBody(requestBody *models.VMUpdateVpcNicParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVMNicVpcInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
