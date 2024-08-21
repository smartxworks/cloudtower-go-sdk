// Code generated by go-swagger; DO NOT EDIT.

package virtual_private_cloud

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

// NewGetVirtualPrivateCloudsParams creates a new GetVirtualPrivateCloudsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVirtualPrivateCloudsParams() *GetVirtualPrivateCloudsParams {
	return &GetVirtualPrivateCloudsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVirtualPrivateCloudsParamsWithTimeout creates a new GetVirtualPrivateCloudsParams object
// with the ability to set a timeout on a request.
func NewGetVirtualPrivateCloudsParamsWithTimeout(timeout time.Duration) *GetVirtualPrivateCloudsParams {
	return &GetVirtualPrivateCloudsParams{
		timeout: timeout,
	}
}

// NewGetVirtualPrivateCloudsParamsWithContext creates a new GetVirtualPrivateCloudsParams object
// with the ability to set a context for a request.
func NewGetVirtualPrivateCloudsParamsWithContext(ctx context.Context) *GetVirtualPrivateCloudsParams {
	return &GetVirtualPrivateCloudsParams{
		Context: ctx,
	}
}

// NewGetVirtualPrivateCloudsParamsWithHTTPClient creates a new GetVirtualPrivateCloudsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVirtualPrivateCloudsParamsWithHTTPClient(client *http.Client) *GetVirtualPrivateCloudsParams {
	return &GetVirtualPrivateCloudsParams{
		HTTPClient: client,
	}
}

/* GetVirtualPrivateCloudsParams contains all the parameters to send to the API endpoint
   for the get virtual private clouds operation.

   Typically these are written to a http.Request.
*/
type GetVirtualPrivateCloudsParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetVirtualPrivateCloudsRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get virtual private clouds params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVirtualPrivateCloudsParams) WithDefaults() *GetVirtualPrivateCloudsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get virtual private clouds params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVirtualPrivateCloudsParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetVirtualPrivateCloudsParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get virtual private clouds params
func (o *GetVirtualPrivateCloudsParams) WithTimeout(timeout time.Duration) *GetVirtualPrivateCloudsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get virtual private clouds params
func (o *GetVirtualPrivateCloudsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get virtual private clouds params
func (o *GetVirtualPrivateCloudsParams) WithContext(ctx context.Context) *GetVirtualPrivateCloudsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get virtual private clouds params
func (o *GetVirtualPrivateCloudsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get virtual private clouds params
func (o *GetVirtualPrivateCloudsParams) WithHTTPClient(client *http.Client) *GetVirtualPrivateCloudsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get virtual private clouds params
func (o *GetVirtualPrivateCloudsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get virtual private clouds params
func (o *GetVirtualPrivateCloudsParams) WithContentLanguage(contentLanguage *string) *GetVirtualPrivateCloudsParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get virtual private clouds params
func (o *GetVirtualPrivateCloudsParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get virtual private clouds params
func (o *GetVirtualPrivateCloudsParams) WithRequestBody(requestBody *models.GetVirtualPrivateCloudsRequestBody) *GetVirtualPrivateCloudsParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get virtual private clouds params
func (o *GetVirtualPrivateCloudsParams) SetRequestBody(requestBody *models.GetVirtualPrivateCloudsRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetVirtualPrivateCloudsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
