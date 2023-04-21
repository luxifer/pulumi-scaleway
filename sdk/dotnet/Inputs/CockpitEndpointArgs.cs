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

    public sealed class CockpitEndpointArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The alertmanager URL
        /// </summary>
        [Input("alertmanagerUrl")]
        public Input<string>? AlertmanagerUrl { get; set; }

        /// <summary>
        /// The grafana URL
        /// </summary>
        [Input("grafanaUrl")]
        public Input<string>? GrafanaUrl { get; set; }

        /// <summary>
        /// The logs URL
        /// </summary>
        [Input("logsUrl")]
        public Input<string>? LogsUrl { get; set; }

        /// <summary>
        /// The metrics URL
        /// </summary>
        [Input("metricsUrl")]
        public Input<string>? MetricsUrl { get; set; }

        public CockpitEndpointArgs()
        {
        }
        public static new CockpitEndpointArgs Empty => new CockpitEndpointArgs();
    }
}
