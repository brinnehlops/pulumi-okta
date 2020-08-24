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

__all__ = [
    'GetSamlResult',
    'AwaitableGetSamlResult',
    'get_saml',
]

@pulumi.output_type
class GetSamlResult:
    """
    A collection of values returned by getSaml.
    """
    def __init__(__self__, accessibility_error_redirect_url=None, accessibility_login_redirect_url=None, accessibility_self_service=None, active_only=None, app_settings_json=None, assertion_signed=None, attribute_statements=None, audience=None, authn_context_class_ref=None, auto_submit_toolbar=None, default_relay_state=None, description=None, destination=None, digest_algorithm=None, features=None, hide_ios=None, hide_web=None, honor_force_authn=None, id=None, idp_issuer=None, key_id=None, label=None, label_prefix=None, name=None, recipient=None, request_compressed=None, response_signed=None, signature_algorithm=None, sp_issuer=None, sso_url=None, status=None, subject_name_id_format=None, subject_name_id_template=None, user_name_template=None, user_name_template_suffix=None, user_name_template_type=None):
        if accessibility_error_redirect_url and not isinstance(accessibility_error_redirect_url, str):
            raise TypeError("Expected argument 'accessibility_error_redirect_url' to be a str")
        pulumi.set(__self__, "accessibility_error_redirect_url", accessibility_error_redirect_url)
        if accessibility_login_redirect_url and not isinstance(accessibility_login_redirect_url, str):
            raise TypeError("Expected argument 'accessibility_login_redirect_url' to be a str")
        pulumi.set(__self__, "accessibility_login_redirect_url", accessibility_login_redirect_url)
        if accessibility_self_service and not isinstance(accessibility_self_service, bool):
            raise TypeError("Expected argument 'accessibility_self_service' to be a bool")
        pulumi.set(__self__, "accessibility_self_service", accessibility_self_service)
        if active_only and not isinstance(active_only, bool):
            raise TypeError("Expected argument 'active_only' to be a bool")
        pulumi.set(__self__, "active_only", active_only)
        if app_settings_json and not isinstance(app_settings_json, str):
            raise TypeError("Expected argument 'app_settings_json' to be a str")
        pulumi.set(__self__, "app_settings_json", app_settings_json)
        if assertion_signed and not isinstance(assertion_signed, bool):
            raise TypeError("Expected argument 'assertion_signed' to be a bool")
        pulumi.set(__self__, "assertion_signed", assertion_signed)
        if attribute_statements and not isinstance(attribute_statements, list):
            raise TypeError("Expected argument 'attribute_statements' to be a list")
        pulumi.set(__self__, "attribute_statements", attribute_statements)
        if audience and not isinstance(audience, str):
            raise TypeError("Expected argument 'audience' to be a str")
        pulumi.set(__self__, "audience", audience)
        if authn_context_class_ref and not isinstance(authn_context_class_ref, str):
            raise TypeError("Expected argument 'authn_context_class_ref' to be a str")
        pulumi.set(__self__, "authn_context_class_ref", authn_context_class_ref)
        if auto_submit_toolbar and not isinstance(auto_submit_toolbar, bool):
            raise TypeError("Expected argument 'auto_submit_toolbar' to be a bool")
        pulumi.set(__self__, "auto_submit_toolbar", auto_submit_toolbar)
        if default_relay_state and not isinstance(default_relay_state, str):
            raise TypeError("Expected argument 'default_relay_state' to be a str")
        pulumi.set(__self__, "default_relay_state", default_relay_state)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if destination and not isinstance(destination, str):
            raise TypeError("Expected argument 'destination' to be a str")
        pulumi.set(__self__, "destination", destination)
        if digest_algorithm and not isinstance(digest_algorithm, str):
            raise TypeError("Expected argument 'digest_algorithm' to be a str")
        pulumi.set(__self__, "digest_algorithm", digest_algorithm)
        if features and not isinstance(features, list):
            raise TypeError("Expected argument 'features' to be a list")
        pulumi.set(__self__, "features", features)
        if hide_ios and not isinstance(hide_ios, bool):
            raise TypeError("Expected argument 'hide_ios' to be a bool")
        pulumi.set(__self__, "hide_ios", hide_ios)
        if hide_web and not isinstance(hide_web, bool):
            raise TypeError("Expected argument 'hide_web' to be a bool")
        pulumi.set(__self__, "hide_web", hide_web)
        if honor_force_authn and not isinstance(honor_force_authn, bool):
            raise TypeError("Expected argument 'honor_force_authn' to be a bool")
        pulumi.set(__self__, "honor_force_authn", honor_force_authn)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if idp_issuer and not isinstance(idp_issuer, str):
            raise TypeError("Expected argument 'idp_issuer' to be a str")
        pulumi.set(__self__, "idp_issuer", idp_issuer)
        if key_id and not isinstance(key_id, str):
            raise TypeError("Expected argument 'key_id' to be a str")
        pulumi.set(__self__, "key_id", key_id)
        if label and not isinstance(label, str):
            raise TypeError("Expected argument 'label' to be a str")
        pulumi.set(__self__, "label", label)
        if label_prefix and not isinstance(label_prefix, str):
            raise TypeError("Expected argument 'label_prefix' to be a str")
        pulumi.set(__self__, "label_prefix", label_prefix)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if recipient and not isinstance(recipient, str):
            raise TypeError("Expected argument 'recipient' to be a str")
        pulumi.set(__self__, "recipient", recipient)
        if request_compressed and not isinstance(request_compressed, bool):
            raise TypeError("Expected argument 'request_compressed' to be a bool")
        pulumi.set(__self__, "request_compressed", request_compressed)
        if response_signed and not isinstance(response_signed, bool):
            raise TypeError("Expected argument 'response_signed' to be a bool")
        pulumi.set(__self__, "response_signed", response_signed)
        if signature_algorithm and not isinstance(signature_algorithm, str):
            raise TypeError("Expected argument 'signature_algorithm' to be a str")
        pulumi.set(__self__, "signature_algorithm", signature_algorithm)
        if sp_issuer and not isinstance(sp_issuer, str):
            raise TypeError("Expected argument 'sp_issuer' to be a str")
        pulumi.set(__self__, "sp_issuer", sp_issuer)
        if sso_url and not isinstance(sso_url, str):
            raise TypeError("Expected argument 'sso_url' to be a str")
        pulumi.set(__self__, "sso_url", sso_url)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if subject_name_id_format and not isinstance(subject_name_id_format, str):
            raise TypeError("Expected argument 'subject_name_id_format' to be a str")
        pulumi.set(__self__, "subject_name_id_format", subject_name_id_format)
        if subject_name_id_template and not isinstance(subject_name_id_template, str):
            raise TypeError("Expected argument 'subject_name_id_template' to be a str")
        pulumi.set(__self__, "subject_name_id_template", subject_name_id_template)
        if user_name_template and not isinstance(user_name_template, str):
            raise TypeError("Expected argument 'user_name_template' to be a str")
        pulumi.set(__self__, "user_name_template", user_name_template)
        if user_name_template_suffix and not isinstance(user_name_template_suffix, str):
            raise TypeError("Expected argument 'user_name_template_suffix' to be a str")
        pulumi.set(__self__, "user_name_template_suffix", user_name_template_suffix)
        if user_name_template_type and not isinstance(user_name_template_type, str):
            raise TypeError("Expected argument 'user_name_template_type' to be a str")
        pulumi.set(__self__, "user_name_template_type", user_name_template_type)

    @property
    @pulumi.getter(name="accessibilityErrorRedirectUrl")
    def accessibility_error_redirect_url(self) -> Optional[str]:
        """
        Custom error page URL.
        """
        return pulumi.get(self, "accessibility_error_redirect_url")

    @property
    @pulumi.getter(name="accessibilityLoginRedirectUrl")
    def accessibility_login_redirect_url(self) -> Optional[str]:
        """
        Custom login page URL.
        """
        return pulumi.get(self, "accessibility_login_redirect_url")

    @property
    @pulumi.getter(name="accessibilitySelfService")
    def accessibility_self_service(self) -> Optional[bool]:
        """
        Enable self service.
        """
        return pulumi.get(self, "accessibility_self_service")

    @property
    @pulumi.getter(name="activeOnly")
    def active_only(self) -> Optional[bool]:
        return pulumi.get(self, "active_only")

    @property
    @pulumi.getter(name="appSettingsJson")
    def app_settings_json(self) -> Optional[str]:
        """
        Application settings in JSON format.
        """
        return pulumi.get(self, "app_settings_json")

    @property
    @pulumi.getter(name="assertionSigned")
    def assertion_signed(self) -> Optional[bool]:
        """
        Determines whether the SAML assertion is digitally signed.
        """
        return pulumi.get(self, "assertion_signed")

    @property
    @pulumi.getter(name="attributeStatements")
    def attribute_statements(self) -> Optional[List['outputs.GetSamlAttributeStatementResult']]:
        """
        SAML Attribute statements.
        """
        return pulumi.get(self, "attribute_statements")

    @property
    @pulumi.getter
    def audience(self) -> Optional[str]:
        """
        Audience restriction.
        """
        return pulumi.get(self, "audience")

    @property
    @pulumi.getter(name="authnContextClassRef")
    def authn_context_class_ref(self) -> Optional[str]:
        """
        Identifies the SAML authentication context class for the assertion’s authentication statement.
        """
        return pulumi.get(self, "authn_context_class_ref")

    @property
    @pulumi.getter(name="autoSubmitToolbar")
    def auto_submit_toolbar(self) -> Optional[bool]:
        """
        Display auto submit toolbar.
        """
        return pulumi.get(self, "auto_submit_toolbar")

    @property
    @pulumi.getter(name="defaultRelayState")
    def default_relay_state(self) -> Optional[str]:
        """
        Identifies a specific application resource in an IDP initiated SSO scenario.
        """
        return pulumi.get(self, "default_relay_state")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        description of application.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def destination(self) -> Optional[str]:
        """
        Identifies the location where the SAML response is intended to be sent inside of the SAML assertion.
        """
        return pulumi.get(self, "destination")

    @property
    @pulumi.getter(name="digestAlgorithm")
    def digest_algorithm(self) -> Optional[str]:
        """
        Determines the digest algorithm used to digitally sign the SAML assertion and response.
        """
        return pulumi.get(self, "digest_algorithm")

    @property
    @pulumi.getter
    def features(self) -> Optional[List[str]]:
        """
        features enabled.
        """
        return pulumi.get(self, "features")

    @property
    @pulumi.getter(name="hideIos")
    def hide_ios(self) -> Optional[bool]:
        """
        Do not display application icon on mobile app.
        """
        return pulumi.get(self, "hide_ios")

    @property
    @pulumi.getter(name="hideWeb")
    def hide_web(self) -> Optional[bool]:
        """
        Do not display application icon to users
        """
        return pulumi.get(self, "hide_web")

    @property
    @pulumi.getter(name="honorForceAuthn")
    def honor_force_authn(self) -> Optional[bool]:
        """
        Prompt user to re-authenticate if SP asks for it.
        """
        return pulumi.get(self, "honor_force_authn")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        id of application.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="idpIssuer")
    def idp_issuer(self) -> Optional[str]:
        """
        SAML issuer ID.
        """
        return pulumi.get(self, "idp_issuer")

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> str:
        """
        Certificate key ID.
        """
        return pulumi.get(self, "key_id")

    @property
    @pulumi.getter
    def label(self) -> Optional[str]:
        """
        label of application.
        """
        return pulumi.get(self, "label")

    @property
    @pulumi.getter(name="labelPrefix")
    def label_prefix(self) -> Optional[str]:
        return pulumi.get(self, "label_prefix")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        name of application.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def recipient(self) -> Optional[str]:
        """
        The location where the app may present the SAML assertion.
        """
        return pulumi.get(self, "recipient")

    @property
    @pulumi.getter(name="requestCompressed")
    def request_compressed(self) -> Optional[bool]:
        """
        Denotes whether the request is compressed or not.
        """
        return pulumi.get(self, "request_compressed")

    @property
    @pulumi.getter(name="responseSigned")
    def response_signed(self) -> Optional[bool]:
        """
        Determines whether the SAML auth response message is digitally signed.
        """
        return pulumi.get(self, "response_signed")

    @property
    @pulumi.getter(name="signatureAlgorithm")
    def signature_algorithm(self) -> Optional[str]:
        """
        Signature algorithm used ot digitally sign the assertion and response.
        """
        return pulumi.get(self, "signature_algorithm")

    @property
    @pulumi.getter(name="spIssuer")
    def sp_issuer(self) -> Optional[str]:
        """
        SAML service provider issuer.
        """
        return pulumi.get(self, "sp_issuer")

    @property
    @pulumi.getter(name="ssoUrl")
    def sso_url(self) -> Optional[str]:
        """
        Single Sign on Url.
        """
        return pulumi.get(self, "sso_url")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        status of application.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subjectNameIdFormat")
    def subject_name_id_format(self) -> Optional[str]:
        """
        Identifies the SAML processing rules.
        """
        return pulumi.get(self, "subject_name_id_format")

    @property
    @pulumi.getter(name="subjectNameIdTemplate")
    def subject_name_id_template(self) -> Optional[str]:
        """
        Template for app user's username when a user is assigned to the app.
        """
        return pulumi.get(self, "subject_name_id_template")

    @property
    @pulumi.getter(name="userNameTemplate")
    def user_name_template(self) -> Optional[str]:
        """
        Username template.
        """
        return pulumi.get(self, "user_name_template")

    @property
    @pulumi.getter(name="userNameTemplateSuffix")
    def user_name_template_suffix(self) -> Optional[str]:
        """
        Username template suffix.
        """
        return pulumi.get(self, "user_name_template_suffix")

    @property
    @pulumi.getter(name="userNameTemplateType")
    def user_name_template_type(self) -> Optional[str]:
        """
        Username template type.
        """
        return pulumi.get(self, "user_name_template_type")


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


