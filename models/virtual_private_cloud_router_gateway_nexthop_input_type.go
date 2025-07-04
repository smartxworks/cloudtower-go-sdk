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

// VirtualPrivateCloudRouterGatewayNexthopInputType virtual private cloud router gateway nexthop input type
//
// swagger:model VirtualPrivateCloudRouterGatewayNexthopInputType
type VirtualPrivateCloudRouterGatewayNexthopInputType struct {

	// edge gateway id
	// Required: true
	EdgeGatewayID *string `json:"edge_gateway_id"`

	// nexthop
	Nexthop *string `json:"nexthop,omitempty"`

	MarshalOpts *VirtualPrivateCloudRouterGatewayNexthopInputTypeMarshalOpts `json:"-"`
}

type VirtualPrivateCloudRouterGatewayNexthopInputTypeMarshalOpts struct {
	EdgeGatewayID_Explicit_Null_When_Empty bool

	Nexthop_Explicit_Null_When_Empty bool
}

func (m VirtualPrivateCloudRouterGatewayNexthopInputType) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

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

	// handle nullable field nexthop
	if m.Nexthop != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"nexthop\":")
		bytes, err := swag.WriteJSON(m.Nexthop)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Nexthop_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"nexthop\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this virtual private cloud router gateway nexthop input type
func (m *VirtualPrivateCloudRouterGatewayNexthopInputType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEdgeGatewayID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualPrivateCloudRouterGatewayNexthopInputType) validateEdgeGatewayID(formats strfmt.Registry) error {

	if err := validate.Required("edge_gateway_id", "body", m.EdgeGatewayID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this virtual private cloud router gateway nexthop input type based on context it is used
func (m *VirtualPrivateCloudRouterGatewayNexthopInputType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VirtualPrivateCloudRouterGatewayNexthopInputType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualPrivateCloudRouterGatewayNexthopInputType) UnmarshalBinary(b []byte) error {
	var res VirtualPrivateCloudRouterGatewayNexthopInputType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
