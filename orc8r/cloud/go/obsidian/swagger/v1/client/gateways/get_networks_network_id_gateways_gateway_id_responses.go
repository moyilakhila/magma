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

// GetNetworksNetworkIDGatewaysGatewayIDReader is a Reader for the GetNetworksNetworkIDGatewaysGatewayID structure.
type GetNetworksNetworkIDGatewaysGatewayIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworksNetworkIDGatewaysGatewayIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworksNetworkIDGatewaysGatewayIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNetworksNetworkIDGatewaysGatewayIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworksNetworkIDGatewaysGatewayIDOK creates a GetNetworksNetworkIDGatewaysGatewayIDOK with default headers values
func NewGetNetworksNetworkIDGatewaysGatewayIDOK() *GetNetworksNetworkIDGatewaysGatewayIDOK {
	return &GetNetworksNetworkIDGatewaysGatewayIDOK{}
}

/*GetNetworksNetworkIDGatewaysGatewayIDOK handles this case with default header values.

The requested gateway
*/
type GetNetworksNetworkIDGatewaysGatewayIDOK struct {
	Payload *models.MagmadGateway
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDOK) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/gateways/{gateway_id}][%d] getNetworksNetworkIdGatewaysGatewayIdOK  %+v", 200, o.Payload)
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDOK) GetPayload() *models.MagmadGateway {
	return o.Payload
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MagmadGateway)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworksNetworkIDGatewaysGatewayIDDefault creates a GetNetworksNetworkIDGatewaysGatewayIDDefault with default headers values
func NewGetNetworksNetworkIDGatewaysGatewayIDDefault(code int) *GetNetworksNetworkIDGatewaysGatewayIDDefault {
	return &GetNetworksNetworkIDGatewaysGatewayIDDefault{
		_statusCode: code,
	}
}

/*GetNetworksNetworkIDGatewaysGatewayIDDefault handles this case with default header values.

Unexpected Error
*/
type GetNetworksNetworkIDGatewaysGatewayIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get networks network ID gateways gateway ID default response
func (o *GetNetworksNetworkIDGatewaysGatewayIDDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDDefault) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/gateways/{gateway_id}][%d] GetNetworksNetworkIDGatewaysGatewayID default  %+v", o._statusCode, o.Payload)
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworksNetworkIDGatewaysGatewayIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
