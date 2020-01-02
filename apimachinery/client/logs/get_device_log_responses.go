// Code generated by go-swagger; DO NOT EDIT.

package logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cloustone/pandas/models"
)

// GetDeviceLogReader is a Reader for the GetDeviceLog structure.
type GetDeviceLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceLogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetDeviceLogNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDeviceLogOK creates a GetDeviceLogOK with default headers values
func NewGetDeviceLogOK() *GetDeviceLogOK {
	return &GetDeviceLogOK{}
}

/*GetDeviceLogOK handles this case with default header values.

successful operation
*/
type GetDeviceLogOK struct {
	Payload []models.DeviceLog
}

func (o *GetDeviceLogOK) Error() string {
	return fmt.Sprintf("[GET /logs/devices/{deviceId}][%d] getDeviceLogOK  %+v", 200, o.Payload)
}

func (o *GetDeviceLogOK) GetPayload() []models.DeviceLog {
	return o.Payload
}

func (o *GetDeviceLogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceLogNotFound creates a GetDeviceLogNotFound with default headers values
func NewGetDeviceLogNotFound() *GetDeviceLogNotFound {
	return &GetDeviceLogNotFound{}
}

/*GetDeviceLogNotFound handles this case with default header values.

device not found.
*/
type GetDeviceLogNotFound struct {
}

func (o *GetDeviceLogNotFound) Error() string {
	return fmt.Sprintf("[GET /logs/devices/{deviceId}][%d] getDeviceLogNotFound ", 404)
}

func (o *GetDeviceLogNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
