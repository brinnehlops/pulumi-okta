// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package deprecated

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type SignonPolicyRule struct {
	s *pulumi.ResourceState
}

// NewSignonPolicyRule registers a new resource with the given unique name, arguments, and options.
func NewSignonPolicyRule(ctx *pulumi.Context,
	name string, args *SignonPolicyRuleArgs, opts ...pulumi.ResourceOpt) (*SignonPolicyRule, error) {
	if args == nil || args.Policyid == nil {
		return nil, errors.New("missing required argument 'Policyid'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["access"] = nil
		inputs["authtype"] = nil
		inputs["enroll"] = nil
		inputs["mfaLifetime"] = nil
		inputs["mfaPrompt"] = nil
		inputs["mfaRememberDevice"] = nil
		inputs["mfaRequired"] = nil
		inputs["name"] = nil
		inputs["networkConnection"] = nil
		inputs["networkExcludes"] = nil
		inputs["networkIncludes"] = nil
		inputs["passwordChange"] = nil
		inputs["passwordReset"] = nil
		inputs["passwordUnlock"] = nil
		inputs["policyid"] = nil
		inputs["priority"] = nil
		inputs["sessionIdle"] = nil
		inputs["sessionLifetime"] = nil
		inputs["sessionPersistent"] = nil
		inputs["status"] = nil
		inputs["usersExcludeds"] = nil
	} else {
		inputs["access"] = args.Access
		inputs["authtype"] = args.Authtype
		inputs["enroll"] = args.Enroll
		inputs["mfaLifetime"] = args.MfaLifetime
		inputs["mfaPrompt"] = args.MfaPrompt
		inputs["mfaRememberDevice"] = args.MfaRememberDevice
		inputs["mfaRequired"] = args.MfaRequired
		inputs["name"] = args.Name
		inputs["networkConnection"] = args.NetworkConnection
		inputs["networkExcludes"] = args.NetworkExcludes
		inputs["networkIncludes"] = args.NetworkIncludes
		inputs["passwordChange"] = args.PasswordChange
		inputs["passwordReset"] = args.PasswordReset
		inputs["passwordUnlock"] = args.PasswordUnlock
		inputs["policyid"] = args.Policyid
		inputs["priority"] = args.Priority
		inputs["sessionIdle"] = args.SessionIdle
		inputs["sessionLifetime"] = args.SessionLifetime
		inputs["sessionPersistent"] = args.SessionPersistent
		inputs["status"] = args.Status
		inputs["usersExcludeds"] = args.UsersExcludeds
	}
	s, err := ctx.RegisterResource("okta:deprecated/signonPolicyRule:SignonPolicyRule", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SignonPolicyRule{s: s}, nil
}

// GetSignonPolicyRule gets an existing SignonPolicyRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSignonPolicyRule(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SignonPolicyRuleState, opts ...pulumi.ResourceOpt) (*SignonPolicyRule, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["access"] = state.Access
		inputs["authtype"] = state.Authtype
		inputs["enroll"] = state.Enroll
		inputs["mfaLifetime"] = state.MfaLifetime
		inputs["mfaPrompt"] = state.MfaPrompt
		inputs["mfaRememberDevice"] = state.MfaRememberDevice
		inputs["mfaRequired"] = state.MfaRequired
		inputs["name"] = state.Name
		inputs["networkConnection"] = state.NetworkConnection
		inputs["networkExcludes"] = state.NetworkExcludes
		inputs["networkIncludes"] = state.NetworkIncludes
		inputs["passwordChange"] = state.PasswordChange
		inputs["passwordReset"] = state.PasswordReset
		inputs["passwordUnlock"] = state.PasswordUnlock
		inputs["policyid"] = state.Policyid
		inputs["priority"] = state.Priority
		inputs["sessionIdle"] = state.SessionIdle
		inputs["sessionLifetime"] = state.SessionLifetime
		inputs["sessionPersistent"] = state.SessionPersistent
		inputs["status"] = state.Status
		inputs["usersExcludeds"] = state.UsersExcludeds
	}
	s, err := ctx.ReadResource("okta:deprecated/signonPolicyRule:SignonPolicyRule", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SignonPolicyRule{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *SignonPolicyRule) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *SignonPolicyRule) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Allow or deny access based on the rule conditions: ALLOW or DENY.
func (r *SignonPolicyRule) Access() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["access"])
}

// Authentication entrypoint: ANY or RADIUS.
func (r *SignonPolicyRule) Authtype() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["authtype"])
}

// Should the user be enrolled the first time they LOGIN, the next time they are CHALLENGEd, or NEVER?
func (r *SignonPolicyRule) Enroll() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["enroll"])
}

// Elapsed time before the next MFA challenge
func (r *SignonPolicyRule) MfaLifetime() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["mfaLifetime"])
}

// Prompt for MFA based on the device used, a factor session lifetime, or every sign on attempt: DEVICE, SESSION or ALWAYS
func (r *SignonPolicyRule) MfaPrompt() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["mfaPrompt"])
}

// Remember MFA device.
func (r *SignonPolicyRule) MfaRememberDevice() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["mfaRememberDevice"])
}

// Require MFA.
func (r *SignonPolicyRule) MfaRequired() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["mfaRequired"])
}

// Policy Rule Name
func (r *SignonPolicyRule) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Network selection mode: ANYWHERE, ZONE, ON_NETWORK, or OFF_NETWORK.
func (r *SignonPolicyRule) NetworkConnection() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["networkConnection"])
}

// The zones to exclude
func (r *SignonPolicyRule) NetworkExcludes() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["networkExcludes"])
}

