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

// PatchAPITokenByAdminIDReader is a Reader for the PatchAPITokenByAdminID structure.
type PatchAPITokenByAdminIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPITokenByAdminIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchAPITokenByAdminIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPatchAPITokenByAdminIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchAPITokenByAdminIDOK creates a PatchAPITokenByAdminIDOK with default headers values
func NewPatchAPITokenByAdminIDOK() *PatchAPITokenByAdminIDOK {
	return &PatchAPITokenByAdminIDOK{}
}

/*PatchAPITokenByAdminIDOK handles this case with default header values.

successful operation
*/
type PatchAPITokenByAdminIDOK struct {
	Payload *models.APIToken
}

func (o *PatchAPITokenByAdminIDOK) Error() string {
	return fmt.Sprintf("[PATCH /setting/admins/{adminId}/apitokens/{apitokenId}][%d] patchApiTokenByAdminIdOK  %+v", 200, o.Payload)
}

func (o *PatchAPITokenByAdminIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.APIToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAPITokenByAdminIDDefault creates a PatchAPITokenByAdminIDDefault with default headers values
func NewPatchAPITokenByAdminIDDefault(code int) *PatchAPITokenByAdminIDDefault {
	return &PatchAPITokenByAdminIDDefault{
		_statusCode: code,
	}
}

/*PatchAPITokenByAdminIDDefault handles this case with default header values.

Error
*/
type PatchAPITokenByAdminIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch Api token by admin Id default response
func (o *PatchAPITokenByAdminIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchAPITokenByAdminIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /setting/admins/{adminId}/apitokens/{apitokenId}][%d] patchApiTokenByAdminId default  %+v", o._statusCode, o.Payload)
}

func (o *PatchAPITokenByAdminIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
