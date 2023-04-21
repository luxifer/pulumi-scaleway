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
    public sealed class IotRouteDatabase
    {
        public readonly string Dbname;
        public readonly string Host;
        public readonly string Password;
        public readonly int Port;
        public readonly string Query;
        public readonly string Username;

        [OutputConstructor]
        private IotRouteDatabase(
            string dbname,

            string host,

            string password,

            int port,

            string query,

            string username)
        {
            Dbname = dbname;
            Host = host;
            Password = password;
            Port = port;
            Query = query;
            Username = username;
        }
    }
}
