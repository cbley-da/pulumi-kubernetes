// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Events.V1
{

    [OutputType]
    public sealed class Event
    {
        /// <summary>
        /// action is what action was taken/failed regarding to the regarding object. It is machine-readable. This field can have at most 128 characters.
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        public readonly string ApiVersion;
        /// <summary>
        /// deprecatedCount is the deprecated field assuring backward compatibility with core.v1 Event type.
        /// </summary>
        public readonly int DeprecatedCount;
        /// <summary>
        /// deprecatedFirstTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
        /// </summary>
        public readonly string DeprecatedFirstTimestamp;
        /// <summary>
        /// deprecatedLastTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
        /// </summary>
        public readonly string DeprecatedLastTimestamp;
        /// <summary>
        /// deprecatedSource is the deprecated field assuring backward compatibility with core.v1 Event type.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.EventSource DeprecatedSource;
        /// <summary>
        /// eventTime is the time when this Event was first observed. It is required.
        /// </summary>
        public readonly string EventTime;
        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        public readonly string Kind;
        public readonly Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMeta Metadata;
        /// <summary>
        /// note is a human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
        /// </summary>
        public readonly string Note;
        /// <summary>
        /// reason is why the action was taken. It is human-readable. This field can have at most 128 characters.
        /// </summary>
        public readonly string Reason;
        /// <summary>
        /// regarding contains the object this Event is about. In most cases it's an Object reporting controller implements, e.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.ObjectReference Regarding;
        /// <summary>
        /// related is the optional secondary object for more complex actions. E.g. when regarding object triggers a creation or deletion of related object.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.ObjectReference Related;
        /// <summary>
        /// reportingController is the name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`. This field cannot be empty for new Events.
        /// </summary>
        public readonly string ReportingController;
        /// <summary>
        /// reportingInstance is the ID of the controller instance, e.g. `kubelet-xyzf`. This field cannot be empty for new Events and it can have at most 128 characters.
        /// </summary>
        public readonly string ReportingInstance;
        /// <summary>
        /// series is data about the Event series this event represents or nil if it's a singleton Event.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Events.V1.EventSeries Series;
        /// <summary>
        /// type is the type of this event (Normal, Warning), new types could be added in the future. It is machine-readable.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private Event(
            string action,

            string apiVersion,

            int deprecatedCount,

            string deprecatedFirstTimestamp,

            string deprecatedLastTimestamp,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.EventSource deprecatedSource,

            string eventTime,

            string kind,

            Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMeta metadata,

            string note,

            string reason,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.ObjectReference regarding,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.ObjectReference related,

            string reportingController,

            string reportingInstance,

            Pulumi.Kubernetes.Types.Outputs.Events.V1.EventSeries series,

            string type)
        {
            Action = action;
            ApiVersion = apiVersion;
            DeprecatedCount = deprecatedCount;
            DeprecatedFirstTimestamp = deprecatedFirstTimestamp;
            DeprecatedLastTimestamp = deprecatedLastTimestamp;
            DeprecatedSource = deprecatedSource;
            EventTime = eventTime;
            Kind = kind;
            Metadata = metadata;
            Note = note;
            Reason = reason;
            Regarding = regarding;
            Related = related;
            ReportingController = reportingController;
            ReportingInstance = reportingInstance;
            Series = series;
            Type = type;
        }
    }
}
