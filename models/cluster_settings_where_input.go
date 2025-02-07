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
)

// ClusterSettingsWhereInput cluster settings where input
//
// swagger:model ClusterSettingsWhereInput
type ClusterSettingsWhereInput struct {

	// a n d
	AND []*ClusterSettingsWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*ClusterSettingsWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*ClusterSettingsWhereInput `json:"OR,omitempty"`

	// cluster
	Cluster *ClusterWhereInput `json:"cluster,omitempty"`

	// default ha
	DefaultHa *bool `json:"default_ha,omitempty"`

	// default ha not
	DefaultHaNot *bool `json:"default_ha_not,omitempty"`

	// default storage policy
	DefaultStoragePolicy *VMVolumeElfStoragePolicyType `json:"default_storage_policy,omitempty"`

	// default storage policy in
	DefaultStoragePolicyIn []VMVolumeElfStoragePolicyType `json:"default_storage_policy_in,omitempty"`

	// default storage policy not
	DefaultStoragePolicyNot *VMVolumeElfStoragePolicyType `json:"default_storage_policy_not,omitempty"`

	// default storage policy not in
	DefaultStoragePolicyNotIn []VMVolumeElfStoragePolicyType `json:"default_storage_policy_not_in,omitempty"`

	// enabled iscsi
	EnabledIscsi *bool `json:"enabled_iscsi,omitempty"`

	// enabled iscsi not
	EnabledIscsiNot *bool `json:"enabled_iscsi_not,omitempty"`

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

	MarshalOpts *ClusterSettingsWhereInputMarshalOpts `json:"-"`
}

type ClusterSettingsWhereInputMarshalOpts struct {
	AND_Explicit_Null_When_Empty bool

	NOT_Explicit_Null_When_Empty bool

	OR_Explicit_Null_When_Empty bool

	Cluster_Explicit_Null_When_Empty bool

	DefaultHa_Explicit_Null_When_Empty bool

	DefaultHaNot_Explicit_Null_When_Empty bool

	DefaultStoragePolicy_Explicit_Null_When_Empty bool

	DefaultStoragePolicyIn_Explicit_Null_When_Empty bool

	DefaultStoragePolicyNot_Explicit_Null_When_Empty bool

	DefaultStoragePolicyNotIn_Explicit_Null_When_Empty bool

	EnabledIscsi_Explicit_Null_When_Empty bool

	EnabledIscsiNot_Explicit_Null_When_Empty bool

	ID_Explicit_Null_When_Empty bool

	IDContains_Explicit_Null_When_Empty bool

	IDEndsWith_Explicit_Null_When_Empty bool

	IDGt_Explicit_Null_When_Empty bool

	IDGte_Explicit_Null_When_Empty bool

	IDIn_Explicit_Null_When_Empty bool

	IDLt_Explicit_Null_When_Empty bool

	IDLte_Explicit_Null_When_Empty bool

	IDNot_Explicit_Null_When_Empty bool

	IDNotContains_Explicit_Null_When_Empty bool

	IDNotEndsWith_Explicit_Null_When_Empty bool

	IDNotIn_Explicit_Null_When_Empty bool

	IDNotStartsWith_Explicit_Null_When_Empty bool

	IDStartsWith_Explicit_Null_When_Empty bool
}

