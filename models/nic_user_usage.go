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

// NicUserUsage nic user usage
//
// swagger:model NicUserUsage
type NicUserUsage string

func NewNicUserUsage(value NicUserUsage) *NicUserUsage {
	return &value
}

// Pointer returns a pointer to a freshly-allocated NicUserUsage.
func (m NicUserUsage) Pointer() *NicUserUsage {
	return &m
}

const (

	// NicUserUsageDEFAULT captures enum value "DEFAULT"
	NicUserUsageDEFAULT NicUserUsage = "DEFAULT"

	// NicUserUsagePASSTHROUGH captures enum value "PASS_THROUGH"
	NicUserUsagePASSTHROUGH NicUserUsage = "PASS_THROUGH"

	// NicUserUsageSRIOV captures enum value "SRIOV"
	NicUserUsageSRIOV NicUserUsage = "SRIOV"
)

// for schema
var nicUserUsageEnum []interface{}

func init() {
	var res []NicUserUsage
	if err := json.Unmarshal([]byte(`["DEFAULT","PASS_THROUGH","SRIOV"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nicUserUsageEnum = append(nicUserUsageEnum, v)
	}
}

func (m NicUserUsage) validateNicUserUsageEnum(path, location string, value NicUserUsage) error {
	if err := validate.EnumCase(path, location, value, nicUserUsageEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this nic user usage
func (m NicUserUsage) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateNicUserUsageEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this nic user usage based on context it is used
func (m NicUserUsage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
