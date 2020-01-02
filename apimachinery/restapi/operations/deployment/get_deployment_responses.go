// Code generated by go-swagger; DO NOT EDIT.

package deployment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cloustone/pandas/models"
)

// GetDeploymentOKCode is the HTTP code returned for type GetDeploymentOK
const GetDeploymentOKCode int = 200

/*GetDeploymentOK successfully operation

swagger:response getDeploymentOK
*/
type GetDeploymentOK struct {

	/*
	  In: Body
	*/
	Payload models.Deployment `json:"body,omitempty"`
}

// NewGetDeploymentOK creates GetDeploymentOK with default headers values
func NewGetDeploymentOK() *GetDeploymentOK {

	return &GetDeploymentOK{}
}

// WithPayload adds the payload to the get deployment o k response
func (o *GetDeploymentOK) WithPayload(payload models.Deployment) *GetDeploymentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get deployment o k response
func (o *GetDeploymentOK) SetPayload(payload models.Deployment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDeploymentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetDeploymentNotFoundCode is the HTTP code returned for type GetDeploymentNotFound
const GetDeploymentNotFoundCode int = 404

/*GetDeploymentNotFound deployment not found

swagger:response getDeploymentNotFound
*/
type GetDeploymentNotFound struct {
}

// NewGetDeploymentNotFound creates GetDeploymentNotFound with default headers values
func NewGetDeploymentNotFound() *GetDeploymentNotFound {

	return &GetDeploymentNotFound{}
}

// WriteResponse to the client
func (o *GetDeploymentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
