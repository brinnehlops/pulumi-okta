// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package group

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates an Okta Group Rule.
// 
// This resource allows you to create and configure an Okta Group Rule.
//
// > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/group_rule.html.markdown.
type Rule struct {
	s *pulumi.ResourceState
}

// NewRule registers a new resource with the given unique name, arguments, and options.
func NewRule(ctx *pulumi.Context,
	name string, args *RuleArgs, opts ...pulumi.ResourceOpt) (*Rule, error) {
	if args == nil || args.ExpressionValue == nil {
		return nil, errors.New("missing required argument 'ExpressionValue'")
	}
	if args == nil || args.GroupAssignments == nil {
		return nil, errors.New("missing required argument 'GroupAssignments'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["expressionType"] = nil
		inputs["expressionValue"] = nil
		inputs["groupAssignments"] = nil
		inputs["name"] = nil
		inputs["status"] = nil
	} else {
		inputs["expressionType"] = args.ExpressionType
		inputs["expressionValue"] = args.ExpressionValue
		inputs["groupAssignments"] = args.GroupAssignments
		inputs["name"] = args.Name
		inputs["status"] = args.Status
	}
	s, err := ctx.RegisterResource("okta:group/rule:Rule", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Rule{s: s}, nil
}

// GetRule gets an existing Rule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRule(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RuleState, opts ...pulumi.ResourceOpt) (*Rule, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["expressionType"] = state.ExpressionType
		inputs["expressionValue"] = state.ExpressionValue
		inputs["groupAssignments"] = state.GroupAssignments
		inputs["name"] = state.Name
		inputs["status"] = state.Status
	}
	s, err := ctx.ReadResource("okta:group/rule:Rule", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Rule{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Rule) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Rule) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The expression type to use to invoke the rule. The default is `"urn:okta:expression:1.0"`.
func (r *Rule) ExpressionType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["expressionType"])
}

// The expression value.
func (r *Rule) ExpressionValue() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["expressionValue"])
}

// The list of group ids to assign the users to.
func (r *Rule) GroupAssignments() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["groupAssignments"])
}

// The name of the Group Rule.
func (r *Rule) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The status of the group rule.
func (r *Rule) Status() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["status"])
}

// Input properties used for looking up and filtering Rule resources.
type RuleState struct {
	// The expression type to use to invoke the rule. The default is `"urn:okta:expression:1.0"`.
	ExpressionType interface{}
	// The expression value.
	ExpressionValue interface{}
	// The list of group ids to assign the users to.
	GroupAssignments interface{}
	// The name of the Group Rule.
	Name interface{}
	// The status of the group rule.
	Status interface{}
}

// The set of arguments for constructing a Rule resource.
type RuleArgs struct {
	// The expression type to use to invoke the rule. The default is `"urn:okta:expression:1.0"`.
	ExpressionType interface{}
	// The expression value.
	ExpressionValue interface{}
	// The list of group ids to assign the users to.
	GroupAssignments interface{}
	// The name of the Group Rule.
	Name interface{}
	// The status of the group rule.
	Status interface{}
}
