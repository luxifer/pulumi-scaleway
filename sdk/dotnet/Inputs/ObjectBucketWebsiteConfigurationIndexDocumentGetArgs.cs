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

    public sealed class ObjectBucketWebsiteConfigurationIndexDocumentGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("suffix", required: true)]
        public Input<string> Suffix { get; set; } = null!;

        public ObjectBucketWebsiteConfigurationIndexDocumentGetArgs()
        {
        }
        public static new ObjectBucketWebsiteConfigurationIndexDocumentGetArgs Empty => new ObjectBucketWebsiteConfigurationIndexDocumentGetArgs();
    }
}
