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
    /// using System.Linq;
    /// using Pulumi;
    /// using Minio = Pulumi.Minio;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var developer = new Minio.IamGroup("developer");
    /// 
    ///     var testPolicy = new Minio.IamGroupPolicy("testPolicy", new()
    ///     {
    ///         Group = developer.Id,
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
    /// 
    /// ",
    ///     });
    /// 
    ///     return new Dictionary&lt;string, object?&gt;
    ///     {
    ///         ["minioName"] = minio_iam_group_membership.Developer.Id,
    ///         ["minioPolicy"] = minio_iam_group_membership.Developer.Policy,
    ///         ["minioGroup"] = minio_iam_group_membership.Developer.Group,
    ///     };
    /// });
    /// ```
    /// </summary>
    [MinioResourceType("minio:index/iamGroupPolicy:IamGroupPolicy")]
    public partial class IamGroupPolicy : global::Pulumi.CustomResource
    {
        [Output("group")]
        public Output<string> Group { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("namePrefix")]
        public Output<string?> NamePrefix { get; private set; } = null!;

        [Output("policy")]
        public Output<string> Policy { get; private set; } = null!;


        /// <summary>
        /// Create a IamGroupPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IamGroupPolicy(string name, IamGroupPolicyArgs args, CustomResourceOptions? options = null)
            : base("minio:index/iamGroupPolicy:IamGroupPolicy", name, args ?? new IamGroupPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IamGroupPolicy(string name, Input<string> id, IamGroupPolicyState? state = null, CustomResourceOptions? options = null)
            : base("minio:index/iamGroupPolicy:IamGroupPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IamGroupPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IamGroupPolicy Get(string name, Input<string> id, IamGroupPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new IamGroupPolicy(name, id, state, options);
        }
    }

    public sealed class IamGroupPolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("group", required: true)]
        public Input<string> Group { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        [Input("policy", required: true)]
        public Input<string> Policy { get; set; } = null!;

        public IamGroupPolicyArgs()
        {
        }
        public static new IamGroupPolicyArgs Empty => new IamGroupPolicyArgs();
    }

    public sealed class IamGroupPolicyState : global::Pulumi.ResourceArgs
    {
        [Input("group")]
        public Input<string>? Group { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        [Input("policy")]
        public Input<string>? Policy { get; set; }

        public IamGroupPolicyState()
        {
        }
        public static new IamGroupPolicyState Empty => new IamGroupPolicyState();
    }
}
