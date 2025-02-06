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

// AbortMigrateVMAcrossClusterParams abort migrate Vm across cluster params
//
// swagger:model AbortMigrateVmAcrossClusterParams
type AbortMigrateVMAcrossClusterParams struct {

	// tasks
	// Required: true
	Tasks *TaskWhereInput `json:"tasks"`

	MarshalOpts *AbortMigrateVMAcrossClusterParamsMarshalOpts `json:"-"`
}

type AbortMigrateVMAcrossClusterParamsMarshalOpts struct {
	Tasks_Explicit_Null_When_Empty bool
}

func (m AbortMigrateVMAcrossClusterParams) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field tasks
	if m.Tasks != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"tasks\":")
		bytes, err := swag.WriteJSON(m.Tasks)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Tasks_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"tasks\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this abort migrate Vm across cluster params
func (m *AbortMigrateVMAcrossClusterParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTasks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AbortMigrateVMAcrossClusterParams) validateTasks(formats strfmt.Registry) error {

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

// ContextValidate validate this abort migrate Vm across cluster params based on the context it is used
func (m *AbortMigrateVMAcrossClusterParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTasks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AbortMigrateVMAcrossClusterParams) contextValidateTasks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AbortMigrateVMAcrossClusterParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AbortMigrateVMAcrossClusterParams) UnmarshalBinary(b []byte) error {
	var res AbortMigrateVMAcrossClusterParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
