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

// GetPublicSubscribeReader is a Reader for the GetPublicSubscribe structure.
type GetPublicSubscribeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicSubscribeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPublicSubscribeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetPublicSubscribeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicSubscribeOK creates a GetPublicSubscribeOK with default headers values
func NewGetPublicSubscribeOK() *GetPublicSubscribeOK {
	return &GetPublicSubscribeOK{}
}

/*GetPublicSubscribeOK handles this case with default header values.

ok response
*/
type GetPublicSubscribeOK struct {
	Payload *models.PrivateSubscribeResponse
}

func (o *GetPublicSubscribeOK) Error() string {
	return fmt.Sprintf("[GET /public/subscribe][%d] getPublicSubscribeOK  %+v", 200, o.Payload)
}

func (o *GetPublicSubscribeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateSubscribeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPublicSubscribeUnauthorized creates a GetPublicSubscribeUnauthorized with default headers values
func NewGetPublicSubscribeUnauthorized() *GetPublicSubscribeUnauthorized {
	return &GetPublicSubscribeUnauthorized{}
}

/*GetPublicSubscribeUnauthorized handles this case with default header values.

not authorised
*/
type GetPublicSubscribeUnauthorized struct {
	Payload *models.ErrorMessage
}

func (o *GetPublicSubscribeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /public/subscribe][%d] getPublicSubscribeUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPublicSubscribeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
