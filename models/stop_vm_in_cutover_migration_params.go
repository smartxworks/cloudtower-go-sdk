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

// StopVMInCutoverMigrationParams stop Vm in cutover migration params
//
// swagger:model StopVmInCutoverMigrationParams
type StopVMInCutoverMigrationParams struct {

	// force
	Force *bool `json:"force,omitempty"`

	// tasks
	// Required: true
	Tasks *TaskWhereInput `json:"tasks"`
}

// Validate validates this stop Vm in cutover migration params
func (m *StopVMInCutoverMigrationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTasks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StopVMInCutoverMigrationParams) validateTasks(formats strfmt.Registry) error {

	if err := validate.Required("tasks", "body", m.Tasks); err != nil {
		return err
	}

	if m.Tasks != nil {
		if err := m.Tasks.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tasks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tasks")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this stop Vm in cutover migration params based on the context it is used
func (m *StopVMInCutoverMigrationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTasks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StopVMInCutoverMigrationParams) contextValidateTasks(ctx context.Context, formats strfmt.Registry) error {

	if m.Tasks != nil {
		if err := m.Tasks.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tasks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tasks")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StopVMInCutoverMigrationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StopVMInCutoverMigrationParams) UnmarshalBinary(b []byte) error {
	var res StopVMInCutoverMigrationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
