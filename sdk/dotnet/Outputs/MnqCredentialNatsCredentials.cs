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
    public sealed class MnqCredentialNatsCredentials
    {
        /// <summary>
        /// Raw content of the NATS credentials file.
        /// </summary>
        public readonly string? Content;

        [OutputConstructor]
        private MnqCredentialNatsCredentials(string? content)
        {
            Content = content;
        }
    }
}
