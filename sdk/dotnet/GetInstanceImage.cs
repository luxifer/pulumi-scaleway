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
    public static class GetInstanceImage
    {
        /// <summary>
        /// Gets information about an instance image.
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
        ///     var myImage = Scaleway.GetInstanceImage.Invoke(new()
        ///     {
        ///         ImageId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetInstanceImageResult> InvokeAsync(GetInstanceImageArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceImageResult>("scaleway:index/getInstanceImage:getInstanceImage", args ?? new GetInstanceImageArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about an instance image.
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
        ///     var myImage = Scaleway.GetInstanceImage.Invoke(new()
        ///     {
        ///         ImageId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetInstanceImageResult> Invoke(GetInstanceImageInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInstanceImageResult>("scaleway:index/getInstanceImage:getInstanceImage", args ?? new GetInstanceImageInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceImageArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The architecture the image is compatible with. Possible values are: `x86_64` or `arm`.
        /// </summary>
        [Input("architecture")]
        public string? Architecture { get; set; }

        /// <summary>
        /// The image id. Only one of `name` and `image_id` should be specified.
        /// </summary>
        [Input("imageId")]
        public string? ImageId { get; set; }

        /// <summary>
        /// Use the latest image ID.
        /// </summary>
        [Input("latest")]
        public bool? Latest { get; set; }

        /// <summary>
        /// The image name. Only one of `name` and `image_id` should be specified.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The ID of the project the image is associated with.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// `zone`) The zone in which the image exists.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetInstanceImageArgs()
        {
        }
        public static new GetInstanceImageArgs Empty => new GetInstanceImageArgs();
    }

    public sealed class GetInstanceImageInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The architecture the image is compatible with. Possible values are: `x86_64` or `arm`.
        /// </summary>
        [Input("architecture")]
        public Input<string>? Architecture { get; set; }

        /// <summary>
        /// The image id. Only one of `name` and `image_id` should be specified.
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        /// <summary>
        /// Use the latest image ID.
        /// </summary>
        [Input("latest")]
        public Input<bool>? Latest { get; set; }

        /// <summary>
        /// The image name. Only one of `name` and `image_id` should be specified.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project the image is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `zone`) The zone in which the image exists.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetInstanceImageInvokeArgs()
        {
        }
        public static new GetInstanceImageInvokeArgs Empty => new GetInstanceImageInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstanceImageResult
    {
        /// <summary>
        /// IDs of the additional volumes in this image.
        /// </summary>
        public readonly ImmutableArray<string> AdditionalVolumeIds;
        public readonly string? Architecture;
        /// <summary>
        /// Date of the image creation.
        /// </summary>
        public readonly string CreationDate;
        /// <summary>
        /// ID of the default bootscript for this image.
        /// </summary>
        public readonly string DefaultBootscriptId;
        /// <summary>
        /// ID of the server the image if based from.
        /// </summary>
        public readonly string FromServerId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ImageId;
        public readonly bool? Latest;
        /// <summary>
        /// Date of image latest update.
        /// </summary>
        public readonly string ModificationDate;
        public readonly string? Name;
        /// <summary>
        /// The ID of the organization the image is associated with.
        /// </summary>
        public readonly string OrganizationId;
        /// <summary>
        /// The ID of the project the image is associated with.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// Set to `true` if the image is public.
        /// </summary>
        public readonly bool Public;
        /// <summary>
        /// ID of the root volume in this image.
        /// </summary>
        public readonly string RootVolumeId;
        /// <summary>
        /// State of the image. Possible values are: `available`, `creating` or `error`.
        /// </summary>
        public readonly string State;
        public readonly string Zone;

        [OutputConstructor]
        private GetInstanceImageResult(
            ImmutableArray<string> additionalVolumeIds,

            string? architecture,

            string creationDate,

            string defaultBootscriptId,

            string fromServerId,

            string id,

            string? imageId,

            bool? latest,

            string modificationDate,

            string? name,

            string organizationId,

            string projectId,

            bool @public,

            string rootVolumeId,

            string state,

            string zone)
        {
            AdditionalVolumeIds = additionalVolumeIds;
            Architecture = architecture;
            CreationDate = creationDate;
            DefaultBootscriptId = defaultBootscriptId;
            FromServerId = fromServerId;
            Id = id;
            ImageId = imageId;
            Latest = latest;
            ModificationDate = modificationDate;
            Name = name;
            OrganizationId = organizationId;
            ProjectId = projectId;
            Public = @public;
            RootVolumeId = rootVolumeId;
            State = state;
            Zone = zone;
        }
    }
}
