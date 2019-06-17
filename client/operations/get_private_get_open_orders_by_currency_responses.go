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

// GetPrivateGetOpenOrdersByCurrencyReader is a Reader for the GetPrivateGetOpenOrdersByCurrency structure.
type GetPrivateGetOpenOrdersByCurrencyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateGetOpenOrdersByCurrencyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateGetOpenOrdersByCurrencyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateGetOpenOrdersByCurrencyOK creates a GetPrivateGetOpenOrdersByCurrencyOK with default headers values
func NewGetPrivateGetOpenOrdersByCurrencyOK() *GetPrivateGetOpenOrdersByCurrencyOK {
	return &GetPrivateGetOpenOrdersByCurrencyOK{}
}

/*GetPrivateGetOpenOrdersByCurrencyOK handles this case with default header values.

GetPrivateGetOpenOrdersByCurrencyOK get private get open orders by currency o k
*/
type GetPrivateGetOpenOrdersByCurrencyOK struct {
	Payload *models.PrivateGetOpenOrdersResponse
}

func (o *GetPrivateGetOpenOrdersByCurrencyOK) Error() string {
	return fmt.Sprintf("[GET /private/get_open_orders_by_currency][%d] getPrivateGetOpenOrdersByCurrencyOK  %+v", 200, o.Payload)
}

func (o *GetPrivateGetOpenOrdersByCurrencyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateGetOpenOrdersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
