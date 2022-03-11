// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewGetUsersConnectionParams creates a new GetUsersConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUsersConnectionParams() *GetUsersConnectionParams {
	return &GetUsersConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsersConnectionParamsWithTimeout creates a new GetUsersConnectionParams object
// with the ability to set a timeout on a request.
func NewGetUsersConnectionParamsWithTimeout(timeout time.Duration) *GetUsersConnectionParams {
	return &GetUsersConnectionParams{
		timeout: timeout,
	}
}

// NewGetUsersConnectionParamsWithContext creates a new GetUsersConnectionParams object
// with the ability to set a context for a request.
func NewGetUsersConnectionParamsWithContext(ctx context.Context) *GetUsersConnectionParams {
	return &GetUsersConnectionParams{
		Context: ctx,
	}
}

// NewGetUsersConnectionParamsWithHTTPClient creates a new GetUsersConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUsersConnectionParamsWithHTTPClient(client *http.Client) *GetUsersConnectionParams {
	return &GetUsersConnectionParams{
		HTTPClient: client,
	}
}

/* GetUsersConnectionParams contains all the parameters to send to the API endpoint
   for the get users connection operation.

   Typically these are written to a http.Request.
*/
type GetUsersConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetUsersConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get users connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsersConnectionParams) WithDefaults() *GetUsersConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get users connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsersConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetUsersConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get users connection params
func (o *GetUsersConnectionParams) WithTimeout(timeout time.Duration) *GetUsersConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get users connection params
func (o *GetUsersConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get users connection params
func (o *GetUsersConnectionParams) WithContext(ctx context.Context) *GetUsersConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get users connection params
func (o *GetUsersConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get users connection params
func (o *GetUsersConnectionParams) WithHTTPClient(client *http.Client) *GetUsersConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get users connection params
func (o *GetUsersConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get users connection params
func (o *GetUsersConnectionParams) WithContentLanguage(contentLanguage *string) *GetUsersConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get users connection params
func (o *GetUsersConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get users connection params
func (o *GetUsersConnectionParams) WithRequestBody(requestBody *models.GetUsersConnectionRequestBody) *GetUsersConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get users connection params
func (o *GetUsersConnectionParams) SetRequestBody(requestBody *models.GetUsersConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsersConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
