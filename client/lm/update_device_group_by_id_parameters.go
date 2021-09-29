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

// NewUpdateDeviceGroupByIDParams creates a new UpdateDeviceGroupByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDeviceGroupByIDParams() *UpdateDeviceGroupByIDParams {
	return &UpdateDeviceGroupByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDeviceGroupByIDParamsWithTimeout creates a new UpdateDeviceGroupByIDParams object
// with the ability to set a timeout on a request.
func NewUpdateDeviceGroupByIDParamsWithTimeout(timeout time.Duration) *UpdateDeviceGroupByIDParams {
	return &UpdateDeviceGroupByIDParams{
		timeout: timeout,
	}
}

// NewUpdateDeviceGroupByIDParamsWithContext creates a new UpdateDeviceGroupByIDParams object
// with the ability to set a context for a request.
func NewUpdateDeviceGroupByIDParamsWithContext(ctx context.Context) *UpdateDeviceGroupByIDParams {
	return &UpdateDeviceGroupByIDParams{
		Context: ctx,
	}
}

// NewUpdateDeviceGroupByIDParamsWithHTTPClient creates a new UpdateDeviceGroupByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDeviceGroupByIDParamsWithHTTPClient(client *http.Client) *UpdateDeviceGroupByIDParams {
	return &UpdateDeviceGroupByIDParams{
		HTTPClient: client,
	}
}

/* UpdateDeviceGroupByIDParams contains all the parameters to send to the API endpoint
   for the update device group by Id operation.

   Typically these are written to a http.Request.
*/
type UpdateDeviceGroupByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-95bb3f4-dirty"
	UserAgent *string

	// Body.
	Body *models.DeviceGroup

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update device group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDeviceGroupByIDParams) WithDefaults() *UpdateDeviceGroupByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update device group by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDeviceGroupByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-95bb3f4-dirty")
	)

	val := UpdateDeviceGroupByIDParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update device group by Id params
func (o *UpdateDeviceGroupByIDParams) WithTimeout(timeout time.Duration) *UpdateDeviceGroupByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update device group by Id params
func (o *UpdateDeviceGroupByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update device group by Id params
func (o *UpdateDeviceGroupByIDParams) WithContext(ctx context.Context) *UpdateDeviceGroupByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update device group by Id params
func (o *UpdateDeviceGroupByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update device group by Id params
func (o *UpdateDeviceGroupByIDParams) WithHTTPClient(client *http.Client) *UpdateDeviceGroupByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update device group by Id params
func (o *UpdateDeviceGroupByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the update device group by Id params
func (o *UpdateDeviceGroupByIDParams) WithUserAgent(userAgent *string) *UpdateDeviceGroupByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the update device group by Id params
func (o *UpdateDeviceGroupByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithBody adds the body to the update device group by Id params
func (o *UpdateDeviceGroupByIDParams) WithBody(body *models.DeviceGroup) *UpdateDeviceGroupByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update device group by Id params
func (o *UpdateDeviceGroupByIDParams) SetBody(body *models.DeviceGroup) {
	o.Body = body
}

// WithID adds the id to the update device group by Id params
func (o *UpdateDeviceGroupByIDParams) WithID(id int32) *UpdateDeviceGroupByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update device group by Id params
func (o *UpdateDeviceGroupByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDeviceGroupByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
