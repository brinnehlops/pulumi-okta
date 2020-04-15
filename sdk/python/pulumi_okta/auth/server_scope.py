# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ServerScope(pulumi.CustomResource):
    auth_server_id: pulumi.Output[str]
    """
    Auth Server ID.
    """
    consent: pulumi.Output[str]
    """
    Indicates whether a consent dialog is needed for the scope. It can be set to `"REQUIRED"` or `"IMPLICIT"`.
    """
    default: pulumi.Output[bool]
    """
    A default scope will be returned in an access token when the client omits the scope parameter in a token request, provided this scope is allowed as part of the access policy rule.
    """
    description: pulumi.Output[str]
    """
    Description of the Auth Server Scope.
    """
    metadata_publish: pulumi.Output[str]
    """
    Whether to publish metadata or not. It can be set to `"ALL_CLIENTS"` or `"NO_CLIENTS"`.
    """
    name: pulumi.Output[str]
    """
    Auth Server scope name.
    """
    def __init__(__self__, resource_name, opts=None, auth_server_id=None, consent=None, default=None, description=None, metadata_publish=None, name=None, __props__=None, __name__=None, __opts__=None):
        """
        Creates an Authorization Server Scope.

        This resource allows you to create and configure an Authorization Server Scope.



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_server_id: Auth Server ID.
        :param pulumi.Input[str] consent: Indicates whether a consent dialog is needed for the scope. It can be set to `"REQUIRED"` or `"IMPLICIT"`.
        :param pulumi.Input[bool] default: A default scope will be returned in an access token when the client omits the scope parameter in a token request, provided this scope is allowed as part of the access policy rule.
        :param pulumi.Input[str] description: Description of the Auth Server Scope.
        :param pulumi.Input[str] metadata_publish: Whether to publish metadata or not. It can be set to `"ALL_CLIENTS"` or `"NO_CLIENTS"`.
        :param pulumi.Input[str] name: Auth Server scope name.
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

            if auth_server_id is None:
                raise TypeError("Missing required property 'auth_server_id'")
            __props__['auth_server_id'] = auth_server_id
            __props__['consent'] = consent
            __props__['default'] = default
            __props__['description'] = description
            __props__['metadata_publish'] = metadata_publish
            __props__['name'] = name
        super(ServerScope, __self__).__init__(
            'okta:auth/serverScope:ServerScope',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, auth_server_id=None, consent=None, default=None, description=None, metadata_publish=None, name=None):
        """
        Get an existing ServerScope resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_server_id: Auth Server ID.
        :param pulumi.Input[str] consent: Indicates whether a consent dialog is needed for the scope. It can be set to `"REQUIRED"` or `"IMPLICIT"`.
        :param pulumi.Input[bool] default: A default scope will be returned in an access token when the client omits the scope parameter in a token request, provided this scope is allowed as part of the access policy rule.
        :param pulumi.Input[str] description: Description of the Auth Server Scope.
        :param pulumi.Input[str] metadata_publish: Whether to publish metadata or not. It can be set to `"ALL_CLIENTS"` or `"NO_CLIENTS"`.
        :param pulumi.Input[str] name: Auth Server scope name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["auth_server_id"] = auth_server_id
        __props__["consent"] = consent
        __props__["default"] = default
        __props__["description"] = description
        __props__["metadata_publish"] = metadata_publish
        __props__["name"] = name
        return ServerScope(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

