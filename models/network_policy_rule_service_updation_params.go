// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkPolicyRuleServiceUpdationParams network policy rule service updation params
//
// swagger:model NetworkPolicyRuleServiceUpdationParams
type NetworkPolicyRuleServiceUpdationParams struct {

	// data
	// Required: true
	Data *NetworkPolicyRuleServiceUpdationParamsData `json:"data"`

	// where
	// Required: true
	Where *NetworkPolicyRuleServiceWhereInput `json:"where"`

	MarshalOpts *NetworkPolicyRuleServiceUpdationParamsMarshalOpts `json:"-"`
}

type NetworkPolicyRuleServiceUpdationParamsMarshalOpts struct {
	Data_Explicit_Null_When_Empty bool

	Where_Explicit_Null_When_Empty bool
}

func (m NetworkPolicyRuleServiceUpdationParams) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field data
	if m.Data != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"data\":")
		bytes, err := swag.WriteJSON(m.Data)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Data_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"data\":null")
		first = false
	}

	// handle nullable field where
	if m.Where != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"where\":")
		bytes, err := swag.WriteJSON(m.Where)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Where_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"where\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this network policy rule service updation params
func (m *NetworkPolicyRuleServiceUpdationParams) Validate(formats strfmt.Registry) error {
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

func (m *NetworkPolicyRuleServiceUpdationParams) validateData(formats strfmt.Registry) error {

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

func (m *NetworkPolicyRuleServiceUpdationParams) validateWhere(formats strfmt.Registry) error {

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

// ContextValidate validate this network policy rule service updation params based on the context it is used
func (m *NetworkPolicyRuleServiceUpdationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *NetworkPolicyRuleServiceUpdationParams) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

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

func (m *NetworkPolicyRuleServiceUpdationParams) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

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
func (m *NetworkPolicyRuleServiceUpdationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkPolicyRuleServiceUpdationParams) UnmarshalBinary(b []byte) error {
	var res NetworkPolicyRuleServiceUpdationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NetworkPolicyRuleServiceUpdationParamsData network policy rule service updation params data
//
// swagger:model NetworkPolicyRuleServiceUpdationParamsData
type NetworkPolicyRuleServiceUpdationParamsData struct {

	// description
	Description *string `json:"description,omitempty"`

	// members
	Members []*NetworkPolicyRuleServiceMemberParams `json:"members,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	MarshalOpts *NetworkPolicyRuleServiceUpdationParamsDataMarshalOpts `json:"-"`
}

type NetworkPolicyRuleServiceUpdationParamsDataMarshalOpts struct {
	Description_Explicit_Null_When_Empty bool

	Members_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool
}

func (m NetworkPolicyRuleServiceUpdationParamsData) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field description
	if m.Description != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description\":")
		bytes, err := swag.WriteJSON(m.Description)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Description_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description\":null")
		first = false
	}

	// handle non nullable field members with omitempty
	if !swag.IsZero(m.Members) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"members\":")
		bytes, err := swag.WriteJSON(m.Members)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field name
	if m.Name != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name\":")
		bytes, err := swag.WriteJSON(m.Name)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Name_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this network policy rule service updation params data
func (m *NetworkPolicyRuleServiceUpdationParamsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMembers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkPolicyRuleServiceUpdationParamsData) validateMembers(formats strfmt.Registry) error {
	if swag.IsZero(m.Members) { // not required
		return nil
	}

	for i := 0; i < len(m.Members); i++ {
		if swag.IsZero(m.Members[i]) { // not required
			continue
		}

		if m.Members[i] != nil {
			if err := m.Members[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + "members" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + "members" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this network policy rule service updation params data based on the context it is used
func (m *NetworkPolicyRuleServiceUpdationParamsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMembers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkPolicyRuleServiceUpdationParamsData) contextValidateMembers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Members); i++ {

		if m.Members[i] != nil {
			if err := m.Members[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + "members" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + "members" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkPolicyRuleServiceUpdationParamsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkPolicyRuleServiceUpdationParamsData) UnmarshalBinary(b []byte) error {
	var res NetworkPolicyRuleServiceUpdationParamsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
