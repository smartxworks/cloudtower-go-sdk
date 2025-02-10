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

// UserAuditLogMessage user audit log message
//
// swagger:model UserAuditLogMessage
type UserAuditLogMessage struct {

	// en u s
	// Required: true
	EnUS *string `json:"en-US"`

	// zh c n
	// Required: true
	ZhCN *string `json:"zh-CN"`

	MarshalOpts *UserAuditLogMessageMarshalOpts `json:"-"`
}

type UserAuditLogMessageMarshalOpts struct {
	EnUS_Explicit_Null_When_Empty bool

	ZhCN_Explicit_Null_When_Empty bool
}

func (m UserAuditLogMessage) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field en-US
	if m.EnUS != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"en-US\":")
		bytes, err := swag.WriteJSON(m.EnUS)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.EnUS_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"en-US\":null")
		first = false
	}

	// handle nullable field zh-CN
	if m.ZhCN != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"zh-CN\":")
		bytes, err := swag.WriteJSON(m.ZhCN)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ZhCN_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"zh-CN\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this user audit log message
func (m *UserAuditLogMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnUS(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZhCN(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserAuditLogMessage) validateEnUS(formats strfmt.Registry) error {

	if err := validate.Required("en-US", "body", m.EnUS); err != nil {
		return err
	}

	return nil
}

func (m *UserAuditLogMessage) validateZhCN(formats strfmt.Registry) error {

	if err := validate.Required("zh-CN", "body", m.ZhCN); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this user audit log message based on context it is used
func (m *UserAuditLogMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserAuditLogMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserAuditLogMessage) UnmarshalBinary(b []byte) error {
	var res UserAuditLogMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
