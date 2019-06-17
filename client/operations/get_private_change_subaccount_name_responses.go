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

// GetPrivateChangeSubaccountNameReader is a Reader for the GetPrivateChangeSubaccountName structure.
type GetPrivateChangeSubaccountNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateChangeSubaccountNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateChangeSubaccountNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateChangeSubaccountNameOK creates a GetPrivateChangeSubaccountNameOK with default headers values
func NewGetPrivateChangeSubaccountNameOK() *GetPrivateChangeSubaccountNameOK {
	return &GetPrivateChangeSubaccountNameOK{}
}

/*GetPrivateChangeSubaccountNameOK handles this case with default header values.

ok response
*/
type GetPrivateChangeSubaccountNameOK struct {
	Payload *models.OkResponse
}

func (o *GetPrivateChangeSubaccountNameOK) Error() string {
	return fmt.Sprintf("[GET /private/change_subaccount_name][%d] getPrivateChangeSubaccountNameOK  %+v", 200, o.Payload)
}

func (o *GetPrivateChangeSubaccountNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
