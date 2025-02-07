// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NestedMetroAvailabilityChecklist nested metro availability checklist
//
// swagger:model NestedMetroAvailabilityChecklist
type NestedMetroAvailabilityChecklist struct {

	// primary zone
	// Required: true
	PrimaryZone *NestedMetroCheckResult `json:"primaryZone"`

	// primary zone and witness
	// Required: true
	PrimaryZoneAndWitness *NestedMetroCheckResult `json:"primaryZoneAndWitness"`

	// secondary zone
	// Required: true
	SecondaryZone *NestedMetroCheckResult `json:"secondaryZone"`

	// secondary zone and witness
	// Required: true
	SecondaryZoneAndWitness *NestedMetroCheckResult `json:"secondaryZoneAndWitness"`

	// witness
	// Required: true
	Witness *NestedMetroCheckResult `json:"witness"`

	// zone and zone
	// Required: true
	ZoneAndZone *NestedMetroCheckResult `json:"zoneAndZone"`

	MarshalOpts *NestedMetroAvailabilityChecklistMarshalOpts `json:"-"`
}

type NestedMetroAvailabilityChecklistMarshalOpts struct {
	PrimaryZone_Explicit_Null_When_Empty bool

	PrimaryZoneAndWitness_Explicit_Null_When_Empty bool

	SecondaryZone_Explicit_Null_When_Empty bool

	SecondaryZoneAndWitness_Explicit_Null_When_Empty bool

	Witness_Explicit_Null_When_Empty bool

	ZoneAndZone_Explicit_Null_When_Empty bool
}

func (m NestedMetroAvailabilityChecklist) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field primaryZone
	if m.PrimaryZone != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"primaryZone\":")
		bytes, err := swag.WriteJSON(m.PrimaryZone)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.PrimaryZone_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"primaryZone\":null")
		first = false
	}

	// handle nullable field primaryZoneAndWitness
	if m.PrimaryZoneAndWitness != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"primaryZoneAndWitness\":")
		bytes, err := swag.WriteJSON(m.PrimaryZoneAndWitness)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.PrimaryZoneAndWitness_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"primaryZoneAndWitness\":null")
		first = false
	}

	// handle nullable field secondaryZone
	if m.SecondaryZone != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"secondaryZone\":")
		bytes, err := swag.WriteJSON(m.SecondaryZone)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.SecondaryZone_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"secondaryZone\":null")
		first = false
	}

	// handle nullable field secondaryZoneAndWitness
	if m.SecondaryZoneAndWitness != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"secondaryZoneAndWitness\":")
		bytes, err := swag.WriteJSON(m.SecondaryZoneAndWitness)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.SecondaryZoneAndWitness_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"secondaryZoneAndWitness\":null")
		first = false
	}

	// handle nullable field witness
	if m.Witness != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"witness\":")
		bytes, err := swag.WriteJSON(m.Witness)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Witness_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"witness\":null")
		first = false
	}

	// handle nullable field zoneAndZone
	if m.ZoneAndZone != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"zoneAndZone\":")
		bytes, err := swag.WriteJSON(m.ZoneAndZone)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ZoneAndZone_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"zoneAndZone\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this nested metro availability checklist
func (m *NestedMetroAvailabilityChecklist) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrimaryZone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryZoneAndWitness(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondaryZone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondaryZoneAndWitness(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWitness(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZoneAndZone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedMetroAvailabilityChecklist) validatePrimaryZone(formats strfmt.Registry) error {

	if err := validate.Required("primaryZone", "body", m.PrimaryZone); err != nil {
		return err
	}

	if m.PrimaryZone != nil {
		if err := m.PrimaryZone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primaryZone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primaryZone")
			}
			return err
		}
	}

	return nil
}

func (m *NestedMetroAvailabilityChecklist) validatePrimaryZoneAndWitness(formats strfmt.Registry) error {

	if err := validate.Required("primaryZoneAndWitness", "body", m.PrimaryZoneAndWitness); err != nil {
		return err
	}

	if m.PrimaryZoneAndWitness != nil {
		if err := m.PrimaryZoneAndWitness.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primaryZoneAndWitness")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primaryZoneAndWitness")
			}
			return err
		}
	}

	return nil
}

