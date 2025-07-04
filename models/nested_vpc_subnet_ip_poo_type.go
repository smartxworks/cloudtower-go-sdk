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

// NestedVpcSubnetIPPooType nested vpc subnet Ip poo type
//
// swagger:model NestedVpcSubnetIpPooType
type NestedVpcSubnetIPPooType struct {

	// end
	// Required: true
	End *string `json:"end"`

	// start
	// Required: true
	Start *string `json:"start"`

	MarshalOpts *NestedVpcSubnetIPPooTypeMarshalOpts `json:"-"`
}

type NestedVpcSubnetIPPooTypeMarshalOpts struct {
	End_Explicit_Null_When_Empty bool

	Start_Explicit_Null_When_Empty bool
}

func (m NestedVpcSubnetIPPooType) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field end
	if m.End != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"end\":")
		bytes, err := swag.WriteJSON(m.End)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.End_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"end\":null")
		first = false
	}

	// handle nullable field start
	if m.Start != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"start\":")
		bytes, err := swag.WriteJSON(m.Start)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Start_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"start\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this nested vpc subnet Ip poo type
func (m *NestedVpcSubnetIPPooType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedVpcSubnetIPPooType) validateEnd(formats strfmt.Registry) error {

	if err := validate.Required("end", "body", m.End); err != nil {
		return err
	}

	return nil
}

func (m *NestedVpcSubnetIPPooType) validateStart(formats strfmt.Registry) error {

	if err := validate.Required("start", "body", m.Start); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this nested vpc subnet Ip poo type based on context it is used
func (m *NestedVpcSubnetIPPooType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NestedVpcSubnetIPPooType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedVpcSubnetIPPooType) UnmarshalBinary(b []byte) error {
	var res NestedVpcSubnetIPPooType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
