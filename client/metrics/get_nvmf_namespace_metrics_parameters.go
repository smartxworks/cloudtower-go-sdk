// Code generated by go-swagger; DO NOT EDIT.

package metrics

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

// NewGetNvmfNamespaceMetricsParams creates a new GetNvmfNamespaceMetricsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNvmfNamespaceMetricsParams() *GetNvmfNamespaceMetricsParams {
	return &GetNvmfNamespaceMetricsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNvmfNamespaceMetricsParamsWithTimeout creates a new GetNvmfNamespaceMetricsParams object
// with the ability to set a timeout on a request.
func NewGetNvmfNamespaceMetricsParamsWithTimeout(timeout time.Duration) *GetNvmfNamespaceMetricsParams {
	return &GetNvmfNamespaceMetricsParams{
		timeout: timeout,
	}
}

// NewGetNvmfNamespaceMetricsParamsWithContext creates a new GetNvmfNamespaceMetricsParams object
// with the ability to set a context for a request.
func NewGetNvmfNamespaceMetricsParamsWithContext(ctx context.Context) *GetNvmfNamespaceMetricsParams {
	return &GetNvmfNamespaceMetricsParams{
		Context: ctx,
	}
}

// NewGetNvmfNamespaceMetricsParamsWithHTTPClient creates a new GetNvmfNamespaceMetricsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNvmfNamespaceMetricsParamsWithHTTPClient(client *http.Client) *GetNvmfNamespaceMetricsParams {
	return &GetNvmfNamespaceMetricsParams{
		HTTPClient: client,
	}
}

/* GetNvmfNamespaceMetricsParams contains all the parameters to send to the API endpoint
   for the get nvmf namespace metrics operation.

   Typically these are written to a http.Request.
*/
type GetNvmfNamespaceMetricsParams struct {

	// RequestBody.
	RequestBody *models.GetNvmfNamespaceMetricInput

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get nvmf namespace metrics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNvmfNamespaceMetricsParams) WithDefaults() *GetNvmfNamespaceMetricsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get nvmf namespace metrics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNvmfNamespaceMetricsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get nvmf namespace metrics params
func (o *GetNvmfNamespaceMetricsParams) WithTimeout(timeout time.Duration) *GetNvmfNamespaceMetricsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nvmf namespace metrics params
func (o *GetNvmfNamespaceMetricsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nvmf namespace metrics params
func (o *GetNvmfNamespaceMetricsParams) WithContext(ctx context.Context) *GetNvmfNamespaceMetricsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nvmf namespace metrics params
func (o *GetNvmfNamespaceMetricsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nvmf namespace metrics params
func (o *GetNvmfNamespaceMetricsParams) WithHTTPClient(client *http.Client) *GetNvmfNamespaceMetricsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nvmf namespace metrics params
func (o *GetNvmfNamespaceMetricsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get nvmf namespace metrics params
func (o *GetNvmfNamespaceMetricsParams) WithRequestBody(requestBody *models.GetNvmfNamespaceMetricInput) *GetNvmfNamespaceMetricsParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get nvmf namespace metrics params
func (o *GetNvmfNamespaceMetricsParams) SetRequestBody(requestBody *models.GetNvmfNamespaceMetricInput) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetNvmfNamespaceMetricsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
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
