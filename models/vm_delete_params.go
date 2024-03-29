// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VMDeleteParams Vm delete params
//
// swagger:model VmDeleteParams
type VMDeleteParams struct {

	// effect
	Effect *VMDeleteParamsEffect `json:"effect,omitempty"`

	// where
	// Required: true
	Where *VMWhereInput `json:"where"`
}

// Validate validates this Vm delete params
func (m *VMDeleteParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEffect(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWhere(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMDeleteParams) validateEffect(formats strfmt.Registry) error {
	if swag.IsZero(m.Effect) { // not required
		return nil
	}

	if m.Effect != nil {
		if err := m.Effect.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("effect")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("effect")
			}
			return err
		}
	}

	return nil
}

func (m *VMDeleteParams) validateWhere(formats strfmt.Registry) error {

	if err := validate.Required("where", "body", m.Where); err != nil {
		return err
	}

	if m.Where != nil {
		if err := m.Where.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("where")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("where")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this Vm delete params based on the context it is used
func (m *VMDeleteParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEffect(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWhere(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMDeleteParams) contextValidateEffect(ctx context.Context, formats strfmt.Registry) error {

	if m.Effect != nil {
		if err := m.Effect.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("effect")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("effect")
			}
			return err
		}
	}

	return nil
}

func (m *VMDeleteParams) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

	if m.Where != nil {
		if err := m.Where.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("where")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("where")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMDeleteParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMDeleteParams) UnmarshalBinary(b []byte) error {
	var res VMDeleteParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VMDeleteParamsEffect VM delete params effect
//
// swagger:model VMDeleteParamsEffect
type VMDeleteParamsEffect struct {

	// include snapshots
	IncludeSnapshots *bool `json:"include_snapshots,omitempty"`
}

// Validate validates this VM delete params effect
func (m *VMDeleteParamsEffect) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this VM delete params effect based on context it is used
func (m *VMDeleteParamsEffect) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VMDeleteParamsEffect) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMDeleteParamsEffect) UnmarshalBinary(b []byte) error {
	var res VMDeleteParamsEffect
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
