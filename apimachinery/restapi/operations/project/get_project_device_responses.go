// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cloustone/pandas/models"
)

// GetProjectDeviceOKCode is the HTTP code returned for type GetProjectDeviceOK
const GetProjectDeviceOKCode int = 200

/*GetProjectDeviceOK successful operation

swagger:response getProjectDeviceOK
*/
type GetProjectDeviceOK struct {

	/*
	  In: Body
	*/
	Payload models.Device `json:"body,omitempty"`
}

// NewGetProjectDeviceOK creates GetProjectDeviceOK with default headers values
func NewGetProjectDeviceOK() *GetProjectDeviceOK {

	return &GetProjectDeviceOK{}
}

// WithPayload adds the payload to the get project device o k response
func (o *GetProjectDeviceOK) WithPayload(payload models.Device) *GetProjectDeviceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project device o k response
func (o *GetProjectDeviceOK) SetPayload(payload models.Device) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectDeviceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetProjectDeviceNotFoundCode is the HTTP code returned for type GetProjectDeviceNotFound
const GetProjectDeviceNotFoundCode int = 404

/*GetProjectDeviceNotFound project or device not found

swagger:response getProjectDeviceNotFound
*/
type GetProjectDeviceNotFound struct {
}

// NewGetProjectDeviceNotFound creates GetProjectDeviceNotFound with default headers values
func NewGetProjectDeviceNotFound() *GetProjectDeviceNotFound {

	return &GetProjectDeviceNotFound{}
}

// WriteResponse to the client
func (o *GetProjectDeviceNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
