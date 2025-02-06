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

// VMRemoveCdRomParams Vm remove cd rom params
//
// swagger:model VmRemoveCdRomParams
type VMRemoveCdRomParams struct {

	// data
	// Required: true
	Data *VMRemoveCdRomParamsData `json:"data"`

	// where
	// Required: true
	Where *VMWhereInput `json:"where"`

	MarshalOpts *VMRemoveCdRomParamsMarshalOpts `json:"-"`
}

type VMRemoveCdRomParamsMarshalOpts struct {
	Data_Explicit_Null_When_Empty bool

	Where_Explicit_Null_When_Empty bool
}

func (m VMRemoveCdRomParams) MarshalJSON() ([]byte, error) {
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

// Validate validates this Vm remove cd rom params
func (m *VMRemoveCdRomParams) Validate(formats strfmt.Registry) error {
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

func (m *VMRemoveCdRomParams) validateData(formats strfmt.Registry) error {

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

func (m *VMRemoveCdRomParams) validateWhere(formats strfmt.Registry) error {

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

// ContextValidate validate this Vm remove cd rom params based on the context it is used
func (m *VMRemoveCdRomParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *VMRemoveCdRomParams) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMRemoveCdRomParams) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

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
func (m *VMRemoveCdRomParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMRemoveCdRomParams) UnmarshalBinary(b []byte) error {
	var res VMRemoveCdRomParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VMRemoveCdRomParamsData VM remove cd rom params data
//
// swagger:model VMRemoveCdRomParamsData
type VMRemoveCdRomParamsData struct {

	// cd rom ids
	// Required: true
	CdRomIds []string `json:"cd_rom_ids"`

	MarshalOpts *VMRemoveCdRomParamsDataMarshalOpts `json:"-"`
}

type VMRemoveCdRomParamsDataMarshalOpts struct {
}

func (m VMRemoveCdRomParamsData) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle non nullable field cd_rom_ids without omitempty
	if !first {
		b.WriteString(",")
	}
	b.WriteString("\"cd_rom_ids\":")
	bytes, err := swag.WriteJSON(m.CdRomIds)
	if err != nil {
		return nil, err
	}
	b.Write(bytes)
	first = false

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this VM remove cd rom params data
func (m *VMRemoveCdRomParamsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCdRomIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMRemoveCdRomParamsData) validateCdRomIds(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"cd_rom_ids", "body", m.CdRomIds); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this VM remove cd rom params data based on context it is used
func (m *VMRemoveCdRomParamsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VMRemoveCdRomParamsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMRemoveCdRomParamsData) UnmarshalBinary(b []byte) error {
	var res VMRemoveCdRomParamsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
