// Code generated by go-swagger; DO NOT EDIT.

package virtual_private_cloud_security_policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new virtual private cloud security policy API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for virtual private cloud security policy API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateVirtualPrivateCloudSecurityPolicy(params *CreateVirtualPrivateCloudSecurityPolicyParams, opts ...ClientOption) (*CreateVirtualPrivateCloudSecurityPolicyOK, error)

	DeleteVirtualPrivateCloudSecurityPolicy(params *DeleteVirtualPrivateCloudSecurityPolicyParams, opts ...ClientOption) (*DeleteVirtualPrivateCloudSecurityPolicyOK, error)

	GetVirtualPrivateCloudSecurityPolicies(params *GetVirtualPrivateCloudSecurityPoliciesParams, opts ...ClientOption) (*GetVirtualPrivateCloudSecurityPoliciesOK, error)

	GetVirtualPrivateCloudSecurityPoliciesConnection(params *GetVirtualPrivateCloudSecurityPoliciesConnectionParams, opts ...ClientOption) (*GetVirtualPrivateCloudSecurityPoliciesConnectionOK, error)

	UpdateVirtualPrivateCloudSecurityPolicy(params *UpdateVirtualPrivateCloudSecurityPolicyParams, opts ...ClientOption) (*UpdateVirtualPrivateCloudSecurityPolicyOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateVirtualPrivateCloudSecurityPolicy create virtual private cloud security policy API
*/
func (a *Client) CreateVirtualPrivateCloudSecurityPolicy(params *CreateVirtualPrivateCloudSecurityPolicyParams, opts ...ClientOption) (*CreateVirtualPrivateCloudSecurityPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateVirtualPrivateCloudSecurityPolicyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateVirtualPrivateCloudSecurityPolicy",
		Method:             "POST",
		PathPattern:        "/create-virtual-private-cloud-security-policy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateVirtualPrivateCloudSecurityPolicyReader{formats: a.formats},
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
	success, ok := result.(*CreateVirtualPrivateCloudSecurityPolicyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateVirtualPrivateCloudSecurityPolicy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteVirtualPrivateCloudSecurityPolicy delete virtual private cloud security policy API
*/
func (a *Client) DeleteVirtualPrivateCloudSecurityPolicy(params *DeleteVirtualPrivateCloudSecurityPolicyParams, opts ...ClientOption) (*DeleteVirtualPrivateCloudSecurityPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVirtualPrivateCloudSecurityPolicyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteVirtualPrivateCloudSecurityPolicy",
		Method:             "POST",
		PathPattern:        "/delete-virtual-private-cloud-security-policy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteVirtualPrivateCloudSecurityPolicyReader{formats: a.formats},
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
	success, ok := result.(*DeleteVirtualPrivateCloudSecurityPolicyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteVirtualPrivateCloudSecurityPolicy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVirtualPrivateCloudSecurityPolicies get virtual private cloud security policies API
*/
func (a *Client) GetVirtualPrivateCloudSecurityPolicies(params *GetVirtualPrivateCloudSecurityPoliciesParams, opts ...ClientOption) (*GetVirtualPrivateCloudSecurityPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVirtualPrivateCloudSecurityPoliciesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetVirtualPrivateCloudSecurityPolicies",
		Method:             "POST",
		PathPattern:        "/get-virtual-private-cloud-security-policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVirtualPrivateCloudSecurityPoliciesReader{formats: a.formats},
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
	success, ok := result.(*GetVirtualPrivateCloudSecurityPoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetVirtualPrivateCloudSecurityPolicies: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVirtualPrivateCloudSecurityPoliciesConnection get virtual private cloud security policies connection API
*/
func (a *Client) GetVirtualPrivateCloudSecurityPoliciesConnection(params *GetVirtualPrivateCloudSecurityPoliciesConnectionParams, opts ...ClientOption) (*GetVirtualPrivateCloudSecurityPoliciesConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVirtualPrivateCloudSecurityPoliciesConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetVirtualPrivateCloudSecurityPoliciesConnection",
		Method:             "POST",
		PathPattern:        "/get-virtual-private-cloud-security-policies-connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVirtualPrivateCloudSecurityPoliciesConnectionReader{formats: a.formats},
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
	success, ok := result.(*GetVirtualPrivateCloudSecurityPoliciesConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetVirtualPrivateCloudSecurityPoliciesConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateVirtualPrivateCloudSecurityPolicy update virtual private cloud security policy API
*/
func (a *Client) UpdateVirtualPrivateCloudSecurityPolicy(params *UpdateVirtualPrivateCloudSecurityPolicyParams, opts ...ClientOption) (*UpdateVirtualPrivateCloudSecurityPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateVirtualPrivateCloudSecurityPolicyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateVirtualPrivateCloudSecurityPolicy",
		Method:             "POST",
		PathPattern:        "/update-virtual-private-cloud-security-policy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateVirtualPrivateCloudSecurityPolicyReader{formats: a.formats},
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
	success, ok := result.(*UpdateVirtualPrivateCloudSecurityPolicyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateVirtualPrivateCloudSecurityPolicy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}