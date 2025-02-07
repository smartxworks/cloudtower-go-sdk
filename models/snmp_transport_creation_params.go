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

// SnmpTransportCreationParams snmp transport creation params
//
// swagger:model SnmpTransportCreationParams
type SnmpTransportCreationParams struct {

	// auth pass phrase
	AuthPassPhrase *string `json:"auth_pass_phrase,omitempty"`

	// auth protocol
	AuthProtocol *SnmpAuthProtocol `json:"auth_protocol,omitempty"`

	// cluster id
	// Required: true
	ClusterID *string `json:"cluster_id"`

	// community
	Community *string `json:"community,omitempty"`

	// disabled
	Disabled *bool `json:"disabled,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// port
	// Required: true
	Port *int32 `json:"port"`

	// privacy pass phrase
	PrivacyPassPhrase *string `json:"privacy_pass_phrase,omitempty"`

	// privacy protocol
	PrivacyProtocol *SnmpPrivacyProtocol `json:"privacy_protocol,omitempty"`

	// protocol
	// Required: true
	Protocol *SnmpProtocol `json:"protocol"`

	// username
	Username *string `json:"username,omitempty"`

	// version
	// Required: true
	Version *SnmpVersion `json:"version"`

	MarshalOpts *SnmpTransportCreationParamsMarshalOpts `json:"-"`
}

type SnmpTransportCreationParamsMarshalOpts struct {
	AuthPassPhrase_Explicit_Null_When_Empty bool

	AuthProtocol_Explicit_Null_When_Empty bool

	ClusterID_Explicit_Null_When_Empty bool

	Community_Explicit_Null_When_Empty bool

	Disabled_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	Port_Explicit_Null_When_Empty bool

	PrivacyPassPhrase_Explicit_Null_When_Empty bool

	PrivacyProtocol_Explicit_Null_When_Empty bool

	Protocol_Explicit_Null_When_Empty bool

	Username_Explicit_Null_When_Empty bool

	Version_Explicit_Null_When_Empty bool
}

func (m SnmpTransportCreationParams) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field auth_pass_phrase
	if m.AuthPassPhrase != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"auth_pass_phrase\":")
		bytes, err := swag.WriteJSON(m.AuthPassPhrase)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.AuthPassPhrase_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"auth_pass_phrase\":null")
		first = false
	}

	// handle nullable field auth_protocol
	if m.AuthProtocol != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"auth_protocol\":")
		bytes, err := swag.WriteJSON(m.AuthProtocol)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.AuthProtocol_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"auth_protocol\":null")
		first = false
	}

	// handle nullable field cluster_id
	if m.ClusterID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster_id\":")
		bytes, err := swag.WriteJSON(m.ClusterID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ClusterID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster_id\":null")
		first = false
	}

	// handle nullable field community
	if m.Community != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"community\":")
		bytes, err := swag.WriteJSON(m.Community)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Community_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"community\":null")
		first = false
	}

	// handle nullable field disabled
	if m.Disabled != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"disabled\":")
		bytes, err := swag.WriteJSON(m.Disabled)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Disabled_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"disabled\":null")
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

	// handle nullable field privacy_pass_phrase
	if m.PrivacyPassPhrase != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"privacy_pass_phrase\":")
		bytes, err := swag.WriteJSON(m.PrivacyPassPhrase)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.PrivacyPassPhrase_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"privacy_pass_phrase\":null")
		first = false
	}

	// handle nullable field privacy_protocol
	if m.PrivacyProtocol != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"privacy_protocol\":")
		bytes, err := swag.WriteJSON(m.PrivacyProtocol)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.PrivacyProtocol_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"privacy_protocol\":null")
		first = false
	}

	// handle nullable field protocol
	if m.Protocol != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"protocol\":")
		bytes, err := swag.WriteJSON(m.Protocol)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Protocol_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"protocol\":null")
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

	// handle nullable field version
	if m.Version != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"version\":")
		bytes, err := swag.WriteJSON(m.Version)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Version_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"version\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this snmp transport creation params
func (m *SnmpTransportCreationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivacyProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnmpTransportCreationParams) validateAuthProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthProtocol) { // not required
		return nil
	}

	if m.AuthProtocol != nil {
		if err := m.AuthProtocol.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auth_protocol")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("auth_protocol")
			}
			return err
		}
	}

	return nil
}

func (m *SnmpTransportCreationParams) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("cluster_id", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *SnmpTransportCreationParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *SnmpTransportCreationParams) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

func (m *SnmpTransportCreationParams) validatePrivacyProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.PrivacyProtocol) { // not required
		return nil
	}

	if m.PrivacyProtocol != nil {
		if err := m.PrivacyProtocol.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("privacy_protocol")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("privacy_protocol")
			}
			return err
		}
	}

	return nil
}

func (m *SnmpTransportCreationParams) validateProtocol(formats strfmt.Registry) error {

	if err := validate.Required("protocol", "body", m.Protocol); err != nil {
		return err
	}

	if err := validate.Required("protocol", "body", m.Protocol); err != nil {
		return err
	}

	if m.Protocol != nil {
		if err := m.Protocol.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protocol")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("protocol")
			}
			return err
		}
	}

	return nil
}

func (m *SnmpTransportCreationParams) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	if m.Version != nil {
		if err := m.Version.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("version")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snmp transport creation params based on the context it is used
func (m *SnmpTransportCreationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthProtocol(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrivacyProtocol(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtocol(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnmpTransportCreationParams) contextValidateAuthProtocol(ctx context.Context, formats strfmt.Registry) error {

	if m.AuthProtocol != nil {
		if err := m.AuthProtocol.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auth_protocol")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("auth_protocol")
			}
			return err
		}
	}

	return nil
}

func (m *SnmpTransportCreationParams) contextValidatePrivacyProtocol(ctx context.Context, formats strfmt.Registry) error {

	if m.PrivacyProtocol != nil {
		if err := m.PrivacyProtocol.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("privacy_protocol")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("privacy_protocol")
			}
			return err
		}
	}

	return nil
}

func (m *SnmpTransportCreationParams) contextValidateProtocol(ctx context.Context, formats strfmt.Registry) error {

	if m.Protocol != nil {
		if err := m.Protocol.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protocol")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("protocol")
			}
			return err
		}
	}

	return nil
}

func (m *SnmpTransportCreationParams) contextValidateVersion(ctx context.Context, formats strfmt.Registry) error {

	if m.Version != nil {
		if err := m.Version.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("version")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnmpTransportCreationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnmpTransportCreationParams) UnmarshalBinary(b []byte) error {
	var res SnmpTransportCreationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
