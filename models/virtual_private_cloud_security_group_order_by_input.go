// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// VirtualPrivateCloudSecurityGroupOrderByInput virtual private cloud security group order by input
//
// swagger:model VirtualPrivateCloudSecurityGroupOrderByInput
type VirtualPrivateCloudSecurityGroupOrderByInput string

func NewVirtualPrivateCloudSecurityGroupOrderByInput(value VirtualPrivateCloudSecurityGroupOrderByInput) *VirtualPrivateCloudSecurityGroupOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VirtualPrivateCloudSecurityGroupOrderByInput.
func (m VirtualPrivateCloudSecurityGroupOrderByInput) Pointer() *VirtualPrivateCloudSecurityGroupOrderByInput {
	return &m
}

const (

	// VirtualPrivateCloudSecurityGroupOrderByInputDefaultForVpcASC captures enum value "default_for_vpc_ASC"
	VirtualPrivateCloudSecurityGroupOrderByInputDefaultForVpcASC VirtualPrivateCloudSecurityGroupOrderByInput = "default_for_vpc_ASC"

	// VirtualPrivateCloudSecurityGroupOrderByInputDefaultForVpcDESC captures enum value "default_for_vpc_DESC"
	VirtualPrivateCloudSecurityGroupOrderByInputDefaultForVpcDESC VirtualPrivateCloudSecurityGroupOrderByInput = "default_for_vpc_DESC"

	// VirtualPrivateCloudSecurityGroupOrderByInputDescriptionASC captures enum value "description_ASC"
	VirtualPrivateCloudSecurityGroupOrderByInputDescriptionASC VirtualPrivateCloudSecurityGroupOrderByInput = "description_ASC"

	// VirtualPrivateCloudSecurityGroupOrderByInputDescriptionDESC captures enum value "description_DESC"
	VirtualPrivateCloudSecurityGroupOrderByInputDescriptionDESC VirtualPrivateCloudSecurityGroupOrderByInput = "description_DESC"

	// VirtualPrivateCloudSecurityGroupOrderByInputEntityAsyncStatusASC captures enum value "entityAsyncStatus_ASC"
	VirtualPrivateCloudSecurityGroupOrderByInputEntityAsyncStatusASC VirtualPrivateCloudSecurityGroupOrderByInput = "entityAsyncStatus_ASC"

	// VirtualPrivateCloudSecurityGroupOrderByInputEntityAsyncStatusDESC captures enum value "entityAsyncStatus_DESC"
	VirtualPrivateCloudSecurityGroupOrderByInputEntityAsyncStatusDESC VirtualPrivateCloudSecurityGroupOrderByInput = "entityAsyncStatus_DESC"

	// VirtualPrivateCloudSecurityGroupOrderByInputIDASC captures enum value "id_ASC"
	VirtualPrivateCloudSecurityGroupOrderByInputIDASC VirtualPrivateCloudSecurityGroupOrderByInput = "id_ASC"

	// VirtualPrivateCloudSecurityGroupOrderByInputIDDESC captures enum value "id_DESC"
	VirtualPrivateCloudSecurityGroupOrderByInputIDDESC VirtualPrivateCloudSecurityGroupOrderByInput = "id_DESC"

	// VirtualPrivateCloudSecurityGroupOrderByInputLabelGroupsASC captures enum value "label_groups_ASC"
	VirtualPrivateCloudSecurityGroupOrderByInputLabelGroupsASC VirtualPrivateCloudSecurityGroupOrderByInput = "label_groups_ASC"

	// VirtualPrivateCloudSecurityGroupOrderByInputLabelGroupsDESC captures enum value "label_groups_DESC"
	VirtualPrivateCloudSecurityGroupOrderByInputLabelGroupsDESC VirtualPrivateCloudSecurityGroupOrderByInput = "label_groups_DESC"

	// VirtualPrivateCloudSecurityGroupOrderByInputLocalIDASC captures enum value "local_id_ASC"
	VirtualPrivateCloudSecurityGroupOrderByInputLocalIDASC VirtualPrivateCloudSecurityGroupOrderByInput = "local_id_ASC"

	// VirtualPrivateCloudSecurityGroupOrderByInputLocalIDDESC captures enum value "local_id_DESC"
	VirtualPrivateCloudSecurityGroupOrderByInputLocalIDDESC VirtualPrivateCloudSecurityGroupOrderByInput = "local_id_DESC"

	// VirtualPrivateCloudSecurityGroupOrderByInputNameASC captures enum value "name_ASC"
	VirtualPrivateCloudSecurityGroupOrderByInputNameASC VirtualPrivateCloudSecurityGroupOrderByInput = "name_ASC"

	// VirtualPrivateCloudSecurityGroupOrderByInputNameDESC captures enum value "name_DESC"
	VirtualPrivateCloudSecurityGroupOrderByInputNameDESC VirtualPrivateCloudSecurityGroupOrderByInput = "name_DESC"
)

// for schema
var virtualPrivateCloudSecurityGroupOrderByInputEnum []interface{}

func init() {
	var res []VirtualPrivateCloudSecurityGroupOrderByInput
	if err := json.Unmarshal([]byte(`["default_for_vpc_ASC","default_for_vpc_DESC","description_ASC","description_DESC","entityAsyncStatus_ASC","entityAsyncStatus_DESC","id_ASC","id_DESC","label_groups_ASC","label_groups_DESC","local_id_ASC","local_id_DESC","name_ASC","name_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		virtualPrivateCloudSecurityGroupOrderByInputEnum = append(virtualPrivateCloudSecurityGroupOrderByInputEnum, v)
	}
}

func (m VirtualPrivateCloudSecurityGroupOrderByInput) validateVirtualPrivateCloudSecurityGroupOrderByInputEnum(path, location string, value VirtualPrivateCloudSecurityGroupOrderByInput) error {
	if err := validate.EnumCase(path, location, value, virtualPrivateCloudSecurityGroupOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this virtual private cloud security group order by input
func (m VirtualPrivateCloudSecurityGroupOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateVirtualPrivateCloudSecurityGroupOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this virtual private cloud security group order by input based on context it is used
func (m VirtualPrivateCloudSecurityGroupOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
