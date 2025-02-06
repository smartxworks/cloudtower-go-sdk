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

// VirtualPrivateCloudExternalSubnetGroup virtual private cloud external subnet group
//
// swagger:model VirtualPrivateCloudExternalSubnetGroup
type VirtualPrivateCloudExternalSubnetGroup struct {

	// description
	Description *string `json:"description,omitempty"`

	// edge gateway group
	// Required: true
	EdgeGatewayGroup *NestedVirtualPrivateCloudEdgeGatewayGroup `json:"edge_gateway_group"`

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// exclusive
	// Required: true
	Exclusive *bool `json:"exclusive"`

	// external subnets spec
	// Required: true
	ExternalSubnetsSpec []*NestedVirtualPrivateCloudExternalSubnetGroupExternalSubnetsSpec `json:"external_subnets_spec"`

	// floating ips
	FloatingIps []*NestedVirtualPrivateCloudFloatingIP `json:"floating_ips,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// nat gateways
	NatGateways []*NestedVirtualPrivateCloudNatGateway `json:"nat_gateways,omitempty"`

	// router gateways
	RouterGateways []*NestedVirtualPrivateCloudRouterGateway `json:"router_gateways,omitempty"`

	// shared in edge gateway group
	// Required: true
	SharedInEdgeGatewayGroup *bool `json:"shared_in_edge_gateway_group"`

	// vpc
	Vpc *NestedVirtualPrivateCloud `json:"vpc,omitempty"`

	MarshalOpts *VirtualPrivateCloudExternalSubnetGroupMarshalOpts `json:"-"`
}

type VirtualPrivateCloudExternalSubnetGroupMarshalOpts struct {
	Description_Explicit_Null_When_Empty bool

	EdgeGatewayGroup_Explicit_Null_When_Empty bool

	EntityAsyncStatus_Explicit_Null_When_Empty bool

	Exclusive_Explicit_Null_When_Empty bool

	ExternalSubnetsSpec_Explicit_Null_When_Empty bool

	FloatingIps_Explicit_Null_When_Empty bool

	ID_Explicit_Null_When_Empty bool

	LocalID_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	NatGateways_Explicit_Null_When_Empty bool

	RouterGateways_Explicit_Null_When_Empty bool

	SharedInEdgeGatewayGroup_Explicit_Null_When_Empty bool

	Vpc_Explicit_Null_When_Empty bool
}

func (m VirtualPrivateCloudExternalSubnetGroup) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

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

	// handle nullable field edge_gateway_group
	if m.EdgeGatewayGroup != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"edge_gateway_group\":")
		bytes, err := swag.WriteJSON(m.EdgeGatewayGroup)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.EdgeGatewayGroup_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"edge_gateway_group\":null")
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

	// handle non nullable field external_subnets_spec without omitempty
	if !first {
		b.WriteString(",")
	}
	b.WriteString("\"external_subnets_spec\":")
	{
		bytes, err := swag.WriteJSON(m.ExternalSubnetsSpec)
		if err != nil {
			return nil, err
		}
	}
	b.Write(bytes)
	first = false

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

	// handle nullable field shared_in_edge_gateway_group
	if m.SharedInEdgeGatewayGroup != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"shared_in_edge_gateway_group\":")
		bytes, err := swag.WriteJSON(m.SharedInEdgeGatewayGroup)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.SharedInEdgeGatewayGroup_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"shared_in_edge_gateway_group\":null")
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

// Validate validates this virtual private cloud external subnet group
func (m *VirtualPrivateCloudExternalSubnetGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEdgeGatewayGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExclusive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalSubnetsSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFloatingIps(formats); err != nil {
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

	if err := m.validateSharedInEdgeGatewayGroup(formats); err != nil {
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

func (m *VirtualPrivateCloudExternalSubnetGroup) validateEdgeGatewayGroup(formats strfmt.Registry) error {

	if err := validate.Required("edge_gateway_group", "body", m.EdgeGatewayGroup); err != nil {
		return err
	}

	if m.EdgeGatewayGroup != nil {
		if err := m.EdgeGatewayGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edge_gateway_group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("edge_gateway_group")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudExternalSubnetGroup) validateEntityAsyncStatus(formats strfmt.Registry) error {
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

func (m *VirtualPrivateCloudExternalSubnetGroup) validateExclusive(formats strfmt.Registry) error {

	if err := validate.Required("exclusive", "body", m.Exclusive); err != nil {
		return err
	}

	return nil
}

func (m *VirtualPrivateCloudExternalSubnetGroup) validateExternalSubnetsSpec(formats strfmt.Registry) error {

	if err := validate.Required("external_subnets_spec", "body", m.ExternalSubnetsSpec); err != nil {
		return err
	}

	for i := 0; i < len(m.ExternalSubnetsSpec); i++ {
		if swag.IsZero(m.ExternalSubnetsSpec[i]) { // not required
			continue
		}

		if m.ExternalSubnetsSpec[i] != nil {
			if err := m.ExternalSubnetsSpec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("external_subnets_spec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("external_subnets_spec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudExternalSubnetGroup) validateFloatingIps(formats strfmt.Registry) error {
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

func (m *VirtualPrivateCloudExternalSubnetGroup) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *VirtualPrivateCloudExternalSubnetGroup) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *VirtualPrivateCloudExternalSubnetGroup) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VirtualPrivateCloudExternalSubnetGroup) validateNatGateways(formats strfmt.Registry) error {
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

func (m *VirtualPrivateCloudExternalSubnetGroup) validateRouterGateways(formats strfmt.Registry) error {
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

func (m *VirtualPrivateCloudExternalSubnetGroup) validateSharedInEdgeGatewayGroup(formats strfmt.Registry) error {

	if err := validate.Required("shared_in_edge_gateway_group", "body", m.SharedInEdgeGatewayGroup); err != nil {
		return err
	}

	return nil
}

func (m *VirtualPrivateCloudExternalSubnetGroup) validateVpc(formats strfmt.Registry) error {
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

// ContextValidate validate this virtual private cloud external subnet group based on the context it is used
func (m *VirtualPrivateCloudExternalSubnetGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEdgeGatewayGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExternalSubnetsSpec(ctx, formats); err != nil {
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

	if err := m.contextValidateVpc(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualPrivateCloudExternalSubnetGroup) contextValidateEdgeGatewayGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.EdgeGatewayGroup != nil {
		if err := m.EdgeGatewayGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edge_gateway_group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("edge_gateway_group")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudExternalSubnetGroup) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VirtualPrivateCloudExternalSubnetGroup) contextValidateExternalSubnetsSpec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExternalSubnetsSpec); i++ {

		if m.ExternalSubnetsSpec[i] != nil {
			if err := m.ExternalSubnetsSpec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("external_subnets_spec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("external_subnets_spec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudExternalSubnetGroup) contextValidateFloatingIps(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VirtualPrivateCloudExternalSubnetGroup) contextValidateNatGateways(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VirtualPrivateCloudExternalSubnetGroup) contextValidateRouterGateways(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VirtualPrivateCloudExternalSubnetGroup) contextValidateVpc(ctx context.Context, formats strfmt.Registry) error {

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
func (m *VirtualPrivateCloudExternalSubnetGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualPrivateCloudExternalSubnetGroup) UnmarshalBinary(b []byte) error {
	var res VirtualPrivateCloudExternalSubnetGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
