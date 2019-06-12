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

// GetPrivateGetSettlementHistoryByInstrumentReader is a Reader for the GetPrivateGetSettlementHistoryByInstrument structure.
type GetPrivateGetSettlementHistoryByInstrumentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateGetSettlementHistoryByInstrumentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateGetSettlementHistoryByInstrumentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateGetSettlementHistoryByInstrumentOK creates a GetPrivateGetSettlementHistoryByInstrumentOK with default headers values
func NewGetPrivateGetSettlementHistoryByInstrumentOK() *GetPrivateGetSettlementHistoryByInstrumentOK {
	return &GetPrivateGetSettlementHistoryByInstrumentOK{}
}

/*GetPrivateGetSettlementHistoryByInstrumentOK handles this case with default header values.

foo
*/
type GetPrivateGetSettlementHistoryByInstrumentOK struct {
	Payload *models.PrivateSettlementResponse
}

func (o *GetPrivateGetSettlementHistoryByInstrumentOK) Error() string {
	return fmt.Sprintf("[GET /private/get_settlement_history_by_instrument][%d] getPrivateGetSettlementHistoryByInstrumentOK  %+v", 200, o.Payload)
}

func (o *GetPrivateGetSettlementHistoryByInstrumentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateSettlementResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
