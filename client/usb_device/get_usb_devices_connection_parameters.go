// Code generated by go-swagger; DO NOT EDIT.

package usb_device

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

// NewGetUsbDevicesConnectionParams creates a new GetUsbDevicesConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUsbDevicesConnectionParams() *GetUsbDevicesConnectionParams {
	return &GetUsbDevicesConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsbDevicesConnectionParamsWithTimeout creates a new GetUsbDevicesConnectionParams object
// with the ability to set a timeout on a request.
func NewGetUsbDevicesConnectionParamsWithTimeout(timeout time.Duration) *GetUsbDevicesConnectionParams {
	return &GetUsbDevicesConnectionParams{
		timeout: timeout,
	}
}

// NewGetUsbDevicesConnectionParamsWithContext creates a new GetUsbDevicesConnectionParams object
// with the ability to set a context for a request.
func NewGetUsbDevicesConnectionParamsWithContext(ctx context.Context) *GetUsbDevicesConnectionParams {
	return &GetUsbDevicesConnectionParams{
		Context: ctx,
	}
}

// NewGetUsbDevicesConnectionParamsWithHTTPClient creates a new GetUsbDevicesConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUsbDevicesConnectionParamsWithHTTPClient(client *http.Client) *GetUsbDevicesConnectionParams {
	return &GetUsbDevicesConnectionParams{
		HTTPClient: client,
	}
}

/* GetUsbDevicesConnectionParams contains all the parameters to send to the API endpoint
   for the get usb devices connection operation.

   Typically these are written to a http.Request.
*/
type GetUsbDevicesConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetUsbDevicesConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get usb devices connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsbDevicesConnectionParams) WithDefaults() *GetUsbDevicesConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get usb devices connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsbDevicesConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetUsbDevicesConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get usb devices connection params
func (o *GetUsbDevicesConnectionParams) WithTimeout(timeout time.Duration) *GetUsbDevicesConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get usb devices connection params
func (o *GetUsbDevicesConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get usb devices connection params
func (o *GetUsbDevicesConnectionParams) WithContext(ctx context.Context) *GetUsbDevicesConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get usb devices connection params
func (o *GetUsbDevicesConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get usb devices connection params
func (o *GetUsbDevicesConnectionParams) WithHTTPClient(client *http.Client) *GetUsbDevicesConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get usb devices connection params
func (o *GetUsbDevicesConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get usb devices connection params
func (o *GetUsbDevicesConnectionParams) WithContentLanguage(contentLanguage *string) *GetUsbDevicesConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get usb devices connection params
func (o *GetUsbDevicesConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get usb devices connection params
func (o *GetUsbDevicesConnectionParams) WithRequestBody(requestBody *models.GetUsbDevicesConnectionRequestBody) *GetUsbDevicesConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get usb devices connection params
func (o *GetUsbDevicesConnectionParams) SetRequestBody(requestBody *models.GetUsbDevicesConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsbDevicesConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
