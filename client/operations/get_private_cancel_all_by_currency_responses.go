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

// GetPrivateCancelAllByCurrencyReader is a Reader for the GetPrivateCancelAllByCurrency structure.
type GetPrivateCancelAllByCurrencyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateCancelAllByCurrencyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateCancelAllByCurrencyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateCancelAllByCurrencyOK creates a GetPrivateCancelAllByCurrencyOK with default headers values
func NewGetPrivateCancelAllByCurrencyOK() *GetPrivateCancelAllByCurrencyOK {
	return &GetPrivateCancelAllByCurrencyOK{}
}

/*GetPrivateCancelAllByCurrencyOK handles this case with default header values.

GetPrivateCancelAllByCurrencyOK get private cancel all by currency o k
*/
type GetPrivateCancelAllByCurrencyOK struct {
	Payload *models.PrivateCancelAllResponse
}

func (o *GetPrivateCancelAllByCurrencyOK) Error() string {
	return fmt.Sprintf("[GET /private/cancel_all_by_currency][%d] getPrivateCancelAllByCurrencyOK  %+v", 200, o.Payload)
}

func (o *GetPrivateCancelAllByCurrencyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateCancelAllResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
