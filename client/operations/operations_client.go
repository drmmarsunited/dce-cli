// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteAccountsID deletes an account by ID
*/
func (a *Client) DeleteAccountsID(params *DeleteAccountsIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteAccountsIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAccountsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteAccountsID",
		Method:             "DELETE",
		PathPattern:        "/accounts/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAccountsIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteAccountsIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteAccountsID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteLeases removes a lease
*/
func (a *Client) DeleteLeases(params *DeleteLeasesParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteLeasesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLeasesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteLeases",
		Method:             "DELETE",
		PathPattern:        "/leases",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteLeasesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteLeasesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteLeases: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAccounts lists accounts
*/
func (a *Client) GetAccounts(params *GetAccountsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAccounts",
		Method:             "GET",
		PathPattern:        "/accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccountsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAccounts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAccountsID gets a specific account by an account ID
*/
func (a *Client) GetAccountsID(params *GetAccountsIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetAccountsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAccountsID",
		Method:             "GET",
		PathPattern:        "/accounts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccountsIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAccountsIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAccountsID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetLeases gets leases
*/
func (a *Client) GetLeases(params *GetLeasesParams, authInfo runtime.ClientAuthInfoWriter) (*GetLeasesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLeasesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLeases",
		Method:             "GET",
		PathPattern:        "/leases",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLeasesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLeasesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetLeases: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetLeasesID gets a lease by Id
*/
func (a *Client) GetLeasesID(params *GetLeasesIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetLeasesIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLeasesIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLeasesID",
		Method:             "GET",
		PathPattern:        "/leases/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLeasesIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLeasesIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetLeasesID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetUsage gets usage records by date range
*/
func (a *Client) GetUsage(params *GetUsageParams, authInfo runtime.ClientAuthInfoWriter) (*GetUsageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetUsage",
		Method:             "GET",
		PathPattern:        "/usage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUsageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUsageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetUsage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostAccounts adds an a w s account to the account pool
*/
func (a *Client) PostAccounts(params *PostAccountsParams, authInfo runtime.ClientAuthInfoWriter) (*PostAccountsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAccountsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostAccounts",
		Method:             "POST",
		PathPattern:        "/accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAccountsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostAccountsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAccounts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostLeases creates a new lease
*/
func (a *Client) PostLeases(params *PostLeasesParams, authInfo runtime.ClientAuthInfoWriter) (*PostLeasesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLeasesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostLeases",
		Method:             "POST",
		PathPattern:        "/leases",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostLeasesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostLeasesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostLeases: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostLeasesIDAuth creates lease authentication by Id
*/
func (a *Client) PostLeasesIDAuth(params *PostLeasesIDAuthParams, authInfo runtime.ClientAuthInfoWriter) (*PostLeasesIDAuthCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLeasesIDAuthParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostLeasesIDAuth",
		Method:             "POST",
		PathPattern:        "/leases/{id}/auth",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostLeasesIDAuthReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostLeasesIDAuthCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostLeasesIDAuth: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
