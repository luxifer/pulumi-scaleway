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
    public sealed class LbFrontendAcl
    {
        /// <summary>
        /// Action to undertake when an ACL filter matches.
        /// </summary>
        public readonly Outputs.LbFrontendAclAction Action;
        /// <summary>
        /// The ACL match rule. At least `ip_subnet` or `http_filter` and `http_filter_value` are required.
        /// </summary>
        public readonly Outputs.LbFrontendAclMatch Match;
        /// <summary>
        /// The ACL name. If not provided it will be randomly generated.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private LbFrontendAcl(
            Outputs.LbFrontendAclAction action,

            Outputs.LbFrontendAclMatch match,

            string? name)
        {
            Action = action;
            Match = match;
            Name = name;
        }
    }
}
