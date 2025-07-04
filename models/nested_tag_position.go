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

// NestedTagPosition nested tag position
//
// swagger:model NestedTagPosition
type NestedTagPosition struct {

	// column
	// Required: true
	Column *int32 `json:"column"`

	// row
	// Required: true
	Row *int32 `json:"row"`

	// tag
	// Required: true
	Tag *string `json:"tag"`

	MarshalOpts *NestedTagPositionMarshalOpts `json:"-"`
}

type NestedTagPositionMarshalOpts struct {
	Column_Explicit_Null_When_Empty bool

	Row_Explicit_Null_When_Empty bool

	Tag_Explicit_Null_When_Empty bool
}

func (m NestedTagPosition) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field column
	if m.Column != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"column\":")
		bytes, err := swag.WriteJSON(m.Column)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Column_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"column\":null")
		first = false
	}

	// handle nullable field row
	if m.Row != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"row\":")
		bytes, err := swag.WriteJSON(m.Row)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Row_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"row\":null")
		first = false
	}

	// handle nullable field tag
	if m.Tag != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"tag\":")
		bytes, err := swag.WriteJSON(m.Tag)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Tag_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"tag\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this nested tag position
func (m *NestedTagPosition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColumn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTag(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedTagPosition) validateColumn(formats strfmt.Registry) error {

	if err := validate.Required("column", "body", m.Column); err != nil {
		return err
	}

	return nil
}

func (m *NestedTagPosition) validateRow(formats strfmt.Registry) error {

	if err := validate.Required("row", "body", m.Row); err != nil {
		return err
	}

	return nil
}

func (m *NestedTagPosition) validateTag(formats strfmt.Registry) error {

	if err := validate.Required("tag", "body", m.Tag); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this nested tag position based on context it is used
func (m *NestedTagPosition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NestedTagPosition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedTagPosition) UnmarshalBinary(b []byte) error {
	var res NestedTagPosition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
