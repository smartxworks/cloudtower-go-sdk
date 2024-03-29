// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// Priority priority
//
// swagger:model Priority
type Priority float64

// Validate validates this priority
func (m Priority) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Minimum("", "body", float64(m), 0, false); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this priority based on context it is used
func (m Priority) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
