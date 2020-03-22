# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class BookmarkApp(pulumi.CustomResource):
    auto_submit_toolbar: pulumi.Output[bool]
    """
    Display auto submit toolbar
    """
    groups: pulumi.Output[list]
    """
    Groups associated with the application
    """
    hide_ios: pulumi.Output[bool]
    """
    Do not display application icon on mobile app
    """
    hide_web: pulumi.Output[bool]
    """
    Do not display application icon to users
    """
    label: pulumi.Output[str]
    """
    Pretty name of app.
    """
    name: pulumi.Output[str]
    """
    name of app.
    """
    request_integration: pulumi.Output[bool]
    sign_on_mode: pulumi.Output[str]
    """
    Sign on mode of application.
    """
    status: pulumi.Output[str]
    """
    Status of application.
    """
    url: pulumi.Output[str]
    users: pulumi.Output[list]
    """
    Users associated with the application

      * `id` (`str`)
      * `password` (`str`)
      * `scope` (`str`)
      * `username` (`str`)
    """
    def __init__(__self__, resource_name, opts=None, auto_submit_toolbar=None, groups=None, hide_ios=None, hide_web=None, label=None, request_integration=None, status=None, url=None, users=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a BookmarkApp resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_submit_toolbar: Display auto submit toolbar
        :param pulumi.Input[list] groups: Groups associated with the application
        :param pulumi.Input[bool] hide_ios: Do not display application icon on mobile app
        :param pulumi.Input[bool] hide_web: Do not display application icon to users
        :param pulumi.Input[str] label: Pretty name of app.
        :param pulumi.Input[str] status: Status of application.
        :param pulumi.Input[list] users: Users associated with the application

        The **users** object supports the following:

          * `id` (`pulumi.Input[str]`)
          * `password` (`pulumi.Input[str]`)
          * `scope` (`pulumi.Input[str]`)
          * `username` (`pulumi.Input[str]`)
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
        super(BookmarkApp, __self__).__init__(
            'okta:deprecated/bookmarkApp:BookmarkApp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, auto_submit_toolbar=None, groups=None, hide_ios=None, hide_web=None, label=None, name=None, request_integration=None, sign_on_mode=None, status=None, url=None, users=None):
        """
        Get an existing BookmarkApp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_submit_toolbar: Display auto submit toolbar
        :param pulumi.Input[list] groups: Groups associated with the application
        :param pulumi.Input[bool] hide_ios: Do not display application icon on mobile app
        :param pulumi.Input[bool] hide_web: Do not display application icon to users
        :param pulumi.Input[str] label: Pretty name of app.
        :param pulumi.Input[str] name: name of app.
        :param pulumi.Input[str] sign_on_mode: Sign on mode of application.
        :param pulumi.Input[str] status: Status of application.
        :param pulumi.Input[list] users: Users associated with the application

        The **users** object supports the following:

          * `id` (`pulumi.Input[str]`)
          * `password` (`pulumi.Input[str]`)
          * `scope` (`pulumi.Input[str]`)
          * `username` (`pulumi.Input[str]`)
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
        return BookmarkApp(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

