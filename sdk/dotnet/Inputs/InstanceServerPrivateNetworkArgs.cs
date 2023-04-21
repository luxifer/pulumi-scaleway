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

    public sealed class InstanceServerPrivateNetworkArgs : global::Pulumi.ResourceArgs
    {
        [Input("macAddress")]
        public Input<string>? MacAddress { get; set; }

        [Input("pnId", required: true)]
        public Input<string> PnId { get; set; } = null!;

        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// `zone`) The zone in which the server should be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public InstanceServerPrivateNetworkArgs()
        {
        }
        public static new InstanceServerPrivateNetworkArgs Empty => new InstanceServerPrivateNetworkArgs();
    }
}
