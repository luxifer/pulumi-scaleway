// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway IAM SSH Keys.
 * For more information,
 * see [the documentation](https://developers.scaleway.com/en/products/iam/api/v1alpha1/#ssh-keys-d8ccd4).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = new scaleway.IamSshKey("main", {
 *     publicKey: "<YOUR-PUBLIC-SSH-KEY>",
 * });
 * ```
 *
 * ## Import
 *
 * SSH keys can be imported using the `id`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:index/iamSshKey:IamSshKey main 11111111-1111-1111-1111-111111111111
 * ```
 */
export class IamSshKey extends pulumi.CustomResource {
    /**
     * Get an existing IamSshKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IamSshKeyState, opts?: pulumi.CustomResourceOptions): IamSshKey {
        return new IamSshKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/iamSshKey:IamSshKey';

    /**
     * Returns true if the given object is an instance of IamSshKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IamSshKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IamSshKey.__pulumiType;
    }

    /**
     * The date and time of the creation of the SSH key.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The SSH key status.
     */
    public readonly disabled!: pulumi.Output<boolean | undefined>;
    /**
     * The fingerprint of the iam SSH key.
     */
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * The name of the SSH key.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the organization the SSH key is associated with.
     */
    public /*out*/ readonly organizationId!: pulumi.Output<string>;
    /**
     * `projectId`) The ID of the project the SSH key is
     * associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The public SSH key to be added.
     */
    public readonly publicKey!: pulumi.Output<string>;
    /**
     * The date and time of the last update of the SSH key.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a IamSshKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IamSshKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IamSshKeyArgs | IamSshKeyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IamSshKeyState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["disabled"] = state ? state.disabled : undefined;
            resourceInputs["fingerprint"] = state ? state.fingerprint : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["publicKey"] = state ? state.publicKey : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as IamSshKeyArgs | undefined;
            if ((!args || args.publicKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'publicKey'");
            }
            resourceInputs["disabled"] = args ? args.disabled : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["publicKey"] = args ? args.publicKey : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["fingerprint"] = undefined /*out*/;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IamSshKey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IamSshKey resources.
 */
export interface IamSshKeyState {
    /**
     * The date and time of the creation of the SSH key.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The SSH key status.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * The fingerprint of the iam SSH key.
     */
    fingerprint?: pulumi.Input<string>;
    /**
     * The name of the SSH key.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the organization the SSH key is associated with.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the SSH key is
     * associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The public SSH key to be added.
     */
    publicKey?: pulumi.Input<string>;
    /**
     * The date and time of the last update of the SSH key.
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IamSshKey resource.
 */
export interface IamSshKeyArgs {
    /**
     * The SSH key status.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * The name of the SSH key.
     */
    name?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the SSH key is
     * associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The public SSH key to be added.
     */
    publicKey: pulumi.Input<string>;
}
