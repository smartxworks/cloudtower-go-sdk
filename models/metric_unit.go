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

// MetricUnit metric unit
//
// swagger:model MetricUnit
type MetricUnit string

func NewMetricUnit(value MetricUnit) *MetricUnit {
	return &value
}

// Pointer returns a pointer to a freshly-allocated MetricUnit.
func (m MetricUnit) Pointer() *MetricUnit {
	return &m
}

const (

	// MetricUnitCOUNT captures enum value "COUNT"
	MetricUnitCOUNT MetricUnit = "COUNT"

	// MetricUnitDATARATEBIT captures enum value "DATA_RATE_BIT"
	MetricUnitDATARATEBIT MetricUnit = "DATA_RATE_BIT"

	// MetricUnitDATARATEBYTE captures enum value "DATA_RATE_BYTE"
	MetricUnitDATARATEBYTE MetricUnit = "DATA_RATE_BYTE"

	// MetricUnitDATASIZE captures enum value "DATA_SIZE"
	MetricUnitDATASIZE MetricUnit = "DATA_SIZE"

	// MetricUnitFREQUENCY captures enum value "FREQUENCY"
	MetricUnitFREQUENCY MetricUnit = "FREQUENCY"

	// MetricUnitLOAD captures enum value "LOAD"
	MetricUnitLOAD MetricUnit = "LOAD"

	// MetricUnitPERCENT captures enum value "PERCENT"
	MetricUnitPERCENT MetricUnit = "PERCENT"

	// MetricUnitRATIO captures enum value "RATIO"
	MetricUnitRATIO MetricUnit = "RATIO"

	// MetricUnitSTORAGEBANDWIDTH captures enum value "STORAGE_BAND_WIDTH"
	MetricUnitSTORAGEBANDWIDTH MetricUnit = "STORAGE_BAND_WIDTH"

	// MetricUnitTEMPERATURE captures enum value "TEMPERATURE"
	MetricUnitTEMPERATURE MetricUnit = "TEMPERATURE"

	// MetricUnitTIME captures enum value "TIME"
	MetricUnitTIME MetricUnit = "TIME"
)

// for schema
var metricUnitEnum []interface{}

func init() {
	var res []MetricUnit
	if err := json.Unmarshal([]byte(`["COUNT","DATA_RATE_BIT","DATA_RATE_BYTE","DATA_SIZE","FREQUENCY","LOAD","PERCENT","RATIO","STORAGE_BAND_WIDTH","TEMPERATURE","TIME"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metricUnitEnum = append(metricUnitEnum, v)
	}
}

func (m MetricUnit) validateMetricUnitEnum(path, location string, value MetricUnit) error {
	if err := validate.EnumCase(path, location, value, metricUnitEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this metric unit
func (m MetricUnit) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateMetricUnitEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this metric unit based on context it is used
func (m MetricUnit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
