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

// MountNewCreateDisksParams mount new create disks params
//
// swagger:model MountNewCreateDisksParams
type MountNewCreateDisksParams struct {

	// boot
	// Required: true
	Boot *int32 `json:"boot"`

	// bus
	// Required: true
	Bus *Bus `json:"bus"`

	// index
	Index *int32 `json:"index,omitempty"`

	// key
	Key *int32 `json:"key,omitempty"`

	// max bandwidth
	MaxBandwidth *int64 `json:"max_bandwidth,omitempty"`

	// max bandwidth policy
	MaxBandwidthPolicy *VMDiskIoRestrictType `json:"max_bandwidth_policy,omitempty"`

	// max bandwidth unit
	MaxBandwidthUnit *BPSUnit `json:"max_bandwidth_unit,omitempty"`

	// max iops
	MaxIops *int64 `json:"max_iops,omitempty"`

	// max iops policy
	MaxIopsPolicy *VMDiskIoRestrictType `json:"max_iops_policy,omitempty"`

	// vm volume
	// Required: true
	VMVolume *MountNewCreateDisksParamsVMVolume `json:"vm_volume"`

	MarshalOpts *MountNewCreateDisksParamsMarshalOpts `json:"-"`
}

type MountNewCreateDisksParamsMarshalOpts struct {
	Boot_Explicit_Null_When_Empty bool

	Bus_Explicit_Null_When_Empty bool

	Index_Explicit_Null_When_Empty bool

	Key_Explicit_Null_When_Empty bool

	MaxBandwidth_Explicit_Null_When_Empty bool

	MaxBandwidthPolicy_Explicit_Null_When_Empty bool

	MaxBandwidthUnit_Explicit_Null_When_Empty bool

	MaxIops_Explicit_Null_When_Empty bool

	MaxIopsPolicy_Explicit_Null_When_Empty bool

	VMVolume_Explicit_Null_When_Empty bool
}

func (m MountNewCreateDisksParams) MarshalJSON() ([]byte, error) {
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

	// handle nullable field index
	if m.Index != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"index\":")
		bytes, err := swag.WriteJSON(m.Index)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Index_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"index\":null")
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

	// handle nullable field max_bandwidth
	if m.MaxBandwidth != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"max_bandwidth\":")
		bytes, err := swag.WriteJSON(m.MaxBandwidth)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.MaxBandwidth_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"max_bandwidth\":null")
		first = false
	}

	// handle nullable field max_bandwidth_policy
	if m.MaxBandwidthPolicy != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"max_bandwidth_policy\":")
		bytes, err := swag.WriteJSON(m.MaxBandwidthPolicy)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.MaxBandwidthPolicy_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"max_bandwidth_policy\":null")
		first = false
	}

	// handle nullable field max_bandwidth_unit
	if m.MaxBandwidthUnit != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"max_bandwidth_unit\":")
		bytes, err := swag.WriteJSON(m.MaxBandwidthUnit)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.MaxBandwidthUnit_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"max_bandwidth_unit\":null")
		first = false
	}

	// handle nullable field max_iops
	if m.MaxIops != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"max_iops\":")
		bytes, err := swag.WriteJSON(m.MaxIops)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.MaxIops_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"max_iops\":null")
		first = false
	}

	// handle nullable field max_iops_policy
	if m.MaxIopsPolicy != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"max_iops_policy\":")
		bytes, err := swag.WriteJSON(m.MaxIopsPolicy)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.MaxIopsPolicy_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"max_iops_policy\":null")
		first = false
	}

	// handle nullable field vm_volume
	if m.VMVolume != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_volume\":")
		bytes, err := swag.WriteJSON(m.VMVolume)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VMVolume_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_volume\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this mount new create disks params
func (m *MountNewCreateDisksParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBoot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxBandwidthPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxBandwidthUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxIopsPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMVolume(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MountNewCreateDisksParams) validateBoot(formats strfmt.Registry) error {

	if err := validate.Required("boot", "body", m.Boot); err != nil {
		return err
	}

	return nil
}

func (m *MountNewCreateDisksParams) validateBus(formats strfmt.Registry) error {

	if err := validate.Required("bus", "body", m.Bus); err != nil {
		return err
	}

	if err := validate.Required("bus", "body", m.Bus); err != nil {
		return err
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

func (m *MountNewCreateDisksParams) validateMaxBandwidthPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxBandwidthPolicy) { // not required
		return nil
	}

	if m.MaxBandwidthPolicy != nil {
		if err := m.MaxBandwidthPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("max_bandwidth_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("max_bandwidth_policy")
			}
			return err
		}
	}

	return nil
}

func (m *MountNewCreateDisksParams) validateMaxBandwidthUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxBandwidthUnit) { // not required
		return nil
	}

	if m.MaxBandwidthUnit != nil {
		if err := m.MaxBandwidthUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("max_bandwidth_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("max_bandwidth_unit")
			}
			return err
		}
	}

	return nil
}

func (m *MountNewCreateDisksParams) validateMaxIopsPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxIopsPolicy) { // not required
		return nil
	}

	if m.MaxIopsPolicy != nil {
		if err := m.MaxIopsPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("max_iops_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("max_iops_policy")
			}
			return err
		}
	}

	return nil
}

