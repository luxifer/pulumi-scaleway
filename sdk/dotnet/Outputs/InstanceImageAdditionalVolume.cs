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
    public sealed class InstanceImageAdditionalVolume
    {
        /// <summary>
        /// Date of the volume creation.
        /// </summary>
        public readonly string? CreationDate;
        /// <summary>
        /// The export URI of the volume.
        /// </summary>
        public readonly string? ExportUri;
        /// <summary>
        /// ID of the server containing the volume.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Date of volume latest update.
        /// </summary>
        public readonly string? ModificationDate;
        /// <summary>
        /// The name of the image. If not provided it will be randomly generated.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The organization ID the volume is associated with.
        /// </summary>
        public readonly string? Organization;
        /// <summary>
        /// ID of the project the volume is associated with
        /// </summary>
        public readonly string? Project;
        /// <summary>
        /// Description of the server containing the volume (in case the image is a backup from a server).
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Server;
        /// <summary>
        /// The size of the volume.
        /// </summary>
        public readonly int? Size;
        /// <summary>
        /// State of the volume.
        /// </summary>
        public readonly string? State;
        /// <summary>
        /// A list of tags to apply to the image.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// The type of volume, possible values are `l_ssd` and `b_ssd`.
        /// </summary>
        public readonly string? VolumeType;
        /// <summary>
        /// The zone in which the image should be created.
        /// </summary>
        public readonly string? Zone;

        [OutputConstructor]
        private InstanceImageAdditionalVolume(
            string? creationDate,

            string? exportUri,

            string? id,

            string? modificationDate,

            string? name,

            string? organization,

            string? project,

            ImmutableDictionary<string, string>? server,

            int? size,

            string? state,

            ImmutableArray<string> tags,

            string? volumeType,

            string? zone)
        {
            CreationDate = creationDate;
            ExportUri = exportUri;
            Id = id;
            ModificationDate = modificationDate;
            Name = name;
            Organization = organization;
            Project = project;
            Server = server;
            Size = size;
            State = state;
            Tags = tags;
            VolumeType = volumeType;
            Zone = zone;
        }
    }
}
