// Code generated by go-swagger; DO NOT EDIT.

package rulechain

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new rulechain API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for rulechain API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteRuleChain deletes rule chain

delete rule chain with Id
*/
func (a *Client) DeleteRuleChain(params *DeleteRuleChainParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteRuleChainOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRuleChainParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteRuleChain",
		Method:             "DELETE",
		PathPattern:        "/rulechains/{ruleChainId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteRuleChainReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteRuleChainOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteRuleChain: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DownloadRuleChain downloads all infomation of one rule chain to the local pc

download all infomation of one rule chain to the local pc
*/
func (a *Client) DownloadRuleChain(params *DownloadRuleChainParams, authInfo runtime.ClientAuthInfoWriter) (*DownloadRuleChainOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDownloadRuleChainParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "downloadRuleChain",
		Method:             "POST",
		PathPattern:        "/rulechains/{ruleChainId}/download",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DownloadRuleChainReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DownloadRuleChainOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for downloadRuleChain: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetRuleChain gets rule chain by id

get  rule chain by id
*/
func (a *Client) GetRuleChain(params *GetRuleChainParams, authInfo runtime.ClientAuthInfoWriter) (*GetRuleChainOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRuleChainParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRuleChain",
		Method:             "GET",
		PathPattern:        "/rulechains/{ruleChainId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRuleChainReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRuleChainOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRuleChain: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetRuleChainMetadata gets meta data of perticular rule chain

get meta data of perticular rule chain
*/
func (a *Client) GetRuleChainMetadata(params *GetRuleChainMetadataParams, authInfo runtime.ClientAuthInfoWriter) (*GetRuleChainMetadataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRuleChainMetadataParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRuleChainMetadata",
		Method:             "GET",
		PathPattern:        "/rulechains/{ruleChainId}/metadata",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRuleChainMetadataReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRuleChainMetadataOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRuleChainMetadata: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetRuleChains gets all of rule chains

get all of rule chains
*/
func (a *Client) GetRuleChains(params *GetRuleChainsParams, authInfo runtime.ClientAuthInfoWriter) (*GetRuleChainsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRuleChainsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRuleChains",
		Method:             "GET",
		PathPattern:        "/rulechains",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRuleChainsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRuleChainsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRuleChains: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SaveRuleChain saves one rule chain

save one rule chain
*/
func (a *Client) SaveRuleChain(params *SaveRuleChainParams, authInfo runtime.ClientAuthInfoWriter) (*SaveRuleChainOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSaveRuleChainParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "saveRuleChain",
		Method:             "POST",
		PathPattern:        "/rulechains/{ruleChainId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SaveRuleChainReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SaveRuleChainOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for saveRuleChain: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SaveRuleChainMetadata saves meta data

save meta data
*/
func (a *Client) SaveRuleChainMetadata(params *SaveRuleChainMetadataParams, authInfo runtime.ClientAuthInfoWriter) (*SaveRuleChainMetadataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSaveRuleChainMetadataParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "saveRuleChainMetadata",
		Method:             "POST",
		PathPattern:        "/rulechains/{ruleChainId}/metadata",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SaveRuleChainMetadataReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SaveRuleChainMetadataOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for saveRuleChainMetadata: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SetRootRuleChain sets root id

set root id
*/
func (a *Client) SetRootRuleChain(params *SetRootRuleChainParams, authInfo runtime.ClientAuthInfoWriter) (*SetRootRuleChainOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetRootRuleChainParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setRootRuleChain",
		Method:             "POST",
		PathPattern:        "/rulechains/{ruleChainId}/root",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SetRootRuleChainReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetRootRuleChainOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for setRootRuleChain: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UploadRuleChain uploads all infomation of one rule chain to the system

upload all infomation of one rule chain to the system
*/
func (a *Client) UploadRuleChain(params *UploadRuleChainParams, authInfo runtime.ClientAuthInfoWriter) (*UploadRuleChainOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadRuleChainParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "uploadRuleChain",
		Method:             "POST",
		PathPattern:        "/rulechains/{ruleChainId}/upload",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UploadRuleChainReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UploadRuleChainOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for uploadRuleChain: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
