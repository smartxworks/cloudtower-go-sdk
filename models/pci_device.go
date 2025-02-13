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
	"github.com/go-openapi/validate"
)

// PciDevice pci device
//
// swagger:model PciDevice
type PciDevice struct {

	// bus
	// Required: true
	Bus *string `json:"bus"`

	// bus location
	// Required: true
	BusLocation *string `json:"bus_location"`

	// class code
	// Required: true
	ClassCode *string `json:"class_code"`

	// device type
	DeviceType *string `json:"device_type,omitempty"`

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// function num
	// Required: true
	FunctionNum *string `json:"function_num"`

	// host
	// Required: true
	Host *NestedHost `json:"host"`

	// id
	// Required: true
	ID *string `json:"id"`

	// iommu status
	IommuStatus *IommuStatus `json:"iommu_status,omitempty"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// mdev state
	MdevState *PciDeviceMdevState `json:"mdev_state,omitempty"`

	// mdev type id
	MdevTypeID *string `json:"mdev_type_id,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// partitioning state
	PartitioningState *PciDevicePartitionState `json:"partitioning_state,omitempty"`

	// product id
	// Required: true
	ProductID *string `json:"product_id"`

	// slot
	// Required: true
	Slot *string `json:"slot"`

	// sriov state
	SriovState *PciDeviceSriovState `json:"sriov_state,omitempty"`

	// subsystem product id
	// Required: true
	SubsystemProductID *string `json:"subsystem_product_id"`

	// subsystem vendor id
	// Required: true
	SubsystemVendorID *string `json:"subsystem_vendor_id"`

	// total mdev num
	TotalMdevNum *int32 `json:"total_mdev_num,omitempty"`

	// total partitioning num
	TotalPartitioningNum *int32 `json:"total_partitioning_num,omitempty"`

	// total vf num
	TotalVfNum *int32 `json:"total_vf_num,omitempty"`

	// usage type
	UsageType *PciDeviceType `json:"usage_type,omitempty"`

	// used mdev num
	UsedMdevNum *int32 `json:"used_mdev_num,omitempty"`

	// used partitioning num
	UsedPartitioningNum *int32 `json:"used_partitioning_num,omitempty"`

	// used vf num
	UsedVfNum *int32 `json:"used_vf_num,omitempty"`

	// user usage
	UserUsage *PciDeviceUsage `json:"user_usage,omitempty"`

	// vendor id
	// Required: true
	VendorID *string `json:"vendor_id"`

	// verdor name
	// Required: true
	VerdorName *string `json:"verdor_name"`

	MarshalOpts *PciDeviceMarshalOpts `json:"-"`
}

type PciDeviceMarshalOpts struct {
	Bus_Explicit_Null_When_Empty bool

	BusLocation_Explicit_Null_When_Empty bool

	ClassCode_Explicit_Null_When_Empty bool

	DeviceType_Explicit_Null_When_Empty bool

	EntityAsyncStatus_Explicit_Null_When_Empty bool

	FunctionNum_Explicit_Null_When_Empty bool

	Host_Explicit_Null_When_Empty bool

	ID_Explicit_Null_When_Empty bool

	IommuStatus_Explicit_Null_When_Empty bool

	LocalID_Explicit_Null_When_Empty bool

	MdevState_Explicit_Null_When_Empty bool

	MdevTypeID_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	PartitioningState_Explicit_Null_When_Empty bool

	ProductID_Explicit_Null_When_Empty bool

	Slot_Explicit_Null_When_Empty bool

	SriovState_Explicit_Null_When_Empty bool

	SubsystemProductID_Explicit_Null_When_Empty bool

	SubsystemVendorID_Explicit_Null_When_Empty bool

	TotalMdevNum_Explicit_Null_When_Empty bool

	TotalPartitioningNum_Explicit_Null_When_Empty bool

	TotalVfNum_Explicit_Null_When_Empty bool

	UsageType_Explicit_Null_When_Empty bool

	UsedMdevNum_Explicit_Null_When_Empty bool

	UsedPartitioningNum_Explicit_Null_When_Empty bool

	UsedVfNum_Explicit_Null_When_Empty bool

	UserUsage_Explicit_Null_When_Empty bool

	VendorID_Explicit_Null_When_Empty bool

	VerdorName_Explicit_Null_When_Empty bool
}

