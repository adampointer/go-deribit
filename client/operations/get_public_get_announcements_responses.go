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

// GetPublicGetAnnouncementsReader is a Reader for the GetPublicGetAnnouncements structure.
type GetPublicGetAnnouncementsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicGetAnnouncementsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPublicGetAnnouncementsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicGetAnnouncementsOK creates a GetPublicGetAnnouncementsOK with default headers values
func NewGetPublicGetAnnouncementsOK() *GetPublicGetAnnouncementsOK {
	return &GetPublicGetAnnouncementsOK{}
}

/*GetPublicGetAnnouncementsOK handles this case with default header values.

foo
*/
type GetPublicGetAnnouncementsOK struct {
	Payload *models.GetAnnouncementsResponse
}

func (o *GetPublicGetAnnouncementsOK) Error() string {
	return fmt.Sprintf("[GET /public/get_announcements][%d] getPublicGetAnnouncementsOK  %+v", 200, o.Payload)
}

func (o *GetPublicGetAnnouncementsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetAnnouncementsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
