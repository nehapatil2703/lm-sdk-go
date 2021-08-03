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
)

// NewDeleteDeviceGroupClusterAlertConfByIDParams creates a new DeleteDeviceGroupClusterAlertConfByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDeviceGroupClusterAlertConfByIDParams() *DeleteDeviceGroupClusterAlertConfByIDParams {
	return &DeleteDeviceGroupClusterAlertConfByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDeviceGroupClusterAlertConfByIDParamsWithTimeout creates a new DeleteDeviceGroupClusterAlertConfByIDParams object
// with the ability to set a timeout on a request.
func NewDeleteDeviceGroupClusterAlertConfByIDParamsWithTimeout(timeout time.Duration) *DeleteDeviceGroupClusterAlertConfByIDParams {
	return &DeleteDeviceGroupClusterAlertConfByIDParams{
		timeout: timeout,
	}
}

// NewDeleteDeviceGroupClusterAlertConfByIDParamsWithContext creates a new DeleteDeviceGroupClusterAlertConfByIDParams object
// with the ability to set a context for a request.
func NewDeleteDeviceGroupClusterAlertConfByIDParamsWithContext(ctx context.Context) *DeleteDeviceGroupClusterAlertConfByIDParams {
	return &DeleteDeviceGroupClusterAlertConfByIDParams{
		Context: ctx,
	}
}

// NewDeleteDeviceGroupClusterAlertConfByIDParamsWithHTTPClient creates a new DeleteDeviceGroupClusterAlertConfByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDeviceGroupClusterAlertConfByIDParamsWithHTTPClient(client *http.Client) *DeleteDeviceGroupClusterAlertConfByIDParams {
	return &DeleteDeviceGroupClusterAlertConfByIDParams{
		HTTPClient: client,
	}
}

/* DeleteDeviceGroupClusterAlertConfByIDParams contains all the parameters to send to the API endpoint
   for the delete device group cluster alert conf by Id operation.

   Typically these are written to a http.Request.
*/
type DeleteDeviceGroupClusterAlertConfByIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-4-gd30bcfd-dirty"
	UserAgent *string

	// DeviceGroupID.
	//
	// Format: int32
	DeviceGroupID int32

	// ID.
	//
	// Format: int32
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete device group cluster alert conf by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDeviceGroupClusterAlertConfByIDParams) WithDefaults() *DeleteDeviceGroupClusterAlertConfByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete device group cluster alert conf by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDeviceGroupClusterAlertConfByIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-4-gd30bcfd-dirty")
	)

	val := DeleteDeviceGroupClusterAlertConfByIDParams{
		UserAgent: &userAgentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete device group cluster alert conf by Id params
func (o *DeleteDeviceGroupClusterAlertConfByIDParams) WithTimeout(timeout time.Duration) *DeleteDeviceGroupClusterAlertConfByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete device group cluster alert conf by Id params
func (o *DeleteDeviceGroupClusterAlertConfByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete device group cluster alert conf by Id params
func (o *DeleteDeviceGroupClusterAlertConfByIDParams) WithContext(ctx context.Context) *DeleteDeviceGroupClusterAlertConfByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete device group cluster alert conf by Id params
func (o *DeleteDeviceGroupClusterAlertConfByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete device group cluster alert conf by Id params
func (o *DeleteDeviceGroupClusterAlertConfByIDParams) WithHTTPClient(client *http.Client) *DeleteDeviceGroupClusterAlertConfByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete device group cluster alert conf by Id params
func (o *DeleteDeviceGroupClusterAlertConfByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the delete device group cluster alert conf by Id params
func (o *DeleteDeviceGroupClusterAlertConfByIDParams) WithUserAgent(userAgent *string) *DeleteDeviceGroupClusterAlertConfByIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the delete device group cluster alert conf by Id params
func (o *DeleteDeviceGroupClusterAlertConfByIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithDeviceGroupID adds the deviceGroupID to the delete device group cluster alert conf by Id params
func (o *DeleteDeviceGroupClusterAlertConfByIDParams) WithDeviceGroupID(deviceGroupID int32) *DeleteDeviceGroupClusterAlertConfByIDParams {
	o.SetDeviceGroupID(deviceGroupID)
	return o
}

// SetDeviceGroupID adds the deviceGroupId to the delete device group cluster alert conf by Id params
func (o *DeleteDeviceGroupClusterAlertConfByIDParams) SetDeviceGroupID(deviceGroupID int32) {
	o.DeviceGroupID = deviceGroupID
}

// WithID adds the id to the delete device group cluster alert conf by Id params
func (o *DeleteDeviceGroupClusterAlertConfByIDParams) WithID(id int32) *DeleteDeviceGroupClusterAlertConfByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete device group cluster alert conf by Id params
func (o *DeleteDeviceGroupClusterAlertConfByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDeviceGroupClusterAlertConfByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param deviceGroupId
	if err := r.SetPathParam("deviceGroupId", swag.FormatInt32(o.DeviceGroupID)); err != nil {
		return err
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
