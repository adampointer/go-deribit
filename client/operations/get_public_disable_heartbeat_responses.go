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

// GetPublicDisableHeartbeatReader is a Reader for the GetPublicDisableHeartbeat structure.
type GetPublicDisableHeartbeatReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicDisableHeartbeatReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPublicDisableHeartbeatOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetPublicDisableHeartbeatBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicDisableHeartbeatOK creates a GetPublicDisableHeartbeatOK with default headers values
func NewGetPublicDisableHeartbeatOK() *GetPublicDisableHeartbeatOK {
	return &GetPublicDisableHeartbeatOK{}
}

/*GetPublicDisableHeartbeatOK handles this case with default header values.

foo
*/
type GetPublicDisableHeartbeatOK struct {
	Payload *models.OkResponse
}

func (o *GetPublicDisableHeartbeatOK) Error() string {
	return fmt.Sprintf("[GET /public/disable_heartbeat][%d] getPublicDisableHeartbeatOK  %+v", 200, o.Payload)
}

func (o *GetPublicDisableHeartbeatOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPublicDisableHeartbeatBadRequest creates a GetPublicDisableHeartbeatBadRequest with default headers values
func NewGetPublicDisableHeartbeatBadRequest() *GetPublicDisableHeartbeatBadRequest {
	return &GetPublicDisableHeartbeatBadRequest{}
}

/*GetPublicDisableHeartbeatBadRequest handles this case with default header values.

result when used via rest/HTTP
*/
type GetPublicDisableHeartbeatBadRequest struct {
	Payload *models.ErrorMessage
}

func (o *GetPublicDisableHeartbeatBadRequest) Error() string {
	return fmt.Sprintf("[GET /public/disable_heartbeat][%d] getPublicDisableHeartbeatBadRequest  %+v", 400, o.Payload)
}

func (o *GetPublicDisableHeartbeatBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
