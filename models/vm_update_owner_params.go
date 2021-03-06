// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VMUpdateOwnerParams Vm update owner params
//
// swagger:model VmUpdateOwnerParams
type VMUpdateOwnerParams struct {

	// data
	// Required: true
	Data *VMUpdateOwnerParamsData `json:"data"`

	// where
	// Required: true
	Where *VMWhereInput `json:"where"`
}

// Validate validates this Vm update owner params
func (m *VMUpdateOwnerParams) Validate(formats strfmt.Registry) error {
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

func (m *VMUpdateOwnerParams) validateData(formats strfmt.Registry) error {

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

func (m *VMUpdateOwnerParams) validateWhere(formats strfmt.Registry) error {

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

// ContextValidate validate this Vm update owner params based on the context it is used
func (m *VMUpdateOwnerParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *VMUpdateOwnerParams) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMUpdateOwnerParams) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

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
func (m *VMUpdateOwnerParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMUpdateOwnerParams) UnmarshalBinary(b []byte) error {
	var res VMUpdateOwnerParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VMUpdateOwnerParamsData VM update owner params data
//
// swagger:model VMUpdateOwnerParamsData
type VMUpdateOwnerParamsData struct {

	// search for
	// Required: true
	// Enum: [username id]
	SearchFor string `json:"search_for"`

	// value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this VM update owner params data
func (m *VMUpdateOwnerParamsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSearchFor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var vmUpdateOwnerParamsDataTypeSearchForPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["username","id"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vmUpdateOwnerParamsDataTypeSearchForPropEnum = append(vmUpdateOwnerParamsDataTypeSearchForPropEnum, v)
	}
}

const (

	// VMUpdateOwnerParamsDataSearchForUsername captures enum value "username"
	VMUpdateOwnerParamsDataSearchForUsername string = "username"

	// VMUpdateOwnerParamsDataSearchForID captures enum value "id"
	VMUpdateOwnerParamsDataSearchForID string = "id"
)

// prop value enum
func (m *VMUpdateOwnerParamsData) validateSearchForEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, vmUpdateOwnerParamsDataTypeSearchForPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VMUpdateOwnerParamsData) validateSearchFor(formats strfmt.Registry) error {

	if err := validate.RequiredString("data"+"."+"search_for", "body", m.SearchFor); err != nil {
		return err
	}

	// value enum
	if err := m.validateSearchForEnum("data"+"."+"search_for", "body", m.SearchFor); err != nil {
		return err
	}

	return nil
}

func (m *VMUpdateOwnerParamsData) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this VM update owner params data based on context it is used
func (m *VMUpdateOwnerParamsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VMUpdateOwnerParamsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMUpdateOwnerParamsData) UnmarshalBinary(b []byte) error {
	var res VMUpdateOwnerParamsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
