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

// VMRemoveNicByWhereParams Vm remove nic by where params
//
// swagger:model VmRemoveNicByWhereParams
type VMRemoveNicByWhereParams struct {

	// effect
	// Required: true
	Effect *VMRemoveNicByWhereParamsEffect `json:"effect"`

	// where
	// Required: true
	Where *VMNicWhereInput `json:"where"`

	MarshalOpts *VMRemoveNicByWhereParamsMarshalOpts `json:"-"`
}

type VMRemoveNicByWhereParamsMarshalOpts struct {
	Effect_Explicit_Null_When_Empty bool

	Where_Explicit_Null_When_Empty bool
}

func (m VMRemoveNicByWhereParams) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field effect
	if m.Effect != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"effect\":")
		bytes, err := swag.WriteJSON(m.Effect)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Effect_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"effect\":null")
		first = false
	}

	// handle nullable field where
	if m.Where != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"where\":")
		bytes, err := swag.WriteJSON(m.Where)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Where_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"where\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this Vm remove nic by where params
func (m *VMRemoveNicByWhereParams) Validate(formats strfmt.Registry) error {
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

func (m *VMRemoveNicByWhereParams) validateEffect(formats strfmt.Registry) error {

	if err := validate.Required("effect", "body", m.Effect); err != nil {
		return err
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

func (m *VMRemoveNicByWhereParams) validateWhere(formats strfmt.Registry) error {

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

// ContextValidate validate this Vm remove nic by where params based on the context it is used
func (m *VMRemoveNicByWhereParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *VMRemoveNicByWhereParams) contextValidateEffect(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMRemoveNicByWhereParams) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

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
func (m *VMRemoveNicByWhereParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMRemoveNicByWhereParams) UnmarshalBinary(b []byte) error {
	var res VMRemoveNicByWhereParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VMRemoveNicByWhereParamsEffect VM remove nic by where params effect
//
// swagger:model VMRemoveNicByWhereParamsEffect
type VMRemoveNicByWhereParamsEffect struct {

	// vm ids
	// Min Items: 1
	VMIds []string `json:"vm_ids,omitempty"`

	MarshalOpts *VMRemoveNicByWhereParamsEffectMarshalOpts `json:"-"`
}

type VMRemoveNicByWhereParamsEffectMarshalOpts struct {
	VMIds_Explicit_Null_When_Empty bool
}

func (m VMRemoveNicByWhereParamsEffect) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle non nullable field vm_ids with omitempty
	if !swag.IsZero(m.VMIds) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_ids\":")
		bytes, err := swag.WriteJSON(m.VMIds)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this VM remove nic by where params effect
func (m *VMRemoveNicByWhereParamsEffect) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVMIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMRemoveNicByWhereParamsEffect) validateVMIds(formats strfmt.Registry) error {
	if swag.IsZero(m.VMIds) { // not required
		return nil
	}

	iVMIdsSize := int64(len(m.VMIds))

	if err := validate.MinItems("effect"+"."+"vm_ids", "body", iVMIdsSize, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this VM remove nic by where params effect based on context it is used
func (m *VMRemoveNicByWhereParamsEffect) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VMRemoveNicByWhereParamsEffect) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMRemoveNicByWhereParamsEffect) UnmarshalBinary(b []byte) error {
	var res VMRemoveNicByWhereParamsEffect
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
