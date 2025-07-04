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

// UserRolePreset user role preset
//
// swagger:model UserRolePreset
type UserRolePreset string

func NewUserRolePreset(value UserRolePreset) *UserRolePreset {
	return &value
}

// Pointer returns a pointer to a freshly-allocated UserRolePreset.
func (m UserRolePreset) Pointer() *UserRolePreset {
	return &m
}

const (

	// UserRolePresetADMIN captures enum value "ADMIN"
	UserRolePresetADMIN UserRolePreset = "ADMIN"

	// UserRolePresetAUDITOR captures enum value "AUDITOR"
	UserRolePresetAUDITOR UserRolePreset = "AUDITOR"

	// UserRolePresetREADONLY captures enum value "READ_ONLY"
	UserRolePresetREADONLY UserRolePreset = "READ_ONLY"

	// UserRolePresetROOT captures enum value "ROOT"
	UserRolePresetROOT UserRolePreset = "ROOT"

	// UserRolePresetUSERADMIN captures enum value "USER_ADMIN"
	UserRolePresetUSERADMIN UserRolePreset = "USER_ADMIN"

	// UserRolePresetVMUSER captures enum value "VM_USER"
	UserRolePresetVMUSER UserRolePreset = "VM_USER"

	// UserRolePresetWORKLOADCLUSTERUSER captures enum value "WORKLOAD_CLUSTER_USER"
	UserRolePresetWORKLOADCLUSTERUSER UserRolePreset = "WORKLOAD_CLUSTER_USER"
)

// for schema
var userRolePresetEnum []interface{}

func init() {
	var res []UserRolePreset
	if err := json.Unmarshal([]byte(`["ADMIN","AUDITOR","READ_ONLY","ROOT","USER_ADMIN","VM_USER","WORKLOAD_CLUSTER_USER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userRolePresetEnum = append(userRolePresetEnum, v)
	}
}

func (m UserRolePreset) validateUserRolePresetEnum(path, location string, value UserRolePreset) error {
	if err := validate.EnumCase(path, location, value, userRolePresetEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this user role preset
func (m UserRolePreset) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateUserRolePresetEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this user role preset based on context it is used
func (m UserRolePreset) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
