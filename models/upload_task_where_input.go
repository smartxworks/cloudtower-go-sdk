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
)

// UploadTaskWhereInput upload task where input
// Example: {"AND":"UploadTaskWhereInput[]","NOT":"UploadTaskWhereInput[]","OR":"UploadTaskWhereInput[]","chunk_size":0,"chunk_size_gt":0,"chunk_size_gte":0,"chunk_size_in":[0],"chunk_size_lt":0,"chunk_size_lte":0,"chunk_size_not":0,"chunk_size_not_in":[0],"current_chunk":0,"current_chunk_gt":0,"current_chunk_gte":0,"current_chunk_in":[0],"current_chunk_lt":0,"current_chunk_lte":0,"current_chunk_not":0,"current_chunk_not_in":[0],"finished_at":"string","finished_at_gt":"string","finished_at_gte":"string","finished_at_in":["string"],"finished_at_lt":"string","finished_at_lte":"string","finished_at_not":"string","finished_at_not_in":["string"],"id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","resource_type":"CLUSTER_IMAGE","resource_type_in":["CLUSTER_IMAGE"],"resource_type_not":"CLUSTER_IMAGE","resource_type_not_in":["CLUSTER_IMAGE"],"size":0,"size_gt":0,"size_gte":0,"size_in":[0],"size_lt":0,"size_lte":0,"size_not":0,"size_not_in":[0],"started_at":"string","started_at_gt":"string","started_at_gte":"string","started_at_in":["string"],"started_at_lt":"string","started_at_lte":"string","started_at_not":"string","started_at_not_in":["string"],"status":"FAILED","status_in":["FAILED"],"status_not":"FAILED","status_not_in":["FAILED"],"updatedAt":"string","updatedAt_gt":"string","updatedAt_gte":"string","updatedAt_in":["string"],"updatedAt_lt":"string","updatedAt_lte":"string","updatedAt_not":"string","updatedAt_not_in":["string"]}
//
// swagger:model UploadTaskWhereInput
type UploadTaskWhereInput struct {

	// a n d
	AND []*UploadTaskWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*UploadTaskWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*UploadTaskWhereInput `json:"OR,omitempty"`

	// chunk size
	ChunkSize *int64 `json:"chunk_size,omitempty"`

	// chunk size gt
	ChunkSizeGt *int64 `json:"chunk_size_gt,omitempty"`

	// chunk size gte
	ChunkSizeGte *int64 `json:"chunk_size_gte,omitempty"`

	// chunk size in
	ChunkSizeIn []int64 `json:"chunk_size_in,omitempty"`

	// chunk size lt
	ChunkSizeLt *int64 `json:"chunk_size_lt,omitempty"`

	// chunk size lte
	ChunkSizeLte *int64 `json:"chunk_size_lte,omitempty"`

	// chunk size not
	ChunkSizeNot *int64 `json:"chunk_size_not,omitempty"`

	// chunk size not in
	ChunkSizeNotIn []int64 `json:"chunk_size_not_in,omitempty"`

	// current chunk
	CurrentChunk *int32 `json:"current_chunk,omitempty"`

	// current chunk gt
	CurrentChunkGt *int32 `json:"current_chunk_gt,omitempty"`

	// current chunk gte
	CurrentChunkGte *int32 `json:"current_chunk_gte,omitempty"`

	// current chunk in
	CurrentChunkIn []int32 `json:"current_chunk_in,omitempty"`

	// current chunk lt
	CurrentChunkLt *int32 `json:"current_chunk_lt,omitempty"`

	// current chunk lte
	CurrentChunkLte *int32 `json:"current_chunk_lte,omitempty"`

	// current chunk not
	CurrentChunkNot *int32 `json:"current_chunk_not,omitempty"`

	// current chunk not in
	CurrentChunkNotIn []int32 `json:"current_chunk_not_in,omitempty"`

	// finished at
	FinishedAt *string `json:"finished_at,omitempty"`

	// finished at gt
	FinishedAtGt *string `json:"finished_at_gt,omitempty"`

	// finished at gte
	FinishedAtGte *string `json:"finished_at_gte,omitempty"`

	// finished at in
	FinishedAtIn []string `json:"finished_at_in,omitempty"`

	// finished at lt
	FinishedAtLt *string `json:"finished_at_lt,omitempty"`

	// finished at lte
	FinishedAtLte *string `json:"finished_at_lte,omitempty"`

	// finished at not
	FinishedAtNot *string `json:"finished_at_not,omitempty"`

	// finished at not in
	FinishedAtNotIn []string `json:"finished_at_not_in,omitempty"`

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

	// resource type
	ResourceType *UploadResourceType `json:"resource_type,omitempty"`

	// resource type in
	ResourceTypeIn []UploadResourceType `json:"resource_type_in,omitempty"`

	// resource type not
	ResourceTypeNot *UploadResourceType `json:"resource_type_not,omitempty"`

	// resource type not in
	ResourceTypeNotIn []UploadResourceType `json:"resource_type_not_in,omitempty"`

	// size
	Size *int64 `json:"size,omitempty"`

	// size gt
	SizeGt *int64 `json:"size_gt,omitempty"`

	// size gte
	SizeGte *int64 `json:"size_gte,omitempty"`

	// size in
	SizeIn []int64 `json:"size_in,omitempty"`

	// size lt
	SizeLt *int64 `json:"size_lt,omitempty"`

	// size lte
	SizeLte *int64 `json:"size_lte,omitempty"`

	// size not
	SizeNot *int64 `json:"size_not,omitempty"`

	// size not in
	SizeNotIn []int64 `json:"size_not_in,omitempty"`

	// started at
	StartedAt *string `json:"started_at,omitempty"`

	// started at gt
	StartedAtGt *string `json:"started_at_gt,omitempty"`

	// started at gte
	StartedAtGte *string `json:"started_at_gte,omitempty"`

	// started at in
	StartedAtIn []string `json:"started_at_in,omitempty"`

	// started at lt
	StartedAtLt *string `json:"started_at_lt,omitempty"`

	// started at lte
	StartedAtLte *string `json:"started_at_lte,omitempty"`

	// started at not
	StartedAtNot *string `json:"started_at_not,omitempty"`

	// started at not in
	StartedAtNotIn []string `json:"started_at_not_in,omitempty"`

	// status
	Status *UploadTaskStatus `json:"status,omitempty"`

	// status in
	StatusIn []UploadTaskStatus `json:"status_in,omitempty"`

	// status not
	StatusNot *UploadTaskStatus `json:"status_not,omitempty"`

	// status not in
	StatusNotIn []UploadTaskStatus `json:"status_not_in,omitempty"`

	// updated at
	UpdatedAt *string `json:"updatedAt,omitempty"`

	// updated at gt
	UpdatedAtGt *string `json:"updatedAt_gt,omitempty"`

	// updated at gte
	UpdatedAtGte *string `json:"updatedAt_gte,omitempty"`

	// updated at in
	UpdatedAtIn []string `json:"updatedAt_in,omitempty"`

	// updated at lt
	UpdatedAtLt *string `json:"updatedAt_lt,omitempty"`

	// updated at lte
	UpdatedAtLte *string `json:"updatedAt_lte,omitempty"`

	// updated at not
	UpdatedAtNot *string `json:"updatedAt_not,omitempty"`

	// updated at not in
	UpdatedAtNotIn []string `json:"updatedAt_not_in,omitempty"`
}

