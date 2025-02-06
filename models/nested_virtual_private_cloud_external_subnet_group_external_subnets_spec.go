// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NestedVirtualPrivateCloudExternalSubnetGroupExternalSubnetsSpec nested virtual private cloud external subnet group external subnets spec
//
// swagger:model NestedVirtualPrivateCloudExternalSubnetGroupExternalSubnetsSpec
type NestedVirtualPrivateCloudExternalSubnetGroupExternalSubnetsSpec struct {

	// cidr
	// Required: true
	Cidr *string `json:"cidr"`

	// edge gateway id
	// Required: true
	EdgeGatewayID *string `json:"edge_gateway_id"`

	// floating ip cidr
	FloatingIPCidr *string `json:"floating_ip_cidr,omitempty"`

	// gateway
	// Required: true
	Gateway *string `json:"gateway"`

	// id
	// Required: true
	ID *string `json:"id"`

	// nat gateway cidr
	NatGatewayCidr *string `json:"nat_gateway_cidr,omitempty"`

	// router gateway cidr
	RouterGatewayCidr *string `json:"router_gateway_cidr,omitempty"`

	// vlan id
	// Required: true
	VlanID *string `json:"vlan_id"`

	MarshalOpts *NestedVirtualPrivateCloudExternalSubnetGroupExternalSubnetsSpecMarshalOpts `json:"-"`
}

type NestedVirtualPrivateCloudExternalSubnetGroupExternalSubnetsSpecMarshalOpts struct {
	Cidr_Explicit_Null_When_Empty bool

	EdgeGatewayID_Explicit_Null_When_Empty bool

	FloatingIPCidr_Explicit_Null_When_Empty bool

	Gateway_Explicit_Null_When_Empty bool

	ID_Explicit_Null_When_Empty bool

	NatGatewayCidr_Explicit_Null_When_Empty bool

	RouterGatewayCidr_Explicit_Null_When_Empty bool

	VlanID_Explicit_Null_When_Empty bool
}

func (m NestedVirtualPrivateCloudExternalSubnetGroupExternalSubnetsSpec) MarshalJSON() ([]byte, error) {
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

	// handle nullable field edge_gateway_id
	if m.EdgeGatewayID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"edge_gateway_id\":")
		bytes, err := swag.WriteJSON(m.EdgeGatewayID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.EdgeGatewayID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"edge_gateway_id\":null")
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

	// handle nullable field vlan_id
	if m.VlanID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vlan_id\":")
		bytes, err := swag.WriteJSON(m.VlanID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VlanID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vlan_id\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this nested virtual private cloud external subnet group external subnets spec
func (m *NestedVirtualPrivateCloudExternalSubnetGroupExternalSubnetsSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCidr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeGatewayID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGateway(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlanID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedVirtualPrivateCloudExternalSubnetGroupExternalSubnetsSpec) validateCidr(formats strfmt.Registry) error {

	if err := validate.Required("cidr", "body", m.Cidr); err != nil {
		return err
	}

	return nil
}

func (m *NestedVirtualPrivateCloudExternalSubnetGroupExternalSubnetsSpec) validateEdgeGatewayID(formats strfmt.Registry) error {

	if err := validate.Required("edge_gateway_id", "body", m.EdgeGatewayID); err != nil {
		return err
	}

	return nil
}

func (m *NestedVirtualPrivateCloudExternalSubnetGroupExternalSubnetsSpec) validateGateway(formats strfmt.Registry) error {

	if err := validate.Required("gateway", "body", m.Gateway); err != nil {
		return err
	}

	return nil
}

func (m *NestedVirtualPrivateCloudExternalSubnetGroupExternalSubnetsSpec) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *NestedVirtualPrivateCloudExternalSubnetGroupExternalSubnetsSpec) validateVlanID(formats strfmt.Registry) error {

	if err := validate.Required("vlan_id", "body", m.VlanID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this nested virtual private cloud external subnet group external subnets spec based on context it is used
func (m *NestedVirtualPrivateCloudExternalSubnetGroupExternalSubnetsSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NestedVirtualPrivateCloudExternalSubnetGroupExternalSubnetsSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedVirtualPrivateCloudExternalSubnetGroupExternalSubnetsSpec) UnmarshalBinary(b []byte) error {
	var res NestedVirtualPrivateCloudExternalSubnetGroupExternalSubnetsSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
