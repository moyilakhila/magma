// Code generated by go-swagger; DO NOT EDIT.

package gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// PutNetworksNetworkIDGatewaysGatewayIDTierReader is a Reader for the PutNetworksNetworkIDGatewaysGatewayIDTier structure.
type PutNetworksNetworkIDGatewaysGatewayIDTierReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutNetworksNetworkIDGatewaysGatewayIDTierReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutNetworksNetworkIDGatewaysGatewayIDTierNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutNetworksNetworkIDGatewaysGatewayIDTierDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutNetworksNetworkIDGatewaysGatewayIDTierNoContent creates a PutNetworksNetworkIDGatewaysGatewayIDTierNoContent with default headers values
func NewPutNetworksNetworkIDGatewaysGatewayIDTierNoContent() *PutNetworksNetworkIDGatewaysGatewayIDTierNoContent {
	return &PutNetworksNetworkIDGatewaysGatewayIDTierNoContent{}
}

/*PutNetworksNetworkIDGatewaysGatewayIDTierNoContent handles this case with default header values.

Success
*/
type PutNetworksNetworkIDGatewaysGatewayIDTierNoContent struct {
}

func (o *PutNetworksNetworkIDGatewaysGatewayIDTierNoContent) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/gateways/{gateway_id}/tier][%d] putNetworksNetworkIdGatewaysGatewayIdTierNoContent ", 204)
}

func (o *PutNetworksNetworkIDGatewaysGatewayIDTierNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutNetworksNetworkIDGatewaysGatewayIDTierDefault creates a PutNetworksNetworkIDGatewaysGatewayIDTierDefault with default headers values
func NewPutNetworksNetworkIDGatewaysGatewayIDTierDefault(code int) *PutNetworksNetworkIDGatewaysGatewayIDTierDefault {
	return &PutNetworksNetworkIDGatewaysGatewayIDTierDefault{
		_statusCode: code,
	}
}

/*PutNetworksNetworkIDGatewaysGatewayIDTierDefault handles this case with default header values.

Unexpected Error
*/
type PutNetworksNetworkIDGatewaysGatewayIDTierDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put networks network ID gateways gateway ID tier default response
func (o *PutNetworksNetworkIDGatewaysGatewayIDTierDefault) Code() int {
	return o._statusCode
}

func (o *PutNetworksNetworkIDGatewaysGatewayIDTierDefault) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/gateways/{gateway_id}/tier][%d] PutNetworksNetworkIDGatewaysGatewayIDTier default  %+v", o._statusCode, o.Payload)
}

func (o *PutNetworksNetworkIDGatewaysGatewayIDTierDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutNetworksNetworkIDGatewaysGatewayIDTierDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
