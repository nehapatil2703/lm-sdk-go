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

// NewGetImmediateDeviceListByDeviceGroupIDParams creates a new GetImmediateDeviceListByDeviceGroupIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetImmediateDeviceListByDeviceGroupIDParams() *GetImmediateDeviceListByDeviceGroupIDParams {
	return &GetImmediateDeviceListByDeviceGroupIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetImmediateDeviceListByDeviceGroupIDParamsWithTimeout creates a new GetImmediateDeviceListByDeviceGroupIDParams object
// with the ability to set a timeout on a request.
func NewGetImmediateDeviceListByDeviceGroupIDParamsWithTimeout(timeout time.Duration) *GetImmediateDeviceListByDeviceGroupIDParams {
	return &GetImmediateDeviceListByDeviceGroupIDParams{
		timeout: timeout,
	}
}

// NewGetImmediateDeviceListByDeviceGroupIDParamsWithContext creates a new GetImmediateDeviceListByDeviceGroupIDParams object
// with the ability to set a context for a request.
func NewGetImmediateDeviceListByDeviceGroupIDParamsWithContext(ctx context.Context) *GetImmediateDeviceListByDeviceGroupIDParams {
	return &GetImmediateDeviceListByDeviceGroupIDParams{
		Context: ctx,
	}
}

// NewGetImmediateDeviceListByDeviceGroupIDParamsWithHTTPClient creates a new GetImmediateDeviceListByDeviceGroupIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetImmediateDeviceListByDeviceGroupIDParamsWithHTTPClient(client *http.Client) *GetImmediateDeviceListByDeviceGroupIDParams {
	return &GetImmediateDeviceListByDeviceGroupIDParams{
		HTTPClient: client,
	}
}

/* GetImmediateDeviceListByDeviceGroupIDParams contains all the parameters to send to the API endpoint
   for the get immediate device list by device group Id operation.

   Typically these are written to a http.Request.
*/
type GetImmediateDeviceListByDeviceGroupIDParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-4-gd30bcfd-dirty"
	UserAgent *string

	// Fields.
	Fields *string

	// Filter.
	Filter *string

	// ID.
	//
	// Format: int32
	ID int32

	// Offset.
	//
	// Format: int32
	Offset *int32

	// Size.
	//
	// Format: int32
	// Default: 50
	Size *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get immediate device list by device group Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetImmediateDeviceListByDeviceGroupIDParams) WithDefaults() *GetImmediateDeviceListByDeviceGroupIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get immediate device list by device group Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetImmediateDeviceListByDeviceGroupIDParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-4-gd30bcfd-dirty")

		offsetDefault = int32(0)

		sizeDefault = int32(50)
	)

	val := GetImmediateDeviceListByDeviceGroupIDParams{
		UserAgent: &userAgentDefault,
		Offset:    &offsetDefault,
		Size:      &sizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get immediate device list by device group Id params
func (o *GetImmediateDeviceListByDeviceGroupIDParams) WithTimeout(timeout time.Duration) *GetImmediateDeviceListByDeviceGroupIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get immediate device list by device group Id params
func (o *GetImmediateDeviceListByDeviceGroupIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get immediate device list by device group Id params
func (o *GetImmediateDeviceListByDeviceGroupIDParams) WithContext(ctx context.Context) *GetImmediateDeviceListByDeviceGroupIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get immediate device list by device group Id params
func (o *GetImmediateDeviceListByDeviceGroupIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get immediate device list by device group Id params
func (o *GetImmediateDeviceListByDeviceGroupIDParams) WithHTTPClient(client *http.Client) *GetImmediateDeviceListByDeviceGroupIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get immediate device list by device group Id params
func (o *GetImmediateDeviceListByDeviceGroupIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get immediate device list by device group Id params
func (o *GetImmediateDeviceListByDeviceGroupIDParams) WithUserAgent(userAgent *string) *GetImmediateDeviceListByDeviceGroupIDParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get immediate device list by device group Id params
func (o *GetImmediateDeviceListByDeviceGroupIDParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithFields adds the fields to the get immediate device list by device group Id params
func (o *GetImmediateDeviceListByDeviceGroupIDParams) WithFields(fields *string) *GetImmediateDeviceListByDeviceGroupIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get immediate device list by device group Id params
func (o *GetImmediateDeviceListByDeviceGroupIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get immediate device list by device group Id params
func (o *GetImmediateDeviceListByDeviceGroupIDParams) WithFilter(filter *string) *GetImmediateDeviceListByDeviceGroupIDParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get immediate device list by device group Id params
func (o *GetImmediateDeviceListByDeviceGroupIDParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithID adds the id to the get immediate device list by device group Id params
func (o *GetImmediateDeviceListByDeviceGroupIDParams) WithID(id int32) *GetImmediateDeviceListByDeviceGroupIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get immediate device list by device group Id params
func (o *GetImmediateDeviceListByDeviceGroupIDParams) SetID(id int32) {
	o.ID = id
}

// WithOffset adds the offset to the get immediate device list by device group Id params
func (o *GetImmediateDeviceListByDeviceGroupIDParams) WithOffset(offset *int32) *GetImmediateDeviceListByDeviceGroupIDParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get immediate device list by device group Id params
func (o *GetImmediateDeviceListByDeviceGroupIDParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSize adds the size to the get immediate device list by device group Id params
func (o *GetImmediateDeviceListByDeviceGroupIDParams) WithSize(size *int32) *GetImmediateDeviceListByDeviceGroupIDParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get immediate device list by device group Id params
func (o *GetImmediateDeviceListByDeviceGroupIDParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetImmediateDeviceListByDeviceGroupIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Fields != nil {

		// query param fields
		var qrFields string

		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {

			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}
	}

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Size != nil {

		// query param size
		var qrSize int32

		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt32(qrSize)
		if qSize != "" {

			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
