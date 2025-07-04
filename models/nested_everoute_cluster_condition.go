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

// NestedEverouteClusterCondition nested everoute cluster condition
//
// swagger:model NestedEverouteClusterCondition
type NestedEverouteClusterCondition struct {

	// last probe time
	LastProbeTime *string `json:"lastProbeTime,omitempty"`

	// type
	// Required: true
	Type *string `json:"type"`

	MarshalOpts *NestedEverouteClusterConditionMarshalOpts `json:"-"`
}

type NestedEverouteClusterConditionMarshalOpts struct {
	LastProbeTime_Explicit_Null_When_Empty bool

	Type_Explicit_Null_When_Empty bool
}

func (m NestedEverouteClusterCondition) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field lastProbeTime
	if m.LastProbeTime != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"lastProbeTime\":")
		bytes, err := swag.WriteJSON(m.LastProbeTime)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.LastProbeTime_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"lastProbeTime\":null")
		first = false
	}

	// handle nullable field type
	if m.Type != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"type\":")
		bytes, err := swag.WriteJSON(m.Type)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Type_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"type\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this nested everoute cluster condition
func (m *NestedEverouteClusterCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedEverouteClusterCondition) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this nested everoute cluster condition based on context it is used
func (m *NestedEverouteClusterCondition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NestedEverouteClusterCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedEverouteClusterCondition) UnmarshalBinary(b []byte) error {
	var res NestedEverouteClusterCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
