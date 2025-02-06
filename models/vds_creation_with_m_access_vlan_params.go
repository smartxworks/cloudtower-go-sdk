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

// VdsCreationWithMAccessVlanParams vds creation with m access vlan params
//
// swagger:model VdsCreationWithMAccessVlanParams
type VdsCreationWithMAccessVlanParams struct {
	VdsCreationParams

	// vlan
	// Required: true
	Vlan *VdsCreationWithMAccessVlanParamsAO1Vlan `json:"vlan"`

	MarshalOpts *VdsCreationWithMAccessVlanParamsMarshalOpts `json:"-"`
}

type VdsCreationWithMAccessVlanParamsMarshalOpts struct {
	Vlan_Explicit_Null_When_Empty bool
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VdsCreationWithMAccessVlanParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 VdsCreationParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.VdsCreationParams = aO0

	// AO1
	var dataAO1 struct {
		Vlan *VdsCreationWithMAccessVlanParamsAO1Vlan `json:"vlan"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Vlan = dataAO1.Vlan

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VdsCreationWithMAccessVlanParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.VdsCreationParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var b bytes.Buffer
	b.WriteString("{")
	first := true

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
	b.WriteString("}")
	_parts = append(_parts, b.Bytes())
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this vds creation with m access vlan params
func (m *VdsCreationWithMAccessVlanParams) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with VdsCreationParams
	if err := m.VdsCreationParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VdsCreationWithMAccessVlanParams) validateVlan(formats strfmt.Registry) error {

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

// ContextValidate validate this vds creation with m access vlan params based on the context it is used
func (m *VdsCreationWithMAccessVlanParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with VdsCreationParams
	if err := m.VdsCreationParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVlan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VdsCreationWithMAccessVlanParams) contextValidateVlan(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *VdsCreationWithMAccessVlanParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VdsCreationWithMAccessVlanParams) UnmarshalBinary(b []byte) error {
	var res VdsCreationWithMAccessVlanParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VdsCreationWithMAccessVlanParamsAO1Vlan vds creation with m access vlan params a o1 vlan
//
// swagger:model VdsCreationWithMAccessVlanParamsAO1Vlan
type VdsCreationWithMAccessVlanParamsAO1Vlan struct {

	// extra ip
	// Required: true
	ExtraIP []*VdsCreationWithMAccessVlanParamsAO1VlanExtraIPItems0 `json:"extra_ip"`

	// gateway ip
	GatewayIP *string `json:"gateway_ip,omitempty"`

	// gateway subnetmask
	GatewaySubnetmask *string `json:"gateway_subnetmask,omitempty"`

	// subnetmask
	// Required: true
	Subnetmask *string `json:"subnetmask"`

	// vlan id
	// Required: true
	VlanID *int32 `json:"vlan_id"`

	MarshalOpts *VdsCreationWithMAccessVlanParamsAO1VlanMarshalOpts `json:"-"`
}

type VdsCreationWithMAccessVlanParamsAO1VlanMarshalOpts struct {
	ExtraIP_Explicit_Null_When_Empty bool

	GatewayIP_Explicit_Null_When_Empty bool

	GatewaySubnetmask_Explicit_Null_When_Empty bool

	Subnetmask_Explicit_Null_When_Empty bool

	VlanID_Explicit_Null_When_Empty bool
}

func (m VdsCreationWithMAccessVlanParamsAO1Vlan) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle non nullable field extra_ip without omitempty
	if !first {
		b.WriteString(",")
	}
	b.WriteString("\"extra_ip\":")
	{
		bytes, err := swag.WriteJSON(m.ExtraIP)
		if err != nil {
			return nil, err
		}
	}
	b.Write(bytes)
	first = false

	// handle nullable field gateway_ip
	if m.GatewayIP != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"gateway_ip\":")
		bytes, err := swag.WriteJSON(m.GatewayIP)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.GatewayIP_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"gateway_ip\":null")
		first = false
	}

	// handle nullable field gateway_subnetmask
	if m.GatewaySubnetmask != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"gateway_subnetmask\":")
		bytes, err := swag.WriteJSON(m.GatewaySubnetmask)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.GatewaySubnetmask_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"gateway_subnetmask\":null")
		first = false
	}

	// handle nullable field subnetmask
	if m.Subnetmask != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"subnetmask\":")
		bytes, err := swag.WriteJSON(m.Subnetmask)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Subnetmask_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"subnetmask\":null")
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

// Validate validates this vds creation with m access vlan params a o1 vlan
func (m *VdsCreationWithMAccessVlanParamsAO1Vlan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExtraIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnetmask(formats); err != nil {
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

func (m *VdsCreationWithMAccessVlanParamsAO1Vlan) validateExtraIP(formats strfmt.Registry) error {

	if err := validate.Required("vlan"+"."+"extra_ip", "body", m.ExtraIP); err != nil {
		return err
	}

	for i := 0; i < len(m.ExtraIP); i++ {
		if swag.IsZero(m.ExtraIP[i]) { // not required
			continue
		}

		if m.ExtraIP[i] != nil {
			if err := m.ExtraIP[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vlan" + "." + "extra_ip" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vlan" + "." + "extra_ip" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VdsCreationWithMAccessVlanParamsAO1Vlan) validateSubnetmask(formats strfmt.Registry) error {

	if err := validate.Required("vlan"+"."+"subnetmask", "body", m.Subnetmask); err != nil {
		return err
	}

	return nil
}

func (m *VdsCreationWithMAccessVlanParamsAO1Vlan) validateVlanID(formats strfmt.Registry) error {

	if err := validate.Required("vlan"+"."+"vlan_id", "body", m.VlanID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this vds creation with m access vlan params a o1 vlan based on the context it is used
func (m *VdsCreationWithMAccessVlanParamsAO1Vlan) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExtraIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VdsCreationWithMAccessVlanParamsAO1Vlan) contextValidateExtraIP(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExtraIP); i++ {

		if m.ExtraIP[i] != nil {
			if err := m.ExtraIP[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vlan" + "." + "extra_ip" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vlan" + "." + "extra_ip" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VdsCreationWithMAccessVlanParamsAO1Vlan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VdsCreationWithMAccessVlanParamsAO1Vlan) UnmarshalBinary(b []byte) error {
	var res VdsCreationWithMAccessVlanParamsAO1Vlan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VdsCreationWithMAccessVlanParamsAO1VlanExtraIPItems0 vds creation with m access vlan params a o1 vlan extra IP items0
//
// swagger:model VdsCreationWithMAccessVlanParamsAO1VlanExtraIPItems0
type VdsCreationWithMAccessVlanParamsAO1VlanExtraIPItems0 struct {

	// host id
	// Required: true
	HostID *string `json:"host_id"`

	// management ip
	// Required: true
	ManagementIP *string `json:"management_ip"`

	MarshalOpts *VdsCreationWithMAccessVlanParamsAO1VlanExtraIPItems0MarshalOpts `json:"-"`
}

type VdsCreationWithMAccessVlanParamsAO1VlanExtraIPItems0MarshalOpts struct {
	HostID_Explicit_Null_When_Empty bool

	ManagementIP_Explicit_Null_When_Empty bool
}

func (m VdsCreationWithMAccessVlanParamsAO1VlanExtraIPItems0) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field host_id
	if m.HostID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"host_id\":")
		bytes, err := swag.WriteJSON(m.HostID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.HostID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"host_id\":null")
		first = false
	}

	// handle nullable field management_ip
	if m.ManagementIP != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"management_ip\":")
		bytes, err := swag.WriteJSON(m.ManagementIP)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ManagementIP_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"management_ip\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this vds creation with m access vlan params a o1 vlan extra IP items0
func (m *VdsCreationWithMAccessVlanParamsAO1VlanExtraIPItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManagementIP(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VdsCreationWithMAccessVlanParamsAO1VlanExtraIPItems0) validateHostID(formats strfmt.Registry) error {

	if err := validate.Required("host_id", "body", m.HostID); err != nil {
		return err
	}

	return nil
}

func (m *VdsCreationWithMAccessVlanParamsAO1VlanExtraIPItems0) validateManagementIP(formats strfmt.Registry) error {

	if err := validate.Required("management_ip", "body", m.ManagementIP); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this vds creation with m access vlan params a o1 vlan extra IP items0 based on context it is used
func (m *VdsCreationWithMAccessVlanParamsAO1VlanExtraIPItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VdsCreationWithMAccessVlanParamsAO1VlanExtraIPItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VdsCreationWithMAccessVlanParamsAO1VlanExtraIPItems0) UnmarshalBinary(b []byte) error {
	var res VdsCreationWithMAccessVlanParamsAO1VlanExtraIPItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