func (m PciDevice) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle nullable field bus
	if m.Bus != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bus\":")
		bytes, err := swag.WriteJSON(m.Bus)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Bus_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bus\":null")
		first = false
	}

	// handle nullable field bus_location
	if m.BusLocation != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bus_location\":")
		bytes, err := swag.WriteJSON(m.BusLocation)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.BusLocation_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"bus_location\":null")
		first = false
	}

	// handle nullable field class_code
	if m.ClassCode != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"class_code\":")
		bytes, err := swag.WriteJSON(m.ClassCode)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ClassCode_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"class_code\":null")
		first = false
	}

	// handle nullable field device_type
	if m.DeviceType != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"device_type\":")
		bytes, err := swag.WriteJSON(m.DeviceType)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.DeviceType_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"device_type\":null")
		first = false
	}

	// handle nullable field entityAsyncStatus
	if m.EntityAsyncStatus != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"entityAsyncStatus\":")
		bytes, err := swag.WriteJSON(m.EntityAsyncStatus)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.EntityAsyncStatus_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"entityAsyncStatus\":null")
		first = false
	}

	// handle nullable field function_num
	if m.FunctionNum != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"function_num\":")
		bytes, err := swag.WriteJSON(m.FunctionNum)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.FunctionNum_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"function_num\":null")
		first = false
	}

	// handle nullable field host
	if m.Host != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"host\":")
		bytes, err := swag.WriteJSON(m.Host)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Host_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"host\":null")
		first = false
	}

	// handle nullable field id
	if m.ID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id\":")
		bytes, err := swag.WriteJSON(m.ID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id\":null")
		first = false
	}

	// handle nullable field iommu_status
	if m.IommuStatus != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"iommu_status\":")
		bytes, err := swag.WriteJSON(m.IommuStatus)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IommuStatus_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"iommu_status\":null")
		first = false
	}

	// handle nullable field local_id
	if m.LocalID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id\":")
		bytes, err := swag.WriteJSON(m.LocalID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.LocalID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id\":null")
		first = false
	}

	// handle nullable field mdev_state
	if m.MdevState != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"mdev_state\":")
		bytes, err := swag.WriteJSON(m.MdevState)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.MdevState_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"mdev_state\":null")
		first = false
	}

	// handle nullable field mdev_type_id
	if m.MdevTypeID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"mdev_type_id\":")
		bytes, err := swag.WriteJSON(m.MdevTypeID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.MdevTypeID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"mdev_type_id\":null")
		first = false
	}

	// handle nullable field name
	if m.Name != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name\":")
		bytes, err := swag.WriteJSON(m.Name)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Name_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name\":null")
		first = false
	}

	// handle nullable field partitioning_state
	if m.PartitioningState != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"partitioning_state\":")
		bytes, err := swag.WriteJSON(m.PartitioningState)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.PartitioningState_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"partitioning_state\":null")
		first = false
	}

	// handle nullable field product_id
	if m.ProductID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"product_id\":")
		bytes, err := swag.WriteJSON(m.ProductID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.ProductID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"product_id\":null")
		first = false
	}

	// handle nullable field slot
	if m.Slot != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"slot\":")
		bytes, err := swag.WriteJSON(m.Slot)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.Slot_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"slot\":null")
		first = false
	}

	// handle nullable field sriov_state
	if m.SriovState != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"sriov_state\":")
		bytes, err := swag.WriteJSON(m.SriovState)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.SriovState_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"sriov_state\":null")
		first = false
	}

	// handle nullable field subsystem_product_id
	if m.SubsystemProductID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"subsystem_product_id\":")
		bytes, err := swag.WriteJSON(m.SubsystemProductID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.SubsystemProductID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"subsystem_product_id\":null")
		first = false
	}

	// handle nullable field subsystem_vendor_id
	if m.SubsystemVendorID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"subsystem_vendor_id\":")
		bytes, err := swag.WriteJSON(m.SubsystemVendorID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.SubsystemVendorID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"subsystem_vendor_id\":null")
		first = false
	}

	// handle nullable field total_mdev_num
	if m.TotalMdevNum != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"total_mdev_num\":")
		bytes, err := swag.WriteJSON(m.TotalMdevNum)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.TotalMdevNum_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"total_mdev_num\":null")
		first = false
	}

	// handle nullable field total_partitioning_num
	if m.TotalPartitioningNum != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"total_partitioning_num\":")
		bytes, err := swag.WriteJSON(m.TotalPartitioningNum)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.TotalPartitioningNum_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"total_partitioning_num\":null")
		first = false
	}

	// handle nullable field total_vf_num
	if m.TotalVfNum != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"total_vf_num\":")
		bytes, err := swag.WriteJSON(m.TotalVfNum)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.TotalVfNum_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"total_vf_num\":null")
		first = false
	}

	// handle nullable field usage_type
	if m.UsageType != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"usage_type\":")
		bytes, err := swag.WriteJSON(m.UsageType)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.UsageType_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"usage_type\":null")
		first = false
	}

	// handle nullable field used_mdev_num
	if m.UsedMdevNum != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"used_mdev_num\":")
		bytes, err := swag.WriteJSON(m.UsedMdevNum)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.UsedMdevNum_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"used_mdev_num\":null")
		first = false
	}

	// handle nullable field used_partitioning_num
	if m.UsedPartitioningNum != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"used_partitioning_num\":")
		bytes, err := swag.WriteJSON(m.UsedPartitioningNum)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.UsedPartitioningNum_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"used_partitioning_num\":null")
		first = false
	}

	// handle nullable field used_vf_num
	if m.UsedVfNum != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"used_vf_num\":")
		bytes, err := swag.WriteJSON(m.UsedVfNum)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.UsedVfNum_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"used_vf_num\":null")
		first = false
	}

	// handle nullable field user_usage
	if m.UserUsage != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"user_usage\":")
		bytes, err := swag.WriteJSON(m.UserUsage)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.UserUsage_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"user_usage\":null")
		first = false
	}

	// handle nullable field vendor_id
	if m.VendorID != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vendor_id\":")
		bytes, err := swag.WriteJSON(m.VendorID)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VendorID_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"vendor_id\":null")
		first = false
	}

	// handle nullable field verdor_name
	if m.VerdorName != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"verdor_name\":")
		bytes, err := swag.WriteJSON(m.VerdorName)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.VerdorName_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"verdor_name\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this pci device
