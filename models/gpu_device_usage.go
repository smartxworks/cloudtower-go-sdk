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

// GpuDeviceUsage gpu device usage
//
// swagger:model GpuDeviceUsage
type GpuDeviceUsage string

func NewGpuDeviceUsage(value GpuDeviceUsage) *GpuDeviceUsage {
	return &value
}

// Pointer returns a pointer to a freshly-allocated GpuDeviceUsage.
func (m GpuDeviceUsage) Pointer() *GpuDeviceUsage {
	return &m
}

const (

	// GpuDeviceUsagePASSTHROUGH captures enum value "PASS_THROUGH"
	GpuDeviceUsagePASSTHROUGH GpuDeviceUsage = "PASS_THROUGH"

	// GpuDeviceUsageVGPU captures enum value "VGPU"
	GpuDeviceUsageVGPU GpuDeviceUsage = "VGPU"
)

// for schema
var gpuDeviceUsageEnum []interface{}

func init() {
	var res []GpuDeviceUsage
	if err := json.Unmarshal([]byte(`["PASS_THROUGH","VGPU"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gpuDeviceUsageEnum = append(gpuDeviceUsageEnum, v)
	}
}

func (m GpuDeviceUsage) validateGpuDeviceUsageEnum(path, location string, value GpuDeviceUsage) error {
	if err := validate.EnumCase(path, location, value, gpuDeviceUsageEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this gpu device usage
func (m GpuDeviceUsage) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateGpuDeviceUsageEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this gpu device usage based on context it is used
func (m GpuDeviceUsage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
