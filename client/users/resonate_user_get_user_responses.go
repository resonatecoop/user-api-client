// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/resonatecoop/user-api-client/models"
)

// ResonateUserGetUserReader is a Reader for the ResonateUserGetUser structure.
type ResonateUserGetUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResonateUserGetUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResonateUserGetUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewResonateUserGetUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewResonateUserGetUserOK creates a ResonateUserGetUserOK with default headers values
func NewResonateUserGetUserOK() *ResonateUserGetUserOK {
	return &ResonateUserGetUserOK{}
}

/* ResonateUserGetUserOK describes a response with status code 200, with default header values.

A successful response.
*/
type ResonateUserGetUserOK struct {
	Payload *models.UserUserPublicResponse
}

func (o *ResonateUserGetUserOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/{id}][%d] resonateUserGetUserOK  %+v", 200, o.Payload)
}
func (o *ResonateUserGetUserOK) GetPayload() *models.UserUserPublicResponse {
	return o.Payload
}

func (o *ResonateUserGetUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserUserPublicResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResonateUserGetUserDefault creates a ResonateUserGetUserDefault with default headers values
func NewResonateUserGetUserDefault(code int) *ResonateUserGetUserDefault {
	return &ResonateUserGetUserDefault{
		_statusCode: code,
	}
}

/* ResonateUserGetUserDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ResonateUserGetUserDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the resonate user get user default response
func (o *ResonateUserGetUserDefault) Code() int {
	return o._statusCode
}

func (o *ResonateUserGetUserDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/{id}][%d] ResonateUser_GetUser default  %+v", o._statusCode, o.Payload)
}
func (o *ResonateUserGetUserDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *ResonateUserGetUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
