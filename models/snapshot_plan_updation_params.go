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

// SnapshotPlanUpdationParams snapshot plan updation params
//
// swagger:model SnapshotPlanUpdationParams
type SnapshotPlanUpdationParams struct {

	// data
	Data *SnapshotPlanUpdationParamsData `json:"data,omitempty"`

	// where
	// Required: true
	Where *SnapshotPlanWhereInput `json:"where"`

	MarshalOpts *SnapshotPlanUpdationParamsMarshalOpts `json:"-"`
}

type SnapshotPlanUpdationParamsMarshalOpts struct {
	Data_Explicit_Null_When_Empty bool

	Where_Explicit_Null_When_Empty bool
}

func (m SnapshotPlanUpdationParams) MarshalJSON() ([]byte, error) {
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

// Validate validates this snapshot plan updation params
func (m *SnapshotPlanUpdationParams) Validate(formats strfmt.Registry) error {
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

func (m *SnapshotPlanUpdationParams) validateData(formats strfmt.Registry) error {
	if swag.IsZero(m.Data) { // not required
		return nil
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

func (m *SnapshotPlanUpdationParams) validateWhere(formats strfmt.Registry) error {

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

// ContextValidate validate this snapshot plan updation params based on the context it is used
func (m *SnapshotPlanUpdationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *SnapshotPlanUpdationParams) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SnapshotPlanUpdationParams) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SnapshotPlanUpdationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotPlanUpdationParams) UnmarshalBinary(b []byte) error {
	var res SnapshotPlanUpdationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapshotPlanUpdationParamsData snapshot plan updation params data
//
// swagger:model SnapshotPlanUpdationParamsData
type SnapshotPlanUpdationParamsData struct {

	// end time
	EndTime *string `json:"end_time,omitempty"`

	// exec h m
	Exechm *string `json:"exec_h_m,omitempty"`

	// execute intervals
	ExecuteIntervals []int32 `json:"execute_intervals,omitempty"`

	// execute plan type
	ExecutePlanType *SnapshotPlanExecuteType `json:"execute_plan_type,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// remain snapshot num
	RemainSnapshotNum *int32 `json:"remain_snapshot_num,omitempty"`

	// vm ids
	VMIds []string `json:"vm_ids,omitempty"`

	MarshalOpts *SnapshotPlanUpdationParamsDataMarshalOpts `json:"-"`
}

type SnapshotPlanUpdationParamsDataMarshalOpts struct {
	EndTime_Explicit_Null_When_Empty bool

	Exechm_Explicit_Null_When_Empty bool

	ExecuteIntervals_Explicit_Null_When_Empty bool

	ExecutePlanType_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	RemainSnapshotNum_Explicit_Null_When_Empty bool

	VMIds_Explicit_Null_When_Empty bool
}

func (m SnapshotPlanUpdationParamsData) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field end_time
	if m.EndTime != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"end_time\":")
		bytes, err := swag.WriteJSON(m.EndTime)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.EndTime_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"end_time\":null")
		first = false
	}

	// handle nullable field exec_h_m
	if m.Exechm != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"exec_h_m\":")
		bytes, err := swag.WriteJSON(m.Exechm)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Exechm_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"exec_h_m\":null")
		first = false
	}

	// handle non nullable field execute_intervals with omitempty
	if swag.IsZero(m.ExecuteIntervals) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"execute_intervals\":")
		bytes, err := swag.WriteJSON(m.ExecuteIntervals)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field execute_plan_type
	if m.ExecutePlanType != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"execute_plan_type\":")
		bytes, err := swag.WriteJSON(m.ExecutePlanType)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ExecutePlanType_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"execute_plan_type\":null")
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

	// handle nullable field remain_snapshot_num
	if m.RemainSnapshotNum != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"remain_snapshot_num\":")
		bytes, err := swag.WriteJSON(m.RemainSnapshotNum)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.RemainSnapshotNum_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"remain_snapshot_num\":null")
		first = false
	}

	// handle non nullable field vm_ids with omitempty
	if swag.IsZero(m.VMIds) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_ids\":")
		bytes, err := swag.WriteJSON(m.VMIds)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this snapshot plan updation params data
func (m *SnapshotPlanUpdationParamsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExecutePlanType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotPlanUpdationParamsData) validateExecutePlanType(formats strfmt.Registry) error {
	if swag.IsZero(m.ExecutePlanType) { // not required
		return nil
	}

	if m.ExecutePlanType != nil {
		if err := m.ExecutePlanType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "execute_plan_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "execute_plan_type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snapshot plan updation params data based on the context it is used
func (m *SnapshotPlanUpdationParamsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExecutePlanType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotPlanUpdationParamsData) contextValidateExecutePlanType(ctx context.Context, formats strfmt.Registry) error {

	if m.ExecutePlanType != nil {
		if err := m.ExecutePlanType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "execute_plan_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "execute_plan_type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotPlanUpdationParamsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotPlanUpdationParamsData) UnmarshalBinary(b []byte) error {
	var res SnapshotPlanUpdationParamsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
