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

// Hypervisor hypervisor
//
// swagger:model Hypervisor
type Hypervisor string

func NewHypervisor(value Hypervisor) *Hypervisor {
	return &value
}

// Pointer returns a pointer to a freshly-allocated Hypervisor.
func (m Hypervisor) Pointer() *Hypervisor {
	return &m
}

const (

	// HypervisorBLUESHARK captures enum value "BLUESHARK"
	HypervisorBLUESHARK Hypervisor = "BLUESHARK"

	// HypervisorELF captures enum value "ELF"
	HypervisorELF Hypervisor = "ELF"

	// HypervisorVMWARE captures enum value "VMWARE"
	HypervisorVMWARE Hypervisor = "VMWARE"

	// HypervisorXENSERVER captures enum value "XENSERVER"
	HypervisorXENSERVER Hypervisor = "XENSERVER"
)

// for schema
var hypervisorEnum []interface{}

func init() {
	var res []Hypervisor
	if err := json.Unmarshal([]byte(`["BLUESHARK","ELF","VMWARE","XENSERVER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hypervisorEnum = append(hypervisorEnum, v)
	}
}

func (m Hypervisor) validateHypervisorEnum(path, location string, value Hypervisor) error {
	if err := validate.EnumCase(path, location, value, hypervisorEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hypervisor
func (m Hypervisor) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHypervisorEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hypervisor based on context it is used
func (m Hypervisor) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
