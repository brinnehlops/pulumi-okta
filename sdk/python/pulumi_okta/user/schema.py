# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Schema']


class Schema(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 array_enums: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 array_one_ofs: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['SchemaArrayOneOfArgs']]]]] = None,
                 array_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enums: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 external_name: Optional[pulumi.Input[str]] = None,
                 index: Optional[pulumi.Input[str]] = None,
                 master: Optional[pulumi.Input[str]] = None,
                 max_length: Optional[pulumi.Input[float]] = None,
                 min_length: Optional[pulumi.Input[float]] = None,
                 one_ofs: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['SchemaOneOfArgs']]]]] = None,
                 permissions: Optional[pulumi.Input[str]] = None,
                 required: Optional[pulumi.Input[bool]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates a User Schema property.

        This resource allows you to create and configure a custom user schema property.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_okta as okta

        example = okta.user.Schema("example",
            description="My custom property name",
            index="customPropertyName",
            master="OKTA",
            scope="SELF",
            title="customPropertyName",
            type="string")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[str]]] array_enums: Array of values that an array property's items can be set to.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['SchemaArrayOneOfArgs']]]] array_one_ofs: Display name and value an enum array can be set to.
        :param pulumi.Input[str] array_type: The type of the array elements if `type` is set to `"array"`.
        :param pulumi.Input[str] description: The description of the user schema property.
        :param pulumi.Input[List[pulumi.Input[str]]] enums: Array of values a primitive property can be set to. See `array_enum` for arrays.
        :param pulumi.Input[str] external_name: External name of the user schema property.
        :param pulumi.Input[str] index: The property name.
        :param pulumi.Input[str] master: Master priority for the user schema property. It can be set to `"PROFILE_MASTER"` or `"OKTA"`.
        :param pulumi.Input[float] max_length: The maximum length of the user property value. Only applies to type `"string"`.
        :param pulumi.Input[float] min_length: The minimum length of the user property value. Only applies to type `"string"`.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['SchemaOneOfArgs']]]] one_ofs: Array of maps containing a mapping for display name to enum value.
        :param pulumi.Input[str] permissions: Access control permissions for the property. It can be set to `"READ_WRITE"`, `"READ_ONLY"`, `"HIDE"`.
        :param pulumi.Input[bool] required: Whether the property is required for this application's users.
        :param pulumi.Input[str] scope: determines whether an app user attribute can be set at the Individual or Group Level.
        :param pulumi.Input[str] title: display name for the enum value.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

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
        super(Schema, __self__).__init__(
            'okta:user/schema:Schema',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            array_enums: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            array_one_ofs: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['SchemaArrayOneOfArgs']]]]] = None,
            array_type: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            enums: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            external_name: Optional[pulumi.Input[str]] = None,
            index: Optional[pulumi.Input[str]] = None,
            master: Optional[pulumi.Input[str]] = None,
            max_length: Optional[pulumi.Input[float]] = None,
            min_length: Optional[pulumi.Input[float]] = None,
            one_ofs: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['SchemaOneOfArgs']]]]] = None,
            permissions: Optional[pulumi.Input[str]] = None,
            required: Optional[pulumi.Input[bool]] = None,
            scope: Optional[pulumi.Input[str]] = None,
            title: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'Schema':
        """
        Get an existing Schema resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[str]]] array_enums: Array of values that an array property's items can be set to.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['SchemaArrayOneOfArgs']]]] array_one_ofs: Display name and value an enum array can be set to.
        :param pulumi.Input[str] array_type: The type of the array elements if `type` is set to `"array"`.
        :param pulumi.Input[str] description: The description of the user schema property.
        :param pulumi.Input[List[pulumi.Input[str]]] enums: Array of values a primitive property can be set to. See `array_enum` for arrays.
        :param pulumi.Input[str] external_name: External name of the user schema property.
        :param pulumi.Input[str] index: The property name.
        :param pulumi.Input[str] master: Master priority for the user schema property. It can be set to `"PROFILE_MASTER"` or `"OKTA"`.
        :param pulumi.Input[float] max_length: The maximum length of the user property value. Only applies to type `"string"`.
        :param pulumi.Input[float] min_length: The minimum length of the user property value. Only applies to type `"string"`.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['SchemaOneOfArgs']]]] one_ofs: Array of maps containing a mapping for display name to enum value.
        :param pulumi.Input[str] permissions: Access control permissions for the property. It can be set to `"READ_WRITE"`, `"READ_ONLY"`, `"HIDE"`.
        :param pulumi.Input[bool] required: Whether the property is required for this application's users.
        :param pulumi.Input[str] scope: determines whether an app user attribute can be set at the Individual or Group Level.
        :param pulumi.Input[str] title: display name for the enum value.
        :param pulumi.Input[str] type: The type of the schema property. It can be `"string"`, `"boolean"`, `"number"`, `"integer"`, `"array"`, or `"object"`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

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
        return Schema(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="arrayEnums")
    def array_enums(self) -> pulumi.Output[Optional[List[str]]]:
        """
        Array of values that an array property's items can be set to.
        """
        return pulumi.get(self, "array_enums")

    @property
    @pulumi.getter(name="arrayOneOfs")
    def array_one_ofs(self) -> pulumi.Output[Optional[List['outputs.SchemaArrayOneOf']]]:
        """
        Display name and value an enum array can be set to.
        """
        return pulumi.get(self, "array_one_ofs")

    @property
    @pulumi.getter(name="arrayType")
    def array_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of the array elements if `type` is set to `"array"`.
        """
        return pulumi.get(self, "array_type")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the user schema property.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def enums(self) -> pulumi.Output[Optional[List[str]]]:
        """
        Array of values a primitive property can be set to. See `array_enum` for arrays.
        """
        return pulumi.get(self, "enums")

    @property
    @pulumi.getter(name="externalName")
    def external_name(self) -> pulumi.Output[Optional[str]]:
        """
        External name of the user schema property.
        """
        return pulumi.get(self, "external_name")

    @property
    @pulumi.getter
    def index(self) -> pulumi.Output[str]:
        """
        The property name.
        """
        return pulumi.get(self, "index")

    @property
    @pulumi.getter
    def master(self) -> pulumi.Output[Optional[str]]:
        """
        Master priority for the user schema property. It can be set to `"PROFILE_MASTER"` or `"OKTA"`.
        """
        return pulumi.get(self, "master")

    @property
    @pulumi.getter(name="maxLength")
    def max_length(self) -> pulumi.Output[Optional[float]]:
        """
        The maximum length of the user property value. Only applies to type `"string"`.
        """
        return pulumi.get(self, "max_length")

    @property
    @pulumi.getter(name="minLength")
    def min_length(self) -> pulumi.Output[Optional[float]]:
        """
        The minimum length of the user property value. Only applies to type `"string"`.
        """
        return pulumi.get(self, "min_length")

    @property
    @pulumi.getter(name="oneOfs")
    def one_ofs(self) -> pulumi.Output[Optional[List['outputs.SchemaOneOf']]]:
        """
        Array of maps containing a mapping for display name to enum value.
        """
        return pulumi.get(self, "one_ofs")

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Output[Optional[str]]:
        """
        Access control permissions for the property. It can be set to `"READ_WRITE"`, `"READ_ONLY"`, `"HIDE"`.
        """
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter
    def required(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the property is required for this application's users.
        """
        return pulumi.get(self, "required")

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Output[Optional[str]]:
        """
        determines whether an app user attribute can be set at the Individual or Group Level.
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter
    def title(self) -> pulumi.Output[str]:
        """
        display name for the enum value.
        """
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the schema property. It can be `"string"`, `"boolean"`, `"number"`, `"integer"`, `"array"`, or `"object"`.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

