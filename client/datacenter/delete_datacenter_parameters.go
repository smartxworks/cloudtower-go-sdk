// Code generated by go-swagger; DO NOT EDIT.

package datacenter

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

// NewDeleteDatacenterParams creates a new DeleteDatacenterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDatacenterParams() *DeleteDatacenterParams {
	return &DeleteDatacenterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDatacenterParamsWithTimeout creates a new DeleteDatacenterParams object
// with the ability to set a timeout on a request.
func NewDeleteDatacenterParamsWithTimeout(timeout time.Duration) *DeleteDatacenterParams {
	return &DeleteDatacenterParams{
		timeout: timeout,
	}
}

// NewDeleteDatacenterParamsWithContext creates a new DeleteDatacenterParams object
// with the ability to set a context for a request.
func NewDeleteDatacenterParamsWithContext(ctx context.Context) *DeleteDatacenterParams {
	return &DeleteDatacenterParams{
		Context: ctx,
	}
}

// NewDeleteDatacenterParamsWithHTTPClient creates a new DeleteDatacenterParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDatacenterParamsWithHTTPClient(client *http.Client) *DeleteDatacenterParams {
	return &DeleteDatacenterParams{
		HTTPClient: client,
	}
}

/* DeleteDatacenterParams contains all the parameters to send to the API endpoint
   for the delete datacenter operation.

   Typically these are written to a http.Request.
*/
type DeleteDatacenterParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.DatacenterDeletionParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete datacenter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDatacenterParams) WithDefaults() *DeleteDatacenterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete datacenter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDatacenterParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := DeleteDatacenterParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete datacenter params
func (o *DeleteDatacenterParams) WithTimeout(timeout time.Duration) *DeleteDatacenterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete datacenter params
func (o *DeleteDatacenterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete datacenter params
func (o *DeleteDatacenterParams) WithContext(ctx context.Context) *DeleteDatacenterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete datacenter params
func (o *DeleteDatacenterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete datacenter params
func (o *DeleteDatacenterParams) WithHTTPClient(client *http.Client) *DeleteDatacenterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete datacenter params
func (o *DeleteDatacenterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the delete datacenter params
func (o *DeleteDatacenterParams) WithContentLanguage(contentLanguage *string) *DeleteDatacenterParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the delete datacenter params
func (o *DeleteDatacenterParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the delete datacenter params
func (o *DeleteDatacenterParams) WithRequestBody(requestBody *models.DatacenterDeletionParams) *DeleteDatacenterParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the delete datacenter params
func (o *DeleteDatacenterParams) SetRequestBody(requestBody *models.DatacenterDeletionParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDatacenterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