func (m *PciDevice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBusLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClassCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFunctionNum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIommuStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMdevState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartitioningState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSriovState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubsystemProductID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubsystemVendorID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsageType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserUsage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVendorID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVerdorName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PciDevice) validateBus(formats strfmt.Registry) error {

	if err := validate.Required("bus", "body", m.Bus); err != nil {
		return err
	}

	return nil
}

func (m *PciDevice) validateBusLocation(formats strfmt.Registry) error {

	if err := validate.Required("bus_location", "body", m.BusLocation); err != nil {
		return err
	}

	return nil
}

func (m *PciDevice) validateClassCode(formats strfmt.Registry) error {

	if err := validate.Required("class_code", "body", m.ClassCode); err != nil {
		return err
	}

	return nil
}

func (m *PciDevice) validateEntityAsyncStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatus) { // not required
		return nil
	}

	if m.EntityAsyncStatus != nil {
		if err := m.EntityAsyncStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus")
			}
			return err
		}
	}

	return nil
}

func (m *PciDevice) validateFunctionNum(formats strfmt.Registry) error {

	if err := validate.Required("function_num", "body", m.FunctionNum); err != nil {
		return err
	}

	return nil
}

func (m *PciDevice) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", m.Host); err != nil {
		return err
	}

	if m.Host != nil {
		if err := m.Host.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("host")
			}
			return err
		}
	}

	return nil
}

func (m *PciDevice) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *PciDevice) validateIommuStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.IommuStatus) { // not required
		return nil
	}

	if m.IommuStatus != nil {
		if err := m.IommuStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iommu_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iommu_status")
			}
			return err
		}
	}

	return nil
}

func (m *PciDevice) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *PciDevice) validateMdevState(formats strfmt.Registry) error {
	if swag.IsZero(m.MdevState) { // not required
		return nil
	}

	if m.MdevState != nil {
		if err := m.MdevState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mdev_state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mdev_state")
			}
			return err
		}
	}

	return nil
}

