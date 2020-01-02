// Code generated by go-swagger; DO NOT EDIT.

package rulechain

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetRuleChainMetadataParams creates a new GetRuleChainMetadataParams object
// with the default values initialized.
func NewGetRuleChainMetadataParams() *GetRuleChainMetadataParams {
	var ()
	return &GetRuleChainMetadataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRuleChainMetadataParamsWithTimeout creates a new GetRuleChainMetadataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRuleChainMetadataParamsWithTimeout(timeout time.Duration) *GetRuleChainMetadataParams {
	var ()
	return &GetRuleChainMetadataParams{

		timeout: timeout,
	}
}

// NewGetRuleChainMetadataParamsWithContext creates a new GetRuleChainMetadataParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRuleChainMetadataParamsWithContext(ctx context.Context) *GetRuleChainMetadataParams {
	var ()
	return &GetRuleChainMetadataParams{

		Context: ctx,
	}
}

// NewGetRuleChainMetadataParamsWithHTTPClient creates a new GetRuleChainMetadataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRuleChainMetadataParamsWithHTTPClient(client *http.Client) *GetRuleChainMetadataParams {
	var ()
	return &GetRuleChainMetadataParams{
		HTTPClient: client,
	}
}

/*GetRuleChainMetadataParams contains all the parameters to send to the API endpoint
for the get rule chain metadata operation typically these are written to a http.Request
*/
type GetRuleChainMetadataParams struct {

	/*RuleChainID
	  rule chain id

	*/
	RuleChainID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get rule chain metadata params
func (o *GetRuleChainMetadataParams) WithTimeout(timeout time.Duration) *GetRuleChainMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get rule chain metadata params
func (o *GetRuleChainMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get rule chain metadata params
func (o *GetRuleChainMetadataParams) WithContext(ctx context.Context) *GetRuleChainMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get rule chain metadata params
func (o *GetRuleChainMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get rule chain metadata params
func (o *GetRuleChainMetadataParams) WithHTTPClient(client *http.Client) *GetRuleChainMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get rule chain metadata params
func (o *GetRuleChainMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRuleChainID adds the ruleChainID to the get rule chain metadata params
func (o *GetRuleChainMetadataParams) WithRuleChainID(ruleChainID string) *GetRuleChainMetadataParams {
	o.SetRuleChainID(ruleChainID)
	return o
}

// SetRuleChainID adds the ruleChainId to the get rule chain metadata params
func (o *GetRuleChainMetadataParams) SetRuleChainID(ruleChainID string) {
	o.RuleChainID = ruleChainID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRuleChainMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ruleChainId
	if err := r.SetPathParam("ruleChainId", o.RuleChainID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
