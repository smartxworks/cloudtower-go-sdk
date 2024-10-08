// Code generated by go-swagger; DO NOT EDIT.

package virtual_private_cloud_nat_gateway

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

// NewCreateVirtualPrivateCloudNatGatewayParams creates a new CreateVirtualPrivateCloudNatGatewayParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateVirtualPrivateCloudNatGatewayParams() *CreateVirtualPrivateCloudNatGatewayParams {
	return &CreateVirtualPrivateCloudNatGatewayParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVirtualPrivateCloudNatGatewayParamsWithTimeout creates a new CreateVirtualPrivateCloudNatGatewayParams object
// with the ability to set a timeout on a request.
func NewCreateVirtualPrivateCloudNatGatewayParamsWithTimeout(timeout time.Duration) *CreateVirtualPrivateCloudNatGatewayParams {
	return &CreateVirtualPrivateCloudNatGatewayParams{
		timeout: timeout,
	}
}

// NewCreateVirtualPrivateCloudNatGatewayParamsWithContext creates a new CreateVirtualPrivateCloudNatGatewayParams object
// with the ability to set a context for a request.
func NewCreateVirtualPrivateCloudNatGatewayParamsWithContext(ctx context.Context) *CreateVirtualPrivateCloudNatGatewayParams {
	return &CreateVirtualPrivateCloudNatGatewayParams{
		Context: ctx,
	}
}

// NewCreateVirtualPrivateCloudNatGatewayParamsWithHTTPClient creates a new CreateVirtualPrivateCloudNatGatewayParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateVirtualPrivateCloudNatGatewayParamsWithHTTPClient(client *http.Client) *CreateVirtualPrivateCloudNatGatewayParams {
	return &CreateVirtualPrivateCloudNatGatewayParams{
		HTTPClient: client,
	}
}

/* CreateVirtualPrivateCloudNatGatewayParams contains all the parameters to send to the API endpoint
   for the create virtual private cloud nat gateway operation.

   Typically these are written to a http.Request.
*/
type CreateVirtualPrivateCloudNatGatewayParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody []*models.VirtualPrivateCloudNatGatewayCreationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create virtual private cloud nat gateway params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVirtualPrivateCloudNatGatewayParams) WithDefaults() *CreateVirtualPrivateCloudNatGatewayParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create virtual private cloud nat gateway params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVirtualPrivateCloudNatGatewayParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := CreateVirtualPrivateCloudNatGatewayParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create virtual private cloud nat gateway params
func (o *CreateVirtualPrivateCloudNatGatewayParams) WithTimeout(timeout time.Duration) *CreateVirtualPrivateCloudNatGatewayParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create virtual private cloud nat gateway params
func (o *CreateVirtualPrivateCloudNatGatewayParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create virtual private cloud nat gateway params
func (o *CreateVirtualPrivateCloudNatGatewayParams) WithContext(ctx context.Context) *CreateVirtualPrivateCloudNatGatewayParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create virtual private cloud nat gateway params
func (o *CreateVirtualPrivateCloudNatGatewayParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create virtual private cloud nat gateway params
func (o *CreateVirtualPrivateCloudNatGatewayParams) WithHTTPClient(client *http.Client) *CreateVirtualPrivateCloudNatGatewayParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create virtual private cloud nat gateway params
func (o *CreateVirtualPrivateCloudNatGatewayParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the create virtual private cloud nat gateway params
func (o *CreateVirtualPrivateCloudNatGatewayParams) WithContentLanguage(contentLanguage *string) *CreateVirtualPrivateCloudNatGatewayParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the create virtual private cloud nat gateway params
func (o *CreateVirtualPrivateCloudNatGatewayParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the create virtual private cloud nat gateway params
func (o *CreateVirtualPrivateCloudNatGatewayParams) WithRequestBody(requestBody []*models.VirtualPrivateCloudNatGatewayCreationParams) *CreateVirtualPrivateCloudNatGatewayParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the create virtual private cloud nat gateway params
func (o *CreateVirtualPrivateCloudNatGatewayParams) SetRequestBody(requestBody []*models.VirtualPrivateCloudNatGatewayCreationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVirtualPrivateCloudNatGatewayParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
