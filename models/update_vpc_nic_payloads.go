// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateVpcNicPayloads update vpc nic payloads
//
// swagger:model UpdateVpcNicPayloads
type UpdateVpcNicPayloads struct {

	// floating ip id
	FloatingIPID *string `json:"floating_ip_id,omitempty"`

	// ip addresses
	IPAddresses []string `json:"ip_addresses,omitempty"`

	// vpc subnet id
	VpcSubnetID *string `json:"vpc_subnet_id,omitempty"`

	MarshalOpts *UpdateVpcNicPayloadsMarshalOpts `json:"-"`
}

type UpdateVpcNicPayloadsMarshalOpts struct {
	FloatingIPID_Explicit_Null_When_Empty bool

	IPAddresses_Explicit_Null_When_Empty bool

	VpcSubnetID_Explicit_Null_When_Empty bool
}

func (m UpdateVpcNicPayloads) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field floating_ip_id
	if m.FloatingIPID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"floating_ip_id\":")
		bytes, err := swag.WriteJSON(m.FloatingIPID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.FloatingIPID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"floating_ip_id\":null")
		first = false
	}

	// handle non nullable field ip_addresses with omitempty
	if !swag.IsZero(m.IPAddresses) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ip_addresses\":")
		bytes, err := swag.WriteJSON(m.IPAddresses)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field vpc_subnet_id
	if m.VpcSubnetID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vpc_subnet_id\":")
		bytes, err := swag.WriteJSON(m.VpcSubnetID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VpcSubnetID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vpc_subnet_id\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this update vpc nic payloads
func (m *UpdateVpcNicPayloads) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update vpc nic payloads based on context it is used
func (m *UpdateVpcNicPayloads) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateVpcNicPayloads) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateVpcNicPayloads) UnmarshalBinary(b []byte) error {
	var res UpdateVpcNicPayloads
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
