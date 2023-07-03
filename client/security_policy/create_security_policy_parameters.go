// Code generated by go-swagger; DO NOT EDIT.

package security_policy

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

// NewCreateSecurityPolicyParams creates a new CreateSecurityPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateSecurityPolicyParams() *CreateSecurityPolicyParams {
	return &CreateSecurityPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSecurityPolicyParamsWithTimeout creates a new CreateSecurityPolicyParams object
// with the ability to set a timeout on a request.
func NewCreateSecurityPolicyParamsWithTimeout(timeout time.Duration) *CreateSecurityPolicyParams {
	return &CreateSecurityPolicyParams{
		timeout: timeout,
	}
}

// NewCreateSecurityPolicyParamsWithContext creates a new CreateSecurityPolicyParams object
// with the ability to set a context for a request.
func NewCreateSecurityPolicyParamsWithContext(ctx context.Context) *CreateSecurityPolicyParams {
	return &CreateSecurityPolicyParams{
		Context: ctx,
	}
}

// NewCreateSecurityPolicyParamsWithHTTPClient creates a new CreateSecurityPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateSecurityPolicyParamsWithHTTPClient(client *http.Client) *CreateSecurityPolicyParams {
	return &CreateSecurityPolicyParams{
		HTTPClient: client,
	}
}

/* CreateSecurityPolicyParams contains all the parameters to send to the API endpoint
   for the create security policy operation.

   Typically these are written to a http.Request.
*/
type CreateSecurityPolicyParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.SecurityPolicyCreateParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create security policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSecurityPolicyParams) WithDefaults() *CreateSecurityPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create security policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSecurityPolicyParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := CreateSecurityPolicyParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create security policy params
func (o *CreateSecurityPolicyParams) WithTimeout(timeout time.Duration) *CreateSecurityPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create security policy params
func (o *CreateSecurityPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create security policy params
func (o *CreateSecurityPolicyParams) WithContext(ctx context.Context) *CreateSecurityPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create security policy params
func (o *CreateSecurityPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create security policy params
func (o *CreateSecurityPolicyParams) WithHTTPClient(client *http.Client) *CreateSecurityPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create security policy params
func (o *CreateSecurityPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the create security policy params
func (o *CreateSecurityPolicyParams) WithContentLanguage(contentLanguage *string) *CreateSecurityPolicyParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the create security policy params
func (o *CreateSecurityPolicyParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the create security policy params
func (o *CreateSecurityPolicyParams) WithRequestBody(requestBody *models.SecurityPolicyCreateParams) *CreateSecurityPolicyParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the create security policy params
func (o *CreateSecurityPolicyParams) SetRequestBody(requestBody *models.SecurityPolicyCreateParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSecurityPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
