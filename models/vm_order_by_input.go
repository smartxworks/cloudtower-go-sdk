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

// VMOrderByInput Vm order by input
//
// swagger:model VmOrderByInput
type VMOrderByInput string

func NewVMOrderByInput(value VMOrderByInput) *VMOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VMOrderByInput.
func (m VMOrderByInput) Pointer() *VMOrderByInput {
	return &m
}

const (

	// VMOrderByInputBiosUUIDASC captures enum value "bios_uuid_ASC"
	VMOrderByInputBiosUUIDASC VMOrderByInput = "bios_uuid_ASC"

	// VMOrderByInputBiosUUIDDESC captures enum value "bios_uuid_DESC"
	VMOrderByInputBiosUUIDDESC VMOrderByInput = "bios_uuid_DESC"

	// VMOrderByInputClockOffsetASC captures enum value "clock_offset_ASC"
	VMOrderByInputClockOffsetASC VMOrderByInput = "clock_offset_ASC"

	// VMOrderByInputClockOffsetDESC captures enum value "clock_offset_DESC"
	VMOrderByInputClockOffsetDESC VMOrderByInput = "clock_offset_DESC"

	// VMOrderByInputCloudInitSupportedASC captures enum value "cloud_init_supported_ASC"
	VMOrderByInputCloudInitSupportedASC VMOrderByInput = "cloud_init_supported_ASC"

	// VMOrderByInputCloudInitSupportedDESC captures enum value "cloud_init_supported_DESC"
	VMOrderByInputCloudInitSupportedDESC VMOrderByInput = "cloud_init_supported_DESC"

	// VMOrderByInputCPUASC captures enum value "cpu_ASC"
	VMOrderByInputCPUASC VMOrderByInput = "cpu_ASC"

	// VMOrderByInputCPUDESC captures enum value "cpu_DESC"
	VMOrderByInputCPUDESC VMOrderByInput = "cpu_DESC"

	// VMOrderByInputCPUModelASC captures enum value "cpu_model_ASC"
	VMOrderByInputCPUModelASC VMOrderByInput = "cpu_model_ASC"

	// VMOrderByInputCPUModelDESC captures enum value "cpu_model_DESC"
	VMOrderByInputCPUModelDESC VMOrderByInput = "cpu_model_DESC"

	// VMOrderByInputCPUUsageASC captures enum value "cpu_usage_ASC"
	VMOrderByInputCPUUsageASC VMOrderByInput = "cpu_usage_ASC"

	// VMOrderByInputCPUUsageDESC captures enum value "cpu_usage_DESC"
	VMOrderByInputCPUUsageDESC VMOrderByInput = "cpu_usage_DESC"

	// VMOrderByInputDeletedAtASC captures enum value "deleted_at_ASC"
	VMOrderByInputDeletedAtASC VMOrderByInput = "deleted_at_ASC"

	// VMOrderByInputDeletedAtDESC captures enum value "deleted_at_DESC"
	VMOrderByInputDeletedAtDESC VMOrderByInput = "deleted_at_DESC"

	// VMOrderByInputDescriptionASC captures enum value "description_ASC"
	VMOrderByInputDescriptionASC VMOrderByInput = "description_ASC"

	// VMOrderByInputDescriptionDESC captures enum value "description_DESC"
	VMOrderByInputDescriptionDESC VMOrderByInput = "description_DESC"

	// VMOrderByInputDNSServersASC captures enum value "dns_servers_ASC"
	VMOrderByInputDNSServersASC VMOrderByInput = "dns_servers_ASC"

	// VMOrderByInputDNSServersDESC captures enum value "dns_servers_DESC"
	VMOrderByInputDNSServersDESC VMOrderByInput = "dns_servers_DESC"

	// VMOrderByInputEntityAsyncStatusASC captures enum value "entityAsyncStatus_ASC"
	VMOrderByInputEntityAsyncStatusASC VMOrderByInput = "entityAsyncStatus_ASC"

	// VMOrderByInputEntityAsyncStatusDESC captures enum value "entityAsyncStatus_DESC"
	VMOrderByInputEntityAsyncStatusDESC VMOrderByInput = "entityAsyncStatus_DESC"

	// VMOrderByInputFirmwareASC captures enum value "firmware_ASC"
	VMOrderByInputFirmwareASC VMOrderByInput = "firmware_ASC"

	// VMOrderByInputFirmwareDESC captures enum value "firmware_DESC"
	VMOrderByInputFirmwareDESC VMOrderByInput = "firmware_DESC"

	// VMOrderByInputGuestCPUModelASC captures enum value "guest_cpu_model_ASC"
	VMOrderByInputGuestCPUModelASC VMOrderByInput = "guest_cpu_model_ASC"

	// VMOrderByInputGuestCPUModelDESC captures enum value "guest_cpu_model_DESC"
	VMOrderByInputGuestCPUModelDESC VMOrderByInput = "guest_cpu_model_DESC"

	// VMOrderByInputGuestOsTypeASC captures enum value "guest_os_type_ASC"
	VMOrderByInputGuestOsTypeASC VMOrderByInput = "guest_os_type_ASC"

	// VMOrderByInputGuestOsTypeDESC captures enum value "guest_os_type_DESC"
	VMOrderByInputGuestOsTypeDESC VMOrderByInput = "guest_os_type_DESC"

	// VMOrderByInputGuestSizeUsageASC captures enum value "guest_size_usage_ASC"
	VMOrderByInputGuestSizeUsageASC VMOrderByInput = "guest_size_usage_ASC"

	// VMOrderByInputGuestSizeUsageDESC captures enum value "guest_size_usage_DESC"
	VMOrderByInputGuestSizeUsageDESC VMOrderByInput = "guest_size_usage_DESC"

	// VMOrderByInputGuestUsedSizeASC captures enum value "guest_used_size_ASC"
	VMOrderByInputGuestUsedSizeASC VMOrderByInput = "guest_used_size_ASC"

	// VMOrderByInputGuestUsedSizeDESC captures enum value "guest_used_size_DESC"
	VMOrderByInputGuestUsedSizeDESC VMOrderByInput = "guest_used_size_DESC"

	// VMOrderByInputHaASC captures enum value "ha_ASC"
	VMOrderByInputHaASC VMOrderByInput = "ha_ASC"

	// VMOrderByInputHaDESC captures enum value "ha_DESC"
	VMOrderByInputHaDESC VMOrderByInput = "ha_DESC"

	// VMOrderByInputHostnameASC captures enum value "hostname_ASC"
	VMOrderByInputHostnameASC VMOrderByInput = "hostname_ASC"

	// VMOrderByInputHostnameDESC captures enum value "hostname_DESC"
	VMOrderByInputHostnameDESC VMOrderByInput = "hostname_DESC"

	// VMOrderByInputIDASC captures enum value "id_ASC"
	VMOrderByInputIDASC VMOrderByInput = "id_ASC"

	// VMOrderByInputIDDESC captures enum value "id_DESC"
	VMOrderByInputIDDESC VMOrderByInput = "id_DESC"

	// VMOrderByInputInRecycleBinASC captures enum value "in_recycle_bin_ASC"
	VMOrderByInputInRecycleBinASC VMOrderByInput = "in_recycle_bin_ASC"

	// VMOrderByInputInRecycleBinDESC captures enum value "in_recycle_bin_DESC"
	VMOrderByInputInRecycleBinDESC VMOrderByInput = "in_recycle_bin_DESC"

	// VMOrderByInputInternalASC captures enum value "internal_ASC"
	VMOrderByInputInternalASC VMOrderByInput = "internal_ASC"

	// VMOrderByInputInternalDESC captures enum value "internal_DESC"
	VMOrderByInputInternalDESC VMOrderByInput = "internal_DESC"

	// VMOrderByInputIoPolicyASC captures enum value "io_policy_ASC"
	VMOrderByInputIoPolicyASC VMOrderByInput = "io_policy_ASC"

	// VMOrderByInputIoPolicyDESC captures enum value "io_policy_DESC"
	VMOrderByInputIoPolicyDESC VMOrderByInput = "io_policy_DESC"

	// VMOrderByInputIpsASC captures enum value "ips_ASC"
	VMOrderByInputIpsASC VMOrderByInput = "ips_ASC"

	// VMOrderByInputIpsDESC captures enum value "ips_DESC"
	VMOrderByInputIpsDESC VMOrderByInput = "ips_DESC"

	// VMOrderByInputKernelInfoASC captures enum value "kernel_info_ASC"
	VMOrderByInputKernelInfoASC VMOrderByInput = "kernel_info_ASC"

	// VMOrderByInputKernelInfoDESC captures enum value "kernel_info_DESC"
	VMOrderByInputKernelInfoDESC VMOrderByInput = "kernel_info_DESC"

	// VMOrderByInputLastShutdownTimeASC captures enum value "last_shutdown_time_ASC"
	VMOrderByInputLastShutdownTimeASC VMOrderByInput = "last_shutdown_time_ASC"

	// VMOrderByInputLastShutdownTimeDESC captures enum value "last_shutdown_time_DESC"
	VMOrderByInputLastShutdownTimeDESC VMOrderByInput = "last_shutdown_time_DESC"

	// VMOrderByInputLocalCreatedAtASC captures enum value "local_created_at_ASC"
	VMOrderByInputLocalCreatedAtASC VMOrderByInput = "local_created_at_ASC"

	// VMOrderByInputLocalCreatedAtDESC captures enum value "local_created_at_DESC"
	VMOrderByInputLocalCreatedAtDESC VMOrderByInput = "local_created_at_DESC"

	// VMOrderByInputLocalIDASC captures enum value "local_id_ASC"
	VMOrderByInputLocalIDASC VMOrderByInput = "local_id_ASC"

	// VMOrderByInputLocalIDDESC captures enum value "local_id_DESC"
	VMOrderByInputLocalIDDESC VMOrderByInput = "local_id_DESC"

	// VMOrderByInputLogicalSizeBytesASC captures enum value "logical_size_bytes_ASC"
	VMOrderByInputLogicalSizeBytesASC VMOrderByInput = "logical_size_bytes_ASC"

	// VMOrderByInputLogicalSizeBytesDESC captures enum value "logical_size_bytes_DESC"
	VMOrderByInputLogicalSizeBytesDESC VMOrderByInput = "logical_size_bytes_DESC"

	// VMOrderByInputMaxBandwidthASC captures enum value "max_bandwidth_ASC"
	VMOrderByInputMaxBandwidthASC VMOrderByInput = "max_bandwidth_ASC"

	// VMOrderByInputMaxBandwidthDESC captures enum value "max_bandwidth_DESC"
	VMOrderByInputMaxBandwidthDESC VMOrderByInput = "max_bandwidth_DESC"

	// VMOrderByInputMaxBandwidthPolicyASC captures enum value "max_bandwidth_policy_ASC"
	VMOrderByInputMaxBandwidthPolicyASC VMOrderByInput = "max_bandwidth_policy_ASC"

	// VMOrderByInputMaxBandwidthPolicyDESC captures enum value "max_bandwidth_policy_DESC"
	VMOrderByInputMaxBandwidthPolicyDESC VMOrderByInput = "max_bandwidth_policy_DESC"

	// VMOrderByInputMaxIopsASC captures enum value "max_iops_ASC"
	VMOrderByInputMaxIopsASC VMOrderByInput = "max_iops_ASC"

	// VMOrderByInputMaxIopsDESC captures enum value "max_iops_DESC"
	VMOrderByInputMaxIopsDESC VMOrderByInput = "max_iops_DESC"

	// VMOrderByInputMaxIopsPolicyASC captures enum value "max_iops_policy_ASC"
	VMOrderByInputMaxIopsPolicyASC VMOrderByInput = "max_iops_policy_ASC"

	// VMOrderByInputMaxIopsPolicyDESC captures enum value "max_iops_policy_DESC"
	VMOrderByInputMaxIopsPolicyDESC VMOrderByInput = "max_iops_policy_DESC"

	// VMOrderByInputMemoryASC captures enum value "memory_ASC"
	VMOrderByInputMemoryASC VMOrderByInput = "memory_ASC"

	// VMOrderByInputMemoryDESC captures enum value "memory_DESC"
	VMOrderByInputMemoryDESC VMOrderByInput = "memory_DESC"

	// VMOrderByInputMemoryUsageASC captures enum value "memory_usage_ASC"
	VMOrderByInputMemoryUsageASC VMOrderByInput = "memory_usage_ASC"

	// VMOrderByInputMemoryUsageDESC captures enum value "memory_usage_DESC"
	VMOrderByInputMemoryUsageDESC VMOrderByInput = "memory_usage_DESC"

	// VMOrderByInputNameASC captures enum value "name_ASC"
	VMOrderByInputNameASC VMOrderByInput = "name_ASC"

	// VMOrderByInputNameDESC captures enum value "name_DESC"
	VMOrderByInputNameDESC VMOrderByInput = "name_DESC"

	// VMOrderByInputNestedVirtualizationASC captures enum value "nested_virtualization_ASC"
	VMOrderByInputNestedVirtualizationASC VMOrderByInput = "nested_virtualization_ASC"

	// VMOrderByInputNestedVirtualizationDESC captures enum value "nested_virtualization_DESC"
	VMOrderByInputNestedVirtualizationDESC VMOrderByInput = "nested_virtualization_DESC"

	// VMOrderByInputNodeIPASC captures enum value "node_ip_ASC"
	VMOrderByInputNodeIPASC VMOrderByInput = "node_ip_ASC"

	// VMOrderByInputNodeIPDESC captures enum value "node_ip_DESC"
	VMOrderByInputNodeIPDESC VMOrderByInput = "node_ip_DESC"

	// VMOrderByInputOriginalNameASC captures enum value "original_name_ASC"
	VMOrderByInputOriginalNameASC VMOrderByInput = "original_name_ASC"

	// VMOrderByInputOriginalNameDESC captures enum value "original_name_DESC"
	VMOrderByInputOriginalNameDESC VMOrderByInput = "original_name_DESC"

	// VMOrderByInputOsASC captures enum value "os_ASC"
	VMOrderByInputOsASC VMOrderByInput = "os_ASC"

	// VMOrderByInputOsDESC captures enum value "os_DESC"
	VMOrderByInputOsDESC VMOrderByInput = "os_DESC"

	// VMOrderByInputProtectedASC captures enum value "protected_ASC"
	VMOrderByInputProtectedASC VMOrderByInput = "protected_ASC"

	// VMOrderByInputProtectedDESC captures enum value "protected_DESC"
	VMOrderByInputProtectedDESC VMOrderByInput = "protected_DESC"

	// VMOrderByInputProvisionedSizeASC captures enum value "provisioned_size_ASC"
	VMOrderByInputProvisionedSizeASC VMOrderByInput = "provisioned_size_ASC"

	// VMOrderByInputProvisionedSizeDESC captures enum value "provisioned_size_DESC"
	VMOrderByInputProvisionedSizeDESC VMOrderByInput = "provisioned_size_DESC"

	// VMOrderByInputSizeASC captures enum value "size_ASC"
	VMOrderByInputSizeASC VMOrderByInput = "size_ASC"

	// VMOrderByInputSizeDESC captures enum value "size_DESC"
	VMOrderByInputSizeDESC VMOrderByInput = "size_DESC"

	// VMOrderByInputStatusASC captures enum value "status_ASC"
	VMOrderByInputStatusASC VMOrderByInput = "status_ASC"

	// VMOrderByInputStatusDESC captures enum value "status_DESC"
	VMOrderByInputStatusDESC VMOrderByInput = "status_DESC"

	// VMOrderByInputUniqueLogicalSizeASC captures enum value "unique_logical_size_ASC"
	VMOrderByInputUniqueLogicalSizeASC VMOrderByInput = "unique_logical_size_ASC"

	// VMOrderByInputUniqueLogicalSizeDESC captures enum value "unique_logical_size_DESC"
	VMOrderByInputUniqueLogicalSizeDESC VMOrderByInput = "unique_logical_size_DESC"

	// VMOrderByInputUniqueSizeASC captures enum value "unique_size_ASC"
	VMOrderByInputUniqueSizeASC VMOrderByInput = "unique_size_ASC"

	// VMOrderByInputUniqueSizeDESC captures enum value "unique_size_DESC"
	VMOrderByInputUniqueSizeDESC VMOrderByInput = "unique_size_DESC"

	// VMOrderByInputUsedSizeASC captures enum value "used_size_ASC"
	VMOrderByInputUsedSizeASC VMOrderByInput = "used_size_ASC"

	// VMOrderByInputUsedSizeDESC captures enum value "used_size_DESC"
	VMOrderByInputUsedSizeDESC VMOrderByInput = "used_size_DESC"

	// VMOrderByInputUsedSizeUsageASC captures enum value "used_size_usage_ASC"
	VMOrderByInputUsedSizeUsageASC VMOrderByInput = "used_size_usage_ASC"

	// VMOrderByInputUsedSizeUsageDESC captures enum value "used_size_usage_DESC"
	VMOrderByInputUsedSizeUsageDESC VMOrderByInput = "used_size_usage_DESC"

	// VMOrderByInputVcpuASC captures enum value "vcpu_ASC"
	VMOrderByInputVcpuASC VMOrderByInput = "vcpu_ASC"

	// VMOrderByInputVcpuDESC captures enum value "vcpu_DESC"
	VMOrderByInputVcpuDESC VMOrderByInput = "vcpu_DESC"

	// VMOrderByInputVideoTypeASC captures enum value "video_type_ASC"
	VMOrderByInputVideoTypeASC VMOrderByInput = "video_type_ASC"

	// VMOrderByInputVideoTypeDESC captures enum value "video_type_DESC"
	VMOrderByInputVideoTypeDESC VMOrderByInput = "video_type_DESC"

	// VMOrderByInputVMToolsStatusASC captures enum value "vm_tools_status_ASC"
	VMOrderByInputVMToolsStatusASC VMOrderByInput = "vm_tools_status_ASC"

	// VMOrderByInputVMToolsStatusDESC captures enum value "vm_tools_status_DESC"
	VMOrderByInputVMToolsStatusDESC VMOrderByInput = "vm_tools_status_DESC"

	// VMOrderByInputVMToolsVersionASC captures enum value "vm_tools_version_ASC"
	VMOrderByInputVMToolsVersionASC VMOrderByInput = "vm_tools_version_ASC"

	// VMOrderByInputVMToolsVersionDESC captures enum value "vm_tools_version_DESC"
	VMOrderByInputVMToolsVersionDESC VMOrderByInput = "vm_tools_version_DESC"

	// VMOrderByInputVMUsageASC captures enum value "vm_usage_ASC"
	VMOrderByInputVMUsageASC VMOrderByInput = "vm_usage_ASC"

	// VMOrderByInputVMUsageDESC captures enum value "vm_usage_DESC"
	VMOrderByInputVMUsageDESC VMOrderByInput = "vm_usage_DESC"

	// VMOrderByInputWinOptASC captures enum value "win_opt_ASC"
	VMOrderByInputWinOptASC VMOrderByInput = "win_opt_ASC"

	// VMOrderByInputWinOptDESC captures enum value "win_opt_DESC"
	VMOrderByInputWinOptDESC VMOrderByInput = "win_opt_DESC"
)

