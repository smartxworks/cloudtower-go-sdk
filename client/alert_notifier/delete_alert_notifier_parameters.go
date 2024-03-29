// Code generated by go-swagger; DO NOT EDIT.

package alert_notifier

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

// NewDeleteAlertNotifierParams creates a new DeleteAlertNotifierParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAlertNotifierParams() *DeleteAlertNotifierParams {
	return &DeleteAlertNotifierParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAlertNotifierParamsWithTimeout creates a new DeleteAlertNotifierParams object
// with the ability to set a timeout on a request.
func NewDeleteAlertNotifierParamsWithTimeout(timeout time.Duration) *DeleteAlertNotifierParams {
	return &DeleteAlertNotifierParams{
		timeout: timeout,
	}
}

// NewDeleteAlertNotifierParamsWithContext creates a new DeleteAlertNotifierParams object
// with the ability to set a context for a request.
func NewDeleteAlertNotifierParamsWithContext(ctx context.Context) *DeleteAlertNotifierParams {
	return &DeleteAlertNotifierParams{
		Context: ctx,
	}
}

// NewDeleteAlertNotifierParamsWithHTTPClient creates a new DeleteAlertNotifierParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAlertNotifierParamsWithHTTPClient(client *http.Client) *DeleteAlertNotifierParams {
	return &DeleteAlertNotifierParams{
		HTTPClient: client,
	}
}

/* DeleteAlertNotifierParams contains all the parameters to send to the API endpoint
   for the delete alert notifier operation.

   Typically these are written to a http.Request.
*/
type DeleteAlertNotifierParams struct {

	// RequestBody.
	RequestBody *models.DeleteAlertNotifierParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete alert notifier params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAlertNotifierParams) WithDefaults() *DeleteAlertNotifierParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete alert notifier params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAlertNotifierParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete alert notifier params
func (o *DeleteAlertNotifierParams) WithTimeout(timeout time.Duration) *DeleteAlertNotifierParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete alert notifier params
func (o *DeleteAlertNotifierParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete alert notifier params
func (o *DeleteAlertNotifierParams) WithContext(ctx context.Context) *DeleteAlertNotifierParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete alert notifier params
func (o *DeleteAlertNotifierParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete alert notifier params
func (o *DeleteAlertNotifierParams) WithHTTPClient(client *http.Client) *DeleteAlertNotifierParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete alert notifier params
func (o *DeleteAlertNotifierParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the delete alert notifier params
func (o *DeleteAlertNotifierParams) WithRequestBody(requestBody *models.DeleteAlertNotifierParams) *DeleteAlertNotifierParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the delete alert notifier params
func (o *DeleteAlertNotifierParams) SetRequestBody(requestBody *models.DeleteAlertNotifierParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAlertNotifierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
