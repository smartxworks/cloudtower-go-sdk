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

// NewDeleteBackupRestorePointParams creates a new DeleteBackupRestorePointParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteBackupRestorePointParams() *DeleteBackupRestorePointParams {
	return &DeleteBackupRestorePointParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBackupRestorePointParamsWithTimeout creates a new DeleteBackupRestorePointParams object
// with the ability to set a timeout on a request.
func NewDeleteBackupRestorePointParamsWithTimeout(timeout time.Duration) *DeleteBackupRestorePointParams {
	return &DeleteBackupRestorePointParams{
		timeout: timeout,
	}
}

// NewDeleteBackupRestorePointParamsWithContext creates a new DeleteBackupRestorePointParams object
// with the ability to set a context for a request.
func NewDeleteBackupRestorePointParamsWithContext(ctx context.Context) *DeleteBackupRestorePointParams {
	return &DeleteBackupRestorePointParams{
		Context: ctx,
	}
}

// NewDeleteBackupRestorePointParamsWithHTTPClient creates a new DeleteBackupRestorePointParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteBackupRestorePointParamsWithHTTPClient(client *http.Client) *DeleteBackupRestorePointParams {
	return &DeleteBackupRestorePointParams{
		HTTPClient: client,
	}
}

/* DeleteBackupRestorePointParams contains all the parameters to send to the API endpoint
   for the delete backup restore point operation.

   Typically these are written to a http.Request.
*/
type DeleteBackupRestorePointParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.BackupRestorePointDeletionParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete backup restore point params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBackupRestorePointParams) WithDefaults() *DeleteBackupRestorePointParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete backup restore point params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBackupRestorePointParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := DeleteBackupRestorePointParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete backup restore point params
func (o *DeleteBackupRestorePointParams) WithTimeout(timeout time.Duration) *DeleteBackupRestorePointParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete backup restore point params
func (o *DeleteBackupRestorePointParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete backup restore point params
func (o *DeleteBackupRestorePointParams) WithContext(ctx context.Context) *DeleteBackupRestorePointParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete backup restore point params
func (o *DeleteBackupRestorePointParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete backup restore point params
func (o *DeleteBackupRestorePointParams) WithHTTPClient(client *http.Client) *DeleteBackupRestorePointParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete backup restore point params
func (o *DeleteBackupRestorePointParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the delete backup restore point params
func (o *DeleteBackupRestorePointParams) WithContentLanguage(contentLanguage *string) *DeleteBackupRestorePointParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the delete backup restore point params
func (o *DeleteBackupRestorePointParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the delete backup restore point params
func (o *DeleteBackupRestorePointParams) WithRequestBody(requestBody *models.BackupRestorePointDeletionParams) *DeleteBackupRestorePointParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the delete backup restore point params
func (o *DeleteBackupRestorePointParams) SetRequestBody(requestBody *models.BackupRestorePointDeletionParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBackupRestorePointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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