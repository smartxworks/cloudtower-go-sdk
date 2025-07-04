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

// EnterMaintenanceModeCheckResult enter maintenance mode check result
//
// swagger:model EnterMaintenanceModeCheckResult
type EnterMaintenanceModeCheckResult struct {

	// task id
	// Required: true
	TaskID *string `json:"task_id"`

	MarshalOpts *EnterMaintenanceModeCheckResultMarshalOpts `json:"-"`
}

type EnterMaintenanceModeCheckResultMarshalOpts struct {
	TaskID_Explicit_Null_When_Empty bool
}

func (m EnterMaintenanceModeCheckResult) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field task_id
	if m.TaskID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"task_id\":")
		bytes, err := swag.WriteJSON(m.TaskID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.TaskID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"task_id\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this enter maintenance mode check result
func (m *EnterMaintenanceModeCheckResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTaskID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnterMaintenanceModeCheckResult) validateTaskID(formats strfmt.Registry) error {

	if err := validate.Required("task_id", "body", m.TaskID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this enter maintenance mode check result based on context it is used
func (m *EnterMaintenanceModeCheckResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EnterMaintenanceModeCheckResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnterMaintenanceModeCheckResult) UnmarshalBinary(b []byte) error {
	var res EnterMaintenanceModeCheckResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
