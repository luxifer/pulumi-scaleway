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
    public static class GetInstanceServer
    {
        /// <summary>
        /// Gets information about an instance server.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myKey = Scaleway.GetInstanceServer.Invoke(new()
        ///     {
        ///         ServerId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetInstanceServerResult> InvokeAsync(GetInstanceServerArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceServerResult>("scaleway:index/getInstanceServer:getInstanceServer", args ?? new GetInstanceServerArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about an instance server.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myKey = Scaleway.GetInstanceServer.Invoke(new()
        ///     {
        ///         ServerId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetInstanceServerResult> Invoke(GetInstanceServerInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInstanceServerResult>("scaleway:index/getInstanceServer:getInstanceServer", args ?? new GetInstanceServerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceServerArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The server name. Only one of `name` and `server_id` should be specified.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The server id. Only one of `name` and `server_id` should be specified.
        /// </summary>
        [Input("serverId")]
        public string? ServerId { get; set; }

        /// <summary>
        /// `zone`) The zone in which the server exists.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetInstanceServerArgs()
        {
        }
        public static new GetInstanceServerArgs Empty => new GetInstanceServerArgs();
    }

    public sealed class GetInstanceServerInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The server name. Only one of `name` and `server_id` should be specified.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The server id. Only one of `name` and `server_id` should be specified.
        /// </summary>
        [Input("serverId")]
        public Input<string>? ServerId { get; set; }

        /// <summary>
        /// `zone`) The zone in which the server exists.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetInstanceServerInvokeArgs()
        {
        }
        public static new GetInstanceServerInvokeArgs Empty => new GetInstanceServerInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstanceServerResult
    {
        /// <summary>
        /// The [additional volumes](https://developers.scaleway.com/en/products/instance/api/#volumes-7e8a39)
        /// attached to the server.
        /// </summary>
        public readonly ImmutableArray<string> AdditionalVolumeIds;
        public readonly string BootType;
        public readonly string BootscriptId;
        /// <summary>
        /// The cloud init script associated with this server.
        /// </summary>
        public readonly string CloudInit;
        /// <summary>
        /// True is dynamic IP in enable on the server.
        /// </summary>
        public readonly bool EnableDynamicIp;
        /// <summary>
        /// Determines if IPv6 is enabled for the server.
        /// </summary>
        public readonly bool EnableIpv6;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The UUID and the label of the base image used by the server.
        /// </summary>
        public readonly string Image;
        public readonly string IpId;
        /// <summary>
        /// The default ipv6 address routed to the server. ( Only set when enable_ipv6 is set to true )
        /// </summary>
        public readonly string Ipv6Address;
        /// <summary>
        /// The ipv6 gateway address. ( Only set when enable_ipv6 is set to true )
        /// </summary>
        public readonly string Ipv6Gateway;
        /// <summary>
        /// The prefix length of the ipv6 subnet routed to the server. ( Only set when enable_ipv6 is set to true )
        /// </summary>
        public readonly int Ipv6PrefixLength;
        public readonly string? Name;
        /// <summary>
        /// The ID of the organization the server is associated with.
        /// </summary>
        public readonly string OrganizationId;
        /// <summary>
        /// The [placement group](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) the server is attached to.
        /// </summary>
        public readonly string PlacementGroupId;
        /// <summary>
        /// True when the placement group policy is respected.
        /// </summary>
        public readonly bool PlacementGroupPolicyRespected;
        /// <summary>
        /// The Scaleway internal IP address of the server.
        /// </summary>
        public readonly string PrivateIp;
        public readonly ImmutableArray<Outputs.GetInstanceServerPrivateNetworkResult> PrivateNetworks;
        /// <summary>
        /// The ID of the project the server is associated with.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// The public IPv4 address of the server.
        /// </summary>
        public readonly string PublicIp;
        /// <summary>
        /// Root [volume](https://developers.scaleway.com/en/products/instance/api/#volumes-7e8a39) attached to the server on creation.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceServerRootVolumeResult> RootVolumes;
        /// <summary>
        /// The [security group](https://developers.scaleway.com/en/products/instance/api/#security-groups-8d7f89) the server is attached to.
        /// </summary>
        public readonly string SecurityGroupId;
        public readonly string? ServerId;
        /// <summary>
        /// The state of the server. Possible values are: `started`, `stopped` or `standby`.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The tags associated with the server.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// The commercial type of the server.
        /// You find all the available types on the [pricing page](https://www.scaleway.com/en/pricing/).
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The user data associated with the server.
        /// </summary>
        public readonly ImmutableDictionary<string, string> UserData;
        public readonly string? Zone;

        [OutputConstructor]
        private GetInstanceServerResult(
            ImmutableArray<string> additionalVolumeIds,

            string bootType,

            string bootscriptId,

            string cloudInit,

            bool enableDynamicIp,

            bool enableIpv6,

            string id,

            string image,

            string ipId,

            string ipv6Address,

            string ipv6Gateway,

            int ipv6PrefixLength,

            string? name,

            string organizationId,

            string placementGroupId,

            bool placementGroupPolicyRespected,

            string privateIp,

            ImmutableArray<Outputs.GetInstanceServerPrivateNetworkResult> privateNetworks,

            string projectId,

            string publicIp,

            ImmutableArray<Outputs.GetInstanceServerRootVolumeResult> rootVolumes,

            string securityGroupId,

            string? serverId,

            string state,

            ImmutableArray<string> tags,

            string type,

            ImmutableDictionary<string, string> userData,

            string? zone)
        {
            AdditionalVolumeIds = additionalVolumeIds;
            BootType = bootType;
            BootscriptId = bootscriptId;
            CloudInit = cloudInit;
            EnableDynamicIp = enableDynamicIp;
            EnableIpv6 = enableIpv6;
            Id = id;
            Image = image;
            IpId = ipId;
            Ipv6Address = ipv6Address;
            Ipv6Gateway = ipv6Gateway;
            Ipv6PrefixLength = ipv6PrefixLength;
            Name = name;
            OrganizationId = organizationId;
            PlacementGroupId = placementGroupId;
            PlacementGroupPolicyRespected = placementGroupPolicyRespected;
            PrivateIp = privateIp;
            PrivateNetworks = privateNetworks;
            ProjectId = projectId;
            PublicIp = publicIp;
            RootVolumes = rootVolumes;
            SecurityGroupId = securityGroupId;
            ServerId = serverId;
            State = state;
            Tags = tags;
            Type = type;
            UserData = userData;
            Zone = zone;
        }
    }
}
