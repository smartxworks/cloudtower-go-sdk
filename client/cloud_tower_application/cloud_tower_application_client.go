// Code generated by go-swagger; DO NOT EDIT.

package cloud_tower_application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new cloud tower application API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cloud tower application API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteCloudTowerApplicationPackage(params *DeleteCloudTowerApplicationPackageParams, opts ...ClientOption) (*DeleteCloudTowerApplicationPackageOK, error)

	DeployCloudTowerApplication(params *DeployCloudTowerApplicationParams, opts ...ClientOption) (*DeployCloudTowerApplicationOK, error)

	GetCloudTowerApplications(params *GetCloudTowerApplicationsParams, opts ...ClientOption) (*GetCloudTowerApplicationsOK, error)

	GetCloudTowerApplicationsConnection(params *GetCloudTowerApplicationsConnectionParams, opts ...ClientOption) (*GetCloudTowerApplicationsConnectionOK, error)

	UninstallCloudTowerApplication(params *UninstallCloudTowerApplicationParams, opts ...ClientOption) (*UninstallCloudTowerApplicationOK, error)

	UpgradeCloudTowerApplication(params *UpgradeCloudTowerApplicationParams, opts ...ClientOption) (*UpgradeCloudTowerApplicationOK, error)

	UploadCloudTowerApplicationPackage(params *UploadCloudTowerApplicationPackageParams, opts ...ClientOption) (*UploadCloudTowerApplicationPackageOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteCloudTowerApplicationPackage delete cloud tower application package API
*/
func (a *Client) DeleteCloudTowerApplicationPackage(params *DeleteCloudTowerApplicationPackageParams, opts ...ClientOption) (*DeleteCloudTowerApplicationPackageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCloudTowerApplicationPackageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteCloudTowerApplicationPackage",
		Method:             "POST",
		PathPattern:        "/delete-cloudtower-application-package",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteCloudTowerApplicationPackageReader{formats: a.formats},
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
	success, ok := result.(*DeleteCloudTowerApplicationPackageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteCloudTowerApplicationPackage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeployCloudTowerApplication deploy cloud tower application API
*/
func (a *Client) DeployCloudTowerApplication(params *DeployCloudTowerApplicationParams, opts ...ClientOption) (*DeployCloudTowerApplicationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeployCloudTowerApplicationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeployCloudTowerApplication",
		Method:             "POST",
		PathPattern:        "/deploy-cloudtower-application",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeployCloudTowerApplicationReader{formats: a.formats},
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
	success, ok := result.(*DeployCloudTowerApplicationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeployCloudTowerApplication: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCloudTowerApplications get cloud tower applications API
*/
func (a *Client) GetCloudTowerApplications(params *GetCloudTowerApplicationsParams, opts ...ClientOption) (*GetCloudTowerApplicationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCloudTowerApplicationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCloudTowerApplications",
		Method:             "POST",
		PathPattern:        "/get-cloudtower-applications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCloudTowerApplicationsReader{formats: a.formats},
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
	success, ok := result.(*GetCloudTowerApplicationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetCloudTowerApplications: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCloudTowerApplicationsConnection get cloud tower applications connection API
*/
func (a *Client) GetCloudTowerApplicationsConnection(params *GetCloudTowerApplicationsConnectionParams, opts ...ClientOption) (*GetCloudTowerApplicationsConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCloudTowerApplicationsConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCloudTowerApplicationsConnection",
		Method:             "POST",
		PathPattern:        "/get-cloud-tower-applications-connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCloudTowerApplicationsConnectionReader{formats: a.formats},
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
	success, ok := result.(*GetCloudTowerApplicationsConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetCloudTowerApplicationsConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UninstallCloudTowerApplication uninstall cloud tower application API
*/
func (a *Client) UninstallCloudTowerApplication(params *UninstallCloudTowerApplicationParams, opts ...ClientOption) (*UninstallCloudTowerApplicationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUninstallCloudTowerApplicationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UninstallCloudTowerApplication",
		Method:             "POST",
		PathPattern:        "/uninstall-cloudtower-application",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UninstallCloudTowerApplicationReader{formats: a.formats},
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
	success, ok := result.(*UninstallCloudTowerApplicationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UninstallCloudTowerApplication: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpgradeCloudTowerApplication upgrade cloud tower application API
*/
func (a *Client) UpgradeCloudTowerApplication(params *UpgradeCloudTowerApplicationParams, opts ...ClientOption) (*UpgradeCloudTowerApplicationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpgradeCloudTowerApplicationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpgradeCloudTowerApplication",
		Method:             "POST",
		PathPattern:        "/upgrade-cloudtower-application",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpgradeCloudTowerApplicationReader{formats: a.formats},
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
	success, ok := result.(*UpgradeCloudTowerApplicationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpgradeCloudTowerApplication: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UploadCloudTowerApplicationPackage upload cloud tower application package API
*/
func (a *Client) UploadCloudTowerApplicationPackage(params *UploadCloudTowerApplicationPackageParams, opts ...ClientOption) (*UploadCloudTowerApplicationPackageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadCloudTowerApplicationPackageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UploadCloudTowerApplicationPackage",
		Method:             "POST",
		PathPattern:        "/upload-cloudtower-application-package",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UploadCloudTowerApplicationPackageReader{formats: a.formats},
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
	success, ok := result.(*UploadCloudTowerApplicationPackageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UploadCloudTowerApplicationPackage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}