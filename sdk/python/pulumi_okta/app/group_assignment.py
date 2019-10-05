# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GroupAssignment(pulumi.CustomResource):
    app_id: pulumi.Output[str]
    """
    The ID of the application to assign a group to.
    """
    group_id: pulumi.Output[str]
    """
    The ID of the group to assign the app to.
    """
    priority: pulumi.Output[float]
    profile: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, app_id=None, group_id=None, priority=None, profile=None, __props__=None, __name__=None, __opts__=None):
        """
        Assigns a group to an application.
        
        This resource allows you to create an App Group assignment.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: The ID of the application to assign a group to.
        :param pulumi.Input[str] group_id: The ID of the group to assign the app to.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-okta/blob/master/website/docs/r/app_group_assignment.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if app_id is None:
                raise TypeError("Missing required property 'app_id'")
            __props__['app_id'] = app_id
            if group_id is None:
                raise TypeError("Missing required property 'group_id'")
            __props__['group_id'] = group_id
            __props__['priority'] = priority
            __props__['profile'] = profile
        super(GroupAssignment, __self__).__init__(
            'okta:app/groupAssignment:GroupAssignment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, app_id=None, group_id=None, priority=None, profile=None):
        """
        Get an existing GroupAssignment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: The ID of the application to assign a group to.
        :param pulumi.Input[str] group_id: The ID of the group to assign the app to.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-okta/blob/master/website/docs/r/app_group_assignment.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["app_id"] = app_id
        __props__["group_id"] = group_id
        __props__["priority"] = priority
        __props__["profile"] = profile
        return GroupAssignment(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

