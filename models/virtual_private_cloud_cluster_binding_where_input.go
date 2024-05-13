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

// VirtualPrivateCloudClusterBindingWhereInput virtual private cloud cluster binding where input
//
// swagger:model VirtualPrivateCloudClusterBindingWhereInput
type VirtualPrivateCloudClusterBindingWhereInput struct {

	// a n d
	AND []*VirtualPrivateCloudClusterBindingWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*VirtualPrivateCloudClusterBindingWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*VirtualPrivateCloudClusterBindingWhereInput `json:"OR,omitempty"`

	// cluster
	Cluster *ClusterWhereInput `json:"cluster,omitempty"`

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// entity async status in
	EntityAsyncStatusIn []EntityAsyncStatus `json:"entityAsyncStatus_in,omitempty"`

	// entity async status not
	EntityAsyncStatusNot *EntityAsyncStatus `json:"entityAsyncStatus_not,omitempty"`

	// entity async status not in
	EntityAsyncStatusNotIn []EntityAsyncStatus `json:"entityAsyncStatus_not_in,omitempty"`

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

	// phase
	Phase *EverouteClusterPhase `json:"phase,omitempty"`

	// phase in
	PhaseIn []EverouteClusterPhase `json:"phase_in,omitempty"`

	// phase not
	PhaseNot *EverouteClusterPhase `json:"phase_not,omitempty"`

	// phase not in
	PhaseNotIn []EverouteClusterPhase `json:"phase_not_in,omitempty"`

	// tep ip pool id
	TepIPPoolID *string `json:"tep_ip_pool_id,omitempty"`

	// tep ip pool id contains
	TepIPPoolIDContains *string `json:"tep_ip_pool_id_contains,omitempty"`

	// tep ip pool id ends with
	TepIPPoolIDEndsWith *string `json:"tep_ip_pool_id_ends_with,omitempty"`

	// tep ip pool id gt
	TepIPPoolIDGt *string `json:"tep_ip_pool_id_gt,omitempty"`

	// tep ip pool id gte
	TepIPPoolIDGte *string `json:"tep_ip_pool_id_gte,omitempty"`

	// tep ip pool id in
	TepIPPoolIDIn []string `json:"tep_ip_pool_id_in,omitempty"`

	// tep ip pool id lt
	TepIPPoolIDLt *string `json:"tep_ip_pool_id_lt,omitempty"`

	// tep ip pool id lte
	TepIPPoolIDLte *string `json:"tep_ip_pool_id_lte,omitempty"`

	// tep ip pool id not
	TepIPPoolIDNot *string `json:"tep_ip_pool_id_not,omitempty"`

	// tep ip pool id not contains
	TepIPPoolIDNotContains *string `json:"tep_ip_pool_id_not_contains,omitempty"`

	// tep ip pool id not ends with
	TepIPPoolIDNotEndsWith *string `json:"tep_ip_pool_id_not_ends_with,omitempty"`

	// tep ip pool id not in
	TepIPPoolIDNotIn []string `json:"tep_ip_pool_id_not_in,omitempty"`

	// tep ip pool id not starts with
	TepIPPoolIDNotStartsWith *string `json:"tep_ip_pool_id_not_starts_with,omitempty"`

	// tep ip pool id starts with
	TepIPPoolIDStartsWith *string `json:"tep_ip_pool_id_starts_with,omitempty"`

	// vds
	Vds *VdsWhereInput `json:"vds,omitempty"`

	// vlan id
	VlanID *int32 `json:"vlan_id,omitempty"`

	// vlan id gt
	VlanIDGt *int32 `json:"vlan_id_gt,omitempty"`

	// vlan id gte
	VlanIDGte *int32 `json:"vlan_id_gte,omitempty"`

	// vlan id in
	VlanIDIn []int32 `json:"vlan_id_in,omitempty"`

	// vlan id lt
	VlanIDLt *int32 `json:"vlan_id_lt,omitempty"`

	// vlan id lte
	VlanIDLte *int32 `json:"vlan_id_lte,omitempty"`

	// vlan id not
	VlanIDNot *int32 `json:"vlan_id_not,omitempty"`

	// vlan id not in
	VlanIDNotIn []int32 `json:"vlan_id_not_in,omitempty"`

	// vpc service
	VpcService *VirtualPrivateCloudServiceWhereInput `json:"vpc_service,omitempty"`
}

