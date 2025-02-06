// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Metric metric
//
// swagger:model Metric
type Metric struct {

	// typename
	// Enum: [Metric]
	Typename *string `json:"__typename,omitempty"`

	// dropped
	// Required: true
	Dropped *bool `json:"dropped"`

	// sample streams
	SampleStreams []*MetricStream `json:"sample_streams,omitempty"`

	// samples
	Samples []*MetricSample `json:"samples,omitempty"`

	// step
	// Required: true
	Step *int32 `json:"step"`

	// unit
	// Required: true
	Unit *MetricUnit `json:"unit"`

	MarshalOpts *MetricMarshalOpts `json:"-"`
}

type MetricMarshalOpts struct {
	Typename_Explicit_Null_When_Empty bool

	Dropped_Explicit_Null_When_Empty bool

	Step_Explicit_Null_When_Empty bool

	Unit_Explicit_Null_When_Empty bool
}

func (m Metric) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field __typename
	if m.Typename != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"__typename\":")
		bytes, err := swag.WriteJSON(m.Typename)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Typename_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"__typename\":null")
		first = false
	}

	// handle nullable field dropped
	if m.Dropped != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"dropped\":")
		bytes, err := swag.WriteJSON(m.Dropped)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Dropped_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"dropped\":null")
		first = false
	}

	// handle non nullable field sample_streams with omitempty
	if swag.IsZero(m.SampleStreams) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"sample_streams\":")
		bytes, err := swag.WriteJSON(m.SampleStreams)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle non nullable field samples with omitempty
	if swag.IsZero(m.Samples) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"samples\":")
		bytes, err := swag.WriteJSON(m.Samples)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field step
	if m.Step != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"step\":")
		bytes, err := swag.WriteJSON(m.Step)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Step_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"step\":null")
		first = false
	}

	// handle nullable field unit
	if m.Unit != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"unit\":")
		bytes, err := swag.WriteJSON(m.Unit)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Unit_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"unit\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this metric
func (m *Metric) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTypename(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDropped(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSampleStreams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSamples(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStep(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var metricTypeTypenamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Metric"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metricTypeTypenamePropEnum = append(metricTypeTypenamePropEnum, v)
	}
}

const (

	// MetricTypenameMetric captures enum value "Metric"
	MetricTypenameMetric string = "Metric"
)

// prop value enum
func (m *Metric) validateTypenameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, metricTypeTypenamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Metric) validateTypename(formats strfmt.Registry) error {
	if swag.IsZero(m.Typename) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypenameEnum("__typename", "body", *m.Typename); err != nil {
		return err
	}

	return nil
}

func (m *Metric) validateDropped(formats strfmt.Registry) error {

	if err := validate.Required("dropped", "body", m.Dropped); err != nil {
		return err
	}

	return nil
}

func (m *Metric) validateSampleStreams(formats strfmt.Registry) error {
	if swag.IsZero(m.SampleStreams) { // not required
		return nil
	}

	for i := 0; i < len(m.SampleStreams); i++ {
		if swag.IsZero(m.SampleStreams[i]) { // not required
			continue
		}

		if m.SampleStreams[i] != nil {
			if err := m.SampleStreams[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sample_streams" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sample_streams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Metric) validateSamples(formats strfmt.Registry) error {
	if swag.IsZero(m.Samples) { // not required
		return nil
	}

	for i := 0; i < len(m.Samples); i++ {
		if swag.IsZero(m.Samples[i]) { // not required
			continue
		}

		if m.Samples[i] != nil {
			if err := m.Samples[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("samples" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("samples" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Metric) validateStep(formats strfmt.Registry) error {

	if err := validate.Required("step", "body", m.Step); err != nil {
		return err
	}

	return nil
}

func (m *Metric) validateUnit(formats strfmt.Registry) error {

	if err := validate.Required("unit", "body", m.Unit); err != nil {
		return err
	}

	if err := validate.Required("unit", "body", m.Unit); err != nil {
		return err
	}

	if m.Unit != nil {
		if err := m.Unit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("unit")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this metric based on the context it is used
func (m *Metric) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSampleStreams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSamples(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Metric) contextValidateSampleStreams(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SampleStreams); i++ {

		if m.SampleStreams[i] != nil {
			if err := m.SampleStreams[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sample_streams" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sample_streams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Metric) contextValidateSamples(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Samples); i++ {

		if m.Samples[i] != nil {
			if err := m.Samples[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("samples" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("samples" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Metric) contextValidateUnit(ctx context.Context, formats strfmt.Registry) error {

	if m.Unit != nil {
		if err := m.Unit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("unit")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Metric) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Metric) UnmarshalBinary(b []byte) error {
	var res Metric
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
