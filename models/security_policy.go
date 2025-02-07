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

// SecurityPolicy security policy
//
// swagger:model SecurityPolicy
type SecurityPolicy struct {

	// apply to
	// Required: true
	ApplyTo []*NestedSecurityPolicyApply `json:"apply_to"`

	// description
	// Required: true
	Description *string `json:"description"`

	// egress
	Egress []*NestedNetworkPolicyRule `json:"egress,omitempty"`

	// everoute cluster
	// Required: true
	EverouteCluster *NestedEverouteCluster `json:"everoute_cluster"`

	// id
	// Required: true
	ID *string `json:"id"`

	// ingress
	Ingress []*NestedNetworkPolicyRule `json:"ingress,omitempty"`

	// is blocklist
	// Required: true
	IsBlocklist *bool `json:"is_blocklist"`

	// name
	// Required: true
	Name *string `json:"name"`

	// policy mode
	PolicyMode *PolicyMode `json:"policy_mode,omitempty"`

	// statistics
	Statistics *NestedSecurityPolicyStatistics `json:"statistics,omitempty"`

	MarshalOpts *SecurityPolicyMarshalOpts `json:"-"`
}

type SecurityPolicyMarshalOpts struct {
	ApplyTo_Explicit_Null_When_Empty bool

	Description_Explicit_Null_When_Empty bool

	Egress_Explicit_Null_When_Empty bool

	EverouteCluster_Explicit_Null_When_Empty bool

	ID_Explicit_Null_When_Empty bool

	Ingress_Explicit_Null_When_Empty bool

	IsBlocklist_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	PolicyMode_Explicit_Null_When_Empty bool

	Statistics_Explicit_Null_When_Empty bool
}

func (m SecurityPolicy) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle non nullable field apply_to without omitempty
	{
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"apply_to\":")
		bytes, err := swag.WriteJSON(m.ApplyTo)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

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

	// handle non nullable field egress with omitempty
	if swag.IsZero(m.Egress) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"egress\":")
		bytes, err := swag.WriteJSON(m.Egress)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field everoute_cluster
	if m.EverouteCluster != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"everoute_cluster\":")
		bytes, err := swag.WriteJSON(m.EverouteCluster)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.EverouteCluster_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"everoute_cluster\":null")
		first = false
	}

	// handle nullable field id
	if m.ID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id\":")
		bytes, err := swag.WriteJSON(m.ID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id\":null")
		first = false
	}

	// handle non nullable field ingress with omitempty
	if swag.IsZero(m.Ingress) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ingress\":")
		bytes, err := swag.WriteJSON(m.Ingress)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field is_blocklist
	if m.IsBlocklist != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"is_blocklist\":")
		bytes, err := swag.WriteJSON(m.IsBlocklist)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IsBlocklist_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"is_blocklist\":null")
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

	// handle nullable field policy_mode
	if m.PolicyMode != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"policy_mode\":")
		bytes, err := swag.WriteJSON(m.PolicyMode)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.PolicyMode_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"policy_mode\":null")
		first = false
	}

	// handle nullable field statistics
	if m.Statistics != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"statistics\":")
		bytes, err := swag.WriteJSON(m.Statistics)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Statistics_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"statistics\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this security policy
func (m *SecurityPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplyTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEgress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEverouteCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsBlocklist(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatistics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityPolicy) validateApplyTo(formats strfmt.Registry) error {

	if err := validate.Required("apply_to", "body", m.ApplyTo); err != nil {
		return err
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

func (m *SecurityPolicy) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *SecurityPolicy) validateEgress(formats strfmt.Registry) error {
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

func (m *SecurityPolicy) validateEverouteCluster(formats strfmt.Registry) error {

	if err := validate.Required("everoute_cluster", "body", m.EverouteCluster); err != nil {
		return err
	}

	if m.EverouteCluster != nil {
		if err := m.EverouteCluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("everoute_cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("everoute_cluster")
			}
			return err
		}
	}

	return nil
}

func (m *SecurityPolicy) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *SecurityPolicy) validateIngress(formats strfmt.Registry) error {
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

func (m *SecurityPolicy) validateIsBlocklist(formats strfmt.Registry) error {

	if err := validate.Required("is_blocklist", "body", m.IsBlocklist); err != nil {
		return err
	}

	return nil
}

func (m *SecurityPolicy) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *SecurityPolicy) validatePolicyMode(formats strfmt.Registry) error {
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

func (m *SecurityPolicy) validateStatistics(formats strfmt.Registry) error {
	if swag.IsZero(m.Statistics) { // not required
		return nil
	}

	if m.Statistics != nil {
		if err := m.Statistics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statistics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("statistics")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this security policy based on the context it is used
func (m *SecurityPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApplyTo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEgress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEverouteCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIngress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePolicyMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatistics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecurityPolicy) contextValidateApplyTo(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SecurityPolicy) contextValidateEgress(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SecurityPolicy) contextValidateEverouteCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.EverouteCluster != nil {
		if err := m.EverouteCluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("everoute_cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("everoute_cluster")
			}
			return err
		}
	}

	return nil
}

func (m *SecurityPolicy) contextValidateIngress(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SecurityPolicy) contextValidatePolicyMode(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SecurityPolicy) contextValidateStatistics(ctx context.Context, formats strfmt.Registry) error {

	if m.Statistics != nil {
		if err := m.Statistics.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statistics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("statistics")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityPolicy) UnmarshalBinary(b []byte) error {
	var res SecurityPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
