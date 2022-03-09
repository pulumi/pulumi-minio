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
    ///         var testIamUser = new Minio.IamUser("testIamUser", new Minio.IamUserArgs
    ///         {
    ///             ForceDestroy = true,
    ///             Tags = 
    ///             {
    ///                 { "tag-key", "tag-value" },
    ///             },
    ///         });
    ///         this.Test = testIamUser.Id;
    ///         this.Status = testIamUser.Status;
    ///         this.Secret = testIamUser.Secret;
    ///     }
    /// 
    ///     [Output("test")]
    ///     public Output&lt;string&gt; Test { get; set; }
    ///     [Output("status")]
    ///     public Output&lt;string&gt; Status { get; set; }
    ///     [Output("secret")]
    ///     public Output&lt;string&gt; Secret { get; set; }
    /// }
    /// ```
    /// </summary>
    [MinioResourceType("minio:index/iamUser:IamUser")]
    public partial class IamUser : Pulumi.CustomResource
    {
        /// <summary>
        /// Disable user
        /// </summary>
        [Output("disableUser")]
        public Output<bool?> DisableUser { get; private set; } = null!;

        /// <summary>
        /// Delete user even if it has non-Terraform-managed IAM access keys
        /// </summary>
        [Output("forceDestroy")]
        public Output<bool?> ForceDestroy { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("secret")]
        public Output<string> Secret { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Rotate Minio User Secret Key
        /// </summary>
        [Output("updateSecret")]
        public Output<bool?> UpdateSecret { get; private set; } = null!;


        /// <summary>
        /// Create a IamUser resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IamUser(string name, IamUserArgs? args = null, CustomResourceOptions? options = null)
            : base("minio:index/iamUser:IamUser", name, args ?? new IamUserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IamUser(string name, Input<string> id, IamUserState? state = null, CustomResourceOptions? options = null)
            : base("minio:index/iamUser:IamUser", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IamUser resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IamUser Get(string name, Input<string> id, IamUserState? state = null, CustomResourceOptions? options = null)
        {
            return new IamUser(name, id, state, options);
        }
    }

    public sealed class IamUserArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Disable user
        /// </summary>
        [Input("disableUser")]
        public Input<bool>? DisableUser { get; set; }

        /// <summary>
        /// Delete user even if it has non-Terraform-managed IAM access keys
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("secret")]
        public Input<string>? Secret { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Rotate Minio User Secret Key
        /// </summary>
        [Input("updateSecret")]
        public Input<bool>? UpdateSecret { get; set; }

        public IamUserArgs()
        {
        }
    }

    public sealed class IamUserState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Disable user
        /// </summary>
        [Input("disableUser")]
        public Input<bool>? DisableUser { get; set; }

        /// <summary>
        /// Delete user even if it has non-Terraform-managed IAM access keys
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("secret")]
        public Input<string>? Secret { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Rotate Minio User Secret Key
        /// </summary>
        [Input("updateSecret")]
        public Input<bool>? UpdateSecret { get; set; }

        public IamUserState()
        {
        }
    }
}
