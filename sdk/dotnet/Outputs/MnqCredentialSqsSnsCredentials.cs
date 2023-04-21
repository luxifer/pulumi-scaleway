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
    public sealed class MnqCredentialSqsSnsCredentials
    {
        /// <summary>
        /// The ID of the key.
        /// </summary>
        public readonly string? AccessKey;
        /// <summary>
        /// List of permissions associated to this Credential. Only one of permissions may be set.
        /// </summary>
        public readonly Outputs.MnqCredentialSqsSnsCredentialsPermissions? Permissions;
        /// <summary>
        /// The Secret value of the key.
        /// </summary>
        public readonly string? SecretKey;

        [OutputConstructor]
        private MnqCredentialSqsSnsCredentials(
            string? accessKey,

            Outputs.MnqCredentialSqsSnsCredentialsPermissions? permissions,

            string? secretKey)
        {
            AccessKey = accessKey;
            Permissions = permissions;
            SecretKey = secretKey;
        }
    }
}