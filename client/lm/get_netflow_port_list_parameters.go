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

// NewGetNetflowPortListParams creates a new GetNetflowPortListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetflowPortListParams() *GetNetflowPortListParams {
	return &GetNetflowPortListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetflowPortListParamsWithTimeout creates a new GetNetflowPortListParams object
// with the ability to set a timeout on a request.
func NewGetNetflowPortListParamsWithTimeout(timeout time.Duration) *GetNetflowPortListParams {
	return &GetNetflowPortListParams{
		timeout: timeout,
	}
}

// NewGetNetflowPortListParamsWithContext creates a new GetNetflowPortListParams object
// with the ability to set a context for a request.
func NewGetNetflowPortListParamsWithContext(ctx context.Context) *GetNetflowPortListParams {
	return &GetNetflowPortListParams{
		Context: ctx,
	}
}

// NewGetNetflowPortListParamsWithHTTPClient creates a new GetNetflowPortListParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetflowPortListParamsWithHTTPClient(client *http.Client) *GetNetflowPortListParams {
	return &GetNetflowPortListParams{
		HTTPClient: client,
	}
}

/* GetNetflowPortListParams contains all the parameters to send to the API endpoint
   for the get netflow port list operation.

   Typically these are written to a http.Request.
*/
type GetNetflowPortListParams struct {

	// UserAgent.
	//
	// Default: "Logicmonitor/SDK: Argus Dist-95bb3f4-dirty"
	UserAgent *string

	// End.
	//
	// Format: int64
	End *int64

	// Fields.
	Fields *string

	// Filter.
	Filter *string

	// ID.
	//
	// Format: int32
	ID int32

	// IP.
	IP *string

	// NetflowFilter.
	NetflowFilter *string

	// Offset.
	//
	// Format: int32
	Offset *int32

	// Size.
	//
	// Format: int32
	// Default: 50
	Size *int32

	// Start.
	//
	// Format: int64
	Start *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get netflow port list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetflowPortListParams) WithDefaults() *GetNetflowPortListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get netflow port list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetflowPortListParams) SetDefaults() {
	var (
		userAgentDefault = string("Logicmonitor/SDK: Argus Dist-95bb3f4-dirty")

		offsetDefault = int32(0)

		sizeDefault = int32(50)
	)

	val := GetNetflowPortListParams{
		UserAgent: &userAgentDefault,
		Offset:    &offsetDefault,
		Size:      &sizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get netflow port list params
func (o *GetNetflowPortListParams) WithTimeout(timeout time.Duration) *GetNetflowPortListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get netflow port list params
func (o *GetNetflowPortListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get netflow port list params
func (o *GetNetflowPortListParams) WithContext(ctx context.Context) *GetNetflowPortListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get netflow port list params
func (o *GetNetflowPortListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get netflow port list params
func (o *GetNetflowPortListParams) WithHTTPClient(client *http.Client) *GetNetflowPortListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get netflow port list params
func (o *GetNetflowPortListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserAgent adds the userAgent to the get netflow port list params
func (o *GetNetflowPortListParams) WithUserAgent(userAgent *string) *GetNetflowPortListParams {
	o.SetUserAgent(userAgent)
	return o
}

// SetUserAgent adds the userAgent to the get netflow port list params
func (o *GetNetflowPortListParams) SetUserAgent(userAgent *string) {
	o.UserAgent = userAgent
}

// WithEnd adds the end to the get netflow port list params
func (o *GetNetflowPortListParams) WithEnd(end *int64) *GetNetflowPortListParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get netflow port list params
func (o *GetNetflowPortListParams) SetEnd(end *int64) {
	o.End = end
}

// WithFields adds the fields to the get netflow port list params
func (o *GetNetflowPortListParams) WithFields(fields *string) *GetNetflowPortListParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get netflow port list params
func (o *GetNetflowPortListParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get netflow port list params
func (o *GetNetflowPortListParams) WithFilter(filter *string) *GetNetflowPortListParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get netflow port list params
func (o *GetNetflowPortListParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithID adds the id to the get netflow port list params
func (o *GetNetflowPortListParams) WithID(id int32) *GetNetflowPortListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get netflow port list params
func (o *GetNetflowPortListParams) SetID(id int32) {
	o.ID = id
}

// WithIP adds the ip to the get netflow port list params
func (o *GetNetflowPortListParams) WithIP(ip *string) *GetNetflowPortListParams {
	o.SetIP(ip)
	return o
}

// SetIP adds the ip to the get netflow port list params
func (o *GetNetflowPortListParams) SetIP(ip *string) {
	o.IP = ip
}

// WithNetflowFilter adds the netflowFilter to the get netflow port list params
func (o *GetNetflowPortListParams) WithNetflowFilter(netflowFilter *string) *GetNetflowPortListParams {
	o.SetNetflowFilter(netflowFilter)
	return o
}

// SetNetflowFilter adds the netflowFilter to the get netflow port list params
func (o *GetNetflowPortListParams) SetNetflowFilter(netflowFilter *string) {
	o.NetflowFilter = netflowFilter
}

// WithOffset adds the offset to the get netflow port list params
func (o *GetNetflowPortListParams) WithOffset(offset *int32) *GetNetflowPortListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get netflow port list params
func (o *GetNetflowPortListParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSize adds the size to the get netflow port list params
func (o *GetNetflowPortListParams) WithSize(size *int32) *GetNetflowPortListParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get netflow port list params
func (o *GetNetflowPortListParams) SetSize(size *int32) {
	o.Size = size
}

// WithStart adds the start to the get netflow port list params
func (o *GetNetflowPortListParams) WithStart(start *int64) *GetNetflowPortListParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get netflow port list params
func (o *GetNetflowPortListParams) SetStart(start *int64) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetflowPortListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.End != nil {

		// query param end
		var qrEnd int64

		if o.End != nil {
			qrEnd = *o.End
		}
		qEnd := swag.FormatInt64(qrEnd)
		if qEnd != "" {

			if err := r.SetQueryParam("end", qEnd); err != nil {
				return err
			}
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

	if o.IP != nil {

		// query param ip
		var qrIP string

		if o.IP != nil {
			qrIP = *o.IP
		}
		qIP := qrIP
		if qIP != "" {

			if err := r.SetQueryParam("ip", qIP); err != nil {
				return err
			}
		}
	}

	if o.NetflowFilter != nil {

		// query param netflowFilter
		var qrNetflowFilter string

		if o.NetflowFilter != nil {
			qrNetflowFilter = *o.NetflowFilter
		}
		qNetflowFilter := qrNetflowFilter
		if qNetflowFilter != "" {

			if err := r.SetQueryParam("netflowFilter", qNetflowFilter); err != nil {
				return err
			}
		}
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

	if o.Start != nil {

		// query param start
		var qrStart int64

		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := swag.FormatInt64(qrStart)
		if qStart != "" {

			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
