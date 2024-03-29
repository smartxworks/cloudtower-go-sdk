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

// AlertNotifierOrderByInput alert notifier order by input
//
// swagger:model AlertNotifierOrderByInput
type AlertNotifierOrderByInput string

func NewAlertNotifierOrderByInput(value AlertNotifierOrderByInput) *AlertNotifierOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AlertNotifierOrderByInput.
func (m AlertNotifierOrderByInput) Pointer() *AlertNotifierOrderByInput {
	return &m
}

const (

	// AlertNotifierOrderByInputDisabledASC captures enum value "disabled_ASC"
	AlertNotifierOrderByInputDisabledASC AlertNotifierOrderByInput = "disabled_ASC"

	// AlertNotifierOrderByInputDisabledDESC captures enum value "disabled_DESC"
	AlertNotifierOrderByInputDisabledDESC AlertNotifierOrderByInput = "disabled_DESC"

	// AlertNotifierOrderByInputEmailFromASC captures enum value "email_from_ASC"
	AlertNotifierOrderByInputEmailFromASC AlertNotifierOrderByInput = "email_from_ASC"

	// AlertNotifierOrderByInputEmailFromDESC captures enum value "email_from_DESC"
	AlertNotifierOrderByInputEmailFromDESC AlertNotifierOrderByInput = "email_from_DESC"

	// AlertNotifierOrderByInputEntityAsyncStatusASC captures enum value "entityAsyncStatus_ASC"
	AlertNotifierOrderByInputEntityAsyncStatusASC AlertNotifierOrderByInput = "entityAsyncStatus_ASC"

	// AlertNotifierOrderByInputEntityAsyncStatusDESC captures enum value "entityAsyncStatus_DESC"
	AlertNotifierOrderByInputEntityAsyncStatusDESC AlertNotifierOrderByInput = "entityAsyncStatus_DESC"

	// AlertNotifierOrderByInputIDASC captures enum value "id_ASC"
	AlertNotifierOrderByInputIDASC AlertNotifierOrderByInput = "id_ASC"

	// AlertNotifierOrderByInputIDDESC captures enum value "id_DESC"
	AlertNotifierOrderByInputIDDESC AlertNotifierOrderByInput = "id_DESC"

	// AlertNotifierOrderByInputLanguageCodeASC captures enum value "language_code_ASC"
	AlertNotifierOrderByInputLanguageCodeASC AlertNotifierOrderByInput = "language_code_ASC"

	// AlertNotifierOrderByInputLanguageCodeDESC captures enum value "language_code_DESC"
	AlertNotifierOrderByInputLanguageCodeDESC AlertNotifierOrderByInput = "language_code_DESC"

	// AlertNotifierOrderByInputNameASC captures enum value "name_ASC"
	AlertNotifierOrderByInputNameASC AlertNotifierOrderByInput = "name_ASC"

	// AlertNotifierOrderByInputNameDESC captures enum value "name_DESC"
	AlertNotifierOrderByInputNameDESC AlertNotifierOrderByInput = "name_DESC"

	// AlertNotifierOrderByInputSecurityModeASC captures enum value "security_mode_ASC"
	AlertNotifierOrderByInputSecurityModeASC AlertNotifierOrderByInput = "security_mode_ASC"

	// AlertNotifierOrderByInputSecurityModeDESC captures enum value "security_mode_DESC"
	AlertNotifierOrderByInputSecurityModeDESC AlertNotifierOrderByInput = "security_mode_DESC"

	// AlertNotifierOrderByInputSMTPServerHostASC captures enum value "smtp_server_host_ASC"
	AlertNotifierOrderByInputSMTPServerHostASC AlertNotifierOrderByInput = "smtp_server_host_ASC"

	// AlertNotifierOrderByInputSMTPServerHostDESC captures enum value "smtp_server_host_DESC"
	AlertNotifierOrderByInputSMTPServerHostDESC AlertNotifierOrderByInput = "smtp_server_host_DESC"

	// AlertNotifierOrderByInputSMTPServerPortASC captures enum value "smtp_server_port_ASC"
	AlertNotifierOrderByInputSMTPServerPortASC AlertNotifierOrderByInput = "smtp_server_port_ASC"

	// AlertNotifierOrderByInputSMTPServerPortDESC captures enum value "smtp_server_port_DESC"
	AlertNotifierOrderByInputSMTPServerPortDESC AlertNotifierOrderByInput = "smtp_server_port_DESC"

	// AlertNotifierOrderByInputUsernameASC captures enum value "username_ASC"
	AlertNotifierOrderByInputUsernameASC AlertNotifierOrderByInput = "username_ASC"

	// AlertNotifierOrderByInputUsernameDESC captures enum value "username_DESC"
	AlertNotifierOrderByInputUsernameDESC AlertNotifierOrderByInput = "username_DESC"
)

// for schema
var alertNotifierOrderByInputEnum []interface{}

func init() {
	var res []AlertNotifierOrderByInput
	if err := json.Unmarshal([]byte(`["disabled_ASC","disabled_DESC","email_from_ASC","email_from_DESC","entityAsyncStatus_ASC","entityAsyncStatus_DESC","id_ASC","id_DESC","language_code_ASC","language_code_DESC","name_ASC","name_DESC","security_mode_ASC","security_mode_DESC","smtp_server_host_ASC","smtp_server_host_DESC","smtp_server_port_ASC","smtp_server_port_DESC","username_ASC","username_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		alertNotifierOrderByInputEnum = append(alertNotifierOrderByInputEnum, v)
	}
}

func (m AlertNotifierOrderByInput) validateAlertNotifierOrderByInputEnum(path, location string, value AlertNotifierOrderByInput) error {
	if err := validate.EnumCase(path, location, value, alertNotifierOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this alert notifier order by input
func (m AlertNotifierOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAlertNotifierOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this alert notifier order by input based on context it is used
func (m AlertNotifierOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
