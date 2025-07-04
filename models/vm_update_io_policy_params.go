// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VMUpdateIoPolicyParams Vm update io policy params
//
// swagger:model VmUpdateIoPolicyParams
type VMUpdateIoPolicyParams struct {

	// data
	// Required: true
	Data *VMUpdateIoPolicyParamsData `json:"data"`

	// where
	// Required: true
	Where *VMWhereInput `json:"where"`

	MarshalOpts *VMUpdateIoPolicyParamsMarshalOpts `json:"-"`
}

type VMUpdateIoPolicyParamsMarshalOpts struct {
	Data_Explicit_Null_When_Empty bool

	Where_Explicit_Null_When_Empty bool
}

func (m VMUpdateIoPolicyParams) MarshalJSON() ([]byte, error) {
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

// Validate validates this Vm update io policy params
func (m *VMUpdateIoPolicyParams) Validate(formats strfmt.Registry) error {
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

func (m *VMUpdateIoPolicyParams) validateData(formats strfmt.Registry) error {

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

func (m *VMUpdateIoPolicyParams) validateWhere(formats strfmt.Registry) error {

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

// ContextValidate validate this Vm update io policy params based on the context it is used
func (m *VMUpdateIoPolicyParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *VMUpdateIoPolicyParams) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMUpdateIoPolicyParams) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

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
func (m *VMUpdateIoPolicyParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMUpdateIoPolicyParams) UnmarshalBinary(b []byte) error {
	var res VMUpdateIoPolicyParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VMUpdateIoPolicyParamsData VM update io policy params data
//
// swagger:model VMUpdateIoPolicyParamsData
type VMUpdateIoPolicyParamsData struct {

	// each disk
	EachDisk []*VMUpdateEachDiskIoPolicyParams `json:"each_disk,omitempty"`

	// io policy
	IoPolicy *VMDiskIoPolicy `json:"io_policy,omitempty"`

	// whole vm
	WholeVM *VMRestrictIoParamsData `json:"whole_vm,omitempty"`

	MarshalOpts *VMUpdateIoPolicyParamsDataMarshalOpts `json:"-"`
}

type VMUpdateIoPolicyParamsDataMarshalOpts struct {
	EachDisk_Explicit_Null_When_Empty bool

	IoPolicy_Explicit_Null_When_Empty bool

	WholeVM_Explicit_Null_When_Empty bool
}

func (m VMUpdateIoPolicyParamsData) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle non nullable field each_disk with omitempty
	if !swag.IsZero(m.EachDisk) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"each_disk\":")
		bytes, err := swag.WriteJSON(m.EachDisk)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field io_policy
	if m.IoPolicy != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"io_policy\":")
		bytes, err := swag.WriteJSON(m.IoPolicy)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IoPolicy_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"io_policy\":null")
		first = false
	}

	// handle nullable field whole_vm
	if m.WholeVM != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"whole_vm\":")
		bytes, err := swag.WriteJSON(m.WholeVM)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.WholeVM_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"whole_vm\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this VM update io policy params data
func (m *VMUpdateIoPolicyParamsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEachDisk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIoPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWholeVM(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMUpdateIoPolicyParamsData) validateEachDisk(formats strfmt.Registry) error {
	if swag.IsZero(m.EachDisk) { // not required
		return nil
	}

	for i := 0; i < len(m.EachDisk); i++ {
		if swag.IsZero(m.EachDisk[i]) { // not required
			continue
		}

		if m.EachDisk[i] != nil {
			if err := m.EachDisk[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + "each_disk" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + "each_disk" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMUpdateIoPolicyParamsData) validateIoPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.IoPolicy) { // not required
		return nil
	}

	if m.IoPolicy != nil {
		if err := m.IoPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "io_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "io_policy")
			}
			return err
		}
	}

	return nil
}

func (m *VMUpdateIoPolicyParamsData) validateWholeVM(formats strfmt.Registry) error {
	if swag.IsZero(m.WholeVM) { // not required
		return nil
	}

	if m.WholeVM != nil {
		if err := m.WholeVM.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "whole_vm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "whole_vm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this VM update io policy params data based on the context it is used
func (m *VMUpdateIoPolicyParamsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEachDisk(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIoPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWholeVM(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMUpdateIoPolicyParamsData) contextValidateEachDisk(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EachDisk); i++ {

		if m.EachDisk[i] != nil {
			if err := m.EachDisk[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + "each_disk" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + "each_disk" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMUpdateIoPolicyParamsData) contextValidateIoPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.IoPolicy != nil {
		if err := m.IoPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "io_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "io_policy")
			}
			return err
		}
	}

	return nil
}

func (m *VMUpdateIoPolicyParamsData) contextValidateWholeVM(ctx context.Context, formats strfmt.Registry) error {

	if m.WholeVM != nil {
		if err := m.WholeVM.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "whole_vm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "whole_vm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMUpdateIoPolicyParamsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMUpdateIoPolicyParamsData) UnmarshalBinary(b []byte) error {
	var res VMUpdateIoPolicyParamsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
