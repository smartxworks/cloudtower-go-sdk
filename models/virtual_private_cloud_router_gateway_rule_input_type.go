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

// VirtualPrivateCloudRouterGatewayRuleInputType virtual private cloud router gateway rule input type
//
// swagger:model VirtualPrivateCloudRouterGatewayRuleInputType
type VirtualPrivateCloudRouterGatewayRuleInputType struct {

	// dst
	// Required: true
	Dst *string `json:"dst"`

	// nexthop
	Nexthop *string `json:"nexthop,omitempty"`
}

// Validate validates this virtual private cloud router gateway rule input type
func (m *VirtualPrivateCloudRouterGatewayRuleInputType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDst(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualPrivateCloudRouterGatewayRuleInputType) validateDst(formats strfmt.Registry) error {

	if err := validate.Required("dst", "body", m.Dst); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this virtual private cloud router gateway rule input type based on context it is used
func (m *VirtualPrivateCloudRouterGatewayRuleInputType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VirtualPrivateCloudRouterGatewayRuleInputType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualPrivateCloudRouterGatewayRuleInputType) UnmarshalBinary(b []byte) error {
	var res VirtualPrivateCloudRouterGatewayRuleInputType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
