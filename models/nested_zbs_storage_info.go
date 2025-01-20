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

// NestedZbsStorageInfo nested zbs storage info
//
// swagger:model NestedZbsStorageInfo
type NestedZbsStorageInfo struct {

	// cluster id
	// Required: true
	ClusterID *string `json:"cluster_id"`

	// snapshot info
	SnapshotInfo []*NestedSnapshotInfo `json:"snapshot_info,omitempty"`
}

// Validate validates this nested zbs storage info
func (m *NestedZbsStorageInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedZbsStorageInfo) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("cluster_id", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *NestedZbsStorageInfo) validateSnapshotInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.SnapshotInfo) { // not required
		return nil
	}

	for i := 0; i < len(m.SnapshotInfo); i++ {
		if swag.IsZero(m.SnapshotInfo[i]) { // not required
			continue
		}

		if m.SnapshotInfo[i] != nil {
			if err := m.SnapshotInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("snapshot_info" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("snapshot_info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this nested zbs storage info based on the context it is used
func (m *NestedZbsStorageInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSnapshotInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedZbsStorageInfo) contextValidateSnapshotInfo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SnapshotInfo); i++ {

		if m.SnapshotInfo[i] != nil {
			if err := m.SnapshotInfo[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("snapshot_info" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("snapshot_info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NestedZbsStorageInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedZbsStorageInfo) UnmarshalBinary(b []byte) error {
	var res NestedZbsStorageInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
