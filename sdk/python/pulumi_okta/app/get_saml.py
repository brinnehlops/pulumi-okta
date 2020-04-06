# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetSamlResult:
    """
    A collection of values returned by getSaml.
    """
    def __init__(__self__, accessibility_error_redirect_url=None, accessibility_login_redirect_url=None, accessibility_self_service=None, active_only=None, app_settings_json=None, assertion_signed=None, attribute_statements=None, audience=None, authn_context_class_ref=None, auto_submit_toolbar=None, default_relay_state=None, description=None, destination=None, digest_algorithm=None, features=None, hide_ios=None, hide_web=None, honor_force_authn=None, id=None, idp_issuer=None, key_id=None, label=None, label_prefix=None, name=None, recipient=None, request_compressed=None, response_signed=None, signature_algorithm=None, sp_issuer=None, sso_url=None, status=None, subject_name_id_format=None, subject_name_id_template=None, user_name_template=None, user_name_template_suffix=None, user_name_template_type=None):
        if accessibility_error_redirect_url and not isinstance(accessibility_error_redirect_url, str):
            raise TypeError("Expected argument 'accessibility_error_redirect_url' to be a str")
        __self__.accessibility_error_redirect_url = accessibility_error_redirect_url
        """
        Custom error page URL.
        """
        if accessibility_login_redirect_url and not isinstance(accessibility_login_redirect_url, str):
            raise TypeError("Expected argument 'accessibility_login_redirect_url' to be a str")
        __self__.accessibility_login_redirect_url = accessibility_login_redirect_url
        """
        Custom login page URL.
        """
        if accessibility_self_service and not isinstance(accessibility_self_service, bool):
            raise TypeError("Expected argument 'accessibility_self_service' to be a bool")
        __self__.accessibility_self_service = accessibility_self_service
        """
        Enable self service.
        """
        if active_only and not isinstance(active_only, bool):
            raise TypeError("Expected argument 'active_only' to be a bool")
        __self__.active_only = active_only
        if app_settings_json and not isinstance(app_settings_json, str):
            raise TypeError("Expected argument 'app_settings_json' to be a str")
        __self__.app_settings_json = app_settings_json
        """
        Application settings in JSON format.
        """
        if assertion_signed and not isinstance(assertion_signed, bool):
            raise TypeError("Expected argument 'assertion_signed' to be a bool")
        __self__.assertion_signed = assertion_signed
        """
        Determines whether the SAML assertion is digitally signed.
        """
        if attribute_statements and not isinstance(attribute_statements, list):
            raise TypeError("Expected argument 'attribute_statements' to be a list")
        __self__.attribute_statements = attribute_statements
        """
        SAML Attribute statements.
        """
        if audience and not isinstance(audience, str):
            raise TypeError("Expected argument 'audience' to be a str")
        __self__.audience = audience
        """
        Audience restriction.
        """
        if authn_context_class_ref and not isinstance(authn_context_class_ref, str):
            raise TypeError("Expected argument 'authn_context_class_ref' to be a str")
        __self__.authn_context_class_ref = authn_context_class_ref
        """
        Identifies the SAML authentication context class for the assertion’s authentication statement.
        """
        if auto_submit_toolbar and not isinstance(auto_submit_toolbar, bool):
            raise TypeError("Expected argument 'auto_submit_toolbar' to be a bool")
        __self__.auto_submit_toolbar = auto_submit_toolbar
        """
        Display auto submit toolbar.
        """
        if default_relay_state and not isinstance(default_relay_state, str):
            raise TypeError("Expected argument 'default_relay_state' to be a str")
        __self__.default_relay_state = default_relay_state
        """
        Identifies a specific application resource in an IDP initiated SSO scenario.
        """
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        description of application.
        """
        if destination and not isinstance(destination, str):
            raise TypeError("Expected argument 'destination' to be a str")
        __self__.destination = destination
        """
        Identifies the location where the SAML response is intended to be sent inside of the SAML assertion.
        """
        if digest_algorithm and not isinstance(digest_algorithm, str):
            raise TypeError("Expected argument 'digest_algorithm' to be a str")
        __self__.digest_algorithm = digest_algorithm
        """
        Determines the digest algorithm used to digitally sign the SAML assertion and response.
        """
        if features and not isinstance(features, list):
            raise TypeError("Expected argument 'features' to be a list")
        __self__.features = features
        """
        features enabled.
        """
        if hide_ios and not isinstance(hide_ios, bool):
            raise TypeError("Expected argument 'hide_ios' to be a bool")
        __self__.hide_ios = hide_ios
        """
        Do not display application icon on mobile app.
        """
        if hide_web and not isinstance(hide_web, bool):
            raise TypeError("Expected argument 'hide_web' to be a bool")
        __self__.hide_web = hide_web
        """
        Do not display application icon to users
        """
        if honor_force_authn and not isinstance(honor_force_authn, bool):
            raise TypeError("Expected argument 'honor_force_authn' to be a bool")
        __self__.honor_force_authn = honor_force_authn
        """
        Prompt user to re-authenticate if SP asks for it.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id of application.
        """
        if idp_issuer and not isinstance(idp_issuer, str):
            raise TypeError("Expected argument 'idp_issuer' to be a str")
        __self__.idp_issuer = idp_issuer
        """
        SAML issuer ID.
        """
        if key_id and not isinstance(key_id, str):
            raise TypeError("Expected argument 'key_id' to be a str")
        __self__.key_id = key_id
        """
        Certificate key ID.
        """
        if label and not isinstance(label, str):
            raise TypeError("Expected argument 'label' to be a str")
        __self__.label = label
        """
        label of application.
        """
        if label_prefix and not isinstance(label_prefix, str):
            raise TypeError("Expected argument 'label_prefix' to be a str")
        __self__.label_prefix = label_prefix
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        name of application.
        """
        if recipient and not isinstance(recipient, str):
            raise TypeError("Expected argument 'recipient' to be a str")
        __self__.recipient = recipient
        """
        The location where the app may present the SAML assertion.
        """
        if request_compressed and not isinstance(request_compressed, bool):
            raise TypeError("Expected argument 'request_compressed' to be a bool")
        __self__.request_compressed = request_compressed
        """
        Denotes whether the request is compressed or not.
        """
        if response_signed and not isinstance(response_signed, bool):
            raise TypeError("Expected argument 'response_signed' to be a bool")
        __self__.response_signed = response_signed
        """
        Determines whether the SAML auth response message is digitally signed.
        """
        if signature_algorithm and not isinstance(signature_algorithm, str):
            raise TypeError("Expected argument 'signature_algorithm' to be a str")
        __self__.signature_algorithm = signature_algorithm
        """
        Signature algorithm used ot digitally sign the assertion and response.
        """
        if sp_issuer and not isinstance(sp_issuer, str):
            raise TypeError("Expected argument 'sp_issuer' to be a str")
        __self__.sp_issuer = sp_issuer
        """
        SAML service provider issuer.
        """
        if sso_url and not isinstance(sso_url, str):
            raise TypeError("Expected argument 'sso_url' to be a str")
        __self__.sso_url = sso_url
        """
        Single Sign on Url.
        """
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        __self__.status = status
        """
        status of application.
        """
        if subject_name_id_format and not isinstance(subject_name_id_format, str):
            raise TypeError("Expected argument 'subject_name_id_format' to be a str")
        __self__.subject_name_id_format = subject_name_id_format
        """
        Identifies the SAML processing rules.
        """
        if subject_name_id_template and not isinstance(subject_name_id_template, str):
            raise TypeError("Expected argument 'subject_name_id_template' to be a str")
        __self__.subject_name_id_template = subject_name_id_template
        """
        Template for app user's username when a user is assigned to the app.
        """
        if user_name_template and not isinstance(user_name_template, str):
            raise TypeError("Expected argument 'user_name_template' to be a str")
        __self__.user_name_template = user_name_template
        """
        Username template.
        """
        if user_name_template_suffix and not isinstance(user_name_template_suffix, str):
            raise TypeError("Expected argument 'user_name_template_suffix' to be a str")
        __self__.user_name_template_suffix = user_name_template_suffix
        """
        Username template suffix.
        """
        if user_name_template_type and not isinstance(user_name_template_type, str):
            raise TypeError("Expected argument 'user_name_template_type' to be a str")
        __self__.user_name_template_type = user_name_template_type
        """
        Username template type.
        """
