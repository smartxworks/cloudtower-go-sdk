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

// UsbDeviceUnmountParams usb device unmount params
//
// swagger:model UsbDeviceUnmountParams
type UsbDeviceUnmountParams struct {

	// data
	Data *UsbDeviceUnmountParamsData `json:"data,omitempty"`

	// where
	// Required: true
	Where *UsbDeviceWhereInput `json:"where"`
}

// Validate validates this usb device unmount params
func (m *UsbDeviceUnmountParams) Validate(formats strfmt.Registry) error {
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

func (m *UsbDeviceUnmountParams) validateData(formats strfmt.Registry) error {
	if swag.IsZero(m.Data) { // not required
		return nil
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

func (m *UsbDeviceUnmountParams) validateWhere(formats strfmt.Registry) error {

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

// ContextValidate validate this usb device unmount params based on the context it is used
func (m *UsbDeviceUnmountParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *UsbDeviceUnmountParams) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

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

func (m *UsbDeviceUnmountParams) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

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
func (m *UsbDeviceUnmountParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UsbDeviceUnmountParams) UnmarshalBinary(b []byte) error {
	var res UsbDeviceUnmountParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UsbDeviceUnmountParamsData usb device unmount params data
//
// swagger:model UsbDeviceUnmountParamsData
type UsbDeviceUnmountParamsData struct {

	// vms
	// Required: true
	Vms *VMWhereInput `json:"vms"`
}

// Validate validates this usb device unmount params data
func (m *UsbDeviceUnmountParamsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVms(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UsbDeviceUnmountParamsData) validateVms(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"vms", "body", m.Vms); err != nil {
		return err
	}

	if m.Vms != nil {
		if err := m.Vms.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "vms")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "vms")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this usb device unmount params data based on the context it is used
func (m *UsbDeviceUnmountParamsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UsbDeviceUnmountParamsData) contextValidateVms(ctx context.Context, formats strfmt.Registry) error {

	if m.Vms != nil {
		if err := m.Vms.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "vms")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "vms")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UsbDeviceUnmountParamsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UsbDeviceUnmountParamsData) UnmarshalBinary(b []byte) error {
	var res UsbDeviceUnmountParamsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
