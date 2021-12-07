// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Core.V1
{

    /// <summary>
    /// LimitRangeItem defines a min/max usage limit for any resource that matches on kind.
    /// </summary>
    public class LimitRangeItemArgs : Pulumi.ResourceArgs
    {
        [Input("default")]
        private InputMap<string>? _default;

        /// <summary>
        /// Default resource requirement limit value by resource name if resource limit is omitted.
        /// </summary>
        public InputMap<string> Default
        {
            get => _default ?? (_default = new InputMap<string>());
            set => _default = value;
        }

        [Input("defaultRequest")]
        private InputMap<string>? _defaultRequest;

        /// <summary>
        /// DefaultRequest is the default resource requirement request value by resource name if resource request is omitted.
        /// </summary>
        public InputMap<string> DefaultRequest
        {
            get => _defaultRequest ?? (_defaultRequest = new InputMap<string>());
            set => _defaultRequest = value;
        }

        [Input("max")]
        private InputMap<string>? _max;

        /// <summary>
        /// Max usage constraints on this kind by resource name.
        /// </summary>
        public InputMap<string> Max
        {
            get => _max ?? (_max = new InputMap<string>());
            set => _max = value;
        }

        [Input("maxLimitRequestRatio")]
        private InputMap<string>? _maxLimitRequestRatio;

        /// <summary>
        /// MaxLimitRequestRatio if specified, the named resource must have a request and limit that are both non-zero where limit divided by request is less than or equal to the enumerated value; this represents the max burst for the named resource.
        /// </summary>
        public InputMap<string> MaxLimitRequestRatio
        {
            get => _maxLimitRequestRatio ?? (_maxLimitRequestRatio = new InputMap<string>());
            set => _maxLimitRequestRatio = value;
        }

        [Input("min")]
        private InputMap<string>? _min;

        /// <summary>
        /// Min usage constraints on this kind by resource name.
        /// </summary>
        public InputMap<string> Min
        {
            get => _min ?? (_min = new InputMap<string>());
            set => _min = value;
        }

        /// <summary>
        /// Type of resource that this limit applies to.
        /// 
        /// Possible enum values:
        ///  - `"Container"` Limit that applies to all containers in a namespace
        ///  - `"PersistentVolumeClaim"` Limit that applies to all persistent volume claims in a namespace
        ///  - `"Pod"` Limit that applies to all pods in a namespace
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public LimitRangeItemArgs()
        {
        }
    }
}
