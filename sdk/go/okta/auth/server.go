// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package auth

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates an Authorization Server.
// 
// This resource allows you to create and configure an Authorization Server.
//
// > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/auth_server.html.markdown.
type Server struct {
	s *pulumi.ResourceState
}

// NewServer registers a new resource with the given unique name, arguments, and options.
func NewServer(ctx *pulumi.Context,
	name string, args *ServerArgs, opts ...pulumi.ResourceOpt) (*Server, error) {
	if args == nil || args.Audiences == nil {
		return nil, errors.New("missing required argument 'Audiences'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["audiences"] = nil
		inputs["credentialsRotationMode"] = nil
		inputs["description"] = nil
		inputs["issuerMode"] = nil
		inputs["name"] = nil
		inputs["status"] = nil
	} else {
		inputs["audiences"] = args.Audiences
		inputs["credentialsRotationMode"] = args.CredentialsRotationMode
		inputs["description"] = args.Description
		inputs["issuerMode"] = args.IssuerMode
		inputs["name"] = args.Name
		inputs["status"] = args.Status
	}
	inputs["credentialsLastRotated"] = nil
	inputs["credentialsNextRotation"] = nil
	inputs["issuer"] = nil
	inputs["kid"] = nil
	s, err := ctx.RegisterResource("okta:auth/server:Server", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Server{s: s}, nil
}

// GetServer gets an existing Server resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServer(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ServerState, opts ...pulumi.ResourceOpt) (*Server, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["audiences"] = state.Audiences
		inputs["credentialsLastRotated"] = state.CredentialsLastRotated
		inputs["credentialsNextRotation"] = state.CredentialsNextRotation
		inputs["credentialsRotationMode"] = state.CredentialsRotationMode
		inputs["description"] = state.Description
		inputs["issuer"] = state.Issuer
		inputs["issuerMode"] = state.IssuerMode
		inputs["kid"] = state.Kid
		inputs["name"] = state.Name
		inputs["status"] = state.Status
	}
	s, err := ctx.ReadResource("okta:auth/server:Server", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Server{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Server) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Server) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The recipients that the tokens are intended for. This becomes the `aud` claim in an access token.
func (r *Server) Audiences() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["audiences"])
}

// The timestamp when the authorization server started to use the `kid` for signing tokens.
func (r *Server) CredentialsLastRotated() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["credentialsLastRotated"])
}

// The timestamp when the authorization server changes the key for signing tokens. Only returned when `credentialsRotationMode` is `"AUTO"`.
func (r *Server) CredentialsNextRotation() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["credentialsNextRotation"])
}

// The key rotation mode for the authorization server. Can be `"AUTO"` or `"MANUAL"`.
func (r *Server) CredentialsRotationMode() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["credentialsRotationMode"])
}

// The description of the authorization server.
func (r *Server) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// The complete URL for a Custom Authorization Server. This becomes the `iss` claim in an access token.
func (r *Server) Issuer() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["issuer"])
}

// Allows you to use a custom issuer URL. It can be set to `"CUSTOM_URL"` or `"ORG_URL"`
func (r *Server) IssuerMode() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["issuerMode"])
}

// The ID of the JSON Web Key used for signing tokens issued by the authorization server.
func (r *Server) Kid() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["kid"])
}

// The name of the authorization server.
func (r *Server) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The status of the auth server. It defaults to `"ACTIVE"`
func (r *Server) Status() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["status"])
}

// Input properties used for looking up and filtering Server resources.
type ServerState struct {
	// The recipients that the tokens are intended for. This becomes the `aud` claim in an access token.
	Audiences interface{}
	// The timestamp when the authorization server started to use the `kid` for signing tokens.
	CredentialsLastRotated interface{}
	// The timestamp when the authorization server changes the key for signing tokens. Only returned when `credentialsRotationMode` is `"AUTO"`.
	CredentialsNextRotation interface{}
	// The key rotation mode for the authorization server. Can be `"AUTO"` or `"MANUAL"`.
	CredentialsRotationMode interface{}
	// The description of the authorization server.
	Description interface{}
	// The complete URL for a Custom Authorization Server. This becomes the `iss` claim in an access token.
	Issuer interface{}
	// Allows you to use a custom issuer URL. It can be set to `"CUSTOM_URL"` or `"ORG_URL"`
	IssuerMode interface{}
	// The ID of the JSON Web Key used for signing tokens issued by the authorization server.
	Kid interface{}
	// The name of the authorization server.
	Name interface{}
	// The status of the auth server. It defaults to `"ACTIVE"`
	Status interface{}
}

// The set of arguments for constructing a Server resource.
type ServerArgs struct {
	// The recipients that the tokens are intended for. This becomes the `aud` claim in an access token.
	Audiences interface{}
	// The key rotation mode for the authorization server. Can be `"AUTO"` or `"MANUAL"`.
	CredentialsRotationMode interface{}
	// The description of the authorization server.
	Description interface{}
	// Allows you to use a custom issuer URL. It can be set to `"CUSTOM_URL"` or `"ORG_URL"`
	IssuerMode interface{}
	// The name of the authorization server.
	Name interface{}
	// The status of the auth server. It defaults to `"ACTIVE"`
	Status interface{}
}
