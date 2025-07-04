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

// ExternalCloudTowerStatus external cloud tower status
//
// swagger:model ExternalCloudTowerStatus
type ExternalCloudTowerStatus string

func NewExternalCloudTowerStatus(value ExternalCloudTowerStatus) *ExternalCloudTowerStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ExternalCloudTowerStatus.
func (m ExternalCloudTowerStatus) Pointer() *ExternalCloudTowerStatus {
	return &m
}

const (

	// ExternalCloudTowerStatusCONNECTED captures enum value "CONNECTED"
	ExternalCloudTowerStatusCONNECTED ExternalCloudTowerStatus = "CONNECTED"

	// ExternalCloudTowerStatusCONNECTING captures enum value "CONNECTING"
	ExternalCloudTowerStatusCONNECTING ExternalCloudTowerStatus = "CONNECTING"

	// ExternalCloudTowerStatusDISCONNECTED captures enum value "DISCONNECTED"
	ExternalCloudTowerStatusDISCONNECTED ExternalCloudTowerStatus = "DISCONNECTED"
)

// for schema
var externalCloudTowerStatusEnum []interface{}

func init() {
	var res []ExternalCloudTowerStatus
	if err := json.Unmarshal([]byte(`["CONNECTED","CONNECTING","DISCONNECTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		externalCloudTowerStatusEnum = append(externalCloudTowerStatusEnum, v)
	}
}

func (m ExternalCloudTowerStatus) validateExternalCloudTowerStatusEnum(path, location string, value ExternalCloudTowerStatus) error {
	if err := validate.EnumCase(path, location, value, externalCloudTowerStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this external cloud tower status
func (m ExternalCloudTowerStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateExternalCloudTowerStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this external cloud tower status based on context it is used
func (m ExternalCloudTowerStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
