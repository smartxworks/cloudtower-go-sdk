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
)

// VirtualPrivateCloudSubnetUpdateDataParams virtual private cloud subnet update data params
//
// swagger:model VirtualPrivateCloudSubnetUpdateDataParams
type VirtualPrivateCloudSubnetUpdateDataParams struct {

	// cidr
	Cidr *string `json:"cidr,omitempty"`

	// description
	Description *string `json:"description,omitempty"`

	// gateway
	Gateway *string `json:"gateway,omitempty"`

	// ip pools
	IPPools []*VirtualPrivateCloudSubnetIPPoolParams `json:"ip_pools,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// route table id
	RouteTableID *string `json:"route_table_id,omitempty"`

	MarshalOpts *VirtualPrivateCloudSubnetUpdateDataParamsMarshalOpts `json:"-"`
}

type VirtualPrivateCloudSubnetUpdateDataParamsMarshalOpts struct {
	Cidr_Explicit_Null_When_Empty bool

	Description_Explicit_Null_When_Empty bool

	Gateway_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	RouteTableID_Explicit_Null_When_Empty bool
}

func (m VirtualPrivateCloudSubnetUpdateDataParams) MarshalJSON() ([]byte, error) {
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

	// handle non nullable field ip_pools with omitempty
	if swag.IsZero(m.IPPools) {
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

	// handle nullable field route_table_id
	if m.RouteTableID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"route_table_id\":")
		bytes, err := swag.WriteJSON(m.RouteTableID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.RouteTableID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"route_table_id\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this virtual private cloud subnet update data params
func (m *VirtualPrivateCloudSubnetUpdateDataParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPPools(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualPrivateCloudSubnetUpdateDataParams) validateIPPools(formats strfmt.Registry) error {
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

// ContextValidate validate this virtual private cloud subnet update data params based on the context it is used
func (m *VirtualPrivateCloudSubnetUpdateDataParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIPPools(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualPrivateCloudSubnetUpdateDataParams) contextValidateIPPools(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *VirtualPrivateCloudSubnetUpdateDataParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualPrivateCloudSubnetUpdateDataParams) UnmarshalBinary(b []byte) error {
	var res VirtualPrivateCloudSubnetUpdateDataParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
