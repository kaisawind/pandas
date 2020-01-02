// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/cloustone/pandas/models"
)

// NewCreateModelParams creates a new CreateModelParams object
// no default values defined in spec.
func NewCreateModelParams() CreateModelParams {

	return CreateModelParams{}
}

// CreateModelParams contains all the bound params for the create model operation
// typically these are obtained from a http.Request
//
// swagger:parameters createModel
type CreateModelParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	DeviceModel models.DeviceModel
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateModelParams() beforehand.
func (o *CreateModelParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.DeviceModel
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("deviceModel", "body", "", err))
		} else {
			// no validation on generic interface
			o.DeviceModel = body
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
