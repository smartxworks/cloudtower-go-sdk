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

// V2EverouteLicenseOrderByInput v2 everoute license order by input
//
// swagger:model V2EverouteLicenseOrderByInput
type V2EverouteLicenseOrderByInput string

func NewV2EverouteLicenseOrderByInput(value V2EverouteLicenseOrderByInput) *V2EverouteLicenseOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated V2EverouteLicenseOrderByInput.
func (m V2EverouteLicenseOrderByInput) Pointer() *V2EverouteLicenseOrderByInput {
	return &m
}

const (

	// V2EverouteLicenseOrderByInputCodeASC captures enum value "code_ASC"
	V2EverouteLicenseOrderByInputCodeASC V2EverouteLicenseOrderByInput = "code_ASC"

	// V2EverouteLicenseOrderByInputCodeDESC captures enum value "code_DESC"
	V2EverouteLicenseOrderByInputCodeDESC V2EverouteLicenseOrderByInput = "code_DESC"

	// V2EverouteLicenseOrderByInputExpireDateASC captures enum value "expire_date_ASC"
	V2EverouteLicenseOrderByInputExpireDateASC V2EverouteLicenseOrderByInput = "expire_date_ASC"

	// V2EverouteLicenseOrderByInputExpireDateDESC captures enum value "expire_date_DESC"
	V2EverouteLicenseOrderByInputExpireDateDESC V2EverouteLicenseOrderByInput = "expire_date_DESC"

	// V2EverouteLicenseOrderByInputFeatureTypeASC captures enum value "feature_type_ASC"
	V2EverouteLicenseOrderByInputFeatureTypeASC V2EverouteLicenseOrderByInput = "feature_type_ASC"

	// V2EverouteLicenseOrderByInputFeatureTypeDESC captures enum value "feature_type_DESC"
	V2EverouteLicenseOrderByInputFeatureTypeDESC V2EverouteLicenseOrderByInput = "feature_type_DESC"

	// V2EverouteLicenseOrderByInputIDASC captures enum value "id_ASC"
	V2EverouteLicenseOrderByInputIDASC V2EverouteLicenseOrderByInput = "id_ASC"

	// V2EverouteLicenseOrderByInputIDDESC captures enum value "id_DESC"
	V2EverouteLicenseOrderByInputIDDESC V2EverouteLicenseOrderByInput = "id_DESC"

	// V2EverouteLicenseOrderByInputMaxSocketNumASC captures enum value "max_socket_num_ASC"
	V2EverouteLicenseOrderByInputMaxSocketNumASC V2EverouteLicenseOrderByInput = "max_socket_num_ASC"

	// V2EverouteLicenseOrderByInputMaxSocketNumDESC captures enum value "max_socket_num_DESC"
	V2EverouteLicenseOrderByInputMaxSocketNumDESC V2EverouteLicenseOrderByInput = "max_socket_num_DESC"

	// V2EverouteLicenseOrderByInputMaxVcpuNumASC captures enum value "max_vcpu_num_ASC"
	V2EverouteLicenseOrderByInputMaxVcpuNumASC V2EverouteLicenseOrderByInput = "max_vcpu_num_ASC"

	// V2EverouteLicenseOrderByInputMaxVcpuNumDESC captures enum value "max_vcpu_num_DESC"
	V2EverouteLicenseOrderByInputMaxVcpuNumDESC V2EverouteLicenseOrderByInput = "max_vcpu_num_DESC"

	// V2EverouteLicenseOrderByInputMaxVMNumASC captures enum value "max_vm_num_ASC"
	V2EverouteLicenseOrderByInputMaxVMNumASC V2EverouteLicenseOrderByInput = "max_vm_num_ASC"

	// V2EverouteLicenseOrderByInputMaxVMNumDESC captures enum value "max_vm_num_DESC"
	V2EverouteLicenseOrderByInputMaxVMNumDESC V2EverouteLicenseOrderByInput = "max_vm_num_DESC"

	// V2EverouteLicenseOrderByInputMaxVpcSocketNumASC captures enum value "max_vpc_socket_num_ASC"
	V2EverouteLicenseOrderByInputMaxVpcSocketNumASC V2EverouteLicenseOrderByInput = "max_vpc_socket_num_ASC"

	// V2EverouteLicenseOrderByInputMaxVpcSocketNumDESC captures enum value "max_vpc_socket_num_DESC"
	V2EverouteLicenseOrderByInputMaxVpcSocketNumDESC V2EverouteLicenseOrderByInput = "max_vpc_socket_num_DESC"

	// V2EverouteLicenseOrderByInputPricingTypeASC captures enum value "pricing_type_ASC"
	V2EverouteLicenseOrderByInputPricingTypeASC V2EverouteLicenseOrderByInput = "pricing_type_ASC"

	// V2EverouteLicenseOrderByInputPricingTypeDESC captures enum value "pricing_type_DESC"
	V2EverouteLicenseOrderByInputPricingTypeDESC V2EverouteLicenseOrderByInput = "pricing_type_DESC"

	// V2EverouteLicenseOrderByInputSerialASC captures enum value "serial_ASC"
	V2EverouteLicenseOrderByInputSerialASC V2EverouteLicenseOrderByInput = "serial_ASC"

	// V2EverouteLicenseOrderByInputSerialDESC captures enum value "serial_DESC"
	V2EverouteLicenseOrderByInputSerialDESC V2EverouteLicenseOrderByInput = "serial_DESC"

	// V2EverouteLicenseOrderByInputSignDateASC captures enum value "sign_date_ASC"
	V2EverouteLicenseOrderByInputSignDateASC V2EverouteLicenseOrderByInput = "sign_date_ASC"

	// V2EverouteLicenseOrderByInputSignDateDESC captures enum value "sign_date_DESC"
	V2EverouteLicenseOrderByInputSignDateDESC V2EverouteLicenseOrderByInput = "sign_date_DESC"

	// V2EverouteLicenseOrderByInputSoftwareEditionASC captures enum value "software_edition_ASC"
	V2EverouteLicenseOrderByInputSoftwareEditionASC V2EverouteLicenseOrderByInput = "software_edition_ASC"

	// V2EverouteLicenseOrderByInputSoftwareEditionDESC captures enum value "software_edition_DESC"
	V2EverouteLicenseOrderByInputSoftwareEditionDESC V2EverouteLicenseOrderByInput = "software_edition_DESC"

	// V2EverouteLicenseOrderByInputTypeASC captures enum value "type_ASC"
	V2EverouteLicenseOrderByInputTypeASC V2EverouteLicenseOrderByInput = "type_ASC"

	// V2EverouteLicenseOrderByInputTypeDESC captures enum value "type_DESC"
	V2EverouteLicenseOrderByInputTypeDESC V2EverouteLicenseOrderByInput = "type_DESC"

	// V2EverouteLicenseOrderByInputUIDASC captures enum value "uid_ASC"
	V2EverouteLicenseOrderByInputUIDASC V2EverouteLicenseOrderByInput = "uid_ASC"

	// V2EverouteLicenseOrderByInputUIDDESC captures enum value "uid_DESC"
	V2EverouteLicenseOrderByInputUIDDESC V2EverouteLicenseOrderByInput = "uid_DESC"

	// V2EverouteLicenseOrderByInputVersionASC captures enum value "version_ASC"
	V2EverouteLicenseOrderByInputVersionASC V2EverouteLicenseOrderByInput = "version_ASC"

	// V2EverouteLicenseOrderByInputVersionDESC captures enum value "version_DESC"
	V2EverouteLicenseOrderByInputVersionDESC V2EverouteLicenseOrderByInput = "version_DESC"
)

