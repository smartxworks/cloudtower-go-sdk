// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NvmfSubsystemCommonParams nvmf subsystem common params
//
// swagger:model NvmfSubsystemCommonParams
type NvmfSubsystemCommonParams struct {

	// bps
	Bps *int64 `json:"bps,omitempty"`

	// bps max
	BpsMax *int64 `json:"bps_max,omitempty"`

	// bps max length
	BpsMaxLength *int64 `json:"bps_max_length,omitempty"`

	// bps max unit
	BpsMaxUnit *BPSUnit `json:"bps_max_unit,omitempty"`

	// bps rd
	BpsRd *int64 `json:"bps_rd,omitempty"`

	// bps rd max
	BpsRdMax *int64 `json:"bps_rd_max,omitempty"`

	// bps rd max length
	BpsRdMaxLength *int64 `json:"bps_rd_max_length,omitempty"`

	// bps rd max unit
	BpsRdMaxUnit *BPSUnit `json:"bps_rd_max_unit,omitempty"`

	// bps rd unit
	BpsRdUnit *BPSUnit `json:"bps_rd_unit,omitempty"`

	// bps unit
	BpsUnit *BPSUnit `json:"bps_unit,omitempty"`

	// bps wr
	BpsWr *int64 `json:"bps_wr,omitempty"`

	// bps wr max
	BpsWrMax *int64 `json:"bps_wr_max,omitempty"`

	// bps wr max length
	BpsWrMaxLength *int64 `json:"bps_wr_max_length,omitempty"`

	// bps wr max unit
	BpsWrMaxUnit *BPSUnit `json:"bps_wr_max_unit,omitempty"`

	// bps wr unit
	BpsWrUnit *BPSUnit `json:"bps_wr_unit,omitempty"`

	// description
	Description *string `json:"description,omitempty"`

	// iops
	Iops *int64 `json:"iops,omitempty"`

	// iops max
	IopsMax *int64 `json:"iops_max,omitempty"`

	// iops max length
	IopsMaxLength *int64 `json:"iops_max_length,omitempty"`

	// iops rd
	IopsRd *int64 `json:"iops_rd,omitempty"`

	// iops rd max
	IopsRdMax *int64 `json:"iops_rd_max,omitempty"`

	// iops rd max length
	IopsRdMaxLength *int64 `json:"iops_rd_max_length,omitempty"`

	// iops wr
	IopsWr *int64 `json:"iops_wr,omitempty"`

	// iops wr max
	IopsWrMax *int64 `json:"iops_wr_max,omitempty"`

	// iops wr max length
	IopsWrMaxLength *int64 `json:"iops_wr_max_length,omitempty"`

	// ip whitelist
	IPWhitelist *string `json:"ip_whitelist,omitempty"`

	// nqn whitelist
	NqnWhitelist *string `json:"nqn_whitelist,omitempty"`

	MarshalOpts *NvmfSubsystemCommonParamsMarshalOpts `json:"-"`
}

type NvmfSubsystemCommonParamsMarshalOpts struct {
	Bps_Explicit_Null_When_Empty bool

	BpsMax_Explicit_Null_When_Empty bool

	BpsMaxLength_Explicit_Null_When_Empty bool

	BpsMaxUnit_Explicit_Null_When_Empty bool

	BpsRd_Explicit_Null_When_Empty bool

	BpsRdMax_Explicit_Null_When_Empty bool

	BpsRdMaxLength_Explicit_Null_When_Empty bool

	BpsRdMaxUnit_Explicit_Null_When_Empty bool

	BpsRdUnit_Explicit_Null_When_Empty bool

	BpsUnit_Explicit_Null_When_Empty bool

	BpsWr_Explicit_Null_When_Empty bool

	BpsWrMax_Explicit_Null_When_Empty bool

	BpsWrMaxLength_Explicit_Null_When_Empty bool

	BpsWrMaxUnit_Explicit_Null_When_Empty bool

	BpsWrUnit_Explicit_Null_When_Empty bool

	Description_Explicit_Null_When_Empty bool

	Iops_Explicit_Null_When_Empty bool

	IopsMax_Explicit_Null_When_Empty bool

	IopsMaxLength_Explicit_Null_When_Empty bool

	IopsRd_Explicit_Null_When_Empty bool

	IopsRdMax_Explicit_Null_When_Empty bool

	IopsRdMaxLength_Explicit_Null_When_Empty bool

	IopsWr_Explicit_Null_When_Empty bool

	IopsWrMax_Explicit_Null_When_Empty bool

	IopsWrMaxLength_Explicit_Null_When_Empty bool

	IPWhitelist_Explicit_Null_When_Empty bool

	NqnWhitelist_Explicit_Null_When_Empty bool
}

