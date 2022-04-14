// Code generated by go-swagger; DO NOT EDIT.

package membership

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/resonatecoop/user-api-client/models"
)

// ResonateUserGetUserMembershipReader is a Reader for the ResonateUserGetUserMembership structure.
type ResonateUserGetUserMembershipReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResonateUserGetUserMembershipReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResonateUserGetUserMembershipOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewResonateUserGetUserMembershipDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewResonateUserGetUserMembershipOK creates a ResonateUserGetUserMembershipOK with default headers values
func NewResonateUserGetUserMembershipOK() *ResonateUserGetUserMembershipOK {
	return &ResonateUserGetUserMembershipOK{}
}

/* ResonateUserGetUserMembershipOK describes a response with status code 200, with default header values.

A successful response.
*/
type ResonateUserGetUserMembershipOK struct {
	Payload *models.UserUserMembershipResponse
}

func (o *ResonateUserGetUserMembershipOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/{id}/membership][%d] resonateUserGetUserMembershipOK  %+v", 200, o.Payload)
}
func (o *ResonateUserGetUserMembershipOK) GetPayload() *models.UserUserMembershipResponse {
	return o.Payload
}

func (o *ResonateUserGetUserMembershipOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserUserMembershipResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResonateUserGetUserMembershipDefault creates a ResonateUserGetUserMembershipDefault with default headers values
func NewResonateUserGetUserMembershipDefault(code int) *ResonateUserGetUserMembershipDefault {
	return &ResonateUserGetUserMembershipDefault{
		_statusCode: code,
	}
}

/* ResonateUserGetUserMembershipDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ResonateUserGetUserMembershipDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the resonate user get user membership default response
func (o *ResonateUserGetUserMembershipDefault) Code() int {
	return o._statusCode
}

func (o *ResonateUserGetUserMembershipDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/{id}/membership][%d] ResonateUser_GetUserMembership default  %+v", o._statusCode, o.Payload)
}
func (o *ResonateUserGetUserMembershipDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *ResonateUserGetUserMembershipDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}