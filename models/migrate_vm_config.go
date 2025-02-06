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

// MigrateVMConfig migrate Vm config
//
// swagger:model MigrateVmConfig
type MigrateVMConfig struct {

	// delete src vm
	DeleteSrcVM *bool `json:"delete_src_vm,omitempty"`

	// elf storage policy
	// Required: true
	ElfStoragePolicy *VMVolumeElfStoragePolicyType `json:"elf_storage_policy"`

	// migrate type
	// Required: true
	MigrateType *MigrateType `json:"migrate_type"`

	// network mapping
	// Required: true
	NetworkMapping []*VlanMapping `json:"network_mapping"`

	// new name
	NewName *string `json:"new_name,omitempty"`

	// remove unmovable devices
	RemoveUnmovableDevices *bool `json:"remove_unmovable_devices,omitempty"`

	MarshalOpts *MigrateVMConfigMarshalOpts `json:"-"`
}

type MigrateVMConfigMarshalOpts struct {
	DeleteSrcVM_Explicit_Null_When_Empty bool

	ElfStoragePolicy_Explicit_Null_When_Empty bool

	MigrateType_Explicit_Null_When_Empty bool

	NetworkMapping_Explicit_Null_When_Empty bool

	NewName_Explicit_Null_When_Empty bool

	RemoveUnmovableDevices_Explicit_Null_When_Empty bool
}

func (m MigrateVMConfig) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field delete_src_vm
	if m.DeleteSrcVM != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"delete_src_vm\":")
		bytes, err := swag.WriteJSON(m.DeleteSrcVM)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.DeleteSrcVM_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"delete_src_vm\":null")
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

	// handle nullable field migrate_type
	if m.MigrateType != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"migrate_type\":")
		bytes, err := swag.WriteJSON(m.MigrateType)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.MigrateType_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"migrate_type\":null")
		first = false
	}

	// handle non nullable field network_mapping without omitempty
	if !first {
		b.WriteString(",")
	}
	b.WriteString("\"network_mapping\":")
	{
		bytes, err := swag.WriteJSON(m.NetworkMapping)
		if err != nil {
			return nil, err
		}
	}
	b.Write(bytes)
	first = false

	// handle nullable field new_name
	if m.NewName != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"new_name\":")
		bytes, err := swag.WriteJSON(m.NewName)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NewName_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"new_name\":null")
		first = false
	}

	// handle nullable field remove_unmovable_devices
	if m.RemoveUnmovableDevices != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"remove_unmovable_devices\":")
		bytes, err := swag.WriteJSON(m.RemoveUnmovableDevices)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.RemoveUnmovableDevices_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"remove_unmovable_devices\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this migrate Vm config
func (m *MigrateVMConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateElfStoragePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMigrateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkMapping(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MigrateVMConfig) validateElfStoragePolicy(formats strfmt.Registry) error {

	if err := validate.Required("elf_storage_policy", "body", m.ElfStoragePolicy); err != nil {
		return err
	}

	if err := validate.Required("elf_storage_policy", "body", m.ElfStoragePolicy); err != nil {
		return err
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

func (m *MigrateVMConfig) validateMigrateType(formats strfmt.Registry) error {

	if err := validate.Required("migrate_type", "body", m.MigrateType); err != nil {
		return err
	}

	if err := validate.Required("migrate_type", "body", m.MigrateType); err != nil {
		return err
	}

	if m.MigrateType != nil {
		if err := m.MigrateType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("migrate_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("migrate_type")
			}
			return err
		}
	}

	return nil
}

func (m *MigrateVMConfig) validateNetworkMapping(formats strfmt.Registry) error {

	if err := validate.Required("network_mapping", "body", m.NetworkMapping); err != nil {
		return err
	}

	for i := 0; i < len(m.NetworkMapping); i++ {
		if swag.IsZero(m.NetworkMapping[i]) { // not required
			continue
		}

		if m.NetworkMapping[i] != nil {
			if err := m.NetworkMapping[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("network_mapping" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("network_mapping" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this migrate Vm config based on the context it is used
func (m *MigrateVMConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateElfStoragePolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMigrateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkMapping(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MigrateVMConfig) contextValidateElfStoragePolicy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *MigrateVMConfig) contextValidateMigrateType(ctx context.Context, formats strfmt.Registry) error {

	if m.MigrateType != nil {
		if err := m.MigrateType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("migrate_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("migrate_type")
			}
			return err
		}
	}

	return nil
}

func (m *MigrateVMConfig) contextValidateNetworkMapping(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NetworkMapping); i++ {

		if m.NetworkMapping[i] != nil {
			if err := m.NetworkMapping[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("network_mapping" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("network_mapping" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MigrateVMConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MigrateVMConfig) UnmarshalBinary(b []byte) error {
	var res MigrateVMConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
