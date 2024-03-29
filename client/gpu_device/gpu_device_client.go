// Code generated by go-swagger; DO NOT EDIT.

package gpu_device

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new gpu device API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for gpu device API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetDetailVMInfoByGpuDevices(params *GetDetailVMInfoByGpuDevicesParams, opts ...ClientOption) (*GetDetailVMInfoByGpuDevicesOK, error)

	GetGpuDevices(params *GetGpuDevicesParams, opts ...ClientOption) (*GetGpuDevicesOK, error)

	GetGpuDevicesConnection(params *GetGpuDevicesConnectionParams, opts ...ClientOption) (*GetGpuDevicesConnectionOK, error)

	SwitchGpuDeviceSriov(params *SwitchGpuDeviceSriovParams, opts ...ClientOption) (*SwitchGpuDeviceSriovOK, error)

	UpdateGpuDeviceDescription(params *UpdateGpuDeviceDescriptionParams, opts ...ClientOption) (*UpdateGpuDeviceDescriptionOK, error)

	UpdateGpuDeviceUsage(params *UpdateGpuDeviceUsageParams, opts ...ClientOption) (*UpdateGpuDeviceUsageOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetDetailVMInfoByGpuDevices get detail Vm info by gpu devices API
*/
func (a *Client) GetDetailVMInfoByGpuDevices(params *GetDetailVMInfoByGpuDevicesParams, opts ...ClientOption) (*GetDetailVMInfoByGpuDevicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDetailVMInfoByGpuDevicesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDetailVmInfoByGpuDevices",
		Method:             "POST",
		PathPattern:        "/get-detail-vm-info-by-gpu-devices",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDetailVMInfoByGpuDevicesReader{formats: a.formats},
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
	success, ok := result.(*GetDetailVMInfoByGpuDevicesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDetailVmInfoByGpuDevices: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetGpuDevices get gpu devices API
*/
func (a *Client) GetGpuDevices(params *GetGpuDevicesParams, opts ...ClientOption) (*GetGpuDevicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGpuDevicesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetGpuDevices",
		Method:             "POST",
		PathPattern:        "/get-gpu-devices",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetGpuDevicesReader{formats: a.formats},
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
	success, ok := result.(*GetGpuDevicesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetGpuDevices: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetGpuDevicesConnection get gpu devices connection API
*/
func (a *Client) GetGpuDevicesConnection(params *GetGpuDevicesConnectionParams, opts ...ClientOption) (*GetGpuDevicesConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGpuDevicesConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetGpuDevicesConnection",
		Method:             "POST",
		PathPattern:        "/get-gpu-devices-connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetGpuDevicesConnectionReader{formats: a.formats},
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
	success, ok := result.(*GetGpuDevicesConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetGpuDevicesConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SwitchGpuDeviceSriov switch gpu device sriov API
*/
func (a *Client) SwitchGpuDeviceSriov(params *SwitchGpuDeviceSriovParams, opts ...ClientOption) (*SwitchGpuDeviceSriovOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSwitchGpuDeviceSriovParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SwitchGpuDeviceSriov",
		Method:             "POST",
		PathPattern:        "/switch-gpu-device-sriov",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SwitchGpuDeviceSriovReader{formats: a.formats},
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
	success, ok := result.(*SwitchGpuDeviceSriovOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SwitchGpuDeviceSriov: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateGpuDeviceDescription update gpu device description API
*/
func (a *Client) UpdateGpuDeviceDescription(params *UpdateGpuDeviceDescriptionParams, opts ...ClientOption) (*UpdateGpuDeviceDescriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateGpuDeviceDescriptionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateGpuDeviceDescription",
		Method:             "POST",
		PathPattern:        "/update-gpu-device-description",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateGpuDeviceDescriptionReader{formats: a.formats},
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
	success, ok := result.(*UpdateGpuDeviceDescriptionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateGpuDeviceDescription: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateGpuDeviceUsage update gpu device usage API
*/
func (a *Client) UpdateGpuDeviceUsage(params *UpdateGpuDeviceUsageParams, opts ...ClientOption) (*UpdateGpuDeviceUsageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateGpuDeviceUsageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateGpuDeviceUsage",
		Method:             "POST",
		PathPattern:        "/update-gpu-device-usage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateGpuDeviceUsageReader{formats: a.formats},
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
	success, ok := result.(*UpdateGpuDeviceUsageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateGpuDeviceUsage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
