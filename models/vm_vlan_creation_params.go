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

// VMVlanCreationParams Vm vlan creation params
//
// swagger:model VmVlanCreationParams
type VMVlanCreationParams struct {

	// name
	// Required: true
	Name *string `json:"name"`

	// vds id
	// Required: true
	VdsID *string `json:"vds_id"`

	// vlan id
	// Required: true
	VlanID *VlanID `json:"vlan_id"`
}

// Validate validates this Vm vlan creation params
func (m *VMVlanCreationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVdsID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlanID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMVlanCreationParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VMVlanCreationParams) validateVdsID(formats strfmt.Registry) error {

	if err := validate.Required("vds_id", "body", m.VdsID); err != nil {
		return err
	}

	return nil
}

func (m *VMVlanCreationParams) validateVlanID(formats strfmt.Registry) error {

	if err := validate.Required("vlan_id", "body", m.VlanID); err != nil {
		return err
	}

	if err := validate.Required("vlan_id", "body", m.VlanID); err != nil {
		return err
	}

	if m.VlanID != nil {
		if err := m.VlanID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlan_id")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlan_id")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this Vm vlan creation params based on the context it is used
func (m *VMVlanCreationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVlanID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMVlanCreationParams) contextValidateVlanID(ctx context.Context, formats strfmt.Registry) error {

	if m.VlanID != nil {
		if err := m.VlanID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlan_id")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlan_id")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMVlanCreationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMVlanCreationParams) UnmarshalBinary(b []byte) error {
	var res VMVlanCreationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
