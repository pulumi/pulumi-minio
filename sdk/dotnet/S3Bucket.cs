// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Minio
{
    [MinioResourceType("minio:index/s3Bucket:S3Bucket")]
    public partial class S3Bucket : Pulumi.CustomResource
    {
        [Output("acl")]
        public Output<string> Acl { get; private set; } = null!;

        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        [Output("bucketDomainName")]
        public Output<string> BucketDomainName { get; private set; } = null!;

        [Output("bucketPrefix")]
        public Output<string?> BucketPrefix { get; private set; } = null!;

        [Output("forceDestroy")]
        public Output<bool?> ForceDestroy { get; private set; } = null!;


        /// <summary>
        /// Create a S3Bucket resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public S3Bucket(string name, S3BucketArgs args, CustomResourceOptions? options = null)
            : base("minio:index/s3Bucket:S3Bucket", name, args ?? new S3BucketArgs(), MakeResourceOptions(options, ""))
        {
        }

        private S3Bucket(string name, Input<string> id, S3BucketState? state = null, CustomResourceOptions? options = null)
            : base("minio:index/s3Bucket:S3Bucket", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing S3Bucket resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static S3Bucket Get(string name, Input<string> id, S3BucketState? state = null, CustomResourceOptions? options = null)
        {
            return new S3Bucket(name, id, state, options);
        }
    }

    public sealed class S3BucketArgs : Pulumi.ResourceArgs
    {
        [Input("acl", required: true)]
        public Input<string> Acl { get; set; } = null!;

        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        [Input("bucketPrefix")]
        public Input<string>? BucketPrefix { get; set; }

        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        public S3BucketArgs()
        {
        }
    }

    public sealed class S3BucketState : Pulumi.ResourceArgs
    {
        [Input("acl")]
        public Input<string>? Acl { get; set; }

        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        [Input("bucketDomainName")]
        public Input<string>? BucketDomainName { get; set; }

        [Input("bucketPrefix")]
        public Input<string>? BucketPrefix { get; set; }

        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        public S3BucketState()
        {
        }
    }
}
