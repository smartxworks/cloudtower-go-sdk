// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MetricLabel metric label
//
// swagger:model MetricLabel
type MetricLabel struct {

	// typename
	// Enum: [MetricLabel]
	Typename *string `json:"__typename,omitempty"`

	// chunk
	Chunk *string `json:"_chunk,omitempty"`

	// cluster
	Cluster *string `json:"_cluster,omitempty"`

	// device
	Device *string `json:"_device,omitempty"`

	// esxi uuid
	EsxiUUID *string `json:"_esxi_uuid,omitempty"`

	// host
	Host *string `json:"_host,omitempty"`

	// mac
	Mac *string `json:"_mac,omitempty"`

	// network
	Network *string `json:"_network,omitempty"`

	// scvm
	Scvm *string `json:"_scvm,omitempty"`

	// service
	Service *string `json:"_service,omitempty"`

	// to uuid
	ToUUID *string `json:"_to_uuid,omitempty"`

	// vm
	VM *string `json:"_vm,omitempty"`

	// volume
	Volume *string `json:"_volume,omitempty"`

	// witness
	Witness *string `json:"_witness,omitempty"`

	// zone
	Zone *string `json:"_zone,omitempty"`

	// instance
	Instance *string `json:"instance,omitempty"`

	// job
	Job *string `json:"job,omitempty"`

	// metric name
	MetricName *string `json:"metric_name,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// pool
	Pool *string `json:"pool,omitempty"`

	// serial number
	SerialNumber *string `json:"serial_number,omitempty"`

	// to hostname
	ToHostname *string `json:"to_hostname,omitempty"`
}

// Validate validates this metric label
func (m *MetricLabel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTypename(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var metricLabelTypeTypenamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MetricLabel"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metricLabelTypeTypenamePropEnum = append(metricLabelTypeTypenamePropEnum, v)
	}
}

const (

	// MetricLabelTypenameMetricLabel captures enum value "MetricLabel"
	MetricLabelTypenameMetricLabel string = "MetricLabel"
)

// prop value enum
func (m *MetricLabel) validateTypenameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, metricLabelTypeTypenamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MetricLabel) validateTypename(formats strfmt.Registry) error {
	if swag.IsZero(m.Typename) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypenameEnum("__typename", "body", *m.Typename); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this metric label based on context it is used
func (m *MetricLabel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MetricLabel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricLabel) UnmarshalBinary(b []byte) error {
	var res MetricLabel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
