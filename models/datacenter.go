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

// Datacenter datacenter
//
// swagger:model Datacenter
type Datacenter struct {

	// cluster num
	ClusterNum *int32 `json:"cluster_num,omitempty"`

	// clusters
	Clusters []*NestedCluster `json:"clusters,omitempty"`

	// failure data space
	FailureDataSpace *int64 `json:"failure_data_space,omitempty"`

	// host num
	HostNum *int32 `json:"host_num,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// labels
	Labels []*NestedLabel `json:"labels,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// organization
	// Required: true
	Organization *NestedOrganization `json:"organization"`

	// total cpu hz
	TotalCPUHz *int64 `json:"total_cpu_hz,omitempty"`

	// total data capacity
	TotalDataCapacity *int64 `json:"total_data_capacity,omitempty"`

	// total memory bytes
	TotalMemoryBytes *int64 `json:"total_memory_bytes,omitempty"`

	// used cpu hz
	UsedCPUHz *float64 `json:"used_cpu_hz,omitempty"`

	// used data space
	UsedDataSpace *int64 `json:"used_data_space,omitempty"`

	// used memory bytes
	UsedMemoryBytes *float64 `json:"used_memory_bytes,omitempty"`

	// vm num
	VMNum *int32 `json:"vm_num,omitempty"`

	MarshalOpts *DatacenterMarshalOpts `json:"-"`
}

type DatacenterMarshalOpts struct {
	ClusterNum_Explicit_Null_When_Empty bool

	Clusters_Explicit_Null_When_Empty bool

	FailureDataSpace_Explicit_Null_When_Empty bool

	HostNum_Explicit_Null_When_Empty bool

	ID_Explicit_Null_When_Empty bool

	Labels_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	Organization_Explicit_Null_When_Empty bool

	TotalCPUHz_Explicit_Null_When_Empty bool

	TotalDataCapacity_Explicit_Null_When_Empty bool

	TotalMemoryBytes_Explicit_Null_When_Empty bool

	UsedCPUHz_Explicit_Null_When_Empty bool

	UsedDataSpace_Explicit_Null_When_Empty bool

	UsedMemoryBytes_Explicit_Null_When_Empty bool

	VMNum_Explicit_Null_When_Empty bool
}

func (m Datacenter) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field cluster_num
	if m.ClusterNum != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster_num\":")
		bytes, err := swag.WriteJSON(m.ClusterNum)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ClusterNum_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"cluster_num\":null")
		first = false
	}

	// handle non nullable field clusters with omitempty
	if !swag.IsZero(m.Clusters) {
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
	}

	// handle nullable field failure_data_space
	if m.FailureDataSpace != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"failure_data_space\":")
		bytes, err := swag.WriteJSON(m.FailureDataSpace)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.FailureDataSpace_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"failure_data_space\":null")
		first = false
	}

	// handle nullable field host_num
	if m.HostNum != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"host_num\":")
		bytes, err := swag.WriteJSON(m.HostNum)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.HostNum_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"host_num\":null")
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

	// handle non nullable field labels with omitempty
	if !swag.IsZero(m.Labels) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"labels\":")
		bytes, err := swag.WriteJSON(m.Labels)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
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

	// handle nullable field organization
	if m.Organization != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"organization\":")
		bytes, err := swag.WriteJSON(m.Organization)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Organization_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"organization\":null")
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

	// handle nullable field total_data_capacity
	if m.TotalDataCapacity != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"total_data_capacity\":")
		bytes, err := swag.WriteJSON(m.TotalDataCapacity)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.TotalDataCapacity_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"total_data_capacity\":null")
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

	// handle nullable field used_cpu_hz
	if m.UsedCPUHz != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"used_cpu_hz\":")
		bytes, err := swag.WriteJSON(m.UsedCPUHz)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.UsedCPUHz_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"used_cpu_hz\":null")
		first = false
	}

	// handle nullable field used_data_space
	if m.UsedDataSpace != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"used_data_space\":")
		bytes, err := swag.WriteJSON(m.UsedDataSpace)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.UsedDataSpace_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"used_data_space\":null")
		first = false
	}

	// handle nullable field used_memory_bytes
	if m.UsedMemoryBytes != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"used_memory_bytes\":")
		bytes, err := swag.WriteJSON(m.UsedMemoryBytes)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.UsedMemoryBytes_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"used_memory_bytes\":null")
		first = false
	}

	// handle nullable field vm_num
	if m.VMNum != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_num\":")
		bytes, err := swag.WriteJSON(m.VMNum)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VMNum_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vm_num\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this datacenter
func (m *Datacenter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Datacenter) validateClusters(formats strfmt.Registry) error {
	if swag.IsZero(m.Clusters) { // not required
		return nil
	}

	for i := 0; i < len(m.Clusters); i++ {
		if swag.IsZero(m.Clusters[i]) { // not required
			continue
		}

		if m.Clusters[i] != nil {
			if err := m.Clusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Datacenter) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Datacenter) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	for i := 0; i < len(m.Labels); i++ {
		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if m.Labels[i] != nil {
			if err := m.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Datacenter) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Datacenter) validateOrganization(formats strfmt.Registry) error {

	if err := validate.Required("organization", "body", m.Organization); err != nil {
		return err
	}

	if m.Organization != nil {
		if err := m.Organization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("organization")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this datacenter based on the context it is used
func (m *Datacenter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrganization(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Datacenter) contextValidateClusters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Clusters); i++ {

		if m.Clusters[i] != nil {
			if err := m.Clusters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Datacenter) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Labels); i++ {

		if m.Labels[i] != nil {
			if err := m.Labels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Datacenter) contextValidateOrganization(ctx context.Context, formats strfmt.Registry) error {

	if m.Organization != nil {
		if err := m.Organization.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("organization")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Datacenter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Datacenter) UnmarshalBinary(b []byte) error {
	var res Datacenter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
