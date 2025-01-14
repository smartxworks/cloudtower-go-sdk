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

// BackupPlanTimePoint backup plan time point
//
// swagger:model BackupPlanTimePoint
type BackupPlanTimePoint struct {

	// date
	Date *int32 `json:"date,omitempty"`

	// date position
	DatePosition []int32 `json:"date_position,omitempty"`

	// hour
	// Required: true
	Hour *int32 `json:"hour"`

	// minute
	// Required: true
	Minute *int32 `json:"minute"`

	// weekday
	Weekday *WeekdayTypeEnum `json:"weekday,omitempty"`
}

// Validate validates this backup plan time point
func (m *BackupPlanTimePoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHour(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinute(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeekday(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupPlanTimePoint) validateHour(formats strfmt.Registry) error {

	if err := validate.Required("hour", "body", m.Hour); err != nil {
		return err
	}

	return nil
}

func (m *BackupPlanTimePoint) validateMinute(formats strfmt.Registry) error {

	if err := validate.Required("minute", "body", m.Minute); err != nil {
		return err
	}

	return nil
}

func (m *BackupPlanTimePoint) validateWeekday(formats strfmt.Registry) error {
	if swag.IsZero(m.Weekday) { // not required
		return nil
	}

	if m.Weekday != nil {
		if err := m.Weekday.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("weekday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("weekday")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this backup plan time point based on the context it is used
func (m *BackupPlanTimePoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWeekday(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupPlanTimePoint) contextValidateWeekday(ctx context.Context, formats strfmt.Registry) error {

	if m.Weekday != nil {
		if err := m.Weekday.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("weekday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("weekday")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackupPlanTimePoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupPlanTimePoint) UnmarshalBinary(b []byte) error {
	var res BackupPlanTimePoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
