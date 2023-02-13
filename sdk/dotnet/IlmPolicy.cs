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
    /// `minio.IlmPolicy` handles lifecycle settings for a given `minio.S3Bucket`.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Minio = Pulumi.Minio;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var bucket = new Minio.S3Bucket("bucket", new()
    ///     {
    ///         Bucket = "bucket",
    ///     });
    /// 
    ///     var bucket_lifecycle_rules = new Minio.IlmPolicy("bucket-lifecycle-rules", new()
    ///     {
    ///         Bucket = bucket.Bucket,
    ///         Rules = new[]
    ///         {
    ///             new Minio.Inputs.IlmPolicyRuleArgs
    ///             {
    ///                 Id = "expire-7d",
    ///                 Expiration = "7d",
    ///                 Filter = "prefix/",
    ///                 Tags = 
    ///                 {
    ///                     { "app", "myapp" },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [MinioResourceType("minio:index/ilmPolicy:IlmPolicy")]
    public partial class IlmPolicy : global::Pulumi.CustomResource
    {
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        [Output("rules")]
        public Output<ImmutableArray<Outputs.IlmPolicyRule>> Rules { get; private set; } = null!;


        /// <summary>
        /// Create a IlmPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IlmPolicy(string name, IlmPolicyArgs args, CustomResourceOptions? options = null)
            : base("minio:index/ilmPolicy:IlmPolicy", name, args ?? new IlmPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IlmPolicy(string name, Input<string> id, IlmPolicyState? state = null, CustomResourceOptions? options = null)
            : base("minio:index/ilmPolicy:IlmPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IlmPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IlmPolicy Get(string name, Input<string> id, IlmPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new IlmPolicy(name, id, state, options);
        }
    }

    public sealed class IlmPolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        [Input("rules", required: true)]
        private InputList<Inputs.IlmPolicyRuleArgs>? _rules;
        public InputList<Inputs.IlmPolicyRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.IlmPolicyRuleArgs>());
            set => _rules = value;
        }

        public IlmPolicyArgs()
        {
        }
        public static new IlmPolicyArgs Empty => new IlmPolicyArgs();
    }

    public sealed class IlmPolicyState : global::Pulumi.ResourceArgs
    {
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        [Input("rules")]
        private InputList<Inputs.IlmPolicyRuleGetArgs>? _rules;
        public InputList<Inputs.IlmPolicyRuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.IlmPolicyRuleGetArgs>());
            set => _rules = value;
        }

        public IlmPolicyState()
        {
        }
        public static new IlmPolicyState Empty => new IlmPolicyState();
    }
}
