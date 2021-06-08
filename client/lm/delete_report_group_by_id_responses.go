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

// DeleteReportGroupByIDReader is a Reader for the DeleteReportGroupByID structure.
type DeleteReportGroupByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteReportGroupByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteReportGroupByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteReportGroupByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteReportGroupByIDOK creates a DeleteReportGroupByIDOK with default headers values
func NewDeleteReportGroupByIDOK() *DeleteReportGroupByIDOK {
	return &DeleteReportGroupByIDOK{}
}

/*DeleteReportGroupByIDOK handles this case with default header values.

successful operation
*/
type DeleteReportGroupByIDOK struct {
	Payload interface{}
}

func (o *DeleteReportGroupByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /report/groups/{id}][%d] deleteReportGroupByIdOK  %+v", 200, o.Payload)
}

func (o *DeleteReportGroupByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteReportGroupByIDDefault creates a DeleteReportGroupByIDDefault with default headers values
func NewDeleteReportGroupByIDDefault(code int) *DeleteReportGroupByIDDefault {
	return &DeleteReportGroupByIDDefault{
		_statusCode: code,
	}
}

/*DeleteReportGroupByIDDefault handles this case with default header values.

Error
*/
type DeleteReportGroupByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete report group by Id default response
func (o *DeleteReportGroupByIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteReportGroupByIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /report/groups/{id}][%d] deleteReportGroupById default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteReportGroupByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
