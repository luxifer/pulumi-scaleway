// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Serverless Container Namespace.
 * For more information see [the documentation](https://developers.scaleway.com/en/products/containers/api/#namespaces-cdce79).
 *
 * ## Examples
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = new scaleway.ContainerNamespace("main", {
 *     description: "Main container namespace",
 * });
 * ```
 *
 * ## Import
 *
 * Namespaces can be imported using the `{region}/{id}`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:index/containerNamespace:ContainerNamespace main fr-par/11111111-1111-1111-1111-111111111111
 * ```
 */
export class ContainerNamespace extends pulumi.CustomResource {
    /**
     * Get an existing ContainerNamespace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContainerNamespaceState, opts?: pulumi.CustomResourceOptions): ContainerNamespace {
        return new ContainerNamespace(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/containerNamespace:ContainerNamespace';

    /**
     * Returns true if the given object is an instance of ContainerNamespace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ContainerNamespace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ContainerNamespace.__pulumiType;
    }

    /**
     * The description of the namespace.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Destroy registry on deletion
     *
     * @deprecated Registry namespace is automatically destroyed with namespace
     */
    public readonly destroyRegistry!: pulumi.Output<boolean | undefined>;
    /**
     * The environment variables of the namespace.
     */
    public readonly environmentVariables!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The unique name of the container namespace.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The organization ID the namespace is associated with.
     */
    public /*out*/ readonly organizationId!: pulumi.Output<string>;
    /**
     * `projectId`) The ID of the project the namespace is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * `region`). The region in which the namespace should be created.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The registry endpoint of the namespace.
     */
    public /*out*/ readonly registryEndpoint!: pulumi.Output<string>;
    /**
     * The registry namespace ID of the namespace.
     */
    public /*out*/ readonly registryNamespaceId!: pulumi.Output<string>;
    /**
     * The secret environment variables of the namespace.
     */
    public readonly secretEnvironmentVariables!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a ContainerNamespace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ContainerNamespaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ContainerNamespaceArgs | ContainerNamespaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ContainerNamespaceState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["destroyRegistry"] = state ? state.destroyRegistry : undefined;
            resourceInputs["environmentVariables"] = state ? state.environmentVariables : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["registryEndpoint"] = state ? state.registryEndpoint : undefined;
            resourceInputs["registryNamespaceId"] = state ? state.registryNamespaceId : undefined;
            resourceInputs["secretEnvironmentVariables"] = state ? state.secretEnvironmentVariables : undefined;
        } else {
            const args = argsOrState as ContainerNamespaceArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["destroyRegistry"] = args ? args.destroyRegistry : undefined;
            resourceInputs["environmentVariables"] = args ? args.environmentVariables : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["secretEnvironmentVariables"] = args ? args.secretEnvironmentVariables : undefined;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["registryEndpoint"] = undefined /*out*/;
            resourceInputs["registryNamespaceId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ContainerNamespace.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ContainerNamespace resources.
 */
export interface ContainerNamespaceState {
    /**
     * The description of the namespace.
     */
    description?: pulumi.Input<string>;
    /**
     * Destroy registry on deletion
     *
     * @deprecated Registry namespace is automatically destroyed with namespace
     */
    destroyRegistry?: pulumi.Input<boolean>;
    /**
     * The environment variables of the namespace.
     */
    environmentVariables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The unique name of the container namespace.
     */
    name?: pulumi.Input<string>;
    /**
     * The organization ID the namespace is associated with.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the namespace is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`). The region in which the namespace should be created.
     */
    region?: pulumi.Input<string>;
    /**
     * The registry endpoint of the namespace.
     */
    registryEndpoint?: pulumi.Input<string>;
    /**
     * The registry namespace ID of the namespace.
     */
    registryNamespaceId?: pulumi.Input<string>;
    /**
     * The secret environment variables of the namespace.
     */
    secretEnvironmentVariables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a ContainerNamespace resource.
 */
export interface ContainerNamespaceArgs {
    /**
     * The description of the namespace.
     */
    description?: pulumi.Input<string>;
    /**
     * Destroy registry on deletion
     *
     * @deprecated Registry namespace is automatically destroyed with namespace
     */
    destroyRegistry?: pulumi.Input<boolean>;
    /**
     * The environment variables of the namespace.
     */
    environmentVariables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The unique name of the container namespace.
     */
    name?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the namespace is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`). The region in which the namespace should be created.
     */
    region?: pulumi.Input<string>;
    /**
     * The secret environment variables of the namespace.
     */
    secretEnvironmentVariables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}