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

// OperateActionEnum operate action enum
//
// swagger:model OperateActionEnum
type OperateActionEnum string

func NewOperateActionEnum(value OperateActionEnum) *OperateActionEnum {
	return &value
}

// Pointer returns a pointer to a freshly-allocated OperateActionEnum.
func (m OperateActionEnum) Pointer() *OperateActionEnum {
	return &m
}

const (

	// OperateActionEnumPoweroff captures enum value "poweroff"
	OperateActionEnumPoweroff OperateActionEnum = "poweroff"

	// OperateActionEnumReboot captures enum value "reboot"
	OperateActionEnumReboot OperateActionEnum = "reboot"
)

// for schema
var operateActionEnumEnum []interface{}

func init() {
	var res []OperateActionEnum
	if err := json.Unmarshal([]byte(`["poweroff","reboot"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		operateActionEnumEnum = append(operateActionEnumEnum, v)
	}
}

func (m OperateActionEnum) validateOperateActionEnumEnum(path, location string, value OperateActionEnum) error {
	if err := validate.EnumCase(path, location, value, operateActionEnumEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this operate action enum
func (m OperateActionEnum) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOperateActionEnumEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this operate action enum based on context it is used
func (m OperateActionEnum) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
