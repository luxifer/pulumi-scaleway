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
    /// Creates and manages Scaleway Function Token.
    /// For more information see [the documentation](https://developers.scaleway.com/en/products/functions/api/#tokens-26b085).
    /// 
    /// ## Examples
    /// 
    /// ### Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Scaleway = luxifer.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var mainFunctionNamespace = new Scaleway.FunctionNamespace("mainFunctionNamespace");
    /// 
    ///     var mainFunction = new Scaleway.Function("mainFunction", new()
    ///     {
    ///         NamespaceId = mainFunctionNamespace.Id,
    ///         Runtime = "go118",
    ///         Handler = "Handle",
    ///         Privacy = "private",
    ///     });
    /// 
    ///     // Namespace Token
    ///     var @namespace = new Scaleway.FunctionToken("namespace", new()
    ///     {
    ///         NamespaceId = mainFunctionNamespace.Id,
    ///         ExpiresAt = "2022-10-18T11:35:15+02:00",
    ///     });
    /// 
    ///     // Function Token
    ///     var function = new Scaleway.FunctionToken("function", new()
    ///     {
    ///         FunctionId = mainFunction.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Tokens can be imported using the `{region}/{id}`, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:index/functionToken:FunctionToken main fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/functionToken:FunctionToken")]
    public partial class FunctionToken : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the token.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The expiration date of the token.
        /// </summary>
        [Output("expiresAt")]
        public Output<string?> ExpiresAt { get; private set; } = null!;

        /// <summary>
        /// The ID of the function.
        /// </summary>
        [Output("functionId")]
        public Output<string?> FunctionId { get; private set; } = null!;

        /// <summary>
        /// The ID of the function namespace.
        /// </summary>
        [Output("namespaceId")]
        public Output<string?> NamespaceId { get; private set; } = null!;

        /// <summary>
        /// `region`). The region in which the namespace should be created.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The token.
        /// </summary>
        [Output("token")]
        public Output<string> Token { get; private set; } = null!;


        /// <summary>
        /// Create a FunctionToken resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FunctionToken(string name, FunctionTokenArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway:index/functionToken:FunctionToken", name, args ?? new FunctionTokenArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FunctionToken(string name, Input<string> id, FunctionTokenState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/functionToken:FunctionToken", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FunctionToken resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FunctionToken Get(string name, Input<string> id, FunctionTokenState? state = null, CustomResourceOptions? options = null)
        {
            return new FunctionToken(name, id, state, options);
        }
    }

    public sealed class FunctionTokenArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the token.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The expiration date of the token.
        /// </summary>
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        /// <summary>
        /// The ID of the function.
        /// </summary>
        [Input("functionId")]
        public Input<string>? FunctionId { get; set; }

        /// <summary>
        /// The ID of the function namespace.
        /// </summary>
        [Input("namespaceId")]
        public Input<string>? NamespaceId { get; set; }

        /// <summary>
        /// `region`). The region in which the namespace should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public FunctionTokenArgs()
        {
        }
        public static new FunctionTokenArgs Empty => new FunctionTokenArgs();
    }

    public sealed class FunctionTokenState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the token.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The expiration date of the token.
        /// </summary>
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        /// <summary>
        /// The ID of the function.
        /// </summary>
        [Input("functionId")]
        public Input<string>? FunctionId { get; set; }

        /// <summary>
        /// The ID of the function namespace.
        /// </summary>
        [Input("namespaceId")]
        public Input<string>? NamespaceId { get; set; }

        /// <summary>
        /// `region`). The region in which the namespace should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The token.
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        public FunctionTokenState()
        {
        }
        public static new FunctionTokenState Empty => new FunctionTokenState();
    }
}
