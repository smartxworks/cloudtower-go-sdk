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

// VMNicWhereInput Vm nic where input
//
// swagger:model VmNicWhereInput
type VMNicWhereInput struct {

	// a n d
	AND []*VMNicWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*VMNicWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*VMNicWhereInput `json:"OR,omitempty"`

	// enabled
	Enabled *bool `json:"enabled,omitempty"`

	// enabled not
	EnabledNot *bool `json:"enabled_not,omitempty"`

	// gateway
	Gateway *string `json:"gateway,omitempty"`

	// gateway contains
	GatewayContains *string `json:"gateway_contains,omitempty"`

	// gateway ends with
	GatewayEndsWith *string `json:"gateway_ends_with,omitempty"`

	// gateway gt
	GatewayGt *string `json:"gateway_gt,omitempty"`

	// gateway gte
	GatewayGte *string `json:"gateway_gte,omitempty"`

	// gateway in
	GatewayIn []string `json:"gateway_in,omitempty"`

	// gateway lt
	GatewayLt *string `json:"gateway_lt,omitempty"`

	// gateway lte
	GatewayLte *string `json:"gateway_lte,omitempty"`

	// gateway not
	GatewayNot *string `json:"gateway_not,omitempty"`

	// gateway not contains
	GatewayNotContains *string `json:"gateway_not_contains,omitempty"`

	// gateway not ends with
	GatewayNotEndsWith *string `json:"gateway_not_ends_with,omitempty"`

	// gateway not in
	GatewayNotIn []string `json:"gateway_not_in,omitempty"`

	// gateway not starts with
	GatewayNotStartsWith *string `json:"gateway_not_starts_with,omitempty"`

	// gateway starts with
	GatewayStartsWith *string `json:"gateway_starts_with,omitempty"`

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

	// interface id
	InterfaceID *string `json:"interface_id,omitempty"`

	// interface id contains
	InterfaceIDContains *string `json:"interface_id_contains,omitempty"`

	// interface id ends with
	InterfaceIDEndsWith *string `json:"interface_id_ends_with,omitempty"`

	// interface id gt
	InterfaceIDGt *string `json:"interface_id_gt,omitempty"`

	// interface id gte
	InterfaceIDGte *string `json:"interface_id_gte,omitempty"`

	// interface id in
	InterfaceIDIn []string `json:"interface_id_in,omitempty"`

	// interface id lt
	InterfaceIDLt *string `json:"interface_id_lt,omitempty"`

	// interface id lte
	InterfaceIDLte *string `json:"interface_id_lte,omitempty"`

	// interface id not
	InterfaceIDNot *string `json:"interface_id_not,omitempty"`

	// interface id not contains
	InterfaceIDNotContains *string `json:"interface_id_not_contains,omitempty"`

	// interface id not ends with
	InterfaceIDNotEndsWith *string `json:"interface_id_not_ends_with,omitempty"`

	// interface id not in
	InterfaceIDNotIn []string `json:"interface_id_not_in,omitempty"`

	// interface id not starts with
	InterfaceIDNotStartsWith *string `json:"interface_id_not_starts_with,omitempty"`

	// interface id starts with
	InterfaceIDStartsWith *string `json:"interface_id_starts_with,omitempty"`

	// ip address
	IPAddress *string `json:"ip_address,omitempty"`

	// ip address contains
	IPAddressContains *string `json:"ip_address_contains,omitempty"`

	// ip address ends with
	IPAddressEndsWith *string `json:"ip_address_ends_with,omitempty"`

	// ip address gt
	IPAddressGt *string `json:"ip_address_gt,omitempty"`

	// ip address gte
	IPAddressGte *string `json:"ip_address_gte,omitempty"`

	// ip address in
	IPAddressIn []string `json:"ip_address_in,omitempty"`

	// ip address lt
	IPAddressLt *string `json:"ip_address_lt,omitempty"`

	// ip address lte
	IPAddressLte *string `json:"ip_address_lte,omitempty"`

	// ip address not
	IPAddressNot *string `json:"ip_address_not,omitempty"`

	// ip address not contains
	IPAddressNotContains *string `json:"ip_address_not_contains,omitempty"`

	// ip address not ends with
	IPAddressNotEndsWith *string `json:"ip_address_not_ends_with,omitempty"`

	// ip address not in
	IPAddressNotIn []string `json:"ip_address_not_in,omitempty"`

	// ip address not starts with
	IPAddressNotStartsWith *string `json:"ip_address_not_starts_with,omitempty"`

	// ip address starts with
	IPAddressStartsWith *string `json:"ip_address_starts_with,omitempty"`

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

	// mirror
	Mirror *bool `json:"mirror,omitempty"`

	// mirror not
	MirrorNot *bool `json:"mirror_not,omitempty"`

	// model
	Model *VMNicModel `json:"model,omitempty"`

	// model in
	ModelIn []VMNicModel `json:"model_in,omitempty"`

	// model not
	ModelNot *VMNicModel `json:"model_not,omitempty"`

	// model not in
	ModelNotIn []VMNicModel `json:"model_not_in,omitempty"`

	// nic
	Nic *NicWhereInput `json:"nic,omitempty"`

	// order
	Order *int32 `json:"order,omitempty"`

	// order gt
	OrderGt *int32 `json:"order_gt,omitempty"`

	// order gte
	OrderGte *int32 `json:"order_gte,omitempty"`

	// order in
	OrderIn []int32 `json:"order_in,omitempty"`

	// order lt
	OrderLt *int32 `json:"order_lt,omitempty"`

	// order lte
	OrderLte *int32 `json:"order_lte,omitempty"`

	// order not
	OrderNot *int32 `json:"order_not,omitempty"`

	// order not in
	OrderNotIn []int32 `json:"order_not_in,omitempty"`

	// subnet mask
	SubnetMask *string `json:"subnet_mask,omitempty"`

	// subnet mask contains
	SubnetMaskContains *string `json:"subnet_mask_contains,omitempty"`

	// subnet mask ends with
	SubnetMaskEndsWith *string `json:"subnet_mask_ends_with,omitempty"`

	// subnet mask gt
	SubnetMaskGt *string `json:"subnet_mask_gt,omitempty"`

	// subnet mask gte
	SubnetMaskGte *string `json:"subnet_mask_gte,omitempty"`

	// subnet mask in
	SubnetMaskIn []string `json:"subnet_mask_in,omitempty"`

	// subnet mask lt
	SubnetMaskLt *string `json:"subnet_mask_lt,omitempty"`

	// subnet mask lte
	SubnetMaskLte *string `json:"subnet_mask_lte,omitempty"`

	// subnet mask not
	SubnetMaskNot *string `json:"subnet_mask_not,omitempty"`

	// subnet mask not contains
	SubnetMaskNotContains *string `json:"subnet_mask_not_contains,omitempty"`

	// subnet mask not ends with
	SubnetMaskNotEndsWith *string `json:"subnet_mask_not_ends_with,omitempty"`

	// subnet mask not in
	SubnetMaskNotIn []string `json:"subnet_mask_not_in,omitempty"`

	// subnet mask not starts with
	SubnetMaskNotStartsWith *string `json:"subnet_mask_not_starts_with,omitempty"`

	// subnet mask starts with
	SubnetMaskStartsWith *string `json:"subnet_mask_starts_with,omitempty"`

	// vlan
	Vlan *VlanWhereInput `json:"vlan,omitempty"`

	// vm
	VM *VMWhereInput `json:"vm,omitempty"`
}

