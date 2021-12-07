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
    /// PersistentVolumeClaimStatus is the current status of a persistent volume claim.
    /// </summary>
    public class PersistentVolumeClaimStatusArgs : Pulumi.ResourceArgs
    {
        [Input("accessModes")]
        private InputList<string>? _accessModes;

        /// <summary>
        /// AccessModes contains the actual access modes the volume backing the PVC has. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1
        /// </summary>
        public InputList<string> AccessModes
        {
            get => _accessModes ?? (_accessModes = new InputList<string>());
            set => _accessModes = value;
        }

        [Input("allocatedResources")]
        private InputMap<string>? _allocatedResources;

        /// <summary>
        /// The storage resource within AllocatedResources tracks the capacity allocated to a PVC. It may be larger than the actual capacity when a volume expansion operation is requested. For storage quota, the larger value from allocatedResources and PVC.spec.resources is used. If allocatedResources is not set, PVC.spec.resources alone is used for quota calculation. If a volume expansion capacity request is lowered, allocatedResources is only lowered if there are no expansion operations in progress and if the actual volume capacity is equal or lower than the requested capacity. This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.
        /// </summary>
        public InputMap<string> AllocatedResources
        {
            get => _allocatedResources ?? (_allocatedResources = new InputMap<string>());
            set => _allocatedResources = value;
        }

        [Input("capacity")]
        private InputMap<string>? _capacity;

        /// <summary>
        /// Represents the actual resources of the underlying volume.
        /// </summary>
        public InputMap<string> Capacity
        {
            get => _capacity ?? (_capacity = new InputMap<string>());
            set => _capacity = value;
        }

        [Input("conditions")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.PersistentVolumeClaimConditionArgs>? _conditions;

        /// <summary>
        /// Current Condition of persistent volume claim. If underlying persistent volume is being resized then the Condition will be set to 'ResizeStarted'.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.PersistentVolumeClaimConditionArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.PersistentVolumeClaimConditionArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// Phase represents the current phase of PersistentVolumeClaim.
        /// 
        /// Possible enum values:
        ///  - `"Bound"` used for PersistentVolumeClaims that are bound
        ///  - `"Lost"` used for PersistentVolumeClaims that lost their underlying PersistentVolume. The claim was bound to a PersistentVolume and this volume does not exist any longer and all data on it was lost.
        ///  - `"Pending"` used for PersistentVolumeClaims that are not yet bound
        /// </summary>
        [Input("phase")]
        public Input<string>? Phase { get; set; }

        /// <summary>
        /// ResizeStatus stores status of resize operation. ResizeStatus is not set by default but when expansion is complete resizeStatus is set to empty string by resize controller or kubelet. This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.
        /// </summary>
        [Input("resizeStatus")]
        public Input<string>? ResizeStatus { get; set; }

        public PersistentVolumeClaimStatusArgs()
        {
        }
    }
}
