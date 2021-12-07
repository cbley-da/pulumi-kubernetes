// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2
{

    /// <summary>
    /// ObjectMetricStatus indicates the current value of a metric describing a kubernetes object (for example, hits-per-second on an Ingress object).
    /// </summary>
    public class ObjectMetricStatusArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// current contains the current value for the given metric
        /// </summary>
        [Input("current", required: true)]
        public Input<Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2.MetricValueStatusArgs> Current { get; set; } = null!;

        /// <summary>
        /// DescribedObject specifies the descriptions of a object,such as kind,name apiVersion
        /// </summary>
        [Input("describedObject", required: true)]
        public Input<Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2.CrossVersionObjectReferenceArgs> DescribedObject { get; set; } = null!;

        /// <summary>
        /// metric identifies the target metric by name and selector
        /// </summary>
        [Input("metric", required: true)]
        public Input<Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2.MetricIdentifierArgs> Metric { get; set; } = null!;

        public ObjectMetricStatusArgs()
        {
        }
    }
}
