// Code generated by go-swagger; DO NOT EDIT.

package rulechain

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"

	"github.com/cloustone/pandas/models"
)

// DownloadRuleChainHandlerFunc turns a function with the right signature into a download rule chain handler
type DownloadRuleChainHandlerFunc func(DownloadRuleChainParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn DownloadRuleChainHandlerFunc) Handle(params DownloadRuleChainParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// DownloadRuleChainHandler interface for that can handle valid download rule chain params
type DownloadRuleChainHandler interface {
	Handle(DownloadRuleChainParams, *models.Principal) middleware.Responder
}

// NewDownloadRuleChain creates a new http.Handler for the download rule chain operation
func NewDownloadRuleChain(ctx *middleware.Context, handler DownloadRuleChainHandler) *DownloadRuleChain {
	return &DownloadRuleChain{Context: ctx, Handler: handler}
}

/*DownloadRuleChain swagger:route POST /rulechains/{ruleChainId}/download Rulechain downloadRuleChain

download all infomation of one rule chain to the local pc

download all infomation of one rule chain to the local pc

*/
type DownloadRuleChain struct {
	Context *middleware.Context
	Handler DownloadRuleChainHandler
}

func (o *DownloadRuleChain) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDownloadRuleChainParams()

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

// DownloadRuleChainOKBody download rule chain o k body
// swagger:model DownloadRuleChainOKBody
type DownloadRuleChainOKBody struct {

	// the result of excution
	Status bool `json:"status,omitempty"`
}

// Validate validates this download rule chain o k body
func (o *DownloadRuleChainOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DownloadRuleChainOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DownloadRuleChainOKBody) UnmarshalBinary(b []byte) error {
	var res DownloadRuleChainOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
