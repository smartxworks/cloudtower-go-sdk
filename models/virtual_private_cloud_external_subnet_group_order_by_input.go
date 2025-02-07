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

// VirtualPrivateCloudExternalSubnetGroupOrderByInput virtual private cloud external subnet group order by input
//
// swagger:model VirtualPrivateCloudExternalSubnetGroupOrderByInput
type VirtualPrivateCloudExternalSubnetGroupOrderByInput string

func NewVirtualPrivateCloudExternalSubnetGroupOrderByInput(value VirtualPrivateCloudExternalSubnetGroupOrderByInput) *VirtualPrivateCloudExternalSubnetGroupOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VirtualPrivateCloudExternalSubnetGroupOrderByInput.
func (m VirtualPrivateCloudExternalSubnetGroupOrderByInput) Pointer() *VirtualPrivateCloudExternalSubnetGroupOrderByInput {
	return &m
}

const (

	// VirtualPrivateCloudExternalSubnetGroupOrderByInputDescriptionASC captures enum value "description_ASC"
	VirtualPrivateCloudExternalSubnetGroupOrderByInputDescriptionASC VirtualPrivateCloudExternalSubnetGroupOrderByInput = "description_ASC"

	// VirtualPrivateCloudExternalSubnetGroupOrderByInputDescriptionDESC captures enum value "description_DESC"
	VirtualPrivateCloudExternalSubnetGroupOrderByInputDescriptionDESC VirtualPrivateCloudExternalSubnetGroupOrderByInput = "description_DESC"

	// VirtualPrivateCloudExternalSubnetGroupOrderByInputEntityAsyncStatusASC captures enum value "entityAsyncStatus_ASC"
	VirtualPrivateCloudExternalSubnetGroupOrderByInputEntityAsyncStatusASC VirtualPrivateCloudExternalSubnetGroupOrderByInput = "entityAsyncStatus_ASC"

	// VirtualPrivateCloudExternalSubnetGroupOrderByInputEntityAsyncStatusDESC captures enum value "entityAsyncStatus_DESC"
	VirtualPrivateCloudExternalSubnetGroupOrderByInputEntityAsyncStatusDESC VirtualPrivateCloudExternalSubnetGroupOrderByInput = "entityAsyncStatus_DESC"

	// VirtualPrivateCloudExternalSubnetGroupOrderByInputExclusiveASC captures enum value "exclusive_ASC"
	VirtualPrivateCloudExternalSubnetGroupOrderByInputExclusiveASC VirtualPrivateCloudExternalSubnetGroupOrderByInput = "exclusive_ASC"

	// VirtualPrivateCloudExternalSubnetGroupOrderByInputExclusiveDESC captures enum value "exclusive_DESC"
	VirtualPrivateCloudExternalSubnetGroupOrderByInputExclusiveDESC VirtualPrivateCloudExternalSubnetGroupOrderByInput = "exclusive_DESC"

	// VirtualPrivateCloudExternalSubnetGroupOrderByInputExternalSubnetsSpecASC captures enum value "external_subnets_spec_ASC"
	VirtualPrivateCloudExternalSubnetGroupOrderByInputExternalSubnetsSpecASC VirtualPrivateCloudExternalSubnetGroupOrderByInput = "external_subnets_spec_ASC"

	// VirtualPrivateCloudExternalSubnetGroupOrderByInputExternalSubnetsSpecDESC captures enum value "external_subnets_spec_DESC"
	VirtualPrivateCloudExternalSubnetGroupOrderByInputExternalSubnetsSpecDESC VirtualPrivateCloudExternalSubnetGroupOrderByInput = "external_subnets_spec_DESC"

	// VirtualPrivateCloudExternalSubnetGroupOrderByInputIDASC captures enum value "id_ASC"
	VirtualPrivateCloudExternalSubnetGroupOrderByInputIDASC VirtualPrivateCloudExternalSubnetGroupOrderByInput = "id_ASC"

	// VirtualPrivateCloudExternalSubnetGroupOrderByInputIDDESC captures enum value "id_DESC"
	VirtualPrivateCloudExternalSubnetGroupOrderByInputIDDESC VirtualPrivateCloudExternalSubnetGroupOrderByInput = "id_DESC"

	// VirtualPrivateCloudExternalSubnetGroupOrderByInputLocalIDASC captures enum value "local_id_ASC"
	VirtualPrivateCloudExternalSubnetGroupOrderByInputLocalIDASC VirtualPrivateCloudExternalSubnetGroupOrderByInput = "local_id_ASC"

	// VirtualPrivateCloudExternalSubnetGroupOrderByInputLocalIDDESC captures enum value "local_id_DESC"
	VirtualPrivateCloudExternalSubnetGroupOrderByInputLocalIDDESC VirtualPrivateCloudExternalSubnetGroupOrderByInput = "local_id_DESC"

	// VirtualPrivateCloudExternalSubnetGroupOrderByInputNameASC captures enum value "name_ASC"
	VirtualPrivateCloudExternalSubnetGroupOrderByInputNameASC VirtualPrivateCloudExternalSubnetGroupOrderByInput = "name_ASC"

	// VirtualPrivateCloudExternalSubnetGroupOrderByInputNameDESC captures enum value "name_DESC"
	VirtualPrivateCloudExternalSubnetGroupOrderByInputNameDESC VirtualPrivateCloudExternalSubnetGroupOrderByInput = "name_DESC"

	// VirtualPrivateCloudExternalSubnetGroupOrderByInputSharedInEdgeGatewayGroupASC captures enum value "shared_in_edge_gateway_group_ASC"
	VirtualPrivateCloudExternalSubnetGroupOrderByInputSharedInEdgeGatewayGroupASC VirtualPrivateCloudExternalSubnetGroupOrderByInput = "shared_in_edge_gateway_group_ASC"

	// VirtualPrivateCloudExternalSubnetGroupOrderByInputSharedInEdgeGatewayGroupDESC captures enum value "shared_in_edge_gateway_group_DESC"
	VirtualPrivateCloudExternalSubnetGroupOrderByInputSharedInEdgeGatewayGroupDESC VirtualPrivateCloudExternalSubnetGroupOrderByInput = "shared_in_edge_gateway_group_DESC"
)

// for schema
var virtualPrivateCloudExternalSubnetGroupOrderByInputEnum []interface{}

func init() {
	var res []VirtualPrivateCloudExternalSubnetGroupOrderByInput
	if err := json.Unmarshal([]byte(`["description_ASC","description_DESC","entityAsyncStatus_ASC","entityAsyncStatus_DESC","exclusive_ASC","exclusive_DESC","external_subnets_spec_ASC","external_subnets_spec_DESC","id_ASC","id_DESC","local_id_ASC","local_id_DESC","name_ASC","name_DESC","shared_in_edge_gateway_group_ASC","shared_in_edge_gateway_group_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		virtualPrivateCloudExternalSubnetGroupOrderByInputEnum = append(virtualPrivateCloudExternalSubnetGroupOrderByInputEnum, v)
	}
}

func (m VirtualPrivateCloudExternalSubnetGroupOrderByInput) validateVirtualPrivateCloudExternalSubnetGroupOrderByInputEnum(path, location string, value VirtualPrivateCloudExternalSubnetGroupOrderByInput) error {
	if err := validate.EnumCase(path, location, value, virtualPrivateCloudExternalSubnetGroupOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this virtual private cloud external subnet group order by input
func (m VirtualPrivateCloudExternalSubnetGroupOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateVirtualPrivateCloudExternalSubnetGroupOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this virtual private cloud external subnet group order by input based on context it is used
func (m VirtualPrivateCloudExternalSubnetGroupOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
