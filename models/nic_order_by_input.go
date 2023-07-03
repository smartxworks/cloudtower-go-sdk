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

// NicOrderByInput nic order by input
//
// swagger:model NicOrderByInput
type NicOrderByInput string

func NewNicOrderByInput(value NicOrderByInput) *NicOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated NicOrderByInput.
func (m NicOrderByInput) Pointer() *NicOrderByInput {
	return &m
}

const (

	// NicOrderByInputDriverASC captures enum value "driver_ASC"
	NicOrderByInputDriverASC NicOrderByInput = "driver_ASC"

	// NicOrderByInputDriverDESC captures enum value "driver_DESC"
	NicOrderByInputDriverDESC NicOrderByInput = "driver_DESC"

	// NicOrderByInputDriverStateASC captures enum value "driver_state_ASC"
	NicOrderByInputDriverStateASC NicOrderByInput = "driver_state_ASC"

	// NicOrderByInputDriverStateDESC captures enum value "driver_state_DESC"
	NicOrderByInputDriverStateDESC NicOrderByInput = "driver_state_DESC"

	// NicOrderByInputEntityAsyncStatusASC captures enum value "entityAsyncStatus_ASC"
	NicOrderByInputEntityAsyncStatusASC NicOrderByInput = "entityAsyncStatus_ASC"

	// NicOrderByInputEntityAsyncStatusDESC captures enum value "entityAsyncStatus_DESC"
	NicOrderByInputEntityAsyncStatusDESC NicOrderByInput = "entityAsyncStatus_DESC"

	// NicOrderByInputGatewayIPASC captures enum value "gateway_ip_ASC"
	NicOrderByInputGatewayIPASC NicOrderByInput = "gateway_ip_ASC"

	// NicOrderByInputGatewayIPDESC captures enum value "gateway_ip_DESC"
	NicOrderByInputGatewayIPDESC NicOrderByInput = "gateway_ip_DESC"

	// NicOrderByInputIbdevASC captures enum value "ibdev_ASC"
	NicOrderByInputIbdevASC NicOrderByInput = "ibdev_ASC"

	// NicOrderByInputIbdevDESC captures enum value "ibdev_DESC"
	NicOrderByInputIbdevDESC NicOrderByInput = "ibdev_DESC"

	// NicOrderByInputIDASC captures enum value "id_ASC"
	NicOrderByInputIDASC NicOrderByInput = "id_ASC"

	// NicOrderByInputIDDESC captures enum value "id_DESC"
	NicOrderByInputIDDESC NicOrderByInput = "id_DESC"

	// NicOrderByInputIPAddressASC captures enum value "ip_address_ASC"
	NicOrderByInputIPAddressASC NicOrderByInput = "ip_address_ASC"

	// NicOrderByInputIPAddressDESC captures enum value "ip_address_DESC"
	NicOrderByInputIPAddressDESC NicOrderByInput = "ip_address_DESC"

	// NicOrderByInputIsSriovASC captures enum value "is_sriov_ASC"
	NicOrderByInputIsSriovASC NicOrderByInput = "is_sriov_ASC"

	// NicOrderByInputIsSriovDESC captures enum value "is_sriov_DESC"
	NicOrderByInputIsSriovDESC NicOrderByInput = "is_sriov_DESC"

	// NicOrderByInputLocalIDASC captures enum value "local_id_ASC"
	NicOrderByInputLocalIDASC NicOrderByInput = "local_id_ASC"

	// NicOrderByInputLocalIDDESC captures enum value "local_id_DESC"
	NicOrderByInputLocalIDDESC NicOrderByInput = "local_id_DESC"

	// NicOrderByInputMacAddressASC captures enum value "mac_address_ASC"
	NicOrderByInputMacAddressASC NicOrderByInput = "mac_address_ASC"

	// NicOrderByInputMacAddressDESC captures enum value "mac_address_DESC"
	NicOrderByInputMacAddressDESC NicOrderByInput = "mac_address_DESC"

	// NicOrderByInputMaxVfNumASC captures enum value "max_vf_num_ASC"
	NicOrderByInputMaxVfNumASC NicOrderByInput = "max_vf_num_ASC"

	// NicOrderByInputMaxVfNumDESC captures enum value "max_vf_num_DESC"
	NicOrderByInputMaxVfNumDESC NicOrderByInput = "max_vf_num_DESC"

	// NicOrderByInputModelASC captures enum value "model_ASC"
	NicOrderByInputModelASC NicOrderByInput = "model_ASC"

	// NicOrderByInputModelDESC captures enum value "model_DESC"
	NicOrderByInputModelDESC NicOrderByInput = "model_DESC"

	// NicOrderByInputMtuASC captures enum value "mtu_ASC"
	NicOrderByInputMtuASC NicOrderByInput = "mtu_ASC"

	// NicOrderByInputMtuDESC captures enum value "mtu_DESC"
	NicOrderByInputMtuDESC NicOrderByInput = "mtu_DESC"

	// NicOrderByInputNameASC captures enum value "name_ASC"
	NicOrderByInputNameASC NicOrderByInput = "name_ASC"

	// NicOrderByInputNameDESC captures enum value "name_DESC"
	NicOrderByInputNameDESC NicOrderByInput = "name_DESC"

	// NicOrderByInputNicUUIDASC captures enum value "nic_uuid_ASC"
	NicOrderByInputNicUUIDASC NicOrderByInput = "nic_uuid_ASC"

	// NicOrderByInputNicUUIDDESC captures enum value "nic_uuid_DESC"
	NicOrderByInputNicUUIDDESC NicOrderByInput = "nic_uuid_DESC"

	// NicOrderByInputPhysicalASC captures enum value "physical_ASC"
	NicOrderByInputPhysicalASC NicOrderByInput = "physical_ASC"

	// NicOrderByInputPhysicalDESC captures enum value "physical_DESC"
	NicOrderByInputPhysicalDESC NicOrderByInput = "physical_DESC"

	// NicOrderByInputRdmaEnabledASC captures enum value "rdma_enabled_ASC"
	NicOrderByInputRdmaEnabledASC NicOrderByInput = "rdma_enabled_ASC"

	// NicOrderByInputRdmaEnabledDESC captures enum value "rdma_enabled_DESC"
	NicOrderByInputRdmaEnabledDESC NicOrderByInput = "rdma_enabled_DESC"

	// NicOrderByInputRunningASC captures enum value "running_ASC"
	NicOrderByInputRunningASC NicOrderByInput = "running_ASC"

	// NicOrderByInputRunningDESC captures enum value "running_DESC"
	NicOrderByInputRunningDESC NicOrderByInput = "running_DESC"

	// NicOrderByInputSpeedASC captures enum value "speed_ASC"
	NicOrderByInputSpeedASC NicOrderByInput = "speed_ASC"

	// NicOrderByInputSpeedDESC captures enum value "speed_DESC"
	NicOrderByInputSpeedDESC NicOrderByInput = "speed_DESC"

	// NicOrderByInputSubnetMaskASC captures enum value "subnet_mask_ASC"
	NicOrderByInputSubnetMaskASC NicOrderByInput = "subnet_mask_ASC"

	// NicOrderByInputSubnetMaskDESC captures enum value "subnet_mask_DESC"
	NicOrderByInputSubnetMaskDESC NicOrderByInput = "subnet_mask_DESC"

	// NicOrderByInputTotalVfNumASC captures enum value "total_vf_num_ASC"
	NicOrderByInputTotalVfNumASC NicOrderByInput = "total_vf_num_ASC"

	// NicOrderByInputTotalVfNumDESC captures enum value "total_vf_num_DESC"
	NicOrderByInputTotalVfNumDESC NicOrderByInput = "total_vf_num_DESC"

	// NicOrderByInputTypeASC captures enum value "type_ASC"
	NicOrderByInputTypeASC NicOrderByInput = "type_ASC"

	// NicOrderByInputTypeDESC captures enum value "type_DESC"
	NicOrderByInputTypeDESC NicOrderByInput = "type_DESC"

	// NicOrderByInputUpASC captures enum value "up_ASC"
	NicOrderByInputUpASC NicOrderByInput = "up_ASC"

	// NicOrderByInputUpDESC captures enum value "up_DESC"
	NicOrderByInputUpDESC NicOrderByInput = "up_DESC"

	// NicOrderByInputUsedVfNumASC captures enum value "used_vf_num_ASC"
	NicOrderByInputUsedVfNumASC NicOrderByInput = "used_vf_num_ASC"

	// NicOrderByInputUsedVfNumDESC captures enum value "used_vf_num_DESC"
	NicOrderByInputUsedVfNumDESC NicOrderByInput = "used_vf_num_DESC"
)

