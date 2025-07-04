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
)

// VMMetaData Vm meta data
//
// swagger:model VmMetaData
type VMMetaData struct {

	// boot with host
	BootWithHost *bool `json:"boot_with_host,omitempty"`

	// clock offset
	ClockOffset *VMClockOffset `json:"clock_offset,omitempty"`

	// cluster architecture
	ClusterArchitecture *Architecture `json:"cluster_architecture,omitempty"`

	// cluster local id
	ClusterLocalID *string `json:"cluster_local_id,omitempty"`

	// cluster type
	ClusterType *ClusterType `json:"cluster_type,omitempty"`

	// cluster version
	ClusterVersion *string `json:"cluster_version,omitempty"`

	// cluster vhost enabled
	ClusterVhostEnabled *bool `json:"cluster_vhost_enabled,omitempty"`

	// cpu cores
	CPUCores *int32 `json:"cpu_cores,omitempty"`

	// cpu exclusive expected enabled
	CPUExclusiveExpectedEnabled *bool `json:"cpu_exclusive_expected_enabled,omitempty"`

	// cpu model
	CPUModel *string `json:"cpu_model,omitempty"`

	// cpu sockets
	CPUSockets *int32 `json:"cpu_sockets,omitempty"`

	// description
	Description *string `json:"description,omitempty"`

	// firmware
	Firmware *VMFirmware `json:"firmware,omitempty"`

	// guest os type
	GuestOsType *VMGuestsOperationSystem `json:"guest_os_type,omitempty"`

	// ha
	Ha *bool `json:"ha,omitempty"`

	// ha priority
	HaPriority *VMHaPriority `json:"ha_priority,omitempty"`

	// host local id
	HostLocalID *string `json:"host_local_id,omitempty"`

	// internal
	Internal *bool `json:"internal,omitempty"`

	// io policy
	IoPolicy *VMDiskIoPolicy `json:"io_policy,omitempty"`

	// local id
	LocalID *string `json:"local_id,omitempty"`

	// max bandwidth
	MaxBandwidth *float64 `json:"max_bandwidth,omitempty"`

	// max bandwidth policy
	MaxBandwidthPolicy *VMDiskIoRestrictType `json:"max_bandwidth_policy,omitempty"`

	// max iops
	MaxIops *int32 `json:"max_iops,omitempty"`

	// max iops policy
	MaxIopsPolicy *VMDiskIoRestrictType `json:"max_iops_policy,omitempty"`

	// memory
	Memory *float64 `json:"memory,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// nested virtualization
	NestedVirtualization *bool `json:"nested_virtualization,omitempty"`

	// protected
	Protected *bool `json:"protected,omitempty"`

	// sync vm time on resume
	SyncVMTimeOnResume *bool `json:"sync_vm_time_on_resume,omitempty"`

	// vcpu
	Vcpu *int32 `json:"vcpu,omitempty"`

	// video type
	VideoType *VMVideoType `json:"video_type,omitempty"`

	// vm disks
	VMDisks []*VMDiskMetaData `json:"vm_disks,omitempty"`

	// vm nics
	VMNics []*VMNicMetaData `json:"vm_nics,omitempty"`

	// vm tools status
	VMToolsStatus *VMToolsStatus `json:"vm_tools_status,omitempty"`

	// win opt
	WinOpt *bool `json:"win_opt,omitempty"`

	MarshalOpts *VMMetaDataMarshalOpts `json:"-"`
}

type VMMetaDataMarshalOpts struct {
	BootWithHost_Explicit_Null_When_Empty bool

	ClockOffset_Explicit_Null_When_Empty bool

	ClusterArchitecture_Explicit_Null_When_Empty bool

	ClusterLocalID_Explicit_Null_When_Empty bool

	ClusterType_Explicit_Null_When_Empty bool

	ClusterVersion_Explicit_Null_When_Empty bool

	ClusterVhostEnabled_Explicit_Null_When_Empty bool

	CPUCores_Explicit_Null_When_Empty bool

	CPUExclusiveExpectedEnabled_Explicit_Null_When_Empty bool

	CPUModel_Explicit_Null_When_Empty bool

	CPUSockets_Explicit_Null_When_Empty bool

	Description_Explicit_Null_When_Empty bool

	Firmware_Explicit_Null_When_Empty bool

	GuestOsType_Explicit_Null_When_Empty bool

	Ha_Explicit_Null_When_Empty bool

	HaPriority_Explicit_Null_When_Empty bool

	HostLocalID_Explicit_Null_When_Empty bool

	Internal_Explicit_Null_When_Empty bool

	IoPolicy_Explicit_Null_When_Empty bool

	LocalID_Explicit_Null_When_Empty bool

	MaxBandwidth_Explicit_Null_When_Empty bool

	MaxBandwidthPolicy_Explicit_Null_When_Empty bool

	MaxIops_Explicit_Null_When_Empty bool

	MaxIopsPolicy_Explicit_Null_When_Empty bool

	Memory_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	NestedVirtualization_Explicit_Null_When_Empty bool

	Protected_Explicit_Null_When_Empty bool

	SyncVMTimeOnResume_Explicit_Null_When_Empty bool

	Vcpu_Explicit_Null_When_Empty bool

	VideoType_Explicit_Null_When_Empty bool

	VMDisks_Explicit_Null_When_Empty bool

	VMNics_Explicit_Null_When_Empty bool

	VMToolsStatus_Explicit_Null_When_Empty bool

	WinOpt_Explicit_Null_When_Empty bool
}

func (m VMMetaData) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field boot_with_host
	if m.BootWithHost != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"boot_with_host\":")
		bytes, err := swag.WriteJSON(m.BootWithHost)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.BootWithHost_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"boot_with_host\":null")
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

	// handle nullable field cluster_architecture
	if m.ClusterArchitecture != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster_architecture\":")
		bytes, err := swag.WriteJSON(m.ClusterArchitecture)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ClusterArchitecture_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster_architecture\":null")
		first = false
	}

	// handle nullable field cluster_local_id
	if m.ClusterLocalID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster_local_id\":")
		bytes, err := swag.WriteJSON(m.ClusterLocalID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ClusterLocalID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster_local_id\":null")
		first = false
	}

	// handle nullable field cluster_type
	if m.ClusterType != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster_type\":")
		bytes, err := swag.WriteJSON(m.ClusterType)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ClusterType_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster_type\":null")
		first = false
	}

	// handle nullable field cluster_version
	if m.ClusterVersion != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster_version\":")
		bytes, err := swag.WriteJSON(m.ClusterVersion)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ClusterVersion_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster_version\":null")
		first = false
	}

	// handle nullable field cluster_vhost_enabled
	if m.ClusterVhostEnabled != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster_vhost_enabled\":")
		bytes, err := swag.WriteJSON(m.ClusterVhostEnabled)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ClusterVhostEnabled_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster_vhost_enabled\":null")
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

	// handle nullable field ha_priority
	if m.HaPriority != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ha_priority\":")
		bytes, err := swag.WriteJSON(m.HaPriority)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.HaPriority_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ha_priority\":null")
		first = false
	}

	// handle nullable field host_local_id
	if m.HostLocalID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"host_local_id\":")
		bytes, err := swag.WriteJSON(m.HostLocalID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.HostLocalID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"host_local_id\":null")
		first = false
	}

	// handle nullable field internal
	if m.Internal != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"internal\":")
		bytes, err := swag.WriteJSON(m.Internal)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Internal_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"internal\":null")
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

	// handle nullable field nested_virtualization
	if m.NestedVirtualization != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"nested_virtualization\":")
		bytes, err := swag.WriteJSON(m.NestedVirtualization)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NestedVirtualization_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"nested_virtualization\":null")
		first = false
	}

	// handle nullable field protected
	if m.Protected != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"protected\":")
		bytes, err := swag.WriteJSON(m.Protected)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Protected_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"protected\":null")
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

	// handle nullable field vm_tools_status
	if m.VMToolsStatus != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_tools_status\":")
		bytes, err := swag.WriteJSON(m.VMToolsStatus)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VMToolsStatus_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_tools_status\":null")
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

// Validate validates this Vm meta data
func (m *VMMetaData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClockOffset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterArchitecture(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirmware(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGuestOsType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHaPriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIoPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxBandwidthPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxIopsPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVideoType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMNics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMToolsStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMMetaData) validateClockOffset(formats strfmt.Registry) error {
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

func (m *VMMetaData) validateClusterArchitecture(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterArchitecture) { // not required
		return nil
	}

	if m.ClusterArchitecture != nil {
		if err := m.ClusterArchitecture.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_architecture")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster_architecture")
			}
			return err
		}
	}

	return nil
}

func (m *VMMetaData) validateClusterType(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterType) { // not required
		return nil
	}

	if m.ClusterType != nil {
		if err := m.ClusterType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster_type")
			}
			return err
		}
	}

	return nil
}

func (m *VMMetaData) validateFirmware(formats strfmt.Registry) error {
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

func (m *VMMetaData) validateGuestOsType(formats strfmt.Registry) error {
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

func (m *VMMetaData) validateHaPriority(formats strfmt.Registry) error {
	if swag.IsZero(m.HaPriority) { // not required
		return nil
	}

	if m.HaPriority != nil {
		if err := m.HaPriority.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ha_priority")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ha_priority")
			}
			return err
		}
	}

	return nil
}

func (m *VMMetaData) validateIoPolicy(formats strfmt.Registry) error {
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

func (m *VMMetaData) validateMaxBandwidthPolicy(formats strfmt.Registry) error {
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

func (m *VMMetaData) validateMaxIopsPolicy(formats strfmt.Registry) error {
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

func (m *VMMetaData) validateVideoType(formats strfmt.Registry) error {
	if swag.IsZero(m.VideoType) { // not required
		return nil
	}

	if m.VideoType != nil {
		if err := m.VideoType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("video_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("video_type")
			}
			return err
		}
	}

	return nil
}

func (m *VMMetaData) validateVMDisks(formats strfmt.Registry) error {
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

func (m *VMMetaData) validateVMNics(formats strfmt.Registry) error {
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

func (m *VMMetaData) validateVMToolsStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.VMToolsStatus) { // not required
		return nil
	}

	if m.VMToolsStatus != nil {
		if err := m.VMToolsStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_tools_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_tools_status")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this Vm meta data based on the context it is used
func (m *VMMetaData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClockOffset(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClusterArchitecture(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClusterType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFirmware(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGuestOsType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHaPriority(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIoPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxBandwidthPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxIopsPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVideoType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMNics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMToolsStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMMetaData) contextValidateClockOffset(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMMetaData) contextValidateClusterArchitecture(ctx context.Context, formats strfmt.Registry) error {

	if m.ClusterArchitecture != nil {
		if err := m.ClusterArchitecture.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_architecture")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster_architecture")
			}
			return err
		}
	}

	return nil
}

func (m *VMMetaData) contextValidateClusterType(ctx context.Context, formats strfmt.Registry) error {

	if m.ClusterType != nil {
		if err := m.ClusterType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster_type")
			}
			return err
		}
	}

	return nil
}

func (m *VMMetaData) contextValidateFirmware(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMMetaData) contextValidateGuestOsType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMMetaData) contextValidateHaPriority(ctx context.Context, formats strfmt.Registry) error {

	if m.HaPriority != nil {
		if err := m.HaPriority.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ha_priority")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ha_priority")
			}
			return err
		}
	}

	return nil
}

func (m *VMMetaData) contextValidateIoPolicy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMMetaData) contextValidateMaxBandwidthPolicy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMMetaData) contextValidateMaxIopsPolicy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMMetaData) contextValidateVideoType(ctx context.Context, formats strfmt.Registry) error {

	if m.VideoType != nil {
		if err := m.VideoType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("video_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("video_type")
			}
			return err
		}
	}

	return nil
}

func (m *VMMetaData) contextValidateVMDisks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMMetaData) contextValidateVMNics(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMMetaData) contextValidateVMToolsStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.VMToolsStatus != nil {
		if err := m.VMToolsStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_tools_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_tools_status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMMetaData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMMetaData) UnmarshalBinary(b []byte) error {
	var res VMMetaData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
