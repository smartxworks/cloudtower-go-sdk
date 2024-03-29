// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NestedEverouteClusterControllerStatus nested everoute cluster controller status
//
// swagger:model NestedEverouteClusterControllerStatus
type NestedEverouteClusterControllerStatus struct {

	// current number
	CurrentNumber *int32 `json:"currentNumber,omitempty"`

	// expect number
	ExpectNumber *int32 `json:"expectNumber,omitempty"`

	// instances
	Instances []*NestedEverouteControllerStatus `json:"instances,omitempty"`

	// number health
	NumberHealth *int32 `json:"numberHealth,omitempty"`
}

// Validate validates this nested everoute cluster controller status
func (m *NestedEverouteClusterControllerStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstances(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedEverouteClusterControllerStatus) validateInstances(formats strfmt.Registry) error {
	if swag.IsZero(m.Instances) { // not required
		return nil
	}

	for i := 0; i < len(m.Instances); i++ {
		if swag.IsZero(m.Instances[i]) { // not required
			continue
		}

		if m.Instances[i] != nil {
			if err := m.Instances[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instances" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("instances" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this nested everoute cluster controller status based on the context it is used
func (m *NestedEverouteClusterControllerStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInstances(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedEverouteClusterControllerStatus) contextValidateInstances(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Instances); i++ {

		if m.Instances[i] != nil {
			if err := m.Instances[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instances" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("instances" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NestedEverouteClusterControllerStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedEverouteClusterControllerStatus) UnmarshalBinary(b []byte) error {
	var res NestedEverouteClusterControllerStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
