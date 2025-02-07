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

// TemplateCloudInit template cloud init
//
// swagger:model TemplateCloudInit
type TemplateCloudInit struct {

	// default user password
	DefaultUserPassword *string `json:"default_user_password,omitempty"`

	// hostname
	Hostname *string `json:"hostname,omitempty"`

	// nameservers
	Nameservers []string `json:"nameservers,omitempty"`

	// networks
	Networks []*CloudInitNetWork `json:"networks,omitempty"`

	// public keys
	PublicKeys []string `json:"public_keys,omitempty"`

	// user data
	UserData *string `json:"user_data,omitempty"`

	MarshalOpts *TemplateCloudInitMarshalOpts `json:"-"`
}

type TemplateCloudInitMarshalOpts struct {
	DefaultUserPassword_Explicit_Null_When_Empty bool

	Hostname_Explicit_Null_When_Empty bool

	Nameservers_Explicit_Null_When_Empty bool

	Networks_Explicit_Null_When_Empty bool

	PublicKeys_Explicit_Null_When_Empty bool

	UserData_Explicit_Null_When_Empty bool
}

func (m TemplateCloudInit) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field default_user_password
	if m.DefaultUserPassword != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"default_user_password\":")
		bytes, err := swag.WriteJSON(m.DefaultUserPassword)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.DefaultUserPassword_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"default_user_password\":null")
		first = false
	}

	// handle nullable field hostname
	if m.Hostname != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"hostname\":")
		bytes, err := swag.WriteJSON(m.Hostname)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Hostname_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"hostname\":null")
		first = false
	}

	// handle non nullable field nameservers with omitempty
	if !swag.IsZero(m.Nameservers) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"nameservers\":")
		bytes, err := swag.WriteJSON(m.Nameservers)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle non nullable field networks with omitempty
	if !swag.IsZero(m.Networks) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"networks\":")
		bytes, err := swag.WriteJSON(m.Networks)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle non nullable field public_keys with omitempty
	if !swag.IsZero(m.PublicKeys) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"public_keys\":")
		bytes, err := swag.WriteJSON(m.PublicKeys)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field user_data
	if m.UserData != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"user_data\":")
		bytes, err := swag.WriteJSON(m.UserData)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.UserData_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"user_data\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this template cloud init
func (m *TemplateCloudInit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetworks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TemplateCloudInit) validateNetworks(formats strfmt.Registry) error {
	if swag.IsZero(m.Networks) { // not required
		return nil
	}

	for i := 0; i < len(m.Networks); i++ {
		if swag.IsZero(m.Networks[i]) { // not required
			continue
		}

		if m.Networks[i] != nil {
			if err := m.Networks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this template cloud init based on the context it is used
func (m *TemplateCloudInit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNetworks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TemplateCloudInit) contextValidateNetworks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Networks); i++ {

		if m.Networks[i] != nil {
			if err := m.Networks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TemplateCloudInit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TemplateCloudInit) UnmarshalBinary(b []byte) error {
	var res TemplateCloudInit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
