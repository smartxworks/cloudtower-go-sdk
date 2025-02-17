// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NestedSnapshotInfo nested snapshot info
//
// swagger:model NestedSnapshotInfo
type NestedSnapshotInfo struct {

	// index
	// Required: true
	Index *int32 `json:"index"`

	// iscsi lun snapshot uuid
	// Required: true
	IscsiLunSnapshotUUID *string `json:"iscsi_lun_snapshot_uuid"`

	// storage encrypted
	StorageEncrypted *bool `json:"storage_encrypted,omitempty"`

	// storage policy config
	StoragePolicyConfig *NestedStoragePolicyConfig `json:"storage_policy_config,omitempty"`
}

// Validate validates this nested snapshot info
func (m *NestedSnapshotInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIscsiLunSnapshotUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoragePolicyConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedSnapshotInfo) validateIndex(formats strfmt.Registry) error {

	if err := validate.Required("index", "body", m.Index); err != nil {
		return err
	}

	return nil
}

func (m *NestedSnapshotInfo) validateIscsiLunSnapshotUUID(formats strfmt.Registry) error {

	if err := validate.Required("iscsi_lun_snapshot_uuid", "body", m.IscsiLunSnapshotUUID); err != nil {
		return err
	}

	return nil
}

func (m *NestedSnapshotInfo) validateStoragePolicyConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.StoragePolicyConfig) { // not required
		return nil
	}

	if m.StoragePolicyConfig != nil {
		if err := m.StoragePolicyConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storage_policy_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storage_policy_config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nested snapshot info based on the context it is used
func (m *NestedSnapshotInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStoragePolicyConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedSnapshotInfo) contextValidateStoragePolicyConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.StoragePolicyConfig != nil {
		if err := m.StoragePolicyConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storage_policy_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storage_policy_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NestedSnapshotInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedSnapshotInfo) UnmarshalBinary(b []byte) error {
	var res NestedSnapshotInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
