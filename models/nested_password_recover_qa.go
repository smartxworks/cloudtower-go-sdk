// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NestedPasswordRecoverQa nested password recover qa
//
// swagger:model NestedPasswordRecoverQa
type NestedPasswordRecoverQa struct {

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// items
	Items []*NestedPasswordReoverQaItem `json:"items,omitempty"`

	MarshalOpts *NestedPasswordRecoverQaMarshalOpts `json:"-"`
}

type NestedPasswordRecoverQaMarshalOpts struct {
	Enabled_Explicit_Null_When_Empty bool

	Items_Explicit_Null_When_Empty bool
}

func (m NestedPasswordRecoverQa) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field enabled
	if m.Enabled != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"enabled\":")
		bytes, err := swag.WriteJSON(m.Enabled)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Enabled_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"enabled\":null")
		first = false
	}

	// handle non nullable field items with omitempty
	if !swag.IsZero(m.Items) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"items\":")
		bytes, err := swag.WriteJSON(m.Items)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this nested password recover qa
func (m *NestedPasswordRecoverQa) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedPasswordRecoverQa) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *NestedPasswordRecoverQa) validateItems(formats strfmt.Registry) error {
	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this nested password recover qa based on the context it is used
func (m *NestedPasswordRecoverQa) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedPasswordRecoverQa) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Items); i++ {

		if m.Items[i] != nil {
			if err := m.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NestedPasswordRecoverQa) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedPasswordRecoverQa) UnmarshalBinary(b []byte) error {
	var res NestedPasswordRecoverQa
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
