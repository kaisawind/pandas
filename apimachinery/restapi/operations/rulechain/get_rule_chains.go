// Code generated by go-swagger; DO NOT EDIT.

package rulechain

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/cloustone/pandas/models"
)

// GetRuleChainsHandlerFunc turns a function with the right signature into a get rule chains handler
type GetRuleChainsHandlerFunc func(GetRuleChainsParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRuleChainsHandlerFunc) Handle(params GetRuleChainsParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetRuleChainsHandler interface for that can handle valid get rule chains params
type GetRuleChainsHandler interface {
	Handle(GetRuleChainsParams, *models.Principal) middleware.Responder
}

// NewGetRuleChains creates a new http.Handler for the get rule chains operation
func NewGetRuleChains(ctx *middleware.Context, handler GetRuleChainsHandler) *GetRuleChains {
	return &GetRuleChains{Context: ctx, Handler: handler}
}

/*GetRuleChains swagger:route GET /rulechains Rulechain getRuleChains

get all of rule chains

get all of rule chains

*/
type GetRuleChains struct {
	Context *middleware.Context
	Handler GetRuleChainsHandler
}

func (o *GetRuleChains) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRuleChainsParams()

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