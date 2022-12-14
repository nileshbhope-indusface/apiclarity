// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openclarity/apiclarity/api/client/models"
)

// GetControlGatewaysGatewayIDReader is a Reader for the GetControlGatewaysGatewayID structure.
type GetControlGatewaysGatewayIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetControlGatewaysGatewayIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetControlGatewaysGatewayIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetControlGatewaysGatewayIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetControlGatewaysGatewayIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetControlGatewaysGatewayIDOK creates a GetControlGatewaysGatewayIDOK with default headers values
func NewGetControlGatewaysGatewayIDOK() *GetControlGatewaysGatewayIDOK {
	return &GetControlGatewaysGatewayIDOK{}
}

/* GetControlGatewaysGatewayIDOK describes a response with status code 200, with default header values.

Gateway information
*/
type GetControlGatewaysGatewayIDOK struct {
	Payload *models.APIGateway
}

func (o *GetControlGatewaysGatewayIDOK) Error() string {
	return fmt.Sprintf("[GET /control/gateways/{gatewayId}][%d] getControlGatewaysGatewayIdOK  %+v", 200, o.Payload)
}
func (o *GetControlGatewaysGatewayIDOK) GetPayload() *models.APIGateway {
	return o.Payload
}

func (o *GetControlGatewaysGatewayIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIGateway)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetControlGatewaysGatewayIDNotFound creates a GetControlGatewaysGatewayIDNotFound with default headers values
func NewGetControlGatewaysGatewayIDNotFound() *GetControlGatewaysGatewayIDNotFound {
	return &GetControlGatewaysGatewayIDNotFound{}
}

/* GetControlGatewaysGatewayIDNotFound describes a response with status code 404, with default header values.

API Gateway not found
*/
type GetControlGatewaysGatewayIDNotFound struct {
	Payload *models.APIResponse
}

func (o *GetControlGatewaysGatewayIDNotFound) Error() string {
	return fmt.Sprintf("[GET /control/gateways/{gatewayId}][%d] getControlGatewaysGatewayIdNotFound  %+v", 404, o.Payload)
}
func (o *GetControlGatewaysGatewayIDNotFound) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *GetControlGatewaysGatewayIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetControlGatewaysGatewayIDDefault creates a GetControlGatewaysGatewayIDDefault with default headers values
func NewGetControlGatewaysGatewayIDDefault(code int) *GetControlGatewaysGatewayIDDefault {
	return &GetControlGatewaysGatewayIDDefault{
		_statusCode: code,
	}
}

/* GetControlGatewaysGatewayIDDefault describes a response with status code -1, with default header values.

unknown error
*/
type GetControlGatewaysGatewayIDDefault struct {
	_statusCode int

	Payload *models.APIResponse
}

// Code gets the status code for the get control gateways gateway ID default response
func (o *GetControlGatewaysGatewayIDDefault) Code() int {
	return o._statusCode
}

func (o *GetControlGatewaysGatewayIDDefault) Error() string {
	return fmt.Sprintf("[GET /control/gateways/{gatewayId}][%d] GetControlGatewaysGatewayID default  %+v", o._statusCode, o.Payload)
}
func (o *GetControlGatewaysGatewayIDDefault) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *GetControlGatewaysGatewayIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
