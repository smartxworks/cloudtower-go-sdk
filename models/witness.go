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

// Witness witness
//
// swagger:model Witness
type Witness struct {

	// cluster
	// Required: true
	Cluster *NestedCluster `json:"cluster"`

	// cpu hz per core
	// Required: true
	CPUHzPerCore *int64 `json:"cpu_hz_per_core"`

	// data ip
	// Required: true
	DataIP *string `json:"data_ip"`

	// id
	// Required: true
	ID *string `json:"id"`

	// local id
	LocalID *string `json:"local_id,omitempty"`

	// management ip
	// Required: true
	ManagementIP *string `json:"management_ip"`

	// name
	// Required: true
	Name *string `json:"name"`

	// system data capacity
	// Required: true
	SystemDataCapacity *int64 `json:"system_data_capacity"`

	// system used data space
	// Required: true
	SystemUsedDataSpace *int64 `json:"system_used_data_space"`

	// total cpu cores
	// Required: true
	TotalCPUCores *int32 `json:"total_cpu_cores"`

	// total cpu hz
	// Required: true
	TotalCPUHz *int64 `json:"total_cpu_hz"`

	// total memory bytes
	// Required: true
	TotalMemoryBytes *int64 `json:"total_memory_bytes"`

	MarshalOpts *WitnessMarshalOpts `json:"-"`
}

type WitnessMarshalOpts struct {
	Cluster_Explicit_Null_When_Empty bool

	CPUHzPerCore_Explicit_Null_When_Empty bool

	DataIP_Explicit_Null_When_Empty bool

	ID_Explicit_Null_When_Empty bool

	LocalID_Explicit_Null_When_Empty bool

	ManagementIP_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	SystemDataCapacity_Explicit_Null_When_Empty bool

	SystemUsedDataSpace_Explicit_Null_When_Empty bool

	TotalCPUCores_Explicit_Null_When_Empty bool

	TotalCPUHz_Explicit_Null_When_Empty bool

	TotalMemoryBytes_Explicit_Null_When_Empty bool
}

func (m Witness) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field cluster
	if m.Cluster != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster\":")
		bytes, err := swag.WriteJSON(m.Cluster)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Cluster_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster\":null")
		first = false
	}

	// handle nullable field cpu_hz_per_core
	if m.CPUHzPerCore != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cpu_hz_per_core\":")
		bytes, err := swag.WriteJSON(m.CPUHzPerCore)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.CPUHzPerCore_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cpu_hz_per_core\":null")
		first = false
	}

	// handle nullable field data_ip
	if m.DataIP != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"data_ip\":")
		bytes, err := swag.WriteJSON(m.DataIP)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.DataIP_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"data_ip\":null")
		first = false
	}

	// handle nullable field id
	if m.ID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id\":")
		bytes, err := swag.WriteJSON(m.ID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id\":null")
		first = false
	}

	// handle nullable field local_id
	if m.LocalID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id\":")
		bytes, err := swag.WriteJSON(m.LocalID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.LocalID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id\":null")
		first = false
	}

	// handle nullable field management_ip
	if m.ManagementIP != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"management_ip\":")
		bytes, err := swag.WriteJSON(m.ManagementIP)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ManagementIP_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"management_ip\":null")
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

	// handle nullable field system_data_capacity
	if m.SystemDataCapacity != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"system_data_capacity\":")
		bytes, err := swag.WriteJSON(m.SystemDataCapacity)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.SystemDataCapacity_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"system_data_capacity\":null")
		first = false
	}

	// handle nullable field system_used_data_space
	if m.SystemUsedDataSpace != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"system_used_data_space\":")
		bytes, err := swag.WriteJSON(m.SystemUsedDataSpace)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.SystemUsedDataSpace_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"system_used_data_space\":null")
		first = false
	}

	// handle nullable field total_cpu_cores
	if m.TotalCPUCores != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"total_cpu_cores\":")
		bytes, err := swag.WriteJSON(m.TotalCPUCores)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.TotalCPUCores_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"total_cpu_cores\":null")
		first = false
	}

	// handle nullable field total_cpu_hz
	if m.TotalCPUHz != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"total_cpu_hz\":")
		bytes, err := swag.WriteJSON(m.TotalCPUHz)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.TotalCPUHz_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"total_cpu_hz\":null")
		first = false
	}

	// handle nullable field total_memory_bytes
	if m.TotalMemoryBytes != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"total_memory_bytes\":")
		bytes, err := swag.WriteJSON(m.TotalMemoryBytes)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.TotalMemoryBytes_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"total_memory_bytes\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this witness
func (m *Witness) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPUHzPerCore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManagementIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystemDataCapacity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystemUsedDataSpace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalCPUCores(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalCPUHz(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalMemoryBytes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Witness) validateCluster(formats strfmt.Registry) error {

	if err := validate.Required("cluster", "body", m.Cluster); err != nil {
		return err
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *Witness) validateCPUHzPerCore(formats strfmt.Registry) error {

	if err := validate.Required("cpu_hz_per_core", "body", m.CPUHzPerCore); err != nil {
		return err
	}

	return nil
}

func (m *Witness) validateDataIP(formats strfmt.Registry) error {

	if err := validate.Required("data_ip", "body", m.DataIP); err != nil {
		return err
	}

	return nil
}

func (m *Witness) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Witness) validateManagementIP(formats strfmt.Registry) error {

	if err := validate.Required("management_ip", "body", m.ManagementIP); err != nil {
		return err
	}

	return nil
}

func (m *Witness) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Witness) validateSystemDataCapacity(formats strfmt.Registry) error {

	if err := validate.Required("system_data_capacity", "body", m.SystemDataCapacity); err != nil {
		return err
	}

	return nil
}

func (m *Witness) validateSystemUsedDataSpace(formats strfmt.Registry) error {

	if err := validate.Required("system_used_data_space", "body", m.SystemUsedDataSpace); err != nil {
		return err
	}

	return nil
}

func (m *Witness) validateTotalCPUCores(formats strfmt.Registry) error {

	if err := validate.Required("total_cpu_cores", "body", m.TotalCPUCores); err != nil {
		return err
	}

	return nil
}

func (m *Witness) validateTotalCPUHz(formats strfmt.Registry) error {

	if err := validate.Required("total_cpu_hz", "body", m.TotalCPUHz); err != nil {
		return err
	}

	return nil
}

func (m *Witness) validateTotalMemoryBytes(formats strfmt.Registry) error {

	if err := validate.Required("total_memory_bytes", "body", m.TotalMemoryBytes); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this witness based on the context it is used
func (m *Witness) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Witness) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {
		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Witness) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Witness) UnmarshalBinary(b []byte) error {
	var res Witness
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
