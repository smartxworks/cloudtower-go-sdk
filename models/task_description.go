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

// TaskDescription task description
//
// swagger:model TaskDescription
type TaskDescription struct {

	// en u s
	// Required: true
	EnUS *string `json:"en-US"`

	// zh c n
	// Required: true
	ZhCN *string `json:"zh-CN"`
}

// Validate validates this task description
func (m *TaskDescription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnUS(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZhCN(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskDescription) validateEnUS(formats strfmt.Registry) error {

	if err := validate.Required("en-US", "body", m.EnUS); err != nil {
		return err
	}

	return nil
}

func (m *TaskDescription) validateZhCN(formats strfmt.Registry) error {

	if err := validate.Required("zh-CN", "body", m.ZhCN); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this task description based on context it is used
func (m *TaskDescription) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TaskDescription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskDescription) UnmarshalBinary(b []byte) error {
	var res TaskDescription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
