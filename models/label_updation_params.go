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

// LabelUpdationParams label updation params
//
// swagger:model LabelUpdationParams
type LabelUpdationParams struct {

	// data
	// Required: true
	Data *LabelUpdationParamsData `json:"data"`

	// where
	// Required: true
	Where *LabelWhereInput `json:"where"`

	MarshalOpts *LabelUpdationParamsMarshalOpts `json:"-"`
}

type LabelUpdationParamsMarshalOpts struct {
	Data_Explicit_Null_When_Empty bool

	Where_Explicit_Null_When_Empty bool
}

func (m LabelUpdationParams) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field data
	if m.Data != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"data\":")
		bytes, err := swag.WriteJSON(m.Data)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Data_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"data\":null")
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

// Validate validates this label updation params
func (m *LabelUpdationParams) Validate(formats strfmt.Registry) error {
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

func (m *LabelUpdationParams) validateData(formats strfmt.Registry) error {

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

func (m *LabelUpdationParams) validateWhere(formats strfmt.Registry) error {

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

// ContextValidate validate this label updation params based on the context it is used
func (m *LabelUpdationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *LabelUpdationParams) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

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

func (m *LabelUpdationParams) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

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
func (m *LabelUpdationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LabelUpdationParams) UnmarshalBinary(b []byte) error {
	var res LabelUpdationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LabelUpdationParamsData label updation params data
//
// swagger:model LabelUpdationParamsData
type LabelUpdationParamsData struct {

	// key
	Key *string `json:"key,omitempty"`

	// value
	Value *string `json:"value,omitempty"`

	MarshalOpts *LabelUpdationParamsDataMarshalOpts `json:"-"`
}

type LabelUpdationParamsDataMarshalOpts struct {
	Key_Explicit_Null_When_Empty bool

	Value_Explicit_Null_When_Empty bool
}

func (m LabelUpdationParamsData) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

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

	// handle nullable field value
	if m.Value != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"value\":")
		bytes, err := swag.WriteJSON(m.Value)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Value_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"value\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this label updation params data
func (m *LabelUpdationParamsData) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this label updation params data based on context it is used
func (m *LabelUpdationParamsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LabelUpdationParamsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LabelUpdationParamsData) UnmarshalBinary(b []byte) error {
	var res LabelUpdationParamsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
