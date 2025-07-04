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

// CloudInitNetWorkRoute cloud init net work route
//
// swagger:model CloudInitNetWorkRoute
type CloudInitNetWorkRoute struct {

	// gateway
	// Required: true
	Gateway *string `json:"gateway"`

	// netmask
	// Required: true
	Netmask *string `json:"netmask"`

	// network
	// Required: true
	Network *string `json:"network"`

	MarshalOpts *CloudInitNetWorkRouteMarshalOpts `json:"-"`
}

type CloudInitNetWorkRouteMarshalOpts struct {
	Gateway_Explicit_Null_When_Empty bool

	Netmask_Explicit_Null_When_Empty bool

	Network_Explicit_Null_When_Empty bool
}

func (m CloudInitNetWorkRoute) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field gateway
	if m.Gateway != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"gateway\":")
		bytes, err := swag.WriteJSON(m.Gateway)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Gateway_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"gateway\":null")
		first = false
	}

	// handle nullable field netmask
	if m.Netmask != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"netmask\":")
		bytes, err := swag.WriteJSON(m.Netmask)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Netmask_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"netmask\":null")
		first = false
	}

	// handle nullable field network
	if m.Network != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"network\":")
		bytes, err := swag.WriteJSON(m.Network)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Network_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"network\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this cloud init net work route
func (m *CloudInitNetWorkRoute) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGateway(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetmask(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudInitNetWorkRoute) validateGateway(formats strfmt.Registry) error {

	if err := validate.Required("gateway", "body", m.Gateway); err != nil {
		return err
	}

	return nil
}

func (m *CloudInitNetWorkRoute) validateNetmask(formats strfmt.Registry) error {

	if err := validate.Required("netmask", "body", m.Netmask); err != nil {
		return err
	}

	return nil
}

func (m *CloudInitNetWorkRoute) validateNetwork(formats strfmt.Registry) error {

	if err := validate.Required("network", "body", m.Network); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cloud init net work route based on context it is used
func (m *CloudInitNetWorkRoute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CloudInitNetWorkRoute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudInitNetWorkRoute) UnmarshalBinary(b []byte) error {
	var res CloudInitNetWorkRoute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
