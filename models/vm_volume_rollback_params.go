// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VMVolumeRollbackParams Vm volume rollback params
//
// swagger:model VmVolumeRollbackParams
type VMVolumeRollbackParams struct {

	// volume snapshot id
	// Required: true
	VolumeSnapshotID *string `json:"volume_snapshot_id"`

	MarshalOpts *VMVolumeRollbackParamsMarshalOpts `json:"-"`
}

type VMVolumeRollbackParamsMarshalOpts struct {
	VolumeSnapshotID_Explicit_Null_When_Empty bool
}

func (m VMVolumeRollbackParams) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field volume_snapshot_id
	if m.VolumeSnapshotID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"volume_snapshot_id\":")
		bytes, err := swag.WriteJSON(m.VolumeSnapshotID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VolumeSnapshotID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"volume_snapshot_id\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this Vm volume rollback params
func (m *VMVolumeRollbackParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVolumeSnapshotID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMVolumeRollbackParams) validateVolumeSnapshotID(formats strfmt.Registry) error {

	if err := validate.Required("volume_snapshot_id", "body", m.VolumeSnapshotID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Vm volume rollback params based on context it is used
func (m *VMVolumeRollbackParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VMVolumeRollbackParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMVolumeRollbackParams) UnmarshalBinary(b []byte) error {
	var res VMVolumeRollbackParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
