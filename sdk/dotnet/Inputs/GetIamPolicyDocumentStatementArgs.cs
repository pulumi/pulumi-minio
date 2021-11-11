// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Minio.Inputs
{

    public sealed class GetIamPolicyDocumentStatementInputArgs : Pulumi.ResourceArgs
    {
        [Input("actions")]
        private InputList<string>? _actions;
        public InputList<string> Actions
        {
            get => _actions ?? (_actions = new InputList<string>());
            set => _actions = value;
        }

        [Input("conditions")]
        private InputList<Inputs.GetIamPolicyDocumentStatementConditionInputArgs>? _conditions;
        public InputList<Inputs.GetIamPolicyDocumentStatementConditionInputArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.GetIamPolicyDocumentStatementConditionInputArgs>());
            set => _conditions = value;
        }

        [Input("effect")]
        public Input<string>? Effect { get; set; }

        [Input("principal")]
        public Input<string>? Principal { get; set; }

        [Input("resources")]
        private InputList<string>? _resources;
        public InputList<string> Resources
        {
            get => _resources ?? (_resources = new InputList<string>());
            set => _resources = value;
        }

        [Input("sid")]
        public Input<string>? Sid { get; set; }

        public GetIamPolicyDocumentStatementInputArgs()
        {
        }
    }
}
