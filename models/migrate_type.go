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

// MigrateType migrate type
//
// swagger:model MigrateType
type MigrateType string

func NewMigrateType(value MigrateType) *MigrateType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated MigrateType.
func (m MigrateType) Pointer() *MigrateType {
	return &m
}

const (

	// MigrateTypeCOLDMIGRATE captures enum value "COLD_MIGRATE"
	MigrateTypeCOLDMIGRATE MigrateType = "COLD_MIGRATE"

	// MigrateTypeCUTOVERMIGRATE captures enum value "CUTOVER_MIGRATE"
	MigrateTypeCUTOVERMIGRATE MigrateType = "CUTOVER_MIGRATE"

	// MigrateTypeLIVEMIGRATE captures enum value "LIVE_MIGRATE"
	MigrateTypeLIVEMIGRATE MigrateType = "LIVE_MIGRATE"
)

// for schema
var migrateTypeEnum []interface{}

func init() {
	var res []MigrateType
	if err := json.Unmarshal([]byte(`["COLD_MIGRATE","CUTOVER_MIGRATE","LIVE_MIGRATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		migrateTypeEnum = append(migrateTypeEnum, v)
	}
}

func (m MigrateType) validateMigrateTypeEnum(path, location string, value MigrateType) error {
	if err := validate.EnumCase(path, location, value, migrateTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this migrate type
func (m MigrateType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateMigrateTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this migrate type based on context it is used
func (m MigrateType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
