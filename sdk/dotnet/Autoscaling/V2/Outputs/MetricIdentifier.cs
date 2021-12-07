// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2
{

    /// <summary>
    /// MetricIdentifier defines the name and optionally selector for a metric
    /// </summary>
    [OutputType]
    public sealed class MetricIdentifier
    {
        /// <summary>
        /// name is the name of the given metric
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Meta.V1.LabelSelector Selector;

        [OutputConstructor]
        private MetricIdentifier(
            string name,

            Pulumi.Kubernetes.Types.Outputs.Meta.V1.LabelSelector selector)
        {
            Name = name;
            Selector = selector;
        }
    }
}