// Validate validates this Vm nic where input
func (m *VMNicWhereInput) Validate(formats strfmt.Registry) error {
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

	if err := m.validateModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModelIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModelNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModelNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVM(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMNicWhereInput) validateAND(formats strfmt.Registry) error {
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

func (m *VMNicWhereInput) validateNOT(formats strfmt.Registry) error {
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

func (m *VMNicWhereInput) validateOR(formats strfmt.Registry) error {
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

func (m *VMNicWhereInput) validateModel(formats strfmt.Registry) error {
	if swag.IsZero(m.Model) { // not required
		return nil
	}

	if m.Model != nil {
		if err := m.Model.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("model")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("model")
			}
			return err
		}
	}

	return nil
}

func (m *VMNicWhereInput) validateModelIn(formats strfmt.Registry) error {
	if swag.IsZero(m.ModelIn) { // not required
		return nil
	}

	for i := 0; i < len(m.ModelIn); i++ {

		if err := m.ModelIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("model_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("model_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VMNicWhereInput) validateModelNot(formats strfmt.Registry) error {
	if swag.IsZero(m.ModelNot) { // not required
		return nil
	}

	if m.ModelNot != nil {
		if err := m.ModelNot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("model_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("model_not")
			}
			return err
		}
	}

	return nil
}

func (m *VMNicWhereInput) validateModelNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.ModelNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.ModelNotIn); i++ {

		if err := m.ModelNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("model_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("model_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VMNicWhereInput) validateNic(formats strfmt.Registry) error {
	if swag.IsZero(m.Nic) { // not required
		return nil
	}

	if m.Nic != nil {
		if err := m.Nic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nic")
			}
			return err
		}
	}

	return nil
}

func (m *VMNicWhereInput) validateVlan(formats strfmt.Registry) error {
	if swag.IsZero(m.Vlan) { // not required
		return nil
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

func (m *VMNicWhereInput) validateVM(formats strfmt.Registry) error {
	if swag.IsZero(m.VM) { // not required
		return nil
	}

	if m.VM != nil {
		if err := m.VM.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this Vm nic where input based on the context it is used
func (m *VMNicWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateModel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModelIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModelNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModelNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNic(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVlan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVM(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMNicWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMNicWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMNicWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMNicWhereInput) contextValidateModel(ctx context.Context, formats strfmt.Registry) error {

	if m.Model != nil {
		if err := m.Model.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("model")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("model")
			}
			return err
		}
	}

	return nil
}

func (m *VMNicWhereInput) contextValidateModelIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ModelIn); i++ {

		if err := m.ModelIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("model_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("model_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VMNicWhereInput) contextValidateModelNot(ctx context.Context, formats strfmt.Registry) error {

	if m.ModelNot != nil {
		if err := m.ModelNot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("model_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("model_not")
			}
			return err
		}
	}

	return nil
}

func (m *VMNicWhereInput) contextValidateModelNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ModelNotIn); i++ {

		if err := m.ModelNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("model_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("model_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VMNicWhereInput) contextValidateNic(ctx context.Context, formats strfmt.Registry) error {

	if m.Nic != nil {
		if err := m.Nic.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nic")
			}
			return err
		}
	}

	return nil
}

func (m *VMNicWhereInput) contextValidateVlan(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMNicWhereInput) contextValidateVM(ctx context.Context, formats strfmt.Registry) error {

	if m.VM != nil {
		if err := m.VM.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMNicWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMNicWhereInput) UnmarshalBinary(b []byte) error {
	var res VMNicWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