func (m *NestedMetroAvailabilityChecklist) validateSecondaryZone(formats strfmt.Registry) error {

	if err := validate.Required("secondaryZone", "body", m.SecondaryZone); err != nil {
		return err
	}

	if m.SecondaryZone != nil {
		if err := m.SecondaryZone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secondaryZone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("secondaryZone")
			}
			return err
		}
	}

	return nil
}

func (m *NestedMetroAvailabilityChecklist) validateSecondaryZoneAndWitness(formats strfmt.Registry) error {

	if err := validate.Required("secondaryZoneAndWitness", "body", m.SecondaryZoneAndWitness); err != nil {
		return err
	}

	if m.SecondaryZoneAndWitness != nil {
		if err := m.SecondaryZoneAndWitness.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secondaryZoneAndWitness")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("secondaryZoneAndWitness")
			}
			return err
		}
	}

	return nil
}

func (m *NestedMetroAvailabilityChecklist) validateWitness(formats strfmt.Registry) error {

	if err := validate.Required("witness", "body", m.Witness); err != nil {
		return err
	}

	if m.Witness != nil {
		if err := m.Witness.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("witness")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("witness")
			}
			return err
		}
	}

	return nil
}

func (m *NestedMetroAvailabilityChecklist) validateZoneAndZone(formats strfmt.Registry) error {

	if err := validate.Required("zoneAndZone", "body", m.ZoneAndZone); err != nil {
		return err
	}

	if m.ZoneAndZone != nil {
		if err := m.ZoneAndZone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("zoneAndZone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("zoneAndZone")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nested metro availability checklist based on the context it is used
func (m *NestedMetroAvailabilityChecklist) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePrimaryZone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrimaryZoneAndWitness(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecondaryZone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecondaryZoneAndWitness(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWitness(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateZoneAndZone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedMetroAvailabilityChecklist) contextValidatePrimaryZone(ctx context.Context, formats strfmt.Registry) error {

	if m.PrimaryZone != nil {
		if err := m.PrimaryZone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primaryZone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primaryZone")
			}
			return err
		}
	}

	return nil
}

func (m *NestedMetroAvailabilityChecklist) contextValidatePrimaryZoneAndWitness(ctx context.Context, formats strfmt.Registry) error {

	if m.PrimaryZoneAndWitness != nil {
		if err := m.PrimaryZoneAndWitness.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primaryZoneAndWitness")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("primaryZoneAndWitness")
			}
			return err
		}
	}

	return nil
}

func (m *NestedMetroAvailabilityChecklist) contextValidateSecondaryZone(ctx context.Context, formats strfmt.Registry) error {

	if m.SecondaryZone != nil {
		if err := m.SecondaryZone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secondaryZone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("secondaryZone")
			}
			return err
		}
	}

	return nil
}

func (m *NestedMetroAvailabilityChecklist) contextValidateSecondaryZoneAndWitness(ctx context.Context, formats strfmt.Registry) error {

	if m.SecondaryZoneAndWitness != nil {
		if err := m.SecondaryZoneAndWitness.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secondaryZoneAndWitness")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("secondaryZoneAndWitness")
			}
			return err
		}
	}

	return nil
}

func (m *NestedMetroAvailabilityChecklist) contextValidateWitness(ctx context.Context, formats strfmt.Registry) error {

	if m.Witness != nil {
		if err := m.Witness.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("witness")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("witness")
			}
			return err
		}
	}

	return nil
}

func (m *NestedMetroAvailabilityChecklist) contextValidateZoneAndZone(ctx context.Context, formats strfmt.Registry) error {

	if m.ZoneAndZone != nil {
		if err := m.ZoneAndZone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("zoneAndZone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("zoneAndZone")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NestedMetroAvailabilityChecklist) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedMetroAvailabilityChecklist) UnmarshalBinary(b []byte) error {
	var res NestedMetroAvailabilityChecklist
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
