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

// PciDeviceSriovState pci device sriov state
//
// swagger:model PciDeviceSriovState
type PciDeviceSriovState string

func NewPciDeviceSriovState(value PciDeviceSriovState) *PciDeviceSriovState {
	return &value
}

// Pointer returns a pointer to a freshly-allocated PciDeviceSriovState.
func (m PciDeviceSriovState) Pointer() *PciDeviceSriovState {
	return &m
}

const (

	// PciDeviceSriovStateENABLED captures enum value "ENABLED"
	PciDeviceSriovStateENABLED PciDeviceSriovState = "ENABLED"

	// PciDeviceSriovStateNOTENABLED captures enum value "NOT_ENABLED"
	PciDeviceSriovStateNOTENABLED PciDeviceSriovState = "NOT_ENABLED"

	// PciDeviceSriovStateNOTSUPPORT captures enum value "NOT_SUPPORT"
	PciDeviceSriovStateNOTSUPPORT PciDeviceSriovState = "NOT_SUPPORT"
)

// for schema
var pciDeviceSriovStateEnum []interface{}

func init() {
	var res []PciDeviceSriovState
	if err := json.Unmarshal([]byte(`["ENABLED","NOT_ENABLED","NOT_SUPPORT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pciDeviceSriovStateEnum = append(pciDeviceSriovStateEnum, v)
	}
}

func (m PciDeviceSriovState) validatePciDeviceSriovStateEnum(path, location string, value PciDeviceSriovState) error {
	if err := validate.EnumCase(path, location, value, pciDeviceSriovStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this pci device sriov state
func (m PciDeviceSriovState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePciDeviceSriovStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this pci device sriov state based on context it is used
func (m PciDeviceSriovState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
