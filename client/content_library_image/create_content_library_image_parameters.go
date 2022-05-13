// Code generated by go-swagger; DO NOT EDIT.

package content_library_image

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
)

// NewCreateContentLibraryImageParams creates a new CreateContentLibraryImageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateContentLibraryImageParams() *CreateContentLibraryImageParams {
	return &CreateContentLibraryImageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateContentLibraryImageParamsWithTimeout creates a new CreateContentLibraryImageParams object
// with the ability to set a timeout on a request.
func NewCreateContentLibraryImageParamsWithTimeout(timeout time.Duration) *CreateContentLibraryImageParams {
	return &CreateContentLibraryImageParams{
		timeout: timeout,
	}
}

// NewCreateContentLibraryImageParamsWithContext creates a new CreateContentLibraryImageParams object
// with the ability to set a context for a request.
func NewCreateContentLibraryImageParamsWithContext(ctx context.Context) *CreateContentLibraryImageParams {
	return &CreateContentLibraryImageParams{
		Context: ctx,
	}
}

// NewCreateContentLibraryImageParamsWithHTTPClient creates a new CreateContentLibraryImageParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateContentLibraryImageParamsWithHTTPClient(client *http.Client) *CreateContentLibraryImageParams {
	return &CreateContentLibraryImageParams{
		HTTPClient: client,
	}
}

/* CreateContentLibraryImageParams contains all the parameters to send to the API endpoint
   for the create content library image operation.

   Typically these are written to a http.Request.
*/
type CreateContentLibraryImageParams struct {

	// Clusters.
	Clusters string

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// Description.
	Description string

	// File.
	File runtime.NamedReadCloser

	// Name.
	Name string

	// Size.
	Size string

	// UploadTaskID.
	UploadTaskID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create content library image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateContentLibraryImageParams) WithDefaults() *CreateContentLibraryImageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create content library image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateContentLibraryImageParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := CreateContentLibraryImageParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create content library image params
func (o *CreateContentLibraryImageParams) WithTimeout(timeout time.Duration) *CreateContentLibraryImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create content library image params
func (o *CreateContentLibraryImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create content library image params
func (o *CreateContentLibraryImageParams) WithContext(ctx context.Context) *CreateContentLibraryImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create content library image params
func (o *CreateContentLibraryImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create content library image params
func (o *CreateContentLibraryImageParams) WithHTTPClient(client *http.Client) *CreateContentLibraryImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create content library image params
func (o *CreateContentLibraryImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusters adds the clusters to the create content library image params
func (o *CreateContentLibraryImageParams) WithClusters(clusters string) *CreateContentLibraryImageParams {
	o.SetClusters(clusters)
	return o
}

// SetClusters adds the clusters to the create content library image params
func (o *CreateContentLibraryImageParams) SetClusters(clusters string) {
	o.Clusters = clusters
}

// WithContentLanguage adds the contentLanguage to the create content library image params
func (o *CreateContentLibraryImageParams) WithContentLanguage(contentLanguage *string) *CreateContentLibraryImageParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the create content library image params
func (o *CreateContentLibraryImageParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithDescription adds the description to the create content library image params
func (o *CreateContentLibraryImageParams) WithDescription(description string) *CreateContentLibraryImageParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the create content library image params
func (o *CreateContentLibraryImageParams) SetDescription(description string) {
	o.Description = description
}

// WithFile adds the file to the create content library image params
func (o *CreateContentLibraryImageParams) WithFile(file runtime.NamedReadCloser) *CreateContentLibraryImageParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the create content library image params
func (o *CreateContentLibraryImageParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WithName adds the name to the create content library image params
func (o *CreateContentLibraryImageParams) WithName(name string) *CreateContentLibraryImageParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the create content library image params
func (o *CreateContentLibraryImageParams) SetName(name string) {
	o.Name = name
}

// WithSize adds the size to the create content library image params
func (o *CreateContentLibraryImageParams) WithSize(size string) *CreateContentLibraryImageParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the create content library image params
func (o *CreateContentLibraryImageParams) SetSize(size string) {
	o.Size = size
}

// WithUploadTaskID adds the uploadTaskID to the create content library image params
func (o *CreateContentLibraryImageParams) WithUploadTaskID(uploadTaskID *string) *CreateContentLibraryImageParams {
	o.SetUploadTaskID(uploadTaskID)
	return o
}

// SetUploadTaskID adds the uploadTaskId to the create content library image params
func (o *CreateContentLibraryImageParams) SetUploadTaskID(uploadTaskID *string) {
	o.UploadTaskID = uploadTaskID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateContentLibraryImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param clusters
	frClusters := o.Clusters
	fClusters := frClusters
	if fClusters != "" {
		if err := r.SetFormParam("clusters", fClusters); err != nil {
			return err
		}
	}

	if o.ContentLanguage != nil {

		// header param content-language
		if err := r.SetHeaderParam("content-language", *o.ContentLanguage); err != nil {
			return err
		}
	}

	// form param description
	frDescription := o.Description
	fDescription := frDescription
	if fDescription != "" {
		if err := r.SetFormParam("description", fDescription); err != nil {
			return err
		}
	}
	// form file param file
	if err := r.SetFileParam("file", o.File); err != nil {
		return err
	}

	// form param name
	frName := o.Name
	fName := frName
	if fName != "" {
		if err := r.SetFormParam("name", fName); err != nil {
			return err
		}
	}

	// form param size
	frSize := o.Size
	fSize := frSize
	if fSize != "" {
		if err := r.SetFormParam("size", fSize); err != nil {
			return err
		}
	}

	if o.UploadTaskID != nil {

		// form param upload_task_id
		var frUploadTaskID string
		if o.UploadTaskID != nil {
			frUploadTaskID = *o.UploadTaskID
		}
		fUploadTaskID := frUploadTaskID
		if fUploadTaskID != "" {
			if err := r.SetFormParam("upload_task_id", fUploadTaskID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}