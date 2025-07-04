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

// LabelCreationParams label creation params
//
// swagger:model LabelCreationParams
type LabelCreationParams struct {

	// key
	// Required: true
	Key *string `json:"key"`

	// value
	Value *string `json:"value,omitempty"`

	MarshalOpts *LabelCreationParamsMarshalOpts `json:"-"`
}

type LabelCreationParamsMarshalOpts struct {
	Key_Explicit_Null_When_Empty bool

	Value_Explicit_Null_When_Empty bool
}

func (m LabelCreationParams) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field key
	if m.Key != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"key\":")
		bytes, err := swag.WriteJSON(m.Key)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Key_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"key\":null")
		first = false
	}

	// handle nullable field value
	if m.Value != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"value\":")
		bytes, err := swag.WriteJSON(m.Value)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Value_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"value\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this label creation params
func (m *LabelCreationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LabelCreationParams) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this label creation params based on context it is used
func (m *LabelCreationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LabelCreationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LabelCreationParams) UnmarshalBinary(b []byte) error {
	var res LabelCreationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
