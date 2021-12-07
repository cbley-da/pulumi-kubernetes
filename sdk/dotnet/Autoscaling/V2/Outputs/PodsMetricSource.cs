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
    /// PodsMetricSource indicates how to scale on a metric describing each pod in the current scale target (for example, transactions-processed-per-second). The values will be averaged together before being compared to the target value.
    /// </summary>
    [OutputType]
    public sealed class PodsMetricSource
    {
        /// <summary>
        /// metric identifies the target metric by name and selector
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2.MetricIdentifier Metric;
        /// <summary>
        /// target specifies the target value for the given metric
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2.MetricTarget Target;

        [OutputConstructor]
        private PodsMetricSource(
            Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2.MetricIdentifier metric,

            Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2.MetricTarget target)
        {
            Metric = metric;
            Target = target;
        }
    }
}
