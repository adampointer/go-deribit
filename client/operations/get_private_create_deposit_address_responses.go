// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/adampointer/go-deribit/v3/models"
)

// GetPrivateCreateDepositAddressReader is a Reader for the GetPrivateCreateDepositAddress structure.
type GetPrivateCreateDepositAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateCreateDepositAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateCreateDepositAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateCreateDepositAddressOK creates a GetPrivateCreateDepositAddressOK with default headers values
func NewGetPrivateCreateDepositAddressOK() *GetPrivateCreateDepositAddressOK {
	return &GetPrivateCreateDepositAddressOK{}
}

/*GetPrivateCreateDepositAddressOK handles this case with default header values.

GetPrivateCreateDepositAddressOK get private create deposit address o k
*/
type GetPrivateCreateDepositAddressOK struct {
	Payload *models.PrivateDepositAddressResponse
}

func (o *GetPrivateCreateDepositAddressOK) Error() string {
	return fmt.Sprintf("[GET /private/create_deposit_address][%d] getPrivateCreateDepositAddressOK  %+v", 200, o.Payload)
}

func (o *GetPrivateCreateDepositAddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateDepositAddressResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
