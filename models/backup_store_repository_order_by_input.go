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

// BackupStoreRepositoryOrderByInput backup store repository order by input
//
// swagger:model BackupStoreRepositoryOrderByInput
type BackupStoreRepositoryOrderByInput string

func NewBackupStoreRepositoryOrderByInput(value BackupStoreRepositoryOrderByInput) *BackupStoreRepositoryOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated BackupStoreRepositoryOrderByInput.
func (m BackupStoreRepositoryOrderByInput) Pointer() *BackupStoreRepositoryOrderByInput {
	return &m
}

const (

	// BackupStoreRepositoryOrderByInputCreatedAtASC captures enum value "createdAt_ASC"
	BackupStoreRepositoryOrderByInputCreatedAtASC BackupStoreRepositoryOrderByInput = "createdAt_ASC"

	// BackupStoreRepositoryOrderByInputCreatedAtDESC captures enum value "createdAt_DESC"
	BackupStoreRepositoryOrderByInputCreatedAtDESC BackupStoreRepositoryOrderByInput = "createdAt_DESC"

	// BackupStoreRepositoryOrderByInputDescriptionASC captures enum value "description_ASC"
	BackupStoreRepositoryOrderByInputDescriptionASC BackupStoreRepositoryOrderByInput = "description_ASC"

	// BackupStoreRepositoryOrderByInputDescriptionDESC captures enum value "description_DESC"
	BackupStoreRepositoryOrderByInputDescriptionDESC BackupStoreRepositoryOrderByInput = "description_DESC"

	// BackupStoreRepositoryOrderByInputEntityAsyncStatusASC captures enum value "entityAsyncStatus_ASC"
	BackupStoreRepositoryOrderByInputEntityAsyncStatusASC BackupStoreRepositoryOrderByInput = "entityAsyncStatus_ASC"

	// BackupStoreRepositoryOrderByInputEntityAsyncStatusDESC captures enum value "entityAsyncStatus_DESC"
	BackupStoreRepositoryOrderByInputEntityAsyncStatusDESC BackupStoreRepositoryOrderByInput = "entityAsyncStatus_DESC"

	// BackupStoreRepositoryOrderByInputErrorCodeASC captures enum value "error_code_ASC"
	BackupStoreRepositoryOrderByInputErrorCodeASC BackupStoreRepositoryOrderByInput = "error_code_ASC"

	// BackupStoreRepositoryOrderByInputErrorCodeDESC captures enum value "error_code_DESC"
	BackupStoreRepositoryOrderByInputErrorCodeDESC BackupStoreRepositoryOrderByInput = "error_code_DESC"

	// BackupStoreRepositoryOrderByInputIDASC captures enum value "id_ASC"
	BackupStoreRepositoryOrderByInputIDASC BackupStoreRepositoryOrderByInput = "id_ASC"

	// BackupStoreRepositoryOrderByInputIDDESC captures enum value "id_DESC"
	BackupStoreRepositoryOrderByInputIDDESC BackupStoreRepositoryOrderByInput = "id_DESC"

	// BackupStoreRepositoryOrderByInputIscsiChapNameASC captures enum value "iscsi_chap_name_ASC"
	BackupStoreRepositoryOrderByInputIscsiChapNameASC BackupStoreRepositoryOrderByInput = "iscsi_chap_name_ASC"

	// BackupStoreRepositoryOrderByInputIscsiChapNameDESC captures enum value "iscsi_chap_name_DESC"
	BackupStoreRepositoryOrderByInputIscsiChapNameDESC BackupStoreRepositoryOrderByInput = "iscsi_chap_name_DESC"

	// BackupStoreRepositoryOrderByInputIscsiChapSecretASC captures enum value "iscsi_chap_secret_ASC"
	BackupStoreRepositoryOrderByInputIscsiChapSecretASC BackupStoreRepositoryOrderByInput = "iscsi_chap_secret_ASC"

	// BackupStoreRepositoryOrderByInputIscsiChapSecretDESC captures enum value "iscsi_chap_secret_DESC"
	BackupStoreRepositoryOrderByInputIscsiChapSecretDESC BackupStoreRepositoryOrderByInput = "iscsi_chap_secret_DESC"

	// BackupStoreRepositoryOrderByInputIscsiIPASC captures enum value "iscsi_ip_ASC"
	BackupStoreRepositoryOrderByInputIscsiIPASC BackupStoreRepositoryOrderByInput = "iscsi_ip_ASC"

	// BackupStoreRepositoryOrderByInputIscsiIPDESC captures enum value "iscsi_ip_DESC"
	BackupStoreRepositoryOrderByInputIscsiIPDESC BackupStoreRepositoryOrderByInput = "iscsi_ip_DESC"

	// BackupStoreRepositoryOrderByInputIscsiLunIDASC captures enum value "iscsi_lun_id_ASC"
	BackupStoreRepositoryOrderByInputIscsiLunIDASC BackupStoreRepositoryOrderByInput = "iscsi_lun_id_ASC"

	// BackupStoreRepositoryOrderByInputIscsiLunIDDESC captures enum value "iscsi_lun_id_DESC"
	BackupStoreRepositoryOrderByInputIscsiLunIDDESC BackupStoreRepositoryOrderByInput = "iscsi_lun_id_DESC"

	// BackupStoreRepositoryOrderByInputIscsiPortASC captures enum value "iscsi_port_ASC"
	BackupStoreRepositoryOrderByInputIscsiPortASC BackupStoreRepositoryOrderByInput = "iscsi_port_ASC"

	// BackupStoreRepositoryOrderByInputIscsiPortDESC captures enum value "iscsi_port_DESC"
	BackupStoreRepositoryOrderByInputIscsiPortDESC BackupStoreRepositoryOrderByInput = "iscsi_port_DESC"

	// BackupStoreRepositoryOrderByInputIscsiTargetIqnASC captures enum value "iscsi_target_iqn_ASC"
	BackupStoreRepositoryOrderByInputIscsiTargetIqnASC BackupStoreRepositoryOrderByInput = "iscsi_target_iqn_ASC"

	// BackupStoreRepositoryOrderByInputIscsiTargetIqnDESC captures enum value "iscsi_target_iqn_DESC"
	BackupStoreRepositoryOrderByInputIscsiTargetIqnDESC BackupStoreRepositoryOrderByInput = "iscsi_target_iqn_DESC"

	// BackupStoreRepositoryOrderByInputNameASC captures enum value "name_ASC"
	BackupStoreRepositoryOrderByInputNameASC BackupStoreRepositoryOrderByInput = "name_ASC"

	// BackupStoreRepositoryOrderByInputNameDESC captures enum value "name_DESC"
	BackupStoreRepositoryOrderByInputNameDESC BackupStoreRepositoryOrderByInput = "name_DESC"

	// BackupStoreRepositoryOrderByInputNfsPathASC captures enum value "nfs_path_ASC"
	BackupStoreRepositoryOrderByInputNfsPathASC BackupStoreRepositoryOrderByInput = "nfs_path_ASC"

	// BackupStoreRepositoryOrderByInputNfsPathDESC captures enum value "nfs_path_DESC"
	BackupStoreRepositoryOrderByInputNfsPathDESC BackupStoreRepositoryOrderByInput = "nfs_path_DESC"

	// BackupStoreRepositoryOrderByInputNfsServerASC captures enum value "nfs_server_ASC"
	BackupStoreRepositoryOrderByInputNfsServerASC BackupStoreRepositoryOrderByInput = "nfs_server_ASC"

	// BackupStoreRepositoryOrderByInputNfsServerDESC captures enum value "nfs_server_DESC"
	BackupStoreRepositoryOrderByInputNfsServerDESC BackupStoreRepositoryOrderByInput = "nfs_server_DESC"

	// BackupStoreRepositoryOrderByInputStatusASC captures enum value "status_ASC"
	BackupStoreRepositoryOrderByInputStatusASC BackupStoreRepositoryOrderByInput = "status_ASC"

	// BackupStoreRepositoryOrderByInputStatusDESC captures enum value "status_DESC"
	BackupStoreRepositoryOrderByInputStatusDESC BackupStoreRepositoryOrderByInput = "status_DESC"

	// BackupStoreRepositoryOrderByInputTotalCapacityASC captures enum value "total_capacity_ASC"
	BackupStoreRepositoryOrderByInputTotalCapacityASC BackupStoreRepositoryOrderByInput = "total_capacity_ASC"

	// BackupStoreRepositoryOrderByInputTotalCapacityDESC captures enum value "total_capacity_DESC"
	BackupStoreRepositoryOrderByInputTotalCapacityDESC BackupStoreRepositoryOrderByInput = "total_capacity_DESC"

	// BackupStoreRepositoryOrderByInputTypeASC captures enum value "type_ASC"
	BackupStoreRepositoryOrderByInputTypeASC BackupStoreRepositoryOrderByInput = "type_ASC"

	// BackupStoreRepositoryOrderByInputTypeDESC captures enum value "type_DESC"
	BackupStoreRepositoryOrderByInputTypeDESC BackupStoreRepositoryOrderByInput = "type_DESC"

	// BackupStoreRepositoryOrderByInputUpdateTimestampASC captures enum value "update_timestamp_ASC"
	BackupStoreRepositoryOrderByInputUpdateTimestampASC BackupStoreRepositoryOrderByInput = "update_timestamp_ASC"

	// BackupStoreRepositoryOrderByInputUpdateTimestampDESC captures enum value "update_timestamp_DESC"
	BackupStoreRepositoryOrderByInputUpdateTimestampDESC BackupStoreRepositoryOrderByInput = "update_timestamp_DESC"

	// BackupStoreRepositoryOrderByInputUsedDataSpaceASC captures enum value "used_data_space_ASC"
	BackupStoreRepositoryOrderByInputUsedDataSpaceASC BackupStoreRepositoryOrderByInput = "used_data_space_ASC"

	// BackupStoreRepositoryOrderByInputUsedDataSpaceDESC captures enum value "used_data_space_DESC"
	BackupStoreRepositoryOrderByInputUsedDataSpaceDESC BackupStoreRepositoryOrderByInput = "used_data_space_DESC"

	// BackupStoreRepositoryOrderByInputUsedDataSpaceUsageASC captures enum value "used_data_space_usage_ASC"
	BackupStoreRepositoryOrderByInputUsedDataSpaceUsageASC BackupStoreRepositoryOrderByInput = "used_data_space_usage_ASC"

	// BackupStoreRepositoryOrderByInputUsedDataSpaceUsageDESC captures enum value "used_data_space_usage_DESC"
	BackupStoreRepositoryOrderByInputUsedDataSpaceUsageDESC BackupStoreRepositoryOrderByInput = "used_data_space_usage_DESC"

	// BackupStoreRepositoryOrderByInputValidDataSpaceASC captures enum value "valid_data_space_ASC"
	BackupStoreRepositoryOrderByInputValidDataSpaceASC BackupStoreRepositoryOrderByInput = "valid_data_space_ASC"

	// BackupStoreRepositoryOrderByInputValidDataSpaceDESC captures enum value "valid_data_space_DESC"
	BackupStoreRepositoryOrderByInputValidDataSpaceDESC BackupStoreRepositoryOrderByInput = "valid_data_space_DESC"
)

