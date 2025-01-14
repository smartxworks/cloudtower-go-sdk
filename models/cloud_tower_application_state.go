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

// CloudTowerApplicationState cloud tower application state
//
// swagger:model CloudTowerApplicationState
type CloudTowerApplicationState string

func NewCloudTowerApplicationState(value CloudTowerApplicationState) *CloudTowerApplicationState {
	return &value
}

// Pointer returns a pointer to a freshly-allocated CloudTowerApplicationState.
func (m CloudTowerApplicationState) Pointer() *CloudTowerApplicationState {
	return &m
}

const (

	// CloudTowerApplicationStateCONFIGUREFAILED captures enum value "CONFIGURE_FAILED"
	CloudTowerApplicationStateCONFIGUREFAILED CloudTowerApplicationState = "CONFIGURE_FAILED"

	// CloudTowerApplicationStateCONFIGURING captures enum value "CONFIGURING"
	CloudTowerApplicationStateCONFIGURING CloudTowerApplicationState = "CONFIGURING"

	// CloudTowerApplicationStateERROR captures enum value "ERROR"
	CloudTowerApplicationStateERROR CloudTowerApplicationState = "ERROR"

	// CloudTowerApplicationStateINSTALLING captures enum value "INSTALLING"
	CloudTowerApplicationStateINSTALLING CloudTowerApplicationState = "INSTALLING"

	// CloudTowerApplicationStateINSTALLFAILED captures enum value "INSTALL_FAILED"
	CloudTowerApplicationStateINSTALLFAILED CloudTowerApplicationState = "INSTALL_FAILED"

	// CloudTowerApplicationStateRUNNING captures enum value "RUNNING"
	CloudTowerApplicationStateRUNNING CloudTowerApplicationState = "RUNNING"

	// CloudTowerApplicationStateSCALEFAILED captures enum value "SCALE_FAILED"
	CloudTowerApplicationStateSCALEFAILED CloudTowerApplicationState = "SCALE_FAILED"

	// CloudTowerApplicationStateSCALING captures enum value "SCALING"
	CloudTowerApplicationStateSCALING CloudTowerApplicationState = "SCALING"

	// CloudTowerApplicationStateTERMINATEFAILED captures enum value "TERMINATE_FAILED"
	CloudTowerApplicationStateTERMINATEFAILED CloudTowerApplicationState = "TERMINATE_FAILED"

	// CloudTowerApplicationStateTERMINATING captures enum value "TERMINATING"
	CloudTowerApplicationStateTERMINATING CloudTowerApplicationState = "TERMINATING"

	// CloudTowerApplicationStateUPGRADEFAILED captures enum value "UPGRADE_FAILED"
	CloudTowerApplicationStateUPGRADEFAILED CloudTowerApplicationState = "UPGRADE_FAILED"

	// CloudTowerApplicationStateUPGRADING captures enum value "UPGRADING"
	CloudTowerApplicationStateUPGRADING CloudTowerApplicationState = "UPGRADING"
)

// for schema
var cloudTowerApplicationStateEnum []interface{}

func init() {
	var res []CloudTowerApplicationState
	if err := json.Unmarshal([]byte(`["CONFIGURE_FAILED","CONFIGURING","ERROR","INSTALLING","INSTALL_FAILED","RUNNING","SCALE_FAILED","SCALING","TERMINATE_FAILED","TERMINATING","UPGRADE_FAILED","UPGRADING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cloudTowerApplicationStateEnum = append(cloudTowerApplicationStateEnum, v)
	}
}

func (m CloudTowerApplicationState) validateCloudTowerApplicationStateEnum(path, location string, value CloudTowerApplicationState) error {
	if err := validate.EnumCase(path, location, value, cloudTowerApplicationStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this cloud tower application state
func (m CloudTowerApplicationState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCloudTowerApplicationStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this cloud tower application state based on context it is used
func (m CloudTowerApplicationState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
