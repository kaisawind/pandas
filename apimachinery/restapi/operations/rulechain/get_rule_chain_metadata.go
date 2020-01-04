// Code generated by go-swagger; DO NOT EDIT.

package rulechain

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/cloustone/pandas/models"
)

// GetRuleChainMetadataHandlerFunc turns a function with the right signature into a get rule chain metadata handler
type GetRuleChainMetadataHandlerFunc func(GetRuleChainMetadataParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRuleChainMetadataHandlerFunc) Handle(params GetRuleChainMetadataParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetRuleChainMetadataHandler interface for that can handle valid get rule chain metadata params
type GetRuleChainMetadataHandler interface {
	Handle(GetRuleChainMetadataParams, *models.Principal) middleware.Responder
}

// NewGetRuleChainMetadata creates a new http.Handler for the get rule chain metadata operation
func NewGetRuleChainMetadata(ctx *middleware.Context, handler GetRuleChainMetadataHandler) *GetRuleChainMetadata {
	return &GetRuleChainMetadata{Context: ctx, Handler: handler}
}

/*GetRuleChainMetadata swagger:route GET /rulechains/{ruleChainId}/metadata Rulechain getRuleChainMetadata

get meta data of perticular rule chain

get meta data of perticular rule chain

*/
type GetRuleChainMetadata struct {
	Context *middleware.Context
	Handler GetRuleChainMetadataHandler
}

func (o *GetRuleChainMetadata) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRuleChainMetadataParams()

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