func (m *PciDevice) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PciDevice) validatePartitioningState(formats strfmt.Registry) error {
	if swag.IsZero(m.PartitioningState) { // not required
		return nil
	}

	if m.PartitioningState != nil {
		if err := m.PartitioningState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("partitioning_state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("partitioning_state")
			}
			return err
		}
	}

	return nil
}

func (m *PciDevice) validateProductID(formats strfmt.Registry) error {

	if err := validate.Required("product_id", "body", m.ProductID); err != nil {
		return err
	}

	return nil
}

func (m *PciDevice) validateSlot(formats strfmt.Registry) error {

	if err := validate.Required("slot", "body", m.Slot); err != nil {
		return err
	}

	return nil
}

func (m *PciDevice) validateSriovState(formats strfmt.Registry) error {
	if swag.IsZero(m.SriovState) { // not required
		return nil
	}

	if m.SriovState != nil {
		if err := m.SriovState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sriov_state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sriov_state")
			}
			return err
		}
	}

	return nil
}

func (m *PciDevice) validateSubsystemProductID(formats strfmt.Registry) error {

	if err := validate.Required("subsystem_product_id", "body", m.SubsystemProductID); err != nil {
		return err
	}

	return nil
}

func (m *PciDevice) validateSubsystemVendorID(formats strfmt.Registry) error {

	if err := validate.Required("subsystem_vendor_id", "body", m.SubsystemVendorID); err != nil {
		return err
	}

	return nil
}

func (m *PciDevice) validateUsageType(formats strfmt.Registry) error {
	if swag.IsZero(m.UsageType) { // not required
		return nil
	}

	if m.UsageType != nil {
		if err := m.UsageType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usage_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usage_type")
			}
			return err
		}
	}

	return nil
}

func (m *PciDevice) validateUserUsage(formats strfmt.Registry) error {
	if swag.IsZero(m.UserUsage) { // not required
		return nil
	}

	if m.UserUsage != nil {
		if err := m.UserUsage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user_usage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user_usage")
			}
			return err
		}
	}

	return nil
}

func (m *PciDevice) validateVendorID(formats strfmt.Registry) error {

	if err := validate.Required("vendor_id", "body", m.VendorID); err != nil {
		return err
	}

	return nil
}

func (m *PciDevice) validateVerdorName(formats strfmt.Registry) error {

	if err := validate.Required("verdor_name", "body", m.VerdorName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this pci device based on the context it is used
func (m *PciDevice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHost(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIommuStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMdevState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePartitioningState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSriovState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsageType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserUsage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PciDevice) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.EntityAsyncStatus != nil {
		if err := m.EntityAsyncStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus")
			}
			return err
		}
	}

	return nil
}

func (m *PciDevice) contextValidateHost(ctx context.Context, formats strfmt.Registry) error {

	if m.Host != nil {
		if err := m.Host.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("host")
			}
			return err
		}
	}

	return nil
}

func (m *PciDevice) contextValidateIommuStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.IommuStatus != nil {
		if err := m.IommuStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iommu_status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("iommu_status")
			}
			return err
		}
	}

	return nil
}

func (m *PciDevice) contextValidateMdevState(ctx context.Context, formats strfmt.Registry) error {

	if m.MdevState != nil {
		if err := m.MdevState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mdev_state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mdev_state")
			}
			return err
		}
	}

	return nil
}

func (m *PciDevice) contextValidatePartitioningState(ctx context.Context, formats strfmt.Registry) error {

	if m.PartitioningState != nil {
		if err := m.PartitioningState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("partitioning_state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("partitioning_state")
			}
			return err
		}
	}

	return nil
}

func (m *PciDevice) contextValidateSriovState(ctx context.Context, formats strfmt.Registry) error {

	if m.SriovState != nil {
		if err := m.SriovState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sriov_state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sriov_state")
			}
			return err
		}
	}

	return nil
}

func (m *PciDevice) contextValidateUsageType(ctx context.Context, formats strfmt.Registry) error {

	if m.UsageType != nil {
		if err := m.UsageType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usage_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usage_type")
			}
			return err
		}
	}

	return nil
}

func (m *PciDevice) contextValidateUserUsage(ctx context.Context, formats strfmt.Registry) error {

	if m.UserUsage != nil {
		if err := m.UserUsage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user_usage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user_usage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PciDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PciDevice) UnmarshalBinary(b []byte) error {
	var res PciDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
