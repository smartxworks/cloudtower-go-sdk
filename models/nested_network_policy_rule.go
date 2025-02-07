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

// NestedNetworkPolicyRule nested network policy rule
//
// swagger:model NestedNetworkPolicyRule
type NestedNetworkPolicyRule struct {

	// ip block
	IPBlock *string `json:"ip_block,omitempty"`

	// ports
	Ports []*NestedNetworkPolicyRulePort `json:"ports,omitempty"`

	// security group
	SecurityGroup *NestedSecurityGroup `json:"security_group,omitempty"`

	// security group id
	SecurityGroupID *string `json:"security_group_id,omitempty"`

	// selector
	Selector []*NestedLabel `json:"selector,omitempty"`

	// selector ids
	SelectorIds []string `json:"selector_ids,omitempty"`

	// type
	// Required: true
	Type *NetworkPolicyRuleType `json:"type"`

	MarshalOpts *NestedNetworkPolicyRuleMarshalOpts `json:"-"`
}

type NestedNetworkPolicyRuleMarshalOpts struct {
	IPBlock_Explicit_Null_When_Empty bool

	Ports_Explicit_Null_When_Empty bool

	SecurityGroup_Explicit_Null_When_Empty bool

	SecurityGroupID_Explicit_Null_When_Empty bool

	Selector_Explicit_Null_When_Empty bool

	SelectorIds_Explicit_Null_When_Empty bool

	Type_Explicit_Null_When_Empty bool
}

func (m NestedNetworkPolicyRule) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field ip_block
	if m.IPBlock != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ip_block\":")
		bytes, err := swag.WriteJSON(m.IPBlock)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IPBlock_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ip_block\":null")
		first = false
	}

	// handle non nullable field ports with omitempty
	if swag.IsZero(m.Ports) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ports\":")
		bytes, err := swag.WriteJSON(m.Ports)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field security_group
	if m.SecurityGroup != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"security_group\":")
		bytes, err := swag.WriteJSON(m.SecurityGroup)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.SecurityGroup_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"security_group\":null")
		first = false
	}

	// handle nullable field security_group_id
	if m.SecurityGroupID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"security_group_id\":")
		bytes, err := swag.WriteJSON(m.SecurityGroupID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.SecurityGroupID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"security_group_id\":null")
		first = false
	}

	// handle non nullable field selector with omitempty
	if swag.IsZero(m.Selector) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"selector\":")
		bytes, err := swag.WriteJSON(m.Selector)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle non nullable field selector_ids with omitempty
	if swag.IsZero(m.SelectorIds) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"selector_ids\":")
		bytes, err := swag.WriteJSON(m.SelectorIds)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field type
	if m.Type != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"type\":")
		bytes, err := swag.WriteJSON(m.Type)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Type_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"type\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this nested network policy rule
func (m *NestedNetworkPolicyRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedNetworkPolicyRule) validatePorts(formats strfmt.Registry) error {
	if swag.IsZero(m.Ports) { // not required
		return nil
	}

	for i := 0; i < len(m.Ports); i++ {
		if swag.IsZero(m.Ports[i]) { // not required
			continue
		}

		if m.Ports[i] != nil {
			if err := m.Ports[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ports" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NestedNetworkPolicyRule) validateSecurityGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityGroup) { // not required
		return nil
	}

	if m.SecurityGroup != nil {
		if err := m.SecurityGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_group")
			}
			return err
		}
	}

	return nil
}

func (m *NestedNetworkPolicyRule) validateSelector(formats strfmt.Registry) error {
	if swag.IsZero(m.Selector) { // not required
		return nil
	}

	for i := 0; i < len(m.Selector); i++ {
		if swag.IsZero(m.Selector[i]) { // not required
			continue
		}

		if m.Selector[i] != nil {
			if err := m.Selector[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("selector" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("selector" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NestedNetworkPolicyRule) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nested network policy rule based on the context it is used
func (m *NestedNetworkPolicyRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurityGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedNetworkPolicyRule) contextValidatePorts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Ports); i++ {

		if m.Ports[i] != nil {
			if err := m.Ports[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ports" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NestedNetworkPolicyRule) contextValidateSecurityGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.SecurityGroup != nil {
		if err := m.SecurityGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_group")
			}
			return err
		}
	}

	return nil
}

func (m *NestedNetworkPolicyRule) contextValidateSelector(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Selector); i++ {

		if m.Selector[i] != nil {
			if err := m.Selector[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("selector" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("selector" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NestedNetworkPolicyRule) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NestedNetworkPolicyRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedNetworkPolicyRule) UnmarshalBinary(b []byte) error {
	var res NestedNetworkPolicyRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
