// Code generated by go-swagger; DO NOT EDIT.

package deployment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// GetDeploymentsReader is a Reader for the GetDeployments structure.
type GetDeploymentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDeploymentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetDeploymentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeploymentsOK creates a GetDeploymentsOK with default headers values
func NewGetDeploymentsOK() *GetDeploymentsOK {
	return &GetDeploymentsOK{}
}

/*GetDeploymentsOK handles this case with default header values.

successfully operation
*/
type GetDeploymentsOK struct {
	Payload []interface{}
}

func (o *GetDeploymentsOK) Error() string {
	return fmt.Sprintf("[GET /deployments][%d] getDeploymentsOK  %+v", 200, o.Payload)
}

func (o *GetDeploymentsOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetDeploymentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentsBadRequest creates a GetDeploymentsBadRequest with default headers values
func NewGetDeploymentsBadRequest() *GetDeploymentsBadRequest {
	return &GetDeploymentsBadRequest{}
}

/*GetDeploymentsBadRequest handles this case with default header values.

Internal server rooro
*/
type GetDeploymentsBadRequest struct {
}

func (o *GetDeploymentsBadRequest) Error() string {
	return fmt.Sprintf("[GET /deployments][%d] getDeploymentsBadRequest ", 400)
}

func (o *GetDeploymentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeploymentsDefault creates a GetDeploymentsDefault with default headers values
func NewGetDeploymentsDefault(code int) *GetDeploymentsDefault {
	return &GetDeploymentsDefault{
		_statusCode: code,
	}
}

/*GetDeploymentsDefault handles this case with default header values.

failed operation
*/
type GetDeploymentsDefault struct {
	_statusCode int

	Payload *GetDeploymentsDefaultBody
}

// Code gets the status code for the get deployments default response
func (o *GetDeploymentsDefault) Code() int {
	return o._statusCode
}

func (o *GetDeploymentsDefault) Error() string {
	return fmt.Sprintf("[GET /deployments][%d] getDeployments default  %+v", o._statusCode, o.Payload)
}

func (o *GetDeploymentsDefault) GetPayload() *GetDeploymentsDefaultBody {
	return o.Payload
}

func (o *GetDeploymentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetDeploymentsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetDeploymentsDefaultBody get deployments default body
swagger:model GetDeploymentsDefaultBody
*/
type GetDeploymentsDefaultBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get deployments default body
func (o *GetDeploymentsDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetDeploymentsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDeploymentsDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetDeploymentsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}