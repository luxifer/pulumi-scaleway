// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Gets information about the Scaleway Cockpit.
 *
 * For more information consult the [documentation](https://www.scaleway.com/en/docs/observability/cockpit/concepts/).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * // Get default project's cockpit
 * const main = pulumi.output(scaleway.getCockpit());
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * // Get a specific project's cockpit
 * const main = pulumi.output(scaleway.getCockpit({
 *     projectId: "11111111-1111-1111-1111-111111111111",
 * }));
 * ```
 */
export function getCockpit(args?: GetCockpitArgs, opts?: pulumi.InvokeOptions): Promise<GetCockpitResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("scaleway:index/getCockpit:getCockpit", {
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getCockpit.
 */
export interface GetCockpitArgs {
    /**
     * `projectId`) The ID of the project the cockpit is associated with.
     */
    projectId?: string;
}

/**
 * A collection of values returned by getCockpit.
 */
export interface GetCockpitResult {
    /**
     * Endpoints
     */
    readonly endpoints: outputs.GetCockpitEndpoint[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly projectId?: string;
}

export function getCockpitOutput(args?: GetCockpitOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCockpitResult> {
    return pulumi.output(args).apply(a => getCockpit(a, opts))
}

/**
 * A collection of arguments for invoking getCockpit.
 */
export interface GetCockpitOutputArgs {
    /**
     * `projectId`) The ID of the project the cockpit is associated with.
     */
    projectId?: pulumi.Input<string>;
}