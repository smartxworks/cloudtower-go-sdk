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

// NestedAggregateWitness nested aggregate witness
//
// swagger:model NestedAggregateWitness
type NestedAggregateWitness struct {

	// count
	// Required: true
	Count *int32 `json:"count"`

	MarshalOpts *NestedAggregateWitnessMarshalOpts `json:"-"`
}

type NestedAggregateWitnessMarshalOpts struct {
	Count_Explicit_Null_When_Empty bool
}

func (m NestedAggregateWitness) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field count
	if m.Count != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"count\":")
		bytes, err := swag.WriteJSON(m.Count)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Count_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"count\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this nested aggregate witness
func (m *NestedAggregateWitness) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedAggregateWitness) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("count", "body", m.Count); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this nested aggregate witness based on context it is used
func (m *NestedAggregateWitness) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NestedAggregateWitness) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedAggregateWitness) UnmarshalBinary(b []byte) error {
	var res NestedAggregateWitness
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
