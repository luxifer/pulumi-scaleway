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
    /// <summary>
    /// Creates and manages Scaleway Domain zone.\
    /// For more information, see [the documentation](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/configure-dns-zones/).
    /// 
    /// ## Examples
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Scaleway = luxifer.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test = new Scaleway.DomainZone("test", new()
    ///     {
    ///         Domain = "scaleway-terraform.com",
    ///         Subdomain = "test",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Zone can be imported using the `{subdomain}.{domain}`, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:index/domainZone:DomainZone test test.scaleway-terraform.com
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/domainZone:DomainZone")]
    public partial class DomainZone : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The domain where the DNS zone will be created.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// Message
        /// </summary>
        [Output("message")]
        public Output<string> Message { get; private set; } = null!;

        /// <summary>
        /// NameServer list for zone.
        /// </summary>
        [Output("ns")]
        public Output<ImmutableArray<string>> Ns { get; private set; } = null!;

        /// <summary>
        /// NameServer default list for zone.
        /// </summary>
        [Output("nsDefaults")]
        public Output<ImmutableArray<string>> NsDefaults { get; private set; } = null!;

        /// <summary>
        /// NameServer master list for zone.
        /// </summary>
        [Output("nsMasters")]
        public Output<ImmutableArray<string>> NsMasters { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the domain is associated with.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The domain zone status.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The subdomain(zone name) to create in the domain.
        /// </summary>
        [Output("subdomain")]
        public Output<string> Subdomain { get; private set; } = null!;

        /// <summary>
        /// The date and time of the last update of the DNS zone.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a DomainZone resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DomainZone(string name, DomainZoneArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/domainZone:DomainZone", name, args ?? new DomainZoneArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DomainZone(string name, Input<string> id, DomainZoneState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/domainZone:DomainZone", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/luxifer/pulumi-scaleway",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DomainZone resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DomainZone Get(string name, Input<string> id, DomainZoneState? state = null, CustomResourceOptions? options = null)
        {
            return new DomainZone(name, id, state, options);
        }
    }

    public sealed class DomainZoneArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain where the DNS zone will be created.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the domain is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The subdomain(zone name) to create in the domain.
        /// </summary>
        [Input("subdomain", required: true)]
        public Input<string> Subdomain { get; set; } = null!;

        public DomainZoneArgs()
        {
        }
        public static new DomainZoneArgs Empty => new DomainZoneArgs();
    }

    public sealed class DomainZoneState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain where the DNS zone will be created.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Message
        /// </summary>
        [Input("message")]
        public Input<string>? Message { get; set; }

        [Input("ns")]
        private InputList<string>? _ns;

        /// <summary>
        /// NameServer list for zone.
        /// </summary>
        public InputList<string> Ns
        {
            get => _ns ?? (_ns = new InputList<string>());
            set => _ns = value;
        }

        [Input("nsDefaults")]
        private InputList<string>? _nsDefaults;

        /// <summary>
        /// NameServer default list for zone.
        /// </summary>
        public InputList<string> NsDefaults
        {
            get => _nsDefaults ?? (_nsDefaults = new InputList<string>());
            set => _nsDefaults = value;
        }

        [Input("nsMasters")]
        private InputList<string>? _nsMasters;

        /// <summary>
        /// NameServer master list for zone.
        /// </summary>
        public InputList<string> NsMasters
        {
            get => _nsMasters ?? (_nsMasters = new InputList<string>());
            set => _nsMasters = value;
        }

        /// <summary>
        /// `project_id`) The ID of the project the domain is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The domain zone status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The subdomain(zone name) to create in the domain.
        /// </summary>
        [Input("subdomain")]
        public Input<string>? Subdomain { get; set; }

        /// <summary>
        /// The date and time of the last update of the DNS zone.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public DomainZoneState()
        {
        }
        public static new DomainZoneState Empty => new DomainZoneState();
    }
}
