// Code generated by go-swagger; DO NOT EDIT.

package replication_plan

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

// NewGetReplicationPlansConnectionParams creates a new GetReplicationPlansConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetReplicationPlansConnectionParams() *GetReplicationPlansConnectionParams {
	return &GetReplicationPlansConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetReplicationPlansConnectionParamsWithTimeout creates a new GetReplicationPlansConnectionParams object
// with the ability to set a timeout on a request.
func NewGetReplicationPlansConnectionParamsWithTimeout(timeout time.Duration) *GetReplicationPlansConnectionParams {
	return &GetReplicationPlansConnectionParams{
		timeout: timeout,
	}
}

// NewGetReplicationPlansConnectionParamsWithContext creates a new GetReplicationPlansConnectionParams object
// with the ability to set a context for a request.
func NewGetReplicationPlansConnectionParamsWithContext(ctx context.Context) *GetReplicationPlansConnectionParams {
	return &GetReplicationPlansConnectionParams{
		Context: ctx,
	}
}

// NewGetReplicationPlansConnectionParamsWithHTTPClient creates a new GetReplicationPlansConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetReplicationPlansConnectionParamsWithHTTPClient(client *http.Client) *GetReplicationPlansConnectionParams {
	return &GetReplicationPlansConnectionParams{
		HTTPClient: client,
	}
}

/* GetReplicationPlansConnectionParams contains all the parameters to send to the API endpoint
   for the get replication plans connection operation.

   Typically these are written to a http.Request.
*/
type GetReplicationPlansConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetReplicationPlansConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get replication plans connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetReplicationPlansConnectionParams) WithDefaults() *GetReplicationPlansConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get replication plans connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetReplicationPlansConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetReplicationPlansConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get replication plans connection params
func (o *GetReplicationPlansConnectionParams) WithTimeout(timeout time.Duration) *GetReplicationPlansConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get replication plans connection params
func (o *GetReplicationPlansConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get replication plans connection params
func (o *GetReplicationPlansConnectionParams) WithContext(ctx context.Context) *GetReplicationPlansConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get replication plans connection params
func (o *GetReplicationPlansConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get replication plans connection params
func (o *GetReplicationPlansConnectionParams) WithHTTPClient(client *http.Client) *GetReplicationPlansConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get replication plans connection params
func (o *GetReplicationPlansConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get replication plans connection params
func (o *GetReplicationPlansConnectionParams) WithContentLanguage(contentLanguage *string) *GetReplicationPlansConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get replication plans connection params
func (o *GetReplicationPlansConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get replication plans connection params
func (o *GetReplicationPlansConnectionParams) WithRequestBody(requestBody *models.GetReplicationPlansConnectionRequestBody) *GetReplicationPlansConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get replication plans connection params
func (o *GetReplicationPlansConnectionParams) SetRequestBody(requestBody *models.GetReplicationPlansConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetReplicationPlansConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
