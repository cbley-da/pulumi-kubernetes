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
    /// NodeCondition contains condition information for a node.
    /// </summary>
    [OutputType]
    public sealed class NodeCondition
    {
        /// <summary>
        /// Last time we got an update on a given condition.
        /// </summary>
        public readonly string LastHeartbeatTime;
        /// <summary>
        /// Last time the condition transit from one status to another.
        /// </summary>
        public readonly string LastTransitionTime;
        /// <summary>
        /// Human readable message indicating details about last transition.
        /// </summary>
        public readonly string Message;
        /// <summary>
        /// (brief) reason for the condition's last transition.
        /// </summary>
        public readonly string Reason;
        /// <summary>
        /// Status of the condition, one of True, False, Unknown.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Type of node condition.
        /// 
        /// Possible enum values:
        ///  - `"DiskPressure"` means the kubelet is under pressure due to insufficient available disk.
        ///  - `"MemoryPressure"` means the kubelet is under pressure due to insufficient available memory.
        ///  - `"NetworkUnavailable"` means that network for the node is not correctly configured.
        ///  - `"PIDPressure"` means the kubelet is under pressure due to insufficient available PID.
        ///  - `"Ready"` means kubelet is healthy and ready to accept pods.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private NodeCondition(
            string lastHeartbeatTime,

            string lastTransitionTime,

            string message,

            string reason,

            string status,

            string type)
        {
            LastHeartbeatTime = lastHeartbeatTime;
            LastTransitionTime = lastTransitionTime;
            Message = message;
            Reason = reason;
            Status = status;
            Type = type;
        }
    }
}
