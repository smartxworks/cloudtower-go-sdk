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

// CustomizeAlertRuleUpdationParams customize alert rule updation params
//
// swagger:model CustomizeAlertRuleUpdationParams
type CustomizeAlertRuleUpdationParams struct {

	// data
	// Required: true
	Data *CustomizeAlertRuleUpdationParamsData `json:"data"`

	// where
	// Required: true
	Where *GlobalAlertRuleWhereInput `json:"where"`

	MarshalOpts *CustomizeAlertRuleUpdationParamsMarshalOpts `json:"-"`
}

type CustomizeAlertRuleUpdationParamsMarshalOpts struct {
	Data_Explicit_Null_When_Empty bool

	Where_Explicit_Null_When_Empty bool
}

func (m CustomizeAlertRuleUpdationParams) MarshalJSON() ([]byte, error) {
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

// Validate validates this customize alert rule updation params
func (m *CustomizeAlertRuleUpdationParams) Validate(formats strfmt.Registry) error {
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

func (m *CustomizeAlertRuleUpdationParams) validateData(formats strfmt.Registry) error {

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

func (m *CustomizeAlertRuleUpdationParams) validateWhere(formats strfmt.Registry) error {

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

// ContextValidate validate this customize alert rule updation params based on the context it is used
func (m *CustomizeAlertRuleUpdationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *CustomizeAlertRuleUpdationParams) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CustomizeAlertRuleUpdationParams) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CustomizeAlertRuleUpdationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomizeAlertRuleUpdationParams) UnmarshalBinary(b []byte) error {
	var res CustomizeAlertRuleUpdationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CustomizeAlertRuleUpdationParamsData customize alert rule updation params data
//
// swagger:model CustomizeAlertRuleUpdationParamsData
type CustomizeAlertRuleUpdationParamsData struct {

	// clusters
	// Required: true
	Clusters *ClusterWhereInput `json:"clusters"`

	// disabled
	Disabled *bool `json:"disabled,omitempty"`

	// thresholds
	Thresholds []*AlertRuleThresholds `json:"thresholds,omitempty"`

	MarshalOpts *CustomizeAlertRuleUpdationParamsDataMarshalOpts `json:"-"`
}

type CustomizeAlertRuleUpdationParamsDataMarshalOpts struct {
	Clusters_Explicit_Null_When_Empty bool

	Disabled_Explicit_Null_When_Empty bool

	Thresholds_Explicit_Null_When_Empty bool
}

func (m CustomizeAlertRuleUpdationParamsData) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field clusters
	if m.Clusters != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"clusters\":")
		bytes, err := swag.WriteJSON(m.Clusters)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Clusters_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"clusters\":null")
		first = false
	}

	// handle nullable field disabled
	if m.Disabled != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"disabled\":")
		bytes, err := swag.WriteJSON(m.Disabled)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Disabled_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"disabled\":null")
		first = false
	}

	// handle non nullable field thresholds with omitempty
	if !swag.IsZero(m.Thresholds) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"thresholds\":")
		bytes, err := swag.WriteJSON(m.Thresholds)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this customize alert rule updation params data
func (m *CustomizeAlertRuleUpdationParamsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThresholds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomizeAlertRuleUpdationParamsData) validateClusters(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"clusters", "body", m.Clusters); err != nil {
		return err
	}

	if m.Clusters != nil {
		if err := m.Clusters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "clusters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "clusters")
			}
			return err
		}
	}

	return nil
}

func (m *CustomizeAlertRuleUpdationParamsData) validateThresholds(formats strfmt.Registry) error {
	if swag.IsZero(m.Thresholds) { // not required
		return nil
	}

	for i := 0; i < len(m.Thresholds); i++ {
		if swag.IsZero(m.Thresholds[i]) { // not required
			continue
		}

		if m.Thresholds[i] != nil {
			if err := m.Thresholds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + "thresholds" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + "thresholds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this customize alert rule updation params data based on the context it is used
func (m *CustomizeAlertRuleUpdationParamsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThresholds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomizeAlertRuleUpdationParamsData) contextValidateClusters(ctx context.Context, formats strfmt.Registry) error {

	if m.Clusters != nil {
		if err := m.Clusters.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "clusters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "clusters")
			}
			return err
		}
	}

	return nil
}

func (m *CustomizeAlertRuleUpdationParamsData) contextValidateThresholds(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Thresholds); i++ {

		if m.Thresholds[i] != nil {
			if err := m.Thresholds[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + "thresholds" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + "thresholds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomizeAlertRuleUpdationParamsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomizeAlertRuleUpdationParamsData) UnmarshalBinary(b []byte) error {
	var res CustomizeAlertRuleUpdationParamsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
