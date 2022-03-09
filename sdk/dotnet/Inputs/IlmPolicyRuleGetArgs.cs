// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Minio.Inputs
{

    public sealed class IlmPolicyRuleGetArgs : Pulumi.ResourceArgs
    {
        [Input("expiration")]
        public Input<string>? Expiration { get; set; }

        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("status")]
        public Input<string>? Status { get; set; }

        public IlmPolicyRuleGetArgs()
        {
        }
    }
}
