// Code generated by go-swagger; DO NOT EDIT.

package rack_topo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// DeleteRackTopoReader is a Reader for the DeleteRackTopo structure.
type DeleteRackTopoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRackTopoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRackTopoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRackTopoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRackTopoOK creates a DeleteRackTopoOK with default headers values
func NewDeleteRackTopoOK() *DeleteRackTopoOK {
	return &DeleteRackTopoOK{}
}

/* DeleteRackTopoOK describes a response with status code 200, with default header values.

Ok
*/
type DeleteRackTopoOK struct {
	Payload []*models.WithTaskDeleteRackTopo
}

func (o *DeleteRackTopoOK) Error() string {
	return fmt.Sprintf("[POST /delete-rack-topo][%d] deleteRackTopoOK  %+v", 200, o.Payload)
}
func (o *DeleteRackTopoOK) GetPayload() []*models.WithTaskDeleteRackTopo {
	return o.Payload
}

func (o *DeleteRackTopoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRackTopoBadRequest creates a DeleteRackTopoBadRequest with default headers values
func NewDeleteRackTopoBadRequest() *DeleteRackTopoBadRequest {
	return &DeleteRackTopoBadRequest{}
}

/* DeleteRackTopoBadRequest describes a response with status code 400, with default header values.

DeleteRackTopoBadRequest delete rack topo bad request
*/
type DeleteRackTopoBadRequest struct {
	Payload string
}

func (o *DeleteRackTopoBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-rack-topo][%d] deleteRackTopoBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteRackTopoBadRequest) GetPayload() string {
	return o.Payload
}

func (o *DeleteRackTopoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
