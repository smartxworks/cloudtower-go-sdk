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

// IsolationPolicyCreateParams isolation policy create params
//
// swagger:model IsolationPolicyCreateParams
type IsolationPolicyCreateParams struct {

	// egress
	Egress []*SecurityPolicyIngressEgressInput `json:"egress,omitempty"`

	// everoute cluster id
	// Required: true
	EverouteClusterID *string `json:"everoute_cluster_id"`

	// ingress
	Ingress []*SecurityPolicyIngressEgressInput `json:"ingress,omitempty"`

	// mode
	// Required: true
	Mode *IsolationMode `json:"mode"`

	// vm id
	// Required: true
	VMID *string `json:"vm_id"`

	MarshalOpts *IsolationPolicyCreateParamsMarshalOpts `json:"-"`
}

type IsolationPolicyCreateParamsMarshalOpts struct {
	Egress_Explicit_Null_When_Empty bool

	EverouteClusterID_Explicit_Null_When_Empty bool

	Ingress_Explicit_Null_When_Empty bool

	Mode_Explicit_Null_When_Empty bool

	VMID_Explicit_Null_When_Empty bool
}

func (m IsolationPolicyCreateParams) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle non nullable field egress with omitempty
	if !swag.IsZero(m.Egress) {
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

	// handle nullable field everoute_cluster_id
	if m.EverouteClusterID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"everoute_cluster_id\":")
		bytes, err := swag.WriteJSON(m.EverouteClusterID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.EverouteClusterID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"everoute_cluster_id\":null")
		first = false
	}

	// handle non nullable field ingress with omitempty
	if !swag.IsZero(m.Ingress) {
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

	// handle nullable field vm_id
	if m.VMID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_id\":")
		bytes, err := swag.WriteJSON(m.VMID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VMID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_id\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this isolation policy create params
func (m *IsolationPolicyCreateParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEgress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEverouteClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IsolationPolicyCreateParams) validateEgress(formats strfmt.Registry) error {
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

func (m *IsolationPolicyCreateParams) validateEverouteClusterID(formats strfmt.Registry) error {

	if err := validate.Required("everoute_cluster_id", "body", m.EverouteClusterID); err != nil {
		return err
	}

	return nil
}

func (m *IsolationPolicyCreateParams) validateIngress(formats strfmt.Registry) error {
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

func (m *IsolationPolicyCreateParams) validateMode(formats strfmt.Registry) error {

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

func (m *IsolationPolicyCreateParams) validateVMID(formats strfmt.Registry) error {

	if err := validate.Required("vm_id", "body", m.VMID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this isolation policy create params based on the context it is used
func (m *IsolationPolicyCreateParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEgress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIngress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IsolationPolicyCreateParams) contextValidateEgress(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IsolationPolicyCreateParams) contextValidateIngress(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IsolationPolicyCreateParams) contextValidateMode(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *IsolationPolicyCreateParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IsolationPolicyCreateParams) UnmarshalBinary(b []byte) error {
	var res IsolationPolicyCreateParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
