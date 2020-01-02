// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/cloustone/pandas/models"
)

// DeleteModelHandlerFunc turns a function with the right signature into a delete model handler
type DeleteModelHandlerFunc func(DeleteModelParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteModelHandlerFunc) Handle(params DeleteModelParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// DeleteModelHandler interface for that can handle valid delete model params
type DeleteModelHandler interface {
	Handle(DeleteModelParams, *models.Principal) middleware.Responder
}

// NewDeleteModel creates a new http.Handler for the delete model operation
func NewDeleteModel(ctx *middleware.Context, handler DeleteModelHandler) *DeleteModel {
	return &DeleteModel{Context: ctx, Handler: handler}
}

/*DeleteModel swagger:route DELETE /models/{modelId} Model deleteModel

delete device's model

*/
type DeleteModel struct {
	Context *middleware.Context
	Handler DeleteModelHandler
}

func (o *DeleteModel) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteModelParams()

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
