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

// LoadBalancerServiceWhereInput load balancer service where input
//
// swagger:model LoadBalancerServiceWhereInput
type LoadBalancerServiceWhereInput struct {

	// a n d
	AND []*LoadBalancerServiceWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*LoadBalancerServiceWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*LoadBalancerServiceWhereInput `json:"OR,omitempty"`

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// entity async status in
	EntityAsyncStatusIn []EntityAsyncStatus `json:"entityAsyncStatus_in,omitempty"`

	// entity async status not
	EntityAsyncStatusNot *EntityAsyncStatus `json:"entityAsyncStatus_not,omitempty"`

	// entity async status not in
	EntityAsyncStatusNotIn []EntityAsyncStatus `json:"entityAsyncStatus_not_in,omitempty"`

	// everoute cluster
	EverouteCluster *EverouteClusterWhereInput `json:"everoute_cluster,omitempty"`

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

	// vm instances every
	VMInstancesEvery *VMWhereInput `json:"vm_instances_every,omitempty"`

	// vm instances none
	VMInstancesNone *VMWhereInput `json:"vm_instances_none,omitempty"`

	// vm instances some
	VMInstancesSome *VMWhereInput `json:"vm_instances_some,omitempty"`

	// vnet bonds every
	VnetBondsEvery *VnetBondWhereInput `json:"vnet_bonds_every,omitempty"`

	// vnet bonds none
	VnetBondsNone *VnetBondWhereInput `json:"vnet_bonds_none,omitempty"`

	// vnet bonds some
	VnetBondsSome *VnetBondWhereInput `json:"vnet_bonds_some,omitempty"`
}

