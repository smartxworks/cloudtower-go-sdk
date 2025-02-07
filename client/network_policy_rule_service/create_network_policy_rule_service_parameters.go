// Code generated by go-swagger; DO NOT EDIT.

package network_policy_rule_service

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

// NewCreateNetworkPolicyRuleServiceParams creates a new CreateNetworkPolicyRuleServiceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateNetworkPolicyRuleServiceParams() *CreateNetworkPolicyRuleServiceParams {
	return &CreateNetworkPolicyRuleServiceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNetworkPolicyRuleServiceParamsWithTimeout creates a new CreateNetworkPolicyRuleServiceParams object
// with the ability to set a timeout on a request.
func NewCreateNetworkPolicyRuleServiceParamsWithTimeout(timeout time.Duration) *CreateNetworkPolicyRuleServiceParams {
	return &CreateNetworkPolicyRuleServiceParams{
		timeout: timeout,
	}
}

// NewCreateNetworkPolicyRuleServiceParamsWithContext creates a new CreateNetworkPolicyRuleServiceParams object
// with the ability to set a context for a request.
func NewCreateNetworkPolicyRuleServiceParamsWithContext(ctx context.Context) *CreateNetworkPolicyRuleServiceParams {
	return &CreateNetworkPolicyRuleServiceParams{
		Context: ctx,
	}
}

// NewCreateNetworkPolicyRuleServiceParamsWithHTTPClient creates a new CreateNetworkPolicyRuleServiceParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateNetworkPolicyRuleServiceParamsWithHTTPClient(client *http.Client) *CreateNetworkPolicyRuleServiceParams {
	return &CreateNetworkPolicyRuleServiceParams{
		HTTPClient: client,
	}
}

/* CreateNetworkPolicyRuleServiceParams contains all the parameters to send to the API endpoint
   for the create network policy rule service operation.

   Typically these are written to a http.Request.
*/
type CreateNetworkPolicyRuleServiceParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody []*models.NetworkPolicyRuleServiceCreationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create network policy rule service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNetworkPolicyRuleServiceParams) WithDefaults() *CreateNetworkPolicyRuleServiceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create network policy rule service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNetworkPolicyRuleServiceParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := CreateNetworkPolicyRuleServiceParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create network policy rule service params
func (o *CreateNetworkPolicyRuleServiceParams) WithTimeout(timeout time.Duration) *CreateNetworkPolicyRuleServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create network policy rule service params
func (o *CreateNetworkPolicyRuleServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create network policy rule service params
func (o *CreateNetworkPolicyRuleServiceParams) WithContext(ctx context.Context) *CreateNetworkPolicyRuleServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create network policy rule service params
func (o *CreateNetworkPolicyRuleServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create network policy rule service params
func (o *CreateNetworkPolicyRuleServiceParams) WithHTTPClient(client *http.Client) *CreateNetworkPolicyRuleServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create network policy rule service params
func (o *CreateNetworkPolicyRuleServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the create network policy rule service params
func (o *CreateNetworkPolicyRuleServiceParams) WithContentLanguage(contentLanguage *string) *CreateNetworkPolicyRuleServiceParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the create network policy rule service params
func (o *CreateNetworkPolicyRuleServiceParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the create network policy rule service params
func (o *CreateNetworkPolicyRuleServiceParams) WithRequestBody(requestBody []*models.NetworkPolicyRuleServiceCreationParams) *CreateNetworkPolicyRuleServiceParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the create network policy rule service params
func (o *CreateNetworkPolicyRuleServiceParams) SetRequestBody(requestBody []*models.NetworkPolicyRuleServiceCreationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNetworkPolicyRuleServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
