// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Host host
//
// swagger:model Host
type Host struct {

	// access ip
	AccessIP *string `json:"access_ip,omitempty"`

	// allocable cpu cores for vm exclusive
	AllocableCPUCoresForVMExclusive *int32 `json:"allocable_cpu_cores_for_vm_exclusive,omitempty"`

	// allocatable memory bytes
	// Required: true
	AllocatableMemoryBytes *int64 `json:"allocatable_memory_bytes"`

	// chunk id
	// Required: true
	ChunkID *string `json:"chunk_id"`

	// cluster
	// Required: true
	Cluster *NestedCluster `json:"cluster"`

	// compatible cpu models
	// Required: true
	CompatibleCPUModels []string `json:"compatible_cpu_models"`

	// cpu brand
	// Required: true
	CPUBrand *string `json:"cpu_brand"`

	// cpu fan speed
	// Required: true
	CPUFanSpeed []float64 `json:"cpu_fan_speed"`

	// cpu fan speed unit
	CPUFanSpeedUnit *CPUFanSpeedUnit `json:"cpu_fan_speed_unit,omitempty"`

	// cpu hz per core
	// Required: true
	CPUHzPerCore *int64 `json:"cpu_hz_per_core"`

	// cpu model
	// Required: true
	CPUModel *string `json:"cpu_model"`

	// cpu temperature celsius
	// Required: true
	CPUTemperatureCelsius []int32 `json:"cpu_temperature_celsius"`

	// cpu vendor
	CPUVendor *string `json:"cpu_vendor,omitempty"`

	// data ip
	DataIP *string `json:"data_ip,omitempty"`

	// disks
	Disks []*NestedDisk `json:"disks,omitempty"`

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// failure data space
	// Required: true
	FailureDataSpace *int64 `json:"failure_data_space"`

	// hdd data capacity
	// Required: true
	HddDataCapacity *int64 `json:"hdd_data_capacity"`

	// hdd disk count
	// Required: true
	HddDiskCount *int32 `json:"hdd_disk_count"`

	// host state
	HostState *NestedMaintenanceHostState `json:"host_state,omitempty"`

	// hypervisor ip
	HypervisorIP *string `json:"hypervisor_ip,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// ipmi
	Ipmi *NestedIpmi `json:"ipmi,omitempty"`

	// is os in raid1
	IsOsInRaid1 *bool `json:"is_os_in_raid1,omitempty"`

	// labels
	Labels []*NestedLabel `json:"labels,omitempty"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// lsm cap disk safe umount
	// Required: true
	LsmCapDiskSafeUmount *bool `json:"lsm_cap_disk_safe_umount"`

	// management ip
	// Required: true
	ManagementIP *string `json:"management_ip"`

	// model
	// Required: true
	Model *string `json:"model"`

	// name
	// Required: true
	Name *string `json:"name"`

	// nested virtualization
	// Required: true
	NestedVirtualization *bool `json:"nested_virtualization"`

	// nic count
	// Required: true
	NicCount *int32 `json:"nic_count"`

	// nics
	Nics []*NestedNic `json:"nics,omitempty"`

	// node topo local id
	NodeTopoLocalID *string `json:"node_topo_local_id,omitempty"`

	// os memory bytes
	// Required: true
	OsMemoryBytes *int64 `json:"os_memory_bytes"`

	// os version
	OsVersion *string `json:"os_version,omitempty"`

	// pmem dimm capacity
	// Required: true
	PmemDimmCapacity *int64 `json:"pmem_dimm_capacity"`

	// pmem dimm count
	// Required: true
	PmemDimmCount *int32 `json:"pmem_dimm_count"`

	// pmem dimms
	PmemDimms []*NestedPmemDimm `json:"pmem_dimms,omitempty"`

	// pmem disk count
	// Required: true
	PmemDiskCount *int32 `json:"pmem_disk_count"`

	// provisioned cpu cores
	// Required: true
	ProvisionedCPUCores *int32 `json:"provisioned_cpu_cores"`

	// provisioned memory bytes
	// Required: true
	ProvisionedMemoryBytes *int64 `json:"provisioned_memory_bytes"`

	// running pause vm memory bytes
	// Required: true
	RunningPauseVMMemoryBytes *int64 `json:"running_pause_vm_memory_bytes"`

	// running vm num
	RunningVMNum *int32 `json:"running_vm_num,omitempty"`

	// scvm cpu
	ScvmCPU *int32 `json:"scvm_cpu,omitempty"`

	// scvm memory
	ScvmMemory *int64 `json:"scvm_memory,omitempty"`

	// scvm name
	ScvmName *string `json:"scvm_name,omitempty"`

	// serial
	Serial *string `json:"serial,omitempty"`

	// ssd data capacity
	// Required: true
	SsdDataCapacity *int64 `json:"ssd_data_capacity"`

	// ssd disk count
	// Required: true
	SsdDiskCount *int32 `json:"ssd_disk_count"`

	// state
	// Required: true
	State *HostState `json:"state"`

	// status
	// Required: true
	Status *HostStatus `json:"status"`

	// stopped vm num
	StoppedVMNum *int32 `json:"stopped_vm_num,omitempty"`

	// suspended vm num
	SuspendedVMNum *int32 `json:"suspended_vm_num,omitempty"`

	// total cache capacity
	TotalCacheCapacity *int64 `json:"total_cache_capacity,omitempty"`

	// total cpu cores
	// Required: true
	TotalCPUCores *int32 `json:"total_cpu_cores"`

	// total cpu hz
	// Required: true
	TotalCPUHz *int64 `json:"total_cpu_hz"`

	// total cpu sockets
	TotalCPUSockets *int32 `json:"total_cpu_sockets,omitempty"`

	// total data capacity
	// Required: true
	TotalDataCapacity *int64 `json:"total_data_capacity"`

	// total memory bytes
	// Required: true
	TotalMemoryBytes *int64 `json:"total_memory_bytes"`

	// usb devices
	UsbDevices []*NestedUsbDevice `json:"usb_devices,omitempty"`

	// used cpu hz
	UsedCPUHz *float64 `json:"used_cpu_hz,omitempty"`

	// used data space
	// Required: true
	UsedDataSpace *int64 `json:"used_data_space"`

	// used memory bytes
	UsedMemoryBytes *float64 `json:"used_memory_bytes,omitempty"`

	// vm num
	VMNum *int32 `json:"vm_num,omitempty"`

	// vmotion ip
	VmotionIP *string `json:"vmotion_ip,omitempty"`

	// vms
	Vms []*NestedVM `json:"vms,omitempty"`

	// vsphere esxi account
	VsphereEsxiAccount *NestedVsphereEsxiAccount `json:"vsphereEsxiAccount,omitempty"`

	// with faster ssd as cache
	WithFasterSsdAsCache *bool `json:"with_faster_ssd_as_cache,omitempty"`

	// zone
	Zone *NestedZone `json:"zone,omitempty"`
}

