// Code generated by go-swagger; DO NOT EDIT.

package virtual_private_cloud_edge_gateway

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

// NewGetVirtualPrivateCloudEdgeGatewaysParams creates a new GetVirtualPrivateCloudEdgeGatewaysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVirtualPrivateCloudEdgeGatewaysParams() *GetVirtualPrivateCloudEdgeGatewaysParams {
	return &GetVirtualPrivateCloudEdgeGatewaysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVirtualPrivateCloudEdgeGatewaysParamsWithTimeout creates a new GetVirtualPrivateCloudEdgeGatewaysParams object
// with the ability to set a timeout on a request.
func NewGetVirtualPrivateCloudEdgeGatewaysParamsWithTimeout(timeout time.Duration) *GetVirtualPrivateCloudEdgeGatewaysParams {
	return &GetVirtualPrivateCloudEdgeGatewaysParams{
		timeout: timeout,
	}
}

// NewGetVirtualPrivateCloudEdgeGatewaysParamsWithContext creates a new GetVirtualPrivateCloudEdgeGatewaysParams object
// with the ability to set a context for a request.
func NewGetVirtualPrivateCloudEdgeGatewaysParamsWithContext(ctx context.Context) *GetVirtualPrivateCloudEdgeGatewaysParams {
	return &GetVirtualPrivateCloudEdgeGatewaysParams{
		Context: ctx,
	}
}

// NewGetVirtualPrivateCloudEdgeGatewaysParamsWithHTTPClient creates a new GetVirtualPrivateCloudEdgeGatewaysParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVirtualPrivateCloudEdgeGatewaysParamsWithHTTPClient(client *http.Client) *GetVirtualPrivateCloudEdgeGatewaysParams {
	return &GetVirtualPrivateCloudEdgeGatewaysParams{
		HTTPClient: client,
	}
}

/* GetVirtualPrivateCloudEdgeGatewaysParams contains all the parameters to send to the API endpoint
   for the get virtual private cloud edge gateways operation.

   Typically these are written to a http.Request.
*/
type GetVirtualPrivateCloudEdgeGatewaysParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetVirtualPrivateCloudEdgeGatewaysRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get virtual private cloud edge gateways params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVirtualPrivateCloudEdgeGatewaysParams) WithDefaults() *GetVirtualPrivateCloudEdgeGatewaysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get virtual private cloud edge gateways params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVirtualPrivateCloudEdgeGatewaysParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetVirtualPrivateCloudEdgeGatewaysParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get virtual private cloud edge gateways params
func (o *GetVirtualPrivateCloudEdgeGatewaysParams) WithTimeout(timeout time.Duration) *GetVirtualPrivateCloudEdgeGatewaysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get virtual private cloud edge gateways params
func (o *GetVirtualPrivateCloudEdgeGatewaysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get virtual private cloud edge gateways params
func (o *GetVirtualPrivateCloudEdgeGatewaysParams) WithContext(ctx context.Context) *GetVirtualPrivateCloudEdgeGatewaysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get virtual private cloud edge gateways params
func (o *GetVirtualPrivateCloudEdgeGatewaysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get virtual private cloud edge gateways params
func (o *GetVirtualPrivateCloudEdgeGatewaysParams) WithHTTPClient(client *http.Client) *GetVirtualPrivateCloudEdgeGatewaysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get virtual private cloud edge gateways params
func (o *GetVirtualPrivateCloudEdgeGatewaysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get virtual private cloud edge gateways params
func (o *GetVirtualPrivateCloudEdgeGatewaysParams) WithContentLanguage(contentLanguage *string) *GetVirtualPrivateCloudEdgeGatewaysParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get virtual private cloud edge gateways params
func (o *GetVirtualPrivateCloudEdgeGatewaysParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get virtual private cloud edge gateways params
func (o *GetVirtualPrivateCloudEdgeGatewaysParams) WithRequestBody(requestBody *models.GetVirtualPrivateCloudEdgeGatewaysRequestBody) *GetVirtualPrivateCloudEdgeGatewaysParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get virtual private cloud edge gateways params
func (o *GetVirtualPrivateCloudEdgeGatewaysParams) SetRequestBody(requestBody *models.GetVirtualPrivateCloudEdgeGatewaysRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetVirtualPrivateCloudEdgeGatewaysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
