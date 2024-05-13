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

// VirtualPrivateCloudSecurityPolicyWhereInput virtual private cloud security policy where input
//
// swagger:model VirtualPrivateCloudSecurityPolicyWhereInput
type VirtualPrivateCloudSecurityPolicyWhereInput struct {

	// a n d
	AND []*VirtualPrivateCloudSecurityPolicyWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*VirtualPrivateCloudSecurityPolicyWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*VirtualPrivateCloudSecurityPolicyWhereInput `json:"OR,omitempty"`

	// description
	Description *string `json:"description,omitempty"`

	// description contains
	DescriptionContains *string `json:"description_contains,omitempty"`

	// description ends with
	DescriptionEndsWith *string `json:"description_ends_with,omitempty"`

	// description gt
	DescriptionGt *string `json:"description_gt,omitempty"`

	// description gte
	DescriptionGte *string `json:"description_gte,omitempty"`

	// description in
	DescriptionIn []string `json:"description_in,omitempty"`

	// description lt
	DescriptionLt *string `json:"description_lt,omitempty"`

	// description lte
	DescriptionLte *string `json:"description_lte,omitempty"`

	// description not
	DescriptionNot *string `json:"description_not,omitempty"`

	// description not contains
	DescriptionNotContains *string `json:"description_not_contains,omitempty"`

	// description not ends with
	DescriptionNotEndsWith *string `json:"description_not_ends_with,omitempty"`

	// description not in
	DescriptionNotIn []string `json:"description_not_in,omitempty"`

	// description not starts with
	DescriptionNotStartsWith *string `json:"description_not_starts_with,omitempty"`

	// description starts with
	DescriptionStartsWith *string `json:"description_starts_with,omitempty"`

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

	// labels every
	LabelsEvery *LabelWhereInput `json:"labels_every,omitempty"`

	// labels none
	LabelsNone *LabelWhereInput `json:"labels_none,omitempty"`

	// labels some
	LabelsSome *LabelWhereInput `json:"labels_some,omitempty"`

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

	// policy mode
	PolicyMode *VirtualPrivateCloudSecurityPolicyMode `json:"policy_mode,omitempty"`

	// policy mode in
	PolicyModeIn []VirtualPrivateCloudSecurityPolicyMode `json:"policy_mode_in,omitempty"`

	// policy mode not
	PolicyModeNot *VirtualPrivateCloudSecurityPolicyMode `json:"policy_mode_not,omitempty"`

	// policy mode not in
	PolicyModeNotIn []VirtualPrivateCloudSecurityPolicyMode `json:"policy_mode_not_in,omitempty"`

	// security groups every
	SecurityGroupsEvery *VirtualPrivateCloudSecurityGroupWhereInput `json:"security_groups_every,omitempty"`

	// security groups none
	SecurityGroupsNone *VirtualPrivateCloudSecurityGroupWhereInput `json:"security_groups_none,omitempty"`

	// security groups some
	SecurityGroupsSome *VirtualPrivateCloudSecurityGroupWhereInput `json:"security_groups_some,omitempty"`

	// vpc
	Vpc *VirtualPrivateCloudWhereInput `json:"vpc,omitempty"`
}

