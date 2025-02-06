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

// VmdkDiskModify vmdk disk modify
//
// swagger:model VmdkDiskModify
type VmdkDiskModify struct {

	// boot
	Boot *int32 `json:"boot,omitempty"`

	// bus
	Bus *Bus `json:"bus,omitempty"`

	// elf storage policy
	ElfStoragePolicy *VMVolumeElfStoragePolicyType `json:"elf_storage_policy,omitempty"`

	// vmdk name
	// Required: true
	VmdkName *string `json:"vmdk_name"`

	// volume name
	VolumeName *string `json:"volume_name,omitempty"`

	MarshalOpts *VmdkDiskModifyMarshalOpts `json:"-"`
}

type VmdkDiskModifyMarshalOpts struct {
	Boot_Explicit_Null_When_Empty bool

	Bus_Explicit_Null_When_Empty bool

	ElfStoragePolicy_Explicit_Null_When_Empty bool

	VmdkName_Explicit_Null_When_Empty bool

	VolumeName_Explicit_Null_When_Empty bool
}

func (m VmdkDiskModify) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field boot
	if m.Boot != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"boot\":")
		bytes, err := swag.WriteJSON(m.Boot)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Boot_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"boot\":null")
		first = false
	}

	// handle nullable field bus
	if m.Bus != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bus\":")
		bytes, err := swag.WriteJSON(m.Bus)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Bus_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bus\":null")
		first = false
	}

	// handle nullable field elf_storage_policy
	if m.ElfStoragePolicy != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"elf_storage_policy\":")
		bytes, err := swag.WriteJSON(m.ElfStoragePolicy)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ElfStoragePolicy_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"elf_storage_policy\":null")
		first = false
	}

	// handle nullable field vmdk_name
	if m.VmdkName != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vmdk_name\":")
		bytes, err := swag.WriteJSON(m.VmdkName)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VmdkName_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vmdk_name\":null")
		first = false
	}

	// handle nullable field volume_name
	if m.VolumeName != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"volume_name\":")
		bytes, err := swag.WriteJSON(m.VolumeName)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VolumeName_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"volume_name\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this vmdk disk modify
func (m *VmdkDiskModify) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElfStoragePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVmdkName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VmdkDiskModify) validateBus(formats strfmt.Registry) error {
	if swag.IsZero(m.Bus) { // not required
		return nil
	}

	if m.Bus != nil {
		if err := m.Bus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bus")
			}
			return err
		}
	}

	return nil
}

func (m *VmdkDiskModify) validateElfStoragePolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.ElfStoragePolicy) { // not required
		return nil
	}

	if m.ElfStoragePolicy != nil {
		if err := m.ElfStoragePolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("elf_storage_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("elf_storage_policy")
			}
			return err
		}
	}

	return nil
}

func (m *VmdkDiskModify) validateVmdkName(formats strfmt.Registry) error {

	if err := validate.Required("vmdk_name", "body", m.VmdkName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this vmdk disk modify based on the context it is used
func (m *VmdkDiskModify) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateElfStoragePolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VmdkDiskModify) contextValidateBus(ctx context.Context, formats strfmt.Registry) error {

	if m.Bus != nil {
		if err := m.Bus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bus")
			}
			return err
		}
	}

	return nil
}

func (m *VmdkDiskModify) contextValidateElfStoragePolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.ElfStoragePolicy != nil {
		if err := m.ElfStoragePolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("elf_storage_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("elf_storage_policy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VmdkDiskModify) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmdkDiskModify) UnmarshalBinary(b []byte) error {
	var res VmdkDiskModify
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
