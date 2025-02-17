// Code generated by go-swagger; DO NOT EDIT.

package backup_plan

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

// NewSuspendBackupPlanParams creates a new SuspendBackupPlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSuspendBackupPlanParams() *SuspendBackupPlanParams {
	return &SuspendBackupPlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSuspendBackupPlanParamsWithTimeout creates a new SuspendBackupPlanParams object
// with the ability to set a timeout on a request.
func NewSuspendBackupPlanParamsWithTimeout(timeout time.Duration) *SuspendBackupPlanParams {
	return &SuspendBackupPlanParams{
		timeout: timeout,
	}
}

// NewSuspendBackupPlanParamsWithContext creates a new SuspendBackupPlanParams object
// with the ability to set a context for a request.
func NewSuspendBackupPlanParamsWithContext(ctx context.Context) *SuspendBackupPlanParams {
	return &SuspendBackupPlanParams{
		Context: ctx,
	}
}

// NewSuspendBackupPlanParamsWithHTTPClient creates a new SuspendBackupPlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewSuspendBackupPlanParamsWithHTTPClient(client *http.Client) *SuspendBackupPlanParams {
	return &SuspendBackupPlanParams{
		HTTPClient: client,
	}
}

/* SuspendBackupPlanParams contains all the parameters to send to the API endpoint
   for the suspend backup plan operation.

   Typically these are written to a http.Request.
*/
type SuspendBackupPlanParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.BackupPlanSuspendParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the suspend backup plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SuspendBackupPlanParams) WithDefaults() *SuspendBackupPlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the suspend backup plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SuspendBackupPlanParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := SuspendBackupPlanParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the suspend backup plan params
func (o *SuspendBackupPlanParams) WithTimeout(timeout time.Duration) *SuspendBackupPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the suspend backup plan params
func (o *SuspendBackupPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the suspend backup plan params
func (o *SuspendBackupPlanParams) WithContext(ctx context.Context) *SuspendBackupPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the suspend backup plan params
func (o *SuspendBackupPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the suspend backup plan params
func (o *SuspendBackupPlanParams) WithHTTPClient(client *http.Client) *SuspendBackupPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the suspend backup plan params
func (o *SuspendBackupPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the suspend backup plan params
func (o *SuspendBackupPlanParams) WithContentLanguage(contentLanguage *string) *SuspendBackupPlanParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the suspend backup plan params
func (o *SuspendBackupPlanParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the suspend backup plan params
func (o *SuspendBackupPlanParams) WithRequestBody(requestBody *models.BackupPlanSuspendParams) *SuspendBackupPlanParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the suspend backup plan params
func (o *SuspendBackupPlanParams) SetRequestBody(requestBody *models.BackupPlanSuspendParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *SuspendBackupPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
