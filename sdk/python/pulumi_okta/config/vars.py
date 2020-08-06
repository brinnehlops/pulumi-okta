# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables

__config__ = pulumi.Config('okta')

api_token = __config__.get('apiToken') or _utilities.get_env('OKTA_API_TOKEN')
"""
API Token granting privileges to Okta API.
"""

backoff = __config__.get('backoff')
"""
Use exponential back off strategy for rate limits.
"""

base_url = __config__.get('baseUrl') or _utilities.get_env('OKTA_BASE_URL')
"""
The Okta url. (Use 'oktapreview.com' for Okta testing)
"""

max_retries = __config__.get('maxRetries')
"""
maximum number of retries to attempt before erroring out.
"""

max_wait_seconds = __config__.get('maxWaitSeconds')
"""
maximum seconds to wait when rate limit is hit. We use exponential backoffs when backoff is enabled.
"""

min_wait_seconds = __config__.get('minWaitSeconds')
"""
minimum seconds to wait when rate limit is hit. We use exponential backoffs when backoff is enabled.
"""

org_name = __config__.get('orgName') or _utilities.get_env('OKTA_ORG_NAME')
"""
The organization to manage in Okta.
"""

parallelism = __config__.get('parallelism')
"""
Number of concurrent requests to make within a resource where bulk operations are not possible. Take note of
https://developer.okta.com/docs/api/getting_started/rate-limits.
"""

