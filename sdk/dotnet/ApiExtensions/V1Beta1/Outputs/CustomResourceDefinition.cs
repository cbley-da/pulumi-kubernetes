// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.ApiExtensions.V1Beta1
{

    /// <summary>
    /// CustomResourceDefinition represents a resource that should be exposed on the API server.  Its name MUST be in the format &lt;.spec.name&gt;.&lt;.spec.group&gt;. Deprecated in v1.16, planned for removal in v1.19. Use apiextensions.k8s.io/v1 CustomResourceDefinition instead.
    /// </summary>
    [OutputType]
    public sealed class CustomResourceDefinition
    {
        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        public readonly string ApiVersion;
        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        public readonly string Kind;
        public readonly Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMeta Metadata;
        /// <summary>
        /// spec describes how the user wants the resources to appear
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.ApiExtensions.V1Beta1.CustomResourceDefinitionSpec Spec;
        /// <summary>
        /// status indicates the actual state of the CustomResourceDefinition
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.ApiExtensions.V1Beta1.CustomResourceDefinitionStatus Status;

        [OutputConstructor]
        private CustomResourceDefinition(
            string apiVersion,

            string kind,

            Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMeta metadata,

            Pulumi.Kubernetes.Types.Outputs.ApiExtensions.V1Beta1.CustomResourceDefinitionSpec spec,

            Pulumi.Kubernetes.Types.Outputs.ApiExtensions.V1Beta1.CustomResourceDefinitionStatus status)
        {
            ApiVersion = apiVersion;
            Kind = kind;
            Metadata = metadata;
            Spec = spec;
            Status = status;
        }
    }
}
