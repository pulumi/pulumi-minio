// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Minio
{
    [MinioResourceType("minio:index/iamServiceAccount:IamServiceAccount")]
    public partial class IamServiceAccount : global::Pulumi.CustomResource
    {
        [Output("accessKey")]
        public Output<string> AccessKey { get; private set; } = null!;

        /// <summary>
        /// Disable service account
        /// </summary>
        [Output("disableUser")]
        public Output<bool?> DisableUser { get; private set; } = null!;

        /// <summary>
        /// policy of service account
        /// </summary>
        [Output("policy")]
        public Output<string?> Policy { get; private set; } = null!;

        [Output("secretKey")]
        public Output<string> SecretKey { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        [Output("targetUser")]
        public Output<string> TargetUser { get; private set; } = null!;

        /// <summary>
        /// rotate secret key
        /// </summary>
        [Output("updateSecret")]
        public Output<bool?> UpdateSecret { get; private set; } = null!;


        /// <summary>
        /// Create a IamServiceAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IamServiceAccount(string name, IamServiceAccountArgs args, CustomResourceOptions? options = null)
            : base("minio:index/iamServiceAccount:IamServiceAccount", name, args ?? new IamServiceAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IamServiceAccount(string name, Input<string> id, IamServiceAccountState? state = null, CustomResourceOptions? options = null)
            : base("minio:index/iamServiceAccount:IamServiceAccount", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "secretKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IamServiceAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IamServiceAccount Get(string name, Input<string> id, IamServiceAccountState? state = null, CustomResourceOptions? options = null)
        {
            return new IamServiceAccount(name, id, state, options);
        }
    }

    public sealed class IamServiceAccountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Disable service account
        /// </summary>
        [Input("disableUser")]
        public Input<bool>? DisableUser { get; set; }

        /// <summary>
        /// policy of service account
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        [Input("targetUser", required: true)]
        public Input<string> TargetUser { get; set; } = null!;

        /// <summary>
        /// rotate secret key
        /// </summary>
        [Input("updateSecret")]
        public Input<bool>? UpdateSecret { get; set; }

        public IamServiceAccountArgs()
        {
        }
        public static new IamServiceAccountArgs Empty => new IamServiceAccountArgs();
    }

    public sealed class IamServiceAccountState : global::Pulumi.ResourceArgs
    {
        [Input("accessKey")]
        public Input<string>? AccessKey { get; set; }

        /// <summary>
        /// Disable service account
        /// </summary>
        [Input("disableUser")]
        public Input<bool>? DisableUser { get; set; }

        /// <summary>
        /// policy of service account
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        [Input("secretKey")]
        private Input<string>? _secretKey;
        public Input<string>? SecretKey
        {
            get => _secretKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secretKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("targetUser")]
        public Input<string>? TargetUser { get; set; }

        /// <summary>
        /// rotate secret key
        /// </summary>
        [Input("updateSecret")]
        public Input<bool>? UpdateSecret { get; set; }

        public IamServiceAccountState()
        {
        }
        public static new IamServiceAccountState Empty => new IamServiceAccountState();
    }
}
