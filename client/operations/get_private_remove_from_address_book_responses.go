// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/DubininaTaisiia/go-deribit/models"
)

// GetPrivateRemoveFromAddressBookReader is a Reader for the GetPrivateRemoveFromAddressBook structure.
type GetPrivateRemoveFromAddressBookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateRemoveFromAddressBookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateRemoveFromAddressBookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateRemoveFromAddressBookOK creates a GetPrivateRemoveFromAddressBookOK with default headers values
func NewGetPrivateRemoveFromAddressBookOK() *GetPrivateRemoveFromAddressBookOK {
	return &GetPrivateRemoveFromAddressBookOK{}
}

/*GetPrivateRemoveFromAddressBookOK handles this case with default header values.

foo
*/
type GetPrivateRemoveFromAddressBookOK struct {
	Payload *models.PrivateRemoveFromAddressBookResponse
}

func (o *GetPrivateRemoveFromAddressBookOK) Error() string {
	return fmt.Sprintf("[GET /private/remove_from_address_book][%d] getPrivateRemoveFromAddressBookOK  %+v", 200, o.Payload)
}

func (o *GetPrivateRemoveFromAddressBookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateRemoveFromAddressBookResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
