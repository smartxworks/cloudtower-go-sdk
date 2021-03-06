// Code generated by go-swagger; DO NOT EDIT.

package node_topo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new node topo API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for node topo API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetNodeTopoes(params *GetNodeTopoesParams, opts ...ClientOption) (*GetNodeTopoesOK, error)

	GetNodeTopoesConnection(params *GetNodeTopoesConnectionParams, opts ...ClientOption) (*GetNodeTopoesConnectionOK, error)

	UpdateNodeTopo(params *UpdateNodeTopoParams, opts ...ClientOption) (*UpdateNodeTopoOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetNodeTopoes get node topoes API
*/
func (a *Client) GetNodeTopoes(params *GetNodeTopoesParams, opts ...ClientOption) (*GetNodeTopoesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodeTopoesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetNodeTopoes",
		Method:             "POST",
		PathPattern:        "/get-node-topoes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetNodeTopoesReader{formats: a.formats},
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
	success, ok := result.(*GetNodeTopoesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetNodeTopoes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNodeTopoesConnection get node topoes connection API
*/
func (a *Client) GetNodeTopoesConnection(params *GetNodeTopoesConnectionParams, opts ...ClientOption) (*GetNodeTopoesConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodeTopoesConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetNodeTopoesConnection",
		Method:             "POST",
		PathPattern:        "/get-node-topoes-connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetNodeTopoesConnectionReader{formats: a.formats},
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
	success, ok := result.(*GetNodeTopoesConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetNodeTopoesConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateNodeTopo update node topo API
*/
func (a *Client) UpdateNodeTopo(params *UpdateNodeTopoParams, opts ...ClientOption) (*UpdateNodeTopoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateNodeTopoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateNodeTopo",
		Method:             "POST",
		PathPattern:        "/move-node-topo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateNodeTopoReader{formats: a.formats},
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
	success, ok := result.(*UpdateNodeTopoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateNodeTopo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
