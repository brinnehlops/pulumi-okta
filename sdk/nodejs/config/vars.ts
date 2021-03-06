// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

let __config = new pulumi.Config("okta");

/**
 * API Token granting privileges to Okta API.
 */
export let apiToken: string | undefined = __config.get("apiToken") || utilities.getEnv("OKTA_API_TOKEN");
/**
 * Use exponential back off strategy for rate limits.
 */
export let backoff: boolean | undefined = __config.getObject<boolean>("backoff");
/**
 * The Okta url. (Use 'oktapreview.com' for Okta testing)
 */
export let baseUrl: string | undefined = __config.get("baseUrl") || utilities.getEnv("OKTA_BASE_URL");
/**
 * maximum number of retries to attempt before erroring out.
 */
export let maxRetries: number | undefined = __config.getObject<number>("maxRetries");
/**
 * maximum seconds to wait when rate limit is hit. We use exponential backoffs when backoff is enabled.
 */
export let maxWaitSeconds: number | undefined = __config.getObject<number>("maxWaitSeconds");
/**
 * minimum seconds to wait when rate limit is hit. We use exponential backoffs when backoff is enabled.
 */
export let minWaitSeconds: number | undefined = __config.getObject<number>("minWaitSeconds");
/**
 * The organization to manage in Okta.
 */
export let orgName: string | undefined = __config.get("orgName") || utilities.getEnv("OKTA_ORG_NAME");
/**
 * Number of concurrent requests to make within a resource where bulk operations are not possible. Take note of
 * https://developer.okta.com/docs/api/getting_started/rate-limits.
 */
export let parallelism: number | undefined = __config.getObject<number>("parallelism");
