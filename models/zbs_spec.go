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

// ZbsSpec zbs spec
//
// swagger:model ZbsSpec
type ZbsSpec string

func NewZbsSpec(value ZbsSpec) *ZbsSpec {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ZbsSpec.
func (m ZbsSpec) Pointer() *ZbsSpec {
	return &m
}

const (

	// ZbsSpecNormal captures enum value "normal"
	ZbsSpecNormal ZbsSpec = "normal"

	// ZbsSpecLarge captures enum value "large"
	ZbsSpecLarge ZbsSpec = "large"
)

// for schema
var zbsSpecEnum []interface{}

func init() {
	var res []ZbsSpec
	if err := json.Unmarshal([]byte(`["normal","large"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		zbsSpecEnum = append(zbsSpecEnum, v)
	}
}

func (m ZbsSpec) validateZbsSpecEnum(path, location string, value ZbsSpec) error {
	if err := validate.EnumCase(path, location, value, zbsSpecEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this zbs spec
func (m ZbsSpec) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateZbsSpecEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this zbs spec based on context it is used
func (m ZbsSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
