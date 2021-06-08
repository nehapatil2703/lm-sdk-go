// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/logicmonitor/lm-sdk-go/models"
)

// GetAuditLogByIDReader is a Reader for the GetAuditLogByID structure.
type GetAuditLogByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuditLogByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAuditLogByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetAuditLogByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAuditLogByIDOK creates a GetAuditLogByIDOK with default headers values
func NewGetAuditLogByIDOK() *GetAuditLogByIDOK {
	return &GetAuditLogByIDOK{}
}

/*GetAuditLogByIDOK handles this case with default header values.

successful operation
*/
type GetAuditLogByIDOK struct {
	Payload *models.AuditLog
}

func (o *GetAuditLogByIDOK) Error() string {
	return fmt.Sprintf("[GET /setting/accesslogs/{id}][%d] getAuditLogByIdOK  %+v", 200, o.Payload)
}

func (o *GetAuditLogByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.AuditLog)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditLogByIDDefault creates a GetAuditLogByIDDefault with default headers values
func NewGetAuditLogByIDDefault(code int) *GetAuditLogByIDDefault {
	return &GetAuditLogByIDDefault{
		_statusCode: code,
	}
}

/*GetAuditLogByIDDefault handles this case with default header values.

Error
*/
type GetAuditLogByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get audit log by Id default response
func (o *GetAuditLogByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetAuditLogByIDDefault) Error() string {
	return fmt.Sprintf("[GET /setting/accesslogs/{id}][%d] getAuditLogById default  %+v", o._statusCode, o.Payload)
}

func (o *GetAuditLogByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
