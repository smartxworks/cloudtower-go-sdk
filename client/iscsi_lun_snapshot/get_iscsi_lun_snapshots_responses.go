// Code generated by go-swagger; DO NOT EDIT.

package iscsi_lun_snapshot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// GetIscsiLunSnapshotsReader is a Reader for the GetIscsiLunSnapshots structure.
type GetIscsiLunSnapshotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIscsiLunSnapshotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIscsiLunSnapshotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIscsiLunSnapshotsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIscsiLunSnapshotsOK creates a GetIscsiLunSnapshotsOK with default headers values
func NewGetIscsiLunSnapshotsOK() *GetIscsiLunSnapshotsOK {
	return &GetIscsiLunSnapshotsOK{}
}

/* GetIscsiLunSnapshotsOK describes a response with status code 200, with default header values.

Ok
*/
type GetIscsiLunSnapshotsOK struct {
	Payload []*models.IscsiLunSnapshot
}

func (o *GetIscsiLunSnapshotsOK) Error() string {
	return fmt.Sprintf("[POST /get-iscsi-lun-snapshots][%d] getIscsiLunSnapshotsOK  %+v", 200, o.Payload)
}
func (o *GetIscsiLunSnapshotsOK) GetPayload() []*models.IscsiLunSnapshot {
	return o.Payload
}

func (o *GetIscsiLunSnapshotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIscsiLunSnapshotsBadRequest creates a GetIscsiLunSnapshotsBadRequest with default headers values
func NewGetIscsiLunSnapshotsBadRequest() *GetIscsiLunSnapshotsBadRequest {
	return &GetIscsiLunSnapshotsBadRequest{}
}

/* GetIscsiLunSnapshotsBadRequest describes a response with status code 400, with default header values.

GetIscsiLunSnapshotsBadRequest get iscsi lun snapshots bad request
*/
type GetIscsiLunSnapshotsBadRequest struct {
	Payload string
}

func (o *GetIscsiLunSnapshotsBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-iscsi-lun-snapshots][%d] getIscsiLunSnapshotsBadRequest  %+v", 400, o.Payload)
}
func (o *GetIscsiLunSnapshotsBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetIscsiLunSnapshotsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