def get_saml(accessibility_error_redirect_url: Optional[str] = None,
             accessibility_login_redirect_url: Optional[str] = None,
             accessibility_self_service: Optional[bool] = None,
             active_only: Optional[bool] = None,
             app_settings_json: Optional[str] = None,
             assertion_signed: Optional[bool] = None,
             attribute_statements: Optional[List[pulumi.InputType['GetSamlAttributeStatementArgs']]] = None,
             audience: Optional[str] = None,
             authn_context_class_ref: Optional[str] = None,
             auto_submit_toolbar: Optional[bool] = None,
             default_relay_state: Optional[str] = None,
             destination: Optional[str] = None,
             digest_algorithm: Optional[str] = None,
             features: Optional[List[str]] = None,
             hide_ios: Optional[bool] = None,
             hide_web: Optional[bool] = None,
             honor_force_authn: Optional[bool] = None,
             id: Optional[str] = None,
             idp_issuer: Optional[str] = None,
             label: Optional[str] = None,
             label_prefix: Optional[str] = None,
             recipient: Optional[str] = None,
             request_compressed: Optional[bool] = None,
             response_signed: Optional[bool] = None,
             signature_algorithm: Optional[str] = None,
             sp_issuer: Optional[str] = None,
             sso_url: Optional[str] = None,
             subject_name_id_format: Optional[str] = None,
             subject_name_id_template: Optional[str] = None,
             user_name_template: Optional[str] = None,
             user_name_template_suffix: Optional[str] = None,
             user_name_template_type: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSamlResult:
    """
    Use this data source to retrieve the collaborators for a given repository.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_okta as okta

    example = okta.app.get_saml(label="Example App")
    ```


    :param str accessibility_error_redirect_url: Custom error page URL.
    :param str accessibility_login_redirect_url: Custom login page URL.
    :param bool accessibility_self_service: Enable self service.
    :param bool active_only: tells the provider to query for only `ACTIVE` applications.
    :param str app_settings_json: Application settings in JSON format.
    :param bool assertion_signed: Determines whether the SAML assertion is digitally signed.
    :param List[pulumi.InputType['GetSamlAttributeStatementArgs']] attribute_statements: SAML Attribute statements.
    :param str audience: Audience restriction.
    :param str authn_context_class_ref: Identifies the SAML authentication context class for the assertion’s authentication statement.
    :param bool auto_submit_toolbar: Display auto submit toolbar.
    :param str default_relay_state: Identifies a specific application resource in an IDP initiated SSO scenario.
    :param str destination: Identifies the location where the SAML response is intended to be sent inside of the SAML assertion.
    :param str digest_algorithm: Determines the digest algorithm used to digitally sign the SAML assertion and response.
    :param List[str] features: features enabled.
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
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('okta:app/getSaml:getSaml', __args__, opts=opts, typ=GetSamlResult).value

    return AwaitableGetSamlResult(
        accessibility_error_redirect_url=__ret__.accessibility_error_redirect_url,
        accessibility_login_redirect_url=__ret__.accessibility_login_redirect_url,
        accessibility_self_service=__ret__.accessibility_self_service,
        active_only=__ret__.active_only,
        app_settings_json=__ret__.app_settings_json,
        assertion_signed=__ret__.assertion_signed,
        attribute_statements=__ret__.attribute_statements,
        audience=__ret__.audience,
        authn_context_class_ref=__ret__.authn_context_class_ref,
        auto_submit_toolbar=__ret__.auto_submit_toolbar,
        default_relay_state=__ret__.default_relay_state,
        description=__ret__.description,
        destination=__ret__.destination,
        digest_algorithm=__ret__.digest_algorithm,
        features=__ret__.features,
        hide_ios=__ret__.hide_ios,
        hide_web=__ret__.hide_web,
        honor_force_authn=__ret__.honor_force_authn,
        id=__ret__.id,
        idp_issuer=__ret__.idp_issuer,
        key_id=__ret__.key_id,
        label=__ret__.label,
        label_prefix=__ret__.label_prefix,
        name=__ret__.name,
        recipient=__ret__.recipient,
        request_compressed=__ret__.request_compressed,
        response_signed=__ret__.response_signed,
        signature_algorithm=__ret__.signature_algorithm,
        sp_issuer=__ret__.sp_issuer,
        sso_url=__ret__.sso_url,
        status=__ret__.status,
        subject_name_id_format=__ret__.subject_name_id_format,
        subject_name_id_template=__ret__.subject_name_id_template,
        user_name_template=__ret__.user_name_template,
        user_name_template_suffix=__ret__.user_name_template_suffix,
        user_name_template_type=__ret__.user_name_template_type)
