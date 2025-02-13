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

// PciDevicePartitionState pci device partition state
//
// swagger:model PciDevicePartitionState
type PciDevicePartitionState string

func NewPciDevicePartitionState(value PciDevicePartitionState) *PciDevicePartitionState {
	return &value
}

// Pointer returns a pointer to a freshly-allocated PciDevicePartitionState.
func (m PciDevicePartitionState) Pointer() *PciDevicePartitionState {
	return &m
}

const (

	// PciDevicePartitionStateENABLED captures enum value "ENABLED"
	PciDevicePartitionStateENABLED PciDevicePartitionState = "ENABLED"

	// PciDevicePartitionStateNOTENABLED captures enum value "NOT_ENABLED"
	PciDevicePartitionStateNOTENABLED PciDevicePartitionState = "NOT_ENABLED"

	// PciDevicePartitionStateNOTSUPPORT captures enum value "NOT_SUPPORT"
	PciDevicePartitionStateNOTSUPPORT PciDevicePartitionState = "NOT_SUPPORT"
)

// for schema
var pciDevicePartitionStateEnum []interface{}

func init() {
	var res []PciDevicePartitionState
	if err := json.Unmarshal([]byte(`["ENABLED","NOT_ENABLED","NOT_SUPPORT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pciDevicePartitionStateEnum = append(pciDevicePartitionStateEnum, v)
	}
}

func (m PciDevicePartitionState) validatePciDevicePartitionStateEnum(path, location string, value PciDevicePartitionState) error {
	if err := validate.EnumCase(path, location, value, pciDevicePartitionStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this pci device partition state
func (m PciDevicePartitionState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePciDevicePartitionStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this pci device partition state based on context it is used
func (m PciDevicePartitionState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
