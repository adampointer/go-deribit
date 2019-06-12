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

// GetPrivateGetUserTradesByOrderReader is a Reader for the GetPrivateGetUserTradesByOrder structure.
type GetPrivateGetUserTradesByOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateGetUserTradesByOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateGetUserTradesByOrderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateGetUserTradesByOrderOK creates a GetPrivateGetUserTradesByOrderOK with default headers values
func NewGetPrivateGetUserTradesByOrderOK() *GetPrivateGetUserTradesByOrderOK {
	return &GetPrivateGetUserTradesByOrderOK{}
}

/*GetPrivateGetUserTradesByOrderOK handles this case with default header values.

foo
*/
type GetPrivateGetUserTradesByOrderOK struct {
	Payload models.UserTradesByOrderResponse
}

func (o *GetPrivateGetUserTradesByOrderOK) Error() string {
	return fmt.Sprintf("[GET /private/get_user_trades_by_order][%d] getPrivateGetUserTradesByOrderOK  %+v", 200, o.Payload)
}

func (o *GetPrivateGetUserTradesByOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
