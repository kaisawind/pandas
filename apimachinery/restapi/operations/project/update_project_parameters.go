// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cloustone/pandas/models"
)

// NewUpdateProjectParams creates a new UpdateProjectParams object
// no default values defined in spec.
func NewUpdateProjectParams() UpdateProjectParams {

	return UpdateProjectParams{}
}

// UpdateProjectParams contains all the bound params for the update project operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateProject
type UpdateProjectParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*updated project
	  Required: true
	  In: body
	*/
	Project models.Project
	/*project identifier
	  Required: true
	  In: path
	*/
	ProjectID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateProjectParams() beforehand.
func (o *UpdateProjectParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Project
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("project", "body"))
			} else {
				res = append(res, errors.NewParseError("project", "body", "", err))
			}
		} else {
			// no validation on generic interface
			o.Project = body
		}
	} else {
		res = append(res, errors.Required("project", "body"))
	}
	rProjectID, rhkProjectID, _ := route.Params.GetOK("projectId")
	if err := o.bindProjectID(rProjectID, rhkProjectID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindProjectID binds and validates parameter ProjectID from path.
func (o *UpdateProjectParams) bindProjectID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ProjectID = raw

	return nil
}
