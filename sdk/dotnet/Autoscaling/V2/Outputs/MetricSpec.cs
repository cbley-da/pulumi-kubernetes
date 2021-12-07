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
    /// MetricSpec specifies how to scale based on a single metric (only `type` and one other matching field should be set at once).
    /// </summary>
    [OutputType]
    public sealed class MetricSpec
    {
        /// <summary>
        /// containerResource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing a single container in each pod of the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source. This is an alpha feature and can be enabled by the HPAContainerMetrics feature flag.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2.ContainerResourceMetricSource ContainerResource;
        /// <summary>
        /// external refers to a global metric that is not associated with any Kubernetes object. It allows autoscaling based on information coming from components running outside of cluster (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2.ExternalMetricSource External;
        /// <summary>
        /// object refers to a metric describing a single kubernetes object (for example, hits-per-second on an Ingress object).
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2.ObjectMetricSource Object;
        /// <summary>
        /// pods refers to a metric describing each pod in the current scale target (for example, transactions-processed-per-second).  The values will be averaged together before being compared to the target value.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2.PodsMetricSource Pods;
        /// <summary>
        /// resource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2.ResourceMetricSource Resource;
        /// <summary>
        /// type is the type of metric source.  It should be one of "ContainerResource", "External", "Object", "Pods" or "Resource", each mapping to a matching field in the object. Note: "ContainerResource" type is available on when the feature-gate HPAContainerMetrics is enabled
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private MetricSpec(
            Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2.ContainerResourceMetricSource containerResource,

            Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2.ExternalMetricSource external,

            Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2.ObjectMetricSource @object,

            Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2.PodsMetricSource pods,

            Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2.ResourceMetricSource resource,

            string type)
        {
            ContainerResource = containerResource;
            External = external;
            Object = @object;
            Pods = pods;
            Resource = resource;
            Type = type;
        }
    }
}
