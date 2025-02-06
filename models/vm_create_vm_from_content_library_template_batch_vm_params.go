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

// VMCreateVMFromContentLibraryTemplateBatchVMParams Vm create Vm from content library template batch Vm params
//
// swagger:model VmCreateVmFromContentLibraryTemplateBatchVmParams
type VMCreateVMFromContentLibraryTemplateBatchVMParams struct {

	// cloud init
	CloudInit *TemplateCloudInit `json:"cloud_init,omitempty"`

	// cluster id
	// Required: true
	ClusterID *string `json:"cluster_id"`

	// cpu cores
	CPUCores *int32 `json:"cpu_cores,omitempty"`

	// cpu sockets
	CPUSockets *int32 `json:"cpu_sockets,omitempty"`

	// description
	Description *string `json:"description,omitempty"`

	// disk operate
	DiskOperate *VMDiskOperate `json:"disk_operate,omitempty"`

	// firmware
	Firmware *VMFirmware `json:"firmware,omitempty"`

	// folder id
	FolderID *string `json:"folder_id,omitempty"`

	// gpu devices
	GpuDevices []*VMGpuOperationParams `json:"gpu_devices,omitempty"`

	// guest os type
	GuestOsType *VMGuestsOperationSystem `json:"guest_os_type,omitempty"`

	// ha
	Ha *bool `json:"ha,omitempty"`

	// host id
	HostID *string `json:"host_id,omitempty"`

	// is full copy
	// Required: true
	IsFullCopy *bool `json:"is_full_copy"`

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

	// memory
	Memory *int64 `json:"memory,omitempty"`

	// memory unit
	MemoryUnit *ByteUnit `json:"memory_unit,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// owner
	Owner *VMOwnerParams `json:"owner,omitempty"`

	// pci nics
	PciNics *NicWhereInput `json:"pci_nics,omitempty"`

	// status
	Status *VMStatus `json:"status,omitempty"`

	// vcpu
	Vcpu *int32 `json:"vcpu,omitempty"`

	// vm nics
	VMNics []*VMNicParams `json:"vm_nics,omitempty"`

	// vm placement group
	VMPlacementGroup *VMPlacementGroupWhereInput `json:"vm_placement_group,omitempty"`

	MarshalOpts *VMCreateVMFromContentLibraryTemplateBatchVMParamsMarshalOpts `json:"-"`
}

type VMCreateVMFromContentLibraryTemplateBatchVMParamsMarshalOpts struct {
	CloudInit_Explicit_Null_When_Empty bool

	ClusterID_Explicit_Null_When_Empty bool

	CPUCores_Explicit_Null_When_Empty bool

	CPUSockets_Explicit_Null_When_Empty bool

	Description_Explicit_Null_When_Empty bool

	DiskOperate_Explicit_Null_When_Empty bool

	Firmware_Explicit_Null_When_Empty bool

	FolderID_Explicit_Null_When_Empty bool

	GuestOsType_Explicit_Null_When_Empty bool

	Ha_Explicit_Null_When_Empty bool

	HostID_Explicit_Null_When_Empty bool

	IsFullCopy_Explicit_Null_When_Empty bool

	MaxBandwidth_Explicit_Null_When_Empty bool

	MaxBandwidthPolicy_Explicit_Null_When_Empty bool

	MaxBandwidthUnit_Explicit_Null_When_Empty bool

	MaxIops_Explicit_Null_When_Empty bool

	MaxIopsPolicy_Explicit_Null_When_Empty bool

	Memory_Explicit_Null_When_Empty bool

	MemoryUnit_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	Owner_Explicit_Null_When_Empty bool

	PciNics_Explicit_Null_When_Empty bool

	Status_Explicit_Null_When_Empty bool

	Vcpu_Explicit_Null_When_Empty bool

	VMPlacementGroup_Explicit_Null_When_Empty bool
}

func (m VMCreateVMFromContentLibraryTemplateBatchVMParams) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field cloud_init
	if m.CloudInit != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cloud_init\":")
		bytes, err := swag.WriteJSON(m.CloudInit)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.CloudInit_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cloud_init\":null")
		first = false
	}

	// handle nullable field cluster_id
	if m.ClusterID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster_id\":")
		bytes, err := swag.WriteJSON(m.ClusterID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ClusterID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster_id\":null")
		first = false
	}

	// handle nullable field cpu_cores
	if m.CPUCores != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cpu_cores\":")
		bytes, err := swag.WriteJSON(m.CPUCores)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.CPUCores_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cpu_cores\":null")
		first = false
	}

	// handle nullable field cpu_sockets
	if m.CPUSockets != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cpu_sockets\":")
		bytes, err := swag.WriteJSON(m.CPUSockets)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.CPUSockets_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cpu_sockets\":null")
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

	// handle nullable field disk_operate
	if m.DiskOperate != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"disk_operate\":")
		bytes, err := swag.WriteJSON(m.DiskOperate)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.DiskOperate_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"disk_operate\":null")
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

	// handle nullable field folder_id
	if m.FolderID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"folder_id\":")
		bytes, err := swag.WriteJSON(m.FolderID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.FolderID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"folder_id\":null")
		first = false
	}

	// handle non nullable field gpu_devices with omitempty
	if swag.IsZero(m.GpuDevices) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"gpu_devices\":")
		bytes, err := swag.WriteJSON(m.GpuDevices)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
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

	// handle nullable field host_id
	if m.HostID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"host_id\":")
		bytes, err := swag.WriteJSON(m.HostID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.HostID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"host_id\":null")
		first = false
	}

	// handle nullable field is_full_copy
	if m.IsFullCopy != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"is_full_copy\":")
		bytes, err := swag.WriteJSON(m.IsFullCopy)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IsFullCopy_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"is_full_copy\":null")
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

	// handle nullable field memory_unit
	if m.MemoryUnit != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"memory_unit\":")
		bytes, err := swag.WriteJSON(m.MemoryUnit)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.MemoryUnit_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"memory_unit\":null")
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

	// handle nullable field owner
	if m.Owner != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"owner\":")
		bytes, err := swag.WriteJSON(m.Owner)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Owner_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"owner\":null")
		first = false
	}

	// handle nullable field pci_nics
	if m.PciNics != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"pci_nics\":")
		bytes, err := swag.WriteJSON(m.PciNics)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.PciNics_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"pci_nics\":null")
		first = false
	}

	// handle nullable field status
	if m.Status != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"status\":")
		bytes, err := swag.WriteJSON(m.Status)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Status_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"status\":null")
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

	// handle nullable field vm_placement_group
	if m.VMPlacementGroup != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_placement_group\":")
		bytes, err := swag.WriteJSON(m.VMPlacementGroup)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VMPlacementGroup_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_placement_group\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this Vm create Vm from content library template batch Vm params
func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudInit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiskOperate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirmware(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGpuDevices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGuestOsType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsFullCopy(formats); err != nil {
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

	if err := m.validateMemoryUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePciNics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMNics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMPlacementGroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) validateCloudInit(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudInit) { // not required
		return nil
	}

	if m.CloudInit != nil {
		if err := m.CloudInit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloud_init")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloud_init")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("cluster_id", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) validateDiskOperate(formats strfmt.Registry) error {
	if swag.IsZero(m.DiskOperate) { // not required
		return nil
	}

	if m.DiskOperate != nil {
		if err := m.DiskOperate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk_operate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disk_operate")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) validateFirmware(formats strfmt.Registry) error {
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

func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) validateGpuDevices(formats strfmt.Registry) error {
	if swag.IsZero(m.GpuDevices) { // not required
		return nil
	}

	for i := 0; i < len(m.GpuDevices); i++ {
		if swag.IsZero(m.GpuDevices[i]) { // not required
			continue
		}

		if m.GpuDevices[i] != nil {
			if err := m.GpuDevices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("gpu_devices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("gpu_devices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) validateGuestOsType(formats strfmt.Registry) error {
	if swag.IsZero(m.GuestOsType) { // not required
		return nil
	}

	if m.GuestOsType != nil {
		if err := m.GuestOsType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("guest_os_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("guest_os_type")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) validateIsFullCopy(formats strfmt.Registry) error {

	if err := validate.Required("is_full_copy", "body", m.IsFullCopy); err != nil {
		return err
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) validateMaxBandwidthPolicy(formats strfmt.Registry) error {
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

func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) validateMaxBandwidthUnit(formats strfmt.Registry) error {
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

func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) validateMaxIopsPolicy(formats strfmt.Registry) error {
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

func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) validateMemoryUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.MemoryUnit) { // not required
		return nil
	}

	if m.MemoryUnit != nil {
		if err := m.MemoryUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("memory_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("memory_unit")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) validateOwner(formats strfmt.Registry) error {
	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) validatePciNics(formats strfmt.Registry) error {
	if swag.IsZero(m.PciNics) { // not required
		return nil
	}

	if m.PciNics != nil {
		if err := m.PciNics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pci_nics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pci_nics")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) validateVMNics(formats strfmt.Registry) error {
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

func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) validateVMPlacementGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.VMPlacementGroup) { // not required
		return nil
	}

	if m.VMPlacementGroup != nil {
		if err := m.VMPlacementGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_placement_group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_placement_group")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this Vm create Vm from content library template batch Vm params based on the context it is used
func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCloudInit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDiskOperate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFirmware(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGpuDevices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGuestOsType(ctx, formats); err != nil {
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

	if err := m.contextValidateMemoryUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePciNics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMNics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMPlacementGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) contextValidateCloudInit(ctx context.Context, formats strfmt.Registry) error {

	if m.CloudInit != nil {
		if err := m.CloudInit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloud_init")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloud_init")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) contextValidateDiskOperate(ctx context.Context, formats strfmt.Registry) error {

	if m.DiskOperate != nil {
		if err := m.DiskOperate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk_operate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disk_operate")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) contextValidateFirmware(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) contextValidateGpuDevices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.GpuDevices); i++ {

		if m.GpuDevices[i] != nil {
			if err := m.GpuDevices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("gpu_devices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("gpu_devices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) contextValidateGuestOsType(ctx context.Context, formats strfmt.Registry) error {

	if m.GuestOsType != nil {
		if err := m.GuestOsType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("guest_os_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("guest_os_type")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) contextValidateMaxBandwidthPolicy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) contextValidateMaxBandwidthUnit(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) contextValidateMaxIopsPolicy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) contextValidateMemoryUnit(ctx context.Context, formats strfmt.Registry) error {

	if m.MemoryUnit != nil {
		if err := m.MemoryUnit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("memory_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("memory_unit")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

	if m.Owner != nil {
		if err := m.Owner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) contextValidatePciNics(ctx context.Context, formats strfmt.Registry) error {

	if m.PciNics != nil {
		if err := m.PciNics.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pci_nics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pci_nics")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) contextValidateVMNics(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) contextValidateVMPlacementGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.VMPlacementGroup != nil {
		if err := m.VMPlacementGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_placement_group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_placement_group")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMCreateVMFromContentLibraryTemplateBatchVMParams) UnmarshalBinary(b []byte) error {
	var res VMCreateVMFromContentLibraryTemplateBatchVMParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
