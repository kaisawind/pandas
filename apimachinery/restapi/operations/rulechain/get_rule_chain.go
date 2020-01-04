// Code generated by go-swagger; DO NOT EDIT.

package rulechain

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/cloustone/pandas/models"
)

// GetRuleChainHandlerFunc turns a function with the right signature into a get rule chain handler
type GetRuleChainHandlerFunc func(GetRuleChainParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRuleChainHandlerFunc) Handle(params GetRuleChainParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetRuleChainHandler interface for that can handle valid get rule chain params
type GetRuleChainHandler interface {
	Handle(GetRuleChainParams, *models.Principal) middleware.Responder
}

// NewGetRuleChain creates a new http.Handler for the get rule chain operation
func NewGetRuleChain(ctx *middleware.Context, handler GetRuleChainHandler) *GetRuleChain {
	return &GetRuleChain{Context: ctx, Handler: handler}
}

/*GetRuleChain swagger:route GET /rulechains/{ruleChainId} Rulechain getRuleChain

get rule chain by id

get  rule chain by id

*/
type GetRuleChain struct {
	Context *middleware.Context
	Handler GetRuleChainHandler
}

func (o *GetRuleChain) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRuleChainParams()

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