# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetMetadataSamlResult:
    """
    A collection of values returned by getMetadataSaml.
    """
    def __init__(__self__, app_id=None, certificate=None, entity_id=None, http_post_binding=None, http_redirect_binding=None, id=None, key_id=None, metadata=None, want_authn_requests_signed=None):
        if app_id and not isinstance(app_id, str):
            raise TypeError("Expected argument 'app_id' to be a str")
        __self__.app_id = app_id
        if certificate and not isinstance(certificate, str):
            raise TypeError("Expected argument 'certificate' to be a str")
        __self__.certificate = certificate
        """
        public certificate from application metadata.
        """
        if entity_id and not isinstance(entity_id, str):
            raise TypeError("Expected argument 'entity_id' to be a str")
        __self__.entity_id = entity_id
        """
        Entity URL for instance `https://www.okta.com/saml2/service-provider/sposcfdmlybtwkdcgtuf`.
        """
        if http_post_binding and not isinstance(http_post_binding, str):
            raise TypeError("Expected argument 'http_post_binding' to be a str")
        __self__.http_post_binding = http_post_binding
        """
        urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Post location from the SAML metadata.
        """
        if http_redirect_binding and not isinstance(http_redirect_binding, str):
            raise TypeError("Expected argument 'http_redirect_binding' to be a str")
        __self__.http_redirect_binding = http_redirect_binding
        """
        urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect location from the SAML metadata.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if key_id and not isinstance(key_id, str):
            raise TypeError("Expected argument 'key_id' to be a str")
        __self__.key_id = key_id
        if metadata and not isinstance(metadata, str):
            raise TypeError("Expected argument 'metadata' to be a str")
        __self__.metadata = metadata
        """
        raw metadata of application.
        """
        if want_authn_requests_signed and not isinstance(want_authn_requests_signed, bool):
            raise TypeError("Expected argument 'want_authn_requests_signed' to be a bool")
        __self__.want_authn_requests_signed = want_authn_requests_signed
        """
        Whether authn requests are signed.
        """
class AwaitableGetMetadataSamlResult(GetMetadataSamlResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMetadataSamlResult(
            app_id=self.app_id,
            certificate=self.certificate,
            entity_id=self.entity_id,
            http_post_binding=self.http_post_binding,
            http_redirect_binding=self.http_redirect_binding,
            id=self.id,
            key_id=self.key_id,
            metadata=self.metadata,
            want_authn_requests_signed=self.want_authn_requests_signed)

def get_metadata_saml(app_id=None,key_id=None,opts=None):
    """
    Use this data source to retrieve the collaborators for a given repository.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_okta as okta

    example = okta.app.get_metadata_saml(app_id="<app id>",
        key_id="<cert key id>")
    ```


    :param str app_id: The application ID.
    :param str key_id: Certificate Key ID.
    """
    __args__ = dict()


    __args__['appId'] = app_id
    __args__['keyId'] = key_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('okta:app/getMetadataSaml:getMetadataSaml', __args__, opts=opts).value

    return AwaitableGetMetadataSamlResult(
        app_id=__ret__.get('appId'),
        certificate=__ret__.get('certificate'),
        entity_id=__ret__.get('entityId'),
        http_post_binding=__ret__.get('httpPostBinding'),
        http_redirect_binding=__ret__.get('httpRedirectBinding'),
        id=__ret__.get('id'),
        key_id=__ret__.get('keyId'),
        metadata=__ret__.get('metadata'),
        want_authn_requests_signed=__ret__.get('wantAuthnRequestsSigned'))
