// Code generated by go-swagger; DO NOT EDIT.

package usergroups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewResonateUserGetUserGroupParams creates a new ResonateUserGetUserGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResonateUserGetUserGroupParams() *ResonateUserGetUserGroupParams {
	return &ResonateUserGetUserGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResonateUserGetUserGroupParamsWithTimeout creates a new ResonateUserGetUserGroupParams object
// with the ability to set a timeout on a request.
func NewResonateUserGetUserGroupParamsWithTimeout(timeout time.Duration) *ResonateUserGetUserGroupParams {
	return &ResonateUserGetUserGroupParams{
		timeout: timeout,
	}
}

// NewResonateUserGetUserGroupParamsWithContext creates a new ResonateUserGetUserGroupParams object
// with the ability to set a context for a request.
func NewResonateUserGetUserGroupParamsWithContext(ctx context.Context) *ResonateUserGetUserGroupParams {
	return &ResonateUserGetUserGroupParams{
		Context: ctx,
	}
}

// NewResonateUserGetUserGroupParamsWithHTTPClient creates a new ResonateUserGetUserGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewResonateUserGetUserGroupParamsWithHTTPClient(client *http.Client) *ResonateUserGetUserGroupParams {
	return &ResonateUserGetUserGroupParams{
		HTTPClient: client,
	}
}

/* ResonateUserGetUserGroupParams contains all the parameters to send to the API endpoint
   for the resonate user get user group operation.

   Typically these are written to a http.Request.
*/
type ResonateUserGetUserGroupParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the resonate user get user group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResonateUserGetUserGroupParams) WithDefaults() *ResonateUserGetUserGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the resonate user get user group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResonateUserGetUserGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the resonate user get user group params
func (o *ResonateUserGetUserGroupParams) WithTimeout(timeout time.Duration) *ResonateUserGetUserGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resonate user get user group params
func (o *ResonateUserGetUserGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resonate user get user group params
func (o *ResonateUserGetUserGroupParams) WithContext(ctx context.Context) *ResonateUserGetUserGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resonate user get user group params
func (o *ResonateUserGetUserGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resonate user get user group params
func (o *ResonateUserGetUserGroupParams) WithHTTPClient(client *http.Client) *ResonateUserGetUserGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resonate user get user group params
func (o *ResonateUserGetUserGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the resonate user get user group params
func (o *ResonateUserGetUserGroupParams) WithID(id string) *ResonateUserGetUserGroupParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the resonate user get user group params
func (o *ResonateUserGetUserGroupParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ResonateUserGetUserGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
