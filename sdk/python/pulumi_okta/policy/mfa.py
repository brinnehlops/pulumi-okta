# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Mfa(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    Policy Description.
    """
    duo: pulumi.Output[dict]
    """
    DUO MFA policy settings.
    
      * `consent_type` (`str`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
      * `enroll` (`str`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
    """
    fido_u2f: pulumi.Output[dict]
    """
    Fido U2F MFA policy settings.
    
      * `consent_type` (`str`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
      * `enroll` (`str`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
    """
    fido_webauthn: pulumi.Output[dict]
    """
    Fido Web Authn MFA policy settings.
    
      * `consent_type` (`str`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
      * `enroll` (`str`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
    """
    google_otp: pulumi.Output[dict]
    """
    Google OTP MFA policy settings.
    
      * `consent_type` (`str`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
      * `enroll` (`str`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
    """
    groups_includeds: pulumi.Output[list]
    """
    List of Group IDs to Include.
    """
    name: pulumi.Output[str]
    """
    Policy Name.
    """
    okta_call: pulumi.Output[dict]
    """
    Okta Call MFA policy settings.
    
      * `consent_type` (`str`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
      * `enroll` (`str`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
    """
    okta_otp: pulumi.Output[dict]
    """
    Okta OTP MFA policy settings.
    
      * `consent_type` (`str`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
      * `enroll` (`str`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
    """
    okta_password: pulumi.Output[dict]
    """
    Okta Password MFA policy settings.
    
      * `consent_type` (`str`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
      * `enroll` (`str`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
    """
    okta_push: pulumi.Output[dict]
    """
    Okta Push MFA policy settings.
    
      * `consent_type` (`str`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
      * `enroll` (`str`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
    """
    okta_question: pulumi.Output[dict]
    """
    Okta Question MFA policy settings.
    
      * `consent_type` (`str`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
      * `enroll` (`str`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
    """
    okta_sms: pulumi.Output[dict]
    """
    Okta SMS MFA policy settings.
    
      * `consent_type` (`str`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
      * `enroll` (`str`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
    """
    priority: pulumi.Output[float]
    """
    Priority of the policy.
    """
    rsa_token: pulumi.Output[dict]
    """
    RSA Token MFA policy settings.
    
      * `consent_type` (`str`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
      * `enroll` (`str`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
    """
    status: pulumi.Output[str]
    """
    Policy Status: `"ACTIVE"` or `"INACTIVE"`.
    """
    symantec_vip: pulumi.Output[dict]
    """
    Symantec VIP MFA policy settings.
    
      * `consent_type` (`str`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
      * `enroll` (`str`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
    """
    yubikey_token: pulumi.Output[dict]
    """
    Yubikey Token MFA policy settings.
    
      * `consent_type` (`str`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
      * `enroll` (`str`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
    """
    def __init__(__self__, resource_name, opts=None, description=None, duo=None, fido_u2f=None, fido_webauthn=None, google_otp=None, groups_includeds=None, name=None, okta_call=None, okta_otp=None, okta_password=None, okta_push=None, okta_question=None, okta_sms=None, priority=None, rsa_token=None, status=None, symantec_vip=None, yubikey_token=None, __props__=None, __name__=None, __opts__=None):
        """
        Creates an MFA Policy.
        
        This resource allows you to create and configure an MFA Policy.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Policy Description.
        :param pulumi.Input[dict] duo: DUO MFA policy settings.
        :param pulumi.Input[dict] fido_u2f: Fido U2F MFA policy settings.
        :param pulumi.Input[dict] fido_webauthn: Fido Web Authn MFA policy settings.
        :param pulumi.Input[dict] google_otp: Google OTP MFA policy settings.
        :param pulumi.Input[list] groups_includeds: List of Group IDs to Include.
        :param pulumi.Input[str] name: Policy Name.
        :param pulumi.Input[dict] okta_call: Okta Call MFA policy settings.
        :param pulumi.Input[dict] okta_otp: Okta OTP MFA policy settings.
        :param pulumi.Input[dict] okta_password: Okta Password MFA policy settings.
        :param pulumi.Input[dict] okta_push: Okta Push MFA policy settings.
        :param pulumi.Input[dict] okta_question: Okta Question MFA policy settings.
        :param pulumi.Input[dict] okta_sms: Okta SMS MFA policy settings.
        :param pulumi.Input[float] priority: Priority of the policy.
        :param pulumi.Input[dict] rsa_token: RSA Token MFA policy settings.
        :param pulumi.Input[str] status: Policy Status: `"ACTIVE"` or `"INACTIVE"`.
        :param pulumi.Input[dict] symantec_vip: Symantec VIP MFA policy settings.
        :param pulumi.Input[dict] yubikey_token: Yubikey Token MFA policy settings.
        
        The **duo** object supports the following:
        
          * `consent_type` (`pulumi.Input[str]`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
          * `enroll` (`pulumi.Input[str]`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
        
        The **fido_u2f** object supports the following:
        
          * `consent_type` (`pulumi.Input[str]`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
          * `enroll` (`pulumi.Input[str]`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
        
        The **fido_webauthn** object supports the following:
        
          * `consent_type` (`pulumi.Input[str]`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
          * `enroll` (`pulumi.Input[str]`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
        
        The **google_otp** object supports the following:
        
          * `consent_type` (`pulumi.Input[str]`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
          * `enroll` (`pulumi.Input[str]`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
        
        The **okta_call** object supports the following:
        
          * `consent_type` (`pulumi.Input[str]`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
          * `enroll` (`pulumi.Input[str]`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
        
        The **okta_otp** object supports the following:
        
          * `consent_type` (`pulumi.Input[str]`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
          * `enroll` (`pulumi.Input[str]`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
        
        The **okta_password** object supports the following:
        
          * `consent_type` (`pulumi.Input[str]`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
          * `enroll` (`pulumi.Input[str]`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
        
        The **okta_push** object supports the following:
        
          * `consent_type` (`pulumi.Input[str]`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
          * `enroll` (`pulumi.Input[str]`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
        
        The **okta_question** object supports the following:
        
          * `consent_type` (`pulumi.Input[str]`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
          * `enroll` (`pulumi.Input[str]`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
        
        The **okta_sms** object supports the following:
        
          * `consent_type` (`pulumi.Input[str]`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
          * `enroll` (`pulumi.Input[str]`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
        
        The **rsa_token** object supports the following:
        
          * `consent_type` (`pulumi.Input[str]`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
          * `enroll` (`pulumi.Input[str]`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
        
        The **symantec_vip** object supports the following:
        
          * `consent_type` (`pulumi.Input[str]`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
          * `enroll` (`pulumi.Input[str]`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
        
        The **yubikey_token** object supports the following:
        
          * `consent_type` (`pulumi.Input[str]`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
          * `enroll` (`pulumi.Input[str]`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.

        > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/policy_mfa.html.markdown.
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
        super(Mfa, __self__).__init__(
            'okta:policy/mfa:Mfa',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, duo=None, fido_u2f=None, fido_webauthn=None, google_otp=None, groups_includeds=None, name=None, okta_call=None, okta_otp=None, okta_password=None, okta_push=None, okta_question=None, okta_sms=None, priority=None, rsa_token=None, status=None, symantec_vip=None, yubikey_token=None):
        """
        Get an existing Mfa resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Policy Description.
        :param pulumi.Input[dict] duo: DUO MFA policy settings.
        :param pulumi.Input[dict] fido_u2f: Fido U2F MFA policy settings.
        :param pulumi.Input[dict] fido_webauthn: Fido Web Authn MFA policy settings.
        :param pulumi.Input[dict] google_otp: Google OTP MFA policy settings.
        :param pulumi.Input[list] groups_includeds: List of Group IDs to Include.
        :param pulumi.Input[str] name: Policy Name.
        :param pulumi.Input[dict] okta_call: Okta Call MFA policy settings.
        :param pulumi.Input[dict] okta_otp: Okta OTP MFA policy settings.
        :param pulumi.Input[dict] okta_password: Okta Password MFA policy settings.
        :param pulumi.Input[dict] okta_push: Okta Push MFA policy settings.
        :param pulumi.Input[dict] okta_question: Okta Question MFA policy settings.
        :param pulumi.Input[dict] okta_sms: Okta SMS MFA policy settings.
        :param pulumi.Input[float] priority: Priority of the policy.
        :param pulumi.Input[dict] rsa_token: RSA Token MFA policy settings.
        :param pulumi.Input[str] status: Policy Status: `"ACTIVE"` or `"INACTIVE"`.
        :param pulumi.Input[dict] symantec_vip: Symantec VIP MFA policy settings.
        :param pulumi.Input[dict] yubikey_token: Yubikey Token MFA policy settings.
        
        The **duo** object supports the following:
        
          * `consent_type` (`pulumi.Input[str]`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
          * `enroll` (`pulumi.Input[str]`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
        
        The **fido_u2f** object supports the following:
        
          * `consent_type` (`pulumi.Input[str]`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
          * `enroll` (`pulumi.Input[str]`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
        
        The **fido_webauthn** object supports the following:
        
          * `consent_type` (`pulumi.Input[str]`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
          * `enroll` (`pulumi.Input[str]`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
        
        The **google_otp** object supports the following:
        
          * `consent_type` (`pulumi.Input[str]`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
          * `enroll` (`pulumi.Input[str]`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
        
        The **okta_call** object supports the following:
        
          * `consent_type` (`pulumi.Input[str]`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
          * `enroll` (`pulumi.Input[str]`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
        
        The **okta_otp** object supports the following:
        
          * `consent_type` (`pulumi.Input[str]`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
          * `enroll` (`pulumi.Input[str]`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
        
        The **okta_password** object supports the following:
        
          * `consent_type` (`pulumi.Input[str]`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
          * `enroll` (`pulumi.Input[str]`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
        
        The **okta_push** object supports the following:
        
          * `consent_type` (`pulumi.Input[str]`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
          * `enroll` (`pulumi.Input[str]`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
        
        The **okta_question** object supports the following:
        
          * `consent_type` (`pulumi.Input[str]`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
          * `enroll` (`pulumi.Input[str]`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
        
        The **okta_sms** object supports the following:
        
          * `consent_type` (`pulumi.Input[str]`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
          * `enroll` (`pulumi.Input[str]`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
        
        The **rsa_token** object supports the following:
        
          * `consent_type` (`pulumi.Input[str]`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
          * `enroll` (`pulumi.Input[str]`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
        
        The **symantec_vip** object supports the following:
        
          * `consent_type` (`pulumi.Input[str]`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
          * `enroll` (`pulumi.Input[str]`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.
        
        The **yubikey_token** object supports the following:
        
          * `consent_type` (`pulumi.Input[str]`) - User consent type required before enrolling in the factor: `"NONE"` or `"TERMS_OF_SERVICE"`. By default it is `"NONE"`.
          * `enroll` (`pulumi.Input[str]`) - Requirements for user initiated enrollment. Can be `"NOT_ALLOWED"`, `"OPTIONAL"`, or `"REQUIRED"`. By default it is `"OPTIONAL"`.

        > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/policy_mfa.html.markdown.
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
        return Mfa(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

