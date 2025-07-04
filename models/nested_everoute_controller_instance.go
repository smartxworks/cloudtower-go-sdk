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

// NestedEverouteControllerInstance nested everoute controller instance
//
// swagger:model NestedEverouteControllerInstance
type NestedEverouteControllerInstance struct {

	// ip addr
	// Required: true
	IPAddr *string `json:"ipAddr"`

	// vlan
	// Required: true
	Vlan *string `json:"vlan"`

	MarshalOpts *NestedEverouteControllerInstanceMarshalOpts `json:"-"`
}

type NestedEverouteControllerInstanceMarshalOpts struct {
	IPAddr_Explicit_Null_When_Empty bool

	Vlan_Explicit_Null_When_Empty bool
}

func (m NestedEverouteControllerInstance) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field ipAddr
	if m.IPAddr != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ipAddr\":")
		bytes, err := swag.WriteJSON(m.IPAddr)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IPAddr_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ipAddr\":null")
		first = false
	}

	// handle nullable field vlan
	if m.Vlan != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vlan\":")
		bytes, err := swag.WriteJSON(m.Vlan)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Vlan_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vlan\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this nested everoute controller instance
func (m *NestedEverouteControllerInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPAddr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedEverouteControllerInstance) validateIPAddr(formats strfmt.Registry) error {

	if err := validate.Required("ipAddr", "body", m.IPAddr); err != nil {
		return err
	}

	return nil
}

func (m *NestedEverouteControllerInstance) validateVlan(formats strfmt.Registry) error {

	if err := validate.Required("vlan", "body", m.Vlan); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this nested everoute controller instance based on context it is used
func (m *NestedEverouteControllerInstance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NestedEverouteControllerInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedEverouteControllerInstance) UnmarshalBinary(b []byte) error {
	var res NestedEverouteControllerInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
