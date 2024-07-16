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
	"github.com/go-openapi/validate"
)

// VirtualPrivateCloudSecurityGroupCreationParams virtual private cloud security group creation params
//
// swagger:model VirtualPrivateCloudSecurityGroupCreationParams
type VirtualPrivateCloudSecurityGroupCreationParams struct {

	// description
	Description *string `json:"description,omitempty"`

	// label groups
	LabelGroups []*LabelGroup `json:"label_groups,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// vm ids
	VMIds []string `json:"vm_ids,omitempty"`

	// vpc id
	// Required: true
	VpcID *string `json:"vpc_id"`
}

// Validate validates this virtual private cloud security group creation params
func (m *VirtualPrivateCloudSecurityGroupCreationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabelGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpcID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualPrivateCloudSecurityGroupCreationParams) validateLabelGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.LabelGroups); i++ {
		if swag.IsZero(m.LabelGroups[i]) { // not required
			continue
		}

		if m.LabelGroups[i] != nil {
			if err := m.LabelGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("label_groups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("label_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudSecurityGroupCreationParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VirtualPrivateCloudSecurityGroupCreationParams) validateVpcID(formats strfmt.Registry) error {

	if err := validate.Required("vpc_id", "body", m.VpcID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this virtual private cloud security group creation params based on the context it is used
func (m *VirtualPrivateCloudSecurityGroupCreationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLabelGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualPrivateCloudSecurityGroupCreationParams) contextValidateLabelGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LabelGroups); i++ {

		if m.LabelGroups[i] != nil {
			if err := m.LabelGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("label_groups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("label_groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VirtualPrivateCloudSecurityGroupCreationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualPrivateCloudSecurityGroupCreationParams) UnmarshalBinary(b []byte) error {
	var res VirtualPrivateCloudSecurityGroupCreationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
