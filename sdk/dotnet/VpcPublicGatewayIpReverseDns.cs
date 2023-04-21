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
    /// Manages Scaleway VPC Public Gateways IPs reverse DNS.
    /// For more information, see [the documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#ips-268151).
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
    ///     var mainVpcPublicGatewayIp = new Scaleway.VpcPublicGatewayIp("mainVpcPublicGatewayIp");
    /// 
    ///     var tfA = new Scaleway.DomainRecord("tfA", new()
    ///     {
    ///         DnsZone = "example.com",
    ///         Type = "A",
    ///         Data = mainVpcPublicGatewayIp.Address,
    ///         Ttl = 3600,
    ///         Priority = 1,
    ///     });
    /// 
    ///     var mainVpcPublicGatewayIpReverseDns = new Scaleway.VpcPublicGatewayIpReverseDns("mainVpcPublicGatewayIpReverseDns", new()
    ///     {
    ///         GatewayIpId = mainVpcPublicGatewayIp.Id,
    ///         Reverse = "tf.example.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Public gateway IPs reverse DNS can be imported using the `{zone}/{id}`, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:index/vpcPublicGatewayIpReverseDns:VpcPublicGatewayIpReverseDns main fr-par-1/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/vpcPublicGatewayIpReverseDns:VpcPublicGatewayIpReverseDns")]
    public partial class VpcPublicGatewayIpReverseDns : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The public gateway IP ID
        /// </summary>
        [Output("gatewayIpId")]
        public Output<string> GatewayIpId { get; private set; } = null!;

        /// <summary>
        /// The reverse domain name for this IP address
        /// </summary>
        [Output("reverse")]
        public Output<string> Reverse { get; private set; } = null!;

        /// <summary>
        /// `zone`) The zone in which the IP should be reserved.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a VpcPublicGatewayIpReverseDns resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcPublicGatewayIpReverseDns(string name, VpcPublicGatewayIpReverseDnsArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/vpcPublicGatewayIpReverseDns:VpcPublicGatewayIpReverseDns", name, args ?? new VpcPublicGatewayIpReverseDnsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcPublicGatewayIpReverseDns(string name, Input<string> id, VpcPublicGatewayIpReverseDnsState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/vpcPublicGatewayIpReverseDns:VpcPublicGatewayIpReverseDns", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpcPublicGatewayIpReverseDns resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcPublicGatewayIpReverseDns Get(string name, Input<string> id, VpcPublicGatewayIpReverseDnsState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcPublicGatewayIpReverseDns(name, id, state, options);
        }
    }

    public sealed class VpcPublicGatewayIpReverseDnsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The public gateway IP ID
        /// </summary>
        [Input("gatewayIpId", required: true)]
        public Input<string> GatewayIpId { get; set; } = null!;

        /// <summary>
        /// The reverse domain name for this IP address
        /// </summary>
        [Input("reverse", required: true)]
        public Input<string> Reverse { get; set; } = null!;

        /// <summary>
        /// `zone`) The zone in which the IP should be reserved.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public VpcPublicGatewayIpReverseDnsArgs()
        {
        }
        public static new VpcPublicGatewayIpReverseDnsArgs Empty => new VpcPublicGatewayIpReverseDnsArgs();
    }

    public sealed class VpcPublicGatewayIpReverseDnsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The public gateway IP ID
        /// </summary>
        [Input("gatewayIpId")]
        public Input<string>? GatewayIpId { get; set; }

        /// <summary>
        /// The reverse domain name for this IP address
        /// </summary>
        [Input("reverse")]
        public Input<string>? Reverse { get; set; }

        /// <summary>
        /// `zone`) The zone in which the IP should be reserved.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public VpcPublicGatewayIpReverseDnsState()
        {
        }
        public static new VpcPublicGatewayIpReverseDnsState Empty => new VpcPublicGatewayIpReverseDnsState();
    }
}
