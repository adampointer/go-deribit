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

// GetPrivateGetUserTradesByInstrumentAndTimeReader is a Reader for the GetPrivateGetUserTradesByInstrumentAndTime structure.
type GetPrivateGetUserTradesByInstrumentAndTimeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateGetUserTradesByInstrumentAndTimeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateGetUserTradesByInstrumentAndTimeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateGetUserTradesByInstrumentAndTimeOK creates a GetPrivateGetUserTradesByInstrumentAndTimeOK with default headers values
func NewGetPrivateGetUserTradesByInstrumentAndTimeOK() *GetPrivateGetUserTradesByInstrumentAndTimeOK {
	return &GetPrivateGetUserTradesByInstrumentAndTimeOK{}
}

/*GetPrivateGetUserTradesByInstrumentAndTimeOK handles this case with default header values.

GetPrivateGetUserTradesByInstrumentAndTimeOK get private get user trades by instrument and time o k
*/
type GetPrivateGetUserTradesByInstrumentAndTimeOK struct {
	Payload *models.PrivateGetUserTradesHistoryResponse
}

func (o *GetPrivateGetUserTradesByInstrumentAndTimeOK) Error() string {
	return fmt.Sprintf("[GET /private/get_user_trades_by_instrument_and_time][%d] getPrivateGetUserTradesByInstrumentAndTimeOK  %+v", 200, o.Payload)
}

func (o *GetPrivateGetUserTradesByInstrumentAndTimeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateGetUserTradesHistoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
