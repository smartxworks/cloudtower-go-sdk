// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BackupPlanUpdationParams backup plan updation params
//
// swagger:model BackupPlanUpdationParams
type BackupPlanUpdationParams struct {

	// data
	// Required: true
	Data *BackupPlanUpdationParamsData `json:"data"`

	// where
	// Required: true
	Where *BackupPlanWhereInput `json:"where"`

	MarshalOpts *BackupPlanUpdationParamsMarshalOpts `json:"-"`
}

type BackupPlanUpdationParamsMarshalOpts struct {
	Data_Explicit_Null_When_Empty bool

	Where_Explicit_Null_When_Empty bool
}

func (m BackupPlanUpdationParams) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field data
	if m.Data != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"data\":")
		bytes, err := swag.WriteJSON(m.Data)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Data_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"data\":null")
		first = false
	}

	// handle nullable field where
	if m.Where != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"where\":")
		bytes, err := swag.WriteJSON(m.Where)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Where_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"where\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this backup plan updation params
func (m *BackupPlanUpdationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWhere(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupPlanUpdationParams) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlanUpdationParams) validateWhere(formats strfmt.Registry) error {

	if err := validate.Required("where", "body", m.Where); err != nil {
		return err
	}

	if m.Where != nil {
		if err := m.Where.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("where")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("where")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this backup plan updation params based on the context it is used
func (m *BackupPlanUpdationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWhere(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupPlanUpdationParams) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if m.Data != nil {
		if err := m.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlanUpdationParams) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

	if m.Where != nil {
		if err := m.Where.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("where")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("where")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackupPlanUpdationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupPlanUpdationParams) UnmarshalBinary(b []byte) error {
	var res BackupPlanUpdationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BackupPlanUpdationParamsData backup plan updation params data
//
// swagger:model BackupPlanUpdationParamsData
type BackupPlanUpdationParamsData struct {

	// backup delay option
	BackupDelayOption *BackupPlanDelayOption `json:"backup_delay_option,omitempty"`

	// compression
	Compression *bool `json:"compression,omitempty"`

	// description
	Description *string `json:"description,omitempty"`

	// disconnect strategy
	DisconnectStrategy *BackupPlanDeleteStrategy `json:"disconnect_strategy,omitempty"`

	// enable window
	EnableWindow *bool `json:"enable_window,omitempty"`

	// full interval
	FullInterval *int32 `json:"full_interval,omitempty"`

	// full period
	FullPeriod *BackupPlanPeriod `json:"full_period,omitempty"`

	// full time point
	FullTimePoint *BackupPlanTimePoint `json:"full_time_point,omitempty"`

	// incremental interval
	IncrementalInterval *int32 `json:"incremental_interval,omitempty"`

	// incremental period
	IncrementalPeriod *BackupPlanPeriod `json:"incremental_period,omitempty"`

	// incremental time points
	IncrementalTimePoints []*BackupPlanTimePoint `json:"incremental_time_points,omitempty"`

	// incremental weekdays
	IncrementalWeekdays []WeekdayTypeEnum `json:"incremental_weekdays,omitempty"`

	// keep policy
	KeepPolicy *BackupPlanKeepPolicy `json:"keep_policy,omitempty"`

	// keep policy value
	KeepPolicyValue *int32 `json:"keep_policy_value,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// snapshot consistent type
	SnapshotConsistentType *ConsistentType `json:"snapshot_consistent_type,omitempty"`

	// vms
	Vms *VMWhereInput `json:"vms,omitempty"`

	// window end
	WindowEnd *string `json:"window_end,omitempty"`

	// window start
	WindowStart *string `json:"window_start,omitempty"`

	MarshalOpts *BackupPlanUpdationParamsDataMarshalOpts `json:"-"`
}

type BackupPlanUpdationParamsDataMarshalOpts struct {
	BackupDelayOption_Explicit_Null_When_Empty bool

	Compression_Explicit_Null_When_Empty bool

	Description_Explicit_Null_When_Empty bool

	DisconnectStrategy_Explicit_Null_When_Empty bool

	EnableWindow_Explicit_Null_When_Empty bool

	FullInterval_Explicit_Null_When_Empty bool

	FullPeriod_Explicit_Null_When_Empty bool

	FullTimePoint_Explicit_Null_When_Empty bool

	IncrementalInterval_Explicit_Null_When_Empty bool

	IncrementalPeriod_Explicit_Null_When_Empty bool

	IncrementalTimePoints_Explicit_Null_When_Empty bool

	IncrementalWeekdays_Explicit_Null_When_Empty bool

	KeepPolicy_Explicit_Null_When_Empty bool

	KeepPolicyValue_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	SnapshotConsistentType_Explicit_Null_When_Empty bool

	Vms_Explicit_Null_When_Empty bool

	WindowEnd_Explicit_Null_When_Empty bool

	WindowStart_Explicit_Null_When_Empty bool
}

func (m BackupPlanUpdationParamsData) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field backup_delay_option
	if m.BackupDelayOption != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"backup_delay_option\":")
		bytes, err := swag.WriteJSON(m.BackupDelayOption)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.BackupDelayOption_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"backup_delay_option\":null")
		first = false
	}

	// handle nullable field compression
	if m.Compression != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"compression\":")
		bytes, err := swag.WriteJSON(m.Compression)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Compression_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"compression\":null")
		first = false
	}

	// handle nullable field description
	if m.Description != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description\":")
		bytes, err := swag.WriteJSON(m.Description)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Description_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description\":null")
		first = false
	}

	// handle nullable field disconnect_strategy
	if m.DisconnectStrategy != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"disconnect_strategy\":")
		bytes, err := swag.WriteJSON(m.DisconnectStrategy)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.DisconnectStrategy_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"disconnect_strategy\":null")
		first = false
	}

	// handle nullable field enable_window
	if m.EnableWindow != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"enable_window\":")
		bytes, err := swag.WriteJSON(m.EnableWindow)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.EnableWindow_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"enable_window\":null")
		first = false
	}

	// handle nullable field full_interval
	if m.FullInterval != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"full_interval\":")
		bytes, err := swag.WriteJSON(m.FullInterval)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.FullInterval_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"full_interval\":null")
		first = false
	}

	// handle nullable field full_period
	if m.FullPeriod != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"full_period\":")
		bytes, err := swag.WriteJSON(m.FullPeriod)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.FullPeriod_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"full_period\":null")
		first = false
	}

	// handle nullable field full_time_point
	if m.FullTimePoint != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"full_time_point\":")
		bytes, err := swag.WriteJSON(m.FullTimePoint)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.FullTimePoint_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"full_time_point\":null")
		first = false
	}

	// handle nullable field incremental_interval
	if m.IncrementalInterval != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"incremental_interval\":")
		bytes, err := swag.WriteJSON(m.IncrementalInterval)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IncrementalInterval_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"incremental_interval\":null")
		first = false
	}

	// handle nullable field incremental_period
	if m.IncrementalPeriod != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"incremental_period\":")
		bytes, err := swag.WriteJSON(m.IncrementalPeriod)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IncrementalPeriod_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"incremental_period\":null")
		first = false
	}

	// handle non nullable field incremental_time_points with omitempty
	if !swag.IsZero(m.IncrementalTimePoints) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"incremental_time_points\":")
		bytes, err := swag.WriteJSON(m.IncrementalTimePoints)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle non nullable field incremental_weekdays with omitempty
	if !swag.IsZero(m.IncrementalWeekdays) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"incremental_weekdays\":")
		bytes, err := swag.WriteJSON(m.IncrementalWeekdays)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field keep_policy
	if m.KeepPolicy != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"keep_policy\":")
		bytes, err := swag.WriteJSON(m.KeepPolicy)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.KeepPolicy_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"keep_policy\":null")
		first = false
	}

	// handle nullable field keep_policy_value
	if m.KeepPolicyValue != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"keep_policy_value\":")
		bytes, err := swag.WriteJSON(m.KeepPolicyValue)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.KeepPolicyValue_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"keep_policy_value\":null")
		first = false
	}

	// handle nullable field name
	if m.Name != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name\":")
		bytes, err := swag.WriteJSON(m.Name)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Name_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name\":null")
		first = false
	}

	// handle nullable field snapshot_consistent_type
	if m.SnapshotConsistentType != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"snapshot_consistent_type\":")
		bytes, err := swag.WriteJSON(m.SnapshotConsistentType)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.SnapshotConsistentType_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"snapshot_consistent_type\":null")
		first = false
	}

	// handle nullable field vms
	if m.Vms != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vms\":")
		bytes, err := swag.WriteJSON(m.Vms)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Vms_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vms\":null")
		first = false
	}

	// handle nullable field window_end
	if m.WindowEnd != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"window_end\":")
		bytes, err := swag.WriteJSON(m.WindowEnd)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.WindowEnd_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"window_end\":null")
		first = false
	}

	// handle nullable field window_start
	if m.WindowStart != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"window_start\":")
		bytes, err := swag.WriteJSON(m.WindowStart)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.WindowStart_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"window_start\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this backup plan updation params data
func (m *BackupPlanUpdationParamsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupDelayOption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisconnectStrategy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFullPeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFullTimePoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncrementalPeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncrementalTimePoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncrementalWeekdays(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeepPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotConsistentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVms(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupPlanUpdationParamsData) validateBackupDelayOption(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupDelayOption) { // not required
		return nil
	}

	if m.BackupDelayOption != nil {
		if err := m.BackupDelayOption.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "backup_delay_option")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "backup_delay_option")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlanUpdationParamsData) validateDisconnectStrategy(formats strfmt.Registry) error {
	if swag.IsZero(m.DisconnectStrategy) { // not required
		return nil
	}

	if m.DisconnectStrategy != nil {
		if err := m.DisconnectStrategy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "disconnect_strategy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "disconnect_strategy")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlanUpdationParamsData) validateFullPeriod(formats strfmt.Registry) error {
	if swag.IsZero(m.FullPeriod) { // not required
		return nil
	}

	if m.FullPeriod != nil {
		if err := m.FullPeriod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "full_period")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "full_period")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlanUpdationParamsData) validateFullTimePoint(formats strfmt.Registry) error {
	if swag.IsZero(m.FullTimePoint) { // not required
		return nil
	}

	if m.FullTimePoint != nil {
		if err := m.FullTimePoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "full_time_point")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "full_time_point")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlanUpdationParamsData) validateIncrementalPeriod(formats strfmt.Registry) error {
	if swag.IsZero(m.IncrementalPeriod) { // not required
		return nil
	}

	if m.IncrementalPeriod != nil {
		if err := m.IncrementalPeriod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "incremental_period")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "incremental_period")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlanUpdationParamsData) validateIncrementalTimePoints(formats strfmt.Registry) error {
	if swag.IsZero(m.IncrementalTimePoints) { // not required
		return nil
	}

	for i := 0; i < len(m.IncrementalTimePoints); i++ {
		if swag.IsZero(m.IncrementalTimePoints[i]) { // not required
			continue
		}

		if m.IncrementalTimePoints[i] != nil {
			if err := m.IncrementalTimePoints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + "incremental_time_points" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + "incremental_time_points" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BackupPlanUpdationParamsData) validateIncrementalWeekdays(formats strfmt.Registry) error {
	if swag.IsZero(m.IncrementalWeekdays) { // not required
		return nil
	}

	for i := 0; i < len(m.IncrementalWeekdays); i++ {

		if err := m.IncrementalWeekdays[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "incremental_weekdays" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "incremental_weekdays" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *BackupPlanUpdationParamsData) validateKeepPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.KeepPolicy) { // not required
		return nil
	}

	if m.KeepPolicy != nil {
		if err := m.KeepPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "keep_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "keep_policy")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlanUpdationParamsData) validateSnapshotConsistentType(formats strfmt.Registry) error {
	if swag.IsZero(m.SnapshotConsistentType) { // not required
		return nil
	}

	if m.SnapshotConsistentType != nil {
		if err := m.SnapshotConsistentType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "snapshot_consistent_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "snapshot_consistent_type")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlanUpdationParamsData) validateVms(formats strfmt.Registry) error {
	if swag.IsZero(m.Vms) { // not required
		return nil
	}

	if m.Vms != nil {
		if err := m.Vms.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "vms")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "vms")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this backup plan updation params data based on the context it is used
func (m *BackupPlanUpdationParamsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBackupDelayOption(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisconnectStrategy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFullPeriod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFullTimePoint(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIncrementalPeriod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIncrementalTimePoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIncrementalWeekdays(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKeepPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnapshotConsistentType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupPlanUpdationParamsData) contextValidateBackupDelayOption(ctx context.Context, formats strfmt.Registry) error {

	if m.BackupDelayOption != nil {
		if err := m.BackupDelayOption.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "backup_delay_option")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "backup_delay_option")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlanUpdationParamsData) contextValidateDisconnectStrategy(ctx context.Context, formats strfmt.Registry) error {

	if m.DisconnectStrategy != nil {
		if err := m.DisconnectStrategy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "disconnect_strategy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "disconnect_strategy")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlanUpdationParamsData) contextValidateFullPeriod(ctx context.Context, formats strfmt.Registry) error {

	if m.FullPeriod != nil {
		if err := m.FullPeriod.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "full_period")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "full_period")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlanUpdationParamsData) contextValidateFullTimePoint(ctx context.Context, formats strfmt.Registry) error {

	if m.FullTimePoint != nil {
		if err := m.FullTimePoint.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "full_time_point")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "full_time_point")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlanUpdationParamsData) contextValidateIncrementalPeriod(ctx context.Context, formats strfmt.Registry) error {

	if m.IncrementalPeriod != nil {
		if err := m.IncrementalPeriod.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "incremental_period")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "incremental_period")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlanUpdationParamsData) contextValidateIncrementalTimePoints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IncrementalTimePoints); i++ {

		if m.IncrementalTimePoints[i] != nil {
			if err := m.IncrementalTimePoints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + "incremental_time_points" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + "incremental_time_points" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BackupPlanUpdationParamsData) contextValidateIncrementalWeekdays(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IncrementalWeekdays); i++ {

		if err := m.IncrementalWeekdays[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "incremental_weekdays" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "incremental_weekdays" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *BackupPlanUpdationParamsData) contextValidateKeepPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.KeepPolicy != nil {
		if err := m.KeepPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "keep_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "keep_policy")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlanUpdationParamsData) contextValidateSnapshotConsistentType(ctx context.Context, formats strfmt.Registry) error {

	if m.SnapshotConsistentType != nil {
		if err := m.SnapshotConsistentType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "snapshot_consistent_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "snapshot_consistent_type")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlanUpdationParamsData) contextValidateVms(ctx context.Context, formats strfmt.Registry) error {

	if m.Vms != nil {
		if err := m.Vms.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "vms")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "vms")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackupPlanUpdationParamsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupPlanUpdationParamsData) UnmarshalBinary(b []byte) error {
	var res BackupPlanUpdationParamsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
