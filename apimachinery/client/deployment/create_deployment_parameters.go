// Code generated by go-swagger; DO NOT EDIT.

package deployment

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

	"github.com/cloustone/pandas/models"
)

// NewCreateDeploymentParams creates a new CreateDeploymentParams object
// with the default values initialized.
func NewCreateDeploymentParams() *CreateDeploymentParams {
	var ()
	return &CreateDeploymentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDeploymentParamsWithTimeout creates a new CreateDeploymentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateDeploymentParamsWithTimeout(timeout time.Duration) *CreateDeploymentParams {
	var ()
	return &CreateDeploymentParams{

		timeout: timeout,
	}
}

// NewCreateDeploymentParamsWithContext creates a new CreateDeploymentParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateDeploymentParamsWithContext(ctx context.Context) *CreateDeploymentParams {
	var ()
	return &CreateDeploymentParams{

		Context: ctx,
	}
}

// NewCreateDeploymentParamsWithHTTPClient creates a new CreateDeploymentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateDeploymentParamsWithHTTPClient(client *http.Client) *CreateDeploymentParams {
	var ()
	return &CreateDeploymentParams{
		HTTPClient: client,
	}
}

/*CreateDeploymentParams contains all the parameters to send to the API endpoint
for the create deployment operation typically these are written to a http.Request
*/
type CreateDeploymentParams struct {

	/*Deployment
	  deployment information

	*/
	Deployment models.Deployment

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create deployment params
func (o *CreateDeploymentParams) WithTimeout(timeout time.Duration) *CreateDeploymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create deployment params
func (o *CreateDeploymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create deployment params
func (o *CreateDeploymentParams) WithContext(ctx context.Context) *CreateDeploymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create deployment params
func (o *CreateDeploymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create deployment params
func (o *CreateDeploymentParams) WithHTTPClient(client *http.Client) *CreateDeploymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create deployment params
func (o *CreateDeploymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeployment adds the deployment to the create deployment params
func (o *CreateDeploymentParams) WithDeployment(deployment models.Deployment) *CreateDeploymentParams {
	o.SetDeployment(deployment)
	return o
}

// SetDeployment adds the deployment to the create deployment params
func (o *CreateDeploymentParams) SetDeployment(deployment models.Deployment) {
	o.Deployment = deployment
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDeploymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Deployment != nil {
		if err := r.SetBodyParam(o.Deployment); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}