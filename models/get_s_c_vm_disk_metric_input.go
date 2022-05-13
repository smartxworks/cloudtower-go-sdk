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

// GetSCVMDiskMetricInput get s c VM disk metric input
//
// swagger:model GetSCVMDiskMetricInput
type GetSCVMDiskMetricInput struct {

	// disks
	// Required: true
	Disks *DiskWhereInput `json:"disks"`

	// metrics
	// Required: true
	Metrics []string `json:"metrics"`

	// range
	// Required: true
	Range *string `json:"range"`
}

// Validate validates this get s c VM disk metric input
func (m *GetSCVMDiskMetricInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetrics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRange(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetSCVMDiskMetricInput) validateDisks(formats strfmt.Registry) error {

	if err := validate.Required("disks", "body", m.Disks); err != nil {
		return err
	}

	if m.Disks != nil {
		if err := m.Disks.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disks")
			}
			return err
		}
	}

	return nil
}

func (m *GetSCVMDiskMetricInput) validateMetrics(formats strfmt.Registry) error {

	if err := validate.Required("metrics", "body", m.Metrics); err != nil {
		return err
	}

	return nil
}

func (m *GetSCVMDiskMetricInput) validateRange(formats strfmt.Registry) error {

	if err := validate.Required("range", "body", m.Range); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get s c VM disk metric input based on the context it is used
func (m *GetSCVMDiskMetricInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetSCVMDiskMetricInput) contextValidateDisks(ctx context.Context, formats strfmt.Registry) error {

	if m.Disks != nil {
		if err := m.Disks.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("disks")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetSCVMDiskMetricInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetSCVMDiskMetricInput) UnmarshalBinary(b []byte) error {
	var res GetSCVMDiskMetricInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
