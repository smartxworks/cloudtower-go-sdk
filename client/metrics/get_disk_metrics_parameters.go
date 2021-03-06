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

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// NewGetDiskMetricsParams creates a new GetDiskMetricsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDiskMetricsParams() *GetDiskMetricsParams {
	return &GetDiskMetricsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDiskMetricsParamsWithTimeout creates a new GetDiskMetricsParams object
// with the ability to set a timeout on a request.
func NewGetDiskMetricsParamsWithTimeout(timeout time.Duration) *GetDiskMetricsParams {
	return &GetDiskMetricsParams{
		timeout: timeout,
	}
}

// NewGetDiskMetricsParamsWithContext creates a new GetDiskMetricsParams object
// with the ability to set a context for a request.
func NewGetDiskMetricsParamsWithContext(ctx context.Context) *GetDiskMetricsParams {
	return &GetDiskMetricsParams{
		Context: ctx,
	}
}

// NewGetDiskMetricsParamsWithHTTPClient creates a new GetDiskMetricsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDiskMetricsParamsWithHTTPClient(client *http.Client) *GetDiskMetricsParams {
	return &GetDiskMetricsParams{
		HTTPClient: client,
	}
}

/* GetDiskMetricsParams contains all the parameters to send to the API endpoint
   for the get disk metrics operation.

   Typically these are written to a http.Request.
*/
type GetDiskMetricsParams struct {

	// RequestBody.
	RequestBody *models.GetDiskMetricInput

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get disk metrics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDiskMetricsParams) WithDefaults() *GetDiskMetricsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get disk metrics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDiskMetricsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get disk metrics params
func (o *GetDiskMetricsParams) WithTimeout(timeout time.Duration) *GetDiskMetricsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get disk metrics params
func (o *GetDiskMetricsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get disk metrics params
func (o *GetDiskMetricsParams) WithContext(ctx context.Context) *GetDiskMetricsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get disk metrics params
func (o *GetDiskMetricsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get disk metrics params
func (o *GetDiskMetricsParams) WithHTTPClient(client *http.Client) *GetDiskMetricsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get disk metrics params
func (o *GetDiskMetricsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the get disk metrics params
func (o *GetDiskMetricsParams) WithRequestBody(requestBody *models.GetDiskMetricInput) *GetDiskMetricsParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get disk metrics params
func (o *GetDiskMetricsParams) SetRequestBody(requestBody *models.GetDiskMetricInput) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetDiskMetricsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