// The zones to include
func (r *SignonPolicyRule) NetworkIncludes() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["networkIncludes"])
}

// Allow or deny a user to change their password: ALLOW or DENY. Default = ALLOW
func (r *SignonPolicyRule) PasswordChange() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["passwordChange"])
}

// Allow or deny a user to reset their password: ALLOW or DENY. Default = ALLOW
func (r *SignonPolicyRule) PasswordReset() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["passwordReset"])
}

// Allow or deny a user to unlock. Default = DENY
func (r *SignonPolicyRule) PasswordUnlock() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["passwordUnlock"])
}

// Policy ID of the Rule
func (r *SignonPolicyRule) Policyid() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["policyid"])
}

// Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an
// invalid priority is provided. API defaults it to the last/lowest if not there.
func (r *SignonPolicyRule) Priority() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["priority"])
}

// Max minutes a session can be idle.
func (r *SignonPolicyRule) SessionIdle() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["sessionIdle"])
}

// Max minutes a session is active: Disable = 0.
func (r *SignonPolicyRule) SessionLifetime() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["sessionLifetime"])
}

// Whether session cookies will last across browser sessions. Okta Administrators can never have persistent session
// cookies.
func (r *SignonPolicyRule) SessionPersistent() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["sessionPersistent"])
}

// Policy Rule Status: ACTIVE or INACTIVE.
func (r *SignonPolicyRule) Status() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["status"])
}

// Set of User IDs to Exclude
func (r *SignonPolicyRule) UsersExcludeds() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["usersExcludeds"])
}

// Input properties used for looking up and filtering SignonPolicyRule resources.
type SignonPolicyRuleState struct {
	// Allow or deny access based on the rule conditions: ALLOW or DENY.
	Access interface{}
	// Authentication entrypoint: ANY or RADIUS.
	Authtype interface{}
	// Should the user be enrolled the first time they LOGIN, the next time they are CHALLENGEd, or NEVER?
	Enroll interface{}
	// Elapsed time before the next MFA challenge
	MfaLifetime interface{}
	// Prompt for MFA based on the device used, a factor session lifetime, or every sign on attempt: DEVICE, SESSION or ALWAYS
	MfaPrompt interface{}
	// Remember MFA device.
	MfaRememberDevice interface{}
	// Require MFA.
	MfaRequired interface{}
	// Policy Rule Name
	Name interface{}
	// Network selection mode: ANYWHERE, ZONE, ON_NETWORK, or OFF_NETWORK.
	NetworkConnection interface{}
	// The zones to exclude
	NetworkExcludes interface{}
	// The zones to include
	NetworkIncludes interface{}
	// Allow or deny a user to change their password: ALLOW or DENY. Default = ALLOW
	PasswordChange interface{}
	// Allow or deny a user to reset their password: ALLOW or DENY. Default = ALLOW
	PasswordReset interface{}
	// Allow or deny a user to unlock. Default = DENY
	PasswordUnlock interface{}
	// Policy ID of the Rule
	Policyid interface{}
	// Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an
	// invalid priority is provided. API defaults it to the last/lowest if not there.
	Priority interface{}
	// Max minutes a session can be idle.
	SessionIdle interface{}
	// Max minutes a session is active: Disable = 0.
	SessionLifetime interface{}
	// Whether session cookies will last across browser sessions. Okta Administrators can never have persistent session
	// cookies.
	SessionPersistent interface{}
	// Policy Rule Status: ACTIVE or INACTIVE.
	Status interface{}
	// Set of User IDs to Exclude
	UsersExcludeds interface{}
}

// The set of arguments for constructing a SignonPolicyRule resource.
type SignonPolicyRuleArgs struct {
	// Allow or deny access based on the rule conditions: ALLOW or DENY.
	Access interface{}
	// Authentication entrypoint: ANY or RADIUS.
	Authtype interface{}
	// Should the user be enrolled the first time they LOGIN, the next time they are CHALLENGEd, or NEVER?
	Enroll interface{}
	// Elapsed time before the next MFA challenge
	MfaLifetime interface{}
	// Prompt for MFA based on the device used, a factor session lifetime, or every sign on attempt: DEVICE, SESSION or ALWAYS
	MfaPrompt interface{}
	// Remember MFA device.
	MfaRememberDevice interface{}
	// Require MFA.
	MfaRequired interface{}
	// Policy Rule Name
	Name interface{}
	// Network selection mode: ANYWHERE, ZONE, ON_NETWORK, or OFF_NETWORK.
	NetworkConnection interface{}
	// The zones to exclude
	NetworkExcludes interface{}
	// The zones to include
	NetworkIncludes interface{}
	// Allow or deny a user to change their password: ALLOW or DENY. Default = ALLOW
	PasswordChange interface{}
	// Allow or deny a user to reset their password: ALLOW or DENY. Default = ALLOW
	PasswordReset interface{}
	// Allow or deny a user to unlock. Default = DENY
	PasswordUnlock interface{}
	// Policy ID of the Rule
	Policyid interface{}
	// Policy Rule Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an
	// invalid priority is provided. API defaults it to the last/lowest if not there.
	Priority interface{}
	// Max minutes a session can be idle.
	SessionIdle interface{}
	// Max minutes a session is active: Disable = 0.
	SessionLifetime interface{}
	// Whether session cookies will last across browser sessions. Okta Administrators can never have persistent session
	// cookies.
	SessionPersistent interface{}
	// Policy Rule Status: ACTIVE or INACTIVE.
	Status interface{}
	// Set of User IDs to Exclude
	UsersExcludeds interface{}
}
