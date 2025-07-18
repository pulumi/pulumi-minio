// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Minio.Inputs
{

    public sealed class GetIamPolicyDocumentStatementConditionArgs : global::Pulumi.InvokeArgs
    {
        [Input("test", required: true)]
        public string Test { get; set; } = null!;

        [Input("values", required: true)]
        private List<string>? _values;
        public List<string> Values
        {
            get => _values ?? (_values = new List<string>());
            set => _values = value;
        }

        [Input("variable", required: true)]
        public string Variable { get; set; } = null!;

        public GetIamPolicyDocumentStatementConditionArgs()
        {
        }
        public static new GetIamPolicyDocumentStatementConditionArgs Empty => new GetIamPolicyDocumentStatementConditionArgs();
    }
}
