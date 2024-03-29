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

// SnapshotPlanExecuteType snapshot plan execute type
//
// swagger:model SnapshotPlanExecuteType
type SnapshotPlanExecuteType string

func NewSnapshotPlanExecuteType(value SnapshotPlanExecuteType) *SnapshotPlanExecuteType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated SnapshotPlanExecuteType.
func (m SnapshotPlanExecuteType) Pointer() *SnapshotPlanExecuteType {
	return &m
}

const (

	// SnapshotPlanExecuteTypeDAY captures enum value "DAY"
	SnapshotPlanExecuteTypeDAY SnapshotPlanExecuteType = "DAY"

	// SnapshotPlanExecuteTypeHOUR captures enum value "HOUR"
	SnapshotPlanExecuteTypeHOUR SnapshotPlanExecuteType = "HOUR"

	// SnapshotPlanExecuteTypeMINUTE captures enum value "MINUTE"
	SnapshotPlanExecuteTypeMINUTE SnapshotPlanExecuteType = "MINUTE"

	// SnapshotPlanExecuteTypeMONTH captures enum value "MONTH"
	SnapshotPlanExecuteTypeMONTH SnapshotPlanExecuteType = "MONTH"

	// SnapshotPlanExecuteTypeWEEK captures enum value "WEEK"
	SnapshotPlanExecuteTypeWEEK SnapshotPlanExecuteType = "WEEK"
)

// for schema
var snapshotPlanExecuteTypeEnum []interface{}

func init() {
	var res []SnapshotPlanExecuteType
	if err := json.Unmarshal([]byte(`["DAY","HOUR","MINUTE","MONTH","WEEK"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		snapshotPlanExecuteTypeEnum = append(snapshotPlanExecuteTypeEnum, v)
	}
}

func (m SnapshotPlanExecuteType) validateSnapshotPlanExecuteTypeEnum(path, location string, value SnapshotPlanExecuteType) error {
	if err := validate.EnumCase(path, location, value, snapshotPlanExecuteTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this snapshot plan execute type
func (m SnapshotPlanExecuteType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSnapshotPlanExecuteTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this snapshot plan execute type based on context it is used
func (m SnapshotPlanExecuteType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
