// Code generated by go-swagger; DO NOT EDIT.

package vm_volume_snapshot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new vm volume snapshot API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for vm volume snapshot API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateVMVolumeSnapshot(params *CreateVMVolumeSnapshotParams, opts ...ClientOption) (*CreateVMVolumeSnapshotOK, error)

	DeleteVMVolumeSnapshot(params *DeleteVMVolumeSnapshotParams, opts ...ClientOption) (*DeleteVMVolumeSnapshotOK, error)

	GetVMVolumeSnapshots(params *GetVMVolumeSnapshotsParams, opts ...ClientOption) (*GetVMVolumeSnapshotsOK, error)

	GetVMVolumeSnapshotsConnection(params *GetVMVolumeSnapshotsConnectionParams, opts ...ClientOption) (*GetVMVolumeSnapshotsConnectionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateVMVolumeSnapshot create Vm volume snapshot API
*/
func (a *Client) CreateVMVolumeSnapshot(params *CreateVMVolumeSnapshotParams, opts ...ClientOption) (*CreateVMVolumeSnapshotOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateVMVolumeSnapshotParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateVmVolumeSnapshot",
		Method:             "POST",
		PathPattern:        "/create-vm-volume-snapshot",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateVMVolumeSnapshotReader{formats: a.formats},
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
	success, ok := result.(*CreateVMVolumeSnapshotOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateVmVolumeSnapshot: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteVMVolumeSnapshot delete Vm volume snapshot API
*/
func (a *Client) DeleteVMVolumeSnapshot(params *DeleteVMVolumeSnapshotParams, opts ...ClientOption) (*DeleteVMVolumeSnapshotOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVMVolumeSnapshotParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteVmVolumeSnapshot",
		Method:             "POST",
		PathPattern:        "/delete-vm-volume-snapshot",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteVMVolumeSnapshotReader{formats: a.formats},
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
	success, ok := result.(*DeleteVMVolumeSnapshotOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteVmVolumeSnapshot: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVMVolumeSnapshots get Vm volume snapshots API
*/
func (a *Client) GetVMVolumeSnapshots(params *GetVMVolumeSnapshotsParams, opts ...ClientOption) (*GetVMVolumeSnapshotsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVMVolumeSnapshotsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetVmVolumeSnapshots",
		Method:             "POST",
		PathPattern:        "/get-vm-volume-snapshots",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVMVolumeSnapshotsReader{formats: a.formats},
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
	success, ok := result.(*GetVMVolumeSnapshotsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetVmVolumeSnapshots: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVMVolumeSnapshotsConnection get Vm volume snapshots connection API
*/
func (a *Client) GetVMVolumeSnapshotsConnection(params *GetVMVolumeSnapshotsConnectionParams, opts ...ClientOption) (*GetVMVolumeSnapshotsConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVMVolumeSnapshotsConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetVmVolumeSnapshotsConnection",
		Method:             "POST",
		PathPattern:        "/get-vm-volume-snapshots-connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVMVolumeSnapshotsConnectionReader{formats: a.formats},
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
	success, ok := result.(*GetVMVolumeSnapshotsConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetVmVolumeSnapshotsConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
