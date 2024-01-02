// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Minio.Inputs
{

    public sealed class S3BucketReplicationRuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Rule ARN genrated by MinIO
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Whether or not to synchronise marker deletion
        /// </summary>
        [Input("deleteMarkerReplication")]
        public Input<bool>? DeleteMarkerReplication { get; set; }

        /// <summary>
        /// Whether or not to propagate deletion
        /// </summary>
        [Input("deleteReplication")]
        public Input<bool>? DeleteReplication { get; set; }

        /// <summary>
        /// Whether or not this rule is enabled
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Whether or not to synchronise object created prior the replication configuration
        /// </summary>
        [Input("existingObjectReplication")]
        public Input<bool>? ExistingObjectReplication { get; set; }

        /// <summary>
        /// Rule ID generated by MinIO
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Whether or not to synchonise buckets and objects metadata (such as locks). This must be enabled to achieve a two-way replication
        /// </summary>
        [Input("metadataSync")]
        public Input<bool>? MetadataSync { get; set; }

        /// <summary>
        /// Bucket prefix object must be in to be syncronised
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        /// <summary>
        /// Rule priority. If omitted, the inverted index will be used as priority. This means that the first rule definition will have the higher priority
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tags which objects must have to be syncronised
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Bucket prefix
        /// </summary>
        [Input("target", required: true)]
        public Input<Inputs.S3BucketReplicationRuleTargetGetArgs> Target { get; set; } = null!;

        public S3BucketReplicationRuleGetArgs()
        {
        }
        public static new S3BucketReplicationRuleGetArgs Empty => new S3BucketReplicationRuleGetArgs();
    }
}
