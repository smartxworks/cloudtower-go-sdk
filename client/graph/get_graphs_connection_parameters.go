// Code generated by go-swagger; DO NOT EDIT.

package graph

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

// NewGetGraphsConnectionParams creates a new GetGraphsConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGraphsConnectionParams() *GetGraphsConnectionParams {
	return &GetGraphsConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGraphsConnectionParamsWithTimeout creates a new GetGraphsConnectionParams object
// with the ability to set a timeout on a request.
func NewGetGraphsConnectionParamsWithTimeout(timeout time.Duration) *GetGraphsConnectionParams {
	return &GetGraphsConnectionParams{
		timeout: timeout,
	}
}

// NewGetGraphsConnectionParamsWithContext creates a new GetGraphsConnectionParams object
// with the ability to set a context for a request.
func NewGetGraphsConnectionParamsWithContext(ctx context.Context) *GetGraphsConnectionParams {
	return &GetGraphsConnectionParams{
		Context: ctx,
	}
}

// NewGetGraphsConnectionParamsWithHTTPClient creates a new GetGraphsConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGraphsConnectionParamsWithHTTPClient(client *http.Client) *GetGraphsConnectionParams {
	return &GetGraphsConnectionParams{
		HTTPClient: client,
	}
}

/* GetGraphsConnectionParams contains all the parameters to send to the API endpoint
   for the get graphs connection operation.

   Typically these are written to a http.Request.
*/
type GetGraphsConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetGraphsConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get graphs connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGraphsConnectionParams) WithDefaults() *GetGraphsConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get graphs connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGraphsConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetGraphsConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get graphs connection params
func (o *GetGraphsConnectionParams) WithTimeout(timeout time.Duration) *GetGraphsConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get graphs connection params
func (o *GetGraphsConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get graphs connection params
func (o *GetGraphsConnectionParams) WithContext(ctx context.Context) *GetGraphsConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get graphs connection params
func (o *GetGraphsConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get graphs connection params
func (o *GetGraphsConnectionParams) WithHTTPClient(client *http.Client) *GetGraphsConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get graphs connection params
func (o *GetGraphsConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get graphs connection params
func (o *GetGraphsConnectionParams) WithContentLanguage(contentLanguage *string) *GetGraphsConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get graphs connection params
func (o *GetGraphsConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get graphs connection params
func (o *GetGraphsConnectionParams) WithRequestBody(requestBody *models.GetGraphsConnectionRequestBody) *GetGraphsConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get graphs connection params
func (o *GetGraphsConnectionParams) SetRequestBody(requestBody *models.GetGraphsConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetGraphsConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
