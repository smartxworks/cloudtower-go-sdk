// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RoleUpdationParams role updation params
//
// swagger:model RoleUpdationParams
type RoleUpdationParams struct {

	// data
	// Required: true
	Data *RoleUpdationParamsData `json:"data"`

	// where
	// Required: true
	Where *UserRoleNextWhereInput `json:"where"`
}

// Validate validates this role updation params
func (m *RoleUpdationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
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

func (m *RoleUpdationParams) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *RoleUpdationParams) validateWhere(formats strfmt.Registry) error {

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

// ContextValidate validate this role updation params based on the context it is used
func (m *RoleUpdationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
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

func (m *RoleUpdationParams) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if m.Data != nil {
		if err := m.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *RoleUpdationParams) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

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
func (m *RoleUpdationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoleUpdationParams) UnmarshalBinary(b []byte) error {
	var res RoleUpdationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RoleUpdationParamsData role updation params data
//
// swagger:model RoleUpdationParamsData
type RoleUpdationParamsData struct {

	// actions
	Actions []ROLEACTION `json:"actions,omitempty"`

	// name
	// Min Length: 1
	Name *string `json:"name,omitempty"`
}

// Validate validates this role updation params data
func (m *RoleUpdationParamsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RoleUpdationParamsData) validateActions(formats strfmt.Registry) error {
	if swag.IsZero(m.Actions) { // not required
		return nil
	}

	for i := 0; i < len(m.Actions); i++ {

		if err := m.Actions[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "actions" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "actions" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *RoleUpdationParamsData) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("data"+"."+"name", "body", *m.Name, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this role updation params data based on the context it is used
func (m *RoleUpdationParamsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RoleUpdationParamsData) contextValidateActions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Actions); i++ {

		if err := m.Actions[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "actions" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "actions" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RoleUpdationParamsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoleUpdationParamsData) UnmarshalBinary(b []byte) error {
	var res RoleUpdationParamsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
