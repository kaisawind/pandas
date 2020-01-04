// Code generated by go-swagger; DO NOT EDIT.

package deployment

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

// NewUpdateDeploymentParams creates a new UpdateDeploymentParams object
// no default values defined in spec.
func NewUpdateDeploymentParams() UpdateDeploymentParams {

	return UpdateDeploymentParams{}
}

// UpdateDeploymentParams contains all the bound params for the update deployment operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateDeployment
type UpdateDeploymentParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*updated deployment
	  Required: true
	  In: body
	*/
	Deployment models.Deployment
	/*deployment identifier
	  Required: true
	  In: path
	*/
	DeploymentID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateDeploymentParams() beforehand.
func (o *UpdateDeploymentParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Deployment
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("deployment", "body"))
			} else {
				res = append(res, errors.NewParseError("deployment", "body", "", err))
			}
		} else {
			// no validation on generic interface
			o.Deployment = body
		}
	} else {
		res = append(res, errors.Required("deployment", "body"))
	}
	rDeploymentID, rhkDeploymentID, _ := route.Params.GetOK("deploymentId")
	if err := o.bindDeploymentID(rDeploymentID, rhkDeploymentID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDeploymentID binds and validates parameter DeploymentID from path.
func (o *UpdateDeploymentParams) bindDeploymentID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.DeploymentID = raw

	return nil
}