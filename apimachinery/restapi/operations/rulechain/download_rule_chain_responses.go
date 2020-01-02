// Code generated by go-swagger; DO NOT EDIT.

package rulechain

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DownloadRuleChainOKCode is the HTTP code returned for type DownloadRuleChainOK
const DownloadRuleChainOKCode int = 200

/*DownloadRuleChainOK excute successfully

swagger:response downloadRuleChainOK
*/
type DownloadRuleChainOK struct {

	/*
	  In: Body
	*/
	Payload *DownloadRuleChainOKBody `json:"body,omitempty"`
}

// NewDownloadRuleChainOK creates DownloadRuleChainOK with default headers values
func NewDownloadRuleChainOK() *DownloadRuleChainOK {

	return &DownloadRuleChainOK{}
}

// WithPayload adds the payload to the download rule chain o k response
func (o *DownloadRuleChainOK) WithPayload(payload *DownloadRuleChainOKBody) *DownloadRuleChainOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download rule chain o k response
func (o *DownloadRuleChainOK) SetPayload(payload *DownloadRuleChainOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadRuleChainOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadRuleChainInternalServerErrorCode is the HTTP code returned for type DownloadRuleChainInternalServerError
const DownloadRuleChainInternalServerErrorCode int = 500

/*DownloadRuleChainInternalServerError Server internal error

swagger:response downloadRuleChainInternalServerError
*/
type DownloadRuleChainInternalServerError struct {
}

// NewDownloadRuleChainInternalServerError creates DownloadRuleChainInternalServerError with default headers values
func NewDownloadRuleChainInternalServerError() *DownloadRuleChainInternalServerError {

	return &DownloadRuleChainInternalServerError{}
}

// WriteResponse to the client
func (o *DownloadRuleChainInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
