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

// IsolationPolicy isolation policy
//
// swagger:model IsolationPolicy
type IsolationPolicy struct {

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

	// labels
	Labels []*NestedLabel `json:"labels,omitempty"`

	// mode
	// Required: true
	Mode *IsolationMode `json:"mode"`

	// statistics
	Statistics *NestedSecurityPolicyStatistics `json:"statistics,omitempty"`

	// vm
	// Required: true
	VM *NestedVM `json:"vm"`

	MarshalOpts *IsolationPolicyMarshalOpts `json:"-"`
}

type IsolationPolicyMarshalOpts struct {
	EverouteCluster_Explicit_Null_When_Empty bool

	ID_Explicit_Null_When_Empty bool

	Mode_Explicit_Null_When_Empty bool

	Statistics_Explicit_Null_When_Empty bool

	VM_Explicit_Null_When_Empty bool
}

func (m IsolationPolicy) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

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

	// handle non nullable field labels with omitempty
	if swag.IsZero(m.Labels) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"labels\":")
		bytes, err := swag.WriteJSON(m.Labels)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field mode
	if m.Mode != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"mode\":")
		bytes, err := swag.WriteJSON(m.Mode)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Mode_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"mode\":null")
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

	// handle nullable field vm
	if m.VM != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm\":")
		bytes, err := swag.WriteJSON(m.VM)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VM_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this isolation policy
func (m *IsolationPolicy) Validate(formats strfmt.Registry) error {
	var res []error

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

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatistics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVM(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IsolationPolicy) validateEgress(formats strfmt.Registry) error {
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

func (m *IsolationPolicy) validateEverouteCluster(formats strfmt.Registry) error {

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

func (m *IsolationPolicy) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *IsolationPolicy) validateIngress(formats strfmt.Registry) error {
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

func (m *IsolationPolicy) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	for i := 0; i < len(m.Labels); i++ {
		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if m.Labels[i] != nil {
			if err := m.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IsolationPolicy) validateMode(formats strfmt.Registry) error {

	if err := validate.Required("mode", "body", m.Mode); err != nil {
		return err
	}

	if err := validate.Required("mode", "body", m.Mode); err != nil {
		return err
	}

	if m.Mode != nil {
		if err := m.Mode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mode")
			}
			return err
		}
	}

	return nil
}

func (m *IsolationPolicy) validateStatistics(formats strfmt.Registry) error {
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

func (m *IsolationPolicy) validateVM(formats strfmt.Registry) error {

	if err := validate.Required("vm", "body", m.VM); err != nil {
		return err
	}

	if m.VM != nil {
		if err := m.VM.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this isolation policy based on the context it is used
func (m *IsolationPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEgress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEverouteCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIngress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatistics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVM(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IsolationPolicy) contextValidateEgress(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IsolationPolicy) contextValidateEverouteCluster(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IsolationPolicy) contextValidateIngress(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IsolationPolicy) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Labels); i++ {

		if m.Labels[i] != nil {
			if err := m.Labels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IsolationPolicy) contextValidateMode(ctx context.Context, formats strfmt.Registry) error {

	if m.Mode != nil {
		if err := m.Mode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mode")
			}
			return err
		}
	}

	return nil
}

func (m *IsolationPolicy) contextValidateStatistics(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IsolationPolicy) contextValidateVM(ctx context.Context, formats strfmt.Registry) error {

	if m.VM != nil {
		if err := m.VM.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IsolationPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IsolationPolicy) UnmarshalBinary(b []byte) error {
	var res IsolationPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
