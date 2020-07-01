# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Origin(pulumi.CustomResource):
    active: pulumi.Output[bool]
    """
    Whether the Trusted Origin is active or not - can only be issued post-creation.
    """
    name: pulumi.Output[str]
    """
    Name of the Trusted Origin Resource.
    """
    origin: pulumi.Output[str]
    """
    The origin to trust.
    """
    scopes: pulumi.Output[list]
    """
    Scopes of the Trusted Origin - can be `"CORS"` and/or `"REDIRECT"`.
    """
    def __init__(__self__, resource_name, opts=None, active=None, name=None, origin=None, scopes=None, __props__=None, __name__=None, __opts__=None):
        """
        Creates a Trusted Origin.

        This resource allows you to create and configure an Trusted Origin.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_okta as okta

        example = okta.trustedorigin.Origin("example",
            origin="https://example.com",
            scopes=["CORS"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] active: Whether the Trusted Origin is active or not - can only be issued post-creation.
        :param pulumi.Input[str] name: Name of the Trusted Origin Resource.
        :param pulumi.Input[str] origin: The origin to trust.
        :param pulumi.Input[list] scopes: Scopes of the Trusted Origin - can be `"CORS"` and/or `"REDIRECT"`.
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

            __props__['active'] = active
            __props__['name'] = name
            if origin is None:
                raise TypeError("Missing required property 'origin'")
            __props__['origin'] = origin
            if scopes is None:
                raise TypeError("Missing required property 'scopes'")
            __props__['scopes'] = scopes
        super(Origin, __self__).__init__(
            'okta:trustedorigin/origin:Origin',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, active=None, name=None, origin=None, scopes=None):
        """
        Get an existing Origin resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] active: Whether the Trusted Origin is active or not - can only be issued post-creation.
        :param pulumi.Input[str] name: Name of the Trusted Origin Resource.
        :param pulumi.Input[str] origin: The origin to trust.
        :param pulumi.Input[list] scopes: Scopes of the Trusted Origin - can be `"CORS"` and/or `"REDIRECT"`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["active"] = active
        __props__["name"] = name
        __props__["origin"] = origin
        __props__["scopes"] = scopes
        return Origin(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