// Validate validates this virtual private cloud cluster binding where input
func (m *VirtualPrivateCloudClusterBindingWhereInput) Validate(formats strfmt.Registry) error {
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

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhaseIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhaseNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhaseNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpcService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualPrivateCloudClusterBindingWhereInput) validateAND(formats strfmt.Registry) error {
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

func (m *VirtualPrivateCloudClusterBindingWhereInput) validateNOT(formats strfmt.Registry) error {
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

func (m *VirtualPrivateCloudClusterBindingWhereInput) validateOR(formats strfmt.Registry) error {
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

func (m *VirtualPrivateCloudClusterBindingWhereInput) validateCluster(formats strfmt.Registry) error {
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

func (m *VirtualPrivateCloudClusterBindingWhereInput) validateEntityAsyncStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatus) { // not required
		return nil
	}

	if m.EntityAsyncStatus != nil {
		if err := m.EntityAsyncStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudClusterBindingWhereInput) validateEntityAsyncStatusIn(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusIn) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityAsyncStatusIn); i++ {

		if err := m.EntityAsyncStatusIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VirtualPrivateCloudClusterBindingWhereInput) validateEntityAsyncStatusNot(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusNot) { // not required
		return nil
	}

	if m.EntityAsyncStatusNot != nil {
		if err := m.EntityAsyncStatusNot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudClusterBindingWhereInput) validateEntityAsyncStatusNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityAsyncStatusNotIn); i++ {

		if err := m.EntityAsyncStatusNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VirtualPrivateCloudClusterBindingWhereInput) validatePhase(formats strfmt.Registry) error {
	if swag.IsZero(m.Phase) { // not required
		return nil
	}

	if m.Phase != nil {
		if err := m.Phase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phase")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phase")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudClusterBindingWhereInput) validatePhaseIn(formats strfmt.Registry) error {
	if swag.IsZero(m.PhaseIn) { // not required
		return nil
	}

	for i := 0; i < len(m.PhaseIn); i++ {

		if err := m.PhaseIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phase_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phase_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VirtualPrivateCloudClusterBindingWhereInput) validatePhaseNot(formats strfmt.Registry) error {
	if swag.IsZero(m.PhaseNot) { // not required
		return nil
	}

	if m.PhaseNot != nil {
		if err := m.PhaseNot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phase_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phase_not")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudClusterBindingWhereInput) validatePhaseNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.PhaseNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.PhaseNotIn); i++ {

		if err := m.PhaseNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phase_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phase_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VirtualPrivateCloudClusterBindingWhereInput) validateVds(formats strfmt.Registry) error {
	if swag.IsZero(m.Vds) { // not required
		return nil
	}

	if m.Vds != nil {
		if err := m.Vds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vds")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vds")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudClusterBindingWhereInput) validateVpcService(formats strfmt.Registry) error {
	if swag.IsZero(m.VpcService) { // not required
		return nil
	}

	if m.VpcService != nil {
		if err := m.VpcService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpc_service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vpc_service")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this virtual private cloud cluster binding where input based on the context it is used
func (m *VirtualPrivateCloudClusterBindingWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhase(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhaseIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhaseNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhaseNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVpcService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualPrivateCloudClusterBindingWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VirtualPrivateCloudClusterBindingWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VirtualPrivateCloudClusterBindingWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VirtualPrivateCloudClusterBindingWhereInput) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VirtualPrivateCloudClusterBindingWhereInput) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.EntityAsyncStatus != nil {
		if err := m.EntityAsyncStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudClusterBindingWhereInput) contextValidateEntityAsyncStatusIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntityAsyncStatusIn); i++ {

		if err := m.EntityAsyncStatusIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VirtualPrivateCloudClusterBindingWhereInput) contextValidateEntityAsyncStatusNot(ctx context.Context, formats strfmt.Registry) error {

	if m.EntityAsyncStatusNot != nil {
		if err := m.EntityAsyncStatusNot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudClusterBindingWhereInput) contextValidateEntityAsyncStatusNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntityAsyncStatusNotIn); i++ {

		if err := m.EntityAsyncStatusNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VirtualPrivateCloudClusterBindingWhereInput) contextValidatePhase(ctx context.Context, formats strfmt.Registry) error {

	if m.Phase != nil {
		if err := m.Phase.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phase")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phase")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudClusterBindingWhereInput) contextValidatePhaseIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PhaseIn); i++ {

		if err := m.PhaseIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phase_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phase_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VirtualPrivateCloudClusterBindingWhereInput) contextValidatePhaseNot(ctx context.Context, formats strfmt.Registry) error {

	if m.PhaseNot != nil {
		if err := m.PhaseNot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phase_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phase_not")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudClusterBindingWhereInput) contextValidatePhaseNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PhaseNotIn); i++ {

		if err := m.PhaseNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("phase_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("phase_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VirtualPrivateCloudClusterBindingWhereInput) contextValidateVds(ctx context.Context, formats strfmt.Registry) error {

	if m.Vds != nil {
		if err := m.Vds.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vds")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vds")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudClusterBindingWhereInput) contextValidateVpcService(ctx context.Context, formats strfmt.Registry) error {

	if m.VpcService != nil {
		if err := m.VpcService.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpc_service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vpc_service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VirtualPrivateCloudClusterBindingWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualPrivateCloudClusterBindingWhereInput) UnmarshalBinary(b []byte) error {
	var res VirtualPrivateCloudClusterBindingWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}