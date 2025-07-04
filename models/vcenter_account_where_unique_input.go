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

// VcenterAccountWhereUniqueInput vcenter account where unique input
//
// swagger:model VcenterAccountWhereUniqueInput
type VcenterAccountWhereUniqueInput struct {

	// id
	ID *string `json:"id,omitempty"`

	// local id
	LocalID *string `json:"local_id,omitempty"`

	MarshalOpts *VcenterAccountWhereUniqueInputMarshalOpts `json:"-"`
}

type VcenterAccountWhereUniqueInputMarshalOpts struct {
	ID_Explicit_Null_When_Empty bool

	LocalID_Explicit_Null_When_Empty bool
}

func (m VcenterAccountWhereUniqueInput) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field id
	if m.ID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id\":")
		bytes, err := swag.WriteJSON(m.ID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id\":null")
		first = false
	}

	// handle nullable field local_id
	if m.LocalID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id\":")
		bytes, err := swag.WriteJSON(m.LocalID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.LocalID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this vcenter account where unique input
func (m *VcenterAccountWhereUniqueInput) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this vcenter account where unique input based on context it is used
func (m *VcenterAccountWhereUniqueInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VcenterAccountWhereUniqueInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VcenterAccountWhereUniqueInput) UnmarshalBinary(b []byte) error {
	var res VcenterAccountWhereUniqueInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
