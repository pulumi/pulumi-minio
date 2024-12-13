// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Minio
{
    public static class GetIamPolicyDocument
    {
        /// <summary>
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Minio = Pulumi.Minio;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Minio.GetIamPolicyDocument.Invoke(new()
        ///     {
        ///         Statements = new[]
        ///         {
        ///             new Minio.Inputs.GetIamPolicyDocumentStatementInputArgs
        ///             {
        ///                 Sid = "1",
        ///                 Actions = new[]
        ///                 {
        ///                     "s3:ListAllMyBuckets",
        ///                     "s3:GetBucketLocation",
        ///                 },
        ///                 Resources = new[]
        ///                 {
        ///                     "arn:aws:s3:::*",
        ///                 },
        ///             },
        ///             new Minio.Inputs.GetIamPolicyDocumentStatementInputArgs
        ///             {
        ///                 Actions = new[]
        ///                 {
        ///                     "s3:ListBucket",
        ///                 },
        ///                 Resources = new[]
        ///                 {
        ///                     "arn:aws:s3:::state-terraform-s3",
        ///                 },
        ///                 Conditions = new[]
        ///                 {
        ///                     new Minio.Inputs.GetIamPolicyDocumentStatementConditionInputArgs
        ///                     {
        ///                         Test = "StringLike",
        ///                         Variable = "s3:prefix",
        ///                         Values = new[]
        ///                         {
        ///                             "",
        ///                             "home/",
        ///                         },
        ///                     },
        ///                 },
        ///             },
        ///             new Minio.Inputs.GetIamPolicyDocumentStatementInputArgs
        ///             {
        ///                 Actions = new[]
        ///                 {
        ///                     "s3:PutObject",
        ///                 },
        ///                 Resources = new[]
        ///                 {
        ///                     "arn:aws:s3:::state-terraform-s3",
        ///                     "arn:aws:s3:::state-terraform-s3/*",
        ///                 },
        ///             },
        ///         },
        ///     });
        /// 
        ///     var testPolicy = new Minio.IamPolicy("test_policy", new()
        ///     {
        ///         Name = "state-terraform-s3",
        ///         Policy = example.Apply(getIamPolicyDocumentResult =&gt; getIamPolicyDocumentResult.Json),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetIamPolicyDocumentResult> InvokeAsync(GetIamPolicyDocumentArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIamPolicyDocumentResult>("minio:index/getIamPolicyDocument:getIamPolicyDocument", args ?? new GetIamPolicyDocumentArgs(), options.WithDefaults());

        /// <summary>
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Minio = Pulumi.Minio;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Minio.GetIamPolicyDocument.Invoke(new()
        ///     {
        ///         Statements = new[]
        ///         {
        ///             new Minio.Inputs.GetIamPolicyDocumentStatementInputArgs
        ///             {
        ///                 Sid = "1",
        ///                 Actions = new[]
        ///                 {
        ///                     "s3:ListAllMyBuckets",
        ///                     "s3:GetBucketLocation",
        ///                 },
        ///                 Resources = new[]
        ///                 {
        ///                     "arn:aws:s3:::*",
        ///                 },
        ///             },
        ///             new Minio.Inputs.GetIamPolicyDocumentStatementInputArgs
        ///             {
        ///                 Actions = new[]
        ///                 {
        ///                     "s3:ListBucket",
        ///                 },
        ///                 Resources = new[]
        ///                 {
        ///                     "arn:aws:s3:::state-terraform-s3",
        ///                 },
        ///                 Conditions = new[]
        ///                 {
        ///                     new Minio.Inputs.GetIamPolicyDocumentStatementConditionInputArgs
        ///                     {
        ///                         Test = "StringLike",
        ///                         Variable = "s3:prefix",
        ///                         Values = new[]
        ///                         {
        ///                             "",
        ///                             "home/",
        ///                         },
        ///                     },
        ///                 },
        ///             },
        ///             new Minio.Inputs.GetIamPolicyDocumentStatementInputArgs
        ///             {
        ///                 Actions = new[]
        ///                 {
        ///                     "s3:PutObject",
        ///                 },
        ///                 Resources = new[]
        ///                 {
        ///                     "arn:aws:s3:::state-terraform-s3",
        ///                     "arn:aws:s3:::state-terraform-s3/*",
        ///                 },
        ///             },
        ///         },
        ///     });
        /// 
        ///     var testPolicy = new Minio.IamPolicy("test_policy", new()
        ///     {
        ///         Name = "state-terraform-s3",
        ///         Policy = example.Apply(getIamPolicyDocumentResult =&gt; getIamPolicyDocumentResult.Json),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetIamPolicyDocumentResult> Invoke(GetIamPolicyDocumentInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIamPolicyDocumentResult>("minio:index/getIamPolicyDocument:getIamPolicyDocument", args ?? new GetIamPolicyDocumentInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Minio = Pulumi.Minio;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Minio.GetIamPolicyDocument.Invoke(new()
        ///     {
        ///         Statements = new[]
        ///         {
        ///             new Minio.Inputs.GetIamPolicyDocumentStatementInputArgs
        ///             {
        ///                 Sid = "1",
        ///                 Actions = new[]
        ///                 {
        ///                     "s3:ListAllMyBuckets",
        ///                     "s3:GetBucketLocation",
        ///                 },
        ///                 Resources = new[]
        ///                 {
        ///                     "arn:aws:s3:::*",
        ///                 },
        ///             },
        ///             new Minio.Inputs.GetIamPolicyDocumentStatementInputArgs
        ///             {
        ///                 Actions = new[]
        ///                 {
        ///                     "s3:ListBucket",
        ///                 },
        ///                 Resources = new[]
        ///                 {
        ///                     "arn:aws:s3:::state-terraform-s3",
        ///                 },
        ///                 Conditions = new[]
        ///                 {
        ///                     new Minio.Inputs.GetIamPolicyDocumentStatementConditionInputArgs
        ///                     {
        ///                         Test = "StringLike",
        ///                         Variable = "s3:prefix",
        ///                         Values = new[]
        ///                         {
        ///                             "",
        ///                             "home/",
        ///                         },
        ///                     },
        ///                 },
        ///             },
        ///             new Minio.Inputs.GetIamPolicyDocumentStatementInputArgs
        ///             {
        ///                 Actions = new[]
        ///                 {
        ///                     "s3:PutObject",
        ///                 },
        ///                 Resources = new[]
        ///                 {
        ///                     "arn:aws:s3:::state-terraform-s3",
        ///                     "arn:aws:s3:::state-terraform-s3/*",
        ///                 },
        ///             },
        ///         },
        ///     });
        /// 
        ///     var testPolicy = new Minio.IamPolicy("test_policy", new()
        ///     {
        ///         Name = "state-terraform-s3",
        ///         Policy = example.Apply(getIamPolicyDocumentResult =&gt; getIamPolicyDocumentResult.Json),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetIamPolicyDocumentResult> Invoke(GetIamPolicyDocumentInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetIamPolicyDocumentResult>("minio:index/getIamPolicyDocument:getIamPolicyDocument", args ?? new GetIamPolicyDocumentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIamPolicyDocumentArgs : global::Pulumi.InvokeArgs
    {
        [Input("overrideJson")]
        public string? OverrideJson { get; set; }

        [Input("policyId")]
        public string? PolicyId { get; set; }

        [Input("sourceJson")]
        public string? SourceJson { get; set; }

        [Input("statements")]
        private List<Inputs.GetIamPolicyDocumentStatementArgs>? _statements;
        public List<Inputs.GetIamPolicyDocumentStatementArgs> Statements
        {
            get => _statements ?? (_statements = new List<Inputs.GetIamPolicyDocumentStatementArgs>());
            set => _statements = value;
        }

        [Input("version")]
        public string? Version { get; set; }

        public GetIamPolicyDocumentArgs()
        {
        }
        public static new GetIamPolicyDocumentArgs Empty => new GetIamPolicyDocumentArgs();
    }

    public sealed class GetIamPolicyDocumentInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("overrideJson")]
        public Input<string>? OverrideJson { get; set; }

        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        [Input("sourceJson")]
        public Input<string>? SourceJson { get; set; }

        [Input("statements")]
        private InputList<Inputs.GetIamPolicyDocumentStatementInputArgs>? _statements;
        public InputList<Inputs.GetIamPolicyDocumentStatementInputArgs> Statements
        {
            get => _statements ?? (_statements = new InputList<Inputs.GetIamPolicyDocumentStatementInputArgs>());
            set => _statements = value;
        }

        [Input("version")]
        public Input<string>? Version { get; set; }

        public GetIamPolicyDocumentInvokeArgs()
        {
        }
        public static new GetIamPolicyDocumentInvokeArgs Empty => new GetIamPolicyDocumentInvokeArgs();
    }


    [OutputType]
    public sealed class GetIamPolicyDocumentResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Json;
        public readonly string? OverrideJson;
        public readonly string? PolicyId;
        public readonly string? SourceJson;
        public readonly ImmutableArray<Outputs.GetIamPolicyDocumentStatementResult> Statements;
        public readonly string? Version;

        [OutputConstructor]
        private GetIamPolicyDocumentResult(
            string id,

            string json,

            string? overrideJson,

            string? policyId,

            string? sourceJson,

            ImmutableArray<Outputs.GetIamPolicyDocumentStatementResult> statements,

            string? version)
        {
            Id = id;
            Json = json;
            OverrideJson = overrideJson;
            PolicyId = policyId;
            SourceJson = sourceJson;
            Statements = statements;
            Version = version;
        }
    }
}
