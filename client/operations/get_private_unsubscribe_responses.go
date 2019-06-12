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

// GetPrivateUnsubscribeReader is a Reader for the GetPrivateUnsubscribe structure.
type GetPrivateUnsubscribeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateUnsubscribeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateUnsubscribeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetPrivateUnsubscribeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateUnsubscribeOK creates a GetPrivateUnsubscribeOK with default headers values
func NewGetPrivateUnsubscribeOK() *GetPrivateUnsubscribeOK {
	return &GetPrivateUnsubscribeOK{}
}

/*GetPrivateUnsubscribeOK handles this case with default header values.

foo
*/
type GetPrivateUnsubscribeOK struct {
	Payload *models.PrivateSubscribeResponse
}

func (o *GetPrivateUnsubscribeOK) Error() string {
	return fmt.Sprintf("[GET /private/unsubscribe][%d] getPrivateUnsubscribeOK  %+v", 200, o.Payload)
}

func (o *GetPrivateUnsubscribeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateSubscribeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPrivateUnsubscribeUnauthorized creates a GetPrivateUnsubscribeUnauthorized with default headers values
func NewGetPrivateUnsubscribeUnauthorized() *GetPrivateUnsubscribeUnauthorized {
	return &GetPrivateUnsubscribeUnauthorized{}
}

/*GetPrivateUnsubscribeUnauthorized handles this case with default header values.

not authorised
*/
type GetPrivateUnsubscribeUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *GetPrivateUnsubscribeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /private/unsubscribe][%d] getPrivateUnsubscribeUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPrivateUnsubscribeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
