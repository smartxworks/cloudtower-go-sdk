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

// NewBackupRestoreVMInPlaceParams creates a new BackupRestoreVMInPlaceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBackupRestoreVMInPlaceParams() *BackupRestoreVMInPlaceParams {
	return &BackupRestoreVMInPlaceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBackupRestoreVMInPlaceParamsWithTimeout creates a new BackupRestoreVMInPlaceParams object
// with the ability to set a timeout on a request.
func NewBackupRestoreVMInPlaceParamsWithTimeout(timeout time.Duration) *BackupRestoreVMInPlaceParams {
	return &BackupRestoreVMInPlaceParams{
		timeout: timeout,
	}
}

// NewBackupRestoreVMInPlaceParamsWithContext creates a new BackupRestoreVMInPlaceParams object
// with the ability to set a context for a request.
func NewBackupRestoreVMInPlaceParamsWithContext(ctx context.Context) *BackupRestoreVMInPlaceParams {
	return &BackupRestoreVMInPlaceParams{
		Context: ctx,
	}
}

// NewBackupRestoreVMInPlaceParamsWithHTTPClient creates a new BackupRestoreVMInPlaceParams object
// with the ability to set a custom HTTPClient for a request.
func NewBackupRestoreVMInPlaceParamsWithHTTPClient(client *http.Client) *BackupRestoreVMInPlaceParams {
	return &BackupRestoreVMInPlaceParams{
		HTTPClient: client,
	}
}

/* BackupRestoreVMInPlaceParams contains all the parameters to send to the API endpoint
   for the backup restore Vm in place operation.

   Typically these are written to a http.Request.
*/
type BackupRestoreVMInPlaceParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.BackupRestorePointRestoreInPlaceParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the backup restore Vm in place params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BackupRestoreVMInPlaceParams) WithDefaults() *BackupRestoreVMInPlaceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the backup restore Vm in place params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BackupRestoreVMInPlaceParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := BackupRestoreVMInPlaceParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the backup restore Vm in place params
func (o *BackupRestoreVMInPlaceParams) WithTimeout(timeout time.Duration) *BackupRestoreVMInPlaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the backup restore Vm in place params
func (o *BackupRestoreVMInPlaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the backup restore Vm in place params
func (o *BackupRestoreVMInPlaceParams) WithContext(ctx context.Context) *BackupRestoreVMInPlaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the backup restore Vm in place params
func (o *BackupRestoreVMInPlaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the backup restore Vm in place params
func (o *BackupRestoreVMInPlaceParams) WithHTTPClient(client *http.Client) *BackupRestoreVMInPlaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the backup restore Vm in place params
func (o *BackupRestoreVMInPlaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the backup restore Vm in place params
func (o *BackupRestoreVMInPlaceParams) WithContentLanguage(contentLanguage *string) *BackupRestoreVMInPlaceParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the backup restore Vm in place params
func (o *BackupRestoreVMInPlaceParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the backup restore Vm in place params
func (o *BackupRestoreVMInPlaceParams) WithRequestBody(requestBody *models.BackupRestorePointRestoreInPlaceParams) *BackupRestoreVMInPlaceParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the backup restore Vm in place params
func (o *BackupRestoreVMInPlaceParams) SetRequestBody(requestBody *models.BackupRestorePointRestoreInPlaceParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *BackupRestoreVMInPlaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
