// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Gets information about a public gateway PAT rule. For further information please check the
 * API [documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#get-8faeea)
 */
export function getVpcPublicGatewayPatRule(args: GetVpcPublicGatewayPatRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcPublicGatewayPatRuleResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("scaleway:index/getVpcPublicGatewayPatRule:getVpcPublicGatewayPatRule", {
        "patRuleId": args.patRuleId,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcPublicGatewayPatRule.
 */
export interface GetVpcPublicGatewayPatRuleArgs {
    /**
     * The ID of the PAT rule to retrieve
     */
    patRuleId: string;
    /**
     * `zone`) The zone in which
     * the image exists.
     */
    zone?: string;
}

/**
 * A collection of values returned by getVpcPublicGatewayPatRule.
 */
export interface GetVpcPublicGatewayPatRuleResult {
    readonly createdAt: string;
    /**
     * The ID of the public gateway.
     */
    readonly gatewayId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly organizationId: string;
    readonly patRuleId: string;
    /**
     * The Private IP to forward data to (IP address).
     */
    readonly privateIp: string;
    /**
     * The Private port to translate to.
     */
    readonly privatePort: number;
    /**
     * The Protocol the rule should apply to. Possible values are both, tcp and udp.
     */
    readonly protocol: string;
    /**
     * The Public port to listen on.
     */
    readonly publicPort: number;
    readonly updatedAt: string;
    readonly zone?: string;
}

export function getVpcPublicGatewayPatRuleOutput(args: GetVpcPublicGatewayPatRuleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVpcPublicGatewayPatRuleResult> {
    return pulumi.output(args).apply(a => getVpcPublicGatewayPatRule(a, opts))
}

/**
 * A collection of arguments for invoking getVpcPublicGatewayPatRule.
 */
export interface GetVpcPublicGatewayPatRuleOutputArgs {
    /**
     * The ID of the PAT rule to retrieve
     */
    patRuleId: pulumi.Input<string>;
    /**
     * `zone`) The zone in which
     * the image exists.
     */
    zone?: pulumi.Input<string>;
}
