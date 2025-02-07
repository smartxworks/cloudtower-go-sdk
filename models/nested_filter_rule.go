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

// NestedFilterRule nested filter rule
//
// swagger:model NestedFilterRule
type NestedFilterRule struct {

	// aggregation
	// Required: true
	Aggregation *FilterRuleAggregationEnum `json:"aggregation"`

	// duration
	// Required: true
	Duration *int32 `json:"duration"`

	// metric
	// Required: true
	Metric *FilterRuleMetricEnum `json:"metric"`

	// op
	// Required: true
	Op *FilterRuleOpEnum `json:"op"`

	// quantile
	// Required: true
	Quantile *int32 `json:"quantile"`

	// threshold
	// Required: true
	Threshold *float64 `json:"threshold"`

	MarshalOpts *NestedFilterRuleMarshalOpts `json:"-"`
}

type NestedFilterRuleMarshalOpts struct {
	Aggregation_Explicit_Null_When_Empty bool

	Duration_Explicit_Null_When_Empty bool

	Metric_Explicit_Null_When_Empty bool

	Op_Explicit_Null_When_Empty bool

	Quantile_Explicit_Null_When_Empty bool

	Threshold_Explicit_Null_When_Empty bool
}

func (m NestedFilterRule) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field aggregation
	if m.Aggregation != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"aggregation\":")
		bytes, err := swag.WriteJSON(m.Aggregation)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Aggregation_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"aggregation\":null")
		first = false
	}

	// handle nullable field duration
	if m.Duration != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"duration\":")
		bytes, err := swag.WriteJSON(m.Duration)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Duration_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"duration\":null")
		first = false
	}

	// handle nullable field metric
	if m.Metric != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"metric\":")
		bytes, err := swag.WriteJSON(m.Metric)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Metric_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"metric\":null")
		first = false
	}

	// handle nullable field op
	if m.Op != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"op\":")
		bytes, err := swag.WriteJSON(m.Op)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Op_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"op\":null")
		first = false
	}

	// handle nullable field quantile
	if m.Quantile != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"quantile\":")
		bytes, err := swag.WriteJSON(m.Quantile)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Quantile_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"quantile\":null")
		first = false
	}

	// handle nullable field threshold
	if m.Threshold != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"threshold\":")
		bytes, err := swag.WriteJSON(m.Threshold)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Threshold_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"threshold\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this nested filter rule
func (m *NestedFilterRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAggregation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetric(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuantile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThreshold(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedFilterRule) validateAggregation(formats strfmt.Registry) error {

	if err := validate.Required("aggregation", "body", m.Aggregation); err != nil {
		return err
	}

	if err := validate.Required("aggregation", "body", m.Aggregation); err != nil {
		return err
	}

	if m.Aggregation != nil {
		if err := m.Aggregation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aggregation")
			}
			return err
		}
	}

	return nil
}

func (m *NestedFilterRule) validateDuration(formats strfmt.Registry) error {

	if err := validate.Required("duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *NestedFilterRule) validateMetric(formats strfmt.Registry) error {

	if err := validate.Required("metric", "body", m.Metric); err != nil {
		return err
	}

	if err := validate.Required("metric", "body", m.Metric); err != nil {
		return err
	}

	if m.Metric != nil {
		if err := m.Metric.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metric")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metric")
			}
			return err
		}
	}

	return nil
}

func (m *NestedFilterRule) validateOp(formats strfmt.Registry) error {

	if err := validate.Required("op", "body", m.Op); err != nil {
		return err
	}

	if err := validate.Required("op", "body", m.Op); err != nil {
		return err
	}

	if m.Op != nil {
		if err := m.Op.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("op")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("op")
			}
			return err
		}
	}

	return nil
}

func (m *NestedFilterRule) validateQuantile(formats strfmt.Registry) error {

	if err := validate.Required("quantile", "body", m.Quantile); err != nil {
		return err
	}

	return nil
}

func (m *NestedFilterRule) validateThreshold(formats strfmt.Registry) error {

	if err := validate.Required("threshold", "body", m.Threshold); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this nested filter rule based on the context it is used
func (m *NestedFilterRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAggregation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetric(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedFilterRule) contextValidateAggregation(ctx context.Context, formats strfmt.Registry) error {

	if m.Aggregation != nil {
		if err := m.Aggregation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aggregation")
			}
			return err
		}
	}

	return nil
}

func (m *NestedFilterRule) contextValidateMetric(ctx context.Context, formats strfmt.Registry) error {

	if m.Metric != nil {
		if err := m.Metric.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metric")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metric")
			}
			return err
		}
	}

	return nil
}

func (m *NestedFilterRule) contextValidateOp(ctx context.Context, formats strfmt.Registry) error {

	if m.Op != nil {
		if err := m.Op.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("op")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("op")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NestedFilterRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedFilterRule) UnmarshalBinary(b []byte) error {
	var res NestedFilterRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
