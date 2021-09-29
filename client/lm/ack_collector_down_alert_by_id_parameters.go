// Code generated by go-swagger; DO NOT EDIT.

package lm

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
	"github.com/go-openapi/swag"

	"github.com/logicmonitor/lm-sdk-go/models"
)

// NewAckCollectorDownAlertByIDParams creates a new AckCollectorDownAlertByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAckCollectorDownAlertByIDParams() *AckCollectorDownAlertByIDParams {
	return &AckCollectorDownAlertByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAckCollectorDownAlertByIDParamsWithTimeout creates a new AckCollectorDownAlertByIDParams object
// with the ability to set a timeout on a request.
func NewAckCollectorDownAlertByIDParamsWithTimeout(timeout time.Duration) *AckCollectorDownAlertByIDParams {
	return &AckCollectorDownAlertByIDParams{
		timeout: timeout,
	}
}

// NewAckCollectorDownAlertByIDParamsWithContext creates a new AckCollectorDownAlertByIDParams object
// with the ability to set a context for a request.
func NewAckCollectorDownAlertByIDParamsWithContext(ctx context.Context) *AckCollectorDownAlertByIDParams {
	return &AckCollectorDownAlertByIDParams{
		Context: ctx,
	}
}

// NewAckCollectorDownAlertByIDParamsWithHTTPClient creates a new AckCollectorDownAlertByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewAckCollectorDownAlertByIDParamsWithHTTPClient(client *http.Client) *AckCollectorDownAlertByIDParams {
	return &AckCollectorDownAlertByIDParams{
		HTTPClient: client,
	}
}

/* AckCollectorDownAlertByIDParams contains all the parameters to send to the API endpoint
   for the ack collector down alert by Id operation.

   Typically these are written to a http.Request.
*/
type AckCollectorDownAlertByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-95bb3f4-dirty"
	UserAgent *string

	// Body.
	Body *models.AckCollectorDown

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ack collector down alert by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AckCollectorDownAlertByIDParams) WithDefaults() *AckCollectorDownAlertByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ack collector down alert by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AckCollectorDownAlertByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-95bb3f4-dirty")
	)

	val := AckCollectorDownAlertByIDParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the ack collector down alert by Id params
func (o *AckCollectorDownAlertByIDParams) WithTimeout(timeout time.Duration) *AckCollectorDownAlertByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ack collector down alert by Id params
func (o *AckCollectorDownAlertByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ack collector down alert by Id params
func (o *AckCollectorDownAlertByIDParams) WithContext(ctx context.Context) *AckCollectorDownAlertByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ack collector down alert by Id params
func (o *AckCollectorDownAlertByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ack collector down alert by Id params
func (o *AckCollectorDownAlertByIDParams) WithHTTPClient(client *http.Client) *AckCollectorDownAlertByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ack collector down alert by Id params
func (o *AckCollectorDownAlertByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the ack collector down alert by Id params
func (o *AckCollectorDownAlertByIDParams) WithUserAgent(userAgent *string) *AckCollectorDownAlertByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the ack collector down alert by Id params
func (o *AckCollectorDownAlertByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the ack collector down alert by Id params
func (o *AckCollectorDownAlertByIDParams) WithBody(body *models.AckCollectorDown) *AckCollectorDownAlertByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the ack collector down alert by Id params
func (o *AckCollectorDownAlertByIDParams) SetBody(body *models.AckCollectorDown) {
	o.Body = body
}

// WithID adds the id to the ack collector down alert by Id params
func (o *AckCollectorDownAlertByIDParams) WithID(id int32) *AckCollectorDownAlertByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ack collector down alert by Id params
func (o *AckCollectorDownAlertByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AckCollectorDownAlertByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UserAgent != nil {

		// header param User-Agent
		if err := r.SetHeaderParam("User-Agent", *o.UserAgent); err != nil {
			return err
		}
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
