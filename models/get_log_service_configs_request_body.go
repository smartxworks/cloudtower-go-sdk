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

// GetLogServiceConfigsRequestBody get log service configs request body
//
// swagger:model GetLogServiceConfigsRequestBody
type GetLogServiceConfigsRequestBody struct {

	// input
	// Required: true
	Input *LogServiceConfigsInput `json:"input"`

	MarshalOpts *GetLogServiceConfigsRequestBodyMarshalOpts `json:"-"`
}

type GetLogServiceConfigsRequestBodyMarshalOpts struct {
	Input_Explicit_Null_When_Empty bool
}

func (m GetLogServiceConfigsRequestBody) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field input
	if m.Input != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"input\":")
		bytes, err := swag.WriteJSON(m.Input)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Input_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"input\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this get log service configs request body
func (m *GetLogServiceConfigsRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetLogServiceConfigsRequestBody) validateInput(formats strfmt.Registry) error {

	if err := validate.Required("input", "body", m.Input); err != nil {
		return err
	}

	if m.Input != nil {
		if err := m.Input.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("input")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get log service configs request body based on the context it is used
func (m *GetLogServiceConfigsRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInput(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetLogServiceConfigsRequestBody) contextValidateInput(ctx context.Context, formats strfmt.Registry) error {

	if m.Input != nil {
		if err := m.Input.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("input")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetLogServiceConfigsRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetLogServiceConfigsRequestBody) UnmarshalBinary(b []byte) error {
	var res GetLogServiceConfigsRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
