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

// BackupRestorePointCreation backup restore point creation
//
// swagger:model BackupRestorePointCreation
type BackupRestorePointCreation string

func NewBackupRestorePointCreation(value BackupRestorePointCreation) *BackupRestorePointCreation {
	return &value
}

// Pointer returns a pointer to a freshly-allocated BackupRestorePointCreation.
func (m BackupRestorePointCreation) Pointer() *BackupRestorePointCreation {
	return &m
}

const (

	// BackupRestorePointCreationAUTO captures enum value "AUTO"
	BackupRestorePointCreationAUTO BackupRestorePointCreation = "AUTO"

	// BackupRestorePointCreationFAILBACK captures enum value "FAILBACK"
	BackupRestorePointCreationFAILBACK BackupRestorePointCreation = "FAILBACK"

	// BackupRestorePointCreationFAILOVER captures enum value "FAILOVER"
	BackupRestorePointCreationFAILOVER BackupRestorePointCreation = "FAILOVER"

	// BackupRestorePointCreationFAILOVERTEST captures enum value "FAILOVER_TEST"
	BackupRestorePointCreationFAILOVERTEST BackupRestorePointCreation = "FAILOVER_TEST"

	// BackupRestorePointCreationMANUAL captures enum value "MANUAL"
	BackupRestorePointCreationMANUAL BackupRestorePointCreation = "MANUAL"
)

// for schema
var backupRestorePointCreationEnum []interface{}

func init() {
	var res []BackupRestorePointCreation
	if err := json.Unmarshal([]byte(`["AUTO","FAILBACK","FAILOVER","FAILOVER_TEST","MANUAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backupRestorePointCreationEnum = append(backupRestorePointCreationEnum, v)
	}
}

func (m BackupRestorePointCreation) validateBackupRestorePointCreationEnum(path, location string, value BackupRestorePointCreation) error {
	if err := validate.EnumCase(path, location, value, backupRestorePointCreationEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this backup restore point creation
func (m BackupRestorePointCreation) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBackupRestorePointCreationEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this backup restore point creation based on context it is used
func (m BackupRestorePointCreation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
