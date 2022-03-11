// Code generated by go-swagger; DO NOT EDIT.

package zone

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

// NewGetZonesParams creates a new GetZonesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetZonesParams() *GetZonesParams {
	return &GetZonesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetZonesParamsWithTimeout creates a new GetZonesParams object
// with the ability to set a timeout on a request.
func NewGetZonesParamsWithTimeout(timeout time.Duration) *GetZonesParams {
	return &GetZonesParams{
		timeout: timeout,
	}
}

// NewGetZonesParamsWithContext creates a new GetZonesParams object
// with the ability to set a context for a request.
func NewGetZonesParamsWithContext(ctx context.Context) *GetZonesParams {
	return &GetZonesParams{
		Context: ctx,
	}
}

// NewGetZonesParamsWithHTTPClient creates a new GetZonesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetZonesParamsWithHTTPClient(client *http.Client) *GetZonesParams {
	return &GetZonesParams{
		HTTPClient: client,
	}
}

/* GetZonesParams contains all the parameters to send to the API endpoint
   for the get zones operation.

   Typically these are written to a http.Request.
*/
type GetZonesParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetZonesRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get zones params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetZonesParams) WithDefaults() *GetZonesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get zones params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetZonesParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetZonesParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get zones params
func (o *GetZonesParams) WithTimeout(timeout time.Duration) *GetZonesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get zones params
func (o *GetZonesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get zones params
func (o *GetZonesParams) WithContext(ctx context.Context) *GetZonesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get zones params
func (o *GetZonesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get zones params
func (o *GetZonesParams) WithHTTPClient(client *http.Client) *GetZonesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get zones params
func (o *GetZonesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get zones params
func (o *GetZonesParams) WithContentLanguage(contentLanguage *string) *GetZonesParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get zones params
func (o *GetZonesParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get zones params
func (o *GetZonesParams) WithRequestBody(requestBody *models.GetZonesRequestBody) *GetZonesParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get zones params
func (o *GetZonesParams) SetRequestBody(requestBody *models.GetZonesRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetZonesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
