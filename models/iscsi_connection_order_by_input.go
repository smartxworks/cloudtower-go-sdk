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

// IscsiConnectionOrderByInput iscsi connection order by input
//
// swagger:model IscsiConnectionOrderByInput
type IscsiConnectionOrderByInput string

func NewIscsiConnectionOrderByInput(value IscsiConnectionOrderByInput) *IscsiConnectionOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated IscsiConnectionOrderByInput.
func (m IscsiConnectionOrderByInput) Pointer() *IscsiConnectionOrderByInput {
	return &m
}

const (

	// IscsiConnectionOrderByInputClientPortASC captures enum value "client_port_ASC"
	IscsiConnectionOrderByInputClientPortASC IscsiConnectionOrderByInput = "client_port_ASC"

	// IscsiConnectionOrderByInputClientPortDESC captures enum value "client_port_DESC"
	IscsiConnectionOrderByInputClientPortDESC IscsiConnectionOrderByInput = "client_port_DESC"

	// IscsiConnectionOrderByInputIDASC captures enum value "id_ASC"
	IscsiConnectionOrderByInputIDASC IscsiConnectionOrderByInput = "id_ASC"

	// IscsiConnectionOrderByInputIDDESC captures enum value "id_DESC"
	IscsiConnectionOrderByInputIDDESC IscsiConnectionOrderByInput = "id_DESC"

	// IscsiConnectionOrderByInputInitiatorIPASC captures enum value "initiator_ip_ASC"
	IscsiConnectionOrderByInputInitiatorIPASC IscsiConnectionOrderByInput = "initiator_ip_ASC"

	// IscsiConnectionOrderByInputInitiatorIPDESC captures enum value "initiator_ip_DESC"
	IscsiConnectionOrderByInputInitiatorIPDESC IscsiConnectionOrderByInput = "initiator_ip_DESC"

	// IscsiConnectionOrderByInputTrTypeASC captures enum value "tr_type_ASC"
	IscsiConnectionOrderByInputTrTypeASC IscsiConnectionOrderByInput = "tr_type_ASC"

	// IscsiConnectionOrderByInputTrTypeDESC captures enum value "tr_type_DESC"
	IscsiConnectionOrderByInputTrTypeDESC IscsiConnectionOrderByInput = "tr_type_DESC"

	// IscsiConnectionOrderByInputTypeASC captures enum value "type_ASC"
	IscsiConnectionOrderByInputTypeASC IscsiConnectionOrderByInput = "type_ASC"

	// IscsiConnectionOrderByInputTypeDESC captures enum value "type_DESC"
	IscsiConnectionOrderByInputTypeDESC IscsiConnectionOrderByInput = "type_DESC"
)

// for schema
var iscsiConnectionOrderByInputEnum []interface{}

func init() {
	var res []IscsiConnectionOrderByInput
	if err := json.Unmarshal([]byte(`["client_port_ASC","client_port_DESC","id_ASC","id_DESC","initiator_ip_ASC","initiator_ip_DESC","tr_type_ASC","tr_type_DESC","type_ASC","type_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iscsiConnectionOrderByInputEnum = append(iscsiConnectionOrderByInputEnum, v)
	}
}

func (m IscsiConnectionOrderByInput) validateIscsiConnectionOrderByInputEnum(path, location string, value IscsiConnectionOrderByInput) error {
	if err := validate.EnumCase(path, location, value, iscsiConnectionOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this iscsi connection order by input
func (m IscsiConnectionOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateIscsiConnectionOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this iscsi connection order by input based on context it is used
func (m IscsiConnectionOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
