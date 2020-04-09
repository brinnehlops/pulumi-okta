# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Password(pulumi.CustomResource):
    auth_provider: pulumi.Output[str]
    """
    Authentication Provider: `"OKTA"` or `"ACTIVE_DIRECTORY"`. Default is `"OKTA"`.
    """
    description: pulumi.Output[str]
    """
    Policy Description.
    """
    email_recovery: pulumi.Output[str]
    """
    Enable or disable email password recovery: ACTIVE or INACTIVE.
    """
    groups_includeds: pulumi.Output[list]
    """
    List of Group IDs to Include.
    """
    name: pulumi.Output[str]
    """
    Policy Name.
    """
    password_auto_unlock_minutes: pulumi.Output[float]
    """
    Number of minutes before a locked account is unlocked: 0 = no limit.
    """
    password_dictionary_lookup: pulumi.Output[bool]
    """
    Check Passwords Against Common Password Dictionary.
    """
    password_exclude_first_name: pulumi.Output[bool]
    """
    User firstName attribute must be excluded from the password.
    """
    password_exclude_last_name: pulumi.Output[bool]
    """
    User lastName attribute must be excluded from the password.
    """
    password_exclude_username: pulumi.Output[bool]
    """
    If the user name must be excluded from the password.
    """
    password_expire_warn_days: pulumi.Output[float]
    """
    Length in days a user will be warned before password expiry: 0 = no warning.
    """
    password_history_count: pulumi.Output[float]
    """
    Number of distinct passwords that must be created before they can be reused: 0 = none.
    """
    password_max_age_days: pulumi.Output[float]
    """
    Length in days a password is valid before expiry: 0 = no limit.",
    """
    password_max_lockout_attempts: pulumi.Output[float]
    """
    Number of unsuccessful login attempts allowed before lockout: 0 = no limit.
    """
    password_min_age_minutes: pulumi.Output[float]
    """
    Minimum time interval in minutes between password changes: 0 = no limit.
    """
    password_min_length: pulumi.Output[float]
    """
    Minimum password length. Default is 8.
    """
    password_min_lowercase: pulumi.Output[float]
    """
    Minimum number of lower case characters in password.
    """
    password_min_number: pulumi.Output[float]
    """
    Minimum number of numbers in password.
    """
    password_min_symbol: pulumi.Output[float]
    """
    Minimum number of symbols in password.
    """
    password_min_uppercase: pulumi.Output[float]
    """
    Minimum number of upper case characters in password.
    """
    password_show_lockout_failures: pulumi.Output[bool]
    """
    If a user should be informed when their account is locked.
    """
    priority: pulumi.Output[float]
    """
    Priority of the policy.
    """
    question_min_length: pulumi.Output[float]
    """
    Min length of the password recovery question answer.
    """
    question_recovery: pulumi.Output[str]
    """
    Enable or disable security question password recovery: ACTIVE or INACTIVE.
    """
    recovery_email_token: pulumi.Output[float]
    """
    Lifetime in minutes of the recovery email token.
    """
    skip_unlock: pulumi.Output[bool]
    """
    When an Active Directory user is locked out of Okta, the Okta unlock operation should also attempt to unlock the user's Windows account.
    """
    sms_recovery: pulumi.Output[str]
    """
    Enable or disable SMS password recovery: ACTIVE or INACTIVE.
    """
    status: pulumi.Output[str]
    """
    Policy Status: `"ACTIVE"` or `"INACTIVE"`.
    """
    def __init__(__self__, resource_name, opts=None, auth_provider=None, description=None, email_recovery=None, groups_includeds=None, name=None, password_auto_unlock_minutes=None, password_dictionary_lookup=None, password_exclude_first_name=None, password_exclude_last_name=None, password_exclude_username=None, password_expire_warn_days=None, password_history_count=None, password_max_age_days=None, password_max_lockout_attempts=None, password_min_age_minutes=None, password_min_length=None, password_min_lowercase=None, password_min_number=None, password_min_symbol=None, password_min_uppercase=None, password_show_lockout_failures=None, priority=None, question_min_length=None, question_recovery=None, recovery_email_token=None, skip_unlock=None, sms_recovery=None, status=None, __props__=None, __name__=None, __opts__=None):
        """
        Creates a Password Policy.

        This resource allows you to create and configure a Password Policy.



        > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/r/policy_password.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_provider: Authentication Provider: `"OKTA"` or `"ACTIVE_DIRECTORY"`. Default is `"OKTA"`.
        :param pulumi.Input[str] description: Policy Description.
        :param pulumi.Input[str] email_recovery: Enable or disable email password recovery: ACTIVE or INACTIVE.
        :param pulumi.Input[list] groups_includeds: List of Group IDs to Include.
        :param pulumi.Input[str] name: Policy Name.
        :param pulumi.Input[float] password_auto_unlock_minutes: Number of minutes before a locked account is unlocked: 0 = no limit.
        :param pulumi.Input[bool] password_dictionary_lookup: Check Passwords Against Common Password Dictionary.
        :param pulumi.Input[bool] password_exclude_first_name: User firstName attribute must be excluded from the password.
        :param pulumi.Input[bool] password_exclude_last_name: User lastName attribute must be excluded from the password.
        :param pulumi.Input[bool] password_exclude_username: If the user name must be excluded from the password.
        :param pulumi.Input[float] password_expire_warn_days: Length in days a user will be warned before password expiry: 0 = no warning.
        :param pulumi.Input[float] password_history_count: Number of distinct passwords that must be created before they can be reused: 0 = none.
        :param pulumi.Input[float] password_max_age_days: Length in days a password is valid before expiry: 0 = no limit.",
        :param pulumi.Input[float] password_max_lockout_attempts: Number of unsuccessful login attempts allowed before lockout: 0 = no limit.
        :param pulumi.Input[float] password_min_age_minutes: Minimum time interval in minutes between password changes: 0 = no limit.
        :param pulumi.Input[float] password_min_length: Minimum password length. Default is 8.
        :param pulumi.Input[float] password_min_lowercase: Minimum number of lower case characters in password.
        :param pulumi.Input[float] password_min_number: Minimum number of numbers in password.
        :param pulumi.Input[float] password_min_symbol: Minimum number of symbols in password.
        :param pulumi.Input[float] password_min_uppercase: Minimum number of upper case characters in password.
        :param pulumi.Input[bool] password_show_lockout_failures: If a user should be informed when their account is locked.
        :param pulumi.Input[float] priority: Priority of the policy.
        :param pulumi.Input[float] question_min_length: Min length of the password recovery question answer.
        :param pulumi.Input[str] question_recovery: Enable or disable security question password recovery: ACTIVE or INACTIVE.
        :param pulumi.Input[float] recovery_email_token: Lifetime in minutes of the recovery email token.
        :param pulumi.Input[bool] skip_unlock: When an Active Directory user is locked out of Okta, the Okta unlock operation should also attempt to unlock the user's Windows account.
        :param pulumi.Input[str] sms_recovery: Enable or disable SMS password recovery: ACTIVE or INACTIVE.
        :param pulumi.Input[str] status: Policy Status: `"ACTIVE"` or `"INACTIVE"`.
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

            __props__['auth_provider'] = auth_provider
            __props__['description'] = description
            __props__['email_recovery'] = email_recovery
            __props__['groups_includeds'] = groups_includeds
            __props__['name'] = name
            __props__['password_auto_unlock_minutes'] = password_auto_unlock_minutes
            __props__['password_dictionary_lookup'] = password_dictionary_lookup
            __props__['password_exclude_first_name'] = password_exclude_first_name
            __props__['password_exclude_last_name'] = password_exclude_last_name
            __props__['password_exclude_username'] = password_exclude_username
            __props__['password_expire_warn_days'] = password_expire_warn_days
            __props__['password_history_count'] = password_history_count
            __props__['password_max_age_days'] = password_max_age_days
            __props__['password_max_lockout_attempts'] = password_max_lockout_attempts
            __props__['password_min_age_minutes'] = password_min_age_minutes
            __props__['password_min_length'] = password_min_length
            __props__['password_min_lowercase'] = password_min_lowercase
            __props__['password_min_number'] = password_min_number
            __props__['password_min_symbol'] = password_min_symbol
            __props__['password_min_uppercase'] = password_min_uppercase
            __props__['password_show_lockout_failures'] = password_show_lockout_failures
            __props__['priority'] = priority
            __props__['question_min_length'] = question_min_length
            __props__['question_recovery'] = question_recovery
            __props__['recovery_email_token'] = recovery_email_token
            __props__['skip_unlock'] = skip_unlock
            __props__['sms_recovery'] = sms_recovery
            __props__['status'] = status
        super(Password, __self__).__init__(
            'okta:policy/password:Password',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, auth_provider=None, description=None, email_recovery=None, groups_includeds=None, name=None, password_auto_unlock_minutes=None, password_dictionary_lookup=None, password_exclude_first_name=None, password_exclude_last_name=None, password_exclude_username=None, password_expire_warn_days=None, password_history_count=None, password_max_age_days=None, password_max_lockout_attempts=None, password_min_age_minutes=None, password_min_length=None, password_min_lowercase=None, password_min_number=None, password_min_symbol=None, password_min_uppercase=None, password_show_lockout_failures=None, priority=None, question_min_length=None, question_recovery=None, recovery_email_token=None, skip_unlock=None, sms_recovery=None, status=None):
        """
        Get an existing Password resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_provider: Authentication Provider: `"OKTA"` or `"ACTIVE_DIRECTORY"`. Default is `"OKTA"`.
        :param pulumi.Input[str] description: Policy Description.
        :param pulumi.Input[str] email_recovery: Enable or disable email password recovery: ACTIVE or INACTIVE.
        :param pulumi.Input[list] groups_includeds: List of Group IDs to Include.
        :param pulumi.Input[str] name: Policy Name.
        :param pulumi.Input[float] password_auto_unlock_minutes: Number of minutes before a locked account is unlocked: 0 = no limit.
        :param pulumi.Input[bool] password_dictionary_lookup: Check Passwords Against Common Password Dictionary.
        :param pulumi.Input[bool] password_exclude_first_name: User firstName attribute must be excluded from the password.
        :param pulumi.Input[bool] password_exclude_last_name: User lastName attribute must be excluded from the password.
        :param pulumi.Input[bool] password_exclude_username: If the user name must be excluded from the password.
        :param pulumi.Input[float] password_expire_warn_days: Length in days a user will be warned before password expiry: 0 = no warning.
        :param pulumi.Input[float] password_history_count: Number of distinct passwords that must be created before they can be reused: 0 = none.
        :param pulumi.Input[float] password_max_age_days: Length in days a password is valid before expiry: 0 = no limit.",
        :param pulumi.Input[float] password_max_lockout_attempts: Number of unsuccessful login attempts allowed before lockout: 0 = no limit.
        :param pulumi.Input[float] password_min_age_minutes: Minimum time interval in minutes between password changes: 0 = no limit.
        :param pulumi.Input[float] password_min_length: Minimum password length. Default is 8.
        :param pulumi.Input[float] password_min_lowercase: Minimum number of lower case characters in password.
        :param pulumi.Input[float] password_min_number: Minimum number of numbers in password.
        :param pulumi.Input[float] password_min_symbol: Minimum number of symbols in password.
        :param pulumi.Input[float] password_min_uppercase: Minimum number of upper case characters in password.
        :param pulumi.Input[bool] password_show_lockout_failures: If a user should be informed when their account is locked.
        :param pulumi.Input[float] priority: Priority of the policy.
        :param pulumi.Input[float] question_min_length: Min length of the password recovery question answer.
        :param pulumi.Input[str] question_recovery: Enable or disable security question password recovery: ACTIVE or INACTIVE.
        :param pulumi.Input[float] recovery_email_token: Lifetime in minutes of the recovery email token.
        :param pulumi.Input[bool] skip_unlock: When an Active Directory user is locked out of Okta, the Okta unlock operation should also attempt to unlock the user's Windows account.
        :param pulumi.Input[str] sms_recovery: Enable or disable SMS password recovery: ACTIVE or INACTIVE.
        :param pulumi.Input[str] status: Policy Status: `"ACTIVE"` or `"INACTIVE"`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["auth_provider"] = auth_provider
        __props__["description"] = description
        __props__["email_recovery"] = email_recovery
        __props__["groups_includeds"] = groups_includeds
        __props__["name"] = name
        __props__["password_auto_unlock_minutes"] = password_auto_unlock_minutes
        __props__["password_dictionary_lookup"] = password_dictionary_lookup
        __props__["password_exclude_first_name"] = password_exclude_first_name
        __props__["password_exclude_last_name"] = password_exclude_last_name
        __props__["password_exclude_username"] = password_exclude_username
        __props__["password_expire_warn_days"] = password_expire_warn_days
        __props__["password_history_count"] = password_history_count
        __props__["password_max_age_days"] = password_max_age_days
        __props__["password_max_lockout_attempts"] = password_max_lockout_attempts
        __props__["password_min_age_minutes"] = password_min_age_minutes
        __props__["password_min_length"] = password_min_length
        __props__["password_min_lowercase"] = password_min_lowercase
        __props__["password_min_number"] = password_min_number
        __props__["password_min_symbol"] = password_min_symbol
        __props__["password_min_uppercase"] = password_min_uppercase
        __props__["password_show_lockout_failures"] = password_show_lockout_failures
        __props__["priority"] = priority
        __props__["question_min_length"] = question_min_length
        __props__["question_recovery"] = question_recovery
        __props__["recovery_email_token"] = recovery_email_token
        __props__["skip_unlock"] = skip_unlock
        __props__["sms_recovery"] = sms_recovery
        __props__["status"] = status
        return Password(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

