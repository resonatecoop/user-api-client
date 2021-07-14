// Code generated by go-swagger; DO NOT EDIT.

package usergroups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/resonatecoop/user-api-client/models"
)

// ResonateUserDeleteUserGroupReader is a Reader for the ResonateUserDeleteUserGroup structure.
type ResonateUserDeleteUserGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResonateUserDeleteUserGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResonateUserDeleteUserGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewResonateUserDeleteUserGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewResonateUserDeleteUserGroupOK creates a ResonateUserDeleteUserGroupOK with default headers values
func NewResonateUserDeleteUserGroupOK() *ResonateUserDeleteUserGroupOK {
	return &ResonateUserDeleteUserGroupOK{}
}

/* ResonateUserDeleteUserGroupOK describes a response with status code 200, with default header values.

A successful response.
*/
type ResonateUserDeleteUserGroupOK struct {
	Payload models.UserEmpty
}

func (o *ResonateUserDeleteUserGroupOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/usergroup/{id}][%d] resonateUserDeleteUserGroupOK  %+v", 200, o.Payload)
}
func (o *ResonateUserDeleteUserGroupOK) GetPayload() models.UserEmpty {
	return o.Payload
}

func (o *ResonateUserDeleteUserGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResonateUserDeleteUserGroupDefault creates a ResonateUserDeleteUserGroupDefault with default headers values
func NewResonateUserDeleteUserGroupDefault(code int) *ResonateUserDeleteUserGroupDefault {
	return &ResonateUserDeleteUserGroupDefault{
		_statusCode: code,
	}
}

/* ResonateUserDeleteUserGroupDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ResonateUserDeleteUserGroupDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the resonate user delete user group default response
func (o *ResonateUserDeleteUserGroupDefault) Code() int {
	return o._statusCode
}

func (o *ResonateUserDeleteUserGroupDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/usergroup/{id}][%d] ResonateUser_DeleteUserGroup default  %+v", o._statusCode, o.Payload)
}
func (o *ResonateUserDeleteUserGroupDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *ResonateUserDeleteUserGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
