// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NestedStoragePolicyConfig nested storage policy config
//
// swagger:model NestedStoragePolicyConfig
type NestedStoragePolicyConfig struct {

	// replica num
	ReplicaNum *int32 `json:"replica_num,omitempty"`

	// thin provision
	ThinProvision *bool `json:"thin_provision,omitempty"`

	MarshalOpts *NestedStoragePolicyConfigMarshalOpts `json:"-"`
}

type NestedStoragePolicyConfigMarshalOpts struct {
	ReplicaNum_Explicit_Null_When_Empty bool

	ThinProvision_Explicit_Null_When_Empty bool
}

func (m NestedStoragePolicyConfig) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field replica_num
	if m.ReplicaNum != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"replica_num\":")
		bytes, err := swag.WriteJSON(m.ReplicaNum)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ReplicaNum_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"replica_num\":null")
		first = false
	}

	// handle nullable field thin_provision
	if m.ThinProvision != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"thin_provision\":")
		bytes, err := swag.WriteJSON(m.ThinProvision)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ThinProvision_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"thin_provision\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this nested storage policy config
func (m *NestedStoragePolicyConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this nested storage policy config based on context it is used
func (m *NestedStoragePolicyConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NestedStoragePolicyConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedStoragePolicyConfig) UnmarshalBinary(b []byte) error {
	var res NestedStoragePolicyConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
