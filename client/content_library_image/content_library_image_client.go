// Code generated by go-swagger; DO NOT EDIT.

package content_library_image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new content library image API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for content library image API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateContentLibraryImage(params *CreateContentLibraryImageParams, opts ...ClientOption) (*CreateContentLibraryImageOK, error)

	DeleteContentLibraryImage(params *DeleteContentLibraryImageParams, opts ...ClientOption) (*DeleteContentLibraryImageOK, error)

	DistributeContentLibraryImageClusters(params *DistributeContentLibraryImageClustersParams, opts ...ClientOption) (*DistributeContentLibraryImageClustersOK, error)

	GetContentLibraryImages(params *GetContentLibraryImagesParams, opts ...ClientOption) (*GetContentLibraryImagesOK, error)

	GetContentLibraryImagesConnection(params *GetContentLibraryImagesConnectionParams, opts ...ClientOption) (*GetContentLibraryImagesConnectionOK, error)

	ImportContentLibraryImage(params *ImportContentLibraryImageParams, opts ...ClientOption) (*ImportContentLibraryImageOK, error)

	RemoveContentLibraryImageClusters(params *RemoveContentLibraryImageClustersParams, opts ...ClientOption) (*RemoveContentLibraryImageClustersOK, error)

	UpdateContentLibraryImage(params *UpdateContentLibraryImageParams, opts ...ClientOption) (*UpdateContentLibraryImageOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateContentLibraryImage create content library image API
*/
func (a *Client) CreateContentLibraryImage(params *CreateContentLibraryImageParams, opts ...ClientOption) (*CreateContentLibraryImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateContentLibraryImageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateContentLibraryImage",
		Method:             "POST",
		PathPattern:        "/upload-content-library-image",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateContentLibraryImageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateContentLibraryImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateContentLibraryImage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteContentLibraryImage delete content library image API
*/
func (a *Client) DeleteContentLibraryImage(params *DeleteContentLibraryImageParams, opts ...ClientOption) (*DeleteContentLibraryImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteContentLibraryImageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteContentLibraryImage",
		Method:             "POST",
		PathPattern:        "/delete-content-library-image",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteContentLibraryImageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteContentLibraryImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteContentLibraryImage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DistributeContentLibraryImageClusters distribute content library image clusters API
*/
func (a *Client) DistributeContentLibraryImageClusters(params *DistributeContentLibraryImageClustersParams, opts ...ClientOption) (*DistributeContentLibraryImageClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDistributeContentLibraryImageClustersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DistributeContentLibraryImageClusters",
		Method:             "POST",
		PathPattern:        "/distribute-content-library-image-clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DistributeContentLibraryImageClustersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DistributeContentLibraryImageClustersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DistributeContentLibraryImageClusters: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetContentLibraryImages get content library images API
*/
func (a *Client) GetContentLibraryImages(params *GetContentLibraryImagesParams, opts ...ClientOption) (*GetContentLibraryImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetContentLibraryImagesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetContentLibraryImages",
		Method:             "POST",
		PathPattern:        "/get-content-library-images",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetContentLibraryImagesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetContentLibraryImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetContentLibraryImages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetContentLibraryImagesConnection get content library images connection API
*/
func (a *Client) GetContentLibraryImagesConnection(params *GetContentLibraryImagesConnectionParams, opts ...ClientOption) (*GetContentLibraryImagesConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetContentLibraryImagesConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetContentLibraryImagesConnection",
		Method:             "POST",
		PathPattern:        "/get-content-library-images-connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetContentLibraryImagesConnectionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetContentLibraryImagesConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetContentLibraryImagesConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ImportContentLibraryImage import content library image API
*/
func (a *Client) ImportContentLibraryImage(params *ImportContentLibraryImageParams, opts ...ClientOption) (*ImportContentLibraryImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImportContentLibraryImageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ImportContentLibraryImage",
		Method:             "POST",
		PathPattern:        "/import-content-library-image",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ImportContentLibraryImageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ImportContentLibraryImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ImportContentLibraryImage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RemoveContentLibraryImageClusters remove content library image clusters API
*/
func (a *Client) RemoveContentLibraryImageClusters(params *RemoveContentLibraryImageClustersParams, opts ...ClientOption) (*RemoveContentLibraryImageClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveContentLibraryImageClustersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RemoveContentLibraryImageClusters",
		Method:             "POST",
		PathPattern:        "/remove-content-library-image-clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RemoveContentLibraryImageClustersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RemoveContentLibraryImageClustersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RemoveContentLibraryImageClusters: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateContentLibraryImage update content library image API
*/
func (a *Client) UpdateContentLibraryImage(params *UpdateContentLibraryImageParams, opts ...ClientOption) (*UpdateContentLibraryImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateContentLibraryImageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateContentLibraryImage",
		Method:             "POST",
		PathPattern:        "/update-content-library-image",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateContentLibraryImageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateContentLibraryImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateContentLibraryImage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
