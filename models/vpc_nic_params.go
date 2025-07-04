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

// VpcNicParams vpc nic params
//
// swagger:model VpcNicParams
type VpcNicParams struct {

	// floating ip id
	FloatingIPID *string `json:"floating_ip_id,omitempty"`

	// ip addresses
	IPAddresses []string `json:"ip_addresses,omitempty"`

	// vpc id
	// Required: true
	VpcID *string `json:"vpc_id"`

	// vpc subnet id
	// Required: true
	VpcSubnetID *string `json:"vpc_subnet_id"`

	MarshalOpts *VpcNicParamsMarshalOpts `json:"-"`
}

type VpcNicParamsMarshalOpts struct {
	FloatingIPID_Explicit_Null_When_Empty bool

	IPAddresses_Explicit_Null_When_Empty bool

	VpcID_Explicit_Null_When_Empty bool

	VpcSubnetID_Explicit_Null_When_Empty bool
}

func (m VpcNicParams) MarshalJSON() ([]byte, error) {
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

	// handle nullable field vpc_id
	if m.VpcID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vpc_id\":")
		bytes, err := swag.WriteJSON(m.VpcID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VpcID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vpc_id\":null")
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

// Validate validates this vpc nic params
func (m *VpcNicParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVpcID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpcSubnetID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VpcNicParams) validateVpcID(formats strfmt.Registry) error {

	if err := validate.Required("vpc_id", "body", m.VpcID); err != nil {
		return err
	}

	return nil
}

func (m *VpcNicParams) validateVpcSubnetID(formats strfmt.Registry) error {

	if err := validate.Required("vpc_subnet_id", "body", m.VpcSubnetID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this vpc nic params based on context it is used
func (m *VpcNicParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VpcNicParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VpcNicParams) UnmarshalBinary(b []byte) error {
	var res VpcNicParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
