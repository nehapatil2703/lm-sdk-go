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

// NewGetDeviceGroupPropertyListParams creates a new GetDeviceGroupPropertyListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeviceGroupPropertyListParams() *GetDeviceGroupPropertyListParams {
	return &GetDeviceGroupPropertyListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceGroupPropertyListParamsWithTimeout creates a new GetDeviceGroupPropertyListParams object
// with the ability to set a timeout on a request.
func NewGetDeviceGroupPropertyListParamsWithTimeout(timeout time.Duration) *GetDeviceGroupPropertyListParams {
	return &GetDeviceGroupPropertyListParams{
		timeout: timeout,
	}
}

// NewGetDeviceGroupPropertyListParamsWithContext creates a new GetDeviceGroupPropertyListParams object
// with the ability to set a context for a request.
func NewGetDeviceGroupPropertyListParamsWithContext(ctx context.Context) *GetDeviceGroupPropertyListParams {
	return &GetDeviceGroupPropertyListParams{
		Context: ctx,
	}
}

// NewGetDeviceGroupPropertyListParamsWithHTTPClient creates a new GetDeviceGroupPropertyListParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeviceGroupPropertyListParamsWithHTTPClient(client *http.Client) *GetDeviceGroupPropertyListParams {
	return &GetDeviceGroupPropertyListParams{
		HTTPClient: client,
	}
}

/* GetDeviceGroupPropertyListParams contains all the parameters to send to the API endpoint
   for the get device group property list operation.

   Typically these are written to a http.Request.
*/
type GetDeviceGroupPropertyListParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-4-gd30bcfd-dirty"
	UserAgent *string

	// Fields.
	Fields *string

	// Filter.
	Filter *string

	/* Gid.

	   group ID

	   Format: int32
	*/
	Gid int32

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

// WithDefaults hydrates default values in the get device group property list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceGroupPropertyListParams) WithDefaults() *GetDeviceGroupPropertyListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get device group property list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeviceGroupPropertyListParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-v2.0.0-argus5-4-gd30bcfd-dirty")

		offsetDefault = int32(0)

		sizeDefault = int32(50)
	)

	val := GetDeviceGroupPropertyListParams{
		UserAgent: &userAgentDefault,
		Offset:    &offsetDefault,
		Size:      &sizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get device group property list params
func (o *GetDeviceGroupPropertyListParams) WithTimeout(timeout time.Duration) *GetDeviceGroupPropertyListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device group property list params
func (o *GetDeviceGroupPropertyListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device group property list params
func (o *GetDeviceGroupPropertyListParams) WithContext(ctx context.Context) *GetDeviceGroupPropertyListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device group property list params
func (o *GetDeviceGroupPropertyListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device group property list params
func (o *GetDeviceGroupPropertyListParams) WithHTTPClient(client *http.Client) *GetDeviceGroupPropertyListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device group property list params
func (o *GetDeviceGroupPropertyListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get device group property list params
func (o *GetDeviceGroupPropertyListParams) WithUserAgent(userAgent *string) *GetDeviceGroupPropertyListParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get device group property list params
func (o *GetDeviceGroupPropertyListParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithFields adds the fields to the get device group property list params
func (o *GetDeviceGroupPropertyListParams) WithFields(fields *string) *GetDeviceGroupPropertyListParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get device group property list params
func (o *GetDeviceGroupPropertyListParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get device group property list params
func (o *GetDeviceGroupPropertyListParams) WithFilter(filter *string) *GetDeviceGroupPropertyListParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get device group property list params
func (o *GetDeviceGroupPropertyListParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithGid adds the gid to the get device group property list params
func (o *GetDeviceGroupPropertyListParams) WithGid(gid int32) *GetDeviceGroupPropertyListParams {
	o.SetGid(gid)
	return o
}

// SetGid adds the gid to the get device group property list params
func (o *GetDeviceGroupPropertyListParams) SetGid(gid int32) {
	o.Gid = gid
}

// WithOffset adds the offset to the get device group property list params
func (o *GetDeviceGroupPropertyListParams) WithOffset(offset *int32) *GetDeviceGroupPropertyListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get device group property list params
func (o *GetDeviceGroupPropertyListParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSize adds the size to the get device group property list params
func (o *GetDeviceGroupPropertyListParams) WithSize(size *int32) *GetDeviceGroupPropertyListParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get device group property list params
func (o *GetDeviceGroupPropertyListParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceGroupPropertyListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param gid
	if err := r.SetPathParam("gid", swag.FormatInt32(o.Gid)); err != nil {
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
