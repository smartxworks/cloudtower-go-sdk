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

// VirtualPrivateCloudSecurityPolicyMode virtual private cloud security policy mode
//
// swagger:model VirtualPrivateCloudSecurityPolicyMode
type VirtualPrivateCloudSecurityPolicyMode string

func NewVirtualPrivateCloudSecurityPolicyMode(value VirtualPrivateCloudSecurityPolicyMode) *VirtualPrivateCloudSecurityPolicyMode {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VirtualPrivateCloudSecurityPolicyMode.
func (m VirtualPrivateCloudSecurityPolicyMode) Pointer() *VirtualPrivateCloudSecurityPolicyMode {
	return &m
}

const (

	// VirtualPrivateCloudSecurityPolicyModeMONITOR captures enum value "MONITOR"
	VirtualPrivateCloudSecurityPolicyModeMONITOR VirtualPrivateCloudSecurityPolicyMode = "MONITOR"

	// VirtualPrivateCloudSecurityPolicyModeWORK captures enum value "WORK"
	VirtualPrivateCloudSecurityPolicyModeWORK VirtualPrivateCloudSecurityPolicyMode = "WORK"
)

// for schema
var virtualPrivateCloudSecurityPolicyModeEnum []interface{}

func init() {
	var res []VirtualPrivateCloudSecurityPolicyMode
	if err := json.Unmarshal([]byte(`["MONITOR","WORK"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		virtualPrivateCloudSecurityPolicyModeEnum = append(virtualPrivateCloudSecurityPolicyModeEnum, v)
	}
}

func (m VirtualPrivateCloudSecurityPolicyMode) validateVirtualPrivateCloudSecurityPolicyModeEnum(path, location string, value VirtualPrivateCloudSecurityPolicyMode) error {
	if err := validate.EnumCase(path, location, value, virtualPrivateCloudSecurityPolicyModeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this virtual private cloud security policy mode
func (m VirtualPrivateCloudSecurityPolicyMode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateVirtualPrivateCloudSecurityPolicyModeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this virtual private cloud security policy mode based on context it is used
func (m VirtualPrivateCloudSecurityPolicyMode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}