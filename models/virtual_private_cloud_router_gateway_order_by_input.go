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

// VirtualPrivateCloudRouterGatewayOrderByInput virtual private cloud router gateway order by input
//
// swagger:model VirtualPrivateCloudRouterGatewayOrderByInput
type VirtualPrivateCloudRouterGatewayOrderByInput string

func NewVirtualPrivateCloudRouterGatewayOrderByInput(value VirtualPrivateCloudRouterGatewayOrderByInput) *VirtualPrivateCloudRouterGatewayOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VirtualPrivateCloudRouterGatewayOrderByInput.
func (m VirtualPrivateCloudRouterGatewayOrderByInput) Pointer() *VirtualPrivateCloudRouterGatewayOrderByInput {
	return &m
}

const (

	// VirtualPrivateCloudRouterGatewayOrderByInputEntityAsyncStatusASC captures enum value "entityAsyncStatus_ASC"
	VirtualPrivateCloudRouterGatewayOrderByInputEntityAsyncStatusASC VirtualPrivateCloudRouterGatewayOrderByInput = "entityAsyncStatus_ASC"

	// VirtualPrivateCloudRouterGatewayOrderByInputEntityAsyncStatusDESC captures enum value "entityAsyncStatus_DESC"
	VirtualPrivateCloudRouterGatewayOrderByInputEntityAsyncStatusDESC VirtualPrivateCloudRouterGatewayOrderByInput = "entityAsyncStatus_DESC"

	// VirtualPrivateCloudRouterGatewayOrderByInputExternalIPASC captures enum value "external_ip_ASC"
	VirtualPrivateCloudRouterGatewayOrderByInputExternalIPASC VirtualPrivateCloudRouterGatewayOrderByInput = "external_ip_ASC"

	// VirtualPrivateCloudRouterGatewayOrderByInputExternalIPDESC captures enum value "external_ip_DESC"
	VirtualPrivateCloudRouterGatewayOrderByInputExternalIPDESC VirtualPrivateCloudRouterGatewayOrderByInput = "external_ip_DESC"

	// VirtualPrivateCloudRouterGatewayOrderByInputExternalIpsASC captures enum value "external_ips_ASC"
	VirtualPrivateCloudRouterGatewayOrderByInputExternalIpsASC VirtualPrivateCloudRouterGatewayOrderByInput = "external_ips_ASC"

	// VirtualPrivateCloudRouterGatewayOrderByInputExternalIpsDESC captures enum value "external_ips_DESC"
	VirtualPrivateCloudRouterGatewayOrderByInputExternalIpsDESC VirtualPrivateCloudRouterGatewayOrderByInput = "external_ips_DESC"

	// VirtualPrivateCloudRouterGatewayOrderByInputIDASC captures enum value "id_ASC"
	VirtualPrivateCloudRouterGatewayOrderByInputIDASC VirtualPrivateCloudRouterGatewayOrderByInput = "id_ASC"

	// VirtualPrivateCloudRouterGatewayOrderByInputIDDESC captures enum value "id_DESC"
	VirtualPrivateCloudRouterGatewayOrderByInputIDDESC VirtualPrivateCloudRouterGatewayOrderByInput = "id_DESC"

	// VirtualPrivateCloudRouterGatewayOrderByInputLocalIDASC captures enum value "local_id_ASC"
	VirtualPrivateCloudRouterGatewayOrderByInputLocalIDASC VirtualPrivateCloudRouterGatewayOrderByInput = "local_id_ASC"

	// VirtualPrivateCloudRouterGatewayOrderByInputLocalIDDESC captures enum value "local_id_DESC"
	VirtualPrivateCloudRouterGatewayOrderByInputLocalIDDESC VirtualPrivateCloudRouterGatewayOrderByInput = "local_id_DESC"

	// VirtualPrivateCloudRouterGatewayOrderByInputNameASC captures enum value "name_ASC"
	VirtualPrivateCloudRouterGatewayOrderByInputNameASC VirtualPrivateCloudRouterGatewayOrderByInput = "name_ASC"

	// VirtualPrivateCloudRouterGatewayOrderByInputNameDESC captures enum value "name_DESC"
	VirtualPrivateCloudRouterGatewayOrderByInputNameDESC VirtualPrivateCloudRouterGatewayOrderByInput = "name_DESC"

	// VirtualPrivateCloudRouterGatewayOrderByInputNexthopIPASC captures enum value "nexthop_ip_ASC"
	VirtualPrivateCloudRouterGatewayOrderByInputNexthopIPASC VirtualPrivateCloudRouterGatewayOrderByInput = "nexthop_ip_ASC"

	// VirtualPrivateCloudRouterGatewayOrderByInputNexthopIPDESC captures enum value "nexthop_ip_DESC"
	VirtualPrivateCloudRouterGatewayOrderByInputNexthopIPDESC VirtualPrivateCloudRouterGatewayOrderByInput = "nexthop_ip_DESC"

	// VirtualPrivateCloudRouterGatewayOrderByInputRulesASC captures enum value "rules_ASC"
	VirtualPrivateCloudRouterGatewayOrderByInputRulesASC VirtualPrivateCloudRouterGatewayOrderByInput = "rules_ASC"

	// VirtualPrivateCloudRouterGatewayOrderByInputRulesDESC captures enum value "rules_DESC"
	VirtualPrivateCloudRouterGatewayOrderByInputRulesDESC VirtualPrivateCloudRouterGatewayOrderByInput = "rules_DESC"
)

// for schema
var virtualPrivateCloudRouterGatewayOrderByInputEnum []interface{}

func init() {
	var res []VirtualPrivateCloudRouterGatewayOrderByInput
	if err := json.Unmarshal([]byte(`["entityAsyncStatus_ASC","entityAsyncStatus_DESC","external_ip_ASC","external_ip_DESC","external_ips_ASC","external_ips_DESC","id_ASC","id_DESC","local_id_ASC","local_id_DESC","name_ASC","name_DESC","nexthop_ip_ASC","nexthop_ip_DESC","rules_ASC","rules_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		virtualPrivateCloudRouterGatewayOrderByInputEnum = append(virtualPrivateCloudRouterGatewayOrderByInputEnum, v)
	}
}

func (m VirtualPrivateCloudRouterGatewayOrderByInput) validateVirtualPrivateCloudRouterGatewayOrderByInputEnum(path, location string, value VirtualPrivateCloudRouterGatewayOrderByInput) error {
	if err := validate.EnumCase(path, location, value, virtualPrivateCloudRouterGatewayOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this virtual private cloud router gateway order by input
func (m VirtualPrivateCloudRouterGatewayOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateVirtualPrivateCloudRouterGatewayOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this virtual private cloud router gateway order by input based on context it is used
func (m VirtualPrivateCloudRouterGatewayOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
