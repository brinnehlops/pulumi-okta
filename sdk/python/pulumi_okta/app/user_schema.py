# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class UserSchema(pulumi.CustomResource):
    app_id: pulumi.Output[str]
    """
    The Application's ID the user custom schema property should be assigned to.
    """
    array_enums: pulumi.Output[list]
    """
    Array of values that an array property's items can be set to.
    """
    array_one_ofs: pulumi.Output[list]
    """
    Display name and value an enum array can be set to.

      * `const` (`str`) - value mapping to member of `enum`.
      * `title` (`str`) - display name for the enum value.
    """
    array_type: pulumi.Output[str]
    """
    The type of the array elements if `type` is set to `"array"`.
    """
    description: pulumi.Output[str]
    """
    The description of the user schema property.
    """
    enums: pulumi.Output[list]
    """
    Array of values a primitive property can be set to. See `array_enum` for arrays.
    """
    external_name: pulumi.Output[str]
    """
    External name of the user schema property.
    """
    index: pulumi.Output[str]
    """
    The property name.
    """
    master: pulumi.Output[str]
    """
    Master priority for the user schema property. It can be set to `"PROFILE_MASTER"` or `"OKTA"`.
    """
    max_length: pulumi.Output[float]
    """
    The maximum length of the user property value. Only applies to type `"string"`.
    """
    min_length: pulumi.Output[float]
    """
    The minimum length of the user property value. Only applies to type `"string"`.
    """
    one_ofs: pulumi.Output[list]
    """
    Array of maps containing a mapping for display name to enum value.

      * `const` (`str`) - value mapping to member of `enum`.
      * `title` (`str`) - display name for the enum value.
    """
    permissions: pulumi.Output[str]
    """
    Access control permissions for the property. It can be set to `"READ_WRITE"`, `"READ_ONLY"`, `"HIDE"`.
    """
    required: pulumi.Output[bool]
    """
    Whether the property is required for this application's users.
    """
    scope: pulumi.Output[str]
    """
    determines whether an app user attribute can be set at the Individual or Group Level.
    """
    title: pulumi.Output[str]
    """
    display name for the enum value.
    """
    type: pulumi.Output[str]
    """
    The type of the schema property. It can be `"string"`, `"boolean"`, `"number"`, `"integer"`, `"array"`, or `"object"`.
    """
    def __init__(__self__, resource_name, opts=None, app_id=None, array_enums=None, array_one_ofs=None, array_type=None, description=None, enums=None, external_name=None, index=None, master=None, max_length=None, min_length=None, one_ofs=None, permissions=None, required=None, scope=None, title=None, type=None, __props__=None, __name__=None, __opts__=None):
        """
        Creates an Application User Schema property.

        This resource allows you to create and configure a custom user schema property and associate it with an application.



        > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/app_user_schema.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: The Application's ID the user custom schema property should be assigned to.
        :param pulumi.Input[list] array_enums: Array of values that an array property's items can be set to.
        :param pulumi.Input[list] array_one_ofs: Display name and value an enum array can be set to.
        :param pulumi.Input[str] array_type: The type of the array elements if `type` is set to `"array"`.
        :param pulumi.Input[str] description: The description of the user schema property.
        :param pulumi.Input[list] enums: Array of values a primitive property can be set to. See `array_enum` for arrays.
        :param pulumi.Input[str] external_name: External name of the user schema property.
        :param pulumi.Input[str] index: The property name.
        :param pulumi.Input[str] master: Master priority for the user schema property. It can be set to `"PROFILE_MASTER"` or `"OKTA"`.
        :param pulumi.Input[float] max_length: The maximum length of the user property value. Only applies to type `"string"`.
        :param pulumi.Input[float] min_length: The minimum length of the user property value. Only applies to type `"string"`.
        :param pulumi.Input[list] one_ofs: Array of maps containing a mapping for display name to enum value.
        :param pulumi.Input[str] permissions: Access control permissions for the property. It can be set to `"READ_WRITE"`, `"READ_ONLY"`, `"HIDE"`.
        :param pulumi.Input[bool] required: Whether the property is required for this application's users.
        :param pulumi.Input[str] scope: determines whether an app user attribute can be set at the Individual or Group Level.
        :param pulumi.Input[str] title: display name for the enum value.
        :param pulumi.Input[str] type: The type of the schema property. It can be `"string"`, `"boolean"`, `"number"`, `"integer"`, `"array"`, or `"object"`.

        The **array_one_ofs** object supports the following:

          * `const` (`pulumi.Input[str]`) - value mapping to member of `enum`.
          * `title` (`pulumi.Input[str]`) - display name for the enum value.

        The **one_ofs** object supports the following:

          * `const` (`pulumi.Input[str]`) - value mapping to member of `enum`.
          * `title` (`pulumi.Input[str]`) - display name for the enum value.
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
            __props__['array_enums'] = array_enums
            __props__['array_one_ofs'] = array_one_ofs
            __props__['array_type'] = array_type
            __props__['description'] = description
            __props__['enums'] = enums
            __props__['external_name'] = external_name
            if index is None:
                raise TypeError("Missing required property 'index'")
            __props__['index'] = index
            __props__['master'] = master
            __props__['max_length'] = max_length
            __props__['min_length'] = min_length
            __props__['one_ofs'] = one_ofs
            __props__['permissions'] = permissions
            __props__['required'] = required
            __props__['scope'] = scope
            if title is None:
                raise TypeError("Missing required property 'title'")
            __props__['title'] = title
            if type is None:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
        super(UserSchema, __self__).__init__(
            'okta:app/userSchema:UserSchema',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, app_id=None, array_enums=None, array_one_ofs=None, array_type=None, description=None, enums=None, external_name=None, index=None, master=None, max_length=None, min_length=None, one_ofs=None, permissions=None, required=None, scope=None, title=None, type=None):
        """
        Get an existing UserSchema resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: The Application's ID the user custom schema property should be assigned to.
        :param pulumi.Input[list] array_enums: Array of values that an array property's items can be set to.
        :param pulumi.Input[list] array_one_ofs: Display name and value an enum array can be set to.
        :param pulumi.Input[str] array_type: The type of the array elements if `type` is set to `"array"`.
        :param pulumi.Input[str] description: The description of the user schema property.
        :param pulumi.Input[list] enums: Array of values a primitive property can be set to. See `array_enum` for arrays.
        :param pulumi.Input[str] external_name: External name of the user schema property.
        :param pulumi.Input[str] index: The property name.
        :param pulumi.Input[str] master: Master priority for the user schema property. It can be set to `"PROFILE_MASTER"` or `"OKTA"`.
        :param pulumi.Input[float] max_length: The maximum length of the user property value. Only applies to type `"string"`.
        :param pulumi.Input[float] min_length: The minimum length of the user property value. Only applies to type `"string"`.
        :param pulumi.Input[list] one_ofs: Array of maps containing a mapping for display name to enum value.
        :param pulumi.Input[str] permissions: Access control permissions for the property. It can be set to `"READ_WRITE"`, `"READ_ONLY"`, `"HIDE"`.
        :param pulumi.Input[bool] required: Whether the property is required for this application's users.
        :param pulumi.Input[str] scope: determines whether an app user attribute can be set at the Individual or Group Level.
        :param pulumi.Input[str] title: display name for the enum value.
        :param pulumi.Input[str] type: The type of the schema property. It can be `"string"`, `"boolean"`, `"number"`, `"integer"`, `"array"`, or `"object"`.

        The **array_one_ofs** object supports the following:

          * `const` (`pulumi.Input[str]`) - value mapping to member of `enum`.
          * `title` (`pulumi.Input[str]`) - display name for the enum value.

        The **one_ofs** object supports the following:

          * `const` (`pulumi.Input[str]`) - value mapping to member of `enum`.
          * `title` (`pulumi.Input[str]`) - display name for the enum value.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["app_id"] = app_id
        __props__["array_enums"] = array_enums
        __props__["array_one_ofs"] = array_one_ofs
        __props__["array_type"] = array_type
        __props__["description"] = description
        __props__["enums"] = enums
        __props__["external_name"] = external_name
        __props__["index"] = index
        __props__["master"] = master
        __props__["max_length"] = max_length
        __props__["min_length"] = min_length
        __props__["one_ofs"] = one_ofs
        __props__["permissions"] = permissions
        __props__["required"] = required
        __props__["scope"] = scope
        __props__["title"] = title
        __props__["type"] = type
        return UserSchema(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

