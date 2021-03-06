// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WitnessWhereInput witness where input
//
// swagger:model WitnessWhereInput
type WitnessWhereInput struct {

	// a n d
	AND []*WitnessWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*WitnessWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*WitnessWhereInput `json:"OR,omitempty"`

	// cluster
	Cluster *ClusterWhereInput `json:"cluster,omitempty"`

	// cpu hz per core
	CPUHzPerCore *int64 `json:"cpu_hz_per_core,omitempty"`

	// cpu hz per core gt
	CPUHzPerCoreGt *int64 `json:"cpu_hz_per_core_gt,omitempty"`

	// cpu hz per core gte
	CPUHzPerCoreGte *int64 `json:"cpu_hz_per_core_gte,omitempty"`

	// cpu hz per core in
	CPUHzPerCoreIn []int64 `json:"cpu_hz_per_core_in,omitempty"`

	// cpu hz per core lt
	CPUHzPerCoreLt *int64 `json:"cpu_hz_per_core_lt,omitempty"`

	// cpu hz per core lte
	CPUHzPerCoreLte *int64 `json:"cpu_hz_per_core_lte,omitempty"`

	// cpu hz per core not
	CPUHzPerCoreNot *int64 `json:"cpu_hz_per_core_not,omitempty"`

	// cpu hz per core not in
	CPUHzPerCoreNotIn []int64 `json:"cpu_hz_per_core_not_in,omitempty"`

	// data ip
	DataIP *string `json:"data_ip,omitempty"`

	// data ip contains
	DataIPContains *string `json:"data_ip_contains,omitempty"`

	// data ip ends with
	DataIPEndsWith *string `json:"data_ip_ends_with,omitempty"`

	// data ip gt
	DataIPGt *string `json:"data_ip_gt,omitempty"`

	// data ip gte
	DataIPGte *string `json:"data_ip_gte,omitempty"`

	// data ip in
	DataIPIn []string `json:"data_ip_in,omitempty"`

	// data ip lt
	DataIPLt *string `json:"data_ip_lt,omitempty"`

	// data ip lte
	DataIPLte *string `json:"data_ip_lte,omitempty"`

	// data ip not
	DataIPNot *string `json:"data_ip_not,omitempty"`

	// data ip not contains
	DataIPNotContains *string `json:"data_ip_not_contains,omitempty"`

	// data ip not ends with
	DataIPNotEndsWith *string `json:"data_ip_not_ends_with,omitempty"`

	// data ip not in
	DataIPNotIn []string `json:"data_ip_not_in,omitempty"`

	// data ip not starts with
	DataIPNotStartsWith *string `json:"data_ip_not_starts_with,omitempty"`

	// data ip starts with
	DataIPStartsWith *string `json:"data_ip_starts_with,omitempty"`

	// id
	ID *string `json:"id,omitempty"`

	// id contains
	IDContains *string `json:"id_contains,omitempty"`

	// id ends with
	IDEndsWith *string `json:"id_ends_with,omitempty"`

	// id gt
	IDGt *string `json:"id_gt,omitempty"`

	// id gte
	IDGte *string `json:"id_gte,omitempty"`

	// id in
	IDIn []string `json:"id_in,omitempty"`

	// id lt
	IDLt *string `json:"id_lt,omitempty"`

	// id lte
	IDLte *string `json:"id_lte,omitempty"`

	// id not
	IDNot *string `json:"id_not,omitempty"`

	// id not contains
	IDNotContains *string `json:"id_not_contains,omitempty"`

	// id not ends with
	IDNotEndsWith *string `json:"id_not_ends_with,omitempty"`

	// id not in
	IDNotIn []string `json:"id_not_in,omitempty"`

	// id not starts with
	IDNotStartsWith *string `json:"id_not_starts_with,omitempty"`

	// id starts with
	IDStartsWith *string `json:"id_starts_with,omitempty"`

	// local id
	LocalID *string `json:"local_id,omitempty"`

	// local id contains
	LocalIDContains *string `json:"local_id_contains,omitempty"`

	// local id ends with
	LocalIDEndsWith *string `json:"local_id_ends_with,omitempty"`

	// local id gt
	LocalIDGt *string `json:"local_id_gt,omitempty"`

	// local id gte
	LocalIDGte *string `json:"local_id_gte,omitempty"`

	// local id in
	LocalIDIn []string `json:"local_id_in,omitempty"`

	// local id lt
	LocalIDLt *string `json:"local_id_lt,omitempty"`

	// local id lte
	LocalIDLte *string `json:"local_id_lte,omitempty"`

	// local id not
	LocalIDNot *string `json:"local_id_not,omitempty"`

	// local id not contains
	LocalIDNotContains *string `json:"local_id_not_contains,omitempty"`

	// local id not ends with
	LocalIDNotEndsWith *string `json:"local_id_not_ends_with,omitempty"`

	// local id not in
	LocalIDNotIn []string `json:"local_id_not_in,omitempty"`

	// local id not starts with
	LocalIDNotStartsWith *string `json:"local_id_not_starts_with,omitempty"`

	// local id starts with
	LocalIDStartsWith *string `json:"local_id_starts_with,omitempty"`

	// management ip
	ManagementIP *string `json:"management_ip,omitempty"`

	// management ip contains
	ManagementIPContains *string `json:"management_ip_contains,omitempty"`

	// management ip ends with
	ManagementIPEndsWith *string `json:"management_ip_ends_with,omitempty"`

	// management ip gt
	ManagementIPGt *string `json:"management_ip_gt,omitempty"`

	// management ip gte
	ManagementIPGte *string `json:"management_ip_gte,omitempty"`

	// management ip in
	ManagementIPIn []string `json:"management_ip_in,omitempty"`

	// management ip lt
	ManagementIPLt *string `json:"management_ip_lt,omitempty"`

	// management ip lte
	ManagementIPLte *string `json:"management_ip_lte,omitempty"`

	// management ip not
	ManagementIPNot *string `json:"management_ip_not,omitempty"`

	// management ip not contains
	ManagementIPNotContains *string `json:"management_ip_not_contains,omitempty"`

	// management ip not ends with
	ManagementIPNotEndsWith *string `json:"management_ip_not_ends_with,omitempty"`

	// management ip not in
	ManagementIPNotIn []string `json:"management_ip_not_in,omitempty"`

	// management ip not starts with
	ManagementIPNotStartsWith *string `json:"management_ip_not_starts_with,omitempty"`

	// management ip starts with
	ManagementIPStartsWith *string `json:"management_ip_starts_with,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// name contains
	NameContains *string `json:"name_contains,omitempty"`

	// name ends with
	NameEndsWith *string `json:"name_ends_with,omitempty"`

	// name gt
	NameGt *string `json:"name_gt,omitempty"`

	// name gte
	NameGte *string `json:"name_gte,omitempty"`

	// name in
	NameIn []string `json:"name_in,omitempty"`

	// name lt
	NameLt *string `json:"name_lt,omitempty"`

	// name lte
	NameLte *string `json:"name_lte,omitempty"`

	// name not
	NameNot *string `json:"name_not,omitempty"`

	// name not contains
	NameNotContains *string `json:"name_not_contains,omitempty"`

	// name not ends with
	NameNotEndsWith *string `json:"name_not_ends_with,omitempty"`

	// name not in
	NameNotIn []string `json:"name_not_in,omitempty"`

	// name not starts with
	NameNotStartsWith *string `json:"name_not_starts_with,omitempty"`

	// name starts with
	NameStartsWith *string `json:"name_starts_with,omitempty"`

	// system data capacity
	SystemDataCapacity *int64 `json:"system_data_capacity,omitempty"`

	// system data capacity gt
	SystemDataCapacityGt *int64 `json:"system_data_capacity_gt,omitempty"`

	// system data capacity gte
	SystemDataCapacityGte *int64 `json:"system_data_capacity_gte,omitempty"`

	// system data capacity in
	SystemDataCapacityIn []int64 `json:"system_data_capacity_in,omitempty"`

	// system data capacity lt
	SystemDataCapacityLt *int64 `json:"system_data_capacity_lt,omitempty"`

	// system data capacity lte
	SystemDataCapacityLte *int64 `json:"system_data_capacity_lte,omitempty"`

	// system data capacity not
	SystemDataCapacityNot *int64 `json:"system_data_capacity_not,omitempty"`

	// system data capacity not in
	SystemDataCapacityNotIn []int64 `json:"system_data_capacity_not_in,omitempty"`

	// system used data space
	SystemUsedDataSpace *int64 `json:"system_used_data_space,omitempty"`

	// system used data space gt
	SystemUsedDataSpaceGt *int64 `json:"system_used_data_space_gt,omitempty"`

	// system used data space gte
	SystemUsedDataSpaceGte *int64 `json:"system_used_data_space_gte,omitempty"`

	// system used data space in
	SystemUsedDataSpaceIn []int64 `json:"system_used_data_space_in,omitempty"`

	// system used data space lt
	SystemUsedDataSpaceLt *int64 `json:"system_used_data_space_lt,omitempty"`

	// system used data space lte
	SystemUsedDataSpaceLte *int64 `json:"system_used_data_space_lte,omitempty"`

	// system used data space not
	SystemUsedDataSpaceNot *int64 `json:"system_used_data_space_not,omitempty"`

	// system used data space not in
	SystemUsedDataSpaceNotIn []int64 `json:"system_used_data_space_not_in,omitempty"`

	// total cpu cores
	TotalCPUCores *int32 `json:"total_cpu_cores,omitempty"`

	// total cpu cores gt
	TotalCPUCoresGt *int32 `json:"total_cpu_cores_gt,omitempty"`

	// total cpu cores gte
	TotalCPUCoresGte *int32 `json:"total_cpu_cores_gte,omitempty"`

	// total cpu cores in
	TotalCPUCoresIn []int32 `json:"total_cpu_cores_in,omitempty"`

	// total cpu cores lt
	TotalCPUCoresLt *int32 `json:"total_cpu_cores_lt,omitempty"`

	// total cpu cores lte
	TotalCPUCoresLte *int32 `json:"total_cpu_cores_lte,omitempty"`

	// total cpu cores not
	TotalCPUCoresNot *int32 `json:"total_cpu_cores_not,omitempty"`

	// total cpu cores not in
	TotalCPUCoresNotIn []int32 `json:"total_cpu_cores_not_in,omitempty"`

	// total cpu hz
	TotalCPUHz *int64 `json:"total_cpu_hz,omitempty"`

	// total cpu hz gt
	TotalCPUHzGt *int64 `json:"total_cpu_hz_gt,omitempty"`

	// total cpu hz gte
	TotalCPUHzGte *int64 `json:"total_cpu_hz_gte,omitempty"`

	// total cpu hz in
	TotalCPUHzIn []int64 `json:"total_cpu_hz_in,omitempty"`

	// total cpu hz lt
	TotalCPUHzLt *int64 `json:"total_cpu_hz_lt,omitempty"`

	// total cpu hz lte
	TotalCPUHzLte *int64 `json:"total_cpu_hz_lte,omitempty"`

	// total cpu hz not
	TotalCPUHzNot *int64 `json:"total_cpu_hz_not,omitempty"`

	// total cpu hz not in
	TotalCPUHzNotIn []int64 `json:"total_cpu_hz_not_in,omitempty"`

	// total memory bytes
	TotalMemoryBytes *int64 `json:"total_memory_bytes,omitempty"`

	// total memory bytes gt
	TotalMemoryBytesGt *int64 `json:"total_memory_bytes_gt,omitempty"`

	// total memory bytes gte
	TotalMemoryBytesGte *int64 `json:"total_memory_bytes_gte,omitempty"`

	// total memory bytes in
	TotalMemoryBytesIn []int64 `json:"total_memory_bytes_in,omitempty"`

	// total memory bytes lt
	TotalMemoryBytesLt *int64 `json:"total_memory_bytes_lt,omitempty"`

	// total memory bytes lte
	TotalMemoryBytesLte *int64 `json:"total_memory_bytes_lte,omitempty"`

	// total memory bytes not
	TotalMemoryBytesNot *int64 `json:"total_memory_bytes_not,omitempty"`

	// total memory bytes not in
	TotalMemoryBytesNotIn []int64 `json:"total_memory_bytes_not_in,omitempty"`
}

// Validate validates this witness where input
func (m *WitnessWhereInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAND(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNOT(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOR(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WitnessWhereInput) validateAND(formats strfmt.Registry) error {
	if swag.IsZero(m.AND) { // not required
		return nil
	}

	for i := 0; i < len(m.AND); i++ {
		if swag.IsZero(m.AND[i]) { // not required
			continue
		}

		if m.AND[i] != nil {
			if err := m.AND[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WitnessWhereInput) validateNOT(formats strfmt.Registry) error {
	if swag.IsZero(m.NOT) { // not required
		return nil
	}

	for i := 0; i < len(m.NOT); i++ {
		if swag.IsZero(m.NOT[i]) { // not required
			continue
		}

		if m.NOT[i] != nil {
			if err := m.NOT[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WitnessWhereInput) validateOR(formats strfmt.Registry) error {
	if swag.IsZero(m.OR) { // not required
		return nil
	}

	for i := 0; i < len(m.OR); i++ {
		if swag.IsZero(m.OR[i]) { // not required
			continue
		}

		if m.OR[i] != nil {
			if err := m.OR[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WitnessWhereInput) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.Cluster) { // not required
		return nil
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

// ContextValidate validate this witness where input based on the context it is used
func (m *WitnessWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAND(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNOT(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOR(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WitnessWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AND); i++ {

		if m.AND[i] != nil {
			if err := m.AND[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WitnessWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NOT); i++ {

		if m.NOT[i] != nil {
			if err := m.NOT[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WitnessWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OR); i++ {

		if m.OR[i] != nil {
			if err := m.OR[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WitnessWhereInput) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

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
func (m *WitnessWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WitnessWhereInput) UnmarshalBinary(b []byte) error {
	var res WitnessWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
