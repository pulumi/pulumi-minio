// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Minio.Inputs
{

    public sealed class GetIamPolicyDocumentStatementConditionInputArgs : Pulumi.ResourceArgs
    {
        [Input("test", required: true)]
        public Input<string> Test { get; set; } = null!;

        [Input("values", required: true)]
        private InputList<string>? _values;
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        [Input("variable", required: true)]
        public Input<string> Variable { get; set; } = null!;

        public GetIamPolicyDocumentStatementConditionInputArgs()
        {
        }
    }
}
