// Code generated by go-swagger; DO NOT EDIT.

package application

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

// NewGetApplicationsConnectionParams creates a new GetApplicationsConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetApplicationsConnectionParams() *GetApplicationsConnectionParams {
	return &GetApplicationsConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetApplicationsConnectionParamsWithTimeout creates a new GetApplicationsConnectionParams object
// with the ability to set a timeout on a request.
func NewGetApplicationsConnectionParamsWithTimeout(timeout time.Duration) *GetApplicationsConnectionParams {
	return &GetApplicationsConnectionParams{
		timeout: timeout,
	}
}

// NewGetApplicationsConnectionParamsWithContext creates a new GetApplicationsConnectionParams object
// with the ability to set a context for a request.
func NewGetApplicationsConnectionParamsWithContext(ctx context.Context) *GetApplicationsConnectionParams {
	return &GetApplicationsConnectionParams{
		Context: ctx,
	}
}

// NewGetApplicationsConnectionParamsWithHTTPClient creates a new GetApplicationsConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetApplicationsConnectionParamsWithHTTPClient(client *http.Client) *GetApplicationsConnectionParams {
	return &GetApplicationsConnectionParams{
		HTTPClient: client,
	}
}

/* GetApplicationsConnectionParams contains all the parameters to send to the API endpoint
   for the get applications connection operation.

   Typically these are written to a http.Request.
*/
type GetApplicationsConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetApplicationsConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get applications connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetApplicationsConnectionParams) WithDefaults() *GetApplicationsConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get applications connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetApplicationsConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetApplicationsConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get applications connection params
func (o *GetApplicationsConnectionParams) WithTimeout(timeout time.Duration) *GetApplicationsConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get applications connection params
func (o *GetApplicationsConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get applications connection params
func (o *GetApplicationsConnectionParams) WithContext(ctx context.Context) *GetApplicationsConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get applications connection params
func (o *GetApplicationsConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get applications connection params
func (o *GetApplicationsConnectionParams) WithHTTPClient(client *http.Client) *GetApplicationsConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get applications connection params
func (o *GetApplicationsConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get applications connection params
func (o *GetApplicationsConnectionParams) WithContentLanguage(contentLanguage *string) *GetApplicationsConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get applications connection params
func (o *GetApplicationsConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get applications connection params
func (o *GetApplicationsConnectionParams) WithRequestBody(requestBody *models.GetApplicationsConnectionRequestBody) *GetApplicationsConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get applications connection params
func (o *GetApplicationsConnectionParams) SetRequestBody(requestBody *models.GetApplicationsConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetApplicationsConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
