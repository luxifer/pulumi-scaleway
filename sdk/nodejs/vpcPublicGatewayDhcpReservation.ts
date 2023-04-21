// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates and manages the [Scaleway DHCP Reservations](https://www.scaleway.com/en/docs/network/vpc/concepts/#dhcp).
 *
 * The static associations are used to assign IP addresses based on the MAC addresses of the Instance.
 *
 * Statically assigned IP addresses should fall within the configured subnet, but be outside of the dynamic range.
 *
 * For more information, see [the documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#dhcp-c05544) and [configuration guide](https://www.scaleway.com/en/docs/network/vpc/how-to/configure-a-public-gateway/#how-to-review-and-configure-dhcp).
 *
 * [DHCP reservations](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#dhcp-entries-e40fb6) hold both dynamic DHCP leases (IP addresses dynamically assigned by the gateway to instances) and static user-created DHCP reservations.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi_scaleway from "@luxifer/pulumi-scaleway";
 *
 * const mainVpcPrivateNetwork = new scaleway.VpcPrivateNetwork("mainVpcPrivateNetwork", {});
 * const mainInstanceServer = new scaleway.InstanceServer("mainInstanceServer", {
 *     image: "ubuntu_jammy",
 *     type: "DEV1-S",
 *     zone: "fr-par-1",
 *     privateNetworks: [{
 *         pnId: mainVpcPrivateNetwork.id,
 *     }],
 * });
 * const mainVpcPublicGatewayIp = new scaleway.VpcPublicGatewayIp("mainVpcPublicGatewayIp", {});
 * const mainVpcPublicGatewayDhcp = new scaleway.VpcPublicGatewayDhcp("mainVpcPublicGatewayDhcp", {subnet: "192.168.1.0/24"});
 * const mainVpcPublicGateway = new scaleway.VpcPublicGateway("mainVpcPublicGateway", {
 *     type: "VPC-GW-S",
 *     ipId: mainVpcPublicGatewayIp.id,
 * });
 * const mainVpcGatewayNetwork = new scaleway.VpcGatewayNetwork("mainVpcGatewayNetwork", {
 *     gatewayId: mainVpcPublicGateway.id,
 *     privateNetworkId: mainVpcPrivateNetwork.id,
 *     dhcpId: mainVpcPublicGatewayDhcp.id,
 *     cleanupDhcp: true,
 *     enableMasquerade: true,
 * }, {
 *     dependsOn: [
 *         mainVpcPublicGatewayIp,
 *         mainVpcPrivateNetwork,
 *     ],
 * });
 * const mainVpcPublicGatewayDhcpReservation = new scaleway.VpcPublicGatewayDhcpReservation("mainVpcPublicGatewayDhcpReservation", {
 *     gatewayNetworkId: mainVpcGatewayNetwork.id,
 *     macAddress: mainInstanceServer.privateNetworks.apply(privateNetworks => privateNetworks?[0]?.macAddress),
 *     ipAddress: "192.168.1.1",
 * });
 * ```
 *
 * ## Import
 *
 * Public gateway DHCP Reservation config can be imported using the `{zone}/{id}`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:index/vpcPublicGatewayDhcpReservation:VpcPublicGatewayDhcpReservation main fr-par-1/11111111-1111-1111-1111-111111111111
 * ```
 */
export class VpcPublicGatewayDhcpReservation extends pulumi.CustomResource {
    /**
     * Get an existing VpcPublicGatewayDhcpReservation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcPublicGatewayDhcpReservationState, opts?: pulumi.CustomResourceOptions): VpcPublicGatewayDhcpReservation {
        return new VpcPublicGatewayDhcpReservation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/vpcPublicGatewayDhcpReservation:VpcPublicGatewayDhcpReservation';

    /**
     * Returns true if the given object is an instance of VpcPublicGatewayDhcpReservation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcPublicGatewayDhcpReservation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcPublicGatewayDhcpReservation.__pulumiType;
    }

    /**
     * The date and time of the creation of the public gateway DHCP config.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The ID of the owning GatewayNetwork.
     */
    public readonly gatewayNetworkId!: pulumi.Output<string>;
    /**
     * The Hostname of the client machine.
     */
    public /*out*/ readonly hostname!: pulumi.Output<string>;
    /**
     * The IP address to give to the machine (IP address).
     */
    public readonly ipAddress!: pulumi.Output<string>;
    /**
     * The MAC address to give a static entry to.
     */
    public readonly macAddress!: pulumi.Output<string>;
    /**
     * The reservation type, either static (DHCP reservation) or dynamic (DHCP lease). Possible values are reservation and lease.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The date and time of the last update of the public gateway DHCP config.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * `zone`) The zone in which the public gateway DHCP config should be created.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a VpcPublicGatewayDhcpReservation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcPublicGatewayDhcpReservationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcPublicGatewayDhcpReservationArgs | VpcPublicGatewayDhcpReservationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcPublicGatewayDhcpReservationState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["gatewayNetworkId"] = state ? state.gatewayNetworkId : undefined;
            resourceInputs["hostname"] = state ? state.hostname : undefined;
            resourceInputs["ipAddress"] = state ? state.ipAddress : undefined;
            resourceInputs["macAddress"] = state ? state.macAddress : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as VpcPublicGatewayDhcpReservationArgs | undefined;
            if ((!args || args.gatewayNetworkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gatewayNetworkId'");
            }
            if ((!args || args.ipAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipAddress'");
            }
            if ((!args || args.macAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'macAddress'");
            }
            resourceInputs["gatewayNetworkId"] = args ? args.gatewayNetworkId : undefined;
            resourceInputs["ipAddress"] = args ? args.ipAddress : undefined;
            resourceInputs["macAddress"] = args ? args.macAddress : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["hostname"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpcPublicGatewayDhcpReservation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcPublicGatewayDhcpReservation resources.
 */
export interface VpcPublicGatewayDhcpReservationState {
    /**
     * The date and time of the creation of the public gateway DHCP config.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The ID of the owning GatewayNetwork.
     */
    gatewayNetworkId?: pulumi.Input<string>;
    /**
     * The Hostname of the client machine.
     */
    hostname?: pulumi.Input<string>;
    /**
     * The IP address to give to the machine (IP address).
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * The MAC address to give a static entry to.
     */
    macAddress?: pulumi.Input<string>;
    /**
     * The reservation type, either static (DHCP reservation) or dynamic (DHCP lease). Possible values are reservation and lease.
     */
    type?: pulumi.Input<string>;
    /**
     * The date and time of the last update of the public gateway DHCP config.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the public gateway DHCP config should be created.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcPublicGatewayDhcpReservation resource.
 */
export interface VpcPublicGatewayDhcpReservationArgs {
    /**
     * The ID of the owning GatewayNetwork.
     */
    gatewayNetworkId: pulumi.Input<string>;
    /**
     * The IP address to give to the machine (IP address).
     */
    ipAddress: pulumi.Input<string>;
    /**
     * The MAC address to give a static entry to.
     */
    macAddress: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the public gateway DHCP config should be created.
     */
    zone?: pulumi.Input<string>;
}