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

// IscsiLunOrderByInput iscsi lun order by input
//
// swagger:model IscsiLunOrderByInput
type IscsiLunOrderByInput string

func NewIscsiLunOrderByInput(value IscsiLunOrderByInput) *IscsiLunOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated IscsiLunOrderByInput.
func (m IscsiLunOrderByInput) Pointer() *IscsiLunOrderByInput {
	return &m
}

const (

	// IscsiLunOrderByInputAllowedInitiatorsASC captures enum value "allowed_initiators_ASC"
	IscsiLunOrderByInputAllowedInitiatorsASC IscsiLunOrderByInput = "allowed_initiators_ASC"

	// IscsiLunOrderByInputAllowedInitiatorsDESC captures enum value "allowed_initiators_DESC"
	IscsiLunOrderByInputAllowedInitiatorsDESC IscsiLunOrderByInput = "allowed_initiators_DESC"

	// IscsiLunOrderByInputAssignedSizeASC captures enum value "assigned_size_ASC"
	IscsiLunOrderByInputAssignedSizeASC IscsiLunOrderByInput = "assigned_size_ASC"

	// IscsiLunOrderByInputAssignedSizeDESC captures enum value "assigned_size_DESC"
	IscsiLunOrderByInputAssignedSizeDESC IscsiLunOrderByInput = "assigned_size_DESC"

	// IscsiLunOrderByInputBpsASC captures enum value "bps_ASC"
	IscsiLunOrderByInputBpsASC IscsiLunOrderByInput = "bps_ASC"

	// IscsiLunOrderByInputBpsDESC captures enum value "bps_DESC"
	IscsiLunOrderByInputBpsDESC IscsiLunOrderByInput = "bps_DESC"

	// IscsiLunOrderByInputBpsMaxASC captures enum value "bps_max_ASC"
	IscsiLunOrderByInputBpsMaxASC IscsiLunOrderByInput = "bps_max_ASC"

	// IscsiLunOrderByInputBpsMaxDESC captures enum value "bps_max_DESC"
	IscsiLunOrderByInputBpsMaxDESC IscsiLunOrderByInput = "bps_max_DESC"

	// IscsiLunOrderByInputBpsMaxLengthASC captures enum value "bps_max_length_ASC"
	IscsiLunOrderByInputBpsMaxLengthASC IscsiLunOrderByInput = "bps_max_length_ASC"

	// IscsiLunOrderByInputBpsMaxLengthDESC captures enum value "bps_max_length_DESC"
	IscsiLunOrderByInputBpsMaxLengthDESC IscsiLunOrderByInput = "bps_max_length_DESC"

	// IscsiLunOrderByInputBpsRdASC captures enum value "bps_rd_ASC"
	IscsiLunOrderByInputBpsRdASC IscsiLunOrderByInput = "bps_rd_ASC"

	// IscsiLunOrderByInputBpsRdDESC captures enum value "bps_rd_DESC"
	IscsiLunOrderByInputBpsRdDESC IscsiLunOrderByInput = "bps_rd_DESC"

	// IscsiLunOrderByInputBpsRdMaxASC captures enum value "bps_rd_max_ASC"
	IscsiLunOrderByInputBpsRdMaxASC IscsiLunOrderByInput = "bps_rd_max_ASC"

	// IscsiLunOrderByInputBpsRdMaxDESC captures enum value "bps_rd_max_DESC"
	IscsiLunOrderByInputBpsRdMaxDESC IscsiLunOrderByInput = "bps_rd_max_DESC"

	// IscsiLunOrderByInputBpsRdMaxLengthASC captures enum value "bps_rd_max_length_ASC"
	IscsiLunOrderByInputBpsRdMaxLengthASC IscsiLunOrderByInput = "bps_rd_max_length_ASC"

	// IscsiLunOrderByInputBpsRdMaxLengthDESC captures enum value "bps_rd_max_length_DESC"
	IscsiLunOrderByInputBpsRdMaxLengthDESC IscsiLunOrderByInput = "bps_rd_max_length_DESC"

	// IscsiLunOrderByInputBpsWrASC captures enum value "bps_wr_ASC"
	IscsiLunOrderByInputBpsWrASC IscsiLunOrderByInput = "bps_wr_ASC"

	// IscsiLunOrderByInputBpsWrDESC captures enum value "bps_wr_DESC"
	IscsiLunOrderByInputBpsWrDESC IscsiLunOrderByInput = "bps_wr_DESC"

	// IscsiLunOrderByInputBpsWrMaxASC captures enum value "bps_wr_max_ASC"
	IscsiLunOrderByInputBpsWrMaxASC IscsiLunOrderByInput = "bps_wr_max_ASC"

	// IscsiLunOrderByInputBpsWrMaxDESC captures enum value "bps_wr_max_DESC"
	IscsiLunOrderByInputBpsWrMaxDESC IscsiLunOrderByInput = "bps_wr_max_DESC"

	// IscsiLunOrderByInputBpsWrMaxLengthASC captures enum value "bps_wr_max_length_ASC"
	IscsiLunOrderByInputBpsWrMaxLengthASC IscsiLunOrderByInput = "bps_wr_max_length_ASC"

	// IscsiLunOrderByInputBpsWrMaxLengthDESC captures enum value "bps_wr_max_length_DESC"
	IscsiLunOrderByInputBpsWrMaxLengthDESC IscsiLunOrderByInput = "bps_wr_max_length_DESC"

	// IscsiLunOrderByInputEntityAsyncStatusASC captures enum value "entityAsyncStatus_ASC"
	IscsiLunOrderByInputEntityAsyncStatusASC IscsiLunOrderByInput = "entityAsyncStatus_ASC"

	// IscsiLunOrderByInputEntityAsyncStatusDESC captures enum value "entityAsyncStatus_DESC"
	IscsiLunOrderByInputEntityAsyncStatusDESC IscsiLunOrderByInput = "entityAsyncStatus_DESC"

	// IscsiLunOrderByInputIDASC captures enum value "id_ASC"
	IscsiLunOrderByInputIDASC IscsiLunOrderByInput = "id_ASC"

	// IscsiLunOrderByInputIDDESC captures enum value "id_DESC"
	IscsiLunOrderByInputIDDESC IscsiLunOrderByInput = "id_DESC"

	// IscsiLunOrderByInputIoSizeASC captures enum value "io_size_ASC"
	IscsiLunOrderByInputIoSizeASC IscsiLunOrderByInput = "io_size_ASC"

	// IscsiLunOrderByInputIoSizeDESC captures enum value "io_size_DESC"
	IscsiLunOrderByInputIoSizeDESC IscsiLunOrderByInput = "io_size_DESC"

	// IscsiLunOrderByInputIopsASC captures enum value "iops_ASC"
	IscsiLunOrderByInputIopsASC IscsiLunOrderByInput = "iops_ASC"

	// IscsiLunOrderByInputIopsDESC captures enum value "iops_DESC"
	IscsiLunOrderByInputIopsDESC IscsiLunOrderByInput = "iops_DESC"

	// IscsiLunOrderByInputIopsMaxASC captures enum value "iops_max_ASC"
	IscsiLunOrderByInputIopsMaxASC IscsiLunOrderByInput = "iops_max_ASC"

	// IscsiLunOrderByInputIopsMaxDESC captures enum value "iops_max_DESC"
	IscsiLunOrderByInputIopsMaxDESC IscsiLunOrderByInput = "iops_max_DESC"

	// IscsiLunOrderByInputIopsMaxLengthASC captures enum value "iops_max_length_ASC"
	IscsiLunOrderByInputIopsMaxLengthASC IscsiLunOrderByInput = "iops_max_length_ASC"

	// IscsiLunOrderByInputIopsMaxLengthDESC captures enum value "iops_max_length_DESC"
	IscsiLunOrderByInputIopsMaxLengthDESC IscsiLunOrderByInput = "iops_max_length_DESC"

	// IscsiLunOrderByInputIopsRdASC captures enum value "iops_rd_ASC"
	IscsiLunOrderByInputIopsRdASC IscsiLunOrderByInput = "iops_rd_ASC"

	// IscsiLunOrderByInputIopsRdDESC captures enum value "iops_rd_DESC"
	IscsiLunOrderByInputIopsRdDESC IscsiLunOrderByInput = "iops_rd_DESC"

	// IscsiLunOrderByInputIopsRdMaxASC captures enum value "iops_rd_max_ASC"
	IscsiLunOrderByInputIopsRdMaxASC IscsiLunOrderByInput = "iops_rd_max_ASC"

	// IscsiLunOrderByInputIopsRdMaxDESC captures enum value "iops_rd_max_DESC"
	IscsiLunOrderByInputIopsRdMaxDESC IscsiLunOrderByInput = "iops_rd_max_DESC"

	// IscsiLunOrderByInputIopsRdMaxLengthASC captures enum value "iops_rd_max_length_ASC"
	IscsiLunOrderByInputIopsRdMaxLengthASC IscsiLunOrderByInput = "iops_rd_max_length_ASC"

	// IscsiLunOrderByInputIopsRdMaxLengthDESC captures enum value "iops_rd_max_length_DESC"
	IscsiLunOrderByInputIopsRdMaxLengthDESC IscsiLunOrderByInput = "iops_rd_max_length_DESC"

	// IscsiLunOrderByInputIopsWrASC captures enum value "iops_wr_ASC"
	IscsiLunOrderByInputIopsWrASC IscsiLunOrderByInput = "iops_wr_ASC"

	// IscsiLunOrderByInputIopsWrDESC captures enum value "iops_wr_DESC"
	IscsiLunOrderByInputIopsWrDESC IscsiLunOrderByInput = "iops_wr_DESC"

	// IscsiLunOrderByInputIopsWrMaxASC captures enum value "iops_wr_max_ASC"
	IscsiLunOrderByInputIopsWrMaxASC IscsiLunOrderByInput = "iops_wr_max_ASC"

	// IscsiLunOrderByInputIopsWrMaxDESC captures enum value "iops_wr_max_DESC"
	IscsiLunOrderByInputIopsWrMaxDESC IscsiLunOrderByInput = "iops_wr_max_DESC"

	// IscsiLunOrderByInputIopsWrMaxLengthASC captures enum value "iops_wr_max_length_ASC"
	IscsiLunOrderByInputIopsWrMaxLengthASC IscsiLunOrderByInput = "iops_wr_max_length_ASC"

	// IscsiLunOrderByInputIopsWrMaxLengthDESC captures enum value "iops_wr_max_length_DESC"
	IscsiLunOrderByInputIopsWrMaxLengthDESC IscsiLunOrderByInput = "iops_wr_max_length_DESC"

	// IscsiLunOrderByInputLocalCreatedAtASC captures enum value "local_created_at_ASC"
	IscsiLunOrderByInputLocalCreatedAtASC IscsiLunOrderByInput = "local_created_at_ASC"

	// IscsiLunOrderByInputLocalCreatedAtDESC captures enum value "local_created_at_DESC"
	IscsiLunOrderByInputLocalCreatedAtDESC IscsiLunOrderByInput = "local_created_at_DESC"

	// IscsiLunOrderByInputLocalIDASC captures enum value "local_id_ASC"
	IscsiLunOrderByInputLocalIDASC IscsiLunOrderByInput = "local_id_ASC"

	// IscsiLunOrderByInputLocalIDDESC captures enum value "local_id_DESC"
	IscsiLunOrderByInputLocalIDDESC IscsiLunOrderByInput = "local_id_DESC"

	// IscsiLunOrderByInputLunIDASC captures enum value "lun_id_ASC"
	IscsiLunOrderByInputLunIDASC IscsiLunOrderByInput = "lun_id_ASC"

	// IscsiLunOrderByInputLunIDDESC captures enum value "lun_id_DESC"
	IscsiLunOrderByInputLunIDDESC IscsiLunOrderByInput = "lun_id_DESC"

	// IscsiLunOrderByInputNameASC captures enum value "name_ASC"
	IscsiLunOrderByInputNameASC IscsiLunOrderByInput = "name_ASC"

	// IscsiLunOrderByInputNameDESC captures enum value "name_DESC"
	IscsiLunOrderByInputNameDESC IscsiLunOrderByInput = "name_DESC"

	// IscsiLunOrderByInputReplicaNumASC captures enum value "replica_num_ASC"
	IscsiLunOrderByInputReplicaNumASC IscsiLunOrderByInput = "replica_num_ASC"

	// IscsiLunOrderByInputReplicaNumDESC captures enum value "replica_num_DESC"
	IscsiLunOrderByInputReplicaNumDESC IscsiLunOrderByInput = "replica_num_DESC"

	// IscsiLunOrderByInputSharedSizeASC captures enum value "shared_size_ASC"
	IscsiLunOrderByInputSharedSizeASC IscsiLunOrderByInput = "shared_size_ASC"

	// IscsiLunOrderByInputSharedSizeDESC captures enum value "shared_size_DESC"
	IscsiLunOrderByInputSharedSizeDESC IscsiLunOrderByInput = "shared_size_DESC"

	// IscsiLunOrderByInputSnapshotNumASC captures enum value "snapshot_num_ASC"
	IscsiLunOrderByInputSnapshotNumASC IscsiLunOrderByInput = "snapshot_num_ASC"

	// IscsiLunOrderByInputSnapshotNumDESC captures enum value "snapshot_num_DESC"
	IscsiLunOrderByInputSnapshotNumDESC IscsiLunOrderByInput = "snapshot_num_DESC"

	// IscsiLunOrderByInputStripeNumASC captures enum value "stripe_num_ASC"
	IscsiLunOrderByInputStripeNumASC IscsiLunOrderByInput = "stripe_num_ASC"

	// IscsiLunOrderByInputStripeNumDESC captures enum value "stripe_num_DESC"
	IscsiLunOrderByInputStripeNumDESC IscsiLunOrderByInput = "stripe_num_DESC"

	// IscsiLunOrderByInputStripeSizeASC captures enum value "stripe_size_ASC"
	IscsiLunOrderByInputStripeSizeASC IscsiLunOrderByInput = "stripe_size_ASC"

	// IscsiLunOrderByInputStripeSizeDESC captures enum value "stripe_size_DESC"
	IscsiLunOrderByInputStripeSizeDESC IscsiLunOrderByInput = "stripe_size_DESC"

	// IscsiLunOrderByInputThinProvisionASC captures enum value "thin_provision_ASC"
	IscsiLunOrderByInputThinProvisionASC IscsiLunOrderByInput = "thin_provision_ASC"

	// IscsiLunOrderByInputThinProvisionDESC captures enum value "thin_provision_DESC"
	IscsiLunOrderByInputThinProvisionDESC IscsiLunOrderByInput = "thin_provision_DESC"

	// IscsiLunOrderByInputUniqueLogicalSizeASC captures enum value "unique_logical_size_ASC"
	IscsiLunOrderByInputUniqueLogicalSizeASC IscsiLunOrderByInput = "unique_logical_size_ASC"

	// IscsiLunOrderByInputUniqueLogicalSizeDESC captures enum value "unique_logical_size_DESC"
	IscsiLunOrderByInputUniqueLogicalSizeDESC IscsiLunOrderByInput = "unique_logical_size_DESC"

	// IscsiLunOrderByInputUniqueSizeASC captures enum value "unique_size_ASC"
	IscsiLunOrderByInputUniqueSizeASC IscsiLunOrderByInput = "unique_size_ASC"

	// IscsiLunOrderByInputUniqueSizeDESC captures enum value "unique_size_DESC"
	IscsiLunOrderByInputUniqueSizeDESC IscsiLunOrderByInput = "unique_size_DESC"

	// IscsiLunOrderByInputZbsVolumeIDASC captures enum value "zbs_volume_id_ASC"
	IscsiLunOrderByInputZbsVolumeIDASC IscsiLunOrderByInput = "zbs_volume_id_ASC"

	// IscsiLunOrderByInputZbsVolumeIDDESC captures enum value "zbs_volume_id_DESC"
	IscsiLunOrderByInputZbsVolumeIDDESC IscsiLunOrderByInput = "zbs_volume_id_DESC"
)

