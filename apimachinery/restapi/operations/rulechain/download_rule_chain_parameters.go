// Code generated by go-swagger; DO NOT EDIT.

package rulechain

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDownloadRuleChainParams creates a new DownloadRuleChainParams object
// no default values defined in spec.
func NewDownloadRuleChainParams() DownloadRuleChainParams {

	return DownloadRuleChainParams{}
}

// DownloadRuleChainParams contains all the bound params for the download rule chain operation
// typically these are obtained from a http.Request
//
// swagger:parameters downloadRuleChain
type DownloadRuleChainParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*download rule chain id
	  Required: true
	  In: path
	*/
	RuleChainID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDownloadRuleChainParams() beforehand.
func (o *DownloadRuleChainParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rRuleChainID, rhkRuleChainID, _ := route.Params.GetOK("ruleChainId")
	if err := o.bindRuleChainID(rRuleChainID, rhkRuleChainID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindRuleChainID binds and validates parameter RuleChainID from path.
func (o *DownloadRuleChainParams) bindRuleChainID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.RuleChainID = raw

	return nil
}