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

// NestedFrozenDisks nested frozen disks
//
// swagger:model NestedFrozenDisks
type NestedFrozenDisks struct {

	// boot
	// Required: true
	Boot *int64 `json:"boot"`

	// bus
	// Required: true
	Bus *Bus `json:"bus"`

	// disabled
	Disabled *bool `json:"disabled,omitempty"`

	// disk name
	DiskName *string `json:"disk_name,omitempty"`

	// elf image local id
	// Required: true
	ElfImageLocalID *string `json:"elf_image_local_id"`

	// image name
	ImageName *string `json:"image_name,omitempty"`

	// index
	// Required: true
	Index *int32 `json:"index"`

	// key
	Key *int32 `json:"key,omitempty"`

	// max bandwidth
	MaxBandwidth *int64 `json:"max_bandwidth,omitempty"`

	// max bandwidth policy
	MaxBandwidthPolicy *VMDiskIoRestrictType `json:"max_bandwidth_policy,omitempty"`

	// max iops
	MaxIops *int32 `json:"max_iops,omitempty"`

	// max iops policy
	MaxIopsPolicy *VMDiskIoRestrictType `json:"max_iops_policy,omitempty"`

	// path
	// Required: true
	Path *string `json:"path"`

	// size
	// Required: true
	Size *int64 `json:"size"`

	// snapshot local id
	SnapshotLocalID *string `json:"snapshot_local_id,omitempty"`

	// storage policy uuid
	// Required: true
	StoragePolicyUUID *string `json:"storage_policy_uuid"`

	// svt image local id
	// Required: true
	SvtImageLocalID *string `json:"svt_image_local_id"`

	// type
	// Required: true
	Type *VMDiskType `json:"type"`

	// vm volume local id
	// Required: true
	VMVolumeLocalID *string `json:"vm_volume_local_id"`

	// vm volume snapshot uuid
	VMVolumeSnapshotUUID *string `json:"vm_volume_snapshot_uuid,omitempty"`

	// vm volume template uuid
	VMVolumeTemplateUUID *string `json:"vm_volume_template_uuid,omitempty"`
}

// Validate validates this nested frozen disks
func (m *NestedFrozenDisks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBoot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElfImageLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxBandwidthPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxIopsPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoragePolicyUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvtImageLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMVolumeLocalID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedFrozenDisks) validateBoot(formats strfmt.Registry) error {

	if err := validate.Required("boot", "body", m.Boot); err != nil {
		return err
	}

	return nil
}

func (m *NestedFrozenDisks) validateBus(formats strfmt.Registry) error {

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

func (m *NestedFrozenDisks) validateElfImageLocalID(formats strfmt.Registry) error {

	if err := validate.Required("elf_image_local_id", "body", m.ElfImageLocalID); err != nil {
		return err
	}

	return nil
}

func (m *NestedFrozenDisks) validateIndex(formats strfmt.Registry) error {

	if err := validate.Required("index", "body", m.Index); err != nil {
		return err
	}

	return nil
}

func (m *NestedFrozenDisks) validateMaxBandwidthPolicy(formats strfmt.Registry) error {
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

func (m *NestedFrozenDisks) validateMaxIopsPolicy(formats strfmt.Registry) error {
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

func (m *NestedFrozenDisks) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

func (m *NestedFrozenDisks) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *NestedFrozenDisks) validateStoragePolicyUUID(formats strfmt.Registry) error {

	if err := validate.Required("storage_policy_uuid", "body", m.StoragePolicyUUID); err != nil {
		return err
	}

	return nil
}

func (m *NestedFrozenDisks) validateSvtImageLocalID(formats strfmt.Registry) error {

	if err := validate.Required("svt_image_local_id", "body", m.SvtImageLocalID); err != nil {
		return err
	}

	return nil
}

func (m *NestedFrozenDisks) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *NestedFrozenDisks) validateVMVolumeLocalID(formats strfmt.Registry) error {

	if err := validate.Required("vm_volume_local_id", "body", m.VMVolumeLocalID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this nested frozen disks based on the context it is used
func (m *NestedFrozenDisks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxBandwidthPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxIopsPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedFrozenDisks) contextValidateBus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *NestedFrozenDisks) contextValidateMaxBandwidthPolicy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *NestedFrozenDisks) contextValidateMaxIopsPolicy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *NestedFrozenDisks) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NestedFrozenDisks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedFrozenDisks) UnmarshalBinary(b []byte) error {
	var res NestedFrozenDisks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
