// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NestedEverouteControllerTemplate nested everoute controller template
//
// swagger:model NestedEverouteControllerTemplate
type NestedEverouteControllerTemplate struct {

	// cluster
	// Required: true
	Cluster *string `json:"cluster"`

	// gateway
	// Required: true
	Gateway *string `json:"gateway"`

	// memory
	// Required: true
	Memory *int64 `json:"memory"`

	// netmask
	// Required: true
	Netmask *string `json:"netmask"`

	// size
	// Required: true
	Size *int64 `json:"size"`

	// vcpu
	// Required: true
	Vcpu *int32 `json:"vcpu"`
}

// Validate validates this nested everoute controller template
func (m *NestedEverouteControllerTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGateway(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetmask(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcpu(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedEverouteControllerTemplate) validateCluster(formats strfmt.Registry) error {

	if err := validate.Required("cluster", "body", m.Cluster); err != nil {
		return err
	}

	return nil
}

func (m *NestedEverouteControllerTemplate) validateGateway(formats strfmt.Registry) error {

	if err := validate.Required("gateway", "body", m.Gateway); err != nil {
		return err
	}

	return nil
}

func (m *NestedEverouteControllerTemplate) validateMemory(formats strfmt.Registry) error {

	if err := validate.Required("memory", "body", m.Memory); err != nil {
		return err
	}

	return nil
}

func (m *NestedEverouteControllerTemplate) validateNetmask(formats strfmt.Registry) error {

	if err := validate.Required("netmask", "body", m.Netmask); err != nil {
		return err
	}

	return nil
}

func (m *NestedEverouteControllerTemplate) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *NestedEverouteControllerTemplate) validateVcpu(formats strfmt.Registry) error {

	if err := validate.Required("vcpu", "body", m.Vcpu); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this nested everoute controller template based on context it is used
func (m *NestedEverouteControllerTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NestedEverouteControllerTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedEverouteControllerTemplate) UnmarshalBinary(b []byte) error {
	var res NestedEverouteControllerTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
