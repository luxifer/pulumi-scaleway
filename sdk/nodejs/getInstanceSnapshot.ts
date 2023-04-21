// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Gets information about an instance snapshot.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * // Get info by snapshot name
 * const byName = pulumi.output(scaleway.getInstanceSnapshot({
 *     name: "my-snapshot-name",
 * }));
 * // Get info by snapshot ID
 * const byId = pulumi.output(scaleway.getInstanceSnapshot({
 *     snapshotId: "11111111-1111-1111-1111-111111111111",
 * }));
 * ```
 */
export function getInstanceSnapshot(args?: GetInstanceSnapshotArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceSnapshotResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("scaleway:index/getInstanceSnapshot:getInstanceSnapshot", {
        "name": args.name,
        "snapshotId": args.snapshotId,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceSnapshot.
 */
export interface GetInstanceSnapshotArgs {
    /**
     * The snapshot name.
     * Only one of `name` and `snapshotId` should be specified.
     */
    name?: string;
    /**
     * The snapshot id.
     * Only one of `name` and `snapshotId` should be specified.
     */
    snapshotId?: string;
    /**
     * `zone`) The zone in which the snapshot exists.
     */
    zone?: string;
}

/**
 * A collection of values returned by getInstanceSnapshot.
 */
export interface GetInstanceSnapshotResult {
    readonly createdAt: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly imports: outputs.GetInstanceSnapshotImport[];
    readonly name?: string;
    readonly organizationId: string;
    readonly projectId: string;
    readonly sizeInGb: number;
    readonly snapshotId?: string;
    readonly tags: string[];
    readonly type: string;
    readonly volumeId: string;
    readonly zone?: string;
}

export function getInstanceSnapshotOutput(args?: GetInstanceSnapshotOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceSnapshotResult> {
    return pulumi.output(args).apply(a => getInstanceSnapshot(a, opts))
}

/**
 * A collection of arguments for invoking getInstanceSnapshot.
 */
export interface GetInstanceSnapshotOutputArgs {
    /**
     * The snapshot name.
     * Only one of `name` and `snapshotId` should be specified.
     */
    name?: pulumi.Input<string>;
    /**
     * The snapshot id.
     * Only one of `name` and `snapshotId` should be specified.
     */
    snapshotId?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the snapshot exists.
     */
    zone?: pulumi.Input<string>;
}
