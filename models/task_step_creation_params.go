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

// TaskStepCreationParams task step creation params
//
// swagger:model TaskStepCreationParams
type TaskStepCreationParams struct {

	// description
	Description *string `json:"description,omitempty"`

	// finished
	// Required: true
	Finished *bool `json:"finished"`

	// key
	// Required: true
	// Min Length: 1
	Key *string `json:"key"`

	MarshalOpts *TaskStepCreationParamsMarshalOpts `json:"-"`
}

type TaskStepCreationParamsMarshalOpts struct {
	Description_Explicit_Null_When_Empty bool

	Finished_Explicit_Null_When_Empty bool

	Key_Explicit_Null_When_Empty bool
}

func (m TaskStepCreationParams) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field description
	if m.Description != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description\":")
		bytes, err := swag.WriteJSON(m.Description)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Description_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description\":null")
		first = false
	}

	// handle nullable field finished
	if m.Finished != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"finished\":")
		bytes, err := swag.WriteJSON(m.Finished)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Finished_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"finished\":null")
		first = false
	}

	// handle nullable field key
	if m.Key != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"key\":")
		bytes, err := swag.WriteJSON(m.Key)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Key_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"key\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this task step creation params
func (m *TaskStepCreationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFinished(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskStepCreationParams) validateFinished(formats strfmt.Registry) error {

	if err := validate.Required("finished", "body", m.Finished); err != nil {
		return err
	}

	return nil
}

func (m *TaskStepCreationParams) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	if err := validate.MinLength("key", "body", *m.Key, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this task step creation params based on context it is used
func (m *TaskStepCreationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TaskStepCreationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskStepCreationParams) UnmarshalBinary(b []byte) error {
	var res TaskStepCreationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
