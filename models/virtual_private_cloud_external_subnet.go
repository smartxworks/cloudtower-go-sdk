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
	"github.com/go-openapi/validate"
)

// VirtualPrivateCloudExternalSubnet virtual private cloud external subnet
//
// swagger:model VirtualPrivateCloudExternalSubnet
type VirtualPrivateCloudExternalSubnet struct {

	// cidr
	// Required: true
	Cidr *string `json:"cidr"`

	// description
	Description *string `json:"description,omitempty"`

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// floating ip cidr
	FloatingIPCidr *string `json:"floating_ip_cidr,omitempty"`

	// floating ips
	FloatingIps []*NestedVirtualPrivateCloudFloatingIP `json:"floating_ips,omitempty"`

	// gateway
	// Required: true
	Gateway *string `json:"gateway"`

	// id
	// Required: true
	ID *string `json:"id"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// nat gateway cidr
	NatGatewayCidr *string `json:"nat_gateway_cidr,omitempty"`

	// nat gateways
	NatGateways []*NestedVirtualPrivateCloudNatGateway `json:"nat_gateways,omitempty"`

	// router gateway cidr
	RouterGatewayCidr *string `json:"router_gateway_cidr,omitempty"`

	// router gateways
	RouterGateways []*NestedVirtualPrivateCloudRouterGateway `json:"router_gateways,omitempty"`

	// vlan
	// Required: true
	Vlan *NestedVlan `json:"vlan"`

	// vpc
	Vpc *NestedVirtualPrivateCloud `json:"vpc,omitempty"`
}

// Validate validates this virtual private cloud external subnet
func (m *VirtualPrivateCloudExternalSubnet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCidr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFloatingIps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGateway(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNatGateways(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouterGateways(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualPrivateCloudExternalSubnet) validateCidr(formats strfmt.Registry) error {

	if err := validate.Required("cidr", "body", m.Cidr); err != nil {
		return err
	}

	return nil
}

func (m *VirtualPrivateCloudExternalSubnet) validateEntityAsyncStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatus) { // not required
		return nil
	}

	if m.EntityAsyncStatus != nil {
		if err := m.EntityAsyncStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudExternalSubnet) validateFloatingIps(formats strfmt.Registry) error {
	if swag.IsZero(m.FloatingIps) { // not required
		return nil
	}

	for i := 0; i < len(m.FloatingIps); i++ {
		if swag.IsZero(m.FloatingIps[i]) { // not required
			continue
		}

		if m.FloatingIps[i] != nil {
			if err := m.FloatingIps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("floating_ips" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("floating_ips" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudExternalSubnet) validateGateway(formats strfmt.Registry) error {

	if err := validate.Required("gateway", "body", m.Gateway); err != nil {
		return err
	}

	return nil
}

func (m *VirtualPrivateCloudExternalSubnet) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *VirtualPrivateCloudExternalSubnet) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *VirtualPrivateCloudExternalSubnet) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VirtualPrivateCloudExternalSubnet) validateNatGateways(formats strfmt.Registry) error {
	if swag.IsZero(m.NatGateways) { // not required
		return nil
	}

	for i := 0; i < len(m.NatGateways); i++ {
		if swag.IsZero(m.NatGateways[i]) { // not required
			continue
		}

		if m.NatGateways[i] != nil {
			if err := m.NatGateways[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nat_gateways" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nat_gateways" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudExternalSubnet) validateRouterGateways(formats strfmt.Registry) error {
	if swag.IsZero(m.RouterGateways) { // not required
		return nil
	}

	for i := 0; i < len(m.RouterGateways); i++ {
		if swag.IsZero(m.RouterGateways[i]) { // not required
			continue
		}

		if m.RouterGateways[i] != nil {
			if err := m.RouterGateways[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("router_gateways" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("router_gateways" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudExternalSubnet) validateVlan(formats strfmt.Registry) error {

	if err := validate.Required("vlan", "body", m.Vlan); err != nil {
		return err
	}

	if m.Vlan != nil {
		if err := m.Vlan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlan")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlan")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudExternalSubnet) validateVpc(formats strfmt.Registry) error {
	if swag.IsZero(m.Vpc) { // not required
		return nil
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

// ContextValidate validate this virtual private cloud external subnet based on the context it is used
func (m *VirtualPrivateCloudExternalSubnet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFloatingIps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNatGateways(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRouterGateways(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVlan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVpc(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualPrivateCloudExternalSubnet) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.EntityAsyncStatus != nil {
		if err := m.EntityAsyncStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudExternalSubnet) contextValidateFloatingIps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FloatingIps); i++ {

		if m.FloatingIps[i] != nil {
			if err := m.FloatingIps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("floating_ips" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("floating_ips" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudExternalSubnet) contextValidateNatGateways(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NatGateways); i++ {

		if m.NatGateways[i] != nil {
			if err := m.NatGateways[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nat_gateways" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nat_gateways" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudExternalSubnet) contextValidateRouterGateways(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RouterGateways); i++ {

		if m.RouterGateways[i] != nil {
			if err := m.RouterGateways[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("router_gateways" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("router_gateways" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudExternalSubnet) contextValidateVlan(ctx context.Context, formats strfmt.Registry) error {

	if m.Vlan != nil {
		if err := m.Vlan.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlan")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlan")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudExternalSubnet) contextValidateVpc(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *VirtualPrivateCloudExternalSubnet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualPrivateCloudExternalSubnet) UnmarshalBinary(b []byte) error {
	var res VirtualPrivateCloudExternalSubnet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
