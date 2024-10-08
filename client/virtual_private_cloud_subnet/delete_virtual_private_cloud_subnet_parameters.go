// Code generated by go-swagger; DO NOT EDIT.

package virtual_private_cloud_subnet

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

// NewDeleteVirtualPrivateCloudSubnetParams creates a new DeleteVirtualPrivateCloudSubnetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteVirtualPrivateCloudSubnetParams() *DeleteVirtualPrivateCloudSubnetParams {
	return &DeleteVirtualPrivateCloudSubnetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVirtualPrivateCloudSubnetParamsWithTimeout creates a new DeleteVirtualPrivateCloudSubnetParams object
// with the ability to set a timeout on a request.
func NewDeleteVirtualPrivateCloudSubnetParamsWithTimeout(timeout time.Duration) *DeleteVirtualPrivateCloudSubnetParams {
	return &DeleteVirtualPrivateCloudSubnetParams{
		timeout: timeout,
	}
}

// NewDeleteVirtualPrivateCloudSubnetParamsWithContext creates a new DeleteVirtualPrivateCloudSubnetParams object
// with the ability to set a context for a request.
func NewDeleteVirtualPrivateCloudSubnetParamsWithContext(ctx context.Context) *DeleteVirtualPrivateCloudSubnetParams {
	return &DeleteVirtualPrivateCloudSubnetParams{
		Context: ctx,
	}
}

// NewDeleteVirtualPrivateCloudSubnetParamsWithHTTPClient creates a new DeleteVirtualPrivateCloudSubnetParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteVirtualPrivateCloudSubnetParamsWithHTTPClient(client *http.Client) *DeleteVirtualPrivateCloudSubnetParams {
	return &DeleteVirtualPrivateCloudSubnetParams{
		HTTPClient: client,
	}
}

/* DeleteVirtualPrivateCloudSubnetParams contains all the parameters to send to the API endpoint
   for the delete virtual private cloud subnet operation.

   Typically these are written to a http.Request.
*/
type DeleteVirtualPrivateCloudSubnetParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.VirtualPrivateCloudSubnetDeletionParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete virtual private cloud subnet params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVirtualPrivateCloudSubnetParams) WithDefaults() *DeleteVirtualPrivateCloudSubnetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete virtual private cloud subnet params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVirtualPrivateCloudSubnetParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := DeleteVirtualPrivateCloudSubnetParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete virtual private cloud subnet params
func (o *DeleteVirtualPrivateCloudSubnetParams) WithTimeout(timeout time.Duration) *DeleteVirtualPrivateCloudSubnetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete virtual private cloud subnet params
func (o *DeleteVirtualPrivateCloudSubnetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete virtual private cloud subnet params
func (o *DeleteVirtualPrivateCloudSubnetParams) WithContext(ctx context.Context) *DeleteVirtualPrivateCloudSubnetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete virtual private cloud subnet params
func (o *DeleteVirtualPrivateCloudSubnetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete virtual private cloud subnet params
func (o *DeleteVirtualPrivateCloudSubnetParams) WithHTTPClient(client *http.Client) *DeleteVirtualPrivateCloudSubnetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete virtual private cloud subnet params
func (o *DeleteVirtualPrivateCloudSubnetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the delete virtual private cloud subnet params
func (o *DeleteVirtualPrivateCloudSubnetParams) WithContentLanguage(contentLanguage *string) *DeleteVirtualPrivateCloudSubnetParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the delete virtual private cloud subnet params
func (o *DeleteVirtualPrivateCloudSubnetParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the delete virtual private cloud subnet params
func (o *DeleteVirtualPrivateCloudSubnetParams) WithRequestBody(requestBody *models.VirtualPrivateCloudSubnetDeletionParams) *DeleteVirtualPrivateCloudSubnetParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the delete virtual private cloud subnet params
func (o *DeleteVirtualPrivateCloudSubnetParams) SetRequestBody(requestBody *models.VirtualPrivateCloudSubnetDeletionParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVirtualPrivateCloudSubnetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