// for schema
var v2EverouteLicenseOrderByInputEnum []interface{}

func init() {
	var res []V2EverouteLicenseOrderByInput
	if err := json.Unmarshal([]byte(`["code_ASC","code_DESC","expire_date_ASC","expire_date_DESC","feature_type_ASC","feature_type_DESC","id_ASC","id_DESC","max_socket_num_ASC","max_socket_num_DESC","max_vcpu_num_ASC","max_vcpu_num_DESC","max_vm_num_ASC","max_vm_num_DESC","max_vpc_socket_num_ASC","max_vpc_socket_num_DESC","pricing_type_ASC","pricing_type_DESC","serial_ASC","serial_DESC","sign_date_ASC","sign_date_DESC","software_edition_ASC","software_edition_DESC","type_ASC","type_DESC","uid_ASC","uid_DESC","version_ASC","version_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v2EverouteLicenseOrderByInputEnum = append(v2EverouteLicenseOrderByInputEnum, v)
	}
}

func (m V2EverouteLicenseOrderByInput) validateV2EverouteLicenseOrderByInputEnum(path, location string, value V2EverouteLicenseOrderByInput) error {
	if err := validate.EnumCase(path, location, value, v2EverouteLicenseOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v2 everoute license order by input
func (m V2EverouteLicenseOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV2EverouteLicenseOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v2 everoute license order by input based on context it is used
func (m V2EverouteLicenseOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}