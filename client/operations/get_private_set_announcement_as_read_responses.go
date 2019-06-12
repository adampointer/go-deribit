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

// GetPrivateSetAnnouncementAsReadReader is a Reader for the GetPrivateSetAnnouncementAsRead structure.
type GetPrivateSetAnnouncementAsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateSetAnnouncementAsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateSetAnnouncementAsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateSetAnnouncementAsReadOK creates a GetPrivateSetAnnouncementAsReadOK with default headers values
func NewGetPrivateSetAnnouncementAsReadOK() *GetPrivateSetAnnouncementAsReadOK {
	return &GetPrivateSetAnnouncementAsReadOK{}
}

/*GetPrivateSetAnnouncementAsReadOK handles this case with default header values.

foo
*/
type GetPrivateSetAnnouncementAsReadOK struct {
	Payload *models.OkResponse
}

func (o *GetPrivateSetAnnouncementAsReadOK) Error() string {
	return fmt.Sprintf("[GET /private/set_announcement_as_read][%d] getPrivateSetAnnouncementAsReadOK  %+v", 200, o.Payload)
}

func (o *GetPrivateSetAnnouncementAsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
