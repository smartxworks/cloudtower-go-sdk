// Code generated by go-swagger; DO NOT EDIT.

package backup_service

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

// NewGetBackupServicesParams creates a new GetBackupServicesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBackupServicesParams() *GetBackupServicesParams {
	return &GetBackupServicesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBackupServicesParamsWithTimeout creates a new GetBackupServicesParams object
// with the ability to set a timeout on a request.
func NewGetBackupServicesParamsWithTimeout(timeout time.Duration) *GetBackupServicesParams {
	return &GetBackupServicesParams{
		timeout: timeout,
	}
}

// NewGetBackupServicesParamsWithContext creates a new GetBackupServicesParams object
// with the ability to set a context for a request.
func NewGetBackupServicesParamsWithContext(ctx context.Context) *GetBackupServicesParams {
	return &GetBackupServicesParams{
		Context: ctx,
	}
}

// NewGetBackupServicesParamsWithHTTPClient creates a new GetBackupServicesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBackupServicesParamsWithHTTPClient(client *http.Client) *GetBackupServicesParams {
	return &GetBackupServicesParams{
		HTTPClient: client,
	}
}

/* GetBackupServicesParams contains all the parameters to send to the API endpoint
   for the get backup services operation.

   Typically these are written to a http.Request.
*/
type GetBackupServicesParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetBackupServicesRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get backup services params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBackupServicesParams) WithDefaults() *GetBackupServicesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get backup services params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBackupServicesParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetBackupServicesParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get backup services params
func (o *GetBackupServicesParams) WithTimeout(timeout time.Duration) *GetBackupServicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get backup services params
func (o *GetBackupServicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get backup services params
func (o *GetBackupServicesParams) WithContext(ctx context.Context) *GetBackupServicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get backup services params
func (o *GetBackupServicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get backup services params
func (o *GetBackupServicesParams) WithHTTPClient(client *http.Client) *GetBackupServicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get backup services params
func (o *GetBackupServicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get backup services params
func (o *GetBackupServicesParams) WithContentLanguage(contentLanguage *string) *GetBackupServicesParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get backup services params
func (o *GetBackupServicesParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get backup services params
func (o *GetBackupServicesParams) WithRequestBody(requestBody *models.GetBackupServicesRequestBody) *GetBackupServicesParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get backup services params
func (o *GetBackupServicesParams) SetRequestBody(requestBody *models.GetBackupServicesRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetBackupServicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