// for schema
var iscsiLunOrderByInputEnum []interface{}

func init() {
	var res []IscsiLunOrderByInput
	if err := json.Unmarshal([]byte(`["allowed_initiators_ASC","allowed_initiators_DESC","assigned_size_ASC","assigned_size_DESC","bps_ASC","bps_DESC","bps_max_ASC","bps_max_DESC","bps_max_length_ASC","bps_max_length_DESC","bps_rd_ASC","bps_rd_DESC","bps_rd_max_ASC","bps_rd_max_DESC","bps_rd_max_length_ASC","bps_rd_max_length_DESC","bps_wr_ASC","bps_wr_DESC","bps_wr_max_ASC","bps_wr_max_DESC","bps_wr_max_length_ASC","bps_wr_max_length_DESC","entityAsyncStatus_ASC","entityAsyncStatus_DESC","id_ASC","id_DESC","io_size_ASC","io_size_DESC","iops_ASC","iops_DESC","iops_max_ASC","iops_max_DESC","iops_max_length_ASC","iops_max_length_DESC","iops_rd_ASC","iops_rd_DESC","iops_rd_max_ASC","iops_rd_max_DESC","iops_rd_max_length_ASC","iops_rd_max_length_DESC","iops_wr_ASC","iops_wr_DESC","iops_wr_max_ASC","iops_wr_max_DESC","iops_wr_max_length_ASC","iops_wr_max_length_DESC","local_created_at_ASC","local_created_at_DESC","local_id_ASC","local_id_DESC","lun_id_ASC","lun_id_DESC","name_ASC","name_DESC","replica_num_ASC","replica_num_DESC","shared_size_ASC","shared_size_DESC","snapshot_num_ASC","snapshot_num_DESC","stripe_num_ASC","stripe_num_DESC","stripe_size_ASC","stripe_size_DESC","thin_provision_ASC","thin_provision_DESC","unique_logical_size_ASC","unique_logical_size_DESC","unique_size_ASC","unique_size_DESC","zbs_volume_id_ASC","zbs_volume_id_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iscsiLunOrderByInputEnum = append(iscsiLunOrderByInputEnum, v)
	}
}

func (m IscsiLunOrderByInput) validateIscsiLunOrderByInputEnum(path, location string, value IscsiLunOrderByInput) error {
	if err := validate.EnumCase(path, location, value, iscsiLunOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this iscsi lun order by input
func (m IscsiLunOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateIscsiLunOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this iscsi lun order by input based on context it is used
func (m IscsiLunOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
