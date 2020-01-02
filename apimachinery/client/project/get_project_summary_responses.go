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

// GetProjectSummaryReader is a Reader for the GetProjectSummary structure.
type GetProjectSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetProjectSummaryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProjectSummaryOK creates a GetProjectSummaryOK with default headers values
func NewGetProjectSummaryOK() *GetProjectSummaryOK {
	return &GetProjectSummaryOK{}
}

/*GetProjectSummaryOK handles this case with default header values.

successful operation
*/
type GetProjectSummaryOK struct {
	Payload models.ProjectSummary
}

func (o *GetProjectSummaryOK) Error() string {
	return fmt.Sprintf("[GET /projects/{projectId}/summary][%d] getProjectSummaryOK  %+v", 200, o.Payload)
}

func (o *GetProjectSummaryOK) GetPayload() models.ProjectSummary {
	return o.Payload
}

func (o *GetProjectSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectSummaryNotFound creates a GetProjectSummaryNotFound with default headers values
func NewGetProjectSummaryNotFound() *GetProjectSummaryNotFound {
	return &GetProjectSummaryNotFound{}
}

/*GetProjectSummaryNotFound handles this case with default header values.

project not found
*/
type GetProjectSummaryNotFound struct {
}

func (o *GetProjectSummaryNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{projectId}/summary][%d] getProjectSummaryNotFound ", 404)
}

func (o *GetProjectSummaryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
