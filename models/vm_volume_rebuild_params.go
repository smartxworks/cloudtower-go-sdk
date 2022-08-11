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

// VMVolumeRebuildParams Vm volume rebuild params
//
// swagger:model VmVolumeRebuildParams
type VMVolumeRebuildParams struct {

	// description
	// Required: true
	Description *string `json:"description"`

	// name
	// Required: true
	Name *string `json:"name"`

	// volume snapshot id
	// Required: true
	VolumeSnapshotID *string `json:"volume_snapshot_id"`
}

// Validate validates this Vm volume rebuild params
func (m *VMVolumeRebuildParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeSnapshotID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMVolumeRebuildParams) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *VMVolumeRebuildParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VMVolumeRebuildParams) validateVolumeSnapshotID(formats strfmt.Registry) error {

	if err := validate.Required("volume_snapshot_id", "body", m.VolumeSnapshotID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Vm volume rebuild params based on context it is used
func (m *VMVolumeRebuildParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VMVolumeRebuildParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMVolumeRebuildParams) UnmarshalBinary(b []byte) error {
	var res VMVolumeRebuildParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
