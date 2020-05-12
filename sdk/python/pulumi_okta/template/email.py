# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Email(pulumi.CustomResource):
    default_language: pulumi.Output[str]
    """
    The default language, by default is set to `"en"`.
    """
    translations: pulumi.Output[list]
    """
    Set of translations for particular template.

      * `language` (`str`) - The language to map tthe template to.
      * `subject` (`str`) - The email subject line.
      * `template` (`str`) - The email body.
    """
    type: pulumi.Output[str]
    """
    Email template type
    """
    def __init__(__self__, resource_name, opts=None, default_language=None, translations=None, type=None, __props__=None, __name__=None, __opts__=None):
        """
        Creates an Okta Email Template.

        This resource allows you to create and configure an Okta Email Template.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_okta as okta

        example = okta.template.Email("example",
            translations=[
                {
                    "language": "en",
                    "subject": "Stuff",
                    "template": f"Hi {user['firstName']},<br/><br/>Blah blah {reset_password_link}",
                },
                {
                    "language": "es",
                    "subject": "Cosas",
                    "template": f"Hola {user['firstName']},<br/><br/>Puedo ir al bano {reset_password_link}",
                },
            ],
            type="email.forgotPassword")
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_language: The default language, by default is set to `"en"`.
        :param pulumi.Input[list] translations: Set of translations for particular template.
        :param pulumi.Input[str] type: Email template type

        The **translations** object supports the following:

          * `language` (`pulumi.Input[str]`) - The language to map tthe template to.
          * `subject` (`pulumi.Input[str]`) - The email subject line.
          * `template` (`pulumi.Input[str]`) - The email body.
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

            __props__['default_language'] = default_language
            if translations is None:
                raise TypeError("Missing required property 'translations'")
            __props__['translations'] = translations
            if type is None:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
        super(Email, __self__).__init__(
            'okta:template/email:Email',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, default_language=None, translations=None, type=None):
        """
        Get an existing Email resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_language: The default language, by default is set to `"en"`.
        :param pulumi.Input[list] translations: Set of translations for particular template.
        :param pulumi.Input[str] type: Email template type

        The **translations** object supports the following:

          * `language` (`pulumi.Input[str]`) - The language to map tthe template to.
          * `subject` (`pulumi.Input[str]`) - The email subject line.
          * `template` (`pulumi.Input[str]`) - The email body.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["default_language"] = default_language
        __props__["translations"] = translations
        __props__["type"] = type
        return Email(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

