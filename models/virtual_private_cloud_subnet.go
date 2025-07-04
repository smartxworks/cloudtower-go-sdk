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

// VirtualPrivateCloudSubnet virtual private cloud subnet
//
// swagger:model VirtualPrivateCloudSubnet
type VirtualPrivateCloudSubnet struct {

	// cidr
	// Required: true
	Cidr *string `json:"cidr"`

	// description
	Description *string `json:"description,omitempty"`

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// gateway
	Gateway *string `json:"gateway,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// ip pools
	IPPools []*NestedVpcSubnetIPPooType `json:"ip_pools,omitempty"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// route table
	// Required: true
	RouteTable *NestedVirtualPrivateCloudRouteTable `json:"route_table"`

	// total ip count
	TotalIPCount *int32 `json:"total_ip_count,omitempty"`

	// unused ip count
	UnusedIPCount *int32 `json:"unused_ip_count,omitempty"`

	// vpc
	// Required: true
	Vpc *NestedVirtualPrivateCloud `json:"vpc"`

	MarshalOpts *VirtualPrivateCloudSubnetMarshalOpts `json:"-"`
}

type VirtualPrivateCloudSubnetMarshalOpts struct {
	Cidr_Explicit_Null_When_Empty bool

	Description_Explicit_Null_When_Empty bool

	EntityAsyncStatus_Explicit_Null_When_Empty bool

	Gateway_Explicit_Null_When_Empty bool

	ID_Explicit_Null_When_Empty bool

	IPPools_Explicit_Null_When_Empty bool

	LocalID_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	RouteTable_Explicit_Null_When_Empty bool

	TotalIPCount_Explicit_Null_When_Empty bool

	UnusedIPCount_Explicit_Null_When_Empty bool

	Vpc_Explicit_Null_When_Empty bool
}

func (m VirtualPrivateCloudSubnet) MarshalJSON() ([]byte, error) {
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

	// handle non nullable field ip_pools with omitempty
	if !swag.IsZero(m.IPPools) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ip_pools\":")
		bytes, err := swag.WriteJSON(m.IPPools)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
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

	// handle nullable field route_table
	if m.RouteTable != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"route_table\":")
		bytes, err := swag.WriteJSON(m.RouteTable)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.RouteTable_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"route_table\":null")
		first = false
	}

	// handle nullable field total_ip_count
	if m.TotalIPCount != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"total_ip_count\":")
		bytes, err := swag.WriteJSON(m.TotalIPCount)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.TotalIPCount_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"total_ip_count\":null")
		first = false
	}

	// handle nullable field unused_ip_count
	if m.UnusedIPCount != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"unused_ip_count\":")
		bytes, err := swag.WriteJSON(m.UnusedIPCount)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.UnusedIPCount_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"unused_ip_count\":null")
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

// Validate validates this virtual private cloud subnet
func (m *VirtualPrivateCloudSubnet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCidr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPPools(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteTable(formats); err != nil {
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

func (m *VirtualPrivateCloudSubnet) validateCidr(formats strfmt.Registry) error {

	if err := validate.Required("cidr", "body", m.Cidr); err != nil {
		return err
	}

	return nil
}

func (m *VirtualPrivateCloudSubnet) validateEntityAsyncStatus(formats strfmt.Registry) error {
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

func (m *VirtualPrivateCloudSubnet) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *VirtualPrivateCloudSubnet) validateIPPools(formats strfmt.Registry) error {
	if swag.IsZero(m.IPPools) { // not required
		return nil
	}

	for i := 0; i < len(m.IPPools); i++ {
		if swag.IsZero(m.IPPools[i]) { // not required
			continue
		}

		if m.IPPools[i] != nil {
			if err := m.IPPools[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ip_pools" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ip_pools" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudSubnet) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *VirtualPrivateCloudSubnet) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VirtualPrivateCloudSubnet) validateRouteTable(formats strfmt.Registry) error {

	if err := validate.Required("route_table", "body", m.RouteTable); err != nil {
		return err
	}

	if m.RouteTable != nil {
		if err := m.RouteTable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("route_table")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("route_table")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudSubnet) validateVpc(formats strfmt.Registry) error {

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

// ContextValidate validate this virtual private cloud subnet based on the context it is used
func (m *VirtualPrivateCloudSubnet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIPPools(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRouteTable(ctx, formats); err != nil {
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

func (m *VirtualPrivateCloudSubnet) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VirtualPrivateCloudSubnet) contextValidateIPPools(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IPPools); i++ {

		if m.IPPools[i] != nil {
			if err := m.IPPools[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ip_pools" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ip_pools" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudSubnet) contextValidateRouteTable(ctx context.Context, formats strfmt.Registry) error {

	if m.RouteTable != nil {
		if err := m.RouteTable.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("route_table")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("route_table")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudSubnet) contextValidateVpc(ctx context.Context, formats strfmt.Registry) error {

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
func (m *VirtualPrivateCloudSubnet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualPrivateCloudSubnet) UnmarshalBinary(b []byte) error {
	var res VirtualPrivateCloudSubnet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
