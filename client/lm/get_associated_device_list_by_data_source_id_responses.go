// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/logicmonitor/lm-sdk-go/models"
)

// GetAssociatedDeviceListByDataSourceIDReader is a Reader for the GetAssociatedDeviceListByDataSourceID structure.
type GetAssociatedDeviceListByDataSourceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAssociatedDeviceListByDataSourceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAssociatedDeviceListByDataSourceIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 429:
		result := NewGetAssociatedDeviceListByDataSourceIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAssociatedDeviceListByDataSourceIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAssociatedDeviceListByDataSourceIDOK creates a GetAssociatedDeviceListByDataSourceIDOK with default headers values
func NewGetAssociatedDeviceListByDataSourceIDOK() *GetAssociatedDeviceListByDataSourceIDOK {
	return &GetAssociatedDeviceListByDataSourceIDOK{}
}

/* GetAssociatedDeviceListByDataSourceIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAssociatedDeviceListByDataSourceIDOK struct {
	Payload *models.DeviceDataSourceAssociatedPaginationResponse
}

func (o *GetAssociatedDeviceListByDataSourceIDOK) Error() string {
	return fmt.Sprintf("[GET /setting/datasources/{id}/devices][%d] getAssociatedDeviceListByDataSourceIdOK  %+v", 200, o.Payload)
}
func (o *GetAssociatedDeviceListByDataSourceIDOK) GetPayload() *models.DeviceDataSourceAssociatedPaginationResponse {
	return o.Payload
}

func (o *GetAssociatedDeviceListByDataSourceIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceDataSourceAssociatedPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAssociatedDeviceListByDataSourceIDTooManyRequests creates a GetAssociatedDeviceListByDataSourceIDTooManyRequests with default headers values
func NewGetAssociatedDeviceListByDataSourceIDTooManyRequests() *GetAssociatedDeviceListByDataSourceIDTooManyRequests {
	return &GetAssociatedDeviceListByDataSourceIDTooManyRequests{}
}

/* GetAssociatedDeviceListByDataSourceIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetAssociatedDeviceListByDataSourceIDTooManyRequests struct {

	/* Request limit per X-Rate-Limit-Window
	 */
	XRateLimitLimit int64

	/* The number of requests left for the time window
	 */
	XRateLimitRemaining int64

	/* The rolling time window length with the unit of second
	 */
	XRateLimitWindow int64
}

func (o *GetAssociatedDeviceListByDataSourceIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /setting/datasources/{id}/devices][%d] getAssociatedDeviceListByDataSourceIdTooManyRequests ", 429)
}

func (o *GetAssociatedDeviceListByDataSourceIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-rate-limit-limit
	hdrXRateLimitLimit := response.GetHeader("x-rate-limit-limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("x-rate-limit-limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header x-rate-limit-remaining
	hdrXRateLimitRemaining := response.GetHeader("x-rate-limit-remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("x-rate-limit-remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header x-rate-limit-window
	hdrXRateLimitWindow := response.GetHeader("x-rate-limit-window")

	if hdrXRateLimitWindow != "" {
		valxRateLimitWindow, err := swag.ConvertInt64(hdrXRateLimitWindow)
		if err != nil {
			return errors.InvalidType("x-rate-limit-window", "header", "int64", hdrXRateLimitWindow)
		}
		o.XRateLimitWindow = valxRateLimitWindow
	}

	return nil
}

// NewGetAssociatedDeviceListByDataSourceIDDefault creates a GetAssociatedDeviceListByDataSourceIDDefault with default headers values
func NewGetAssociatedDeviceListByDataSourceIDDefault(code int) *GetAssociatedDeviceListByDataSourceIDDefault {
	return &GetAssociatedDeviceListByDataSourceIDDefault{
		_statusCode: code,
	}
}

/* GetAssociatedDeviceListByDataSourceIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetAssociatedDeviceListByDataSourceIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get associated device list by data source Id default response
func (o *GetAssociatedDeviceListByDataSourceIDDefault) Code() int {
	return o._statusCode
}

func (o *GetAssociatedDeviceListByDataSourceIDDefault) Error() string {
	return fmt.Sprintf("[GET /setting/datasources/{id}/devices][%d] getAssociatedDeviceListByDataSourceId default  %+v", o._statusCode, o.Payload)
}
func (o *GetAssociatedDeviceListByDataSourceIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAssociatedDeviceListByDataSourceIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
