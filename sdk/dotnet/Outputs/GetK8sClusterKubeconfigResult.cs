// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace luxifer.Scaleway.Outputs
{

    [OutputType]
    public sealed class GetK8sClusterKubeconfigResult
    {
        /// <summary>
        /// The CA certificate of the Kubernetes API server.
        /// </summary>
        public readonly string ClusterCaCertificate;
        /// <summary>
        /// The raw kubeconfig file.
        /// </summary>
        public readonly string ConfigFile;
        /// <summary>
        /// The URL of the Kubernetes API server.
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// The token to connect to the Kubernetes API server.
        /// </summary>
        public readonly string Token;

        [OutputConstructor]
        private GetK8sClusterKubeconfigResult(
            string clusterCaCertificate,

            string configFile,

            string host,

            string token)
        {
            ClusterCaCertificate = clusterCaCertificate;
            ConfigFile = configFile;
            Host = host;
            Token = token;
        }
    }
}
