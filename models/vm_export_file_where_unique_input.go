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

// VMExportFileWhereUniqueInput Vm export file where unique input
//
// swagger:model VmExportFileWhereUniqueInput
type VMExportFileWhereUniqueInput struct {

	// data port id
	DataPortID *string `json:"data_port_id,omitempty"`

	// id
	ID *string `json:"id,omitempty"`

	MarshalOpts *VMExportFileWhereUniqueInputMarshalOpts `json:"-"`
}

type VMExportFileWhereUniqueInputMarshalOpts struct {
	DataPortID_Explicit_Null_When_Empty bool

	ID_Explicit_Null_When_Empty bool
}

func (m VMExportFileWhereUniqueInput) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field data_port_id
	if m.DataPortID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"data_port_id\":")
		bytes, err := swag.WriteJSON(m.DataPortID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.DataPortID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"data_port_id\":null")
		first = false
	}

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

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this Vm export file where unique input
func (m *VMExportFileWhereUniqueInput) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this Vm export file where unique input based on context it is used
func (m *VMExportFileWhereUniqueInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VMExportFileWhereUniqueInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMExportFileWhereUniqueInput) UnmarshalBinary(b []byte) error {
	var res VMExportFileWhereUniqueInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
