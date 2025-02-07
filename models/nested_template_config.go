// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NestedTemplateConfig nested template config
//
// swagger:model NestedTemplateConfig
type NestedTemplateConfig struct {

	// cpu exclusive expected enabled
	CPUExclusiveExpectedEnabled *bool `json:"cpu_exclusive_expected_enabled,omitempty"`

	// guest os type
	GuestOsType *string `json:"guest_os_type,omitempty"`

	// minimum replica
	MinimumReplica *bool `json:"minimum_replica,omitempty"`

	// sync vm time on resume
	SyncVMTimeOnResume *bool `json:"sync_vm_time_on_resume,omitempty"`

	MarshalOpts *NestedTemplateConfigMarshalOpts `json:"-"`
}

type NestedTemplateConfigMarshalOpts struct {
	CPUExclusiveExpectedEnabled_Explicit_Null_When_Empty bool

	GuestOsType_Explicit_Null_When_Empty bool

	MinimumReplica_Explicit_Null_When_Empty bool

	SyncVMTimeOnResume_Explicit_Null_When_Empty bool
}

func (m NestedTemplateConfig) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field cpu_exclusive_expected_enabled
	if m.CPUExclusiveExpectedEnabled != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cpu_exclusive_expected_enabled\":")
		bytes, err := swag.WriteJSON(m.CPUExclusiveExpectedEnabled)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.CPUExclusiveExpectedEnabled_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cpu_exclusive_expected_enabled\":null")
		first = false
	}

	// handle nullable field guest_os_type
	if m.GuestOsType != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"guest_os_type\":")
		bytes, err := swag.WriteJSON(m.GuestOsType)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.GuestOsType_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"guest_os_type\":null")
		first = false
	}

	// handle nullable field minimum_replica
	if m.MinimumReplica != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"minimum_replica\":")
		bytes, err := swag.WriteJSON(m.MinimumReplica)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.MinimumReplica_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"minimum_replica\":null")
		first = false
	}

	// handle nullable field sync_vm_time_on_resume
	if m.SyncVMTimeOnResume != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"sync_vm_time_on_resume\":")
		bytes, err := swag.WriteJSON(m.SyncVMTimeOnResume)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.SyncVMTimeOnResume_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"sync_vm_time_on_resume\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this nested template config
func (m *NestedTemplateConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this nested template config based on context it is used
func (m *NestedTemplateConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NestedTemplateConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedTemplateConfig) UnmarshalBinary(b []byte) error {
	var res NestedTemplateConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
