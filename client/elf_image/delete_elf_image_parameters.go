// Code generated by go-swagger; DO NOT EDIT.

package elf_image

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

// NewDeleteElfImageParams creates a new DeleteElfImageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteElfImageParams() *DeleteElfImageParams {
	return &DeleteElfImageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteElfImageParamsWithTimeout creates a new DeleteElfImageParams object
// with the ability to set a timeout on a request.
func NewDeleteElfImageParamsWithTimeout(timeout time.Duration) *DeleteElfImageParams {
	return &DeleteElfImageParams{
		timeout: timeout,
	}
}

// NewDeleteElfImageParamsWithContext creates a new DeleteElfImageParams object
// with the ability to set a context for a request.
func NewDeleteElfImageParamsWithContext(ctx context.Context) *DeleteElfImageParams {
	return &DeleteElfImageParams{
		Context: ctx,
	}
}

// NewDeleteElfImageParamsWithHTTPClient creates a new DeleteElfImageParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteElfImageParamsWithHTTPClient(client *http.Client) *DeleteElfImageParams {
	return &DeleteElfImageParams{
		HTTPClient: client,
	}
}

/* DeleteElfImageParams contains all the parameters to send to the API endpoint
   for the delete elf image operation.

   Typically these are written to a http.Request.
*/
type DeleteElfImageParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.ElfImageDeletionParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete elf image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteElfImageParams) WithDefaults() *DeleteElfImageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete elf image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteElfImageParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := DeleteElfImageParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete elf image params
func (o *DeleteElfImageParams) WithTimeout(timeout time.Duration) *DeleteElfImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete elf image params
func (o *DeleteElfImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete elf image params
func (o *DeleteElfImageParams) WithContext(ctx context.Context) *DeleteElfImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete elf image params
func (o *DeleteElfImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete elf image params
func (o *DeleteElfImageParams) WithHTTPClient(client *http.Client) *DeleteElfImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete elf image params
func (o *DeleteElfImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the delete elf image params
func (o *DeleteElfImageParams) WithContentLanguage(contentLanguage *string) *DeleteElfImageParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the delete elf image params
func (o *DeleteElfImageParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the delete elf image params
func (o *DeleteElfImageParams) WithRequestBody(requestBody *models.ElfImageDeletionParams) *DeleteElfImageParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the delete elf image params
func (o *DeleteElfImageParams) SetRequestBody(requestBody *models.ElfImageDeletionParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteElfImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