func (m NvmfSubsystemCommonParams) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field bps
	if m.Bps != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bps\":")
		bytes, err := swag.WriteJSON(m.Bps)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Bps_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bps\":null")
		first = false
	}

	// handle nullable field bps_max
	if m.BpsMax != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bps_max\":")
		bytes, err := swag.WriteJSON(m.BpsMax)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.BpsMax_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bps_max\":null")
		first = false
	}

	// handle nullable field bps_max_length
	if m.BpsMaxLength != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bps_max_length\":")
		bytes, err := swag.WriteJSON(m.BpsMaxLength)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.BpsMaxLength_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bps_max_length\":null")
		first = false
	}

	// handle nullable field bps_max_unit
	if m.BpsMaxUnit != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bps_max_unit\":")
		bytes, err := swag.WriteJSON(m.BpsMaxUnit)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.BpsMaxUnit_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bps_max_unit\":null")
		first = false
	}

	// handle nullable field bps_rd
	if m.BpsRd != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bps_rd\":")
		bytes, err := swag.WriteJSON(m.BpsRd)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.BpsRd_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bps_rd\":null")
		first = false
	}

	// handle nullable field bps_rd_max
	if m.BpsRdMax != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bps_rd_max\":")
		bytes, err := swag.WriteJSON(m.BpsRdMax)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.BpsRdMax_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bps_rd_max\":null")
		first = false
	}

	// handle nullable field bps_rd_max_length
	if m.BpsRdMaxLength != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bps_rd_max_length\":")
		bytes, err := swag.WriteJSON(m.BpsRdMaxLength)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.BpsRdMaxLength_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bps_rd_max_length\":null")
		first = false
	}

	// handle nullable field bps_rd_max_unit
	if m.BpsRdMaxUnit != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bps_rd_max_unit\":")
		bytes, err := swag.WriteJSON(m.BpsRdMaxUnit)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.BpsRdMaxUnit_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bps_rd_max_unit\":null")
		first = false
	}

	// handle nullable field bps_rd_unit
	if m.BpsRdUnit != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bps_rd_unit\":")
		bytes, err := swag.WriteJSON(m.BpsRdUnit)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.BpsRdUnit_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bps_rd_unit\":null")
		first = false
	}

	// handle nullable field bps_unit
	if m.BpsUnit != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bps_unit\":")
		bytes, err := swag.WriteJSON(m.BpsUnit)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.BpsUnit_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bps_unit\":null")
		first = false
	}

	// handle nullable field bps_wr
	if m.BpsWr != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bps_wr\":")
		bytes, err := swag.WriteJSON(m.BpsWr)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.BpsWr_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bps_wr\":null")
		first = false
	}

	// handle nullable field bps_wr_max
	if m.BpsWrMax != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bps_wr_max\":")
		bytes, err := swag.WriteJSON(m.BpsWrMax)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.BpsWrMax_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bps_wr_max\":null")
		first = false
	}

	// handle nullable field bps_wr_max_length
	if m.BpsWrMaxLength != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bps_wr_max_length\":")
		bytes, err := swag.WriteJSON(m.BpsWrMaxLength)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.BpsWrMaxLength_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bps_wr_max_length\":null")
		first = false
	}

	// handle nullable field bps_wr_max_unit
	if m.BpsWrMaxUnit != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bps_wr_max_unit\":")
		bytes, err := swag.WriteJSON(m.BpsWrMaxUnit)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.BpsWrMaxUnit_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bps_wr_max_unit\":null")
		first = false
	}

	// handle nullable field bps_wr_unit
	if m.BpsWrUnit != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bps_wr_unit\":")
		bytes, err := swag.WriteJSON(m.BpsWrUnit)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.BpsWrUnit_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bps_wr_unit\":null")
		first = false
	}

	// handle nullable field description
	if m.Description != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description\":")
		bytes, err := swag.WriteJSON(m.Description)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Description_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"description\":null")
		first = false
	}

	// handle nullable field iops
	if m.Iops != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"iops\":")
		bytes, err := swag.WriteJSON(m.Iops)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Iops_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"iops\":null")
		first = false
	}

	// handle nullable field iops_max
	if m.IopsMax != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"iops_max\":")
		bytes, err := swag.WriteJSON(m.IopsMax)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IopsMax_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"iops_max\":null")
		first = false
	}

	// handle nullable field iops_max_length
	if m.IopsMaxLength != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"iops_max_length\":")
		bytes, err := swag.WriteJSON(m.IopsMaxLength)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IopsMaxLength_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"iops_max_length\":null")
		first = false
	}

	// handle nullable field iops_rd
	if m.IopsRd != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"iops_rd\":")
		bytes, err := swag.WriteJSON(m.IopsRd)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IopsRd_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"iops_rd\":null")
		first = false
	}

	// handle nullable field iops_rd_max
	if m.IopsRdMax != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"iops_rd_max\":")
		bytes, err := swag.WriteJSON(m.IopsRdMax)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IopsRdMax_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"iops_rd_max\":null")
		first = false
	}

	// handle nullable field iops_rd_max_length
	if m.IopsRdMaxLength != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"iops_rd_max_length\":")
		bytes, err := swag.WriteJSON(m.IopsRdMaxLength)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IopsRdMaxLength_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"iops_rd_max_length\":null")
		first = false
	}

	// handle nullable field iops_wr
	if m.IopsWr != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"iops_wr\":")
		bytes, err := swag.WriteJSON(m.IopsWr)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IopsWr_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"iops_wr\":null")
		first = false
	}

	// handle nullable field iops_wr_max
	if m.IopsWrMax != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"iops_wr_max\":")
		bytes, err := swag.WriteJSON(m.IopsWrMax)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IopsWrMax_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"iops_wr_max\":null")
		first = false
	}

	// handle nullable field iops_wr_max_length
	if m.IopsWrMaxLength != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"iops_wr_max_length\":")
		bytes, err := swag.WriteJSON(m.IopsWrMaxLength)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IopsWrMaxLength_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"iops_wr_max_length\":null")
		first = false
	}

	// handle nullable field ip_whitelist
	if m.IPWhitelist != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ip_whitelist\":")
		bytes, err := swag.WriteJSON(m.IPWhitelist)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IPWhitelist_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"ip_whitelist\":null")
		first = false
	}

	// handle nullable field nqn_whitelist
	if m.NqnWhitelist != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"nqn_whitelist\":")
		bytes, err := swag.WriteJSON(m.NqnWhitelist)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NqnWhitelist_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"nqn_whitelist\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this nvmf subsystem common params
func (m *NvmfSubsystemCommonParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBpsMaxUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBpsRdMaxUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBpsRdUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBpsUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBpsWrMaxUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBpsWrUnit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmfSubsystemCommonParams) validateBpsMaxUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.BpsMaxUnit) { // not required
		return nil
	}

	if m.BpsMaxUnit != nil {
		if err := m.BpsMaxUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bps_max_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bps_max_unit")
			}
			return err
		}
	}

	return nil
}

func (m *NvmfSubsystemCommonParams) validateBpsRdMaxUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.BpsRdMaxUnit) { // not required
		return nil
	}

	if m.BpsRdMaxUnit != nil {
		if err := m.BpsRdMaxUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bps_rd_max_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bps_rd_max_unit")
			}
			return err
		}
	}

	return nil
}

func (m *NvmfSubsystemCommonParams) validateBpsRdUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.BpsRdUnit) { // not required
		return nil
	}

	if m.BpsRdUnit != nil {
		if err := m.BpsRdUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bps_rd_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bps_rd_unit")
			}
			return err
		}
	}

	return nil
}

func (m *NvmfSubsystemCommonParams) validateBpsUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.BpsUnit) { // not required
		return nil
	}

	if m.BpsUnit != nil {
		if err := m.BpsUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bps_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bps_unit")
			}
			return err
		}
	}

	return nil
}

