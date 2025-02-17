// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ContentLibraryVMTemplateOrderByInput content library Vm template order by input
//
// swagger:model ContentLibraryVmTemplateOrderByInput
type ContentLibraryVMTemplateOrderByInput string

func NewContentLibraryVMTemplateOrderByInput(value ContentLibraryVMTemplateOrderByInput) *ContentLibraryVMTemplateOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ContentLibraryVMTemplateOrderByInput.
func (m ContentLibraryVMTemplateOrderByInput) Pointer() *ContentLibraryVMTemplateOrderByInput {
	return &m
}

const (

	// ContentLibraryVMTemplateOrderByInputArchitectureASC captures enum value "architecture_ASC"
	ContentLibraryVMTemplateOrderByInputArchitectureASC ContentLibraryVMTemplateOrderByInput = "architecture_ASC"

	// ContentLibraryVMTemplateOrderByInputArchitectureDESC captures enum value "architecture_DESC"
	ContentLibraryVMTemplateOrderByInputArchitectureDESC ContentLibraryVMTemplateOrderByInput = "architecture_DESC"

	// ContentLibraryVMTemplateOrderByInputClockOffsetASC captures enum value "clock_offset_ASC"
	ContentLibraryVMTemplateOrderByInputClockOffsetASC ContentLibraryVMTemplateOrderByInput = "clock_offset_ASC"

	// ContentLibraryVMTemplateOrderByInputClockOffsetDESC captures enum value "clock_offset_DESC"
	ContentLibraryVMTemplateOrderByInputClockOffsetDESC ContentLibraryVMTemplateOrderByInput = "clock_offset_DESC"

	// ContentLibraryVMTemplateOrderByInputCloudInitSupportedASC captures enum value "cloud_init_supported_ASC"
	ContentLibraryVMTemplateOrderByInputCloudInitSupportedASC ContentLibraryVMTemplateOrderByInput = "cloud_init_supported_ASC"

	// ContentLibraryVMTemplateOrderByInputCloudInitSupportedDESC captures enum value "cloud_init_supported_DESC"
	ContentLibraryVMTemplateOrderByInputCloudInitSupportedDESC ContentLibraryVMTemplateOrderByInput = "cloud_init_supported_DESC"

	// ContentLibraryVMTemplateOrderByInputCPUASC captures enum value "cpu_ASC"
	ContentLibraryVMTemplateOrderByInputCPUASC ContentLibraryVMTemplateOrderByInput = "cpu_ASC"

	// ContentLibraryVMTemplateOrderByInputCPUDESC captures enum value "cpu_DESC"
	ContentLibraryVMTemplateOrderByInputCPUDESC ContentLibraryVMTemplateOrderByInput = "cpu_DESC"

	// ContentLibraryVMTemplateOrderByInputCPUModelASC captures enum value "cpu_model_ASC"
	ContentLibraryVMTemplateOrderByInputCPUModelASC ContentLibraryVMTemplateOrderByInput = "cpu_model_ASC"

	// ContentLibraryVMTemplateOrderByInputCPUModelDESC captures enum value "cpu_model_DESC"
	ContentLibraryVMTemplateOrderByInputCPUModelDESC ContentLibraryVMTemplateOrderByInput = "cpu_model_DESC"

	// ContentLibraryVMTemplateOrderByInputCreatedAtASC captures enum value "createdAt_ASC"
	ContentLibraryVMTemplateOrderByInputCreatedAtASC ContentLibraryVMTemplateOrderByInput = "createdAt_ASC"

	// ContentLibraryVMTemplateOrderByInputCreatedAtDESC captures enum value "createdAt_DESC"
	ContentLibraryVMTemplateOrderByInputCreatedAtDESC ContentLibraryVMTemplateOrderByInput = "createdAt_DESC"

	// ContentLibraryVMTemplateOrderByInputDescriptionASC captures enum value "description_ASC"
	ContentLibraryVMTemplateOrderByInputDescriptionASC ContentLibraryVMTemplateOrderByInput = "description_ASC"

	// ContentLibraryVMTemplateOrderByInputDescriptionDESC captures enum value "description_DESC"
	ContentLibraryVMTemplateOrderByInputDescriptionDESC ContentLibraryVMTemplateOrderByInput = "description_DESC"

	// ContentLibraryVMTemplateOrderByInputEntityAsyncStatusASC captures enum value "entityAsyncStatus_ASC"
	ContentLibraryVMTemplateOrderByInputEntityAsyncStatusASC ContentLibraryVMTemplateOrderByInput = "entityAsyncStatus_ASC"

	// ContentLibraryVMTemplateOrderByInputEntityAsyncStatusDESC captures enum value "entityAsyncStatus_DESC"
	ContentLibraryVMTemplateOrderByInputEntityAsyncStatusDESC ContentLibraryVMTemplateOrderByInput = "entityAsyncStatus_DESC"

	// ContentLibraryVMTemplateOrderByInputFirmwareASC captures enum value "firmware_ASC"
	ContentLibraryVMTemplateOrderByInputFirmwareASC ContentLibraryVMTemplateOrderByInput = "firmware_ASC"

	// ContentLibraryVMTemplateOrderByInputFirmwareDESC captures enum value "firmware_DESC"
	ContentLibraryVMTemplateOrderByInputFirmwareDESC ContentLibraryVMTemplateOrderByInput = "firmware_DESC"

	// ContentLibraryVMTemplateOrderByInputHaASC captures enum value "ha_ASC"
	ContentLibraryVMTemplateOrderByInputHaASC ContentLibraryVMTemplateOrderByInput = "ha_ASC"

	// ContentLibraryVMTemplateOrderByInputHaDESC captures enum value "ha_DESC"
	ContentLibraryVMTemplateOrderByInputHaDESC ContentLibraryVMTemplateOrderByInput = "ha_DESC"

	// ContentLibraryVMTemplateOrderByInputIDASC captures enum value "id_ASC"
	ContentLibraryVMTemplateOrderByInputIDASC ContentLibraryVMTemplateOrderByInput = "id_ASC"

	// ContentLibraryVMTemplateOrderByInputIDDESC captures enum value "id_DESC"
	ContentLibraryVMTemplateOrderByInputIDDESC ContentLibraryVMTemplateOrderByInput = "id_DESC"

	// ContentLibraryVMTemplateOrderByInputIoPolicyASC captures enum value "io_policy_ASC"
	ContentLibraryVMTemplateOrderByInputIoPolicyASC ContentLibraryVMTemplateOrderByInput = "io_policy_ASC"

	// ContentLibraryVMTemplateOrderByInputIoPolicyDESC captures enum value "io_policy_DESC"
	ContentLibraryVMTemplateOrderByInputIoPolicyDESC ContentLibraryVMTemplateOrderByInput = "io_policy_DESC"

	// ContentLibraryVMTemplateOrderByInputMaxBandwidthASC captures enum value "max_bandwidth_ASC"
	ContentLibraryVMTemplateOrderByInputMaxBandwidthASC ContentLibraryVMTemplateOrderByInput = "max_bandwidth_ASC"

	// ContentLibraryVMTemplateOrderByInputMaxBandwidthDESC captures enum value "max_bandwidth_DESC"
	ContentLibraryVMTemplateOrderByInputMaxBandwidthDESC ContentLibraryVMTemplateOrderByInput = "max_bandwidth_DESC"

	// ContentLibraryVMTemplateOrderByInputMaxBandwidthPolicyASC captures enum value "max_bandwidth_policy_ASC"
	ContentLibraryVMTemplateOrderByInputMaxBandwidthPolicyASC ContentLibraryVMTemplateOrderByInput = "max_bandwidth_policy_ASC"

	// ContentLibraryVMTemplateOrderByInputMaxBandwidthPolicyDESC captures enum value "max_bandwidth_policy_DESC"
	ContentLibraryVMTemplateOrderByInputMaxBandwidthPolicyDESC ContentLibraryVMTemplateOrderByInput = "max_bandwidth_policy_DESC"

	// ContentLibraryVMTemplateOrderByInputMaxIopsASC captures enum value "max_iops_ASC"
	ContentLibraryVMTemplateOrderByInputMaxIopsASC ContentLibraryVMTemplateOrderByInput = "max_iops_ASC"

	// ContentLibraryVMTemplateOrderByInputMaxIopsDESC captures enum value "max_iops_DESC"
	ContentLibraryVMTemplateOrderByInputMaxIopsDESC ContentLibraryVMTemplateOrderByInput = "max_iops_DESC"

	// ContentLibraryVMTemplateOrderByInputMaxIopsPolicyASC captures enum value "max_iops_policy_ASC"
	ContentLibraryVMTemplateOrderByInputMaxIopsPolicyASC ContentLibraryVMTemplateOrderByInput = "max_iops_policy_ASC"

	// ContentLibraryVMTemplateOrderByInputMaxIopsPolicyDESC captures enum value "max_iops_policy_DESC"
	ContentLibraryVMTemplateOrderByInputMaxIopsPolicyDESC ContentLibraryVMTemplateOrderByInput = "max_iops_policy_DESC"

	// ContentLibraryVMTemplateOrderByInputMemoryASC captures enum value "memory_ASC"
	ContentLibraryVMTemplateOrderByInputMemoryASC ContentLibraryVMTemplateOrderByInput = "memory_ASC"

	// ContentLibraryVMTemplateOrderByInputMemoryDESC captures enum value "memory_DESC"
	ContentLibraryVMTemplateOrderByInputMemoryDESC ContentLibraryVMTemplateOrderByInput = "memory_DESC"

	// ContentLibraryVMTemplateOrderByInputNameASC captures enum value "name_ASC"
	ContentLibraryVMTemplateOrderByInputNameASC ContentLibraryVMTemplateOrderByInput = "name_ASC"

	// ContentLibraryVMTemplateOrderByInputNameDESC captures enum value "name_DESC"
	ContentLibraryVMTemplateOrderByInputNameDESC ContentLibraryVMTemplateOrderByInput = "name_DESC"

	// ContentLibraryVMTemplateOrderByInputOsASC captures enum value "os_ASC"
	ContentLibraryVMTemplateOrderByInputOsASC ContentLibraryVMTemplateOrderByInput = "os_ASC"

	// ContentLibraryVMTemplateOrderByInputOsDESC captures enum value "os_DESC"
	ContentLibraryVMTemplateOrderByInputOsDESC ContentLibraryVMTemplateOrderByInput = "os_DESC"

	// ContentLibraryVMTemplateOrderByInputSizeASC captures enum value "size_ASC"
	ContentLibraryVMTemplateOrderByInputSizeASC ContentLibraryVMTemplateOrderByInput = "size_ASC"

	// ContentLibraryVMTemplateOrderByInputSizeDESC captures enum value "size_DESC"
	ContentLibraryVMTemplateOrderByInputSizeDESC ContentLibraryVMTemplateOrderByInput = "size_DESC"

	// ContentLibraryVMTemplateOrderByInputTemplateConfigASC captures enum value "template_config_ASC"
	ContentLibraryVMTemplateOrderByInputTemplateConfigASC ContentLibraryVMTemplateOrderByInput = "template_config_ASC"

	// ContentLibraryVMTemplateOrderByInputTemplateConfigDESC captures enum value "template_config_DESC"
	ContentLibraryVMTemplateOrderByInputTemplateConfigDESC ContentLibraryVMTemplateOrderByInput = "template_config_DESC"

	// ContentLibraryVMTemplateOrderByInputVcpuASC captures enum value "vcpu_ASC"
	ContentLibraryVMTemplateOrderByInputVcpuASC ContentLibraryVMTemplateOrderByInput = "vcpu_ASC"

	// ContentLibraryVMTemplateOrderByInputVcpuDESC captures enum value "vcpu_DESC"
	ContentLibraryVMTemplateOrderByInputVcpuDESC ContentLibraryVMTemplateOrderByInput = "vcpu_DESC"

	// ContentLibraryVMTemplateOrderByInputVideoTypeASC captures enum value "video_type_ASC"
	ContentLibraryVMTemplateOrderByInputVideoTypeASC ContentLibraryVMTemplateOrderByInput = "video_type_ASC"

	// ContentLibraryVMTemplateOrderByInputVideoTypeDESC captures enum value "video_type_DESC"
	ContentLibraryVMTemplateOrderByInputVideoTypeDESC ContentLibraryVMTemplateOrderByInput = "video_type_DESC"

	// ContentLibraryVMTemplateOrderByInputVMDisksASC captures enum value "vm_disks_ASC"
	ContentLibraryVMTemplateOrderByInputVMDisksASC ContentLibraryVMTemplateOrderByInput = "vm_disks_ASC"

	// ContentLibraryVMTemplateOrderByInputVMDisksDESC captures enum value "vm_disks_DESC"
	ContentLibraryVMTemplateOrderByInputVMDisksDESC ContentLibraryVMTemplateOrderByInput = "vm_disks_DESC"

	// ContentLibraryVMTemplateOrderByInputVMNicsASC captures enum value "vm_nics_ASC"
	ContentLibraryVMTemplateOrderByInputVMNicsASC ContentLibraryVMTemplateOrderByInput = "vm_nics_ASC"

	// ContentLibraryVMTemplateOrderByInputVMNicsDESC captures enum value "vm_nics_DESC"
	ContentLibraryVMTemplateOrderByInputVMNicsDESC ContentLibraryVMTemplateOrderByInput = "vm_nics_DESC"

	// ContentLibraryVMTemplateOrderByInputWinOptASC captures enum value "win_opt_ASC"
	ContentLibraryVMTemplateOrderByInputWinOptASC ContentLibraryVMTemplateOrderByInput = "win_opt_ASC"

	// ContentLibraryVMTemplateOrderByInputWinOptDESC captures enum value "win_opt_DESC"
	ContentLibraryVMTemplateOrderByInputWinOptDESC ContentLibraryVMTemplateOrderByInput = "win_opt_DESC"

	// ContentLibraryVMTemplateOrderByInputZbsStorageInfoASC captures enum value "zbs_storage_info_ASC"
	ContentLibraryVMTemplateOrderByInputZbsStorageInfoASC ContentLibraryVMTemplateOrderByInput = "zbs_storage_info_ASC"

	// ContentLibraryVMTemplateOrderByInputZbsStorageInfoDESC captures enum value "zbs_storage_info_DESC"
	ContentLibraryVMTemplateOrderByInputZbsStorageInfoDESC ContentLibraryVMTemplateOrderByInput = "zbs_storage_info_DESC"
)

