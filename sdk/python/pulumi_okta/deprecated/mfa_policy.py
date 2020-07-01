# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class MfaPolicy(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    Policy Description
    """
    duo: pulumi.Output[dict]
    fido_u2f: pulumi.Output[dict]
    fido_webauthn: pulumi.Output[dict]
    google_otp: pulumi.Output[dict]
    groups_includeds: pulumi.Output[list]
    """
    List of Group IDs to Include
    """
    name: pulumi.Output[str]
    """
    Policy Name
    """
    okta_call: pulumi.Output[dict]
    okta_otp: pulumi.Output[dict]
    okta_password: pulumi.Output[dict]
    okta_push: pulumi.Output[dict]
    okta_question: pulumi.Output[dict]
    okta_sms: pulumi.Output[dict]
    priority: pulumi.Output[float]
    """
    Policy Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid
    priority is provided. API defaults it to the last/lowest if not there.
    """
    rsa_token: pulumi.Output[dict]
    status: pulumi.Output[str]
    """
    Policy Status: ACTIVE or INACTIVE.
    """
    symantec_vip: pulumi.Output[dict]
    yubikey_token: pulumi.Output[dict]
    def __init__(__self__, resource_name, opts=None, description=None, duo=None, fido_u2f=None, fido_webauthn=None, google_otp=None, groups_includeds=None, name=None, okta_call=None, okta_otp=None, okta_password=None, okta_push=None, okta_question=None, okta_sms=None, priority=None, rsa_token=None, status=None, symantec_vip=None, yubikey_token=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a MfaPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Policy Description
        :param pulumi.Input[list] groups_includeds: List of Group IDs to Include
        :param pulumi.Input[str] name: Policy Name
        :param pulumi.Input[float] priority: Policy Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid
               priority is provided. API defaults it to the last/lowest if not there.
        :param pulumi.Input[str] status: Policy Status: ACTIVE or INACTIVE.

        The **duo** object supports the following:

          * `consent_type` (`pulumi.Input[str]`)
          * `enroll` (`pulumi.Input[str]`)

        The **fido_u2f** object supports the following:

          * `consent_type` (`pulumi.Input[str]`)
          * `enroll` (`pulumi.Input[str]`)

        The **fido_webauthn** object supports the following:

          * `consent_type` (`pulumi.Input[str]`)
          * `enroll` (`pulumi.Input[str]`)

        The **google_otp** object supports the following:

          * `consent_type` (`pulumi.Input[str]`)
          * `enroll` (`pulumi.Input[str]`)

        The **okta_call** object supports the following:

          * `consent_type` (`pulumi.Input[str]`)
          * `enroll` (`pulumi.Input[str]`)

        The **okta_otp** object supports the following:

          * `consent_type` (`pulumi.Input[str]`)
          * `enroll` (`pulumi.Input[str]`)

        The **okta_password** object supports the following:

          * `consent_type` (`pulumi.Input[str]`)
          * `enroll` (`pulumi.Input[str]`)

        The **okta_push** object supports the following:

          * `consent_type` (`pulumi.Input[str]`)
          * `enroll` (`pulumi.Input[str]`)

        The **okta_question** object supports the following:

          * `consent_type` (`pulumi.Input[str]`)
          * `enroll` (`pulumi.Input[str]`)

        The **okta_sms** object supports the following:

          * `consent_type` (`pulumi.Input[str]`)
          * `enroll` (`pulumi.Input[str]`)

        The **rsa_token** object supports the following:

          * `consent_type` (`pulumi.Input[str]`)
          * `enroll` (`pulumi.Input[str]`)

        The **symantec_vip** object supports the following:

          * `consent_type` (`pulumi.Input[str]`)
          * `enroll` (`pulumi.Input[str]`)

        The **yubikey_token** object supports the following:

          * `consent_type` (`pulumi.Input[str]`)
          * `enroll` (`pulumi.Input[str]`)
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

            __props__['description'] = description
            __props__['duo'] = duo
            __props__['fido_u2f'] = fido_u2f
            __props__['fido_webauthn'] = fido_webauthn
            __props__['google_otp'] = google_otp
            __props__['groups_includeds'] = groups_includeds
            __props__['name'] = name
            __props__['okta_call'] = okta_call
            __props__['okta_otp'] = okta_otp
            __props__['okta_password'] = okta_password
            __props__['okta_push'] = okta_push
            __props__['okta_question'] = okta_question
            __props__['okta_sms'] = okta_sms
            __props__['priority'] = priority
            __props__['rsa_token'] = rsa_token
            __props__['status'] = status
            __props__['symantec_vip'] = symantec_vip
            __props__['yubikey_token'] = yubikey_token
        super(MfaPolicy, __self__).__init__(
            'okta:deprecated/mfaPolicy:MfaPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, duo=None, fido_u2f=None, fido_webauthn=None, google_otp=None, groups_includeds=None, name=None, okta_call=None, okta_otp=None, okta_password=None, okta_push=None, okta_question=None, okta_sms=None, priority=None, rsa_token=None, status=None, symantec_vip=None, yubikey_token=None):
        """
        Get an existing MfaPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Policy Description
        :param pulumi.Input[list] groups_includeds: List of Group IDs to Include
        :param pulumi.Input[str] name: Policy Name
        :param pulumi.Input[float] priority: Policy Priority, this attribute can be set to a valid priority. To avoid endless diff situation we error if an invalid
               priority is provided. API defaults it to the last/lowest if not there.
        :param pulumi.Input[str] status: Policy Status: ACTIVE or INACTIVE.

        The **duo** object supports the following:

          * `consent_type` (`pulumi.Input[str]`)
          * `enroll` (`pulumi.Input[str]`)

        The **fido_u2f** object supports the following:

          * `consent_type` (`pulumi.Input[str]`)
          * `enroll` (`pulumi.Input[str]`)

        The **fido_webauthn** object supports the following:

          * `consent_type` (`pulumi.Input[str]`)
          * `enroll` (`pulumi.Input[str]`)

        The **google_otp** object supports the following:

          * `consent_type` (`pulumi.Input[str]`)
          * `enroll` (`pulumi.Input[str]`)

        The **okta_call** object supports the following:

          * `consent_type` (`pulumi.Input[str]`)
          * `enroll` (`pulumi.Input[str]`)

        The **okta_otp** object supports the following:

          * `consent_type` (`pulumi.Input[str]`)
          * `enroll` (`pulumi.Input[str]`)

        The **okta_password** object supports the following:

          * `consent_type` (`pulumi.Input[str]`)
          * `enroll` (`pulumi.Input[str]`)

        The **okta_push** object supports the following:

          * `consent_type` (`pulumi.Input[str]`)
          * `enroll` (`pulumi.Input[str]`)

        The **okta_question** object supports the following:

          * `consent_type` (`pulumi.Input[str]`)
          * `enroll` (`pulumi.Input[str]`)

        The **okta_sms** object supports the following:

          * `consent_type` (`pulumi.Input[str]`)
          * `enroll` (`pulumi.Input[str]`)

        The **rsa_token** object supports the following:

          * `consent_type` (`pulumi.Input[str]`)
          * `enroll` (`pulumi.Input[str]`)

        The **symantec_vip** object supports the following:

          * `consent_type` (`pulumi.Input[str]`)
          * `enroll` (`pulumi.Input[str]`)

        The **yubikey_token** object supports the following:

          * `consent_type` (`pulumi.Input[str]`)
          * `enroll` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["duo"] = duo
        __props__["fido_u2f"] = fido_u2f
        __props__["fido_webauthn"] = fido_webauthn
        __props__["google_otp"] = google_otp
        __props__["groups_includeds"] = groups_includeds
        __props__["name"] = name
        __props__["okta_call"] = okta_call
        __props__["okta_otp"] = okta_otp
        __props__["okta_password"] = okta_password
        __props__["okta_push"] = okta_push
        __props__["okta_question"] = okta_question
        __props__["okta_sms"] = okta_sms
        __props__["priority"] = priority
        __props__["rsa_token"] = rsa_token
        __props__["status"] = status
        __props__["symantec_vip"] = symantec_vip
        __props__["yubikey_token"] = yubikey_token
        return MfaPolicy(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
