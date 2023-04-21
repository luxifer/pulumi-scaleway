// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace luxifer.Scaleway
{
    /// <summary>
    /// Creates and manages Scaleway VPC Public Gateway DHCP.
    /// For more information, see [the documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#dhcp-c05544).
    /// 
    /// ## Example
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Scaleway = luxifer.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var main = new Scaleway.VpcPublicGatewayDhcp("main", new()
    ///     {
    ///         Subnet = "192.168.1.0/24",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Public gateway DHCP config can be imported using the `{zone}/{id}`, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:index/vpcPublicGatewayDhcp:VpcPublicGatewayDhcp main fr-par-1/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/vpcPublicGatewayDhcp:VpcPublicGatewayDhcp")]
    public partial class VpcPublicGatewayDhcp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The IP address of the public gateway DHCP config.
        /// </summary>
        [Output("address")]
        public Output<string> Address { get; private set; } = null!;

        /// <summary>
        /// The date and time of the creation of the public gateway DHCP config.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// TLD given to hostnames in the Private Network. Allowed characters are `a-z0-9-.`. Defaults to the slugified Private Network name if created along a GatewayNetwork, or else to `priv`.
        /// </summary>
        [Output("dnsLocalName")]
        public Output<string> DnsLocalName { get; private set; } = null!;

        /// <summary>
        /// Additional DNS search paths
        /// </summary>
        [Output("dnsSearches")]
        public Output<ImmutableArray<string>> DnsSearches { get; private set; } = null!;

        /// <summary>
        /// Override the DNS server list pushed to DHCP clients, instead of the gateway itself
        /// </summary>
        [Output("dnsServersOverrides")]
        public Output<ImmutableArray<string>> DnsServersOverrides { get; private set; } = null!;

        /// <summary>
        /// Whether to enable dynamic pooling of IPs. By turning the dynamic pool off, only pre-existing DHCP reservations will be handed out. Defaults to `true`.
        /// </summary>
        [Output("enableDynamic")]
        public Output<bool> EnableDynamic { get; private set; } = null!;

        /// <summary>
        /// The organization ID the public gateway DHCP config is associated with.
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// High IP (excluded) of the dynamic address pool. Defaults to the last address of the subnet.
        /// </summary>
        [Output("poolHigh")]
        public Output<string> PoolHigh { get; private set; } = null!;

        /// <summary>
        /// Low IP (included) of the dynamic address pool. Defaults to the second address of the subnet.
        /// </summary>
        [Output("poolLow")]
        public Output<string> PoolLow { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the public gateway DHCP config is associated with.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Whether the gateway should push a default route to DHCP clients or only hand out IPs. Defaults to `true`.
        /// </summary>
        [Output("pushDefaultRoute")]
        public Output<bool> PushDefaultRoute { get; private set; } = null!;

        /// <summary>
        /// Whether the gateway should push custom DNS servers to clients. This allows for instance hostname &gt; IP resolution. Defaults to `true`.
        /// </summary>
        [Output("pushDnsServer")]
        public Output<bool> PushDnsServer { get; private set; } = null!;

        /// <summary>
        /// After how long, in seconds, a DHCP client will query for a new lease if previous renews fail. Must be 30s lower than `valid_lifetime`. Defaults to 51m (3060s).
        /// </summary>
        [Output("rebindTimer")]
        public Output<int> RebindTimer { get; private set; } = null!;

        /// <summary>
        /// After how long, in seconds, a renewal will be attempted. Must be 30s lower than `rebind_timer`. Defaults to 50m (3000s).
        /// </summary>
        [Output("renewTimer")]
        public Output<int> RenewTimer { get; private set; } = null!;

        /// <summary>
        /// The subnet to associate with the public gateway DHCP config.
        /// </summary>
        [Output("subnet")]
        public Output<string> Subnet { get; private set; } = null!;

        /// <summary>
        /// The date and time of the last update of the public gateway DHCP config.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// For how long, in seconds, will DHCP entries will be valid. Defaults to 1h (3600s).
        /// </summary>
        [Output("validLifetime")]
        public Output<int> ValidLifetime { get; private set; } = null!;

        /// <summary>
        /// `zone`) The zone in which the public gateway DHCP config should be created.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a VpcPublicGatewayDhcp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcPublicGatewayDhcp(string name, VpcPublicGatewayDhcpArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/vpcPublicGatewayDhcp:VpcPublicGatewayDhcp", name, args ?? new VpcPublicGatewayDhcpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcPublicGatewayDhcp(string name, Input<string> id, VpcPublicGatewayDhcpState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/vpcPublicGatewayDhcp:VpcPublicGatewayDhcp", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/luxifer/pulumi-scaleway",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VpcPublicGatewayDhcp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcPublicGatewayDhcp Get(string name, Input<string> id, VpcPublicGatewayDhcpState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcPublicGatewayDhcp(name, id, state, options);
        }
    }

    public sealed class VpcPublicGatewayDhcpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IP address of the public gateway DHCP config.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// TLD given to hostnames in the Private Network. Allowed characters are `a-z0-9-.`. Defaults to the slugified Private Network name if created along a GatewayNetwork, or else to `priv`.
        /// </summary>
        [Input("dnsLocalName")]
        public Input<string>? DnsLocalName { get; set; }

        [Input("dnsSearches")]
        private InputList<string>? _dnsSearches;

        /// <summary>
        /// Additional DNS search paths
        /// </summary>
        public InputList<string> DnsSearches
        {
            get => _dnsSearches ?? (_dnsSearches = new InputList<string>());
            set => _dnsSearches = value;
        }

        [Input("dnsServersOverrides")]
        private InputList<string>? _dnsServersOverrides;

        /// <summary>
        /// Override the DNS server list pushed to DHCP clients, instead of the gateway itself
        /// </summary>
        public InputList<string> DnsServersOverrides
        {
            get => _dnsServersOverrides ?? (_dnsServersOverrides = new InputList<string>());
            set => _dnsServersOverrides = value;
        }

        /// <summary>
        /// Whether to enable dynamic pooling of IPs. By turning the dynamic pool off, only pre-existing DHCP reservations will be handed out. Defaults to `true`.
        /// </summary>
        [Input("enableDynamic")]
        public Input<bool>? EnableDynamic { get; set; }

        /// <summary>
        /// High IP (excluded) of the dynamic address pool. Defaults to the last address of the subnet.
        /// </summary>
        [Input("poolHigh")]
        public Input<string>? PoolHigh { get; set; }

        /// <summary>
        /// Low IP (included) of the dynamic address pool. Defaults to the second address of the subnet.
        /// </summary>
        [Input("poolLow")]
        public Input<string>? PoolLow { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the public gateway DHCP config is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Whether the gateway should push a default route to DHCP clients or only hand out IPs. Defaults to `true`.
        /// </summary>
        [Input("pushDefaultRoute")]
        public Input<bool>? PushDefaultRoute { get; set; }

        /// <summary>
        /// Whether the gateway should push custom DNS servers to clients. This allows for instance hostname &gt; IP resolution. Defaults to `true`.
        /// </summary>
        [Input("pushDnsServer")]
        public Input<bool>? PushDnsServer { get; set; }

        /// <summary>
        /// After how long, in seconds, a DHCP client will query for a new lease if previous renews fail. Must be 30s lower than `valid_lifetime`. Defaults to 51m (3060s).
        /// </summary>
        [Input("rebindTimer")]
        public Input<int>? RebindTimer { get; set; }

        /// <summary>
        /// After how long, in seconds, a renewal will be attempted. Must be 30s lower than `rebind_timer`. Defaults to 50m (3000s).
        /// </summary>
        [Input("renewTimer")]
        public Input<int>? RenewTimer { get; set; }

        /// <summary>
        /// The subnet to associate with the public gateway DHCP config.
        /// </summary>
        [Input("subnet", required: true)]
        public Input<string> Subnet { get; set; } = null!;

        /// <summary>
        /// For how long, in seconds, will DHCP entries will be valid. Defaults to 1h (3600s).
        /// </summary>
        [Input("validLifetime")]
        public Input<int>? ValidLifetime { get; set; }

        /// <summary>
        /// `zone`) The zone in which the public gateway DHCP config should be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public VpcPublicGatewayDhcpArgs()
        {
        }
        public static new VpcPublicGatewayDhcpArgs Empty => new VpcPublicGatewayDhcpArgs();
    }

    public sealed class VpcPublicGatewayDhcpState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IP address of the public gateway DHCP config.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// The date and time of the creation of the public gateway DHCP config.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// TLD given to hostnames in the Private Network. Allowed characters are `a-z0-9-.`. Defaults to the slugified Private Network name if created along a GatewayNetwork, or else to `priv`.
        /// </summary>
        [Input("dnsLocalName")]
        public Input<string>? DnsLocalName { get; set; }

        [Input("dnsSearches")]
        private InputList<string>? _dnsSearches;

        /// <summary>
        /// Additional DNS search paths
        /// </summary>
        public InputList<string> DnsSearches
        {
            get => _dnsSearches ?? (_dnsSearches = new InputList<string>());
            set => _dnsSearches = value;
        }

        [Input("dnsServersOverrides")]
        private InputList<string>? _dnsServersOverrides;

        /// <summary>
        /// Override the DNS server list pushed to DHCP clients, instead of the gateway itself
        /// </summary>
        public InputList<string> DnsServersOverrides
        {
            get => _dnsServersOverrides ?? (_dnsServersOverrides = new InputList<string>());
            set => _dnsServersOverrides = value;
        }

        /// <summary>
        /// Whether to enable dynamic pooling of IPs. By turning the dynamic pool off, only pre-existing DHCP reservations will be handed out. Defaults to `true`.
        /// </summary>
        [Input("enableDynamic")]
        public Input<bool>? EnableDynamic { get; set; }

        /// <summary>
        /// The organization ID the public gateway DHCP config is associated with.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// High IP (excluded) of the dynamic address pool. Defaults to the last address of the subnet.
        /// </summary>
        [Input("poolHigh")]
        public Input<string>? PoolHigh { get; set; }

        /// <summary>
        /// Low IP (included) of the dynamic address pool. Defaults to the second address of the subnet.
        /// </summary>
        [Input("poolLow")]
        public Input<string>? PoolLow { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the public gateway DHCP config is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Whether the gateway should push a default route to DHCP clients or only hand out IPs. Defaults to `true`.
        /// </summary>
        [Input("pushDefaultRoute")]
        public Input<bool>? PushDefaultRoute { get; set; }

        /// <summary>
        /// Whether the gateway should push custom DNS servers to clients. This allows for instance hostname &gt; IP resolution. Defaults to `true`.
        /// </summary>
        [Input("pushDnsServer")]
        public Input<bool>? PushDnsServer { get; set; }

        /// <summary>
        /// After how long, in seconds, a DHCP client will query for a new lease if previous renews fail. Must be 30s lower than `valid_lifetime`. Defaults to 51m (3060s).
        /// </summary>
        [Input("rebindTimer")]
        public Input<int>? RebindTimer { get; set; }

        /// <summary>
        /// After how long, in seconds, a renewal will be attempted. Must be 30s lower than `rebind_timer`. Defaults to 50m (3000s).
        /// </summary>
        [Input("renewTimer")]
        public Input<int>? RenewTimer { get; set; }

        /// <summary>
        /// The subnet to associate with the public gateway DHCP config.
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        /// <summary>
        /// The date and time of the last update of the public gateway DHCP config.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// For how long, in seconds, will DHCP entries will be valid. Defaults to 1h (3600s).
        /// </summary>
        [Input("validLifetime")]
        public Input<int>? ValidLifetime { get; set; }

        /// <summary>
        /// `zone`) The zone in which the public gateway DHCP config should be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public VpcPublicGatewayDhcpState()
        {
        }
        public static new VpcPublicGatewayDhcpState Empty => new VpcPublicGatewayDhcpState();
    }
}
