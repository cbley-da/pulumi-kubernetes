// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Core.V1
{

    /// <summary>
    /// ContainerStatus contains details for the current status of this container.
    /// </summary>
    [OutputType]
    public sealed class ContainerStatus
    {
        /// <summary>
        /// Container's ID in the format 'docker://&lt;container_id&gt;'.
        /// </summary>
        public readonly string ContainerID;
        /// <summary>
        /// The image the container is running. More info: https://kubernetes.io/docs/concepts/containers/images.
        /// </summary>
        public readonly string Image;
        /// <summary>
        /// ImageID of the container's image.
        /// </summary>
        public readonly string ImageID;
        /// <summary>
        /// Details about the container's last termination condition.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.ContainerState LastState;
        /// <summary>
        /// This must be a DNS_LABEL. Each container in a pod must have a unique name. Cannot be updated.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Specifies whether the container has passed its readiness probe.
        /// </summary>
        public readonly bool Ready;
        /// <summary>
        /// The number of times the container has been restarted.
        /// </summary>
        public readonly int RestartCount;
        /// <summary>
        /// Specifies whether the container has passed its startup probe. Initialized as false, becomes true after startupProbe is considered successful. Resets to false when the container is restarted, or if kubelet loses state temporarily. Is always true when no startupProbe is defined.
        /// </summary>
        public readonly bool Started;
        /// <summary>
        /// Details about the container's current condition.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.ContainerState State;

        [OutputConstructor]
        private ContainerStatus(
            string containerID,

            string image,

            string imageID,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.ContainerState lastState,

            string name,

            bool ready,

            int restartCount,

            bool started,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.ContainerState state)
        {
            ContainerID = containerID;
            Image = image;
            ImageID = imageID;
            LastState = lastState;
            Name = name;
            Ready = ready;
            RestartCount = restartCount;
            Started = started;
            State = state;
        }
    }
}
