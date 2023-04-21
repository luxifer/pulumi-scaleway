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

    public sealed class IotRouteDatabaseGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("dbname", required: true)]
        public Input<string> Dbname { get; set; } = null!;

        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        [Input("query", required: true)]
        public Input<string> Query { get; set; } = null!;

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public IotRouteDatabaseGetArgs()
        {
        }
        public static new IotRouteDatabaseGetArgs Empty => new IotRouteDatabaseGetArgs();
    }
}