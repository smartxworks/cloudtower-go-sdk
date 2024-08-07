// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VirtualPrivateCloudSecurityPolicyUpdateParams virtual private cloud security policy update params
//
// swagger:model VirtualPrivateCloudSecurityPolicyUpdateParams
type VirtualPrivateCloudSecurityPolicyUpdateParams struct {

	// apply to
	ApplyTo []*VirtualPrivateCloudSecurityPolicyApplyInput `json:"apply_to,omitempty"`

	// description
	Description *string `json:"description,omitempty"`

	// egress
	Egress []*VirtualPrivateCloudNetworkPolicyRuleInput `json:"egress,omitempty"`

	// ingress
	Ingress []*VirtualPrivateCloudNetworkPolicyRuleInput `json:"ingress,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// policy mode
	PolicyMode *VirtualPrivateCloudSecurityPolicyMode `json:"policy_mode,omitempty"`
}

// Validate validates this virtual private cloud security policy update params
func (m *VirtualPrivateCloudSecurityPolicyUpdateParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplyTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEgress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualPrivateCloudSecurityPolicyUpdateParams) validateApplyTo(formats strfmt.Registry) error {
	if swag.IsZero(m.ApplyTo) { // not required
		return nil
	}

	for i := 0; i < len(m.ApplyTo); i++ {
		if swag.IsZero(m.ApplyTo[i]) { // not required
			continue
		}

		if m.ApplyTo[i] != nil {
			if err := m.ApplyTo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("apply_to" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("apply_to" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudSecurityPolicyUpdateParams) validateEgress(formats strfmt.Registry) error {
	if swag.IsZero(m.Egress) { // not required
		return nil
	}

	for i := 0; i < len(m.Egress); i++ {
		if swag.IsZero(m.Egress[i]) { // not required
			continue
		}

		if m.Egress[i] != nil {
			if err := m.Egress[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("egress" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("egress" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudSecurityPolicyUpdateParams) validateIngress(formats strfmt.Registry) error {
	if swag.IsZero(m.Ingress) { // not required
		return nil
	}

	for i := 0; i < len(m.Ingress); i++ {
		if swag.IsZero(m.Ingress[i]) { // not required
			continue
		}

		if m.Ingress[i] != nil {
			if err := m.Ingress[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ingress" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ingress" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudSecurityPolicyUpdateParams) validatePolicyMode(formats strfmt.Registry) error {
	if swag.IsZero(m.PolicyMode) { // not required
		return nil
	}

	if m.PolicyMode != nil {
		if err := m.PolicyMode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy_mode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("policy_mode")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this virtual private cloud security policy update params based on the context it is used
func (m *VirtualPrivateCloudSecurityPolicyUpdateParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApplyTo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEgress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIngress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePolicyMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualPrivateCloudSecurityPolicyUpdateParams) contextValidateApplyTo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ApplyTo); i++ {

		if m.ApplyTo[i] != nil {
			if err := m.ApplyTo[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("apply_to" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("apply_to" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudSecurityPolicyUpdateParams) contextValidateEgress(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Egress); i++ {

		if m.Egress[i] != nil {
			if err := m.Egress[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("egress" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("egress" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudSecurityPolicyUpdateParams) contextValidateIngress(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Ingress); i++ {

		if m.Ingress[i] != nil {
			if err := m.Ingress[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ingress" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ingress" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudSecurityPolicyUpdateParams) contextValidatePolicyMode(ctx context.Context, formats strfmt.Registry) error {

	if m.PolicyMode != nil {
		if err := m.PolicyMode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy_mode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("policy_mode")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VirtualPrivateCloudSecurityPolicyUpdateParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualPrivateCloudSecurityPolicyUpdateParams) UnmarshalBinary(b []byte) error {
	var res VirtualPrivateCloudSecurityPolicyUpdateParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