// Validate validates this host
func (m *Host) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocatableMemoryBytes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChunkID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompatibleCPUModels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPUBrand(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPUFanSpeed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPUFanSpeedUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPUHzPerCore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPUModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPUTemperatureCelsius(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFailureDataSpace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHddDataCapacity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHddDiskCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIpmi(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLsmCapDiskSafeUmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManagementIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNestedVirtualization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNicCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOsMemoryBytes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePmemDimmCapacity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePmemDimmCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePmemDimms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePmemDiskCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvisionedCPUCores(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvisionedMemoryBytes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunningPauseVMMemoryBytes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSsdDataCapacity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSsdDiskCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalCPUCores(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalCPUHz(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalDataCapacity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalMemoryBytes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsbDevices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsedDataSpace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVsphereEsxiAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Host) validateAllocatableMemoryBytes(formats strfmt.Registry) error {

	if err := validate.Required("allocatable_memory_bytes", "body", m.AllocatableMemoryBytes); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateChunkID(formats strfmt.Registry) error {

	if err := validate.Required("chunk_id", "body", m.ChunkID); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateCluster(formats strfmt.Registry) error {

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

func (m *Host) validateCompatibleCPUModels(formats strfmt.Registry) error {

	if err := validate.Required("compatible_cpu_models", "body", m.CompatibleCPUModels); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateCPUBrand(formats strfmt.Registry) error {

	if err := validate.Required("cpu_brand", "body", m.CPUBrand); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateCPUFanSpeed(formats strfmt.Registry) error {

	if err := validate.Required("cpu_fan_speed", "body", m.CPUFanSpeed); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateCPUFanSpeedUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.CPUFanSpeedUnit) { // not required
		return nil
	}

	if m.CPUFanSpeedUnit != nil {
		if err := m.CPUFanSpeedUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cpu_fan_speed_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cpu_fan_speed_unit")
			}
			return err
		}
	}

	return nil
}

func (m *Host) validateCPUHzPerCore(formats strfmt.Registry) error {

	if err := validate.Required("cpu_hz_per_core", "body", m.CPUHzPerCore); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateCPUModel(formats strfmt.Registry) error {

	if err := validate.Required("cpu_model", "body", m.CPUModel); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateCPUTemperatureCelsius(formats strfmt.Registry) error {

	if err := validate.Required("cpu_temperature_celsius", "body", m.CPUTemperatureCelsius); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateDisks(formats strfmt.Registry) error {
	if swag.IsZero(m.Disks) { // not required
		return nil
	}

	for i := 0; i < len(m.Disks); i++ {
		if swag.IsZero(m.Disks[i]) { // not required
			continue
		}

		if m.Disks[i] != nil {
			if err := m.Disks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Host) validateEntityAsyncStatus(formats strfmt.Registry) error {
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

func (m *Host) validateFailureDataSpace(formats strfmt.Registry) error {

	if err := validate.Required("failure_data_space", "body", m.FailureDataSpace); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateHddDataCapacity(formats strfmt.Registry) error {

	if err := validate.Required("hdd_data_capacity", "body", m.HddDataCapacity); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateHddDiskCount(formats strfmt.Registry) error {

	if err := validate.Required("hdd_disk_count", "body", m.HddDiskCount); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateHostState(formats strfmt.Registry) error {
	if swag.IsZero(m.HostState) { // not required
		return nil
	}

	if m.HostState != nil {
		if err := m.HostState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host_state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("host_state")
			}
			return err
		}
	}

	return nil
}

func (m *Host) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateIpmi(formats strfmt.Registry) error {
	if swag.IsZero(m.Ipmi) { // not required
		return nil
	}

	if m.Ipmi != nil {
		if err := m.Ipmi.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipmi")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipmi")
			}
			return err
		}
	}

	return nil
}

func (m *Host) validateLabels(formats strfmt.Registry) error {
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

func (m *Host) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateLsmCapDiskSafeUmount(formats strfmt.Registry) error {

	if err := validate.Required("lsm_cap_disk_safe_umount", "body", m.LsmCapDiskSafeUmount); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateManagementIP(formats strfmt.Registry) error {

	if err := validate.Required("management_ip", "body", m.ManagementIP); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateModel(formats strfmt.Registry) error {

	if err := validate.Required("model", "body", m.Model); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateNestedVirtualization(formats strfmt.Registry) error {

	if err := validate.Required("nested_virtualization", "body", m.NestedVirtualization); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateNicCount(formats strfmt.Registry) error {

	if err := validate.Required("nic_count", "body", m.NicCount); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateNics(formats strfmt.Registry) error {
	if swag.IsZero(m.Nics) { // not required
		return nil
	}

	for i := 0; i < len(m.Nics); i++ {
		if swag.IsZero(m.Nics[i]) { // not required
			continue
		}

		if m.Nics[i] != nil {
			if err := m.Nics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Host) validateOsMemoryBytes(formats strfmt.Registry) error {

	if err := validate.Required("os_memory_bytes", "body", m.OsMemoryBytes); err != nil {
		return err
	}

	return nil
}

func (m *Host) validatePmemDimmCapacity(formats strfmt.Registry) error {

	if err := validate.Required("pmem_dimm_capacity", "body", m.PmemDimmCapacity); err != nil {
		return err
	}

	return nil
}

func (m *Host) validatePmemDimmCount(formats strfmt.Registry) error {

	if err := validate.Required("pmem_dimm_count", "body", m.PmemDimmCount); err != nil {
		return err
	}

	return nil
}

func (m *Host) validatePmemDimms(formats strfmt.Registry) error {
	if swag.IsZero(m.PmemDimms) { // not required
		return nil
	}

	for i := 0; i < len(m.PmemDimms); i++ {
		if swag.IsZero(m.PmemDimms[i]) { // not required
			continue
		}

		if m.PmemDimms[i] != nil {
			if err := m.PmemDimms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pmem_dimms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pmem_dimms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Host) validatePmemDiskCount(formats strfmt.Registry) error {

	if err := validate.Required("pmem_disk_count", "body", m.PmemDiskCount); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateProvisionedCPUCores(formats strfmt.Registry) error {

	if err := validate.Required("provisioned_cpu_cores", "body", m.ProvisionedCPUCores); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateProvisionedMemoryBytes(formats strfmt.Registry) error {

	if err := validate.Required("provisioned_memory_bytes", "body", m.ProvisionedMemoryBytes); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateRunningPauseVMMemoryBytes(formats strfmt.Registry) error {

	if err := validate.Required("running_pause_vm_memory_bytes", "body", m.RunningPauseVMMemoryBytes); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateSsdDataCapacity(formats strfmt.Registry) error {

	if err := validate.Required("ssd_data_capacity", "body", m.SsdDataCapacity); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateSsdDiskCount(formats strfmt.Registry) error {

	if err := validate.Required("ssd_disk_count", "body", m.SsdDiskCount); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *Host) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
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

func (m *Host) validateTotalCPUCores(formats strfmt.Registry) error {

	if err := validate.Required("total_cpu_cores", "body", m.TotalCPUCores); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateTotalCPUHz(formats strfmt.Registry) error {

	if err := validate.Required("total_cpu_hz", "body", m.TotalCPUHz); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateTotalDataCapacity(formats strfmt.Registry) error {

	if err := validate.Required("total_data_capacity", "body", m.TotalDataCapacity); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateTotalMemoryBytes(formats strfmt.Registry) error {

	if err := validate.Required("total_memory_bytes", "body", m.TotalMemoryBytes); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateUsbDevices(formats strfmt.Registry) error {
	if swag.IsZero(m.UsbDevices) { // not required
		return nil
	}

	for i := 0; i < len(m.UsbDevices); i++ {
		if swag.IsZero(m.UsbDevices[i]) { // not required
			continue
		}

		if m.UsbDevices[i] != nil {
			if err := m.UsbDevices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("usb_devices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("usb_devices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Host) validateUsedDataSpace(formats strfmt.Registry) error {

	if err := validate.Required("used_data_space", "body", m.UsedDataSpace); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateVms(formats strfmt.Registry) error {
	if swag.IsZero(m.Vms) { // not required
		return nil
	}

	for i := 0; i < len(m.Vms); i++ {
		if swag.IsZero(m.Vms[i]) { // not required
			continue
		}

		if m.Vms[i] != nil {
			if err := m.Vms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Host) validateVsphereEsxiAccount(formats strfmt.Registry) error {
	if swag.IsZero(m.VsphereEsxiAccount) { // not required
		return nil
	}

	if m.VsphereEsxiAccount != nil {
		if err := m.VsphereEsxiAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vsphereEsxiAccount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vsphereEsxiAccount")
			}
			return err
		}
	}

	return nil
}

func (m *Host) validateZone(formats strfmt.Registry) error {
	if swag.IsZero(m.Zone) { // not required
		return nil
	}

	if m.Zone != nil {
		if err := m.Zone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("zone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("zone")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this host based on the context it is used
func (m *Host) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCPUFanSpeedUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHostState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIpmi(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePmemDimms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsbDevices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVsphereEsxiAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateZone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Host) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Host) contextValidateCPUFanSpeedUnit(ctx context.Context, formats strfmt.Registry) error {

	if m.CPUFanSpeedUnit != nil {
		if err := m.CPUFanSpeedUnit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cpu_fan_speed_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cpu_fan_speed_unit")
			}
			return err
		}
	}

	return nil
}

func (m *Host) contextValidateDisks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Disks); i++ {

		if m.Disks[i] != nil {
			if err := m.Disks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Host) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Host) contextValidateHostState(ctx context.Context, formats strfmt.Registry) error {

	if m.HostState != nil {
		if err := m.HostState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host_state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("host_state")
			}
			return err
		}
	}

	return nil
}

func (m *Host) contextValidateIpmi(ctx context.Context, formats strfmt.Registry) error {

	if m.Ipmi != nil {
		if err := m.Ipmi.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipmi")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipmi")
			}
			return err
		}
	}

	return nil
}

func (m *Host) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Host) contextValidateNics(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Nics); i++ {

		if m.Nics[i] != nil {
			if err := m.Nics[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Host) contextValidatePmemDimms(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PmemDimms); i++ {

		if m.PmemDimms[i] != nil {
			if err := m.PmemDimms[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pmem_dimms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pmem_dimms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Host) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if m.State != nil {
		if err := m.State.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *Host) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Host) contextValidateUsbDevices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UsbDevices); i++ {

		if m.UsbDevices[i] != nil {
			if err := m.UsbDevices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("usb_devices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("usb_devices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Host) contextValidateVms(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Vms); i++ {

		if m.Vms[i] != nil {
			if err := m.Vms[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Host) contextValidateVsphereEsxiAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.VsphereEsxiAccount != nil {
		if err := m.VsphereEsxiAccount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vsphereEsxiAccount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vsphereEsxiAccount")
			}
			return err
		}
	}

	return nil
}

func (m *Host) contextValidateZone(ctx context.Context, formats strfmt.Registry) error {

	if m.Zone != nil {
		if err := m.Zone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("zone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("zone")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Host) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Host) UnmarshalBinary(b []byte) error {
	var res Host
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
