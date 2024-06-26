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
)

// VirtualPrivateCloudNicSnapshotWhereInput virtual private cloud nic snapshot where input
//
// swagger:model VirtualPrivateCloudNicSnapshotWhereInput
type VirtualPrivateCloudNicSnapshotWhereInput struct {

	// a n d
	AND []*VirtualPrivateCloudNicSnapshotWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*VirtualPrivateCloudNicSnapshotWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*VirtualPrivateCloudNicSnapshotWhereInput `json:"OR,omitempty"`

	// floating ip
	FloatingIP *VirtualPrivateCloudFloatingIPWhereInput `json:"floating_ip,omitempty"`

	// id
	ID *string `json:"id,omitempty"`

	// id contains
	IDContains *string `json:"id_contains,omitempty"`

	// id ends with
	IDEndsWith *string `json:"id_ends_with,omitempty"`

	// id gt
	IDGt *string `json:"id_gt,omitempty"`

	// id gte
	IDGte *string `json:"id_gte,omitempty"`

	// id in
	IDIn []string `json:"id_in,omitempty"`

	// id lt
	IDLt *string `json:"id_lt,omitempty"`

	// id lte
	IDLte *string `json:"id_lte,omitempty"`

	// id not
	IDNot *string `json:"id_not,omitempty"`

	// id not contains
	IDNotContains *string `json:"id_not_contains,omitempty"`

	// id not ends with
	IDNotEndsWith *string `json:"id_not_ends_with,omitempty"`

	// id not in
	IDNotIn []string `json:"id_not_in,omitempty"`

	// id not starts with
	IDNotStartsWith *string `json:"id_not_starts_with,omitempty"`

	// id starts with
	IDStartsWith *string `json:"id_starts_with,omitempty"`

	// local id
	LocalID *string `json:"local_id,omitempty"`

	// local id contains
	LocalIDContains *string `json:"local_id_contains,omitempty"`

	// local id ends with
	LocalIDEndsWith *string `json:"local_id_ends_with,omitempty"`

	// local id gt
	LocalIDGt *string `json:"local_id_gt,omitempty"`

	// local id gte
	LocalIDGte *string `json:"local_id_gte,omitempty"`

	// local id in
	LocalIDIn []string `json:"local_id_in,omitempty"`

	// local id lt
	LocalIDLt *string `json:"local_id_lt,omitempty"`

	// local id lte
	LocalIDLte *string `json:"local_id_lte,omitempty"`

	// local id not
	LocalIDNot *string `json:"local_id_not,omitempty"`

	// local id not contains
	LocalIDNotContains *string `json:"local_id_not_contains,omitempty"`

	// local id not ends with
	LocalIDNotEndsWith *string `json:"local_id_not_ends_with,omitempty"`

	// local id not in
	LocalIDNotIn []string `json:"local_id_not_in,omitempty"`

	// local id not starts with
	LocalIDNotStartsWith *string `json:"local_id_not_starts_with,omitempty"`

	// local id starts with
	LocalIDStartsWith *string `json:"local_id_starts_with,omitempty"`

	// mac address
	MacAddress *string `json:"mac_address,omitempty"`

	// mac address contains
	MacAddressContains *string `json:"mac_address_contains,omitempty"`

	// mac address ends with
	MacAddressEndsWith *string `json:"mac_address_ends_with,omitempty"`

	// mac address gt
	MacAddressGt *string `json:"mac_address_gt,omitempty"`

	// mac address gte
	MacAddressGte *string `json:"mac_address_gte,omitempty"`

	// mac address in
	MacAddressIn []string `json:"mac_address_in,omitempty"`

	// mac address lt
	MacAddressLt *string `json:"mac_address_lt,omitempty"`

	// mac address lte
	MacAddressLte *string `json:"mac_address_lte,omitempty"`

	// mac address not
	MacAddressNot *string `json:"mac_address_not,omitempty"`

	// mac address not contains
	MacAddressNotContains *string `json:"mac_address_not_contains,omitempty"`

	// mac address not ends with
	MacAddressNotEndsWith *string `json:"mac_address_not_ends_with,omitempty"`

	// mac address not in
	MacAddressNotIn []string `json:"mac_address_not_in,omitempty"`

	// mac address not starts with
	MacAddressNotStartsWith *string `json:"mac_address_not_starts_with,omitempty"`

	// mac address starts with
	MacAddressStartsWith *string `json:"mac_address_starts_with,omitempty"`

	// vpc
	Vpc *VirtualPrivateCloudWhereInput `json:"vpc,omitempty"`

	// vpc nic
	VpcNic *VirtualPrivateCloudNicWhereInput `json:"vpc_nic,omitempty"`

	// vpc subnet
	VpcSubnet *VirtualPrivateCloudSubnetWhereInput `json:"vpc_subnet,omitempty"`
}

