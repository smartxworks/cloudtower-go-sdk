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

// SecurityPolicyOrderByInput security policy order by input
//
// swagger:model SecurityPolicyOrderByInput
type SecurityPolicyOrderByInput string

func NewSecurityPolicyOrderByInput(value SecurityPolicyOrderByInput) *SecurityPolicyOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated SecurityPolicyOrderByInput.
func (m SecurityPolicyOrderByInput) Pointer() *SecurityPolicyOrderByInput {
	return &m
}

const (

	// SecurityPolicyOrderByInputApplyToASC captures enum value "apply_to_ASC"
	SecurityPolicyOrderByInputApplyToASC SecurityPolicyOrderByInput = "apply_to_ASC"

	// SecurityPolicyOrderByInputApplyToDESC captures enum value "apply_to_DESC"
	SecurityPolicyOrderByInputApplyToDESC SecurityPolicyOrderByInput = "apply_to_DESC"

	// SecurityPolicyOrderByInputDescriptionASC captures enum value "description_ASC"
	SecurityPolicyOrderByInputDescriptionASC SecurityPolicyOrderByInput = "description_ASC"

	// SecurityPolicyOrderByInputDescriptionDESC captures enum value "description_DESC"
	SecurityPolicyOrderByInputDescriptionDESC SecurityPolicyOrderByInput = "description_DESC"

	// SecurityPolicyOrderByInputEgressASC captures enum value "egress_ASC"
	SecurityPolicyOrderByInputEgressASC SecurityPolicyOrderByInput = "egress_ASC"

	// SecurityPolicyOrderByInputEgressDESC captures enum value "egress_DESC"
	SecurityPolicyOrderByInputEgressDESC SecurityPolicyOrderByInput = "egress_DESC"

	// SecurityPolicyOrderByInputIDASC captures enum value "id_ASC"
	SecurityPolicyOrderByInputIDASC SecurityPolicyOrderByInput = "id_ASC"

	// SecurityPolicyOrderByInputIDDESC captures enum value "id_DESC"
	SecurityPolicyOrderByInputIDDESC SecurityPolicyOrderByInput = "id_DESC"

	// SecurityPolicyOrderByInputIngressASC captures enum value "ingress_ASC"
	SecurityPolicyOrderByInputIngressASC SecurityPolicyOrderByInput = "ingress_ASC"

	// SecurityPolicyOrderByInputIngressDESC captures enum value "ingress_DESC"
	SecurityPolicyOrderByInputIngressDESC SecurityPolicyOrderByInput = "ingress_DESC"

	// SecurityPolicyOrderByInputIsBlocklistASC captures enum value "is_blocklist_ASC"
	SecurityPolicyOrderByInputIsBlocklistASC SecurityPolicyOrderByInput = "is_blocklist_ASC"

	// SecurityPolicyOrderByInputIsBlocklistDESC captures enum value "is_blocklist_DESC"
	SecurityPolicyOrderByInputIsBlocklistDESC SecurityPolicyOrderByInput = "is_blocklist_DESC"

	// SecurityPolicyOrderByInputNameASC captures enum value "name_ASC"
	SecurityPolicyOrderByInputNameASC SecurityPolicyOrderByInput = "name_ASC"

	// SecurityPolicyOrderByInputNameDESC captures enum value "name_DESC"
	SecurityPolicyOrderByInputNameDESC SecurityPolicyOrderByInput = "name_DESC"

	// SecurityPolicyOrderByInputPolicyModeASC captures enum value "policy_mode_ASC"
	SecurityPolicyOrderByInputPolicyModeASC SecurityPolicyOrderByInput = "policy_mode_ASC"

	// SecurityPolicyOrderByInputPolicyModeDESC captures enum value "policy_mode_DESC"
	SecurityPolicyOrderByInputPolicyModeDESC SecurityPolicyOrderByInput = "policy_mode_DESC"
)

// for schema
var securityPolicyOrderByInputEnum []interface{}

func init() {
	var res []SecurityPolicyOrderByInput
	if err := json.Unmarshal([]byte(`["apply_to_ASC","apply_to_DESC","description_ASC","description_DESC","egress_ASC","egress_DESC","id_ASC","id_DESC","ingress_ASC","ingress_DESC","is_blocklist_ASC","is_blocklist_DESC","name_ASC","name_DESC","policy_mode_ASC","policy_mode_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		securityPolicyOrderByInputEnum = append(securityPolicyOrderByInputEnum, v)
	}
}

func (m SecurityPolicyOrderByInput) validateSecurityPolicyOrderByInputEnum(path, location string, value SecurityPolicyOrderByInput) error {
	if err := validate.EnumCase(path, location, value, securityPolicyOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this security policy order by input
func (m SecurityPolicyOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSecurityPolicyOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this security policy order by input based on context it is used
func (m SecurityPolicyOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
