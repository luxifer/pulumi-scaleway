// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * Credential can be imported using the `{region}/{id}`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:index/mnqCredential:MnqCredential main fr-par/11111111111111111111111111111111
 * ```
 */
export class MnqCredential extends pulumi.CustomResource {
    /**
     * Get an existing MnqCredential resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MnqCredentialState, opts?: pulumi.CustomResourceOptions): MnqCredential {
        return new MnqCredential(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/mnqCredential:MnqCredential';

    /**
     * Returns true if the given object is an instance of MnqCredential.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MnqCredential {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MnqCredential.__pulumiType;
    }

    /**
     * The credential name..
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The namespace containing the Credential.
     */
    public readonly namespaceId!: pulumi.Output<string>;
    /**
     * Credentials file used to connect to the NATS service.
     */
    public readonly natsCredentials!: pulumi.Output<outputs.MnqCredentialNatsCredentials>;
    /**
     * The protocol associated to the Credential. Possible values are `nats` and `sqsSns`.
     */
    public /*out*/ readonly protocol!: pulumi.Output<string>;
    /**
     * (Defaults to provider `region`). The region
     * in which the namespace should be created.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Credential used to connect to the SQS/SNS service.
     */
    public readonly sqsSnsCredentials!: pulumi.Output<outputs.MnqCredentialSqsSnsCredentials | undefined>;

    /**
     * Create a MnqCredential resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MnqCredentialArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MnqCredentialArgs | MnqCredentialState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MnqCredentialState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespaceId"] = state ? state.namespaceId : undefined;
            resourceInputs["natsCredentials"] = state ? state.natsCredentials : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["sqsSnsCredentials"] = state ? state.sqsSnsCredentials : undefined;
        } else {
            const args = argsOrState as MnqCredentialArgs | undefined;
            if ((!args || args.namespaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'namespaceId'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespaceId"] = args ? args.namespaceId : undefined;
            resourceInputs["natsCredentials"] = args ? args.natsCredentials : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["sqsSnsCredentials"] = args ? args.sqsSnsCredentials : undefined;
            resourceInputs["protocol"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MnqCredential.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MnqCredential resources.
 */
export interface MnqCredentialState {
    /**
     * The credential name..
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace containing the Credential.
     */
    namespaceId?: pulumi.Input<string>;
    /**
     * Credentials file used to connect to the NATS service.
     */
    natsCredentials?: pulumi.Input<inputs.MnqCredentialNatsCredentials>;
    /**
     * The protocol associated to the Credential. Possible values are `nats` and `sqsSns`.
     */
    protocol?: pulumi.Input<string>;
    /**
     * (Defaults to provider `region`). The region
     * in which the namespace should be created.
     */
    region?: pulumi.Input<string>;
    /**
     * Credential used to connect to the SQS/SNS service.
     */
    sqsSnsCredentials?: pulumi.Input<inputs.MnqCredentialSqsSnsCredentials>;
}

/**
 * The set of arguments for constructing a MnqCredential resource.
 */
export interface MnqCredentialArgs {
    /**
     * The credential name..
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace containing the Credential.
     */
    namespaceId: pulumi.Input<string>;
    /**
     * Credentials file used to connect to the NATS service.
     */
    natsCredentials?: pulumi.Input<inputs.MnqCredentialNatsCredentials>;
    /**
     * (Defaults to provider `region`). The region
     * in which the namespace should be created.
     */
    region?: pulumi.Input<string>;
    /**
     * Credential used to connect to the SQS/SNS service.
     */
    sqsSnsCredentials?: pulumi.Input<inputs.MnqCredentialSqsSnsCredentials>;
}