// Validate validates this virtual private cloud security policy where input
func (m *VirtualPrivateCloudSecurityPolicyWhereInput) Validate(formats strfmt.Registry) error {
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

	if err := m.validateLabelsEvery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelsNone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelsSome(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyModeIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyModeNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyModeNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityGroupsEvery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityGroupsNone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityGroupsSome(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) validateAND(formats strfmt.Registry) error {
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

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) validateNOT(formats strfmt.Registry) error {
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

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) validateOR(formats strfmt.Registry) error {
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

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) validateEntityAsyncStatus(formats strfmt.Registry) error {
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

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) validateEntityAsyncStatusIn(formats strfmt.Registry) error {
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

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) validateEntityAsyncStatusNot(formats strfmt.Registry) error {
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

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) validateEntityAsyncStatusNotIn(formats strfmt.Registry) error {
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

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) validateLabelsEvery(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelsEvery) { // not required
		return nil
	}

	if m.LabelsEvery != nil {
		if err := m.LabelsEvery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labels_every")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) validateLabelsNone(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelsNone) { // not required
		return nil
	}

	if m.LabelsNone != nil {
		if err := m.LabelsNone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labels_none")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) validateLabelsSome(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelsSome) { // not required
		return nil
	}

	if m.LabelsSome != nil {
		if err := m.LabelsSome.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labels_some")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) validatePolicyMode(formats strfmt.Registry) error {
	if swag.IsZero(m.PolicyMode) { // not required
		return nil
	}

	if m.PolicyMode != nil {
		if err := m.PolicyMode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy_mode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("policy_mode")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) validatePolicyModeIn(formats strfmt.Registry) error {
	if swag.IsZero(m.PolicyModeIn) { // not required
		return nil
	}

	for i := 0; i < len(m.PolicyModeIn); i++ {

		if err := m.PolicyModeIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy_mode_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("policy_mode_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) validatePolicyModeNot(formats strfmt.Registry) error {
	if swag.IsZero(m.PolicyModeNot) { // not required
		return nil
	}

	if m.PolicyModeNot != nil {
		if err := m.PolicyModeNot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy_mode_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("policy_mode_not")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) validatePolicyModeNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.PolicyModeNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.PolicyModeNotIn); i++ {

		if err := m.PolicyModeNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy_mode_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("policy_mode_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) validateSecurityGroupsEvery(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityGroupsEvery) { // not required
		return nil
	}

	if m.SecurityGroupsEvery != nil {
		if err := m.SecurityGroupsEvery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_groups_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_groups_every")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) validateSecurityGroupsNone(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityGroupsNone) { // not required
		return nil
	}

	if m.SecurityGroupsNone != nil {
		if err := m.SecurityGroupsNone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_groups_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_groups_none")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) validateSecurityGroupsSome(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityGroupsSome) { // not required
		return nil
	}

	if m.SecurityGroupsSome != nil {
		if err := m.SecurityGroupsSome.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_groups_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_groups_some")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) validateVpc(formats strfmt.Registry) error {
	if swag.IsZero(m.Vpc) { // not required
		return nil
	}

	if m.Vpc != nil {
		if err := m.Vpc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vpc")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this virtual private cloud security policy where input based on the context it is used
func (m *VirtualPrivateCloudSecurityPolicyWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateLabelsEvery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabelsNone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabelsSome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePolicyMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePolicyModeIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePolicyModeNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePolicyModeNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurityGroupsEvery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurityGroupsNone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurityGroupsSome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVpc(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) contextValidateEntityAsyncStatusIn(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) contextValidateEntityAsyncStatusNot(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) contextValidateEntityAsyncStatusNotIn(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) contextValidateLabelsEvery(ctx context.Context, formats strfmt.Registry) error {

	if m.LabelsEvery != nil {
		if err := m.LabelsEvery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labels_every")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) contextValidateLabelsNone(ctx context.Context, formats strfmt.Registry) error {

	if m.LabelsNone != nil {
		if err := m.LabelsNone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labels_none")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) contextValidateLabelsSome(ctx context.Context, formats strfmt.Registry) error {

	if m.LabelsSome != nil {
		if err := m.LabelsSome.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labels_some")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) contextValidatePolicyMode(ctx context.Context, formats strfmt.Registry) error {

	if m.PolicyMode != nil {
		if err := m.PolicyMode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy_mode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("policy_mode")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) contextValidatePolicyModeIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PolicyModeIn); i++ {

		if err := m.PolicyModeIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy_mode_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("policy_mode_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) contextValidatePolicyModeNot(ctx context.Context, formats strfmt.Registry) error {

	if m.PolicyModeNot != nil {
		if err := m.PolicyModeNot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy_mode_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("policy_mode_not")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) contextValidatePolicyModeNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PolicyModeNotIn); i++ {

		if err := m.PolicyModeNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy_mode_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("policy_mode_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) contextValidateSecurityGroupsEvery(ctx context.Context, formats strfmt.Registry) error {

	if m.SecurityGroupsEvery != nil {
		if err := m.SecurityGroupsEvery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_groups_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_groups_every")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) contextValidateSecurityGroupsNone(ctx context.Context, formats strfmt.Registry) error {

	if m.SecurityGroupsNone != nil {
		if err := m.SecurityGroupsNone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_groups_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_groups_none")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) contextValidateSecurityGroupsSome(ctx context.Context, formats strfmt.Registry) error {

	if m.SecurityGroupsSome != nil {
		if err := m.SecurityGroupsSome.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("security_groups_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("security_groups_some")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudSecurityPolicyWhereInput) contextValidateVpc(ctx context.Context, formats strfmt.Registry) error {

	if m.Vpc != nil {
		if err := m.Vpc.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vpc")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VirtualPrivateCloudSecurityPolicyWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualPrivateCloudSecurityPolicyWhereInput) UnmarshalBinary(b []byte) error {
	var res VirtualPrivateCloudSecurityPolicyWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