// for schema
var vmOrderByInputEnum []interface{}

func init() {
	var res []VMOrderByInput
	if err := json.Unmarshal([]byte(`["bios_uuid_ASC","bios_uuid_DESC","clock_offset_ASC","clock_offset_DESC","cloud_init_supported_ASC","cloud_init_supported_DESC","cpu_ASC","cpu_DESC","cpu_model_ASC","cpu_model_DESC","cpu_usage_ASC","cpu_usage_DESC","deleted_at_ASC","deleted_at_DESC","description_ASC","description_DESC","dns_servers_ASC","dns_servers_DESC","entityAsyncStatus_ASC","entityAsyncStatus_DESC","firmware_ASC","firmware_DESC","guest_cpu_model_ASC","guest_cpu_model_DESC","guest_os_type_ASC","guest_os_type_DESC","guest_size_usage_ASC","guest_size_usage_DESC","guest_used_size_ASC","guest_used_size_DESC","ha_ASC","ha_DESC","hostname_ASC","hostname_DESC","id_ASC","id_DESC","in_recycle_bin_ASC","in_recycle_bin_DESC","internal_ASC","internal_DESC","io_policy_ASC","io_policy_DESC","ips_ASC","ips_DESC","kernel_info_ASC","kernel_info_DESC","last_shutdown_time_ASC","last_shutdown_time_DESC","local_created_at_ASC","local_created_at_DESC","local_id_ASC","local_id_DESC","logical_size_bytes_ASC","logical_size_bytes_DESC","max_bandwidth_ASC","max_bandwidth_DESC","max_bandwidth_policy_ASC","max_bandwidth_policy_DESC","max_iops_ASC","max_iops_DESC","max_iops_policy_ASC","max_iops_policy_DESC","memory_ASC","memory_DESC","memory_usage_ASC","memory_usage_DESC","name_ASC","name_DESC","nested_virtualization_ASC","nested_virtualization_DESC","node_ip_ASC","node_ip_DESC","original_name_ASC","original_name_DESC","os_ASC","os_DESC","protected_ASC","protected_DESC","provisioned_size_ASC","provisioned_size_DESC","size_ASC","size_DESC","status_ASC","status_DESC","unique_logical_size_ASC","unique_logical_size_DESC","unique_size_ASC","unique_size_DESC","used_size_ASC","used_size_DESC","used_size_usage_ASC","used_size_usage_DESC","vcpu_ASC","vcpu_DESC","video_type_ASC","video_type_DESC","vm_tools_status_ASC","vm_tools_status_DESC","vm_tools_version_ASC","vm_tools_version_DESC","vm_usage_ASC","vm_usage_DESC","win_opt_ASC","win_opt_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vmOrderByInputEnum = append(vmOrderByInputEnum, v)
	}
}

func (m VMOrderByInput) validateVMOrderByInputEnum(path, location string, value VMOrderByInput) error {
	if err := validate.EnumCase(path, location, value, vmOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this Vm order by input
func (m VMOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateVMOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this Vm order by input based on context it is used
func (m VMOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
