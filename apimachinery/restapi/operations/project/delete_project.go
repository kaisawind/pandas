// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/cloustone/pandas/models"
)

// DeleteProjectHandlerFunc turns a function with the right signature into a delete project handler
type DeleteProjectHandlerFunc func(DeleteProjectParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteProjectHandlerFunc) Handle(params DeleteProjectParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// DeleteProjectHandler interface for that can handle valid delete project params
type DeleteProjectHandler interface {
	Handle(DeleteProjectParams, *models.Principal) middleware.Responder
}

// NewDeleteProject creates a new http.Handler for the delete project operation
func NewDeleteProject(ctx *middleware.Context, handler DeleteProjectHandler) *DeleteProject {
	return &DeleteProject{Context: ctx, Handler: handler}
}

/*DeleteProject swagger:route DELETE /projects/{projectId} Project deleteProject

delete project

delete specifed project

*/
type DeleteProject struct {
	Context *middleware.Context
	Handler DeleteProjectHandler
}

func (o *DeleteProject) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteProjectParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