// for schema
var nicOrderByInputEnum []interface{}

func init() {
	var res []NicOrderByInput
	if err := json.Unmarshal([]byte(`["driver_ASC","driver_DESC","driver_state_ASC","driver_state_DESC","entityAsyncStatus_ASC","entityAsyncStatus_DESC","gateway_ip_ASC","gateway_ip_DESC","ibdev_ASC","ibdev_DESC","id_ASC","id_DESC","ip_address_ASC","ip_address_DESC","is_sriov_ASC","is_sriov_DESC","local_id_ASC","local_id_DESC","mac_address_ASC","mac_address_DESC","max_vf_num_ASC","max_vf_num_DESC","model_ASC","model_DESC","mtu_ASC","mtu_DESC","name_ASC","name_DESC","nic_uuid_ASC","nic_uuid_DESC","physical_ASC","physical_DESC","rdma_enabled_ASC","rdma_enabled_DESC","running_ASC","running_DESC","speed_ASC","speed_DESC","subnet_mask_ASC","subnet_mask_DESC","total_vf_num_ASC","total_vf_num_DESC","type_ASC","type_DESC","up_ASC","up_DESC","used_vf_num_ASC","used_vf_num_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nicOrderByInputEnum = append(nicOrderByInputEnum, v)
	}
}

func (m NicOrderByInput) validateNicOrderByInputEnum(path, location string, value NicOrderByInput) error {
	if err := validate.EnumCase(path, location, value, nicOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this nic order by input
func (m NicOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateNicOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this nic order by input based on context it is used
func (m NicOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
