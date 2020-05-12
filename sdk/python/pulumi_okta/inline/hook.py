# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Hook(pulumi.CustomResource):
    auth: pulumi.Output[dict]
    """
    Authentication required for inline hook request.

      * `key` (`str`) - Key to use for authentication, usually the header name, for example `"Authorization"`.
      * `type` (`str`) - The type of hook to trigger. Currently only `"HTTP"` is supported.
      * `value` (`str`) - Authentication secret.
    """
    channel: pulumi.Output[dict]
    """
    Details of the endpoint the inline hook will hit.

      * `method` (`str`) - The request method to use. Default is `"POST"`.
      * `type` (`bool`) - The type of hook to trigger. Currently only `"HTTP"` is supported.
      * `uri` (`str`) - The URI the hook will hit.
      * `version` (`str`) - The version of the endpoint.
    """
    headers: pulumi.Output[list]
    """
    Map of headers to send along in inline hook request.

      * `key` (`str`) - Key to use for authentication, usually the header name, for example `"Authorization"`.
      * `value` (`str`) - Authentication secret.
    """
    name: pulumi.Output[str]
    """
    The inline hook display name.
    """
    status: pulumi.Output[str]
    type: pulumi.Output[str]
    """
    The type of hook to trigger. Currently only `"HTTP"` is supported.
    """
    version: pulumi.Output[str]
    """
    The version of the endpoint.
    """
    def __init__(__self__, resource_name, opts=None, auth=None, channel=None, headers=None, name=None, status=None, type=None, version=None, __props__=None, __name__=None, __opts__=None):
        """
        Creates an inline hook.

        This resource allows you to create and configure an inline hook.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_okta as okta

        example = okta.inline.Hook("example",
            auth={
                "key": "Authorization",
                "type": "HEADER",
                "value": "secret",
            },
            channel={
                "method": "POST",
                "uri": "https://example.com/test",
                "version": "1.0.0",
            },
            type="com.okta.oauth2.tokens.transform",
            version="1.0.1")
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] auth: Authentication required for inline hook request.
        :param pulumi.Input[dict] channel: Details of the endpoint the inline hook will hit.
        :param pulumi.Input[list] headers: Map of headers to send along in inline hook request.
        :param pulumi.Input[str] name: The inline hook display name.
        :param pulumi.Input[str] type: The type of hook to trigger. Currently only `"HTTP"` is supported.
        :param pulumi.Input[str] version: The version of the endpoint.

        The **auth** object supports the following:

          * `key` (`pulumi.Input[str]`) - Key to use for authentication, usually the header name, for example `"Authorization"`.
          * `type` (`pulumi.Input[str]`) - The type of hook to trigger. Currently only `"HTTP"` is supported.
          * `value` (`pulumi.Input[str]`) - Authentication secret.

        The **channel** object supports the following:

          * `method` (`pulumi.Input[str]`) - The request method to use. Default is `"POST"`.
          * `type` (`pulumi.Input[bool]`) - The type of hook to trigger. Currently only `"HTTP"` is supported.
          * `uri` (`pulumi.Input[str]`) - The URI the hook will hit.
          * `version` (`pulumi.Input[str]`) - The version of the endpoint.

        The **headers** object supports the following:

          * `key` (`pulumi.Input[str]`) - Key to use for authentication, usually the header name, for example `"Authorization"`.
          * `value` (`pulumi.Input[str]`) - Authentication secret.
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

            __props__['auth'] = auth
            __props__['channel'] = channel
            __props__['headers'] = headers
            __props__['name'] = name
            __props__['status'] = status
            if type is None:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
            if version is None:
                raise TypeError("Missing required property 'version'")
            __props__['version'] = version
        super(Hook, __self__).__init__(
            'okta:inline/hook:Hook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, auth=None, channel=None, headers=None, name=None, status=None, type=None, version=None):
        """
        Get an existing Hook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] auth: Authentication required for inline hook request.
        :param pulumi.Input[dict] channel: Details of the endpoint the inline hook will hit.
        :param pulumi.Input[list] headers: Map of headers to send along in inline hook request.
        :param pulumi.Input[str] name: The inline hook display name.
        :param pulumi.Input[str] type: The type of hook to trigger. Currently only `"HTTP"` is supported.
        :param pulumi.Input[str] version: The version of the endpoint.

        The **auth** object supports the following:

          * `key` (`pulumi.Input[str]`) - Key to use for authentication, usually the header name, for example `"Authorization"`.
          * `type` (`pulumi.Input[str]`) - The type of hook to trigger. Currently only `"HTTP"` is supported.
          * `value` (`pulumi.Input[str]`) - Authentication secret.

        The **channel** object supports the following:

          * `method` (`pulumi.Input[str]`) - The request method to use. Default is `"POST"`.
          * `type` (`pulumi.Input[bool]`) - The type of hook to trigger. Currently only `"HTTP"` is supported.
          * `uri` (`pulumi.Input[str]`) - The URI the hook will hit.
          * `version` (`pulumi.Input[str]`) - The version of the endpoint.

        The **headers** object supports the following:

          * `key` (`pulumi.Input[str]`) - Key to use for authentication, usually the header name, for example `"Authorization"`.
          * `value` (`pulumi.Input[str]`) - Authentication secret.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["auth"] = auth
        __props__["channel"] = channel
        __props__["headers"] = headers
        __props__["name"] = name
        __props__["status"] = status
        __props__["type"] = type
        __props__["version"] = version
        return Hook(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