class AwaitableGetSamlResult(GetSamlResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSamlResult(
            accessibility_error_redirect_url=self.accessibility_error_redirect_url,
            accessibility_login_redirect_url=self.accessibility_login_redirect_url,
            accessibility_self_service=self.accessibility_self_service,
            active_only=self.active_only,
            app_settings_json=self.app_settings_json,
            assertion_signed=self.assertion_signed,
            attribute_statements=self.attribute_statements,
            audience=self.audience,
            authn_context_class_ref=self.authn_context_class_ref,
            auto_submit_toolbar=self.auto_submit_toolbar,
            default_relay_state=self.default_relay_state,
            description=self.description,
            destination=self.destination,
            digest_algorithm=self.digest_algorithm,
            features=self.features,
            hide_ios=self.hide_ios,
            hide_web=self.hide_web,
            honor_force_authn=self.honor_force_authn,
            id=self.id,
            idp_issuer=self.idp_issuer,
            key_id=self.key_id,
            label=self.label,
            label_prefix=self.label_prefix,
            name=self.name,
            recipient=self.recipient,
            request_compressed=self.request_compressed,
            response_signed=self.response_signed,
            signature_algorithm=self.signature_algorithm,
            sp_issuer=self.sp_issuer,
            sso_url=self.sso_url,
            status=self.status,
            subject_name_id_format=self.subject_name_id_format,
            subject_name_id_template=self.subject_name_id_template,
            user_name_template=self.user_name_template,
            user_name_template_suffix=self.user_name_template_suffix,
            user_name_template_type=self.user_name_template_type)

def get_saml(accessibility_error_redirect_url=None,accessibility_login_redirect_url=None,accessibility_self_service=None,active_only=None,app_settings_json=None,assertion_signed=None,attribute_statements=None,audience=None,authn_context_class_ref=None,auto_submit_toolbar=None,default_relay_state=None,destination=None,digest_algorithm=None,features=None,hide_ios=None,hide_web=None,honor_force_authn=None,id=None,idp_issuer=None,label=None,label_prefix=None,recipient=None,request_compressed=None,response_signed=None,signature_algorithm=None,sp_issuer=None,sso_url=None,subject_name_id_format=None,subject_name_id_template=None,user_name_template=None,user_name_template_suffix=None,user_name_template_type=None,opts=None):
    """
    Use this data source to retrieve the collaborators for a given repository.

    > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/d/app_saml.html.markdown.


    :param str accessibility_error_redirect_url: Custom error page URL.
    :param str accessibility_login_redirect_url: Custom login page URL.
    :param bool accessibility_self_service: Enable self service.
    :param bool active_only: tells the provider to query for only `ACTIVE` applications.
    :param str app_settings_json: Application settings in JSON format.
    :param bool assertion_signed: Determines whether the SAML assertion is digitally signed.
    :param list attribute_statements: SAML Attribute statements.
    :param str audience: Audience restriction.
    :param str authn_context_class_ref: Identifies the SAML authentication context class for the assertion’s authentication statement.
    :param bool auto_submit_toolbar: Display auto submit toolbar.
    :param str default_relay_state: Identifies a specific application resource in an IDP initiated SSO scenario.
    :param str destination: Identifies the location where the SAML response is intended to be sent inside of the SAML assertion.
    :param str digest_algorithm: Determines the digest algorithm used to digitally sign the SAML assertion and response.
    :param list features: features enabled.
    :param bool hide_ios: Do not display application icon on mobile app.
    :param bool hide_web: Do not display application icon to users
    :param bool honor_force_authn: Prompt user to re-authenticate if SP asks for it.
    :param str id: `id` of application to retrieve, conflicts with `label` and `label_prefix`.
    :param str idp_issuer: SAML issuer ID.
    :param str label: The label of the app to retrieve, conflicts with `label_prefix` and `id`.
    :param str label_prefix: Label prefix of the app to retrieve, conflicts with `label` and `id`. This will tell the provider to do a `starts with` query as opposed to an `equals` query.
    :param str recipient: The location where the app may present the SAML assertion.
    :param bool request_compressed: Denotes whether the request is compressed or not.
    :param bool response_signed: Determines whether the SAML auth response message is digitally signed.
    :param str signature_algorithm: Signature algorithm used ot digitally sign the assertion and response.
    :param str sp_issuer: SAML service provider issuer.
    :param str sso_url: Single Sign on Url.
    :param str subject_name_id_format: Identifies the SAML processing rules.
    :param str subject_name_id_template: Template for app user's username when a user is assigned to the app.
    :param str user_name_template: Username template.
    :param str user_name_template_suffix: Username template suffix.
    :param str user_name_template_type: Username template type.

    The **attribute_statements** object supports the following:

      * `filterType` (`str`)
      * `filterValue` (`str`)
      * `name` (`str`) - name of application.
      * `namespace` (`str`)
      * `type` (`str`)
      * `values` (`list`)
    """
    __args__ = dict()


    __args__['accessibilityErrorRedirectUrl'] = accessibility_error_redirect_url
    __args__['accessibilityLoginRedirectUrl'] = accessibility_login_redirect_url
    __args__['accessibilitySelfService'] = accessibility_self_service
    __args__['activeOnly'] = active_only
    __args__['appSettingsJson'] = app_settings_json
    __args__['assertionSigned'] = assertion_signed
    __args__['attributeStatements'] = attribute_statements
    __args__['audience'] = audience
    __args__['authnContextClassRef'] = authn_context_class_ref
    __args__['autoSubmitToolbar'] = auto_submit_toolbar
    __args__['defaultRelayState'] = default_relay_state
    __args__['destination'] = destination
    __args__['digestAlgorithm'] = digest_algorithm
    __args__['features'] = features
    __args__['hideIos'] = hide_ios
    __args__['hideWeb'] = hide_web
    __args__['honorForceAuthn'] = honor_force_authn
    __args__['id'] = id
    __args__['idpIssuer'] = idp_issuer
    __args__['label'] = label
    __args__['labelPrefix'] = label_prefix
    __args__['recipient'] = recipient
    __args__['requestCompressed'] = request_compressed
    __args__['responseSigned'] = response_signed
    __args__['signatureAlgorithm'] = signature_algorithm
    __args__['spIssuer'] = sp_issuer
    __args__['ssoUrl'] = sso_url
    __args__['subjectNameIdFormat'] = subject_name_id_format
    __args__['subjectNameIdTemplate'] = subject_name_id_template
    __args__['userNameTemplate'] = user_name_template
    __args__['userNameTemplateSuffix'] = user_name_template_suffix
    __args__['userNameTemplateType'] = user_name_template_type
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('okta:app/getSaml:getSaml', __args__, opts=opts).value

    return AwaitableGetSamlResult(
        accessibility_error_redirect_url=__ret__.get('accessibilityErrorRedirectUrl'),
        accessibility_login_redirect_url=__ret__.get('accessibilityLoginRedirectUrl'),
        accessibility_self_service=__ret__.get('accessibilitySelfService'),
        active_only=__ret__.get('activeOnly'),
        app_settings_json=__ret__.get('appSettingsJson'),
        assertion_signed=__ret__.get('assertionSigned'),
        attribute_statements=__ret__.get('attributeStatements'),
        audience=__ret__.get('audience'),
        authn_context_class_ref=__ret__.get('authnContextClassRef'),
        auto_submit_toolbar=__ret__.get('autoSubmitToolbar'),
        default_relay_state=__ret__.get('defaultRelayState'),
        description=__ret__.get('description'),
        destination=__ret__.get('destination'),
        digest_algorithm=__ret__.get('digestAlgorithm'),
        features=__ret__.get('features'),
        hide_ios=__ret__.get('hideIos'),
        hide_web=__ret__.get('hideWeb'),
        honor_force_authn=__ret__.get('honorForceAuthn'),
        id=__ret__.get('id'),
        idp_issuer=__ret__.get('idpIssuer'),
        key_id=__ret__.get('keyId'),
        label=__ret__.get('label'),
        label_prefix=__ret__.get('labelPrefix'),
        name=__ret__.get('name'),
        recipient=__ret__.get('recipient'),
        request_compressed=__ret__.get('requestCompressed'),
        response_signed=__ret__.get('responseSigned'),
        signature_algorithm=__ret__.get('signatureAlgorithm'),
        sp_issuer=__ret__.get('spIssuer'),
        sso_url=__ret__.get('ssoUrl'),
        status=__ret__.get('status'),
        subject_name_id_format=__ret__.get('subjectNameIdFormat'),
        subject_name_id_template=__ret__.get('subjectNameIdTemplate'),
        user_name_template=__ret__.get('userNameTemplate'),
        user_name_template_suffix=__ret__.get('userNameTemplateSuffix'),
        user_name_template_type=__ret__.get('userNameTemplateType'))
