// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LogCollection log collection
//
// swagger:model LogCollection
type LogCollection struct {

	// cluster
	// Required: true
	Cluster *NestedCluster `json:"cluster"`

	// groups
	// Required: true
	Groups []string `json:"groups"`

	// hosts
	Hosts []*NestedHost `json:"hosts,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// log ended at
	// Required: true
	LogEndedAt *string `json:"log_ended_at"`

	// log started at
	// Required: true
	LogStartedAt *string `json:"log_started_at"`

	// owner
	// Required: true
	Owner *string `json:"owner"`

	// path
	// Required: true
	Path *string `json:"path"`

	// progress
	// Required: true
	Progress *float64 `json:"progress"`

	// services
	// Required: true
	Services []string `json:"services"`

	// size
	// Required: true
	Size *int64 `json:"size"`

	// started at
	// Required: true
	StartedAt *string `json:"started_at"`

	// status
	// Required: true
	Status *LogCollectionStatus `json:"status"`

	// witness
	Witness *NestedWitness `json:"witness,omitempty"`
}

// Validate validates this log collection
func (m *LogCollection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHosts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogEndedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProgress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWitness(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogCollection) validateCluster(formats strfmt.Registry) error {

	if err := validate.Required("cluster", "body", m.Cluster); err != nil {
		return err
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *LogCollection) validateGroups(formats strfmt.Registry) error {

	if err := validate.Required("groups", "body", m.Groups); err != nil {
		return err
	}

	return nil
}

func (m *LogCollection) validateHosts(formats strfmt.Registry) error {
	if swag.IsZero(m.Hosts) { // not required
		return nil
	}

	for i := 0; i < len(m.Hosts); i++ {
		if swag.IsZero(m.Hosts[i]) { // not required
			continue
		}

		if m.Hosts[i] != nil {
			if err := m.Hosts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hosts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("hosts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LogCollection) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *LogCollection) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *LogCollection) validateLogEndedAt(formats strfmt.Registry) error {

	if err := validate.Required("log_ended_at", "body", m.LogEndedAt); err != nil {
		return err
	}

	return nil
}

func (m *LogCollection) validateLogStartedAt(formats strfmt.Registry) error {

	if err := validate.Required("log_started_at", "body", m.LogStartedAt); err != nil {
		return err
	}

	return nil
}

func (m *LogCollection) validateOwner(formats strfmt.Registry) error {

	if err := validate.Required("owner", "body", m.Owner); err != nil {
		return err
	}

	return nil
}

func (m *LogCollection) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

func (m *LogCollection) validateProgress(formats strfmt.Registry) error {

	if err := validate.Required("progress", "body", m.Progress); err != nil {
		return err
	}

	return nil
}

func (m *LogCollection) validateServices(formats strfmt.Registry) error {

	if err := validate.Required("services", "body", m.Services); err != nil {
		return err
	}

	return nil
}

func (m *LogCollection) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *LogCollection) validateStartedAt(formats strfmt.Registry) error {

	if err := validate.Required("started_at", "body", m.StartedAt); err != nil {
		return err
	}

	return nil
}

func (m *LogCollection) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *LogCollection) validateWitness(formats strfmt.Registry) error {
	if swag.IsZero(m.Witness) { // not required
		return nil
	}

	if m.Witness != nil {
		if err := m.Witness.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("witness")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("witness")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this log collection based on the context it is used
func (m *LogCollection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHosts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWitness(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogCollection) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {
		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *LogCollection) contextValidateHosts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Hosts); i++ {

		if m.Hosts[i] != nil {
			if err := m.Hosts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hosts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("hosts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LogCollection) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *LogCollection) contextValidateWitness(ctx context.Context, formats strfmt.Registry) error {

	if m.Witness != nil {
		if err := m.Witness.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("witness")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("witness")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LogCollection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogCollection) UnmarshalBinary(b []byte) error {
	var res LogCollection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
