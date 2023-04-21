// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Gets information about multiple Load Balancer Frontends.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * // Find frontends that share the same LB ID
 * const byLBID = scaleway_lb_lb01.id.apply(id => scaleway.getLbFrontends({
 *     lbId: id,
 * }));
 * // Find frontends by LB ID and name
 * const byLBIDAndName = scaleway_lb_lb01.id.apply(id => scaleway.getLbFrontends({
 *     lbId: id,
 *     name: "tf-frontend-datasource",
 * }));
 * ```
 */
export function getLbFrontends(args: GetLbFrontendsArgs, opts?: pulumi.InvokeOptions): Promise<GetLbFrontendsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("scaleway:index/getLbFrontends:getLbFrontends", {
        "lbId": args.lbId,
        "name": args.name,
        "projectId": args.projectId,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getLbFrontends.
 */
export interface GetLbFrontendsArgs {
    /**
     * The load-balancer ID this frontend is attached to. frontends with a LB ID like it are listed.
     */
    lbId: string;
    /**
     * The frontend name used as filter. Frontends with a name like it are listed.
     */
    name?: string;
    projectId?: string;
    /**
     * `zone`) The zone in which frontends exist.
     */
    zone?: string;
}

/**
 * A collection of values returned by getLbFrontends.
 */
export interface GetLbFrontendsResult {
    /**
     * List of found frontends
     */
    readonly frontends: outputs.GetLbFrontendsFrontend[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly lbId: string;
    readonly name?: string;
    readonly organizationId: string;
    readonly projectId: string;
    readonly zone: string;
}

export function getLbFrontendsOutput(args: GetLbFrontendsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLbFrontendsResult> {
    return pulumi.output(args).apply(a => getLbFrontends(a, opts))
}

/**
 * A collection of arguments for invoking getLbFrontends.
 */
export interface GetLbFrontendsOutputArgs {
    /**
     * The load-balancer ID this frontend is attached to. frontends with a LB ID like it are listed.
     */
    lbId: pulumi.Input<string>;
    /**
     * The frontend name used as filter. Frontends with a name like it are listed.
     */
    name?: pulumi.Input<string>;
    projectId?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which frontends exist.
     */
    zone?: pulumi.Input<string>;
}
