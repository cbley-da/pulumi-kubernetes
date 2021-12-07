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
    /// ObjectMetricStatus indicates the current value of a metric describing a kubernetes object (for example, hits-per-second on an Ingress object).
    /// </summary>
    [OutputType]
    public sealed class ObjectMetricStatus
    {
        /// <summary>
        /// current contains the current value for the given metric
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2.MetricValueStatus Current;
        /// <summary>
        /// DescribedObject specifies the descriptions of a object,such as kind,name apiVersion
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2.CrossVersionObjectReference DescribedObject;
        /// <summary>
        /// metric identifies the target metric by name and selector
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2.MetricIdentifier Metric;

        [OutputConstructor]
        private ObjectMetricStatus(
            Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2.MetricValueStatus current,

            Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2.CrossVersionObjectReference describedObject,

            Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2.MetricIdentifier metric)
        {
            Current = current;
            DescribedObject = describedObject;
            Metric = metric;
        }
    }
}
