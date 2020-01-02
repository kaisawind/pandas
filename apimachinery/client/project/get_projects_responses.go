// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cloustone/pandas/models"
)

// GetProjectsReader is a Reader for the GetProjects structure.
type GetProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProjectsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProjectsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProjectsOK creates a GetProjectsOK with default headers values
func NewGetProjectsOK() *GetProjectsOK {
	return &GetProjectsOK{}
}

/*GetProjectsOK handles this case with default header values.

successfully operation
*/
type GetProjectsOK struct {
	Payload []models.Project
}

func (o *GetProjectsOK) Error() string {
	return fmt.Sprintf("[GET /projects][%d] getProjectsOK  %+v", 200, o.Payload)
}

func (o *GetProjectsOK) GetPayload() []models.Project {
	return o.Payload
}

func (o *GetProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectsBadRequest creates a GetProjectsBadRequest with default headers values
func NewGetProjectsBadRequest() *GetProjectsBadRequest {
	return &GetProjectsBadRequest{}
}

/*GetProjectsBadRequest handles this case with default header values.

Bad request
*/
type GetProjectsBadRequest struct {
}

func (o *GetProjectsBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects][%d] getProjectsBadRequest ", 400)
}

func (o *GetProjectsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsInternalServerError creates a GetProjectsInternalServerError with default headers values
func NewGetProjectsInternalServerError() *GetProjectsInternalServerError {
	return &GetProjectsInternalServerError{}
}

/*GetProjectsInternalServerError handles this case with default header values.

Internal error
*/
type GetProjectsInternalServerError struct {
}

func (o *GetProjectsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects][%d] getProjectsInternalServerError ", 500)
}

func (o *GetProjectsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
