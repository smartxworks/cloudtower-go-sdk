// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// GetClustersConnectionReader is a Reader for the GetClustersConnection structure.
type GetClustersConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClustersConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClustersConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetClustersConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetClustersConnectionOK creates a GetClustersConnectionOK with default headers values
func NewGetClustersConnectionOK() *GetClustersConnectionOK {
	return &GetClustersConnectionOK{}
}

/* GetClustersConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetClustersConnectionOK struct {
	Payload *models.ClusterConnection
}

func (o *GetClustersConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-clusters-connection][%d] getClustersConnectionOK  %+v", 200, o.Payload)
}
func (o *GetClustersConnectionOK) GetPayload() *models.ClusterConnection {
	return o.Payload
}

func (o *GetClustersConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClustersConnectionBadRequest creates a GetClustersConnectionBadRequest with default headers values
func NewGetClustersConnectionBadRequest() *GetClustersConnectionBadRequest {
	return &GetClustersConnectionBadRequest{}
}

/* GetClustersConnectionBadRequest describes a response with status code 400, with default header values.

GetClustersConnectionBadRequest get clusters connection bad request
*/
type GetClustersConnectionBadRequest struct {
	Payload string
}

func (o *GetClustersConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-clusters-connection][%d] getClustersConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetClustersConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetClustersConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