func (m ClusterSettingsWhereInput) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle non nullable field AND with omitempty
	if !swag.IsZero(m.AND) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"AND\":")
		bytes, err := swag.WriteJSON(m.AND)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle non nullable field NOT with omitempty
	if !swag.IsZero(m.NOT) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"NOT\":")
		bytes, err := swag.WriteJSON(m.NOT)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle non nullable field OR with omitempty
	if !swag.IsZero(m.OR) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"OR\":")
		bytes, err := swag.WriteJSON(m.OR)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

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

	// handle nullable field default_ha
	if m.DefaultHa != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"default_ha\":")
		bytes, err := swag.WriteJSON(m.DefaultHa)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.DefaultHa_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"default_ha\":null")
		first = false
	}

	// handle nullable field default_ha_not
	if m.DefaultHaNot != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"default_ha_not\":")
		bytes, err := swag.WriteJSON(m.DefaultHaNot)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.DefaultHaNot_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"default_ha_not\":null")
		first = false
	}

	// handle nullable field default_storage_policy
	if m.DefaultStoragePolicy != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"default_storage_policy\":")
		bytes, err := swag.WriteJSON(m.DefaultStoragePolicy)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.DefaultStoragePolicy_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"default_storage_policy\":null")
		first = false
	}

	// handle non nullable field default_storage_policy_in with omitempty
	if !swag.IsZero(m.DefaultStoragePolicyIn) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"default_storage_policy_in\":")
		bytes, err := swag.WriteJSON(m.DefaultStoragePolicyIn)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field default_storage_policy_not
	if m.DefaultStoragePolicyNot != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"default_storage_policy_not\":")
		bytes, err := swag.WriteJSON(m.DefaultStoragePolicyNot)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.DefaultStoragePolicyNot_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"default_storage_policy_not\":null")
		first = false
	}

	// handle non nullable field default_storage_policy_not_in with omitempty
	if !swag.IsZero(m.DefaultStoragePolicyNotIn) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"default_storage_policy_not_in\":")
		bytes, err := swag.WriteJSON(m.DefaultStoragePolicyNotIn)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field enabled_iscsi
	if m.EnabledIscsi != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"enabled_iscsi\":")
		bytes, err := swag.WriteJSON(m.EnabledIscsi)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.EnabledIscsi_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"enabled_iscsi\":null")
		first = false
	}

	// handle nullable field enabled_iscsi_not
	if m.EnabledIscsiNot != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"enabled_iscsi_not\":")
		bytes, err := swag.WriteJSON(m.EnabledIscsiNot)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.EnabledIscsiNot_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"enabled_iscsi_not\":null")
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

	// handle nullable field id_contains
	if m.IDContains != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_contains\":")
		bytes, err := swag.WriteJSON(m.IDContains)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDContains_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_contains\":null")
		first = false
	}

	// handle nullable field id_ends_with
	if m.IDEndsWith != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_ends_with\":")
		bytes, err := swag.WriteJSON(m.IDEndsWith)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDEndsWith_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_ends_with\":null")
		first = false
	}

	// handle nullable field id_gt
	if m.IDGt != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_gt\":")
		bytes, err := swag.WriteJSON(m.IDGt)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDGt_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_gt\":null")
		first = false
	}

	// handle nullable field id_gte
	if m.IDGte != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_gte\":")
		bytes, err := swag.WriteJSON(m.IDGte)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDGte_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_gte\":null")
		first = false
	}

	// handle non nullable field id_in with omitempty
	if !swag.IsZero(m.IDIn) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_in\":")
		bytes, err := swag.WriteJSON(m.IDIn)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field id_lt
	if m.IDLt != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_lt\":")
		bytes, err := swag.WriteJSON(m.IDLt)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDLt_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_lt\":null")
		first = false
	}

	// handle nullable field id_lte
	if m.IDLte != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_lte\":")
		bytes, err := swag.WriteJSON(m.IDLte)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDLte_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_lte\":null")
		first = false
	}

	// handle nullable field id_not
	if m.IDNot != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_not\":")
		bytes, err := swag.WriteJSON(m.IDNot)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDNot_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_not\":null")
		first = false
	}

	// handle nullable field id_not_contains
	if m.IDNotContains != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_not_contains\":")
		bytes, err := swag.WriteJSON(m.IDNotContains)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDNotContains_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_not_contains\":null")
		first = false
	}

	// handle nullable field id_not_ends_with
	if m.IDNotEndsWith != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_not_ends_with\":")
		bytes, err := swag.WriteJSON(m.IDNotEndsWith)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDNotEndsWith_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_not_ends_with\":null")
		first = false
	}

	// handle non nullable field id_not_in with omitempty
	if !swag.IsZero(m.IDNotIn) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_not_in\":")
		bytes, err := swag.WriteJSON(m.IDNotIn)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field id_not_starts_with
	if m.IDNotStartsWith != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_not_starts_with\":")
		bytes, err := swag.WriteJSON(m.IDNotStartsWith)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDNotStartsWith_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_not_starts_with\":null")
		first = false
	}

	// handle nullable field id_starts_with
	if m.IDStartsWith != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_starts_with\":")
		bytes, err := swag.WriteJSON(m.IDStartsWith)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDStartsWith_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_starts_with\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this cluster settings where input
func (m *ClusterSettingsWhereInput) Validate(formats strfmt.Registry) error {
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

	if err := m.validateDefaultStoragePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultStoragePolicyIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultStoragePolicyNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultStoragePolicyNotIn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSettingsWhereInput) validateAND(formats strfmt.Registry) error {
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

func (m *ClusterSettingsWhereInput) validateNOT(formats strfmt.Registry) error {
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

func (m *ClusterSettingsWhereInput) validateOR(formats strfmt.Registry) error {
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

func (m *ClusterSettingsWhereInput) validateCluster(formats strfmt.Registry) error {
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

func (m *ClusterSettingsWhereInput) validateDefaultStoragePolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultStoragePolicy) { // not required
		return nil
	}

	if m.DefaultStoragePolicy != nil {
		if err := m.DefaultStoragePolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_storage_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("default_storage_policy")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSettingsWhereInput) validateDefaultStoragePolicyIn(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultStoragePolicyIn) { // not required
		return nil
	}

	for i := 0; i < len(m.DefaultStoragePolicyIn); i++ {

		if err := m.DefaultStoragePolicyIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_storage_policy_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("default_storage_policy_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ClusterSettingsWhereInput) validateDefaultStoragePolicyNot(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultStoragePolicyNot) { // not required
		return nil
	}

	if m.DefaultStoragePolicyNot != nil {
		if err := m.DefaultStoragePolicyNot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_storage_policy_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("default_storage_policy_not")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSettingsWhereInput) validateDefaultStoragePolicyNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultStoragePolicyNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.DefaultStoragePolicyNotIn); i++ {

		if err := m.DefaultStoragePolicyNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_storage_policy_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("default_storage_policy_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this cluster settings where input based on the context it is used
func (m *ClusterSettingsWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateDefaultStoragePolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDefaultStoragePolicyIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDefaultStoragePolicyNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDefaultStoragePolicyNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSettingsWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ClusterSettingsWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ClusterSettingsWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ClusterSettingsWhereInput) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ClusterSettingsWhereInput) contextValidateDefaultStoragePolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.DefaultStoragePolicy != nil {
		if err := m.DefaultStoragePolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_storage_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("default_storage_policy")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSettingsWhereInput) contextValidateDefaultStoragePolicyIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DefaultStoragePolicyIn); i++ {

		if err := m.DefaultStoragePolicyIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_storage_policy_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("default_storage_policy_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ClusterSettingsWhereInput) contextValidateDefaultStoragePolicyNot(ctx context.Context, formats strfmt.Registry) error {

	if m.DefaultStoragePolicyNot != nil {
		if err := m.DefaultStoragePolicyNot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_storage_policy_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("default_storage_policy_not")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSettingsWhereInput) contextValidateDefaultStoragePolicyNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DefaultStoragePolicyNotIn); i++ {

		if err := m.DefaultStoragePolicyNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_storage_policy_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("default_storage_policy_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterSettingsWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterSettingsWhereInput) UnmarshalBinary(b []byte) error {
	var res ClusterSettingsWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
