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
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Minio = Pulumi.Minio;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var testUser = new Minio.IamUser("testUser");
    /// 
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
    ///     var developerIamUserPolicyAttachment = new Minio.IamUserPolicyAttachment("developerIamUserPolicyAttachment", new()
    ///     {
    ///         UserName = testUser.Id,
    ///         PolicyName = testPolicy.Id,
    ///     });
    /// 
    ///     var developerIndex_iamUserPolicyAttachmentIamUserPolicyAttachment = new Minio.IamUserPolicyAttachment("developerIndex/iamUserPolicyAttachmentIamUserPolicyAttachment", new()
    ///     {
    ///         UserName = "CN=My User,OU=Unit,DC=example,DC=com",
    ///         PolicyName = testPolicy.Id,
    ///     });
    /// 
    ///     return new Dictionary&lt;string, object?&gt;
    ///     {
    ///         ["minioName"] = developerIamUserPolicyAttachment.Id,
    ///         ["minioUsers"] = developerIamUserPolicyAttachment.UserName,
    ///         ["minioGroup"] = developerIamUserPolicyAttachment.PolicyName,
    ///     };
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [MinioResourceType("minio:index/iamUserPolicyAttachment:IamUserPolicyAttachment")]
    public partial class IamUserPolicyAttachment : global::Pulumi.CustomResource
    {
        [Output("policyName")]
        public Output<string> PolicyName { get; private set; } = null!;

        [Output("userName")]
        public Output<string> UserName { get; private set; } = null!;


        /// <summary>
        /// Create a IamUserPolicyAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IamUserPolicyAttachment(string name, IamUserPolicyAttachmentArgs args, CustomResourceOptions? options = null)
            : base("minio:index/iamUserPolicyAttachment:IamUserPolicyAttachment", name, args ?? new IamUserPolicyAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IamUserPolicyAttachment(string name, Input<string> id, IamUserPolicyAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("minio:index/iamUserPolicyAttachment:IamUserPolicyAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IamUserPolicyAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IamUserPolicyAttachment Get(string name, Input<string> id, IamUserPolicyAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new IamUserPolicyAttachment(name, id, state, options);
        }
    }

    public sealed class IamUserPolicyAttachmentArgs : global::Pulumi.ResourceArgs
    {
        [Input("policyName", required: true)]
        public Input<string> PolicyName { get; set; } = null!;

        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        public IamUserPolicyAttachmentArgs()
        {
        }
        public static new IamUserPolicyAttachmentArgs Empty => new IamUserPolicyAttachmentArgs();
    }

    public sealed class IamUserPolicyAttachmentState : global::Pulumi.ResourceArgs
    {
        [Input("policyName")]
        public Input<string>? PolicyName { get; set; }

        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public IamUserPolicyAttachmentState()
        {
        }
        public static new IamUserPolicyAttachmentState Empty => new IamUserPolicyAttachmentState();
    }
}
