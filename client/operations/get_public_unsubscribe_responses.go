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

// GetPublicUnsubscribeReader is a Reader for the GetPublicUnsubscribe structure.
type GetPublicUnsubscribeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicUnsubscribeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPublicUnsubscribeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetPublicUnsubscribeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicUnsubscribeOK creates a GetPublicUnsubscribeOK with default headers values
func NewGetPublicUnsubscribeOK() *GetPublicUnsubscribeOK {
	return &GetPublicUnsubscribeOK{}
}

/*GetPublicUnsubscribeOK handles this case with default header values.

ok response
*/
type GetPublicUnsubscribeOK struct {
	Payload *models.PrivateSubscribeResponse
}

func (o *GetPublicUnsubscribeOK) Error() string {
	return fmt.Sprintf("[GET /public/unsubscribe][%d] getPublicUnsubscribeOK  %+v", 200, o.Payload)
}

func (o *GetPublicUnsubscribeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateSubscribeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPublicUnsubscribeUnauthorized creates a GetPublicUnsubscribeUnauthorized with default headers values
func NewGetPublicUnsubscribeUnauthorized() *GetPublicUnsubscribeUnauthorized {
	return &GetPublicUnsubscribeUnauthorized{}
}

/*GetPublicUnsubscribeUnauthorized handles this case with default header values.

not authorised
*/
type GetPublicUnsubscribeUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *GetPublicUnsubscribeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /public/unsubscribe][%d] getPublicUnsubscribeUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPublicUnsubscribeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
