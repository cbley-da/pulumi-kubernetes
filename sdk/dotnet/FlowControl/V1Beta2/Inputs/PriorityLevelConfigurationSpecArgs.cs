// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.FlowControl.V1Beta2
{

    /// <summary>
    /// PriorityLevelConfigurationSpec specifies the configuration of a priority level.
    /// </summary>
    public class PriorityLevelConfigurationSpecArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// `limited` specifies how requests are handled for a Limited priority level. This field must be non-empty if and only if `type` is `"Limited"`.
        /// </summary>
        [Input("limited")]
        public Input<Pulumi.Kubernetes.Types.Inputs.FlowControl.V1Beta2.LimitedPriorityLevelConfigurationArgs>? Limited { get; set; }

        /// <summary>
        /// `type` indicates whether this priority level is subject to limitation on request execution.  A value of `"Exempt"` means that requests of this priority level are not subject to a limit (and thus are never queued) and do not detract from the capacity made available to other priority levels.  A value of `"Limited"` means that (a) requests of this priority level _are_ subject to limits and (b) some of the server's limited capacity is made available exclusively to this priority level. Required.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public PriorityLevelConfigurationSpecArgs()
        {
        }
        public static new PriorityLevelConfigurationSpecArgs Empty => new PriorityLevelConfigurationSpecArgs();
    }
}
