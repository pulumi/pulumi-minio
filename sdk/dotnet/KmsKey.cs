// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Minio
{
    [MinioResourceType("minio:index/kmsKey:KmsKey")]
    public partial class KmsKey : global::Pulumi.CustomResource
    {
        [Output("keyId")]
        public Output<string> KeyId { get; private set; } = null!;


        /// <summary>
        /// Create a KmsKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KmsKey(string name, KmsKeyArgs args, CustomResourceOptions? options = null)
            : base("minio:index/kmsKey:KmsKey", name, args ?? new KmsKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KmsKey(string name, Input<string> id, KmsKeyState? state = null, CustomResourceOptions? options = null)
            : base("minio:index/kmsKey:KmsKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing KmsKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KmsKey Get(string name, Input<string> id, KmsKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new KmsKey(name, id, state, options);
        }
    }

    public sealed class KmsKeyArgs : global::Pulumi.ResourceArgs
    {
        [Input("keyId", required: true)]
        public Input<string> KeyId { get; set; } = null!;

        public KmsKeyArgs()
        {
        }
        public static new KmsKeyArgs Empty => new KmsKeyArgs();
    }

    public sealed class KmsKeyState : global::Pulumi.ResourceArgs
    {
        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        public KmsKeyState()
        {
        }
        public static new KmsKeyState Empty => new KmsKeyState();
    }
}