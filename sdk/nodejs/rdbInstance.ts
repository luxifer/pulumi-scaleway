// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Database Instances.
 * For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api).
 *
 * ## Examples
 *
 * ### Example Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = new scaleway.RdbInstance("main", {
 *     disableBackup: true,
 *     engine: "PostgreSQL-11",
 *     isHaCluster: true,
 *     nodeType: "DB-DEV-S",
 *     password: "thiZ_is_v&ry_s3cret",
 *     userName: "my_initial_user",
 * });
 * ```
 *
 * ### Example with Settings
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = new scaleway.RdbInstance("main", {
 *     disableBackup: true,
 *     engine: "MySQL-8",
 *     initSettings: {
 *         lower_case_table_names: 1,
 *     },
 *     nodeType: "db-dev-s",
 *     password: "thiZ_is_v&ry_s3cret",
 *     settings: {
 *         max_connections: "350",
 *     },
 *     userName: "my_initial_user",
 * });
 * ```
 *
 * ### Example with backup schedule
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = new scaleway.RdbInstance("main", {
 *     backupScheduleFrequency: 24, // every day
 *     backupScheduleRetention: 7, // keep it one week
 *     disableBackup: false,
 *     engine: "PostgreSQL-11",
 *     isHaCluster: true,
 *     nodeType: "DB-DEV-S",
 *     password: "thiZ_is_v&ry_s3cret",
 *     userName: "my_initial_user",
 * });
 * ```
 *
 * ### Example with private network and dhcp configuration
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi_scaleway from "@luxifer/pulumi-scaleway";
 *
 * const pn02 = new scaleway.VpcPrivateNetwork("pn02", {});
 * const mainVpcPublicGatewayDhcp = new scaleway.VpcPublicGatewayDhcp("mainVpcPublicGatewayDhcp", {subnet: "192.168.1.0/24"});
 * const mainVpcPublicGatewayIp = new scaleway.VpcPublicGatewayIp("mainVpcPublicGatewayIp", {});
 * const mainVpcPublicGateway = new scaleway.VpcPublicGateway("mainVpcPublicGateway", {
 *     type: "VPC-GW-S",
 *     ipId: mainVpcPublicGatewayIp.id,
 * });
 * const mainVpcGatewayNetwork = new scaleway.VpcGatewayNetwork("mainVpcGatewayNetwork", {
 *     gatewayId: mainVpcPublicGateway.id,
 *     privateNetworkId: pn02.id,
 *     dhcpId: mainVpcPublicGatewayDhcp.id,
 *     cleanupDhcp: true,
 *     enableMasquerade: true,
 * }, {
 *     dependsOn: [
 *         mainVpcPublicGatewayIp,
 *         pn02,
 *     ],
 * });
 * const mainRdbInstance = new scaleway.RdbInstance("mainRdbInstance", {
 *     nodeType: "db-dev-s",
 *     engine: "PostgreSQL-11",
 *     isHaCluster: false,
 *     disableBackup: true,
 *     userName: "my_initial_user",
 *     password: "thiZ_is_v&ry_s3cret",
 *     region: "fr-par",
 *     tags: [
 *         "terraform-test",
 *         "scaleway_rdb_instance",
 *         "volume",
 *         "rdb_pn",
 *     ],
 *     volumeType: "bssd",
 *     volumeSizeInGb: 10,
 *     privateNetwork: {
 *         ipNet: "192.168.1.254/24",
 *         pnId: pn02.id,
 *     },
 * });
 * const mainVpcPublicGatewayPatRule = new scaleway.VpcPublicGatewayPatRule("mainVpcPublicGatewayPatRule", {
 *     gatewayId: mainVpcPublicGateway.id,
 *     privateIp: mainVpcPublicGatewayDhcp.address,
 *     privatePort: mainRdbInstance.privateNetwork.apply(privateNetwork => privateNetwork?.port),
 *     publicPort: 42,
 *     protocol: "both",
 * }, {
 *     dependsOn: [
 *         mainVpcGatewayNetwork,
 *         pn02,
 *     ],
 * });
 * ```
 *
 * ## Settings
 *
 * Please consult the [GoDoc](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@v1.0.0-beta.9/api/rdb/v1#EngineVersion) to list all available `settings` and `initSettings` on your `nodeType` of your convenient.
 *
 * ## Private Network
 *
 * > **Important:** Updates to `privateNetwork` will recreate the attachment Instance.
 *
 * - `ipNet` - (Required) The IP network where to con.
 * - `pnId` - (Required) The ID of the private network. If not provided it will be randomly generated.
 *
 * ## Limitations
 *
 * The Managed Database product is only compliant with the private network in the default availability zone (AZ).
 * i.e. `fr-par-1`, `nl-ams-1`, `pl-waw-1`. To learn more, read our section [How to connect a PostgreSQL and MySQL Database Instance to a Private Network](https://www.scaleway.com/en/docs/managed-databases/postgresql-and-mysql/how-to/connect-database-private-network/)
 *
 * ## Import
 *
 * Database Instance can be imported using the `{region}/{id}`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:index/rdbInstance:RdbInstance rdb01 fr-par/11111111-1111-1111-1111-111111111111
 * ```
 */
export class RdbInstance extends pulumi.CustomResource {
    /**
     * Get an existing RdbInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RdbInstanceState, opts?: pulumi.CustomResourceOptions): RdbInstance {
        return new RdbInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/rdbInstance:RdbInstance';

    /**
     * Returns true if the given object is an instance of RdbInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RdbInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RdbInstance.__pulumiType;
    }

    /**
     * Boolean to store logical backups in the same region as the database instance.
     */
    public readonly backupSameRegion!: pulumi.Output<boolean>;
    /**
     * Backup schedule frequency in hours.
     */
    public readonly backupScheduleFrequency!: pulumi.Output<number>;
    /**
     * Backup schedule retention in days.
     */
    public readonly backupScheduleRetention!: pulumi.Output<number>;
    /**
     * Certificate of the database instance.
     */
    public /*out*/ readonly certificate!: pulumi.Output<string>;
    /**
     * Disable automated backup for the database instance.
     */
    public readonly disableBackup!: pulumi.Output<boolean | undefined>;
    /**
     * (Deprecated) The IP of the Database Instance.
     *
     * @deprecated Please use the private_network or the load_balancer attribute
     */
    public /*out*/ readonly endpointIp!: pulumi.Output<string>;
    /**
     * (Deprecated) The port of the Database Instance.
     */
    public /*out*/ readonly endpointPort!: pulumi.Output<number>;
    /**
     * Database Instance's engine version (e.g. `PostgreSQL-11`).
     */
    public readonly engine!: pulumi.Output<string>;
    /**
     * Map of engine settings to be set at database initialisation.
     */
    public readonly initSettings!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Enable or disable high availability for the database instance.
     */
    public readonly isHaCluster!: pulumi.Output<boolean | undefined>;
    /**
     * List of load balancer endpoints of the database instance.
     */
    public /*out*/ readonly loadBalancers!: pulumi.Output<outputs.RdbInstanceLoadBalancer[]>;
    /**
     * The name of the Database Instance.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The type of database instance you want to create (e.g. `db-dev-s`).
     */
    public readonly nodeType!: pulumi.Output<string>;
    /**
     * The organization ID the Database Instance is associated with.
     */
    public /*out*/ readonly organizationId!: pulumi.Output<string>;
    /**
     * Password for the first user of the database instance.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * List of private networks endpoints of the database instance.
     */
    public readonly privateNetwork!: pulumi.Output<outputs.RdbInstancePrivateNetwork | undefined>;
    /**
     * `projectId`) The ID of the project the Database Instance is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * List of read replicas of the database instance.
     */
    public /*out*/ readonly readReplicas!: pulumi.Output<outputs.RdbInstanceReadReplica[]>;
    /**
     * `region`) The region in which the Database Instance should be created.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Map of engine settings to be set. Using this option will override default config.
     */
    public readonly settings!: pulumi.Output<{[key: string]: string}>;
    /**
     * The tags associated with the Database Instance.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * Identifier for the first user of the database instance.
     */
    public readonly userName!: pulumi.Output<string | undefined>;
    /**
     * Volume size (in GB) when `volumeType` is set to `bssd`. Must be a multiple of 5000000000.
     */
    public readonly volumeSizeInGb!: pulumi.Output<number>;
    /**
     * Type of volume where data are stored (`bssd` or `lssd`).
     */
    public readonly volumeType!: pulumi.Output<string | undefined>;

    /**
     * Create a RdbInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RdbInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RdbInstanceArgs | RdbInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RdbInstanceState | undefined;
            resourceInputs["backupSameRegion"] = state ? state.backupSameRegion : undefined;
            resourceInputs["backupScheduleFrequency"] = state ? state.backupScheduleFrequency : undefined;
            resourceInputs["backupScheduleRetention"] = state ? state.backupScheduleRetention : undefined;
            resourceInputs["certificate"] = state ? state.certificate : undefined;
            resourceInputs["disableBackup"] = state ? state.disableBackup : undefined;
            resourceInputs["endpointIp"] = state ? state.endpointIp : undefined;
            resourceInputs["endpointPort"] = state ? state.endpointPort : undefined;
            resourceInputs["engine"] = state ? state.engine : undefined;
            resourceInputs["initSettings"] = state ? state.initSettings : undefined;
            resourceInputs["isHaCluster"] = state ? state.isHaCluster : undefined;
            resourceInputs["loadBalancers"] = state ? state.loadBalancers : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nodeType"] = state ? state.nodeType : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["privateNetwork"] = state ? state.privateNetwork : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["readReplicas"] = state ? state.readReplicas : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["settings"] = state ? state.settings : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
            resourceInputs["volumeSizeInGb"] = state ? state.volumeSizeInGb : undefined;
            resourceInputs["volumeType"] = state ? state.volumeType : undefined;
        } else {
            const args = argsOrState as RdbInstanceArgs | undefined;
            if ((!args || args.engine === undefined) && !opts.urn) {
                throw new Error("Missing required property 'engine'");
            }
            if ((!args || args.nodeType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeType'");
            }
            resourceInputs["backupSameRegion"] = args ? args.backupSameRegion : undefined;
            resourceInputs["backupScheduleFrequency"] = args ? args.backupScheduleFrequency : undefined;
            resourceInputs["backupScheduleRetention"] = args ? args.backupScheduleRetention : undefined;
            resourceInputs["disableBackup"] = args ? args.disableBackup : undefined;
            resourceInputs["engine"] = args ? args.engine : undefined;
            resourceInputs["initSettings"] = args ? args.initSettings : undefined;
            resourceInputs["isHaCluster"] = args ? args.isHaCluster : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nodeType"] = args ? args.nodeType : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["privateNetwork"] = args ? args.privateNetwork : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["settings"] = args ? args.settings : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
            resourceInputs["volumeSizeInGb"] = args ? args.volumeSizeInGb : undefined;
            resourceInputs["volumeType"] = args ? args.volumeType : undefined;
            resourceInputs["certificate"] = undefined /*out*/;
            resourceInputs["endpointIp"] = undefined /*out*/;
            resourceInputs["endpointPort"] = undefined /*out*/;
            resourceInputs["loadBalancers"] = undefined /*out*/;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["readReplicas"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RdbInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RdbInstance resources.
 */
export interface RdbInstanceState {
    /**
     * Boolean to store logical backups in the same region as the database instance.
     */
    backupSameRegion?: pulumi.Input<boolean>;
    /**
     * Backup schedule frequency in hours.
     */
    backupScheduleFrequency?: pulumi.Input<number>;
    /**
     * Backup schedule retention in days.
     */
    backupScheduleRetention?: pulumi.Input<number>;
    /**
     * Certificate of the database instance.
     */
    certificate?: pulumi.Input<string>;
    /**
     * Disable automated backup for the database instance.
     */
    disableBackup?: pulumi.Input<boolean>;
    /**
     * (Deprecated) The IP of the Database Instance.
     *
     * @deprecated Please use the private_network or the load_balancer attribute
     */
    endpointIp?: pulumi.Input<string>;
    /**
     * (Deprecated) The port of the Database Instance.
     */
    endpointPort?: pulumi.Input<number>;
    /**
     * Database Instance's engine version (e.g. `PostgreSQL-11`).
     */
    engine?: pulumi.Input<string>;
    /**
     * Map of engine settings to be set at database initialisation.
     */
    initSettings?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Enable or disable high availability for the database instance.
     */
    isHaCluster?: pulumi.Input<boolean>;
    /**
     * List of load balancer endpoints of the database instance.
     */
    loadBalancers?: pulumi.Input<pulumi.Input<inputs.RdbInstanceLoadBalancer>[]>;
    /**
     * The name of the Database Instance.
     */
    name?: pulumi.Input<string>;
    /**
     * The type of database instance you want to create (e.g. `db-dev-s`).
     */
    nodeType?: pulumi.Input<string>;
    /**
     * The organization ID the Database Instance is associated with.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * Password for the first user of the database instance.
     */
    password?: pulumi.Input<string>;
    /**
     * List of private networks endpoints of the database instance.
     */
    privateNetwork?: pulumi.Input<inputs.RdbInstancePrivateNetwork>;
    /**
     * `projectId`) The ID of the project the Database Instance is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * List of read replicas of the database instance.
     */
    readReplicas?: pulumi.Input<pulumi.Input<inputs.RdbInstanceReadReplica>[]>;
    /**
     * `region`) The region in which the Database Instance should be created.
     */
    region?: pulumi.Input<string>;
    /**
     * Map of engine settings to be set. Using this option will override default config.
     */
    settings?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The tags associated with the Database Instance.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Identifier for the first user of the database instance.
     */
    userName?: pulumi.Input<string>;
    /**
     * Volume size (in GB) when `volumeType` is set to `bssd`. Must be a multiple of 5000000000.
     */
    volumeSizeInGb?: pulumi.Input<number>;
    /**
     * Type of volume where data are stored (`bssd` or `lssd`).
     */
    volumeType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RdbInstance resource.
 */
export interface RdbInstanceArgs {
    /**
     * Boolean to store logical backups in the same region as the database instance.
     */
    backupSameRegion?: pulumi.Input<boolean>;
    /**
     * Backup schedule frequency in hours.
     */
    backupScheduleFrequency?: pulumi.Input<number>;
    /**
     * Backup schedule retention in days.
     */
    backupScheduleRetention?: pulumi.Input<number>;
    /**
     * Disable automated backup for the database instance.
     */
    disableBackup?: pulumi.Input<boolean>;
    /**
     * Database Instance's engine version (e.g. `PostgreSQL-11`).
     */
    engine: pulumi.Input<string>;
    /**
     * Map of engine settings to be set at database initialisation.
     */
    initSettings?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Enable or disable high availability for the database instance.
     */
    isHaCluster?: pulumi.Input<boolean>;
    /**
     * The name of the Database Instance.
     */
    name?: pulumi.Input<string>;
    /**
     * The type of database instance you want to create (e.g. `db-dev-s`).
     */
    nodeType: pulumi.Input<string>;
    /**
     * Password for the first user of the database instance.
     */
    password?: pulumi.Input<string>;
    /**
     * List of private networks endpoints of the database instance.
     */
    privateNetwork?: pulumi.Input<inputs.RdbInstancePrivateNetwork>;
    /**
     * `projectId`) The ID of the project the Database Instance is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`) The region in which the Database Instance should be created.
     */
    region?: pulumi.Input<string>;
    /**
     * Map of engine settings to be set. Using this option will override default config.
     */
    settings?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The tags associated with the Database Instance.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Identifier for the first user of the database instance.
     */
    userName?: pulumi.Input<string>;
    /**
     * Volume size (in GB) when `volumeType` is set to `bssd`. Must be a multiple of 5000000000.
     */
    volumeSizeInGb?: pulumi.Input<number>;
    /**
     * Type of volume where data are stored (`bssd` or `lssd`).
     */
    volumeType?: pulumi.Input<string>;
}
