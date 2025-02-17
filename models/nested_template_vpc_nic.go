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

// NestedTemplateVpcNic nested template vpc nic
//
// swagger:model NestedTemplateVpcNic
type NestedTemplateVpcNic struct {

	// use floating ip
	// Required: true
	UseFloatingIP *bool `json:"use_floating_ip"`

	// vpc local id
	// Required: true
	VpcLocalID *string `json:"vpc_local_id"`

	// vpc subnet local id
	// Required: true
	VpcSubnetLocalID *string `json:"vpc_subnet_local_id"`
}

// Validate validates this nested template vpc nic
func (m *NestedTemplateVpcNic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUseFloatingIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpcLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpcSubnetLocalID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedTemplateVpcNic) validateUseFloatingIP(formats strfmt.Registry) error {

	if err := validate.Required("use_floating_ip", "body", m.UseFloatingIP); err != nil {
		return err
	}

	return nil
}

func (m *NestedTemplateVpcNic) validateVpcLocalID(formats strfmt.Registry) error {

	if err := validate.Required("vpc_local_id", "body", m.VpcLocalID); err != nil {
		return err
	}

	return nil
}

func (m *NestedTemplateVpcNic) validateVpcSubnetLocalID(formats strfmt.Registry) error {

	if err := validate.Required("vpc_subnet_local_id", "body", m.VpcSubnetLocalID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this nested template vpc nic based on context it is used
func (m *NestedTemplateVpcNic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NestedTemplateVpcNic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedTemplateVpcNic) UnmarshalBinary(b []byte) error {
	var res NestedTemplateVpcNic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
