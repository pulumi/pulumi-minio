// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.minio;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.minio.inputs.S3BucketReplicationRuleArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class S3BucketReplicationArgs extends com.pulumi.resources.ResourceArgs {

    public static final S3BucketReplicationArgs Empty = new S3BucketReplicationArgs();

    /**
     * The name of the existing target bucket to replicate into
     * 
     */
    @Import(name="bucket", required=true)
    private Output<String> bucket;

    /**
     * @return The name of the existing target bucket to replicate into
     * 
     */
    public Output<String> bucket() {
        return this.bucket;
    }

    /**
     * Rule definitions
     * 
     */
    @Import(name="rules")
    private @Nullable Output<List<S3BucketReplicationRuleArgs>> rules;

    /**
     * @return Rule definitions
     * 
     */
    public Optional<Output<List<S3BucketReplicationRuleArgs>>> rules() {
        return Optional.ofNullable(this.rules);
    }

    private S3BucketReplicationArgs() {}

    private S3BucketReplicationArgs(S3BucketReplicationArgs $) {
        this.bucket = $.bucket;
        this.rules = $.rules;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(S3BucketReplicationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private S3BucketReplicationArgs $;

        public Builder() {
            $ = new S3BucketReplicationArgs();
        }

        public Builder(S3BucketReplicationArgs defaults) {
            $ = new S3BucketReplicationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param bucket The name of the existing target bucket to replicate into
         * 
         * @return builder
         * 
         */
        public Builder bucket(Output<String> bucket) {
            $.bucket = bucket;
            return this;
        }

        /**
         * @param bucket The name of the existing target bucket to replicate into
         * 
         * @return builder
         * 
         */
        public Builder bucket(String bucket) {
            return bucket(Output.of(bucket));
        }

        /**
         * @param rules Rule definitions
         * 
         * @return builder
         * 
         */
        public Builder rules(@Nullable Output<List<S3BucketReplicationRuleArgs>> rules) {
            $.rules = rules;
            return this;
        }

        /**
         * @param rules Rule definitions
         * 
         * @return builder
         * 
         */
        public Builder rules(List<S3BucketReplicationRuleArgs> rules) {
            return rules(Output.of(rules));
        }

        /**
         * @param rules Rule definitions
         * 
         * @return builder
         * 
         */
        public Builder rules(S3BucketReplicationRuleArgs... rules) {
            return rules(List.of(rules));
        }

        public S3BucketReplicationArgs build() {
            if ($.bucket == null) {
                throw new MissingRequiredPropertyException("S3BucketReplicationArgs", "bucket");
            }
            return $;
        }
    }

}
