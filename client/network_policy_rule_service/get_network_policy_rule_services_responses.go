// Code generated by go-swagger; DO NOT EDIT.

package network_policy_rule_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/v2/models"
)

// GetNetworkPolicyRuleServicesReader is a Reader for the GetNetworkPolicyRuleServices structure.
type GetNetworkPolicyRuleServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkPolicyRuleServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkPolicyRuleServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetNetworkPolicyRuleServicesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetNetworkPolicyRuleServicesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetNetworkPolicyRuleServicesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkPolicyRuleServicesOK creates a GetNetworkPolicyRuleServicesOK with default headers values
func NewGetNetworkPolicyRuleServicesOK() *GetNetworkPolicyRuleServicesOK {
	return &GetNetworkPolicyRuleServicesOK{}
}

/* GetNetworkPolicyRuleServicesOK describes a response with status code 200, with default header values.

GetNetworkPolicyRuleServicesOK get network policy rule services o k
*/
type GetNetworkPolicyRuleServicesOK struct {
	XTowerRequestID string

	Payload []*models.NetworkPolicyRuleService
}

func (o *GetNetworkPolicyRuleServicesOK) Error() string {
	return fmt.Sprintf("[POST /get-network-policy-rule-services][%d] getNetworkPolicyRuleServicesOK  %+v", 200, o.Payload)
}
func (o *GetNetworkPolicyRuleServicesOK) GetPayload() []*models.NetworkPolicyRuleService {
	return o.Payload
}

func (o *GetNetworkPolicyRuleServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkPolicyRuleServicesBadRequest creates a GetNetworkPolicyRuleServicesBadRequest with default headers values
func NewGetNetworkPolicyRuleServicesBadRequest() *GetNetworkPolicyRuleServicesBadRequest {
	return &GetNetworkPolicyRuleServicesBadRequest{}
}

/* GetNetworkPolicyRuleServicesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetNetworkPolicyRuleServicesBadRequest struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetNetworkPolicyRuleServicesBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-network-policy-rule-services][%d] getNetworkPolicyRuleServicesBadRequest  %+v", 400, o.Payload)
}
func (o *GetNetworkPolicyRuleServicesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNetworkPolicyRuleServicesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkPolicyRuleServicesNotFound creates a GetNetworkPolicyRuleServicesNotFound with default headers values
func NewGetNetworkPolicyRuleServicesNotFound() *GetNetworkPolicyRuleServicesNotFound {
	return &GetNetworkPolicyRuleServicesNotFound{}
}

/* GetNetworkPolicyRuleServicesNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetNetworkPolicyRuleServicesNotFound struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetNetworkPolicyRuleServicesNotFound) Error() string {
	return fmt.Sprintf("[POST /get-network-policy-rule-services][%d] getNetworkPolicyRuleServicesNotFound  %+v", 404, o.Payload)
}
func (o *GetNetworkPolicyRuleServicesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNetworkPolicyRuleServicesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkPolicyRuleServicesInternalServerError creates a GetNetworkPolicyRuleServicesInternalServerError with default headers values
func NewGetNetworkPolicyRuleServicesInternalServerError() *GetNetworkPolicyRuleServicesInternalServerError {
	return &GetNetworkPolicyRuleServicesInternalServerError{}
}

/* GetNetworkPolicyRuleServicesInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type GetNetworkPolicyRuleServicesInternalServerError struct {
	XTowerRequestID string

	Payload *models.ErrorBody
}

func (o *GetNetworkPolicyRuleServicesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /get-network-policy-rule-services][%d] getNetworkPolicyRuleServicesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetNetworkPolicyRuleServicesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNetworkPolicyRuleServicesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-tower-request-id
	hdrXTowerRequestID := response.GetHeader("x-tower-request-id")

	if hdrXTowerRequestID != "" {
		o.XTowerRequestID = hdrXTowerRequestID
	}

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