// for schema
var backupStoreRepositoryOrderByInputEnum []interface{}

func init() {
	var res []BackupStoreRepositoryOrderByInput
	if err := json.Unmarshal([]byte(`["createdAt_ASC","createdAt_DESC","description_ASC","description_DESC","entityAsyncStatus_ASC","entityAsyncStatus_DESC","error_code_ASC","error_code_DESC","id_ASC","id_DESC","iscsi_chap_name_ASC","iscsi_chap_name_DESC","iscsi_chap_secret_ASC","iscsi_chap_secret_DESC","iscsi_ip_ASC","iscsi_ip_DESC","iscsi_lun_id_ASC","iscsi_lun_id_DESC","iscsi_port_ASC","iscsi_port_DESC","iscsi_target_iqn_ASC","iscsi_target_iqn_DESC","name_ASC","name_DESC","nfs_path_ASC","nfs_path_DESC","nfs_server_ASC","nfs_server_DESC","status_ASC","status_DESC","total_capacity_ASC","total_capacity_DESC","type_ASC","type_DESC","update_timestamp_ASC","update_timestamp_DESC","used_data_space_ASC","used_data_space_DESC","used_data_space_usage_ASC","used_data_space_usage_DESC","valid_data_space_ASC","valid_data_space_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backupStoreRepositoryOrderByInputEnum = append(backupStoreRepositoryOrderByInputEnum, v)
	}
}

func (m BackupStoreRepositoryOrderByInput) validateBackupStoreRepositoryOrderByInputEnum(path, location string, value BackupStoreRepositoryOrderByInput) error {
	if err := validate.EnumCase(path, location, value, backupStoreRepositoryOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this backup store repository order by input
func (m BackupStoreRepositoryOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBackupStoreRepositoryOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this backup store repository order by input based on context it is used
func (m BackupStoreRepositoryOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
