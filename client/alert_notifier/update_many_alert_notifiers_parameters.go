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

// NewUpdateManyAlertNotifiersParams creates a new UpdateManyAlertNotifiersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateManyAlertNotifiersParams() *UpdateManyAlertNotifiersParams {
	return &UpdateManyAlertNotifiersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateManyAlertNotifiersParamsWithTimeout creates a new UpdateManyAlertNotifiersParams object
// with the ability to set a timeout on a request.
func NewUpdateManyAlertNotifiersParamsWithTimeout(timeout time.Duration) *UpdateManyAlertNotifiersParams {
	return &UpdateManyAlertNotifiersParams{
		timeout: timeout,
	}
}

// NewUpdateManyAlertNotifiersParamsWithContext creates a new UpdateManyAlertNotifiersParams object
// with the ability to set a context for a request.
func NewUpdateManyAlertNotifiersParamsWithContext(ctx context.Context) *UpdateManyAlertNotifiersParams {
	return &UpdateManyAlertNotifiersParams{
		Context: ctx,
	}
}

// NewUpdateManyAlertNotifiersParamsWithHTTPClient creates a new UpdateManyAlertNotifiersParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateManyAlertNotifiersParamsWithHTTPClient(client *http.Client) *UpdateManyAlertNotifiersParams {
	return &UpdateManyAlertNotifiersParams{
		HTTPClient: client,
	}
}

/* UpdateManyAlertNotifiersParams contains all the parameters to send to the API endpoint
   for the update many alert notifiers operation.

   Typically these are written to a http.Request.
*/
type UpdateManyAlertNotifiersParams struct {

	// RequestBody.
	RequestBody *models.AlertNotifierManyUpdationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update many alert notifiers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateManyAlertNotifiersParams) WithDefaults() *UpdateManyAlertNotifiersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update many alert notifiers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateManyAlertNotifiersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update many alert notifiers params
func (o *UpdateManyAlertNotifiersParams) WithTimeout(timeout time.Duration) *UpdateManyAlertNotifiersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update many alert notifiers params
func (o *UpdateManyAlertNotifiersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update many alert notifiers params
func (o *UpdateManyAlertNotifiersParams) WithContext(ctx context.Context) *UpdateManyAlertNotifiersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update many alert notifiers params
func (o *UpdateManyAlertNotifiersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update many alert notifiers params
func (o *UpdateManyAlertNotifiersParams) WithHTTPClient(client *http.Client) *UpdateManyAlertNotifiersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update many alert notifiers params
func (o *UpdateManyAlertNotifiersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestBody adds the requestBody to the update many alert notifiers params
func (o *UpdateManyAlertNotifiersParams) WithRequestBody(requestBody *models.AlertNotifierManyUpdationParams) *UpdateManyAlertNotifiersParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update many alert notifiers params
func (o *UpdateManyAlertNotifiersParams) SetRequestBody(requestBody *models.AlertNotifierManyUpdationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateManyAlertNotifiersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
