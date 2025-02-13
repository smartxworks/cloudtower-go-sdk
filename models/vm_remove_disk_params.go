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

// VMRemoveDiskParams Vm remove disk params
//
// swagger:model VmRemoveDiskParams
type VMRemoveDiskParams struct {

	// data
	// Required: true
	Data *VMRemoveDiskParamsData `json:"data"`

	// where
	// Required: true
	Where *VMWhereInput `json:"where"`

	MarshalOpts *VMRemoveDiskParamsMarshalOpts `json:"-"`
}

type VMRemoveDiskParamsMarshalOpts struct {
	Data_Explicit_Null_When_Empty bool

	Where_Explicit_Null_When_Empty bool
}

func (m VMRemoveDiskParams) MarshalJSON() ([]byte, error) {
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

// Validate validates this Vm remove disk params
func (m *VMRemoveDiskParams) Validate(formats strfmt.Registry) error {
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

func (m *VMRemoveDiskParams) validateData(formats strfmt.Registry) error {

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

func (m *VMRemoveDiskParams) validateWhere(formats strfmt.Registry) error {

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

// ContextValidate validate this Vm remove disk params based on the context it is used
func (m *VMRemoveDiskParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *VMRemoveDiskParams) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMRemoveDiskParams) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

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
func (m *VMRemoveDiskParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMRemoveDiskParams) UnmarshalBinary(b []byte) error {
	var res VMRemoveDiskParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VMRemoveDiskParamsData VM remove disk params data
//
// swagger:model VMRemoveDiskParamsData
type VMRemoveDiskParamsData struct {

	// disk ids
	// Required: true
	DiskIds []string `json:"disk_ids"`

	MarshalOpts *VMRemoveDiskParamsDataMarshalOpts `json:"-"`
}

type VMRemoveDiskParamsDataMarshalOpts struct {
	DiskIds_Explicit_Null_When_Empty bool
}

func (m VMRemoveDiskParamsData) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle non nullable field disk_ids without omitempty
	{
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"disk_ids\":")
		bytes, err := swag.WriteJSON(m.DiskIds)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this VM remove disk params data
func (m *VMRemoveDiskParamsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDiskIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMRemoveDiskParamsData) validateDiskIds(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"disk_ids", "body", m.DiskIds); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this VM remove disk params data based on context it is used
func (m *VMRemoveDiskParamsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VMRemoveDiskParamsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMRemoveDiskParamsData) UnmarshalBinary(b []byte) error {
	var res VMRemoveDiskParamsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
