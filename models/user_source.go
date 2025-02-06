// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// UserSource user source
//
// swagger:model UserSource
type UserSource string

func NewUserSource(value UserSource) *UserSource {
	return &value
}

// Pointer returns a pointer to a freshly-allocated UserSource.
func (m UserSource) Pointer() *UserSource {
	return &m
}

const (

	// UserSourceAUTHN captures enum value "AUTHN"
	UserSourceAUTHN UserSource = "AUTHN"

	// UserSourceLDAP captures enum value "LDAP"
	UserSourceLDAP UserSource = "LDAP"

	// UserSourceLOCAL captures enum value "LOCAL"
	UserSourceLOCAL UserSource = "LOCAL"

	// UserSourceSSO captures enum value "SSO"
	UserSourceSSO UserSource = "SSO"
)

// for schema
var userSourceEnum []interface{}

func init() {
	var res []UserSource
	if err := json.Unmarshal([]byte(`["AUTHN","LDAP","LOCAL","SSO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userSourceEnum = append(userSourceEnum, v)
	}
}

func (m UserSource) validateUserSourceEnum(path, location string, value UserSource) error {
	if err := validate.EnumCase(path, location, value, userSourceEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this user source
func (m UserSource) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateUserSourceEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this user source based on context it is used
func (m UserSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
