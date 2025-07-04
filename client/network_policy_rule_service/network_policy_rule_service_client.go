// Code generated by go-swagger; DO NOT EDIT.

package network_policy_rule_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new network policy rule service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for network policy rule service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateNetworkPolicyRuleService(params *CreateNetworkPolicyRuleServiceParams, opts ...ClientOption) (*CreateNetworkPolicyRuleServiceOK, error)

	DeleteNetworkPolicyRuleService(params *DeleteNetworkPolicyRuleServiceParams, opts ...ClientOption) (*DeleteNetworkPolicyRuleServiceOK, error)

	GetNetworkPolicyRuleServices(params *GetNetworkPolicyRuleServicesParams, opts ...ClientOption) (*GetNetworkPolicyRuleServicesOK, error)

	GetNetworkPolicyRuleServicesConnection(params *GetNetworkPolicyRuleServicesConnectionParams, opts ...ClientOption) (*GetNetworkPolicyRuleServicesConnectionOK, error)

	UpdateNetworkPolicyRuleService(params *UpdateNetworkPolicyRuleServiceParams, opts ...ClientOption) (*UpdateNetworkPolicyRuleServiceOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateNetworkPolicyRuleService create network policy rule service API
*/
func (a *Client) CreateNetworkPolicyRuleService(params *CreateNetworkPolicyRuleServiceParams, opts ...ClientOption) (*CreateNetworkPolicyRuleServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateNetworkPolicyRuleServiceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateNetworkPolicyRuleService",
		Method:             "POST",
		PathPattern:        "/create-network-policy-rule-service",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateNetworkPolicyRuleServiceReader{formats: a.formats},
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
	success, ok := result.(*CreateNetworkPolicyRuleServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateNetworkPolicyRuleService: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteNetworkPolicyRuleService delete network policy rule service API
*/
func (a *Client) DeleteNetworkPolicyRuleService(params *DeleteNetworkPolicyRuleServiceParams, opts ...ClientOption) (*DeleteNetworkPolicyRuleServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNetworkPolicyRuleServiceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteNetworkPolicyRuleService",
		Method:             "POST",
		PathPattern:        "/delete-network-policy-rule-service",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteNetworkPolicyRuleServiceReader{formats: a.formats},
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
	success, ok := result.(*DeleteNetworkPolicyRuleServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteNetworkPolicyRuleService: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNetworkPolicyRuleServices get network policy rule services API
*/
func (a *Client) GetNetworkPolicyRuleServices(params *GetNetworkPolicyRuleServicesParams, opts ...ClientOption) (*GetNetworkPolicyRuleServicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkPolicyRuleServicesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetNetworkPolicyRuleServices",
		Method:             "POST",
		PathPattern:        "/get-network-policy-rule-services",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetNetworkPolicyRuleServicesReader{formats: a.formats},
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
	success, ok := result.(*GetNetworkPolicyRuleServicesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetNetworkPolicyRuleServices: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNetworkPolicyRuleServicesConnection get network policy rule services connection API
*/
func (a *Client) GetNetworkPolicyRuleServicesConnection(params *GetNetworkPolicyRuleServicesConnectionParams, opts ...ClientOption) (*GetNetworkPolicyRuleServicesConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkPolicyRuleServicesConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetNetworkPolicyRuleServicesConnection",
		Method:             "POST",
		PathPattern:        "/get-network-policy-rule-services-connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetNetworkPolicyRuleServicesConnectionReader{formats: a.formats},
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
	success, ok := result.(*GetNetworkPolicyRuleServicesConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetNetworkPolicyRuleServicesConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateNetworkPolicyRuleService update network policy rule service API
*/
func (a *Client) UpdateNetworkPolicyRuleService(params *UpdateNetworkPolicyRuleServiceParams, opts ...ClientOption) (*UpdateNetworkPolicyRuleServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateNetworkPolicyRuleServiceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateNetworkPolicyRuleService",
		Method:             "POST",
		PathPattern:        "/update-network-policy-rule-service",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateNetworkPolicyRuleServiceReader{formats: a.formats},
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
	success, ok := result.(*UpdateNetworkPolicyRuleServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateNetworkPolicyRuleService: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
