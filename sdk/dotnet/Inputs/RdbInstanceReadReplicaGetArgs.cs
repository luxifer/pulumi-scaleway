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

    public sealed class RdbInstanceReadReplicaGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IP of the endpoint.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// The name of the Database Instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Port of the endpoint.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        public RdbInstanceReadReplicaGetArgs()
        {
        }
        public static new RdbInstanceReadReplicaGetArgs Empty => new RdbInstanceReadReplicaGetArgs();
    }
}
