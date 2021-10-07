// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new user API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateUser(params *CreateUserParams, authInfo runtime.ClientAuthInfoWriter) (*CreateUserCreated, error)

	DeleteUser(params *DeleteUserParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteUserOK, error)

	GetCurrentUserInfo(params *GetCurrentUserInfoParams, authInfo runtime.ClientAuthInfoWriter) (*GetCurrentUserInfoOK, error)

	GetCurrentUserPermissions(params *GetCurrentUserPermissionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCurrentUserPermissionsOK, error)

	GetUser(params *GetUserParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserOK, error)

	ListUsers(params *ListUsersParams, authInfo runtime.ClientAuthInfoWriter) (*ListUsersOK, error)

	SearchUsers(params *SearchUsersParams, authInfo runtime.ClientAuthInfoWriter) (*SearchUsersOK, error)

	SetCliSecret(params *SetCliSecretParams, authInfo runtime.ClientAuthInfoWriter) (*SetCliSecretOK, error)

	SetUserSysAdmin(params *SetUserSysAdminParams, authInfo runtime.ClientAuthInfoWriter) (*SetUserSysAdminOK, error)

	UpdateUserPassword(params *UpdateUserPasswordParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateUserPasswordOK, error)

	UpdateUserProfile(params *UpdateUserProfileParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateUserProfileOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateUser creates a local user

  This API can be used only when the authentication mode is for local DB.  When self registration is disabled.
*/
func (a *Client) CreateUser(params *CreateUserParams, authInfo runtime.ClientAuthInfoWriter) (*CreateUserCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createUser",
		Method:             "POST",
		PathPattern:        "/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateUserCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteUser marks a registered user as be removed

  This endpoint let administrator of Harbor mark a registered user as removed.It actually won't be deleted from DB.

*/
func (a *Client) DeleteUser(params *DeleteUserParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteUser",
		Method:             "DELETE",
		PathPattern:        "/users/{user_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCurrentUserInfo gets current user info
*/
func (a *Client) GetCurrentUserInfo(params *GetCurrentUserInfoParams, authInfo runtime.ClientAuthInfoWriter) (*GetCurrentUserInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCurrentUserInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCurrentUserInfo",
		Method:             "GET",
		PathPattern:        "/users/current",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCurrentUserInfoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCurrentUserInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCurrentUserInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCurrentUserPermissions gets current user permissions
*/
func (a *Client) GetCurrentUserPermissions(params *GetCurrentUserPermissionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCurrentUserPermissionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCurrentUserPermissionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCurrentUserPermissions",
		Method:             "GET",
		PathPattern:        "/users/current/permissions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCurrentUserPermissionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCurrentUserPermissionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCurrentUserPermissions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUser gets a user s profile
*/
func (a *Client) GetUser(params *GetUserParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUser",
		Method:             "GET",
		PathPattern:        "/users/{user_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListUsers lists users
*/
func (a *Client) ListUsers(params *ListUsersParams, authInfo runtime.ClientAuthInfoWriter) (*ListUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListUsersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listUsers",
		Method:             "GET",
		PathPattern:        "/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListUsersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listUsers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SearchUsers searches users by username

  This endpoint is to search the users by username.  It's open for all authenticated requests.

*/
func (a *Client) SearchUsers(params *SearchUsersParams, authInfo runtime.ClientAuthInfoWriter) (*SearchUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchUsersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "searchUsers",
		Method:             "GET",
		PathPattern:        "/users/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SearchUsersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SearchUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for searchUsers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SetCliSecret sets c l i secret for a user

  This endpoint let user generate a new CLI secret for himself.  This API only works when auth mode is set to 'OIDC'. Once this API returns with successful status, the old secret will be invalid, as there will be only one CLI secret for a user.
*/
func (a *Client) SetCliSecret(params *SetCliSecretParams, authInfo runtime.ClientAuthInfoWriter) (*SetCliSecretOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetCliSecretParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setCliSecret",
		Method:             "PUT",
		PathPattern:        "/users/{user_id}/cli_secret",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SetCliSecretReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetCliSecretOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for setCliSecret: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SetUserSysAdmin updates a registered user to change to be an administrator of harbor
*/
func (a *Client) SetUserSysAdmin(params *SetUserSysAdminParams, authInfo runtime.ClientAuthInfoWriter) (*SetUserSysAdminOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetUserSysAdminParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setUserSysAdmin",
		Method:             "PUT",
		PathPattern:        "/users/{user_id}/sysadmin",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SetUserSysAdminReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetUserSysAdminOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for setUserSysAdmin: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateUserPassword changes the password on a user that already exists

  This endpoint is for user to update password. Users with the admin role can change any user's password. Regular users can change only their own password.

*/
func (a *Client) UpdateUserPassword(params *UpdateUserPasswordParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateUserPasswordOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUserPasswordParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateUserPassword",
		Method:             "PUT",
		PathPattern:        "/users/{user_id}/password",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateUserPasswordReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateUserPasswordOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateUserPassword: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateUserProfile updates user s profile
*/
func (a *Client) UpdateUserProfile(params *UpdateUserProfileParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateUserProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUserProfileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateUserProfile",
		Method:             "PUT",
		PathPattern:        "/users/{user_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateUserProfileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateUserProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateUserProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}