# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class User(pulumi.CustomResource):
    app_id: pulumi.Output[str]
    """
    App to associate user with.
    """
    password: pulumi.Output[str]
    """
    The password to use.
    """
    profile: pulumi.Output[str]
    """
    The JSON profile of the App User.
    """
    user_id: pulumi.Output[str]
    """
    User to associate the application with.
    """
    username: pulumi.Output[str]
    """
    The username to use for the app user.
    """
    def __init__(__self__, resource_name, opts=None, app_id=None, password=None, profile=None, user_id=None, username=None, __props__=None, __name__=None, __opts__=None):
        """
        Creates an Application User.

        This resource allows you to create and configure an Application User.

        __When using this resource, make sure to add the following `lifefycle` argument to the application resource you are assigning to:__

        ```python
        import pulumi
        ```

        ## Example Usage

        ```python
        import pulumi
        import pulumi_okta as okta

        example = okta.app.User("example",
            app_id="<app_id>",
            user_id="<user id>",
            username="example")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: App to associate user with.
        :param pulumi.Input[str] password: The password to use.
        :param pulumi.Input[str] profile: The JSON profile of the App User.
        :param pulumi.Input[str] user_id: User to associate the application with.
        :param pulumi.Input[str] username: The username to use for the app user.
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
            __props__['password'] = password
            __props__['profile'] = profile
            if user_id is None:
                raise TypeError("Missing required property 'user_id'")
            __props__['user_id'] = user_id
            if username is None:
                raise TypeError("Missing required property 'username'")
            __props__['username'] = username
        super(User, __self__).__init__(
            'okta:app/user:User',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, app_id=None, password=None, profile=None, user_id=None, username=None):
        """
        Get an existing User resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: App to associate user with.
        :param pulumi.Input[str] password: The password to use.
        :param pulumi.Input[str] profile: The JSON profile of the App User.
        :param pulumi.Input[str] user_id: User to associate the application with.
        :param pulumi.Input[str] username: The username to use for the app user.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["app_id"] = app_id
        __props__["password"] = password
        __props__["profile"] = profile
        __props__["user_id"] = user_id
        __props__["username"] = username
        return User(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
