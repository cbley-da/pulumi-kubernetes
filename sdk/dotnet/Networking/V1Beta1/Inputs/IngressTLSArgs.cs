// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Networking.V1Beta1
{

    /// <summary>
    /// IngressTLS describes the transport layer security associated with an Ingress.
    /// </summary>
    public class IngressTLSArgs : Pulumi.ResourceArgs
    {
        [Input("hosts")]
        private InputList<string>? _hosts;

        /// <summary>
        /// Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.
        /// </summary>
        public InputList<string> Hosts
        {
            get => _hosts ?? (_hosts = new InputList<string>());
            set => _hosts = value;
        }

        /// <summary>
        /// SecretName is the name of the secret used to terminate SSL traffic on 443. Field is left optional to allow SSL routing based on SNI hostname alone. If the SNI host in a listener conflicts with the "Host" header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing.
        /// </summary>
        [Input("secretName")]
        public Input<string>? SecretName { get; set; }

        public IngressTLSArgs()
        {
        }
    }
}