func (m *MountNewCreateDisksParams) validateVMVolume(formats strfmt.Registry) error {

	if err := validate.Required("vm_volume", "body", m.VMVolume); err != nil {
		return err
	}

	if m.VMVolume != nil {
		if err := m.VMVolume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_volume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_volume")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this mount new create disks params based on the context it is used
func (m *MountNewCreateDisksParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxBandwidthPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxBandwidthUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxIopsPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MountNewCreateDisksParams) contextValidateBus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *MountNewCreateDisksParams) contextValidateMaxBandwidthPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.MaxBandwidthPolicy != nil {
		if err := m.MaxBandwidthPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("max_bandwidth_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("max_bandwidth_policy")
			}
			return err
		}
	}

	return nil
}

func (m *MountNewCreateDisksParams) contextValidateMaxBandwidthUnit(ctx context.Context, formats strfmt.Registry) error {

	if m.MaxBandwidthUnit != nil {
		if err := m.MaxBandwidthUnit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("max_bandwidth_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("max_bandwidth_unit")
			}
			return err
		}
	}

	return nil
}

func (m *MountNewCreateDisksParams) contextValidateMaxIopsPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.MaxIopsPolicy != nil {
		if err := m.MaxIopsPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("max_iops_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("max_iops_policy")
			}
			return err
		}
	}

	return nil
}

func (m *MountNewCreateDisksParams) contextValidateVMVolume(ctx context.Context, formats strfmt.Registry) error {

	if m.VMVolume != nil {
		if err := m.VMVolume.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_volume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_volume")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MountNewCreateDisksParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MountNewCreateDisksParams) UnmarshalBinary(b []byte) error {
	var res MountNewCreateDisksParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MountNewCreateDisksParamsVMVolume mount new create disks params VM volume
//
// swagger:model MountNewCreateDisksParamsVMVolume
type MountNewCreateDisksParamsVMVolume struct {

	// elf storage policy
	// Required: true
	ElfStoragePolicy *VMVolumeElfStoragePolicyType `json:"elf_storage_policy"`

	// name
	// Required: true
	Name *string `json:"name"`

	// path
	Path *string `json:"path,omitempty"`

	// size
	// Required: true
	Size *int64 `json:"size"`

	// size unit
	SizeUnit *ByteUnit `json:"size_unit,omitempty"`

	MarshalOpts *MountNewCreateDisksParamsVMVolumeMarshalOpts `json:"-"`
}

type MountNewCreateDisksParamsVMVolumeMarshalOpts struct {
	ElfStoragePolicy_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	Path_Explicit_Null_When_Empty bool

	Size_Explicit_Null_When_Empty bool

	SizeUnit_Explicit_Null_When_Empty bool
}

func (m MountNewCreateDisksParamsVMVolume) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

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

	// handle nullable field name
	if m.Name != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name\":")
		bytes, err := swag.WriteJSON(m.Name)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Name_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name\":null")
		first = false
	}

	// handle nullable field path
	if m.Path != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"path\":")
		bytes, err := swag.WriteJSON(m.Path)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Path_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"path\":null")
		first = false
	}

	// handle nullable field size
	if m.Size != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"size\":")
		bytes, err := swag.WriteJSON(m.Size)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Size_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"size\":null")
		first = false
	}

	// handle nullable field size_unit
	if m.SizeUnit != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"size_unit\":")
		bytes, err := swag.WriteJSON(m.SizeUnit)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.SizeUnit_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"size_unit\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this mount new create disks params VM volume
func (m *MountNewCreateDisksParamsVMVolume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateElfStoragePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSizeUnit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MountNewCreateDisksParamsVMVolume) validateElfStoragePolicy(formats strfmt.Registry) error {

	if err := validate.Required("vm_volume"+"."+"elf_storage_policy", "body", m.ElfStoragePolicy); err != nil {
		return err
	}

	if err := validate.Required("vm_volume"+"."+"elf_storage_policy", "body", m.ElfStoragePolicy); err != nil {
		return err
	}

	if m.ElfStoragePolicy != nil {
		if err := m.ElfStoragePolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_volume" + "." + "elf_storage_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_volume" + "." + "elf_storage_policy")
			}
			return err
		}
	}

	return nil
}

func (m *MountNewCreateDisksParamsVMVolume) validateName(formats strfmt.Registry) error {

	if err := validate.Required("vm_volume"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *MountNewCreateDisksParamsVMVolume) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("vm_volume"+"."+"size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *MountNewCreateDisksParamsVMVolume) validateSizeUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.SizeUnit) { // not required
		return nil
	}

	if m.SizeUnit != nil {
		if err := m.SizeUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_volume" + "." + "size_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_volume" + "." + "size_unit")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this mount new create disks params VM volume based on the context it is used
func (m *MountNewCreateDisksParamsVMVolume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateElfStoragePolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSizeUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MountNewCreateDisksParamsVMVolume) contextValidateElfStoragePolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.ElfStoragePolicy != nil {
		if err := m.ElfStoragePolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_volume" + "." + "elf_storage_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_volume" + "." + "elf_storage_policy")
			}
			return err
		}
	}

	return nil
}

func (m *MountNewCreateDisksParamsVMVolume) contextValidateSizeUnit(ctx context.Context, formats strfmt.Registry) error {

	if m.SizeUnit != nil {
		if err := m.SizeUnit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_volume" + "." + "size_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_volume" + "." + "size_unit")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MountNewCreateDisksParamsVMVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MountNewCreateDisksParamsVMVolume) UnmarshalBinary(b []byte) error {
	var res MountNewCreateDisksParamsVMVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
