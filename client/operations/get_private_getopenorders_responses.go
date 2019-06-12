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

// GetPrivateGetopenordersReader is a Reader for the GetPrivateGetopenorders structure.
type GetPrivateGetopenordersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateGetopenordersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateGetopenordersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateGetopenordersOK creates a GetPrivateGetopenordersOK with default headers values
func NewGetPrivateGetopenordersOK() *GetPrivateGetopenordersOK {
	return &GetPrivateGetopenordersOK{}
}

/*GetPrivateGetopenordersOK handles this case with default header values.

foo
*/
type GetPrivateGetopenordersOK struct {
	Payload *models.PrivateGetopenordersResponse
}

func (o *GetPrivateGetopenordersOK) Error() string {
	return fmt.Sprintf("[GET /private/getopenorders][%d] getPrivateGetopenordersOK  %+v", 200, o.Payload)
}

func (o *GetPrivateGetopenordersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateGetopenordersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
