// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Gets information about the RDB instance network Access Control List.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * // Get the database ACL for the instance id 11111111-1111-1111-1111-111111111111 located in the default region e.g: fr-par
 * const myAcl = pulumi.output(scaleway.getRdbAcl({
 *     instanceId: "11111111-1111-1111-1111-111111111111",
 * }));
 * ```
 */
export function getRdbAcl(args: GetRdbAclArgs, opts?: pulumi.InvokeOptions): Promise<GetRdbAclResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("scaleway:index/getRdbAcl:getRdbAcl", {
        "instanceId": args.instanceId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getRdbAcl.
 */
export interface GetRdbAclArgs {
    /**
     * The RDB instance ID.
     */
    instanceId: string;
    /**
     * `region`) The region in which the Database Instance should be created.
     */
    region?: string;
}

/**
 * A collection of values returned by getRdbAcl.
 */
export interface GetRdbAclResult {
    /**
     * A list of ACLs rules (structure is described below)
     */
    readonly aclRules: outputs.GetRdbAclAclRule[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly region?: string;
}

export function getRdbAclOutput(args: GetRdbAclOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRdbAclResult> {
    return pulumi.output(args).apply(a => getRdbAcl(a, opts))
}

/**
 * A collection of arguments for invoking getRdbAcl.
 */
export interface GetRdbAclOutputArgs {
    /**
     * The RDB instance ID.
     */
    instanceId: pulumi.Input<string>;
    /**
     * `region`) The region in which the Database Instance should be created.
     */
    region?: pulumi.Input<string>;
}
