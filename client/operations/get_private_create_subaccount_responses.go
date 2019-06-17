// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/adampointer/go-deribit/models"
)

// GetPrivateCreateSubaccountReader is a Reader for the GetPrivateCreateSubaccount structure.
type GetPrivateCreateSubaccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateCreateSubaccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateCreateSubaccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateCreateSubaccountOK creates a GetPrivateCreateSubaccountOK with default headers values
func NewGetPrivateCreateSubaccountOK() *GetPrivateCreateSubaccountOK {
	return &GetPrivateCreateSubaccountOK{}
}

/*GetPrivateCreateSubaccountOK handles this case with default header values.

ok response
*/
type GetPrivateCreateSubaccountOK struct {
	Payload *models.PrivateCreateSubaccountResponse
}

func (o *GetPrivateCreateSubaccountOK) Error() string {
	return fmt.Sprintf("[GET /private/create_subaccount][%d] getPrivateCreateSubaccountOK  %+v", 200, o.Payload)
}

func (o *GetPrivateCreateSubaccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateCreateSubaccountResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
