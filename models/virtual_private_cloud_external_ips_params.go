// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VirtualPrivateCloudExternalIpsParams virtual private cloud external ips params
//
// swagger:model VirtualPrivateCloudExternalIpsParams
type VirtualPrivateCloudExternalIpsParams struct {

	// edge gateway id
	// Required: true
	EdgeGatewayID *string `json:"edge_gateway_id"`

	// external ip
	// Required: true
	ExternalIP *string `json:"external_ip"`

	MarshalOpts *VirtualPrivateCloudExternalIpsParamsMarshalOpts `json:"-"`
}

type VirtualPrivateCloudExternalIpsParamsMarshalOpts struct {
	EdgeGatewayID_Explicit_Null_When_Empty bool

	ExternalIP_Explicit_Null_When_Empty bool
}

func (m VirtualPrivateCloudExternalIpsParams) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field edge_gateway_id
	if m.EdgeGatewayID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"edge_gateway_id\":")
		bytes, err := swag.WriteJSON(m.EdgeGatewayID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.EdgeGatewayID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"edge_gateway_id\":null")
		first = false
	}

	// handle nullable field external_ip
	if m.ExternalIP != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"external_ip\":")
		bytes, err := swag.WriteJSON(m.ExternalIP)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ExternalIP_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"external_ip\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this virtual private cloud external ips params
func (m *VirtualPrivateCloudExternalIpsParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEdgeGatewayID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalIP(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualPrivateCloudExternalIpsParams) validateEdgeGatewayID(formats strfmt.Registry) error {

	if err := validate.Required("edge_gateway_id", "body", m.EdgeGatewayID); err != nil {
		return err
	}

	return nil
}

func (m *VirtualPrivateCloudExternalIpsParams) validateExternalIP(formats strfmt.Registry) error {

	if err := validate.Required("external_ip", "body", m.ExternalIP); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this virtual private cloud external ips params based on context it is used
func (m *VirtualPrivateCloudExternalIpsParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VirtualPrivateCloudExternalIpsParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualPrivateCloudExternalIpsParams) UnmarshalBinary(b []byte) error {
	var res VirtualPrivateCloudExternalIpsParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
