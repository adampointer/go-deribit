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

// GetPrivateGetSubaccountsReader is a Reader for the GetPrivateGetSubaccounts structure.
type GetPrivateGetSubaccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateGetSubaccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateGetSubaccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetPrivateGetSubaccountsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateGetSubaccountsOK creates a GetPrivateGetSubaccountsOK with default headers values
func NewGetPrivateGetSubaccountsOK() *GetPrivateGetSubaccountsOK {
	return &GetPrivateGetSubaccountsOK{}
}

/*GetPrivateGetSubaccountsOK handles this case with default header values.

ok response
*/
type GetPrivateGetSubaccountsOK struct {
	Payload *models.PrivateGetSubaccountsResponse
}

func (o *GetPrivateGetSubaccountsOK) Error() string {
	return fmt.Sprintf("[GET /private/get_subaccounts][%d] getPrivateGetSubaccountsOK  %+v", 200, o.Payload)
}

func (o *GetPrivateGetSubaccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateGetSubaccountsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrivateGetSubaccountsUnauthorized creates a GetPrivateGetSubaccountsUnauthorized with default headers values
func NewGetPrivateGetSubaccountsUnauthorized() *GetPrivateGetSubaccountsUnauthorized {
	return &GetPrivateGetSubaccountsUnauthorized{}
}

/*GetPrivateGetSubaccountsUnauthorized handles this case with default header values.

not authorised
*/
type GetPrivateGetSubaccountsUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *GetPrivateGetSubaccountsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /private/get_subaccounts][%d] getPrivateGetSubaccountsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPrivateGetSubaccountsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
