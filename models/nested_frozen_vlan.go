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

// NestedFrozenVlan nested frozen vlan
//
// swagger:model NestedFrozenVlan
type NestedFrozenVlan struct {

	// name
	// Required: true
	Name *string `json:"name"`

	// network ids
	NetworkIds []string `json:"network_ids,omitempty"`

	// vds ovs
	// Required: true
	VdsOvs *string `json:"vds_ovs"`

	// vlan id
	// Required: true
	VlanID *int32 `json:"vlan_id"`

	// vlan local id
	// Required: true
	VlanLocalID *string `json:"vlan_local_id"`

	MarshalOpts *NestedFrozenVlanMarshalOpts `json:"-"`
}

type NestedFrozenVlanMarshalOpts struct {
	Name_Explicit_Null_When_Empty bool

	VdsOvs_Explicit_Null_When_Empty bool

	VlanID_Explicit_Null_When_Empty bool

	VlanLocalID_Explicit_Null_When_Empty bool
}

func (m NestedFrozenVlan) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

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

	// handle non nullable field network_ids with omitempty
	if swag.IsZero(m.NetworkIds) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"network_ids\":")
		bytes, err := swag.WriteJSON(m.NetworkIds)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field vds_ovs
	if m.VdsOvs != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vds_ovs\":")
		bytes, err := swag.WriteJSON(m.VdsOvs)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VdsOvs_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vds_ovs\":null")
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

	// handle nullable field vlan_local_id
	if m.VlanLocalID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vlan_local_id\":")
		bytes, err := swag.WriteJSON(m.VlanLocalID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VlanLocalID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vlan_local_id\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this nested frozen vlan
func (m *NestedFrozenVlan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVdsOvs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlanID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlanLocalID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedFrozenVlan) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *NestedFrozenVlan) validateVdsOvs(formats strfmt.Registry) error {

	if err := validate.Required("vds_ovs", "body", m.VdsOvs); err != nil {
		return err
	}

	return nil
}

func (m *NestedFrozenVlan) validateVlanID(formats strfmt.Registry) error {

	if err := validate.Required("vlan_id", "body", m.VlanID); err != nil {
		return err
	}

	return nil
}

func (m *NestedFrozenVlan) validateVlanLocalID(formats strfmt.Registry) error {

	if err := validate.Required("vlan_local_id", "body", m.VlanLocalID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this nested frozen vlan based on context it is used
func (m *NestedFrozenVlan) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NestedFrozenVlan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedFrozenVlan) UnmarshalBinary(b []byte) error {
	var res NestedFrozenVlan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
