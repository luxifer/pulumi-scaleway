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
    public sealed class GetIotDeviceCertificateResult
    {
        public readonly string Crt;
        public readonly string Key;

        [OutputConstructor]
        private GetIotDeviceCertificateResult(
            string crt,

            string key)
        {
            Crt = crt;
            Key = key;
        }
    }
}
