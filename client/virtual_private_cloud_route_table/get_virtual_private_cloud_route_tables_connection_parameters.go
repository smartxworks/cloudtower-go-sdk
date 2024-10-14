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

// NewGetVirtualPrivateCloudRouteTablesConnectionParams creates a new GetVirtualPrivateCloudRouteTablesConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVirtualPrivateCloudRouteTablesConnectionParams() *GetVirtualPrivateCloudRouteTablesConnectionParams {
	return &GetVirtualPrivateCloudRouteTablesConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVirtualPrivateCloudRouteTablesConnectionParamsWithTimeout creates a new GetVirtualPrivateCloudRouteTablesConnectionParams object
// with the ability to set a timeout on a request.
func NewGetVirtualPrivateCloudRouteTablesConnectionParamsWithTimeout(timeout time.Duration) *GetVirtualPrivateCloudRouteTablesConnectionParams {
	return &GetVirtualPrivateCloudRouteTablesConnectionParams{
		timeout: timeout,
	}
}

// NewGetVirtualPrivateCloudRouteTablesConnectionParamsWithContext creates a new GetVirtualPrivateCloudRouteTablesConnectionParams object
// with the ability to set a context for a request.
func NewGetVirtualPrivateCloudRouteTablesConnectionParamsWithContext(ctx context.Context) *GetVirtualPrivateCloudRouteTablesConnectionParams {
	return &GetVirtualPrivateCloudRouteTablesConnectionParams{
		Context: ctx,
	}
}

// NewGetVirtualPrivateCloudRouteTablesConnectionParamsWithHTTPClient creates a new GetVirtualPrivateCloudRouteTablesConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVirtualPrivateCloudRouteTablesConnectionParamsWithHTTPClient(client *http.Client) *GetVirtualPrivateCloudRouteTablesConnectionParams {
	return &GetVirtualPrivateCloudRouteTablesConnectionParams{
		HTTPClient: client,
	}
}

/* GetVirtualPrivateCloudRouteTablesConnectionParams contains all the parameters to send to the API endpoint
   for the get virtual private cloud route tables connection operation.

   Typically these are written to a http.Request.
*/
type GetVirtualPrivateCloudRouteTablesConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetVirtualPrivateCloudRouteTablesConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get virtual private cloud route tables connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVirtualPrivateCloudRouteTablesConnectionParams) WithDefaults() *GetVirtualPrivateCloudRouteTablesConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get virtual private cloud route tables connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVirtualPrivateCloudRouteTablesConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetVirtualPrivateCloudRouteTablesConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get virtual private cloud route tables connection params
func (o *GetVirtualPrivateCloudRouteTablesConnectionParams) WithTimeout(timeout time.Duration) *GetVirtualPrivateCloudRouteTablesConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get virtual private cloud route tables connection params
func (o *GetVirtualPrivateCloudRouteTablesConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get virtual private cloud route tables connection params
func (o *GetVirtualPrivateCloudRouteTablesConnectionParams) WithContext(ctx context.Context) *GetVirtualPrivateCloudRouteTablesConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get virtual private cloud route tables connection params
func (o *GetVirtualPrivateCloudRouteTablesConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get virtual private cloud route tables connection params
func (o *GetVirtualPrivateCloudRouteTablesConnectionParams) WithHTTPClient(client *http.Client) *GetVirtualPrivateCloudRouteTablesConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get virtual private cloud route tables connection params
func (o *GetVirtualPrivateCloudRouteTablesConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get virtual private cloud route tables connection params
func (o *GetVirtualPrivateCloudRouteTablesConnectionParams) WithContentLanguage(contentLanguage *string) *GetVirtualPrivateCloudRouteTablesConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get virtual private cloud route tables connection params
func (o *GetVirtualPrivateCloudRouteTablesConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get virtual private cloud route tables connection params
func (o *GetVirtualPrivateCloudRouteTablesConnectionParams) WithRequestBody(requestBody *models.GetVirtualPrivateCloudRouteTablesConnectionRequestBody) *GetVirtualPrivateCloudRouteTablesConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get virtual private cloud route tables connection params
func (o *GetVirtualPrivateCloudRouteTablesConnectionParams) SetRequestBody(requestBody *models.GetVirtualPrivateCloudRouteTablesConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetVirtualPrivateCloudRouteTablesConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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