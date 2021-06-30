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
    /// using Pulumi;
    /// using Minio = Pulumi.Minio;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var developerIamGroup = new Minio.IamGroup("developerIamGroup", new Minio.IamGroupArgs
    ///         {
    ///         });
    ///         var userOne = new Minio.IamUser("userOne", new Minio.IamUserArgs
    ///         {
    ///         });
    ///         var developerIamGroupUserAttachment = new Minio.IamGroupUserAttachment("developerIamGroupUserAttachment", new Minio.IamGroupUserAttachmentArgs
    ///         {
    ///             GroupName = minio_iam_group.Group.Name,
    ///             UserName = userOne.Name,
    ///         });
    ///         this.MinioName = developerIamGroupUserAttachment.Id;
    ///         this.MinioUsers = developerIamGroupUserAttachment.GroupName;
    ///         this.MinioGroup = developerIamGroupUserAttachment.UserName;
    ///     }
    /// 
    ///     [Output("minioName")]
    ///     public Output&lt;string&gt; MinioName { get; set; }
    ///     [Output("minioUsers")]
    ///     public Output&lt;string&gt; MinioUsers { get; set; }
    ///     [Output("minioGroup")]
    ///     public Output&lt;string&gt; MinioGroup { get; set; }
    /// }
    /// ```
    /// </summary>
    [MinioResourceType("minio:index/iamGroupUserAttachment:IamGroupUserAttachment")]
    public partial class IamGroupUserAttachment : Pulumi.CustomResource
    {
        [Output("groupName")]
        public Output<string> GroupName { get; private set; } = null!;

        [Output("userName")]
        public Output<string> UserName { get; private set; } = null!;


        /// <summary>
        /// Create a IamGroupUserAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IamGroupUserAttachment(string name, IamGroupUserAttachmentArgs args, CustomResourceOptions? options = null)
            : base("minio:index/iamGroupUserAttachment:IamGroupUserAttachment", name, args ?? new IamGroupUserAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IamGroupUserAttachment(string name, Input<string> id, IamGroupUserAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("minio:index/iamGroupUserAttachment:IamGroupUserAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IamGroupUserAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IamGroupUserAttachment Get(string name, Input<string> id, IamGroupUserAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new IamGroupUserAttachment(name, id, state, options);
        }
    }

    public sealed class IamGroupUserAttachmentArgs : Pulumi.ResourceArgs
    {
        [Input("groupName", required: true)]
        public Input<string> GroupName { get; set; } = null!;

        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        public IamGroupUserAttachmentArgs()
        {
        }
    }

    public sealed class IamGroupUserAttachmentState : Pulumi.ResourceArgs
    {
        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public IamGroupUserAttachmentState()
        {
        }
    }
}
