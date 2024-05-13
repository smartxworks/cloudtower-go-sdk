// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// VirtualPrivateCloudRouteNextHopType virtual private cloud route next hop type
//
// swagger:model VirtualPrivateCloudRouteNextHopType
type VirtualPrivateCloudRouteNextHopType string

func NewVirtualPrivateCloudRouteNextHopType(value VirtualPrivateCloudRouteNextHopType) *VirtualPrivateCloudRouteNextHopType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VirtualPrivateCloudRouteNextHopType.
func (m VirtualPrivateCloudRouteNextHopType) Pointer() *VirtualPrivateCloudRouteNextHopType {
	return &m
}

const (

	// VirtualPrivateCloudRouteNextHopTypeNATGATEWAY captures enum value "NAT_GATEWAY"
	VirtualPrivateCloudRouteNextHopTypeNATGATEWAY VirtualPrivateCloudRouteNextHopType = "NAT_GATEWAY"

	// VirtualPrivateCloudRouteNextHopTypeROUTERGATEWAY captures enum value "ROUTER_GATEWAY"
	VirtualPrivateCloudRouteNextHopTypeROUTERGATEWAY VirtualPrivateCloudRouteNextHopType = "ROUTER_GATEWAY"

	// VirtualPrivateCloudRouteNextHopTypeUNKNOWN captures enum value "UNKNOWN"
	VirtualPrivateCloudRouteNextHopTypeUNKNOWN VirtualPrivateCloudRouteNextHopType = "UNKNOWN"

	// VirtualPrivateCloudRouteNextHopTypeVPCPEERING captures enum value "VPC_PEERING"
	VirtualPrivateCloudRouteNextHopTypeVPCPEERING VirtualPrivateCloudRouteNextHopType = "VPC_PEERING"
)

// for schema
var virtualPrivateCloudRouteNextHopTypeEnum []interface{}

func init() {
	var res []VirtualPrivateCloudRouteNextHopType
	if err := json.Unmarshal([]byte(`["NAT_GATEWAY","ROUTER_GATEWAY","UNKNOWN","VPC_PEERING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		virtualPrivateCloudRouteNextHopTypeEnum = append(virtualPrivateCloudRouteNextHopTypeEnum, v)
	}
}

func (m VirtualPrivateCloudRouteNextHopType) validateVirtualPrivateCloudRouteNextHopTypeEnum(path, location string, value VirtualPrivateCloudRouteNextHopType) error {
	if err := validate.EnumCase(path, location, value, virtualPrivateCloudRouteNextHopTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this virtual private cloud route next hop type
func (m VirtualPrivateCloudRouteNextHopType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateVirtualPrivateCloudRouteNextHopTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this virtual private cloud route next hop type based on context it is used
func (m VirtualPrivateCloudRouteNextHopType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
