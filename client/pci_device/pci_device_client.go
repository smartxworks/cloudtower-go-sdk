// Code generated by go-swagger; DO NOT EDIT.

package pci_device

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new pci device API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for pci device API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetPciDevices(params *GetPciDevicesParams, opts ...ClientOption) (*GetPciDevicesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetPciDevices get pci devices API
*/
func (a *Client) GetPciDevices(params *GetPciDevicesParams, opts ...ClientOption) (*GetPciDevicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPciDevicesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetPciDevices",
		Method:             "POST",
		PathPattern:        "/get-pci-devices",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPciDevicesReader{formats: a.formats},
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
	success, ok := result.(*GetPciDevicesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPciDevices: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
