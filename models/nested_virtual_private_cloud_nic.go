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

// NestedVirtualPrivateCloudNic nested virtual private cloud nic
//
// swagger:model NestedVirtualPrivateCloudNic
type NestedVirtualPrivateCloudNic struct {

	// floating ip
	FloatingIP *NestedVirtualPrivateCloudFloatingIP `json:"floating_ip,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// ip addresses
	// Required: true
	IPAddresses []string `json:"ip_addresses"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// vpc
	// Required: true
	Vpc *NestedVirtualPrivateCloud `json:"vpc"`

	// vpc subnet
	// Required: true
	VpcSubnet *NestedVirtualPrivateCloudSubnet `json:"vpc_subnet"`
}

// Validate validates this nested virtual private cloud nic
func (m *NestedVirtualPrivateCloudNic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFloatingIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpcSubnet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedVirtualPrivateCloudNic) validateFloatingIP(formats strfmt.Registry) error {
	if swag.IsZero(m.FloatingIP) { // not required
		return nil
	}

	if m.FloatingIP != nil {
		if err := m.FloatingIP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("floating_ip")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("floating_ip")
			}
			return err
		}
	}

	return nil
}

func (m *NestedVirtualPrivateCloudNic) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *NestedVirtualPrivateCloudNic) validateIPAddresses(formats strfmt.Registry) error {

	if err := validate.Required("ip_addresses", "body", m.IPAddresses); err != nil {
		return err
	}

	return nil
}

func (m *NestedVirtualPrivateCloudNic) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *NestedVirtualPrivateCloudNic) validateVpc(formats strfmt.Registry) error {

	if err := validate.Required("vpc", "body", m.Vpc); err != nil {
		return err
	}

	if m.Vpc != nil {
		if err := m.Vpc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vpc")
			}
			return err
		}
	}

	return nil
}

func (m *NestedVirtualPrivateCloudNic) validateVpcSubnet(formats strfmt.Registry) error {

	if err := validate.Required("vpc_subnet", "body", m.VpcSubnet); err != nil {
		return err
	}

	if m.VpcSubnet != nil {
		if err := m.VpcSubnet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpc_subnet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vpc_subnet")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nested virtual private cloud nic based on the context it is used
func (m *NestedVirtualPrivateCloudNic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFloatingIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVpc(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVpcSubnet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedVirtualPrivateCloudNic) contextValidateFloatingIP(ctx context.Context, formats strfmt.Registry) error {

	if m.FloatingIP != nil {
		if err := m.FloatingIP.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("floating_ip")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("floating_ip")
			}
			return err
		}
	}

	return nil
}

func (m *NestedVirtualPrivateCloudNic) contextValidateVpc(ctx context.Context, formats strfmt.Registry) error {

	if m.Vpc != nil {
		if err := m.Vpc.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vpc")
			}
			return err
		}
	}

	return nil
}

func (m *NestedVirtualPrivateCloudNic) contextValidateVpcSubnet(ctx context.Context, formats strfmt.Registry) error {

	if m.VpcSubnet != nil {
		if err := m.VpcSubnet.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpc_subnet")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vpc_subnet")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NestedVirtualPrivateCloudNic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedVirtualPrivateCloudNic) UnmarshalBinary(b []byte) error {
	var res NestedVirtualPrivateCloudNic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
