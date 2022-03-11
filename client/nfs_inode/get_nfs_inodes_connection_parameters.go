// Code generated by go-swagger; DO NOT EDIT.

package nfs_inode

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

// NewGetNfsInodesConnectionParams creates a new GetNfsInodesConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNfsInodesConnectionParams() *GetNfsInodesConnectionParams {
	return &GetNfsInodesConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNfsInodesConnectionParamsWithTimeout creates a new GetNfsInodesConnectionParams object
// with the ability to set a timeout on a request.
func NewGetNfsInodesConnectionParamsWithTimeout(timeout time.Duration) *GetNfsInodesConnectionParams {
	return &GetNfsInodesConnectionParams{
		timeout: timeout,
	}
}

// NewGetNfsInodesConnectionParamsWithContext creates a new GetNfsInodesConnectionParams object
// with the ability to set a context for a request.
func NewGetNfsInodesConnectionParamsWithContext(ctx context.Context) *GetNfsInodesConnectionParams {
	return &GetNfsInodesConnectionParams{
		Context: ctx,
	}
}

// NewGetNfsInodesConnectionParamsWithHTTPClient creates a new GetNfsInodesConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNfsInodesConnectionParamsWithHTTPClient(client *http.Client) *GetNfsInodesConnectionParams {
	return &GetNfsInodesConnectionParams{
		HTTPClient: client,
	}
}

/* GetNfsInodesConnectionParams contains all the parameters to send to the API endpoint
   for the get nfs inodes connection operation.

   Typically these are written to a http.Request.
*/
type GetNfsInodesConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetNfsInodesConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get nfs inodes connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNfsInodesConnectionParams) WithDefaults() *GetNfsInodesConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get nfs inodes connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNfsInodesConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetNfsInodesConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get nfs inodes connection params
func (o *GetNfsInodesConnectionParams) WithTimeout(timeout time.Duration) *GetNfsInodesConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nfs inodes connection params
func (o *GetNfsInodesConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nfs inodes connection params
func (o *GetNfsInodesConnectionParams) WithContext(ctx context.Context) *GetNfsInodesConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nfs inodes connection params
func (o *GetNfsInodesConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nfs inodes connection params
func (o *GetNfsInodesConnectionParams) WithHTTPClient(client *http.Client) *GetNfsInodesConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nfs inodes connection params
func (o *GetNfsInodesConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get nfs inodes connection params
func (o *GetNfsInodesConnectionParams) WithContentLanguage(contentLanguage *string) *GetNfsInodesConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get nfs inodes connection params
func (o *GetNfsInodesConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get nfs inodes connection params
func (o *GetNfsInodesConnectionParams) WithRequestBody(requestBody *models.GetNfsInodesConnectionRequestBody) *GetNfsInodesConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get nfs inodes connection params
func (o *GetNfsInodesConnectionParams) SetRequestBody(requestBody *models.GetNfsInodesConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetNfsInodesConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
