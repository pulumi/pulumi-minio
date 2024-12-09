// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Minio
{
    [MinioResourceType("minio:index/iamGroupMembership:IamGroupMembership")]
    public partial class IamGroupMembership : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Group name to add users
        /// </summary>
        [Output("group")]
        public Output<string> Group { get; private set; } = null!;

        /// <summary>
        /// Name of group membership
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Add user or list of users such as a group membership
        /// </summary>
        [Output("users")]
        public Output<ImmutableArray<string>> Users { get; private set; } = null!;


        /// <summary>
        /// Create a IamGroupMembership resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IamGroupMembership(string name, IamGroupMembershipArgs args, CustomResourceOptions? options = null)
            : base("minio:index/iamGroupMembership:IamGroupMembership", name, args ?? new IamGroupMembershipArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IamGroupMembership(string name, Input<string> id, IamGroupMembershipState? state = null, CustomResourceOptions? options = null)
            : base("minio:index/iamGroupMembership:IamGroupMembership", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IamGroupMembership resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IamGroupMembership Get(string name, Input<string> id, IamGroupMembershipState? state = null, CustomResourceOptions? options = null)
        {
            return new IamGroupMembership(name, id, state, options);
        }
    }

    public sealed class IamGroupMembershipArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Group name to add users
        /// </summary>
        [Input("group", required: true)]
        public Input<string> Group { get; set; } = null!;

        /// <summary>
        /// Name of group membership
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("users", required: true)]
        private InputList<string>? _users;

        /// <summary>
        /// Add user or list of users such as a group membership
        /// </summary>
        public InputList<string> Users
        {
            get => _users ?? (_users = new InputList<string>());
            set => _users = value;
        }

        public IamGroupMembershipArgs()
        {
        }
        public static new IamGroupMembershipArgs Empty => new IamGroupMembershipArgs();
    }

    public sealed class IamGroupMembershipState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Group name to add users
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        /// <summary>
        /// Name of group membership
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("users")]
        private InputList<string>? _users;

        /// <summary>
        /// Add user or list of users such as a group membership
        /// </summary>
        public InputList<string> Users
        {
            get => _users ?? (_users = new InputList<string>());
            set => _users = value;
        }

        public IamGroupMembershipState()
        {
        }
        public static new IamGroupMembershipState Empty => new IamGroupMembershipState();
    }
}