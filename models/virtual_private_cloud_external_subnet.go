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

// VirtualPrivateCloudExternalSubnet virtual private cloud external subnet
//
// swagger:model VirtualPrivateCloudExternalSubnet
type VirtualPrivateCloudExternalSubnet struct {

	// cidr
	// Required: true
	Cidr *string `json:"cidr"`

	// description
	Description *string `json:"description,omitempty"`

	// edge gateway
	EdgeGateway *NestedVirtualPrivateCloudEdgeGateway `json:"edge_gateway,omitempty"`

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// exclusive
	// Required: true
	Exclusive *bool `json:"exclusive"`

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

	MarshalOpts *VirtualPrivateCloudExternalSubnetMarshalOpts `json:"-"`
}

type VirtualPrivateCloudExternalSubnetMarshalOpts struct {
	Cidr_Explicit_Null_When_Empty bool

	Description_Explicit_Null_When_Empty bool

	EdgeGateway_Explicit_Null_When_Empty bool

	EntityAsyncStatus_Explicit_Null_When_Empty bool

	Exclusive_Explicit_Null_When_Empty bool

	FloatingIPCidr_Explicit_Null_When_Empty bool

	Gateway_Explicit_Null_When_Empty bool

	ID_Explicit_Null_When_Empty bool

	LocalID_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	NatGatewayCidr_Explicit_Null_When_Empty bool

	RouterGatewayCidr_Explicit_Null_When_Empty bool

	Vlan_Explicit_Null_When_Empty bool

	Vpc_Explicit_Null_When_Empty bool
}

func (m VirtualPrivateCloudExternalSubnet) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field cidr
	if m.Cidr != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cidr\":")
		bytes, err := swag.WriteJSON(m.Cidr)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Cidr_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cidr\":null")
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

	// handle nullable field edge_gateway
	if m.EdgeGateway != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"edge_gateway\":")
		bytes, err := swag.WriteJSON(m.EdgeGateway)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.EdgeGateway_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"edge_gateway\":null")
		first = false
	}

	// handle nullable field entityAsyncStatus
	if m.EntityAsyncStatus != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"entityAsyncStatus\":")
		bytes, err := swag.WriteJSON(m.EntityAsyncStatus)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.EntityAsyncStatus_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"entityAsyncStatus\":null")
		first = false
	}

	// handle nullable field exclusive
	if m.Exclusive != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"exclusive\":")
		bytes, err := swag.WriteJSON(m.Exclusive)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Exclusive_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"exclusive\":null")
		first = false
	}

	// handle nullable field floating_ip_cidr
	if m.FloatingIPCidr != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"floating_ip_cidr\":")
		bytes, err := swag.WriteJSON(m.FloatingIPCidr)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.FloatingIPCidr_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"floating_ip_cidr\":null")
		first = false
	}

	// handle non nullable field floating_ips with omitempty
	if swag.IsZero(m.FloatingIps) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"floating_ips\":")
		bytes, err := swag.WriteJSON(m.FloatingIps)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field gateway
	if m.Gateway != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"gateway\":")
		bytes, err := swag.WriteJSON(m.Gateway)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Gateway_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"gateway\":null")
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

	// handle nullable field local_id
	if m.LocalID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id\":")
		bytes, err := swag.WriteJSON(m.LocalID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.LocalID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id\":null")
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

	// handle nullable field nat_gateway_cidr
	if m.NatGatewayCidr != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"nat_gateway_cidr\":")
		bytes, err := swag.WriteJSON(m.NatGatewayCidr)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NatGatewayCidr_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"nat_gateway_cidr\":null")
		first = false
	}

	// handle non nullable field nat_gateways with omitempty
	if swag.IsZero(m.NatGateways) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"nat_gateways\":")
		bytes, err := swag.WriteJSON(m.NatGateways)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field router_gateway_cidr
	if m.RouterGatewayCidr != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"router_gateway_cidr\":")
		bytes, err := swag.WriteJSON(m.RouterGatewayCidr)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.RouterGatewayCidr_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"router_gateway_cidr\":null")
		first = false
	}

	// handle non nullable field router_gateways with omitempty
	if swag.IsZero(m.RouterGateways) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"router_gateways\":")
		bytes, err := swag.WriteJSON(m.RouterGateways)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field vlan
	if m.Vlan != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vlan\":")
		bytes, err := swag.WriteJSON(m.Vlan)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Vlan_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vlan\":null")
		first = false
	}

	// handle nullable field vpc
	if m.Vpc != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vpc\":")
		bytes, err := swag.WriteJSON(m.Vpc)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Vpc_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vpc\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this virtual private cloud external subnet
func (m *VirtualPrivateCloudExternalSubnet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCidr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeGateway(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExclusive(formats); err != nil {
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

func (m *VirtualPrivateCloudExternalSubnet) validateEdgeGateway(formats strfmt.Registry) error {
	if swag.IsZero(m.EdgeGateway) { // not required
		return nil
	}

	if m.EdgeGateway != nil {
		if err := m.EdgeGateway.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edge_gateway")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("edge_gateway")
			}
			return err
		}
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

func (m *VirtualPrivateCloudExternalSubnet) validateExclusive(formats strfmt.Registry) error {

	if err := validate.Required("exclusive", "body", m.Exclusive); err != nil {
		return err
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

	if err := m.contextValidateEdgeGateway(ctx, formats); err != nil {
		res = append(res, err)
	}

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

func (m *VirtualPrivateCloudExternalSubnet) contextValidateEdgeGateway(ctx context.Context, formats strfmt.Registry) error {

	if m.EdgeGateway != nil {
		if err := m.EdgeGateway.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edge_gateway")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("edge_gateway")
			}
			return err
		}
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
