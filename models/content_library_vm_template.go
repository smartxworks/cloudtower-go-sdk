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

// ContentLibraryVMTemplate content library Vm template
//
// swagger:model ContentLibraryVmTemplate
type ContentLibraryVMTemplate struct {

	// architecture
	// Required: true
	Architecture *Architecture `json:"architecture"`

	// clock offset
	ClockOffset *VMClockOffset `json:"clock_offset,omitempty"`

	// cloud init supported
	// Required: true
	CloudInitSupported *bool `json:"cloud_init_supported"`

	// clusters
	Clusters []*NestedCluster `json:"clusters,omitempty"`

	// cpu
	CPU *NestedCPU `json:"cpu,omitempty"`

	// cpu model
	CPUModel *string `json:"cpu_model,omitempty"`

	// created at
	// Required: true
	CreatedAt *string `json:"createdAt"`

	// description
	// Required: true
	Description *string `json:"description"`

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// firmware
	Firmware *VMFirmware `json:"firmware,omitempty"`

	// ha
	Ha *bool `json:"ha,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// io policy
	IoPolicy *VMDiskIoPolicy `json:"io_policy,omitempty"`

	// labels
	Labels []*NestedLabel `json:"labels,omitempty"`

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

	// os
	Os *string `json:"os,omitempty"`

	// size
	// Required: true
	Size *int64 `json:"size"`

	// template config
	TemplateConfig *NestedTemplateConfig `json:"template_config,omitempty"`

	// vcpu
	// Required: true
	Vcpu *int32 `json:"vcpu"`

	// video type
	VideoType *string `json:"video_type,omitempty"`

	// vm disks
	VMDisks []*NestedContentLibraryVMTemplateDisk `json:"vm_disks,omitempty"`

	// vm nics
	VMNics []*NestedContentLibraryVMTemplateNic `json:"vm_nics,omitempty"`

	// vm template uuids
	// Required: true
	VMTemplateUuids []string `json:"vm_template_uuids"`

	// vm templates
	VMTemplates []*NestedVMTemplate `json:"vm_templates,omitempty"`

	// win opt
	WinOpt *bool `json:"win_opt,omitempty"`

	// zbs storage info
	ZbsStorageInfo []*NestedZbsStorageInfo `json:"zbs_storage_info,omitempty"`

	MarshalOpts *ContentLibraryVMTemplateMarshalOpts `json:"-"`
}

type ContentLibraryVMTemplateMarshalOpts struct {
	Architecture_Explicit_Null_When_Empty bool

	ClockOffset_Explicit_Null_When_Empty bool

	CloudInitSupported_Explicit_Null_When_Empty bool

	Clusters_Explicit_Null_When_Empty bool

	CPU_Explicit_Null_When_Empty bool

	CPUModel_Explicit_Null_When_Empty bool

	CreatedAt_Explicit_Null_When_Empty bool

	Description_Explicit_Null_When_Empty bool

	EntityAsyncStatus_Explicit_Null_When_Empty bool

	Firmware_Explicit_Null_When_Empty bool

	Ha_Explicit_Null_When_Empty bool

	ID_Explicit_Null_When_Empty bool

	IoPolicy_Explicit_Null_When_Empty bool

	Labels_Explicit_Null_When_Empty bool

	MaxBandwidth_Explicit_Null_When_Empty bool

	MaxBandwidthPolicy_Explicit_Null_When_Empty bool

	MaxIops_Explicit_Null_When_Empty bool

	MaxIopsPolicy_Explicit_Null_When_Empty bool

	Memory_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	Os_Explicit_Null_When_Empty bool

	Size_Explicit_Null_When_Empty bool

	TemplateConfig_Explicit_Null_When_Empty bool

	Vcpu_Explicit_Null_When_Empty bool

	VideoType_Explicit_Null_When_Empty bool

	VMDisks_Explicit_Null_When_Empty bool

	VMNics_Explicit_Null_When_Empty bool

	VMTemplateUuids_Explicit_Null_When_Empty bool

	VMTemplates_Explicit_Null_When_Empty bool

	WinOpt_Explicit_Null_When_Empty bool

	ZbsStorageInfo_Explicit_Null_When_Empty bool
}

func (m ContentLibraryVMTemplate) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field architecture
	if m.Architecture != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"architecture\":")
		bytes, err := swag.WriteJSON(m.Architecture)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Architecture_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"architecture\":null")
		first = false
	}

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

	// handle non nullable field clusters with omitempty
	if swag.IsZero(m.Clusters) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"clusters\":")
		bytes, err := swag.WriteJSON(m.Clusters)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
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

	// handle nullable field createdAt
	if m.CreatedAt != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"createdAt\":")
		bytes, err := swag.WriteJSON(m.CreatedAt)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.CreatedAt_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"createdAt\":null")
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
	if swag.IsZero(m.Labels) {
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

	// handle nullable field os
	if m.Os != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"os\":")
		bytes, err := swag.WriteJSON(m.Os)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Os_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"os\":null")
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

	// handle nullable field template_config
	if m.TemplateConfig != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"template_config\":")
		bytes, err := swag.WriteJSON(m.TemplateConfig)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.TemplateConfig_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"template_config\":null")
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
	if swag.IsZero(m.VMDisks) {
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
	if swag.IsZero(m.VMNics) {
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

	// handle non nullable field vm_template_uuids without omitempty
	{
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_template_uuids\":")
		bytes, err := swag.WriteJSON(m.VMTemplateUuids)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle non nullable field vm_templates with omitempty
	if swag.IsZero(m.VMTemplates) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_templates\":")
		bytes, err := swag.WriteJSON(m.VMTemplates)
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

	// handle non nullable field zbs_storage_info with omitempty
	if swag.IsZero(m.ZbsStorageInfo) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"zbs_storage_info\":")
		bytes, err := swag.WriteJSON(m.ZbsStorageInfo)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this content library Vm template
func (m *ContentLibraryVMTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArchitecture(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClockOffset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudInitSupported(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPU(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
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

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIoPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
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

	if err := m.validateTemplateConfig(formats); err != nil {
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

	if err := m.validateVMTemplateUuids(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMTemplates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZbsStorageInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentLibraryVMTemplate) validateArchitecture(formats strfmt.Registry) error {

	if err := validate.Required("architecture", "body", m.Architecture); err != nil {
		return err
	}

	if err := validate.Required("architecture", "body", m.Architecture); err != nil {
		return err
	}

	if m.Architecture != nil {
		if err := m.Architecture.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("architecture")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("architecture")
			}
			return err
		}
	}

	return nil
}

func (m *ContentLibraryVMTemplate) validateClockOffset(formats strfmt.Registry) error {
	if swag.IsZero(m.ClockOffset) { // not required
		return nil
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

func (m *ContentLibraryVMTemplate) validateCloudInitSupported(formats strfmt.Registry) error {

	if err := validate.Required("cloud_init_supported", "body", m.CloudInitSupported); err != nil {
		return err
	}

	return nil
}

func (m *ContentLibraryVMTemplate) validateClusters(formats strfmt.Registry) error {
	if swag.IsZero(m.Clusters) { // not required
		return nil
	}

	for i := 0; i < len(m.Clusters); i++ {
		if swag.IsZero(m.Clusters[i]) { // not required
			continue
		}

		if m.Clusters[i] != nil {
			if err := m.Clusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContentLibraryVMTemplate) validateCPU(formats strfmt.Registry) error {
	if swag.IsZero(m.CPU) { // not required
		return nil
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

func (m *ContentLibraryVMTemplate) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *ContentLibraryVMTemplate) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *ContentLibraryVMTemplate) validateEntityAsyncStatus(formats strfmt.Registry) error {
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

func (m *ContentLibraryVMTemplate) validateFirmware(formats strfmt.Registry) error {
	if swag.IsZero(m.Firmware) { // not required
		return nil
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

func (m *ContentLibraryVMTemplate) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ContentLibraryVMTemplate) validateIoPolicy(formats strfmt.Registry) error {
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

func (m *ContentLibraryVMTemplate) validateLabels(formats strfmt.Registry) error {
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

func (m *ContentLibraryVMTemplate) validateMaxBandwidthPolicy(formats strfmt.Registry) error {
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

func (m *ContentLibraryVMTemplate) validateMaxIopsPolicy(formats strfmt.Registry) error {
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

func (m *ContentLibraryVMTemplate) validateMemory(formats strfmt.Registry) error {

	if err := validate.Required("memory", "body", m.Memory); err != nil {
		return err
	}

	return nil
}

func (m *ContentLibraryVMTemplate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ContentLibraryVMTemplate) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *ContentLibraryVMTemplate) validateTemplateConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.TemplateConfig) { // not required
		return nil
	}

	if m.TemplateConfig != nil {
		if err := m.TemplateConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("template_config")
			}
			return err
		}
	}

	return nil
}

func (m *ContentLibraryVMTemplate) validateVcpu(formats strfmt.Registry) error {

	if err := validate.Required("vcpu", "body", m.Vcpu); err != nil {
		return err
	}

	return nil
}

func (m *ContentLibraryVMTemplate) validateVMDisks(formats strfmt.Registry) error {
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

func (m *ContentLibraryVMTemplate) validateVMNics(formats strfmt.Registry) error {
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

func (m *ContentLibraryVMTemplate) validateVMTemplateUuids(formats strfmt.Registry) error {

	if err := validate.Required("vm_template_uuids", "body", m.VMTemplateUuids); err != nil {
		return err
	}

	return nil
}

func (m *ContentLibraryVMTemplate) validateVMTemplates(formats strfmt.Registry) error {
	if swag.IsZero(m.VMTemplates) { // not required
		return nil
	}

	for i := 0; i < len(m.VMTemplates); i++ {
		if swag.IsZero(m.VMTemplates[i]) { // not required
			continue
		}

		if m.VMTemplates[i] != nil {
			if err := m.VMTemplates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_templates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vm_templates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContentLibraryVMTemplate) validateZbsStorageInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.ZbsStorageInfo) { // not required
		return nil
	}

	for i := 0; i < len(m.ZbsStorageInfo); i++ {
		if swag.IsZero(m.ZbsStorageInfo[i]) { // not required
			continue
		}

		if m.ZbsStorageInfo[i] != nil {
			if err := m.ZbsStorageInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("zbs_storage_info" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("zbs_storage_info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this content library Vm template based on the context it is used
func (m *ContentLibraryVMTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateArchitecture(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClockOffset(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClusters(ctx, formats); err != nil {
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

	if err := m.contextValidateTemplateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMNics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMTemplates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateZbsStorageInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentLibraryVMTemplate) contextValidateArchitecture(ctx context.Context, formats strfmt.Registry) error {

	if m.Architecture != nil {
		if err := m.Architecture.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("architecture")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("architecture")
			}
			return err
		}
	}

	return nil
}

func (m *ContentLibraryVMTemplate) contextValidateClockOffset(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ContentLibraryVMTemplate) contextValidateClusters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Clusters); i++ {

		if m.Clusters[i] != nil {
			if err := m.Clusters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContentLibraryVMTemplate) contextValidateCPU(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ContentLibraryVMTemplate) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ContentLibraryVMTemplate) contextValidateFirmware(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ContentLibraryVMTemplate) contextValidateIoPolicy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ContentLibraryVMTemplate) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ContentLibraryVMTemplate) contextValidateMaxBandwidthPolicy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ContentLibraryVMTemplate) contextValidateMaxIopsPolicy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ContentLibraryVMTemplate) contextValidateTemplateConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.TemplateConfig != nil {
		if err := m.TemplateConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("template_config")
			}
			return err
		}
	}

	return nil
}

func (m *ContentLibraryVMTemplate) contextValidateVMDisks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ContentLibraryVMTemplate) contextValidateVMNics(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ContentLibraryVMTemplate) contextValidateVMTemplates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VMTemplates); i++ {

		if m.VMTemplates[i] != nil {
			if err := m.VMTemplates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_templates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vm_templates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContentLibraryVMTemplate) contextValidateZbsStorageInfo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ZbsStorageInfo); i++ {

		if m.ZbsStorageInfo[i] != nil {
			if err := m.ZbsStorageInfo[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("zbs_storage_info" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("zbs_storage_info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContentLibraryVMTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentLibraryVMTemplate) UnmarshalBinary(b []byte) error {
	var res ContentLibraryVMTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
