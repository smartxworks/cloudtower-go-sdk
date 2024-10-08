// Code generated by go-swagger; DO NOT EDIT.

package virtual_private_cloud_router_gateway

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

// NewUpdateVirtualPrivateCloudRouterGatewayParams creates a new UpdateVirtualPrivateCloudRouterGatewayParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateVirtualPrivateCloudRouterGatewayParams() *UpdateVirtualPrivateCloudRouterGatewayParams {
	return &UpdateVirtualPrivateCloudRouterGatewayParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVirtualPrivateCloudRouterGatewayParamsWithTimeout creates a new UpdateVirtualPrivateCloudRouterGatewayParams object
// with the ability to set a timeout on a request.
func NewUpdateVirtualPrivateCloudRouterGatewayParamsWithTimeout(timeout time.Duration) *UpdateVirtualPrivateCloudRouterGatewayParams {
	return &UpdateVirtualPrivateCloudRouterGatewayParams{
		timeout: timeout,
	}
}

// NewUpdateVirtualPrivateCloudRouterGatewayParamsWithContext creates a new UpdateVirtualPrivateCloudRouterGatewayParams object
// with the ability to set a context for a request.
func NewUpdateVirtualPrivateCloudRouterGatewayParamsWithContext(ctx context.Context) *UpdateVirtualPrivateCloudRouterGatewayParams {
	return &UpdateVirtualPrivateCloudRouterGatewayParams{
		Context: ctx,
	}
}

// NewUpdateVirtualPrivateCloudRouterGatewayParamsWithHTTPClient creates a new UpdateVirtualPrivateCloudRouterGatewayParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateVirtualPrivateCloudRouterGatewayParamsWithHTTPClient(client *http.Client) *UpdateVirtualPrivateCloudRouterGatewayParams {
	return &UpdateVirtualPrivateCloudRouterGatewayParams{
		HTTPClient: client,
	}
}

/* UpdateVirtualPrivateCloudRouterGatewayParams contains all the parameters to send to the API endpoint
   for the update virtual private cloud router gateway operation.

   Typically these are written to a http.Request.
*/
type UpdateVirtualPrivateCloudRouterGatewayParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.VirtualPrivateCloudRouterGatewayUpdationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update virtual private cloud router gateway params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVirtualPrivateCloudRouterGatewayParams) WithDefaults() *UpdateVirtualPrivateCloudRouterGatewayParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update virtual private cloud router gateway params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVirtualPrivateCloudRouterGatewayParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := UpdateVirtualPrivateCloudRouterGatewayParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update virtual private cloud router gateway params
func (o *UpdateVirtualPrivateCloudRouterGatewayParams) WithTimeout(timeout time.Duration) *UpdateVirtualPrivateCloudRouterGatewayParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update virtual private cloud router gateway params
func (o *UpdateVirtualPrivateCloudRouterGatewayParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update virtual private cloud router gateway params
func (o *UpdateVirtualPrivateCloudRouterGatewayParams) WithContext(ctx context.Context) *UpdateVirtualPrivateCloudRouterGatewayParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update virtual private cloud router gateway params
func (o *UpdateVirtualPrivateCloudRouterGatewayParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update virtual private cloud router gateway params
func (o *UpdateVirtualPrivateCloudRouterGatewayParams) WithHTTPClient(client *http.Client) *UpdateVirtualPrivateCloudRouterGatewayParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update virtual private cloud router gateway params
func (o *UpdateVirtualPrivateCloudRouterGatewayParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the update virtual private cloud router gateway params
func (o *UpdateVirtualPrivateCloudRouterGatewayParams) WithContentLanguage(contentLanguage *string) *UpdateVirtualPrivateCloudRouterGatewayParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the update virtual private cloud router gateway params
func (o *UpdateVirtualPrivateCloudRouterGatewayParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the update virtual private cloud router gateway params
func (o *UpdateVirtualPrivateCloudRouterGatewayParams) WithRequestBody(requestBody *models.VirtualPrivateCloudRouterGatewayUpdationParams) *UpdateVirtualPrivateCloudRouterGatewayParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update virtual private cloud router gateway params
func (o *UpdateVirtualPrivateCloudRouterGatewayParams) SetRequestBody(requestBody *models.VirtualPrivateCloudRouterGatewayUpdationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVirtualPrivateCloudRouterGatewayParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
