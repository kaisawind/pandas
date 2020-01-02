// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cloustone/pandas/apimachinery/client/deployment"
	"github.com/cloustone/pandas/apimachinery/client/device"
	"github.com/cloustone/pandas/apimachinery/client/logs"
	"github.com/cloustone/pandas/apimachinery/client/model"
	"github.com/cloustone/pandas/apimachinery/client/project"
	"github.com/cloustone/pandas/apimachinery/client/rulechain"
)

// Default swagger pandas HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost:8080"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/v1"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new swagger pandas HTTP client.
func NewHTTPClient(formats strfmt.Registry) *SwaggerPandas {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new swagger pandas HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *SwaggerPandas {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new swagger pandas client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *SwaggerPandas {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(SwaggerPandas)
	cli.Transport = transport

	cli.Deployment = deployment.New(transport, formats)

	cli.Device = device.New(transport, formats)

	cli.Logs = logs.New(transport, formats)

	cli.Model = model.New(transport, formats)

	cli.Project = project.New(transport, formats)

	cli.Rulechain = rulechain.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// SwaggerPandas is a client for swagger pandas
type SwaggerPandas struct {
	Deployment *deployment.Client

	Device *device.Client

	Logs *logs.Client

	Model *model.Client

	Project *project.Client

	Rulechain *rulechain.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *SwaggerPandas) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Deployment.SetTransport(transport)

	c.Device.SetTransport(transport)

	c.Logs.SetTransport(transport)

	c.Model.SetTransport(transport)

	c.Project.SetTransport(transport)

	c.Rulechain.SetTransport(transport)

}
