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

// UpdateAlertRuleByIDReader is a Reader for the UpdateAlertRuleByID structure.
type UpdateAlertRuleByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAlertRuleByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateAlertRuleByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateAlertRuleByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateAlertRuleByIDOK creates a UpdateAlertRuleByIDOK with default headers values
func NewUpdateAlertRuleByIDOK() *UpdateAlertRuleByIDOK {
	return &UpdateAlertRuleByIDOK{}
}

/*UpdateAlertRuleByIDOK handles this case with default header values.

successful operation
*/
type UpdateAlertRuleByIDOK struct {
	Payload *models.AlertRule
}

func (o *UpdateAlertRuleByIDOK) Error() string {
	return fmt.Sprintf("[PUT /setting/alert/rules/{id}][%d] updateAlertRuleByIdOK  %+v", 200, o.Payload)
}

func (o *UpdateAlertRuleByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.AlertRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAlertRuleByIDDefault creates a UpdateAlertRuleByIDDefault with default headers values
func NewUpdateAlertRuleByIDDefault(code int) *UpdateAlertRuleByIDDefault {
	return &UpdateAlertRuleByIDDefault{
		_statusCode: code,
	}
}

/*UpdateAlertRuleByIDDefault handles this case with default header values.

Error
*/
type UpdateAlertRuleByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update alert rule by Id default response
func (o *UpdateAlertRuleByIDDefault) Code() int {
	return o._statusCode
}

func (o *UpdateAlertRuleByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /setting/alert/rules/{id}][%d] updateAlertRuleById default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateAlertRuleByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
