// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openclarity/apiclarity/api/server/models"
)

// GetControlGatewaysGatewayIDOKCode is the HTTP code returned for type GetControlGatewaysGatewayIDOK
const GetControlGatewaysGatewayIDOKCode int = 200

/*GetControlGatewaysGatewayIDOK Gateway information

swagger:response getControlGatewaysGatewayIdOK
*/
type GetControlGatewaysGatewayIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.APIGateway `json:"body,omitempty"`
}

// NewGetControlGatewaysGatewayIDOK creates GetControlGatewaysGatewayIDOK with default headers values
func NewGetControlGatewaysGatewayIDOK() *GetControlGatewaysGatewayIDOK {

	return &GetControlGatewaysGatewayIDOK{}
}

// WithPayload adds the payload to the get control gateways gateway Id o k response
func (o *GetControlGatewaysGatewayIDOK) WithPayload(payload *models.APIGateway) *GetControlGatewaysGatewayIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get control gateways gateway Id o k response
func (o *GetControlGatewaysGatewayIDOK) SetPayload(payload *models.APIGateway) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetControlGatewaysGatewayIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetControlGatewaysGatewayIDNotFoundCode is the HTTP code returned for type GetControlGatewaysGatewayIDNotFound
const GetControlGatewaysGatewayIDNotFoundCode int = 404

/*GetControlGatewaysGatewayIDNotFound API Gateway not found

swagger:response getControlGatewaysGatewayIdNotFound
*/
type GetControlGatewaysGatewayIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetControlGatewaysGatewayIDNotFound creates GetControlGatewaysGatewayIDNotFound with default headers values
func NewGetControlGatewaysGatewayIDNotFound() *GetControlGatewaysGatewayIDNotFound {

	return &GetControlGatewaysGatewayIDNotFound{}
}

// WithPayload adds the payload to the get control gateways gateway Id not found response
func (o *GetControlGatewaysGatewayIDNotFound) WithPayload(payload *models.APIResponse) *GetControlGatewaysGatewayIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get control gateways gateway Id not found response
func (o *GetControlGatewaysGatewayIDNotFound) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetControlGatewaysGatewayIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetControlGatewaysGatewayIDDefault unknown error

swagger:response getControlGatewaysGatewayIdDefault
*/
type GetControlGatewaysGatewayIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewGetControlGatewaysGatewayIDDefault creates GetControlGatewaysGatewayIDDefault with default headers values
func NewGetControlGatewaysGatewayIDDefault(code int) *GetControlGatewaysGatewayIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetControlGatewaysGatewayIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get control gateways gateway ID default response
func (o *GetControlGatewaysGatewayIDDefault) WithStatusCode(code int) *GetControlGatewaysGatewayIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get control gateways gateway ID default response
func (o *GetControlGatewaysGatewayIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get control gateways gateway ID default response
func (o *GetControlGatewaysGatewayIDDefault) WithPayload(payload *models.APIResponse) *GetControlGatewaysGatewayIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get control gateways gateway ID default response
func (o *GetControlGatewaysGatewayIDDefault) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetControlGatewaysGatewayIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
