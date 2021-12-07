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
    /// UserSubject holds detailed information for user-kind subject.
    /// </summary>
    public class UserSubjectArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// `name` is the username that matches, or "*" to match all usernames. Required.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public UserSubjectArgs()
        {
        }
    }
}
