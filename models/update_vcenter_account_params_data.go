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

// UpdateVcenterAccountParamsData update vcenter account params data
//
// swagger:model UpdateVcenterAccountParamsData
type UpdateVcenterAccountParamsData struct {

	// ip
	// Required: true
	IP *string `json:"ip"`

	// password
	// Required: true
	Password *string `json:"password"`

	// port
	// Required: true
	Port *int32 `json:"port"`

	// username
	// Required: true
	Username *string `json:"username"`

	MarshalOpts *UpdateVcenterAccountParamsDataMarshalOpts `json:"-"`
}

type UpdateVcenterAccountParamsDataMarshalOpts struct {
	IP_Explicit_Null_When_Empty bool

	Password_Explicit_Null_When_Empty bool

	Port_Explicit_Null_When_Empty bool

	Username_Explicit_Null_When_Empty bool
}

func (m UpdateVcenterAccountParamsData) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field ip
	if m.IP != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ip\":")
		bytes, err := swag.WriteJSON(m.IP)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IP_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ip\":null")
		first = false
	}

	// handle nullable field password
	if m.Password != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"password\":")
		bytes, err := swag.WriteJSON(m.Password)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Password_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"password\":null")
		first = false
	}

	// handle nullable field port
	if m.Port != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"port\":")
		bytes, err := swag.WriteJSON(m.Port)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Port_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"port\":null")
		first = false
	}

	// handle nullable field username
	if m.Username != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"username\":")
		bytes, err := swag.WriteJSON(m.Username)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Username_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"username\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this update vcenter account params data
func (m *UpdateVcenterAccountParamsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateVcenterAccountParamsData) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("ip", "body", m.IP); err != nil {
		return err
	}

	return nil
}

func (m *UpdateVcenterAccountParamsData) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *UpdateVcenterAccountParamsData) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

func (m *UpdateVcenterAccountParamsData) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update vcenter account params data based on context it is used
func (m *UpdateVcenterAccountParamsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateVcenterAccountParamsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateVcenterAccountParamsData) UnmarshalBinary(b []byte) error {
	var res UpdateVcenterAccountParamsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
