// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package profile

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type Mapping struct {
	pulumi.CustomResourceState

	// When turned on this flag will trigger the provider to delete mapping properties that are not defined in config. By
	// default, we do not delete missing properties.
	DeleteWhenAbsent pulumi.BoolPtrOutput `pulumi:"deleteWhenAbsent"`
	Mappings MappingMappingArrayOutput `pulumi:"mappings"`
	// The source id of the mapping to manage.
	SourceId pulumi.StringOutput `pulumi:"sourceId"`
	SourceName pulumi.StringOutput `pulumi:"sourceName"`
	SourceType pulumi.StringOutput `pulumi:"sourceType"`
	// The target id of the mapping to manage.
	TargetId pulumi.StringOutput `pulumi:"targetId"`
	TargetName pulumi.StringOutput `pulumi:"targetName"`
	TargetType pulumi.StringOutput `pulumi:"targetType"`
}

// NewMapping registers a new resource with the given unique name, arguments, and options.
func NewMapping(ctx *pulumi.Context,
	name string, args *MappingArgs, opts ...pulumi.ResourceOption) (*Mapping, error) {
	if args == nil || args.SourceId == nil {
		return nil, errors.New("missing required argument 'SourceId'")
	}
	if args == nil || args.TargetId == nil {
		return nil, errors.New("missing required argument 'TargetId'")
	}
	if args == nil {
		args = &MappingArgs{}
	}
	var resource Mapping
	err := ctx.RegisterResource("okta:profile/mapping:Mapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMapping gets an existing Mapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MappingState, opts ...pulumi.ResourceOption) (*Mapping, error) {
	var resource Mapping
	err := ctx.ReadResource("okta:profile/mapping:Mapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Mapping resources.
type mappingState struct {
	// When turned on this flag will trigger the provider to delete mapping properties that are not defined in config. By
	// default, we do not delete missing properties.
	DeleteWhenAbsent *bool `pulumi:"deleteWhenAbsent"`
	Mappings []MappingMapping `pulumi:"mappings"`
	// The source id of the mapping to manage.
	SourceId *string `pulumi:"sourceId"`
	SourceName *string `pulumi:"sourceName"`
	SourceType *string `pulumi:"sourceType"`
	// The target id of the mapping to manage.
	TargetId *string `pulumi:"targetId"`
	TargetName *string `pulumi:"targetName"`
	TargetType *string `pulumi:"targetType"`
}

type MappingState struct {
	// When turned on this flag will trigger the provider to delete mapping properties that are not defined in config. By
	// default, we do not delete missing properties.
	DeleteWhenAbsent pulumi.BoolPtrInput
	Mappings MappingMappingArrayInput
	// The source id of the mapping to manage.
	SourceId pulumi.StringPtrInput
	SourceName pulumi.StringPtrInput
	SourceType pulumi.StringPtrInput
	// The target id of the mapping to manage.
	TargetId pulumi.StringPtrInput
	TargetName pulumi.StringPtrInput
	TargetType pulumi.StringPtrInput
}

func (MappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*mappingState)(nil)).Elem()
}

type mappingArgs struct {
	// When turned on this flag will trigger the provider to delete mapping properties that are not defined in config. By
	// default, we do not delete missing properties.
	DeleteWhenAbsent *bool `pulumi:"deleteWhenAbsent"`
	Mappings []MappingMapping `pulumi:"mappings"`
	// The source id of the mapping to manage.
	SourceId string `pulumi:"sourceId"`
	// The target id of the mapping to manage.
	TargetId string `pulumi:"targetId"`
}

// The set of arguments for constructing a Mapping resource.
type MappingArgs struct {
	// When turned on this flag will trigger the provider to delete mapping properties that are not defined in config. By
	// default, we do not delete missing properties.
	DeleteWhenAbsent pulumi.BoolPtrInput
	Mappings MappingMappingArrayInput
	// The source id of the mapping to manage.
	SourceId pulumi.StringInput
	// The target id of the mapping to manage.
	TargetId pulumi.StringInput
}

func (MappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mappingArgs)(nil)).Elem()
}

