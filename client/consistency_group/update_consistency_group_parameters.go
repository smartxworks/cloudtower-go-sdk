// Code generated by go-swagger; DO NOT EDIT.

package consistency_group

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

// NewUpdateConsistencyGroupParams creates a new UpdateConsistencyGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateConsistencyGroupParams() *UpdateConsistencyGroupParams {
	return &UpdateConsistencyGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateConsistencyGroupParamsWithTimeout creates a new UpdateConsistencyGroupParams object
// with the ability to set a timeout on a request.
func NewUpdateConsistencyGroupParamsWithTimeout(timeout time.Duration) *UpdateConsistencyGroupParams {
	return &UpdateConsistencyGroupParams{
		timeout: timeout,
	}
}

// NewUpdateConsistencyGroupParamsWithContext creates a new UpdateConsistencyGroupParams object
// with the ability to set a context for a request.
func NewUpdateConsistencyGroupParamsWithContext(ctx context.Context) *UpdateConsistencyGroupParams {
	return &UpdateConsistencyGroupParams{
		Context: ctx,
	}
}

// NewUpdateConsistencyGroupParamsWithHTTPClient creates a new UpdateConsistencyGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateConsistencyGroupParamsWithHTTPClient(client *http.Client) *UpdateConsistencyGroupParams {
	return &UpdateConsistencyGroupParams{
		HTTPClient: client,
	}
}

/* UpdateConsistencyGroupParams contains all the parameters to send to the API endpoint
   for the update consistency group operation.

   Typically these are written to a http.Request.
*/
type UpdateConsistencyGroupParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.ConsistencyGroupUpdationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update consistency group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateConsistencyGroupParams) WithDefaults() *UpdateConsistencyGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update consistency group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateConsistencyGroupParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := UpdateConsistencyGroupParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update consistency group params
func (o *UpdateConsistencyGroupParams) WithTimeout(timeout time.Duration) *UpdateConsistencyGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update consistency group params
func (o *UpdateConsistencyGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update consistency group params
func (o *UpdateConsistencyGroupParams) WithContext(ctx context.Context) *UpdateConsistencyGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update consistency group params
func (o *UpdateConsistencyGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update consistency group params
func (o *UpdateConsistencyGroupParams) WithHTTPClient(client *http.Client) *UpdateConsistencyGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update consistency group params
func (o *UpdateConsistencyGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the update consistency group params
func (o *UpdateConsistencyGroupParams) WithContentLanguage(contentLanguage *string) *UpdateConsistencyGroupParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the update consistency group params
func (o *UpdateConsistencyGroupParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the update consistency group params
func (o *UpdateConsistencyGroupParams) WithRequestBody(requestBody *models.ConsistencyGroupUpdationParams) *UpdateConsistencyGroupParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update consistency group params
func (o *UpdateConsistencyGroupParams) SetRequestBody(requestBody *models.ConsistencyGroupUpdationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateConsistencyGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
