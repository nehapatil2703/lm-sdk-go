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

// PatchCollectorGroupByIDReader is a Reader for the PatchCollectorGroupByID structure.
type PatchCollectorGroupByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchCollectorGroupByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchCollectorGroupByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPatchCollectorGroupByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchCollectorGroupByIDOK creates a PatchCollectorGroupByIDOK with default headers values
func NewPatchCollectorGroupByIDOK() *PatchCollectorGroupByIDOK {
	return &PatchCollectorGroupByIDOK{}
}

/*PatchCollectorGroupByIDOK handles this case with default header values.

successful operation
*/
type PatchCollectorGroupByIDOK struct {
	Payload *models.CollectorGroup
}

func (o *PatchCollectorGroupByIDOK) Error() string {
	return fmt.Sprintf("[PATCH /setting/collector/groups/{id}][%d] patchCollectorGroupByIdOK  %+v", 200, o.Payload)
}

func (o *PatchCollectorGroupByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.CollectorGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCollectorGroupByIDDefault creates a PatchCollectorGroupByIDDefault with default headers values
func NewPatchCollectorGroupByIDDefault(code int) *PatchCollectorGroupByIDDefault {
	return &PatchCollectorGroupByIDDefault{
		_statusCode: code,
	}
}

/*PatchCollectorGroupByIDDefault handles this case with default header values.

Error
*/
type PatchCollectorGroupByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch collector group by Id default response
func (o *PatchCollectorGroupByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchCollectorGroupByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /setting/collector/groups/{id}][%d] patchCollectorGroupById default  %+v", o._statusCode, o.Payload)
}

func (o *PatchCollectorGroupByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
