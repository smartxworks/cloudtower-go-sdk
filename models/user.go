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

// User user
//
// swagger:model User
type User struct {

	// auth config id
	AuthConfigID *string `json:"auth_config_id,omitempty"`

	// display username
	// Required: true
	DisplayUsername *string `json:"display_username"`

	// email address
	EmailAddress *string `json:"email_address,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// internal
	// Required: true
	Internal *bool `json:"internal"`

	// ldap dn
	LdapDn *string `json:"ldap_dn,omitempty"`

	// mobile phone
	MobilePhone *string `json:"mobile_phone,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// password expired
	PasswordExpired *bool `json:"password_expired,omitempty"`

	// password recover qa
	PasswordRecoverQa *NestedPasswordRecoverQa `json:"password_recover_qa,omitempty"`

	// password updated at
	PasswordUpdatedAt *string `json:"password_updated_at,omitempty"`

	// role
	Role *UserRole `json:"role,omitempty"`

	// roles
	Roles []*NestedUserRoleNext `json:"roles,omitempty"`

	// source
	// Required: true
	Source *UserSource `json:"source"`

	// username
	// Required: true
	Username *string `json:"username"`

	MarshalOpts *UserMarshalOpts `json:"-"`
}

type UserMarshalOpts struct {
	AuthConfigID_Explicit_Null_When_Empty bool

	DisplayUsername_Explicit_Null_When_Empty bool

	EmailAddress_Explicit_Null_When_Empty bool

	ID_Explicit_Null_When_Empty bool

	Internal_Explicit_Null_When_Empty bool

	LdapDn_Explicit_Null_When_Empty bool

	MobilePhone_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	PasswordExpired_Explicit_Null_When_Empty bool

	PasswordRecoverQa_Explicit_Null_When_Empty bool

	PasswordUpdatedAt_Explicit_Null_When_Empty bool

	Role_Explicit_Null_When_Empty bool

	Roles_Explicit_Null_When_Empty bool

	Source_Explicit_Null_When_Empty bool

	Username_Explicit_Null_When_Empty bool
}

func (m User) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field auth_config_id
	if m.AuthConfigID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"auth_config_id\":")
		bytes, err := swag.WriteJSON(m.AuthConfigID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.AuthConfigID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"auth_config_id\":null")
		first = false
	}

	// handle nullable field display_username
	if m.DisplayUsername != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"display_username\":")
		bytes, err := swag.WriteJSON(m.DisplayUsername)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.DisplayUsername_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"display_username\":null")
		first = false
	}

	// handle nullable field email_address
	if m.EmailAddress != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"email_address\":")
		bytes, err := swag.WriteJSON(m.EmailAddress)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.EmailAddress_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"email_address\":null")
		first = false
	}

	// handle nullable field id
	if m.ID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id\":")
		bytes, err := swag.WriteJSON(m.ID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id\":null")
		first = false
	}

	// handle nullable field internal
	if m.Internal != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"internal\":")
		bytes, err := swag.WriteJSON(m.Internal)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Internal_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"internal\":null")
		first = false
	}

	// handle nullable field ldap_dn
	if m.LdapDn != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ldap_dn\":")
		bytes, err := swag.WriteJSON(m.LdapDn)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.LdapDn_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ldap_dn\":null")
		first = false
	}

	// handle nullable field mobile_phone
	if m.MobilePhone != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"mobile_phone\":")
		bytes, err := swag.WriteJSON(m.MobilePhone)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.MobilePhone_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"mobile_phone\":null")
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

	// handle nullable field password_expired
	if m.PasswordExpired != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"password_expired\":")
		bytes, err := swag.WriteJSON(m.PasswordExpired)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.PasswordExpired_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"password_expired\":null")
		first = false
	}

	// handle nullable field password_recover_qa
	if m.PasswordRecoverQa != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"password_recover_qa\":")
		bytes, err := swag.WriteJSON(m.PasswordRecoverQa)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.PasswordRecoverQa_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"password_recover_qa\":null")
		first = false
	}

	// handle nullable field password_updated_at
	if m.PasswordUpdatedAt != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"password_updated_at\":")
		bytes, err := swag.WriteJSON(m.PasswordUpdatedAt)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.PasswordUpdatedAt_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"password_updated_at\":null")
		first = false
	}

	// handle nullable field role
	if m.Role != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"role\":")
		bytes, err := swag.WriteJSON(m.Role)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Role_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"role\":null")
		first = false
	}

	// handle non nullable field roles with omitempty
	if !swag.IsZero(m.Roles) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"roles\":")
		bytes, err := swag.WriteJSON(m.Roles)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field source
	if m.Source != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"source\":")
		bytes, err := swag.WriteJSON(m.Source)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Source_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"source\":null")
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

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplayUsername(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInternal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePasswordRecoverQa(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
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

func (m *User) validateDisplayUsername(formats strfmt.Registry) error {

	if err := validate.Required("display_username", "body", m.DisplayUsername); err != nil {
		return err
	}

	return nil
}

func (m *User) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *User) validateInternal(formats strfmt.Registry) error {

	if err := validate.Required("internal", "body", m.Internal); err != nil {
		return err
	}

	return nil
}

func (m *User) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *User) validatePasswordRecoverQa(formats strfmt.Registry) error {
	if swag.IsZero(m.PasswordRecoverQa) { // not required
		return nil
	}

	if m.PasswordRecoverQa != nil {
		if err := m.PasswordRecoverQa.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("password_recover_qa")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("password_recover_qa")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateRole(formats strfmt.Registry) error {
	if swag.IsZero(m.Role) { // not required
		return nil
	}

	if m.Role != nil {
		if err := m.Role.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateRoles(formats strfmt.Registry) error {
	if swag.IsZero(m.Roles) { // not required
		return nil
	}

	for i := 0; i < len(m.Roles); i++ {
		if swag.IsZero(m.Roles[i]) { // not required
			continue
		}

		if m.Roles[i] != nil {
			if err := m.Roles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("roles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *User) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this user based on the context it is used
func (m *User) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePasswordRecoverQa(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRole(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) contextValidatePasswordRecoverQa(ctx context.Context, formats strfmt.Registry) error {

	if m.PasswordRecoverQa != nil {
		if err := m.PasswordRecoverQa.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("password_recover_qa")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("password_recover_qa")
			}
			return err
		}
	}

	return nil
}

func (m *User) contextValidateRole(ctx context.Context, formats strfmt.Registry) error {

	if m.Role != nil {
		if err := m.Role.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

func (m *User) contextValidateRoles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Roles); i++ {

		if m.Roles[i] != nil {
			if err := m.Roles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("roles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *User) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	if m.Source != nil {
		if err := m.Source.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
