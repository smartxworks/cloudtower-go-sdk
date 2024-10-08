// Code generated by go-swagger; DO NOT EDIT.

package v2_everoute_license

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new v2 everoute license API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v2 everoute license API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetV2EverouteLicenses(params *GetV2EverouteLicensesParams, opts ...ClientOption) (*GetV2EverouteLicensesOK, error)

	GetV2EverouteLicensesConnection(params *GetV2EverouteLicensesConnectionParams, opts ...ClientOption) (*GetV2EverouteLicensesConnectionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetV2EverouteLicenses get v2 everoute licenses API
*/
func (a *Client) GetV2EverouteLicenses(params *GetV2EverouteLicensesParams, opts ...ClientOption) (*GetV2EverouteLicensesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV2EverouteLicensesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV2EverouteLicenses",
		Method:             "POST",
		PathPattern:        "/get-v2-everoute-licenses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetV2EverouteLicensesReader{formats: a.formats},
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
	success, ok := result.(*GetV2EverouteLicensesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV2EverouteLicenses: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV2EverouteLicensesConnection get v2 everoute licenses connection API
*/
func (a *Client) GetV2EverouteLicensesConnection(params *GetV2EverouteLicensesConnectionParams, opts ...ClientOption) (*GetV2EverouteLicensesConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV2EverouteLicensesConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV2EverouteLicensesConnection",
		Method:             "POST",
		PathPattern:        "/get-v-2-everoute-licenses-connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetV2EverouteLicensesConnectionReader{formats: a.formats},
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
	success, ok := result.(*GetV2EverouteLicensesConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV2EverouteLicensesConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
