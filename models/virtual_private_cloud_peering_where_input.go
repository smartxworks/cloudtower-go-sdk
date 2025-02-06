// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VirtualPrivateCloudPeeringWhereInput virtual private cloud peering where input
//
// swagger:model VirtualPrivateCloudPeeringWhereInput
type VirtualPrivateCloudPeeringWhereInput struct {

	// a n d
	AND []*VirtualPrivateCloudPeeringWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*VirtualPrivateCloudPeeringWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*VirtualPrivateCloudPeeringWhereInput `json:"OR,omitempty"`

	// dst vpc
	DstVpc *VirtualPrivateCloudWhereInput `json:"dst_vpc,omitempty"`

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// entity async status in
	EntityAsyncStatusIn []EntityAsyncStatus `json:"entityAsyncStatus_in,omitempty"`

	// entity async status not
	EntityAsyncStatusNot *EntityAsyncStatus `json:"entityAsyncStatus_not,omitempty"`

	// entity async status not in
	EntityAsyncStatusNotIn []EntityAsyncStatus `json:"entityAsyncStatus_not_in,omitempty"`

	// id
	ID *string `json:"id,omitempty"`

	// id contains
	IDContains *string `json:"id_contains,omitempty"`

	// id ends with
	IDEndsWith *string `json:"id_ends_with,omitempty"`

	// id gt
	IDGt *string `json:"id_gt,omitempty"`

	// id gte
	IDGte *string `json:"id_gte,omitempty"`

	// id in
	IDIn []string `json:"id_in,omitempty"`

	// id lt
	IDLt *string `json:"id_lt,omitempty"`

	// id lte
	IDLte *string `json:"id_lte,omitempty"`

	// id not
	IDNot *string `json:"id_not,omitempty"`

	// id not contains
	IDNotContains *string `json:"id_not_contains,omitempty"`

	// id not ends with
	IDNotEndsWith *string `json:"id_not_ends_with,omitempty"`

	// id not in
	IDNotIn []string `json:"id_not_in,omitempty"`

	// id not starts with
	IDNotStartsWith *string `json:"id_not_starts_with,omitempty"`

	// id starts with
	IDStartsWith *string `json:"id_starts_with,omitempty"`

	// local id
	LocalID *string `json:"local_id,omitempty"`

	// local id contains
	LocalIDContains *string `json:"local_id_contains,omitempty"`

	// local id ends with
	LocalIDEndsWith *string `json:"local_id_ends_with,omitempty"`

	// local id gt
	LocalIDGt *string `json:"local_id_gt,omitempty"`

	// local id gte
	LocalIDGte *string `json:"local_id_gte,omitempty"`

	// local id in
	LocalIDIn []string `json:"local_id_in,omitempty"`

	// local id lt
	LocalIDLt *string `json:"local_id_lt,omitempty"`

	// local id lte
	LocalIDLte *string `json:"local_id_lte,omitempty"`

	// local id not
	LocalIDNot *string `json:"local_id_not,omitempty"`

	// local id not contains
	LocalIDNotContains *string `json:"local_id_not_contains,omitempty"`

	// local id not ends with
	LocalIDNotEndsWith *string `json:"local_id_not_ends_with,omitempty"`

	// local id not in
	LocalIDNotIn []string `json:"local_id_not_in,omitempty"`

	// local id not starts with
	LocalIDNotStartsWith *string `json:"local_id_not_starts_with,omitempty"`

	// local id starts with
	LocalIDStartsWith *string `json:"local_id_starts_with,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// name contains
	NameContains *string `json:"name_contains,omitempty"`

	// name ends with
	NameEndsWith *string `json:"name_ends_with,omitempty"`

	// name gt
	NameGt *string `json:"name_gt,omitempty"`

	// name gte
	NameGte *string `json:"name_gte,omitempty"`

	// name in
	NameIn []string `json:"name_in,omitempty"`

	// name lt
	NameLt *string `json:"name_lt,omitempty"`

	// name lte
	NameLte *string `json:"name_lte,omitempty"`

	// name not
	NameNot *string `json:"name_not,omitempty"`

	// name not contains
	NameNotContains *string `json:"name_not_contains,omitempty"`

	// name not ends with
	NameNotEndsWith *string `json:"name_not_ends_with,omitempty"`

	// name not in
	NameNotIn []string `json:"name_not_in,omitempty"`

	// name not starts with
	NameNotStartsWith *string `json:"name_not_starts_with,omitempty"`

	// name starts with
	NameStartsWith *string `json:"name_starts_with,omitempty"`

	// src vpc
	SrcVpc *VirtualPrivateCloudWhereInput `json:"src_vpc,omitempty"`

	MarshalOpts *VirtualPrivateCloudPeeringWhereInputMarshalOpts `json:"-"`
}

type VirtualPrivateCloudPeeringWhereInputMarshalOpts struct {
	DstVpc_Explicit_Null_When_Empty bool

	EntityAsyncStatus_Explicit_Null_When_Empty bool

	EntityAsyncStatusNot_Explicit_Null_When_Empty bool

	ID_Explicit_Null_When_Empty bool

	IDContains_Explicit_Null_When_Empty bool

	IDEndsWith_Explicit_Null_When_Empty bool

	IDGt_Explicit_Null_When_Empty bool

	IDGte_Explicit_Null_When_Empty bool

	IDLt_Explicit_Null_When_Empty bool

	IDLte_Explicit_Null_When_Empty bool

	IDNot_Explicit_Null_When_Empty bool

	IDNotContains_Explicit_Null_When_Empty bool

	IDNotEndsWith_Explicit_Null_When_Empty bool

	IDNotStartsWith_Explicit_Null_When_Empty bool

	IDStartsWith_Explicit_Null_When_Empty bool

	LocalID_Explicit_Null_When_Empty bool

	LocalIDContains_Explicit_Null_When_Empty bool

	LocalIDEndsWith_Explicit_Null_When_Empty bool

	LocalIDGt_Explicit_Null_When_Empty bool

	LocalIDGte_Explicit_Null_When_Empty bool

	LocalIDLt_Explicit_Null_When_Empty bool

	LocalIDLte_Explicit_Null_When_Empty bool

	LocalIDNot_Explicit_Null_When_Empty bool

	LocalIDNotContains_Explicit_Null_When_Empty bool

	LocalIDNotEndsWith_Explicit_Null_When_Empty bool

	LocalIDNotStartsWith_Explicit_Null_When_Empty bool

	LocalIDStartsWith_Explicit_Null_When_Empty bool

	Name_Explicit_Null_When_Empty bool

	NameContains_Explicit_Null_When_Empty bool

	NameEndsWith_Explicit_Null_When_Empty bool

	NameGt_Explicit_Null_When_Empty bool

	NameGte_Explicit_Null_When_Empty bool

	NameLt_Explicit_Null_When_Empty bool

	NameLte_Explicit_Null_When_Empty bool

	NameNot_Explicit_Null_When_Empty bool

	NameNotContains_Explicit_Null_When_Empty bool

	NameNotEndsWith_Explicit_Null_When_Empty bool

	NameNotStartsWith_Explicit_Null_When_Empty bool

	NameStartsWith_Explicit_Null_When_Empty bool

	SrcVpc_Explicit_Null_When_Empty bool
}

func (m VirtualPrivateCloudPeeringWhereInput) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("{")

	first := true

	// handle non nullable field AND with omitempty
	if swag.IsZero(m.AND) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"AND\":")
		bytes, err := swag.WriteJSON(m.AND)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle non nullable field NOT with omitempty
	if swag.IsZero(m.NOT) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"NOT\":")
		bytes, err := swag.WriteJSON(m.NOT)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle non nullable field OR with omitempty
	if swag.IsZero(m.OR) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"OR\":")
		bytes, err := swag.WriteJSON(m.OR)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field dst_vpc
	if m.DstVpc != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"dst_vpc\":")
		bytes, err := swag.WriteJSON(m.DstVpc)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.DstVpc_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"dst_vpc\":null")
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

	// handle non nullable field entityAsyncStatus_in with omitempty
	if swag.IsZero(m.EntityAsyncStatusIn) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"entityAsyncStatus_in\":")
		bytes, err := swag.WriteJSON(m.EntityAsyncStatusIn)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field entityAsyncStatus_not
	if m.EntityAsyncStatusNot != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"entityAsyncStatus_not\":")
		bytes, err := swag.WriteJSON(m.EntityAsyncStatusNot)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.EntityAsyncStatusNot_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"entityAsyncStatus_not\":null")
		first = false
	}

	// handle non nullable field entityAsyncStatus_not_in with omitempty
	if swag.IsZero(m.EntityAsyncStatusNotIn) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"entityAsyncStatus_not_in\":")
		bytes, err := swag.WriteJSON(m.EntityAsyncStatusNotIn)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
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

	// handle nullable field id_contains
	if m.IDContains != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_contains\":")
		bytes, err := swag.WriteJSON(m.IDContains)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDContains_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_contains\":null")
		first = false
	}

	// handle nullable field id_ends_with
	if m.IDEndsWith != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_ends_with\":")
		bytes, err := swag.WriteJSON(m.IDEndsWith)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDEndsWith_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_ends_with\":null")
		first = false
	}

	// handle nullable field id_gt
	if m.IDGt != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_gt\":")
		bytes, err := swag.WriteJSON(m.IDGt)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDGt_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_gt\":null")
		first = false
	}

	// handle nullable field id_gte
	if m.IDGte != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_gte\":")
		bytes, err := swag.WriteJSON(m.IDGte)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDGte_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_gte\":null")
		first = false
	}

	// handle non nullable field id_in with omitempty
	if swag.IsZero(m.IDIn) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_in\":")
		bytes, err := swag.WriteJSON(m.IDIn)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field id_lt
	if m.IDLt != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_lt\":")
		bytes, err := swag.WriteJSON(m.IDLt)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDLt_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_lt\":null")
		first = false
	}

	// handle nullable field id_lte
	if m.IDLte != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_lte\":")
		bytes, err := swag.WriteJSON(m.IDLte)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDLte_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_lte\":null")
		first = false
	}

	// handle nullable field id_not
	if m.IDNot != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_not\":")
		bytes, err := swag.WriteJSON(m.IDNot)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDNot_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_not\":null")
		first = false
	}

	// handle nullable field id_not_contains
	if m.IDNotContains != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_not_contains\":")
		bytes, err := swag.WriteJSON(m.IDNotContains)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDNotContains_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_not_contains\":null")
		first = false
	}

	// handle nullable field id_not_ends_with
	if m.IDNotEndsWith != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_not_ends_with\":")
		bytes, err := swag.WriteJSON(m.IDNotEndsWith)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDNotEndsWith_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_not_ends_with\":null")
		first = false
	}

	// handle non nullable field id_not_in with omitempty
	if swag.IsZero(m.IDNotIn) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_not_in\":")
		bytes, err := swag.WriteJSON(m.IDNotIn)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field id_not_starts_with
	if m.IDNotStartsWith != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_not_starts_with\":")
		bytes, err := swag.WriteJSON(m.IDNotStartsWith)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDNotStartsWith_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_not_starts_with\":null")
		first = false
	}

	// handle nullable field id_starts_with
	if m.IDStartsWith != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_starts_with\":")
		bytes, err := swag.WriteJSON(m.IDStartsWith)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.IDStartsWith_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"id_starts_with\":null")
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

	// handle nullable field local_id_contains
	if m.LocalIDContains != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id_contains\":")
		bytes, err := swag.WriteJSON(m.LocalIDContains)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.LocalIDContains_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id_contains\":null")
		first = false
	}

	// handle nullable field local_id_ends_with
	if m.LocalIDEndsWith != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id_ends_with\":")
		bytes, err := swag.WriteJSON(m.LocalIDEndsWith)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.LocalIDEndsWith_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id_ends_with\":null")
		first = false
	}

	// handle nullable field local_id_gt
	if m.LocalIDGt != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id_gt\":")
		bytes, err := swag.WriteJSON(m.LocalIDGt)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.LocalIDGt_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id_gt\":null")
		first = false
	}

	// handle nullable field local_id_gte
	if m.LocalIDGte != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id_gte\":")
		bytes, err := swag.WriteJSON(m.LocalIDGte)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.LocalIDGte_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id_gte\":null")
		first = false
	}

	// handle non nullable field local_id_in with omitempty
	if swag.IsZero(m.LocalIDIn) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id_in\":")
		bytes, err := swag.WriteJSON(m.LocalIDIn)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field local_id_lt
	if m.LocalIDLt != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id_lt\":")
		bytes, err := swag.WriteJSON(m.LocalIDLt)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.LocalIDLt_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id_lt\":null")
		first = false
	}

	// handle nullable field local_id_lte
	if m.LocalIDLte != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id_lte\":")
		bytes, err := swag.WriteJSON(m.LocalIDLte)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.LocalIDLte_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id_lte\":null")
		first = false
	}

	// handle nullable field local_id_not
	if m.LocalIDNot != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id_not\":")
		bytes, err := swag.WriteJSON(m.LocalIDNot)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.LocalIDNot_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id_not\":null")
		first = false
	}

	// handle nullable field local_id_not_contains
	if m.LocalIDNotContains != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id_not_contains\":")
		bytes, err := swag.WriteJSON(m.LocalIDNotContains)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.LocalIDNotContains_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id_not_contains\":null")
		first = false
	}

	// handle nullable field local_id_not_ends_with
	if m.LocalIDNotEndsWith != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id_not_ends_with\":")
		bytes, err := swag.WriteJSON(m.LocalIDNotEndsWith)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.LocalIDNotEndsWith_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id_not_ends_with\":null")
		first = false
	}

	// handle non nullable field local_id_not_in with omitempty
	if swag.IsZero(m.LocalIDNotIn) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id_not_in\":")
		bytes, err := swag.WriteJSON(m.LocalIDNotIn)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field local_id_not_starts_with
	if m.LocalIDNotStartsWith != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id_not_starts_with\":")
		bytes, err := swag.WriteJSON(m.LocalIDNotStartsWith)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.LocalIDNotStartsWith_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id_not_starts_with\":null")
		first = false
	}

	// handle nullable field local_id_starts_with
	if m.LocalIDStartsWith != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id_starts_with\":")
		bytes, err := swag.WriteJSON(m.LocalIDStartsWith)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.LocalIDStartsWith_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"local_id_starts_with\":null")
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

	// handle nullable field name_contains
	if m.NameContains != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_contains\":")
		bytes, err := swag.WriteJSON(m.NameContains)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NameContains_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_contains\":null")
		first = false
	}

	// handle nullable field name_ends_with
	if m.NameEndsWith != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_ends_with\":")
		bytes, err := swag.WriteJSON(m.NameEndsWith)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NameEndsWith_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_ends_with\":null")
		first = false
	}

	// handle nullable field name_gt
	if m.NameGt != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_gt\":")
		bytes, err := swag.WriteJSON(m.NameGt)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NameGt_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_gt\":null")
		first = false
	}

	// handle nullable field name_gte
	if m.NameGte != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_gte\":")
		bytes, err := swag.WriteJSON(m.NameGte)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NameGte_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_gte\":null")
		first = false
	}

	// handle non nullable field name_in with omitempty
	if swag.IsZero(m.NameIn) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_in\":")
		bytes, err := swag.WriteJSON(m.NameIn)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field name_lt
	if m.NameLt != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_lt\":")
		bytes, err := swag.WriteJSON(m.NameLt)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NameLt_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_lt\":null")
		first = false
	}

	// handle nullable field name_lte
	if m.NameLte != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_lte\":")
		bytes, err := swag.WriteJSON(m.NameLte)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NameLte_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_lte\":null")
		first = false
	}

	// handle nullable field name_not
	if m.NameNot != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_not\":")
		bytes, err := swag.WriteJSON(m.NameNot)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NameNot_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_not\":null")
		first = false
	}

	// handle nullable field name_not_contains
	if m.NameNotContains != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_not_contains\":")
		bytes, err := swag.WriteJSON(m.NameNotContains)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NameNotContains_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_not_contains\":null")
		first = false
	}

	// handle nullable field name_not_ends_with
	if m.NameNotEndsWith != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_not_ends_with\":")
		bytes, err := swag.WriteJSON(m.NameNotEndsWith)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NameNotEndsWith_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_not_ends_with\":null")
		first = false
	}

	// handle non nullable field name_not_in with omitempty
	if swag.IsZero(m.NameNotIn) {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_not_in\":")
		bytes, err := swag.WriteJSON(m.NameNotIn)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	}

	// handle nullable field name_not_starts_with
	if m.NameNotStartsWith != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_not_starts_with\":")
		bytes, err := swag.WriteJSON(m.NameNotStartsWith)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NameNotStartsWith_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_not_starts_with\":null")
		first = false
	}

	// handle nullable field name_starts_with
	if m.NameStartsWith != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_starts_with\":")
		bytes, err := swag.WriteJSON(m.NameStartsWith)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.NameStartsWith_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"name_starts_with\":null")
		first = false
	}

	// handle nullable field src_vpc
	if m.SrcVpc != nil {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"src_vpc\":")
		bytes, err := swag.WriteJSON(m.SrcVpc)
		if err != nil {
			return nil, err
		}
		b.Write(bytes)
		first = false
	} else if m.MarshalOpts != nil && m.MarshalOpts.SrcVpc_Explicit_Null_When_Empty {
		if !first {
			b.WriteString(",")
		}
		b.WriteString("\"src_vpc\":null")
		first = false
	}

	b.WriteString("}")
	return b.Bytes(), nil
}

// Validate validates this virtual private cloud peering where input
func (m *VirtualPrivateCloudPeeringWhereInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAND(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNOT(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOR(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDstVpc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcVpc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualPrivateCloudPeeringWhereInput) validateAND(formats strfmt.Registry) error {
	if swag.IsZero(m.AND) { // not required
		return nil
	}

	for i := 0; i < len(m.AND); i++ {
		if swag.IsZero(m.AND[i]) { // not required
			continue
		}

		if m.AND[i] != nil {
			if err := m.AND[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudPeeringWhereInput) validateNOT(formats strfmt.Registry) error {
	if swag.IsZero(m.NOT) { // not required
		return nil
	}

	for i := 0; i < len(m.NOT); i++ {
		if swag.IsZero(m.NOT[i]) { // not required
			continue
		}

		if m.NOT[i] != nil {
			if err := m.NOT[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudPeeringWhereInput) validateOR(formats strfmt.Registry) error {
	if swag.IsZero(m.OR) { // not required
		return nil
	}

	for i := 0; i < len(m.OR); i++ {
		if swag.IsZero(m.OR[i]) { // not required
			continue
		}

		if m.OR[i] != nil {
			if err := m.OR[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudPeeringWhereInput) validateDstVpc(formats strfmt.Registry) error {
	if swag.IsZero(m.DstVpc) { // not required
		return nil
	}

	if m.DstVpc != nil {
		if err := m.DstVpc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dst_vpc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dst_vpc")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudPeeringWhereInput) validateEntityAsyncStatus(formats strfmt.Registry) error {
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

func (m *VirtualPrivateCloudPeeringWhereInput) validateEntityAsyncStatusIn(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusIn) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityAsyncStatusIn); i++ {

		if err := m.EntityAsyncStatusIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VirtualPrivateCloudPeeringWhereInput) validateEntityAsyncStatusNot(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusNot) { // not required
		return nil
	}

	if m.EntityAsyncStatusNot != nil {
		if err := m.EntityAsyncStatusNot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudPeeringWhereInput) validateEntityAsyncStatusNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityAsyncStatusNotIn); i++ {

		if err := m.EntityAsyncStatusNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VirtualPrivateCloudPeeringWhereInput) validateSrcVpc(formats strfmt.Registry) error {
	if swag.IsZero(m.SrcVpc) { // not required
		return nil
	}

	if m.SrcVpc != nil {
		if err := m.SrcVpc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("src_vpc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("src_vpc")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this virtual private cloud peering where input based on the context it is used
func (m *VirtualPrivateCloudPeeringWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAND(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNOT(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOR(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDstVpc(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSrcVpc(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualPrivateCloudPeeringWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AND); i++ {

		if m.AND[i] != nil {
			if err := m.AND[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudPeeringWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NOT); i++ {

		if m.NOT[i] != nil {
			if err := m.NOT[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudPeeringWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OR); i++ {

		if m.OR[i] != nil {
			if err := m.OR[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualPrivateCloudPeeringWhereInput) contextValidateDstVpc(ctx context.Context, formats strfmt.Registry) error {

	if m.DstVpc != nil {
		if err := m.DstVpc.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dst_vpc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dst_vpc")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudPeeringWhereInput) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VirtualPrivateCloudPeeringWhereInput) contextValidateEntityAsyncStatusIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntityAsyncStatusIn); i++ {

		if err := m.EntityAsyncStatusIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VirtualPrivateCloudPeeringWhereInput) contextValidateEntityAsyncStatusNot(ctx context.Context, formats strfmt.Registry) error {

	if m.EntityAsyncStatusNot != nil {
		if err := m.EntityAsyncStatusNot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualPrivateCloudPeeringWhereInput) contextValidateEntityAsyncStatusNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntityAsyncStatusNotIn); i++ {

		if err := m.EntityAsyncStatusNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VirtualPrivateCloudPeeringWhereInput) contextValidateSrcVpc(ctx context.Context, formats strfmt.Registry) error {

	if m.SrcVpc != nil {
		if err := m.SrcVpc.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("src_vpc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("src_vpc")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VirtualPrivateCloudPeeringWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualPrivateCloudPeeringWhereInput) UnmarshalBinary(b []byte) error {
	var res VirtualPrivateCloudPeeringWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
