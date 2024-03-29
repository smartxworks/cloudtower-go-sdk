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

// IommuStatus iommu status
//
// swagger:model IommuStatus
type IommuStatus string

func NewIommuStatus(value IommuStatus) *IommuStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated IommuStatus.
func (m IommuStatus) Pointer() *IommuStatus {
	return &m
}

const (

	// IommuStatusDISABLE captures enum value "DISABLE"
	IommuStatusDISABLE IommuStatus = "DISABLE"

	// IommuStatusENABLE captures enum value "ENABLE"
	IommuStatusENABLE IommuStatus = "ENABLE"

	// IommuStatusNEEDREBOOT captures enum value "NEED_REBOOT"
	IommuStatusNEEDREBOOT IommuStatus = "NEED_REBOOT"
)

// for schema
var iommuStatusEnum []interface{}

func init() {
	var res []IommuStatus
	if err := json.Unmarshal([]byte(`["DISABLE","ENABLE","NEED_REBOOT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iommuStatusEnum = append(iommuStatusEnum, v)
	}
}

func (m IommuStatus) validateIommuStatusEnum(path, location string, value IommuStatus) error {
	if err := validate.EnumCase(path, location, value, iommuStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this iommu status
func (m IommuStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateIommuStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this iommu status based on context it is used
func (m IommuStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
