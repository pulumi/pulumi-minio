// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Minio
{
    [MinioResourceType("minio:index/iamGroup:IamGroup")]
    public partial class IamGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Disable group
        /// </summary>
        [Output("disableGroup")]
        public Output<bool?> DisableGroup { get; private set; } = null!;

        /// <summary>
        /// Delete group even if it has non-Terraform-managed members
        /// </summary>
        [Output("forceDestroy")]
        public Output<bool?> ForceDestroy { get; private set; } = null!;

        [Output("groupName")]
        public Output<string> GroupName { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a IamGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IamGroup(string name, IamGroupArgs? args = null, CustomResourceOptions? options = null)
            : base("minio:index/iamGroup:IamGroup", name, args ?? new IamGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IamGroup(string name, Input<string> id, IamGroupState? state = null, CustomResourceOptions? options = null)
            : base("minio:index/iamGroup:IamGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IamGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IamGroup Get(string name, Input<string> id, IamGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new IamGroup(name, id, state, options);
        }
    }

    public sealed class IamGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Disable group
        /// </summary>
        [Input("disableGroup")]
        public Input<bool>? DisableGroup { get; set; }

        /// <summary>
        /// Delete group even if it has non-Terraform-managed members
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public IamGroupArgs()
        {
        }
        public static new IamGroupArgs Empty => new IamGroupArgs();
    }

    public sealed class IamGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Disable group
        /// </summary>
        [Input("disableGroup")]
        public Input<bool>? DisableGroup { get; set; }

        /// <summary>
        /// Delete group even if it has non-Terraform-managed members
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public IamGroupState()
        {
        }
        public static new IamGroupState Empty => new IamGroupState();
    }
}
