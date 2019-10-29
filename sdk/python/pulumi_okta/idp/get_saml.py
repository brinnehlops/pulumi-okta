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
    def __init__(__self__, acs_binding=None, acs_type=None, audience=None, id=None, issuer=None, issuer_mode=None, kid=None, name=None, sso_binding=None, sso_destination=None, sso_url=None, subject_filter=None, subject_formats=None, type=None):
        if acs_binding and not isinstance(acs_binding, str):
            raise TypeError("Expected argument 'acs_binding' to be a str")
        __self__.acs_binding = acs_binding
        """
        HTTP binding used to receive a SAMLResponse message from the IdP.
        """
        if acs_type and not isinstance(acs_type, str):
            raise TypeError("Expected argument 'acs_type' to be a str")
        __self__.acs_type = acs_type
        """
        Determines whether to publish an instance-specific (trust) or organization (shared) ACS endpoint in the SAML metadata.
        """
        if audience and not isinstance(audience, str):
            raise TypeError("Expected argument 'audience' to be a str")
        __self__.audience = audience
        """
        URI that identifies the target Okta IdP instance (SP)
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id of idp.
        """
        if issuer and not isinstance(issuer, str):
            raise TypeError("Expected argument 'issuer' to be a str")
        __self__.issuer = issuer
        """
        URI that identifies the issuer (IdP).
        """
        if issuer_mode and not isinstance(issuer_mode, str):
            raise TypeError("Expected argument 'issuer_mode' to be a str")
        __self__.issuer_mode = issuer_mode
        """
        indicates whether Okta uses the original Okta org domain URL, or a custom domain URL in the request to the IdP.
        """
        if kid and not isinstance(kid, str):
            raise TypeError("Expected argument 'kid' to be a str")
        __self__.kid = kid
        """
        Key ID reference to the IdP's X.509 signature certificate.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        name of the idp.
        """
        if sso_binding and not isinstance(sso_binding, str):
            raise TypeError("Expected argument 'sso_binding' to be a str")
        __self__.sso_binding = sso_binding
        """
        single sign on binding.
        """
        if sso_destination and not isinstance(sso_destination, str):
            raise TypeError("Expected argument 'sso_destination' to be a str")
        __self__.sso_destination = sso_destination
        """
        SSO request binding, HTTP-POST or HTTP-REDIRECT.
        """
        if sso_url and not isinstance(sso_url, str):
            raise TypeError("Expected argument 'sso_url' to be a str")
        __self__.sso_url = sso_url
        """
        single sign on url.
        """
        if subject_filter and not isinstance(subject_filter, str):
            raise TypeError("Expected argument 'subject_filter' to be a str")
        __self__.subject_filter = subject_filter
        """
        regular expression pattern used to filter untrusted IdP usernames.
        """
        if subject_formats and not isinstance(subject_formats, list):
            raise TypeError("Expected argument 'subject_formats' to be a list")
        __self__.subject_formats = subject_formats
        """
        Expression to generate or transform a unique username for the IdP user.
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        type of idp.
        """
class AwaitableGetSamlResult(GetSamlResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSamlResult(
            acs_binding=self.acs_binding,
            acs_type=self.acs_type,
            audience=self.audience,
            id=self.id,
            issuer=self.issuer,
            issuer_mode=self.issuer_mode,
            kid=self.kid,
            name=self.name,
            sso_binding=self.sso_binding,
            sso_destination=self.sso_destination,
            sso_url=self.sso_url,
            subject_filter=self.subject_filter,
            subject_formats=self.subject_formats,
            type=self.type)

def get_saml(id=None,name=None,opts=None):
    """
    Use this data source to retrieve a SAML IdP from Okta.
    
    :param str id: The id of the idp to retrieve, conflicts with `name`.
    :param str name: The name of the idp to retrieve, conflicts with `id`.

    > This content is derived from https://github.com/articulate/terraform-provider-okta/blob/master/website/docs/d/idp_saml.html.markdown.
    """
    __args__ = dict()

    __args__['id'] = id
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('okta:idp/getSaml:getSaml', __args__, opts=opts).value

    return AwaitableGetSamlResult(
        acs_binding=__ret__.get('acsBinding'),
        acs_type=__ret__.get('acsType'),
        audience=__ret__.get('audience'),
        id=__ret__.get('id'),
        issuer=__ret__.get('issuer'),
        issuer_mode=__ret__.get('issuerMode'),
        kid=__ret__.get('kid'),
        name=__ret__.get('name'),
        sso_binding=__ret__.get('ssoBinding'),
        sso_destination=__ret__.get('ssoDestination'),
        sso_url=__ret__.get('ssoUrl'),
        subject_filter=__ret__.get('subjectFilter'),
        subject_formats=__ret__.get('subjectFormats'),
        type=__ret__.get('type'))
