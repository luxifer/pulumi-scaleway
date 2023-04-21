// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace luxifer.Scaleway
{
    public static class GetAccountSshKey
    {
        /// <summary>
        /// Use this data source to get SSH key information based on its ID or name.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myKey = Scaleway.GetAccountSshKey.Invoke(new()
        ///     {
        ///         SshKeyId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAccountSshKeyResult> InvokeAsync(GetAccountSshKeyArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAccountSshKeyResult>("scaleway:index/getAccountSshKey:getAccountSshKey", args ?? new GetAccountSshKeyArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get SSH key information based on its ID or name.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myKey = Scaleway.GetAccountSshKey.Invoke(new()
        ///     {
        ///         SshKeyId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAccountSshKeyResult> Invoke(GetAccountSshKeyInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAccountSshKeyResult>("scaleway:index/getAccountSshKey:getAccountSshKey", args ?? new GetAccountSshKeyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccountSshKeyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The SSH key name. Only one of `name` and `ssh_key_id` should be specified.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The SSH key id. Only one of `name` and `ssh_key_id` should be specified.
        /// </summary>
        [Input("sshKeyId")]
        public string? SshKeyId { get; set; }

        public GetAccountSshKeyArgs()
        {
        }
        public static new GetAccountSshKeyArgs Empty => new GetAccountSshKeyArgs();
    }

    public sealed class GetAccountSshKeyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The SSH key name. Only one of `name` and `ssh_key_id` should be specified.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The SSH key id. Only one of `name` and `ssh_key_id` should be specified.
        /// </summary>
        [Input("sshKeyId")]
        public Input<string>? SshKeyId { get; set; }

        public GetAccountSshKeyInvokeArgs()
        {
        }
        public static new GetAccountSshKeyInvokeArgs Empty => new GetAccountSshKeyInvokeArgs();
    }


    [OutputType]
    public sealed class GetAccountSshKeyResult
    {
        public readonly string CreatedAt;
        public readonly bool Disabled;
        public readonly string Fingerprint;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Name;
        /// <summary>
        /// The ID of the organization the SSH key is associated with.
        /// </summary>
        public readonly string OrganizationId;
        public readonly string ProjectId;
        /// <summary>
        /// The SSH public key string
        /// </summary>
        public readonly string PublicKey;
        public readonly string? SshKeyId;
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetAccountSshKeyResult(
            string createdAt,

            bool disabled,

            string fingerprint,

            string id,

            string? name,

            string organizationId,

            string projectId,

            string publicKey,

            string? sshKeyId,

            string updatedAt)
        {
            CreatedAt = createdAt;
            Disabled = disabled;
            Fingerprint = fingerprint;
            Id = id;
            Name = name;
            OrganizationId = organizationId;
            ProjectId = projectId;
            PublicKey = publicKey;
            SshKeyId = sshKeyId;
            UpdatedAt = updatedAt;
        }
    }
}
