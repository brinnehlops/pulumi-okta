# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class UserBaseSchema(pulumi.CustomResource):
    app_id: pulumi.Output[str]
    """
    The Application's ID the user schema property should be assigned to.
    """
    index: pulumi.Output[str]
    """
    The property name.
    """
    master: pulumi.Output[str]
    """
    Master priority for the user schema property. It can be set to `"PROFILE_MASTER"` or `"OKTA"`.
    """
    permissions: pulumi.Output[str]
    """
    Access control permissions for the property. It can be set to `"READ_WRITE"`, `"READ_ONLY"`, `"HIDE"`.
    """
    required: pulumi.Output[bool]
    """
    Whether the property is required for this application's users.
    """
    title: pulumi.Output[str]
    """
    The property display name.
    """
    type: pulumi.Output[str]
    """
    The type of the schema property. It can be `"string"`, `"boolean"`, `"number"`, `"integer"`, `"array"`, or `"object"`.
    """
    def __init__(__self__, resource_name, opts=None, app_id=None, index=None, master=None, permissions=None, required=None, title=None, type=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an Application User Base Schema property.

        This resource allows you to configure a base app user schema property.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_okta as okta

        example = okta.app.UserBaseSchema("example",
            app_id="<app id>",
            index="customPropertyName",
            master="OKTA",
            title="customPropertyName",
            type="string")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: The Application's ID the user schema property should be assigned to.
        :param pulumi.Input[str] index: The property name.
        :param pulumi.Input[str] master: Master priority for the user schema property. It can be set to `"PROFILE_MASTER"` or `"OKTA"`.
        :param pulumi.Input[str] permissions: Access control permissions for the property. It can be set to `"READ_WRITE"`, `"READ_ONLY"`, `"HIDE"`.
        :param pulumi.Input[bool] required: Whether the property is required for this application's users.
        :param pulumi.Input[str] title: The property display name.
        :param pulumi.Input[str] type: The type of the schema property. It can be `"string"`, `"boolean"`, `"number"`, `"integer"`, `"array"`, or `"object"`.
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
            if index is None:
                raise TypeError("Missing required property 'index'")
            __props__['index'] = index
            __props__['master'] = master
            __props__['permissions'] = permissions
            __props__['required'] = required
            if title is None:
                raise TypeError("Missing required property 'title'")
            __props__['title'] = title
            if type is None:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
        super(UserBaseSchema, __self__).__init__(
            'okta:app/userBaseSchema:UserBaseSchema',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, app_id=None, index=None, master=None, permissions=None, required=None, title=None, type=None):
        """
        Get an existing UserBaseSchema resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: The Application's ID the user schema property should be assigned to.
        :param pulumi.Input[str] index: The property name.
        :param pulumi.Input[str] master: Master priority for the user schema property. It can be set to `"PROFILE_MASTER"` or `"OKTA"`.
        :param pulumi.Input[str] permissions: Access control permissions for the property. It can be set to `"READ_WRITE"`, `"READ_ONLY"`, `"HIDE"`.
        :param pulumi.Input[bool] required: Whether the property is required for this application's users.
        :param pulumi.Input[str] title: The property display name.
        :param pulumi.Input[str] type: The type of the schema property. It can be `"string"`, `"boolean"`, `"number"`, `"integer"`, `"array"`, or `"object"`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["app_id"] = app_id
        __props__["index"] = index
        __props__["master"] = master
        __props__["permissions"] = permissions
        __props__["required"] = required
        __props__["title"] = title
        __props__["type"] = type
        return UserBaseSchema(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
