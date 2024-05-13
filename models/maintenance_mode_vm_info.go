// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MaintenanceModeVMInfo maintenance mode Vm info
//
// swagger:model MaintenanceModeVmInfo
type MaintenanceModeVMInfo struct {

	// state
	State *string `json:"state,omitempty"`

	// target host name
	TargetHostName *string `json:"target_host_name,omitempty"`

	// verify
	Verify *MaintenanceModeVerify `json:"verify,omitempty"`

	// vm ha
	VMHa *bool `json:"vm_ha,omitempty"`

	// vm name
	VMName *string `json:"vm_name,omitempty"`

	// vm state
	VMState *string `json:"vm_state,omitempty"`

	// vm uuid
	VMUUID *string `json:"vm_uuid,omitempty"`
}

// Validate validates this maintenance mode Vm info
func (m *MaintenanceModeVMInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVerify(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MaintenanceModeVMInfo) validateVerify(formats strfmt.Registry) error {
	if swag.IsZero(m.Verify) { // not required
		return nil
	}

	if m.Verify != nil {
		if err := m.Verify.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("verify")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("verify")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this maintenance mode Vm info based on the context it is used
func (m *MaintenanceModeVMInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVerify(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MaintenanceModeVMInfo) contextValidateVerify(ctx context.Context, formats strfmt.Registry) error {

	if m.Verify != nil {
		if err := m.Verify.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("verify")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("verify")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MaintenanceModeVMInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MaintenanceModeVMInfo) UnmarshalBinary(b []byte) error {
	var res MaintenanceModeVMInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}