// for schema
var contentLibraryVmTemplateOrderByInputEnum []interface{}

func init() {
	var res []ContentLibraryVMTemplateOrderByInput
	if err := json.Unmarshal([]byte(`["architecture_ASC","architecture_DESC","clock_offset_ASC","clock_offset_DESC","cloud_init_supported_ASC","cloud_init_supported_DESC","cpu_ASC","cpu_DESC","cpu_model_ASC","cpu_model_DESC","createdAt_ASC","createdAt_DESC","description_ASC","description_DESC","entityAsyncStatus_ASC","entityAsyncStatus_DESC","firmware_ASC","firmware_DESC","ha_ASC","ha_DESC","id_ASC","id_DESC","io_policy_ASC","io_policy_DESC","max_bandwidth_ASC","max_bandwidth_DESC","max_bandwidth_policy_ASC","max_bandwidth_policy_DESC","max_iops_ASC","max_iops_DESC","max_iops_policy_ASC","max_iops_policy_DESC","memory_ASC","memory_DESC","name_ASC","name_DESC","os_ASC","os_DESC","size_ASC","size_DESC","template_config_ASC","template_config_DESC","vcpu_ASC","vcpu_DESC","video_type_ASC","video_type_DESC","vm_disks_ASC","vm_disks_DESC","vm_nics_ASC","vm_nics_DESC","win_opt_ASC","win_opt_DESC","zbs_storage_info_ASC","zbs_storage_info_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		contentLibraryVmTemplateOrderByInputEnum = append(contentLibraryVmTemplateOrderByInputEnum, v)
	}
}

func (m ContentLibraryVMTemplateOrderByInput) validateContentLibraryVMTemplateOrderByInputEnum(path, location string, value ContentLibraryVMTemplateOrderByInput) error {
	if err := validate.EnumCase(path, location, value, contentLibraryVmTemplateOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this content library Vm template order by input
func (m ContentLibraryVMTemplateOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateContentLibraryVMTemplateOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this content library Vm template order by input based on context it is used
func (m ContentLibraryVMTemplateOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
