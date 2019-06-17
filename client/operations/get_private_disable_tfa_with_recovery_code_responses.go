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

// GetPrivateDisableTfaWithRecoveryCodeReader is a Reader for the GetPrivateDisableTfaWithRecoveryCode structure.
type GetPrivateDisableTfaWithRecoveryCodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateDisableTfaWithRecoveryCodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateDisableTfaWithRecoveryCodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateDisableTfaWithRecoveryCodeOK creates a GetPrivateDisableTfaWithRecoveryCodeOK with default headers values
func NewGetPrivateDisableTfaWithRecoveryCodeOK() *GetPrivateDisableTfaWithRecoveryCodeOK {
	return &GetPrivateDisableTfaWithRecoveryCodeOK{}
}

/*GetPrivateDisableTfaWithRecoveryCodeOK handles this case with default header values.

ok response
*/
type GetPrivateDisableTfaWithRecoveryCodeOK struct {
	Payload *models.OkResponse
}

func (o *GetPrivateDisableTfaWithRecoveryCodeOK) Error() string {
	return fmt.Sprintf("[GET /private/disable_tfa_with_recovery_code][%d] getPrivateDisableTfaWithRecoveryCodeOK  %+v", 200, o.Payload)
}

func (o *GetPrivateDisableTfaWithRecoveryCodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
