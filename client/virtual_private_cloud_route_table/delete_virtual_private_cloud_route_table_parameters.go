// Code generated by go-swagger; DO NOT EDIT.

package virtual_private_cloud_route_table

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

// NewDeleteVirtualPrivateCloudRouteTableParams creates a new DeleteVirtualPrivateCloudRouteTableParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteVirtualPrivateCloudRouteTableParams() *DeleteVirtualPrivateCloudRouteTableParams {
	return &DeleteVirtualPrivateCloudRouteTableParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVirtualPrivateCloudRouteTableParamsWithTimeout creates a new DeleteVirtualPrivateCloudRouteTableParams object
// with the ability to set a timeout on a request.
func NewDeleteVirtualPrivateCloudRouteTableParamsWithTimeout(timeout time.Duration) *DeleteVirtualPrivateCloudRouteTableParams {
	return &DeleteVirtualPrivateCloudRouteTableParams{
		timeout: timeout,
	}
}

// NewDeleteVirtualPrivateCloudRouteTableParamsWithContext creates a new DeleteVirtualPrivateCloudRouteTableParams object
// with the ability to set a context for a request.
func NewDeleteVirtualPrivateCloudRouteTableParamsWithContext(ctx context.Context) *DeleteVirtualPrivateCloudRouteTableParams {
	return &DeleteVirtualPrivateCloudRouteTableParams{
		Context: ctx,
	}
}

// NewDeleteVirtualPrivateCloudRouteTableParamsWithHTTPClient creates a new DeleteVirtualPrivateCloudRouteTableParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteVirtualPrivateCloudRouteTableParamsWithHTTPClient(client *http.Client) *DeleteVirtualPrivateCloudRouteTableParams {
	return &DeleteVirtualPrivateCloudRouteTableParams{
		HTTPClient: client,
	}
}

/* DeleteVirtualPrivateCloudRouteTableParams contains all the parameters to send to the API endpoint
   for the delete virtual private cloud route table operation.

   Typically these are written to a http.Request.
*/
type DeleteVirtualPrivateCloudRouteTableParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.VirtualPrivateCloudRouteTableDeletionParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete virtual private cloud route table params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVirtualPrivateCloudRouteTableParams) WithDefaults() *DeleteVirtualPrivateCloudRouteTableParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete virtual private cloud route table params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVirtualPrivateCloudRouteTableParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := DeleteVirtualPrivateCloudRouteTableParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete virtual private cloud route table params
func (o *DeleteVirtualPrivateCloudRouteTableParams) WithTimeout(timeout time.Duration) *DeleteVirtualPrivateCloudRouteTableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete virtual private cloud route table params
func (o *DeleteVirtualPrivateCloudRouteTableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete virtual private cloud route table params
func (o *DeleteVirtualPrivateCloudRouteTableParams) WithContext(ctx context.Context) *DeleteVirtualPrivateCloudRouteTableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete virtual private cloud route table params
func (o *DeleteVirtualPrivateCloudRouteTableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete virtual private cloud route table params
func (o *DeleteVirtualPrivateCloudRouteTableParams) WithHTTPClient(client *http.Client) *DeleteVirtualPrivateCloudRouteTableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete virtual private cloud route table params
func (o *DeleteVirtualPrivateCloudRouteTableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the delete virtual private cloud route table params
func (o *DeleteVirtualPrivateCloudRouteTableParams) WithContentLanguage(contentLanguage *string) *DeleteVirtualPrivateCloudRouteTableParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the delete virtual private cloud route table params
func (o *DeleteVirtualPrivateCloudRouteTableParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the delete virtual private cloud route table params
func (o *DeleteVirtualPrivateCloudRouteTableParams) WithRequestBody(requestBody *models.VirtualPrivateCloudRouteTableDeletionParams) *DeleteVirtualPrivateCloudRouteTableParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the delete virtual private cloud route table params
func (o *DeleteVirtualPrivateCloudRouteTableParams) SetRequestBody(requestBody *models.VirtualPrivateCloudRouteTableDeletionParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVirtualPrivateCloudRouteTableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
