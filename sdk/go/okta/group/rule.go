// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package group

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates an Okta Group Rule.
// 
// This resource allows you to create and configure an Okta Group Rule.
// 
// > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/group_rule.html.markdown.
type Rule struct {
	pulumi.CustomResourceState

	// The expression type to use to invoke the rule. The default is `"urn:okta:expression:1.0"`.
	ExpressionType pulumi.StringPtrOutput `pulumi:"expressionType"`
	// The expression value.
	ExpressionValue pulumi.StringOutput `pulumi:"expressionValue"`
	// The list of group ids to assign the users to.
	GroupAssignments pulumi.StringArrayOutput `pulumi:"groupAssignments"`
	// The name of the Group Rule.
	Name pulumi.StringOutput `pulumi:"name"`
	// The status of the group rule.
	Status pulumi.StringPtrOutput `pulumi:"status"`
}

// NewRule registers a new resource with the given unique name, arguments, and options.
func NewRule(ctx *pulumi.Context,
	name string, args *RuleArgs, opts ...pulumi.ResourceOption) (*Rule, error) {
	if args == nil || args.ExpressionValue == nil {
		return nil, errors.New("missing required argument 'ExpressionValue'")
	}
	if args == nil || args.GroupAssignments == nil {
		return nil, errors.New("missing required argument 'GroupAssignments'")
	}
	if args == nil {
		args = &RuleArgs{}
	}
	var resource Rule
	err := ctx.RegisterResource("okta:group/rule:Rule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRule gets an existing Rule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleState, opts ...pulumi.ResourceOption) (*Rule, error) {
	var resource Rule
	err := ctx.ReadResource("okta:group/rule:Rule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rule resources.
type ruleState struct {
	// The expression type to use to invoke the rule. The default is `"urn:okta:expression:1.0"`.
	ExpressionType *string `pulumi:"expressionType"`
	// The expression value.
	ExpressionValue *string `pulumi:"expressionValue"`
	// The list of group ids to assign the users to.
	GroupAssignments []string `pulumi:"groupAssignments"`
	// The name of the Group Rule.
	Name *string `pulumi:"name"`
	// The status of the group rule.
	Status *string `pulumi:"status"`
}

type RuleState struct {
	// The expression type to use to invoke the rule. The default is `"urn:okta:expression:1.0"`.
	ExpressionType pulumi.StringPtrInput
	// The expression value.
	ExpressionValue pulumi.StringPtrInput
	// The list of group ids to assign the users to.
	GroupAssignments pulumi.StringArrayInput
	// The name of the Group Rule.
	Name pulumi.StringPtrInput
	// The status of the group rule.
	Status pulumi.StringPtrInput
}

func (RuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleState)(nil)).Elem()
}

type ruleArgs struct {
	// The expression type to use to invoke the rule. The default is `"urn:okta:expression:1.0"`.
	ExpressionType *string `pulumi:"expressionType"`
	// The expression value.
	ExpressionValue string `pulumi:"expressionValue"`
	// The list of group ids to assign the users to.
	GroupAssignments []string `pulumi:"groupAssignments"`
	// The name of the Group Rule.
	Name *string `pulumi:"name"`
	// The status of the group rule.
	Status *string `pulumi:"status"`
}

// The set of arguments for constructing a Rule resource.
type RuleArgs struct {
	// The expression type to use to invoke the rule. The default is `"urn:okta:expression:1.0"`.
	ExpressionType pulumi.StringPtrInput
	// The expression value.
	ExpressionValue pulumi.StringInput
	// The list of group ids to assign the users to.
	GroupAssignments pulumi.StringArrayInput
	// The name of the Group Rule.
	Name pulumi.StringPtrInput
	// The status of the group rule.
	Status pulumi.StringPtrInput
}

func (RuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleArgs)(nil)).Elem()
}