// Validate validates this virtual private cloud nic snapshot where input
func (m *VirtualPrivateCloudNicSnapshotWhereInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAND(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNOT(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOR(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFloatingIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpcNic(formats); err != nil {
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

func (m *VirtualPrivateCloudNicSnapshotWhereInput) validateAND(formats strfmt.Registry) error {
	if swag.IsZero(m.AND) { // not required
		return nil
	}

	for i := 0; i < len(m.AND); i++ {
		if swag.IsZero(m.AND[i]) { // not required
			continue
		}

		if m.AND[i] != nil {
			if err := m.AND[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudNicSnapshotWhereInput) validateNOT(formats strfmt.Registry) error {
	if swag.IsZero(m.NOT) { // not required
		return nil
	}

	for i := 0; i < len(m.NOT); i++ {
		if swag.IsZero(m.NOT[i]) { // not required
			continue
		}

		if m.NOT[i] != nil {
			if err := m.NOT[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudNicSnapshotWhereInput) validateOR(formats strfmt.Registry) error {
	if swag.IsZero(m.OR) { // not required
		return nil
	}

	for i := 0; i < len(m.OR); i++ {
		if swag.IsZero(m.OR[i]) { // not required
			continue
		}

		if m.OR[i] != nil {
			if err := m.OR[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudNicSnapshotWhereInput) validateFloatingIP(formats strfmt.Registry) error {
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

func (m *VirtualPrivateCloudNicSnapshotWhereInput) validateVpc(formats strfmt.Registry) error {
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

func (m *VirtualPrivateCloudNicSnapshotWhereInput) validateVpcNic(formats strfmt.Registry) error {
	if swag.IsZero(m.VpcNic) { // not required
		return nil
	}

	if m.VpcNic != nil {
		if err := m.VpcNic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpc_nic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vpc_nic")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudNicSnapshotWhereInput) validateVpcSubnet(formats strfmt.Registry) error {
	if swag.IsZero(m.VpcSubnet) { // not required
		return nil
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

// ContextValidate validate this virtual private cloud nic snapshot where input based on the context it is used
func (m *VirtualPrivateCloudNicSnapshotWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAND(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNOT(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOR(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFloatingIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVpc(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVpcNic(ctx, formats); err != nil {
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

func (m *VirtualPrivateCloudNicSnapshotWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AND); i++ {

		if m.AND[i] != nil {
			if err := m.AND[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudNicSnapshotWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NOT); i++ {

		if m.NOT[i] != nil {
			if err := m.NOT[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudNicSnapshotWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OR); i++ {

		if m.OR[i] != nil {
			if err := m.OR[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudNicSnapshotWhereInput) contextValidateFloatingIP(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VirtualPrivateCloudNicSnapshotWhereInput) contextValidateVpc(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VirtualPrivateCloudNicSnapshotWhereInput) contextValidateVpcNic(ctx context.Context, formats strfmt.Registry) error {

	if m.VpcNic != nil {
		if err := m.VpcNic.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpc_nic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vpc_nic")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudNicSnapshotWhereInput) contextValidateVpcSubnet(ctx context.Context, formats strfmt.Registry) error {

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
func (m *VirtualPrivateCloudNicSnapshotWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualPrivateCloudNicSnapshotWhereInput) UnmarshalBinary(b []byte) error {
	var res VirtualPrivateCloudNicSnapshotWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
