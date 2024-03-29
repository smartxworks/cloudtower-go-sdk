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

// VMUpdateDiskParams Vm update disk params
//
// swagger:model VmUpdateDiskParams
type VMUpdateDiskParams struct {

	// data
	// Required: true
	Data *VMUpdateDiskParamsData `json:"data"`

	// where
	// Required: true
	Where *VMWhereInput `json:"where"`
}

// Validate validates this Vm update disk params
func (m *VMUpdateDiskParams) Validate(formats strfmt.Registry) error {
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

func (m *VMUpdateDiskParams) validateData(formats strfmt.Registry) error {

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

func (m *VMUpdateDiskParams) validateWhere(formats strfmt.Registry) error {

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

// ContextValidate validate this Vm update disk params based on the context it is used
func (m *VMUpdateDiskParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *VMUpdateDiskParams) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMUpdateDiskParams) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

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
func (m *VMUpdateDiskParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMUpdateDiskParams) UnmarshalBinary(b []byte) error {
	var res VMUpdateDiskParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VMUpdateDiskParamsData VM update disk params data
//
// swagger:model VMUpdateDiskParamsData
type VMUpdateDiskParamsData struct {

	// bus
	Bus *Bus `json:"bus,omitempty"`

	// content library image id
	ContentLibraryImageID *string `json:"content_library_image_id,omitempty"`

	// elf image id
	ElfImageID *string `json:"elf_image_id,omitempty"`

	// vm disk id
	// Required: true
	VMDiskID *string `json:"vm_disk_id"`

	// vm volume id
	VMVolumeID *string `json:"vm_volume_id,omitempty"`
}

// Validate validates this VM update disk params data
func (m *VMUpdateDiskParamsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMDiskID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMUpdateDiskParamsData) validateBus(formats strfmt.Registry) error {
	if swag.IsZero(m.Bus) { // not required
		return nil
	}

	if m.Bus != nil {
		if err := m.Bus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "bus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "bus")
			}
			return err
		}
	}

	return nil
}

func (m *VMUpdateDiskParamsData) validateVMDiskID(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"vm_disk_id", "body", m.VMDiskID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this VM update disk params data based on the context it is used
func (m *VMUpdateDiskParamsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMUpdateDiskParamsData) contextValidateBus(ctx context.Context, formats strfmt.Registry) error {

	if m.Bus != nil {
		if err := m.Bus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "bus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "bus")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMUpdateDiskParamsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMUpdateDiskParamsData) UnmarshalBinary(b []byte) error {
	var res VMUpdateDiskParamsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
