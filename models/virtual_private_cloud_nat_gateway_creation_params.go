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

// VirtualPrivateCloudNatGatewayCreationParams virtual private cloud nat gateway creation params
//
// swagger:model VirtualPrivateCloudNatGatewayCreationParams
type VirtualPrivateCloudNatGatewayCreationParams struct {

	// dnat rules
	DnatRules []*VirtualPrivateCloudDnatRuleParams `json:"dnat_rules,omitempty"`

	// enable dnat
	// Required: true
	EnableDnat *bool `json:"enable_dnat"`

	// enable snat
	// Required: true
	EnableSnat *bool `json:"enable_snat"`

	// external ip
	ExternalIP *string `json:"external_ip,omitempty"`

	// external subnet id
	// Required: true
	ExternalSubnetID *string `json:"external_subnet_id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// vpc id
	// Required: true
	VpcID *string `json:"vpc_id"`

	MarshalOpts *VirtualPrivateCloudNatGatewayCreationParamsMarshalOpts `json:"-"`
}

type VirtualPrivateCloudNatGatewayCreationParamsMarshalOpts struct {
	EnableDnat_Explicit_Null_When_Empty bool

	EnableSnat_Explicit_Null_When_Empty bool

	ExternalIP_Explicit_Null_When_Empty bool

	ExternalSubnetID_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	VpcID_Explicit_Null_When_Empty bool
}

func (m VirtualPrivateCloudNatGatewayCreationParams) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle non nullable field dnat_rules with omitempty
	if swag.IsZero(m.DnatRules) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"dnat_rules\":")
		bytes, err := swag.WriteJSON(m.DnatRules)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field enable_dnat
	if m.EnableDnat != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"enable_dnat\":")
		bytes, err := swag.WriteJSON(m.EnableDnat)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.EnableDnat_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"enable_dnat\":null")
		first = false
	}

	// handle nullable field enable_snat
	if m.EnableSnat != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"enable_snat\":")
		bytes, err := swag.WriteJSON(m.EnableSnat)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.EnableSnat_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"enable_snat\":null")
		first = false
	}

	// handle nullable field external_ip
	if m.ExternalIP != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"external_ip\":")
		bytes, err := swag.WriteJSON(m.ExternalIP)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ExternalIP_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"external_ip\":null")
		first = false
	}

	// handle nullable field external_subnet_id
	if m.ExternalSubnetID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"external_subnet_id\":")
		bytes, err := swag.WriteJSON(m.ExternalSubnetID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ExternalSubnetID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"external_subnet_id\":null")
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

	// handle nullable field vpc_id
	if m.VpcID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vpc_id\":")
		bytes, err := swag.WriteJSON(m.VpcID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VpcID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vpc_id\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this virtual private cloud nat gateway creation params
func (m *VirtualPrivateCloudNatGatewayCreationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDnatRules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnableDnat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnableSnat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalSubnetID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpcID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualPrivateCloudNatGatewayCreationParams) validateDnatRules(formats strfmt.Registry) error {
	if swag.IsZero(m.DnatRules) { // not required
		return nil
	}

	for i := 0; i < len(m.DnatRules); i++ {
		if swag.IsZero(m.DnatRules[i]) { // not required
			continue
		}

		if m.DnatRules[i] != nil {
			if err := m.DnatRules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dnat_rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dnat_rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudNatGatewayCreationParams) validateEnableDnat(formats strfmt.Registry) error {

	if err := validate.Required("enable_dnat", "body", m.EnableDnat); err != nil {
		return err
	}

	return nil
}

func (m *VirtualPrivateCloudNatGatewayCreationParams) validateEnableSnat(formats strfmt.Registry) error {

	if err := validate.Required("enable_snat", "body", m.EnableSnat); err != nil {
		return err
	}

	return nil
}

func (m *VirtualPrivateCloudNatGatewayCreationParams) validateExternalSubnetID(formats strfmt.Registry) error {

	if err := validate.Required("external_subnet_id", "body", m.ExternalSubnetID); err != nil {
		return err
	}

	return nil
}

func (m *VirtualPrivateCloudNatGatewayCreationParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VirtualPrivateCloudNatGatewayCreationParams) validateVpcID(formats strfmt.Registry) error {

	if err := validate.Required("vpc_id", "body", m.VpcID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this virtual private cloud nat gateway creation params based on the context it is used
func (m *VirtualPrivateCloudNatGatewayCreationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDnatRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualPrivateCloudNatGatewayCreationParams) contextValidateDnatRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DnatRules); i++ {

		if m.DnatRules[i] != nil {
			if err := m.DnatRules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dnat_rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dnat_rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VirtualPrivateCloudNatGatewayCreationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualPrivateCloudNatGatewayCreationParams) UnmarshalBinary(b []byte) error {
	var res VirtualPrivateCloudNatGatewayCreationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
