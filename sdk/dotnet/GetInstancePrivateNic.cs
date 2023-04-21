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
    public static class GetInstancePrivateNic
    {
        /// <summary>
        /// Gets information about an instance private NIC.
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
        ///     var byNicId = Scaleway.GetInstancePrivateNic.Invoke(new()
        ///     {
        ///         PrivateNicId = "11111111-1111-1111-1111-111111111111",
        ///         ServerId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        ///     var byPnId = Scaleway.GetInstancePrivateNic.Invoke(new()
        ///     {
        ///         PrivateNetworkId = "11111111-1111-1111-1111-111111111111",
        ///         ServerId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        ///     var byTags = Scaleway.GetInstancePrivateNic.Invoke(new()
        ///     {
        ///         ServerId = "11111111-1111-1111-1111-111111111111",
        ///         Tags = new[]
        ///         {
        ///             "mytag",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetInstancePrivateNicResult> InvokeAsync(GetInstancePrivateNicArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstancePrivateNicResult>("scaleway:index/getInstancePrivateNic:getInstancePrivateNic", args ?? new GetInstancePrivateNicArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about an instance private NIC.
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
        ///     var byNicId = Scaleway.GetInstancePrivateNic.Invoke(new()
        ///     {
        ///         PrivateNicId = "11111111-1111-1111-1111-111111111111",
        ///         ServerId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        ///     var byPnId = Scaleway.GetInstancePrivateNic.Invoke(new()
        ///     {
        ///         PrivateNetworkId = "11111111-1111-1111-1111-111111111111",
        ///         ServerId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        ///     var byTags = Scaleway.GetInstancePrivateNic.Invoke(new()
        ///     {
        ///         ServerId = "11111111-1111-1111-1111-111111111111",
        ///         Tags = new[]
        ///         {
        ///             "mytag",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetInstancePrivateNicResult> Invoke(GetInstancePrivateNicInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInstancePrivateNicResult>("scaleway:index/getInstancePrivateNic:getInstancePrivateNic", args ?? new GetInstancePrivateNicInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstancePrivateNicArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the private network
        /// Only one of `private_nic_id` and `private_network_id` should be specified.
        /// </summary>
        [Input("privateNetworkId")]
        public string? PrivateNetworkId { get; set; }

        /// <summary>
        /// The ID of the instance server private nic
        /// Only one of `private_nic_id` and `private_network_id` should be specified.
        /// </summary>
        [Input("privateNicId")]
        public string? PrivateNicId { get; set; }

        /// <summary>
        /// The server's id
        /// </summary>
        [Input("serverId", required: true)]
        public string ServerId { get; set; } = null!;

        [Input("tags")]
        private List<string>? _tags;

        /// <summary>
        /// The tags associated with the private NIC.
        /// As datasource only returns one private NIC, the search with given tags must return only one result
        /// </summary>
        public List<string> Tags
        {
            get => _tags ?? (_tags = new List<string>());
            set => _tags = value;
        }

        /// <summary>
        /// `zone`) The zone in which the private nic exists.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetInstancePrivateNicArgs()
        {
        }
        public static new GetInstancePrivateNicArgs Empty => new GetInstancePrivateNicArgs();
    }

    public sealed class GetInstancePrivateNicInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the private network
        /// Only one of `private_nic_id` and `private_network_id` should be specified.
        /// </summary>
        [Input("privateNetworkId")]
        public Input<string>? PrivateNetworkId { get; set; }

        /// <summary>
        /// The ID of the instance server private nic
        /// Only one of `private_nic_id` and `private_network_id` should be specified.
        /// </summary>
        [Input("privateNicId")]
        public Input<string>? PrivateNicId { get; set; }

        /// <summary>
        /// The server's id
        /// </summary>
        [Input("serverId", required: true)]
        public Input<string> ServerId { get; set; } = null!;

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the private NIC.
        /// As datasource only returns one private NIC, the search with given tags must return only one result
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// `zone`) The zone in which the private nic exists.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetInstancePrivateNicInvokeArgs()
        {
        }
        public static new GetInstancePrivateNicInvokeArgs Empty => new GetInstancePrivateNicInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstancePrivateNicResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string MacAddress;
        public readonly string? PrivateNetworkId;
        public readonly string? PrivateNicId;
        public readonly string ServerId;
        public readonly ImmutableArray<string> Tags;
        public readonly string? Zone;

        [OutputConstructor]
        private GetInstancePrivateNicResult(
            string id,

            string macAddress,

            string? privateNetworkId,

            string? privateNicId,

            string serverId,

            ImmutableArray<string> tags,

            string? zone)
        {
            Id = id;
            MacAddress = macAddress;
            PrivateNetworkId = privateNetworkId;
            PrivateNicId = privateNicId;
            ServerId = serverId;
            Tags = tags;
            Zone = zone;
        }
    }
}