// Validate validates this upload task where input
func (m *UploadTaskWhereInput) Validate(formats strfmt.Registry) error {
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

	if err := m.validateResourceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceTypeIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceTypeNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceTypeNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusNotIn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UploadTaskWhereInput) validateAND(formats strfmt.Registry) error {
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

func (m *UploadTaskWhereInput) validateNOT(formats strfmt.Registry) error {
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

func (m *UploadTaskWhereInput) validateOR(formats strfmt.Registry) error {
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

func (m *UploadTaskWhereInput) validateResourceType(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceType) { // not required
		return nil
	}

	if m.ResourceType != nil {
		if err := m.ResourceType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource_type")
			}
			return err
		}
	}

	return nil
}

func (m *UploadTaskWhereInput) validateResourceTypeIn(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceTypeIn) { // not required
		return nil
	}

	for i := 0; i < len(m.ResourceTypeIn); i++ {

		if err := m.ResourceTypeIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource_type_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource_type_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *UploadTaskWhereInput) validateResourceTypeNot(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceTypeNot) { // not required
		return nil
	}

	if m.ResourceTypeNot != nil {
		if err := m.ResourceTypeNot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource_type_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource_type_not")
			}
			return err
		}
	}

	return nil
}

func (m *UploadTaskWhereInput) validateResourceTypeNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceTypeNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.ResourceTypeNotIn); i++ {

		if err := m.ResourceTypeNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource_type_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource_type_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *UploadTaskWhereInput) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
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

func (m *UploadTaskWhereInput) validateStatusIn(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusIn) { // not required
		return nil
	}

	for i := 0; i < len(m.StatusIn); i++ {

		if err := m.StatusIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *UploadTaskWhereInput) validateStatusNot(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusNot) { // not required
		return nil
	}

	if m.StatusNot != nil {
		if err := m.StatusNot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status_not")
			}
			return err
		}
	}

	return nil
}

func (m *UploadTaskWhereInput) validateStatusNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.StatusNotIn); i++ {

		if err := m.StatusNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this upload task where input based on the context it is used
func (m *UploadTaskWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateResourceType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResourceTypeIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResourceTypeNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResourceTypeNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatusIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatusNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatusNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UploadTaskWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

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

func (m *UploadTaskWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

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

func (m *UploadTaskWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

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

func (m *UploadTaskWhereInput) contextValidateResourceType(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourceType != nil {
		if err := m.ResourceType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource_type")
			}
			return err
		}
	}

	return nil
}

func (m *UploadTaskWhereInput) contextValidateResourceTypeIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ResourceTypeIn); i++ {

		if err := m.ResourceTypeIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource_type_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource_type_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *UploadTaskWhereInput) contextValidateResourceTypeNot(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourceTypeNot != nil {
		if err := m.ResourceTypeNot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource_type_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource_type_not")
			}
			return err
		}
	}

	return nil
}

func (m *UploadTaskWhereInput) contextValidateResourceTypeNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ResourceTypeNotIn); i++ {

		if err := m.ResourceTypeNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource_type_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource_type_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *UploadTaskWhereInput) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *UploadTaskWhereInput) contextValidateStatusIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StatusIn); i++ {

		if err := m.StatusIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *UploadTaskWhereInput) contextValidateStatusNot(ctx context.Context, formats strfmt.Registry) error {

	if m.StatusNot != nil {
		if err := m.StatusNot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status_not")
			}
			return err
		}
	}

	return nil
}

func (m *UploadTaskWhereInput) contextValidateStatusNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StatusNotIn); i++ {

		if err := m.StatusNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UploadTaskWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UploadTaskWhereInput) UnmarshalBinary(b []byte) error {
	var res UploadTaskWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
