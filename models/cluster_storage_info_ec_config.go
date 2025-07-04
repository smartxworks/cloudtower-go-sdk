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

// ClusterStorageInfoEcConfig cluster storage info ec config
//
// swagger:model ClusterStorageInfoEcConfig
type ClusterStorageInfoEcConfig struct {

	// k
	// Required: true
	K *int32 `json:"k"`

	// m
	// Required: true
	M *int32 `json:"m"`

	MarshalOpts *ClusterStorageInfoEcConfigMarshalOpts `json:"-"`
}

type ClusterStorageInfoEcConfigMarshalOpts struct {
	K_Explicit_Null_When_Empty bool

	M_Explicit_Null_When_Empty bool
}

func (m ClusterStorageInfoEcConfig) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field k
	if m.K != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"k\":")
		bytes, err := swag.WriteJSON(m.K)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.K_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"k\":null")
		first = false
	}

	// handle nullable field m
	if m.M != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"m\":")
		bytes, err := swag.WriteJSON(m.M)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.M_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"m\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this cluster storage info ec config
func (m *ClusterStorageInfoEcConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateK(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateM(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterStorageInfoEcConfig) validateK(formats strfmt.Registry) error {

	if err := validate.Required("k", "body", m.K); err != nil {
		return err
	}

	return nil
}

func (m *ClusterStorageInfoEcConfig) validateM(formats strfmt.Registry) error {

	if err := validate.Required("m", "body", m.M); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cluster storage info ec config based on context it is used
func (m *ClusterStorageInfoEcConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterStorageInfoEcConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterStorageInfoEcConfig) UnmarshalBinary(b []byte) error {
	var res ClusterStorageInfoEcConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
