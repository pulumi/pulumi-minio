// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Minio
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Minio = Pulumi.Minio;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var testPolicy = new Minio.IamPolicy("testPolicy", new()
    ///     {
    ///         Policy = @"{
    ///   ""Version"":""2012-10-17"",
    ///   ""Statement"": [
    ///     {
    ///       ""Sid"":""ListAllBucket"",
    ///       ""Effect"": ""Allow"",
    ///       ""Action"": [""s3:PutObject""],
    ///       ""Principal"":""*"",
    ///       ""Resource"": ""arn:aws:s3:::state-terraform-s3/*""
    ///     }
    ///   ]
    /// }
    /// ",
    ///     });
    /// 
    ///     return new Dictionary&lt;string, object?&gt;
    ///     {
    ///         ["minioId"] = testPolicy.Id,
    ///         ["minioPolicy"] = testPolicy.Policy,
    ///     };
    /// });
    /// ```
    /// </summary>
    [MinioResourceType("minio:index/iamPolicy:IamPolicy")]
    public partial class IamPolicy : global::Pulumi.CustomResource
    {
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("namePrefix")]
        public Output<string?> NamePrefix { get; private set; } = null!;

        [Output("policy")]
        public Output<string> Policy { get; private set; } = null!;


        /// <summary>
        /// Create a IamPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IamPolicy(string name, IamPolicyArgs args, CustomResourceOptions? options = null)
            : base("minio:index/iamPolicy:IamPolicy", name, args ?? new IamPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IamPolicy(string name, Input<string> id, IamPolicyState? state = null, CustomResourceOptions? options = null)
            : base("minio:index/iamPolicy:IamPolicy", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IamPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IamPolicy Get(string name, Input<string> id, IamPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new IamPolicy(name, id, state, options);
        }
    }

    public sealed class IamPolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        [Input("policy", required: true)]
        public Input<string> Policy { get; set; } = null!;

        public IamPolicyArgs()
        {
        }
        public static new IamPolicyArgs Empty => new IamPolicyArgs();
    }

    public sealed class IamPolicyState : global::Pulumi.ResourceArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        [Input("policy")]
        public Input<string>? Policy { get; set; }

        public IamPolicyState()
        {
        }
        public static new IamPolicyState Empty => new IamPolicyState();
    }
}