func (m *NvmfSubsystemCommonParams) validateBpsWrMaxUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.BpsWrMaxUnit) { // not required
		return nil
	}

	if m.BpsWrMaxUnit != nil {
		if err := m.BpsWrMaxUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bps_wr_max_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bps_wr_max_unit")
			}
			return err
		}
	}

	return nil
}

func (m *NvmfSubsystemCommonParams) validateBpsWrUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.BpsWrUnit) { // not required
		return nil
	}

	if m.BpsWrUnit != nil {
		if err := m.BpsWrUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bps_wr_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bps_wr_unit")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nvmf subsystem common params based on the context it is used
func (m *NvmfSubsystemCommonParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBpsMaxUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBpsRdMaxUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBpsRdUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBpsUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBpsWrMaxUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBpsWrUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmfSubsystemCommonParams) contextValidateBpsMaxUnit(ctx context.Context, formats strfmt.Registry) error {

	if m.BpsMaxUnit != nil {
		if err := m.BpsMaxUnit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bps_max_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bps_max_unit")
			}
			return err
		}
	}

	return nil
}

func (m *NvmfSubsystemCommonParams) contextValidateBpsRdMaxUnit(ctx context.Context, formats strfmt.Registry) error {

	if m.BpsRdMaxUnit != nil {
		if err := m.BpsRdMaxUnit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bps_rd_max_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bps_rd_max_unit")
			}
			return err
		}
	}

	return nil
}

func (m *NvmfSubsystemCommonParams) contextValidateBpsRdUnit(ctx context.Context, formats strfmt.Registry) error {

	if m.BpsRdUnit != nil {
		if err := m.BpsRdUnit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bps_rd_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bps_rd_unit")
			}
			return err
		}
	}

	return nil
}

func (m *NvmfSubsystemCommonParams) contextValidateBpsUnit(ctx context.Context, formats strfmt.Registry) error {

	if m.BpsUnit != nil {
		if err := m.BpsUnit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bps_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bps_unit")
			}
			return err
		}
	}

	return nil
}

func (m *NvmfSubsystemCommonParams) contextValidateBpsWrMaxUnit(ctx context.Context, formats strfmt.Registry) error {

	if m.BpsWrMaxUnit != nil {
		if err := m.BpsWrMaxUnit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bps_wr_max_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bps_wr_max_unit")
			}
			return err
		}
	}

	return nil
}

func (m *NvmfSubsystemCommonParams) contextValidateBpsWrUnit(ctx context.Context, formats strfmt.Registry) error {

	if m.BpsWrUnit != nil {
		if err := m.BpsWrUnit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bps_wr_unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bps_wr_unit")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NvmfSubsystemCommonParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmfSubsystemCommonParams) UnmarshalBinary(b []byte) error {
	var res NvmfSubsystemCommonParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
