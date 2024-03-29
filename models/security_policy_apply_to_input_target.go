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

// SecurityPolicyApplyToInputTarget security policy apply to input target
//
// swagger:model SecurityPolicyApplyToInputTarget
type SecurityPolicyApplyToInputTarget struct {

	// label groups
	LabelGroups []*LabelWhereInput `json:"label_groups,omitempty"`

	// security groups
	SecurityGroups *SecurityGroupWhereInput `json:"security_groups,omitempty"`
}

// Validate validates this security policy apply to input target
func (m *SecurityPolicyApplyToInputTarget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabelGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityGroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityPolicyApplyToInputTarget) validateLabelGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.LabelGroups); i++ {
		if swag.IsZero(m.LabelGroups[i]) { // not required
			continue
		}

		if m.LabelGroups[i] != nil {
			if err := m.LabelGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("label_groups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("label_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SecurityPolicyApplyToInputTarget) validateSecurityGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityGroups) { // not required
		return nil
	}

	if m.SecurityGroups != nil {
		if err := m.SecurityGroups.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_groups")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_groups")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this security policy apply to input target based on the context it is used
func (m *SecurityPolicyApplyToInputTarget) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLabelGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurityGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityPolicyApplyToInputTarget) contextValidateLabelGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LabelGroups); i++ {

		if m.LabelGroups[i] != nil {
			if err := m.LabelGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("label_groups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("label_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SecurityPolicyApplyToInputTarget) contextValidateSecurityGroups(ctx context.Context, formats strfmt.Registry) error {

	if m.SecurityGroups != nil {
		if err := m.SecurityGroups.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_groups")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_groups")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityPolicyApplyToInputTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityPolicyApplyToInputTarget) UnmarshalBinary(b []byte) error {
	var res SecurityPolicyApplyToInputTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