// Validate validates this load balancer service where input
func (m *LoadBalancerServiceWhereInput) Validate(formats strfmt.Registry) error {
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

	if err := m.validateEverouteCluster(formats); err != nil {
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

	if err := m.validateVMInstancesEvery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMInstancesNone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMInstancesSome(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVnetBondsEvery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVnetBondsNone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVnetBondsSome(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoadBalancerServiceWhereInput) validateAND(formats strfmt.Registry) error {
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

func (m *LoadBalancerServiceWhereInput) validateNOT(formats strfmt.Registry) error {
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

func (m *LoadBalancerServiceWhereInput) validateOR(formats strfmt.Registry) error {
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

func (m *LoadBalancerServiceWhereInput) validateEntityAsyncStatus(formats strfmt.Registry) error {
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

func (m *LoadBalancerServiceWhereInput) validateEntityAsyncStatusIn(formats strfmt.Registry) error {
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

func (m *LoadBalancerServiceWhereInput) validateEntityAsyncStatusNot(formats strfmt.Registry) error {
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

func (m *LoadBalancerServiceWhereInput) validateEntityAsyncStatusNotIn(formats strfmt.Registry) error {
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

func (m *LoadBalancerServiceWhereInput) validateEverouteCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.EverouteCluster) { // not required
		return nil
	}

	if m.EverouteCluster != nil {
		if err := m.EverouteCluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("everoute_cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("everoute_cluster")
			}
			return err
		}
	}

	return nil
}

func (m *LoadBalancerServiceWhereInput) validatePhase(formats strfmt.Registry) error {
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

func (m *LoadBalancerServiceWhereInput) validatePhaseIn(formats strfmt.Registry) error {
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

func (m *LoadBalancerServiceWhereInput) validatePhaseNot(formats strfmt.Registry) error {
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

func (m *LoadBalancerServiceWhereInput) validatePhaseNotIn(formats strfmt.Registry) error {
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

func (m *LoadBalancerServiceWhereInput) validateVMInstancesEvery(formats strfmt.Registry) error {
	if swag.IsZero(m.VMInstancesEvery) { // not required
		return nil
	}

	if m.VMInstancesEvery != nil {
		if err := m.VMInstancesEvery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_instances_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_instances_every")
			}
			return err
		}
	}

	return nil
}

func (m *LoadBalancerServiceWhereInput) validateVMInstancesNone(formats strfmt.Registry) error {
	if swag.IsZero(m.VMInstancesNone) { // not required
		return nil
	}

	if m.VMInstancesNone != nil {
		if err := m.VMInstancesNone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_instances_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_instances_none")
			}
			return err
		}
	}

	return nil
}

func (m *LoadBalancerServiceWhereInput) validateVMInstancesSome(formats strfmt.Registry) error {
	if swag.IsZero(m.VMInstancesSome) { // not required
		return nil
	}

	if m.VMInstancesSome != nil {
		if err := m.VMInstancesSome.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_instances_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_instances_some")
			}
			return err
		}
	}

	return nil
}

func (m *LoadBalancerServiceWhereInput) validateVnetBondsEvery(formats strfmt.Registry) error {
	if swag.IsZero(m.VnetBondsEvery) { // not required
		return nil
	}

	if m.VnetBondsEvery != nil {
		if err := m.VnetBondsEvery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vnet_bonds_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vnet_bonds_every")
			}
			return err
		}
	}

	return nil
}

func (m *LoadBalancerServiceWhereInput) validateVnetBondsNone(formats strfmt.Registry) error {
	if swag.IsZero(m.VnetBondsNone) { // not required
		return nil
	}

	if m.VnetBondsNone != nil {
		if err := m.VnetBondsNone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vnet_bonds_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vnet_bonds_none")
			}
			return err
		}
	}

	return nil
}

func (m *LoadBalancerServiceWhereInput) validateVnetBondsSome(formats strfmt.Registry) error {
	if swag.IsZero(m.VnetBondsSome) { // not required
		return nil
	}

	if m.VnetBondsSome != nil {
		if err := m.VnetBondsSome.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vnet_bonds_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vnet_bonds_some")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this load balancer service where input based on the context it is used
func (m *LoadBalancerServiceWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateEverouteCluster(ctx, formats); err != nil {
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

	if err := m.contextValidateVMInstancesEvery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMInstancesNone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMInstancesSome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVnetBondsEvery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVnetBondsNone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVnetBondsSome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoadBalancerServiceWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

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

func (m *LoadBalancerServiceWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

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

func (m *LoadBalancerServiceWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

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

func (m *LoadBalancerServiceWhereInput) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *LoadBalancerServiceWhereInput) contextValidateEntityAsyncStatusIn(ctx context.Context, formats strfmt.Registry) error {

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

func (m *LoadBalancerServiceWhereInput) contextValidateEntityAsyncStatusNot(ctx context.Context, formats strfmt.Registry) error {

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

func (m *LoadBalancerServiceWhereInput) contextValidateEntityAsyncStatusNotIn(ctx context.Context, formats strfmt.Registry) error {

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

func (m *LoadBalancerServiceWhereInput) contextValidateEverouteCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.EverouteCluster != nil {
		if err := m.EverouteCluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("everoute_cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("everoute_cluster")
			}
			return err
		}
	}

	return nil
}

func (m *LoadBalancerServiceWhereInput) contextValidatePhase(ctx context.Context, formats strfmt.Registry) error {

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

func (m *LoadBalancerServiceWhereInput) contextValidatePhaseIn(ctx context.Context, formats strfmt.Registry) error {

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

func (m *LoadBalancerServiceWhereInput) contextValidatePhaseNot(ctx context.Context, formats strfmt.Registry) error {

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

func (m *LoadBalancerServiceWhereInput) contextValidatePhaseNotIn(ctx context.Context, formats strfmt.Registry) error {

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

func (m *LoadBalancerServiceWhereInput) contextValidateVMInstancesEvery(ctx context.Context, formats strfmt.Registry) error {

	if m.VMInstancesEvery != nil {
		if err := m.VMInstancesEvery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_instances_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_instances_every")
			}
			return err
		}
	}

	return nil
}

func (m *LoadBalancerServiceWhereInput) contextValidateVMInstancesNone(ctx context.Context, formats strfmt.Registry) error {

	if m.VMInstancesNone != nil {
		if err := m.VMInstancesNone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_instances_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_instances_none")
			}
			return err
		}
	}

	return nil
}

func (m *LoadBalancerServiceWhereInput) contextValidateVMInstancesSome(ctx context.Context, formats strfmt.Registry) error {

	if m.VMInstancesSome != nil {
		if err := m.VMInstancesSome.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_instances_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_instances_some")
			}
			return err
		}
	}

	return nil
}

func (m *LoadBalancerServiceWhereInput) contextValidateVnetBondsEvery(ctx context.Context, formats strfmt.Registry) error {

	if m.VnetBondsEvery != nil {
		if err := m.VnetBondsEvery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vnet_bonds_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vnet_bonds_every")
			}
			return err
		}
	}

	return nil
}

func (m *LoadBalancerServiceWhereInput) contextValidateVnetBondsNone(ctx context.Context, formats strfmt.Registry) error {

	if m.VnetBondsNone != nil {
		if err := m.VnetBondsNone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vnet_bonds_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vnet_bonds_none")
			}
			return err
		}
	}

	return nil
}

func (m *LoadBalancerServiceWhereInput) contextValidateVnetBondsSome(ctx context.Context, formats strfmt.Registry) error {

	if m.VnetBondsSome != nil {
		if err := m.VnetBondsSome.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vnet_bonds_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vnet_bonds_some")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LoadBalancerServiceWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoadBalancerServiceWhereInput) UnmarshalBinary(b []byte) error {
	var res LoadBalancerServiceWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
