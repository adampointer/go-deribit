// Code generated by go-swagger; DO NOT EDIT.

package private

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/adampointer/go-deribit/v3/models"
)

// GetPrivateClosePositionReader is a Reader for the GetPrivateClosePosition structure.
type GetPrivateClosePositionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateClosePositionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPrivateClosePositionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateClosePositionOK creates a GetPrivateClosePositionOK with default headers values
func NewGetPrivateClosePositionOK() *GetPrivateClosePositionOK {
	return &GetPrivateClosePositionOK{}
}

/*GetPrivateClosePositionOK handles this case with default header values.

ok response
*/
type GetPrivateClosePositionOK struct {
	Payload *models.PrivateBuyAndSellResponse
}

func (o *GetPrivateClosePositionOK) Error() string {
	return fmt.Sprintf("[GET /private/close_position][%d] getPrivateClosePositionOK  %+v", 200, o.Payload)
}

func (o *GetPrivateClosePositionOK) GetPayload() *models.PrivateBuyAndSellResponse {
	return o.Payload
}

func (o *GetPrivateClosePositionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateBuyAndSellResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}