# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Bookmark(pulumi.CustomResource):
    auto_submit_toolbar: pulumi.Output[bool]
    groups: pulumi.Output[list]
    hide_ios: pulumi.Output[bool]
    hide_web: pulumi.Output[bool]
    label: pulumi.Output[str]
    """
    The Application's display name.
    """
    name: pulumi.Output[str]
    request_integration: pulumi.Output[bool]
    """
    Would you like Okta to add an integration for this app?
    """
    sign_on_mode: pulumi.Output[str]
    status: pulumi.Output[str]
    url: pulumi.Output[str]
    """
    The URL of the bookmark.
    """
    users: pulumi.Output[list]
    def __init__(__self__, resource_name, opts=None, auto_submit_toolbar=None, groups=None, hide_ios=None, hide_web=None, label=None, request_integration=None, status=None, url=None, users=None, __props__=None, __name__=None, __opts__=None):
        """
        Creates a Bookmark Application.
        
        This resource allows you to create and configure a Bookmark Application.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] label: The Application's display name.
        :param pulumi.Input[bool] request_integration: Would you like Okta to add an integration for this app?
        :param pulumi.Input[str] url: The URL of the bookmark.
        
        The **users** object supports the following:
        
          * `id` (`pulumi.Input[str]`) - ID of the Application.
          * `password` (`pulumi.Input[str]`)
          * `scope` (`pulumi.Input[str]`)
          * `username` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-okta/blob/master/website/docs/r/app_bookmark.html.markdown.
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

            __props__['auto_submit_toolbar'] = auto_submit_toolbar
            __props__['groups'] = groups
            __props__['hide_ios'] = hide_ios
            __props__['hide_web'] = hide_web
            if label is None:
                raise TypeError("Missing required property 'label'")
            __props__['label'] = label
            __props__['request_integration'] = request_integration
            __props__['status'] = status
            if url is None:
                raise TypeError("Missing required property 'url'")
            __props__['url'] = url
            __props__['users'] = users
            __props__['name'] = None
            __props__['sign_on_mode'] = None
        super(Bookmark, __self__).__init__(
            'okta:app/bookmark:Bookmark',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, auto_submit_toolbar=None, groups=None, hide_ios=None, hide_web=None, label=None, name=None, request_integration=None, sign_on_mode=None, status=None, url=None, users=None):
        """
        Get an existing Bookmark resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] label: The Application's display name.
        :param pulumi.Input[bool] request_integration: Would you like Okta to add an integration for this app?
        :param pulumi.Input[str] url: The URL of the bookmark.
        
        The **users** object supports the following:
        
          * `id` (`pulumi.Input[str]`) - ID of the Application.
          * `password` (`pulumi.Input[str]`)
          * `scope` (`pulumi.Input[str]`)
          * `username` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-okta/blob/master/website/docs/r/app_bookmark.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["auto_submit_toolbar"] = auto_submit_toolbar
        __props__["groups"] = groups
        __props__["hide_ios"] = hide_ios
        __props__["hide_web"] = hide_web
        __props__["label"] = label
        __props__["name"] = name
        __props__["request_integration"] = request_integration
        __props__["sign_on_mode"] = sign_on_mode
        __props__["status"] = status
        __props__["url"] = url
        __props__["users"] = users
        return Bookmark(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

