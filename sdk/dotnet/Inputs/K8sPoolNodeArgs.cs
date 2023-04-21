// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace luxifer.Scaleway.Inputs
{

    public sealed class K8sPoolNodeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name for the pool.
        /// &gt; **Important:** Updates to this field will recreate a new resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The public IPv4.
        /// </summary>
        [Input("publicIp")]
        public Input<string>? PublicIp { get; set; }

        /// <summary>
        /// The public IPv6.
        /// </summary>
        [Input("publicIpV6")]
        public Input<string>? PublicIpV6 { get; set; }

        /// <summary>
        /// The status of the node.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public K8sPoolNodeArgs()
        {
        }
        public static new K8sPoolNodeArgs Empty => new K8sPoolNodeArgs();
    }
}
