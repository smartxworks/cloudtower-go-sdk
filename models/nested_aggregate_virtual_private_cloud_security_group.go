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

// NestedAggregateVirtualPrivateCloudSecurityGroup nested aggregate virtual private cloud security group
//
// swagger:model NestedAggregateVirtualPrivateCloudSecurityGroup
type NestedAggregateVirtualPrivateCloudSecurityGroup struct {

	// count
	// Required: true
	Count *int32 `json:"count"`
}

// Validate validates this nested aggregate virtual private cloud security group
func (m *NestedAggregateVirtualPrivateCloudSecurityGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedAggregateVirtualPrivateCloudSecurityGroup) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("count", "body", m.Count); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this nested aggregate virtual private cloud security group based on context it is used
func (m *NestedAggregateVirtualPrivateCloudSecurityGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NestedAggregateVirtualPrivateCloudSecurityGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedAggregateVirtualPrivateCloudSecurityGroup) UnmarshalBinary(b []byte) error {
	var res NestedAggregateVirtualPrivateCloudSecurityGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
