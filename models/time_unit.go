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

// TimeUnit time unit
//
// swagger:model TimeUnit
type TimeUnit string

func NewTimeUnit(value TimeUnit) *TimeUnit {
	return &value
}

// Pointer returns a pointer to a freshly-allocated TimeUnit.
func (m TimeUnit) Pointer() *TimeUnit {
	return &m
}

const (

	// TimeUnitDAY captures enum value "DAY"
	TimeUnitDAY TimeUnit = "DAY"

	// TimeUnitHOUR captures enum value "HOUR"
	TimeUnitHOUR TimeUnit = "HOUR"

	// TimeUnitMINUTE captures enum value "MINUTE"
	TimeUnitMINUTE TimeUnit = "MINUTE"

	// TimeUnitMONTH captures enum value "MONTH"
	TimeUnitMONTH TimeUnit = "MONTH"
)

// for schema
var timeUnitEnum []interface{}

func init() {
	var res []TimeUnit
	if err := json.Unmarshal([]byte(`["DAY","HOUR","MINUTE","MONTH"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		timeUnitEnum = append(timeUnitEnum, v)
	}
}

func (m TimeUnit) validateTimeUnitEnum(path, location string, value TimeUnit) error {
	if err := validate.EnumCase(path, location, value, timeUnitEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this time unit
func (m TimeUnit) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTimeUnitEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this time unit based on context it is used
func (m TimeUnit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
