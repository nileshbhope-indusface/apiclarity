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

// GetAPIInventoryAPIIDAPIInfoReader is a Reader for the GetAPIInventoryAPIIDAPIInfo structure.
type GetAPIInventoryAPIIDAPIInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIInventoryAPIIDAPIInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIInventoryAPIIDAPIInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAPIInventoryAPIIDAPIInfoDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAPIInventoryAPIIDAPIInfoOK creates a GetAPIInventoryAPIIDAPIInfoOK with default headers values
func NewGetAPIInventoryAPIIDAPIInfoOK() *GetAPIInventoryAPIIDAPIInfoOK {
	return &GetAPIInventoryAPIIDAPIInfoOK{}
}

/* GetAPIInventoryAPIIDAPIInfoOK describes a response with status code 200, with default header values.

Success
*/
type GetAPIInventoryAPIIDAPIInfoOK struct {
	Payload *models.APIInfoWithType
}

func (o *GetAPIInventoryAPIIDAPIInfoOK) Error() string {
	return fmt.Sprintf("[GET /apiInventory/{apiId}/apiInfo][%d] getApiInventoryApiIdApiInfoOK  %+v", 200, o.Payload)
}
func (o *GetAPIInventoryAPIIDAPIInfoOK) GetPayload() *models.APIInfoWithType {
	return o.Payload
}

func (o *GetAPIInventoryAPIIDAPIInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIInfoWithType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIInventoryAPIIDAPIInfoDefault creates a GetAPIInventoryAPIIDAPIInfoDefault with default headers values
func NewGetAPIInventoryAPIIDAPIInfoDefault(code int) *GetAPIInventoryAPIIDAPIInfoDefault {
	return &GetAPIInventoryAPIIDAPIInfoDefault{
		_statusCode: code,
	}
}

/* GetAPIInventoryAPIIDAPIInfoDefault describes a response with status code -1, with default header values.

unknown error
*/
type GetAPIInventoryAPIIDAPIInfoDefault struct {
	_statusCode int

	Payload *models.APIResponse
}

// Code gets the status code for the get API inventory API ID API info default response
func (o *GetAPIInventoryAPIIDAPIInfoDefault) Code() int {
	return o._statusCode
}

func (o *GetAPIInventoryAPIIDAPIInfoDefault) Error() string {
	return fmt.Sprintf("[GET /apiInventory/{apiId}/apiInfo][%d] GetAPIInventoryAPIIDAPIInfo default  %+v", o._statusCode, o.Payload)
}
func (o *GetAPIInventoryAPIIDAPIInfoDefault) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *GetAPIInventoryAPIIDAPIInfoDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
