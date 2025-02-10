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

// VMTemplate Vm template
//
// swagger:model VmTemplate
type VMTemplate struct {

	// clock offset
	// Required: true
	ClockOffset *VMClockOffset `json:"clock_offset"`

	// cloud init supported
	// Required: true
	CloudInitSupported *bool `json:"cloud_init_supported"`

	// cluster
	// Required: true
	Cluster *NestedCluster `json:"cluster"`

	// content library vm template
	ContentLibraryVMTemplate *NestedContentLibraryVMTemplate `json:"content_library_vm_template,omitempty"`

	// cpu
	// Required: true
	CPU *NestedCPU `json:"cpu"`

	// cpu model
	// Required: true
	CPUModel *string `json:"cpu_model"`

	// description
	// Required: true
	Description *string `json:"description"`

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// firmware
	// Required: true
	Firmware *VMFirmware `json:"firmware"`

	// ha
	// Required: true
	Ha *bool `json:"ha"`

	// id
	// Required: true
	ID *string `json:"id"`

	// io policy
	IoPolicy *VMDiskIoPolicy `json:"io_policy,omitempty"`

	// labels
	Labels []*NestedLabel `json:"labels,omitempty"`

	// local created at
	LocalCreatedAt *string `json:"local_created_at,omitempty"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// max bandwidth
	MaxBandwidth *int64 `json:"max_bandwidth,omitempty"`

	// max bandwidth policy
	MaxBandwidthPolicy *VMDiskIoRestrictType `json:"max_bandwidth_policy,omitempty"`

	// max iops
	MaxIops *int32 `json:"max_iops,omitempty"`

	// max iops policy
	MaxIopsPolicy *VMDiskIoRestrictType `json:"max_iops_policy,omitempty"`

	// memory
	// Required: true
	Memory *int64 `json:"memory"`

	// name
	// Required: true
	Name *string `json:"name"`

	// size
	// Required: true
	Size *int64 `json:"size"`

	// vcpu
	// Required: true
	Vcpu *int32 `json:"vcpu"`

	// video type
	VideoType *string `json:"video_type,omitempty"`

	// vm disks
	VMDisks []*NestedFrozenDisks `json:"vm_disks,omitempty"`

	// vm nics
	VMNics []*NestedTemplateNic `json:"vm_nics,omitempty"`

	// win opt
	// Required: true
	WinOpt *bool `json:"win_opt"`

	MarshalOpts *VMTemplateMarshalOpts `json:"-"`
}

type VMTemplateMarshalOpts struct {
	ClockOffset_Explicit_Null_When_Empty bool

	CloudInitSupported_Explicit_Null_When_Empty bool

	Cluster_Explicit_Null_When_Empty bool

	ContentLibraryVMTemplate_Explicit_Null_When_Empty bool

	CPU_Explicit_Null_When_Empty bool

	CPUModel_Explicit_Null_When_Empty bool

	Description_Explicit_Null_When_Empty bool

	EntityAsyncStatus_Explicit_Null_When_Empty bool

	Firmware_Explicit_Null_When_Empty bool

	Ha_Explicit_Null_When_Empty bool

	ID_Explicit_Null_When_Empty bool

	IoPolicy_Explicit_Null_When_Empty bool

	Labels_Explicit_Null_When_Empty bool

	LocalCreatedAt_Explicit_Null_When_Empty bool

	LocalID_Explicit_Null_When_Empty bool

	MaxBandwidth_Explicit_Null_When_Empty bool

	MaxBandwidthPolicy_Explicit_Null_When_Empty bool

	MaxIops_Explicit_Null_When_Empty bool

	MaxIopsPolicy_Explicit_Null_When_Empty bool

	Memory_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	Size_Explicit_Null_When_Empty bool

	Vcpu_Explicit_Null_When_Empty bool

	VideoType_Explicit_Null_When_Empty bool

	VMDisks_Explicit_Null_When_Empty bool

	VMNics_Explicit_Null_When_Empty bool

	WinOpt_Explicit_Null_When_Empty bool
}

func (m VMTemplate) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field clock_offset
	if m.ClockOffset != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"clock_offset\":")
		bytes, err := swag.WriteJSON(m.ClockOffset)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ClockOffset_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"clock_offset\":null")
		first = false
	}

	// handle nullable field cloud_init_supported
	if m.CloudInitSupported != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cloud_init_supported\":")
		bytes, err := swag.WriteJSON(m.CloudInitSupported)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.CloudInitSupported_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cloud_init_supported\":null")
		first = false
	}

	// handle nullable field cluster
	if m.Cluster != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster\":")
		bytes, err := swag.WriteJSON(m.Cluster)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Cluster_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster\":null")
		first = false
	}

	// handle nullable field content_library_vm_template
	if m.ContentLibraryVMTemplate != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"content_library_vm_template\":")
		bytes, err := swag.WriteJSON(m.ContentLibraryVMTemplate)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ContentLibraryVMTemplate_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"content_library_vm_template\":null")
		first = false
	}

	// handle nullable field cpu
	if m.CPU != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cpu\":")
		bytes, err := swag.WriteJSON(m.CPU)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.CPU_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cpu\":null")
		first = false
	}

	// handle nullable field cpu_model
	if m.CPUModel != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cpu_model\":")
		bytes, err := swag.WriteJSON(m.CPUModel)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.CPUModel_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cpu_model\":null")
		first = false
	}

	// handle nullable field description
	if m.Description != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description\":")
		bytes, err := swag.WriteJSON(m.Description)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Description_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description\":null")
		first = false
	}

	// handle nullable field entityAsyncStatus
	if m.EntityAsyncStatus != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"entityAsyncStatus\":")
		bytes, err := swag.WriteJSON(m.EntityAsyncStatus)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.EntityAsyncStatus_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"entityAsyncStatus\":null")
		first = false
	}

	// handle nullable field firmware
	if m.Firmware != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"firmware\":")
		bytes, err := swag.WriteJSON(m.Firmware)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Firmware_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"firmware\":null")
		first = false
	}

	// handle nullable field ha
	if m.Ha != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ha\":")
		bytes, err := swag.WriteJSON(m.Ha)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Ha_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ha\":null")
		first = false
	}

	// handle nullable field id
	if m.ID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id\":")
		bytes, err := swag.WriteJSON(m.ID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id\":null")
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

	// handle non nullable field labels with omitempty
	if !swag.IsZero(m.Labels) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"labels\":")
		bytes, err := swag.WriteJSON(m.Labels)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field local_created_at
	if m.LocalCreatedAt != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_created_at\":")
		bytes, err := swag.WriteJSON(m.LocalCreatedAt)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.LocalCreatedAt_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_created_at\":null")
		first = false
	}

	// handle nullable field local_id
	if m.LocalID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id\":")
		bytes, err := swag.WriteJSON(m.LocalID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.LocalID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id\":null")
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

	// handle nullable field memory
	if m.Memory != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"memory\":")
		bytes, err := swag.WriteJSON(m.Memory)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Memory_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"memory\":null")
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

	// handle nullable field vcpu
	if m.Vcpu != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vcpu\":")
		bytes, err := swag.WriteJSON(m.Vcpu)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Vcpu_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vcpu\":null")
		first = false
	}

	// handle nullable field video_type
	if m.VideoType != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"video_type\":")
		bytes, err := swag.WriteJSON(m.VideoType)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VideoType_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"video_type\":null")
		first = false
	}

	// handle non nullable field vm_disks with omitempty
	if !swag.IsZero(m.VMDisks) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_disks\":")
		bytes, err := swag.WriteJSON(m.VMDisks)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle non nullable field vm_nics with omitempty
	if !swag.IsZero(m.VMNics) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_nics\":")
		bytes, err := swag.WriteJSON(m.VMNics)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field win_opt
	if m.WinOpt != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"win_opt\":")
		bytes, err := swag.WriteJSON(m.WinOpt)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.WinOpt_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"win_opt\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this Vm template
func (m *VMTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClockOffset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudInitSupported(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContentLibraryVMTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPU(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPUModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirmware(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHa(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIoPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxBandwidthPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxIopsPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcpu(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMNics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWinOpt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMTemplate) validateClockOffset(formats strfmt.Registry) error {

	if err := validate.Required("clock_offset", "body", m.ClockOffset); err != nil {
		return err
	}

	if err := validate.Required("clock_offset", "body", m.ClockOffset); err != nil {
		return err
	}

	if m.ClockOffset != nil {
		if err := m.ClockOffset.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clock_offset")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clock_offset")
			}
			return err
		}
	}

	return nil
}

func (m *VMTemplate) validateCloudInitSupported(formats strfmt.Registry) error {

	if err := validate.Required("cloud_init_supported", "body", m.CloudInitSupported); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplate) validateCluster(formats strfmt.Registry) error {

	if err := validate.Required("cluster", "body", m.Cluster); err != nil {
		return err
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *VMTemplate) validateContentLibraryVMTemplate(formats strfmt.Registry) error {
	if swag.IsZero(m.ContentLibraryVMTemplate) { // not required
		return nil
	}

	if m.ContentLibraryVMTemplate != nil {
		if err := m.ContentLibraryVMTemplate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content_library_vm_template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content_library_vm_template")
			}
			return err
		}
	}

	return nil
}

func (m *VMTemplate) validateCPU(formats strfmt.Registry) error {

	if err := validate.Required("cpu", "body", m.CPU); err != nil {
		return err
	}

	if m.CPU != nil {
		if err := m.CPU.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cpu")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cpu")
			}
			return err
		}
	}

	return nil
}

func (m *VMTemplate) validateCPUModel(formats strfmt.Registry) error {

	if err := validate.Required("cpu_model", "body", m.CPUModel); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplate) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplate) validateEntityAsyncStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatus) { // not required
		return nil
	}

	if m.EntityAsyncStatus != nil {
		if err := m.EntityAsyncStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus")
			}
			return err
		}
	}

	return nil
}

func (m *VMTemplate) validateFirmware(formats strfmt.Registry) error {

	if err := validate.Required("firmware", "body", m.Firmware); err != nil {
		return err
	}

	if err := validate.Required("firmware", "body", m.Firmware); err != nil {
		return err
	}

	if m.Firmware != nil {
		if err := m.Firmware.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("firmware")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("firmware")
			}
			return err
		}
	}

	return nil
}

func (m *VMTemplate) validateHa(formats strfmt.Registry) error {

	if err := validate.Required("ha", "body", m.Ha); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplate) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplate) validateIoPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.IoPolicy) { // not required
		return nil
	}

	if m.IoPolicy != nil {
		if err := m.IoPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("io_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("io_policy")
			}
			return err
		}
	}

	return nil
}

func (m *VMTemplate) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	for i := 0; i < len(m.Labels); i++ {
		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if m.Labels[i] != nil {
			if err := m.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMTemplate) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplate) validateMaxBandwidthPolicy(formats strfmt.Registry) error {
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

func (m *VMTemplate) validateMaxIopsPolicy(formats strfmt.Registry) error {
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

func (m *VMTemplate) validateMemory(formats strfmt.Registry) error {

	if err := validate.Required("memory", "body", m.Memory); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplate) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplate) validateVcpu(formats strfmt.Registry) error {

	if err := validate.Required("vcpu", "body", m.Vcpu); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplate) validateVMDisks(formats strfmt.Registry) error {
	if swag.IsZero(m.VMDisks) { // not required
		return nil
	}

	for i := 0; i < len(m.VMDisks); i++ {
		if swag.IsZero(m.VMDisks[i]) { // not required
			continue
		}

		if m.VMDisks[i] != nil {
			if err := m.VMDisks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vm_disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMTemplate) validateVMNics(formats strfmt.Registry) error {
	if swag.IsZero(m.VMNics) { // not required
		return nil
	}

	for i := 0; i < len(m.VMNics); i++ {
		if swag.IsZero(m.VMNics[i]) { // not required
			continue
		}

		if m.VMNics[i] != nil {
			if err := m.VMNics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_nics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vm_nics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMTemplate) validateWinOpt(formats strfmt.Registry) error {

	if err := validate.Required("win_opt", "body", m.WinOpt); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this Vm template based on the context it is used
func (m *VMTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClockOffset(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContentLibraryVMTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCPU(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFirmware(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIoPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxBandwidthPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxIopsPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMNics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMTemplate) contextValidateClockOffset(ctx context.Context, formats strfmt.Registry) error {

	if m.ClockOffset != nil {
		if err := m.ClockOffset.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clock_offset")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clock_offset")
			}
			return err
		}
	}

	return nil
}

func (m *VMTemplate) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {
		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *VMTemplate) contextValidateContentLibraryVMTemplate(ctx context.Context, formats strfmt.Registry) error {

	if m.ContentLibraryVMTemplate != nil {
		if err := m.ContentLibraryVMTemplate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content_library_vm_template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content_library_vm_template")
			}
			return err
		}
	}

	return nil
}

func (m *VMTemplate) contextValidateCPU(ctx context.Context, formats strfmt.Registry) error {

	if m.CPU != nil {
		if err := m.CPU.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cpu")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cpu")
			}
			return err
		}
	}

	return nil
}

func (m *VMTemplate) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.EntityAsyncStatus != nil {
		if err := m.EntityAsyncStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus")
			}
			return err
		}
	}

	return nil
}

func (m *VMTemplate) contextValidateFirmware(ctx context.Context, formats strfmt.Registry) error {

	if m.Firmware != nil {
		if err := m.Firmware.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("firmware")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("firmware")
			}
			return err
		}
	}

	return nil
}

func (m *VMTemplate) contextValidateIoPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.IoPolicy != nil {
		if err := m.IoPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("io_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("io_policy")
			}
			return err
		}
	}

	return nil
}

func (m *VMTemplate) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Labels); i++ {

		if m.Labels[i] != nil {
			if err := m.Labels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMTemplate) contextValidateMaxBandwidthPolicy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMTemplate) contextValidateMaxIopsPolicy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMTemplate) contextValidateVMDisks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VMDisks); i++ {

		if m.VMDisks[i] != nil {
			if err := m.VMDisks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vm_disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMTemplate) contextValidateVMNics(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VMNics); i++ {

		if m.VMNics[i] != nil {
			if err := m.VMNics[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_nics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vm_nics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMTemplate) UnmarshalBinary(b []byte) error {
	var res VMTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
