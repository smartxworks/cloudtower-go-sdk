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

// NicUpdationParams nic updation params
//
// swagger:model NicUpdationParams
type NicUpdationParams struct {

	// data
	// Required: true
	Data *NicUpdationParamsData `json:"data"`

	// where
	// Required: true
	Where *NicWhereInput `json:"where"`
}

// Validate validates this nic updation params
func (m *NicUpdationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWhere(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NicUpdationParams) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *NicUpdationParams) validateWhere(formats strfmt.Registry) error {

	if err := validate.Required("where", "body", m.Where); err != nil {
		return err
	}

	if m.Where != nil {
		if err := m.Where.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("where")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("where")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nic updation params based on the context it is used
func (m *NicUpdationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWhere(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NicUpdationParams) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if m.Data != nil {
		if err := m.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *NicUpdationParams) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

	if m.Where != nil {
		if err := m.Where.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("where")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("where")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NicUpdationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NicUpdationParams) UnmarshalBinary(b []byte) error {
	var res NicUpdationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NicUpdationParamsData nic updation params data
//
// swagger:model NicUpdationParamsData
type NicUpdationParamsData struct {

	// mtu
	Mtu *int32 `json:"mtu,omitempty"`

	// nic user usage
	NicUserUsage *NicUserUsage `json:"nic_user_usage,omitempty"`

	// total vf num
	TotalVfNum *int32 `json:"total_vf_num,omitempty"`
}

// Validate validates this nic updation params data
func (m *NicUpdationParamsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNicUserUsage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NicUpdationParamsData) validateNicUserUsage(formats strfmt.Registry) error {
	if swag.IsZero(m.NicUserUsage) { // not required
		return nil
	}

	if m.NicUserUsage != nil {
		if err := m.NicUserUsage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "nic_user_usage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "nic_user_usage")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nic updation params data based on the context it is used
func (m *NicUpdationParamsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNicUserUsage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NicUpdationParamsData) contextValidateNicUserUsage(ctx context.Context, formats strfmt.Registry) error {

	if m.NicUserUsage != nil {
		if err := m.NicUserUsage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "nic_user_usage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "nic_user_usage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NicUpdationParamsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NicUpdationParamsData) UnmarshalBinary(b []byte) error {
	var res NicUpdationParamsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
