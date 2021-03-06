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

// ResonateUserGetUserCreditsReader is a Reader for the ResonateUserGetUserCredits structure.
type ResonateUserGetUserCreditsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResonateUserGetUserCreditsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResonateUserGetUserCreditsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewResonateUserGetUserCreditsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewResonateUserGetUserCreditsOK creates a ResonateUserGetUserCreditsOK with default headers values
func NewResonateUserGetUserCreditsOK() *ResonateUserGetUserCreditsOK {
	return &ResonateUserGetUserCreditsOK{}
}

/* ResonateUserGetUserCreditsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ResonateUserGetUserCreditsOK struct {
	Payload *models.UserUserCreditResponse
}

func (o *ResonateUserGetUserCreditsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/{id}/credits][%d] resonateUserGetUserCreditsOK  %+v", 200, o.Payload)
}
func (o *ResonateUserGetUserCreditsOK) GetPayload() *models.UserUserCreditResponse {
	return o.Payload
}

func (o *ResonateUserGetUserCreditsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserUserCreditResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResonateUserGetUserCreditsDefault creates a ResonateUserGetUserCreditsDefault with default headers values
func NewResonateUserGetUserCreditsDefault(code int) *ResonateUserGetUserCreditsDefault {
	return &ResonateUserGetUserCreditsDefault{
		_statusCode: code,
	}
}

/* ResonateUserGetUserCreditsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ResonateUserGetUserCreditsDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the resonate user get user credits default response
func (o *ResonateUserGetUserCreditsDefault) Code() int {
	return o._statusCode
}

func (o *ResonateUserGetUserCreditsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/{id}/credits][%d] ResonateUser_GetUserCredits default  %+v", o._statusCode, o.Payload)
}
func (o *ResonateUserGetUserCreditsDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *ResonateUserGetUserCreditsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
