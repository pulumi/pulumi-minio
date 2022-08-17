// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.minio.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class S3BucketState extends com.pulumi.resources.ResourceArgs {

    public static final S3BucketState Empty = new S3BucketState();

    @Import(name="acl")
    private @Nullable Output<String> acl;

    public Optional<Output<String>> acl() {
        return Optional.ofNullable(this.acl);
    }

    @Import(name="bucket")
    private @Nullable Output<String> bucket;

    public Optional<Output<String>> bucket() {
        return Optional.ofNullable(this.bucket);
    }

    @Import(name="bucketDomainName")
    private @Nullable Output<String> bucketDomainName;

    public Optional<Output<String>> bucketDomainName() {
        return Optional.ofNullable(this.bucketDomainName);
    }

    @Import(name="bucketPrefix")
    private @Nullable Output<String> bucketPrefix;

    public Optional<Output<String>> bucketPrefix() {
        return Optional.ofNullable(this.bucketPrefix);
    }

    @Import(name="forceDestroy")
    private @Nullable Output<Boolean> forceDestroy;

    public Optional<Output<Boolean>> forceDestroy() {
        return Optional.ofNullable(this.forceDestroy);
    }

    /**
     * The limit of the amount of data in the bucket (bytes).
     * 
     */
    @Import(name="quota")
    private @Nullable Output<Integer> quota;

    /**
     * @return The limit of the amount of data in the bucket (bytes).
     * 
     */
    public Optional<Output<Integer>> quota() {
        return Optional.ofNullable(this.quota);
    }

    private S3BucketState() {}

    private S3BucketState(S3BucketState $) {
        this.acl = $.acl;
        this.bucket = $.bucket;
        this.bucketDomainName = $.bucketDomainName;
        this.bucketPrefix = $.bucketPrefix;
        this.forceDestroy = $.forceDestroy;
        this.quota = $.quota;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(S3BucketState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private S3BucketState $;

        public Builder() {
            $ = new S3BucketState();
        }

        public Builder(S3BucketState defaults) {
            $ = new S3BucketState(Objects.requireNonNull(defaults));
        }

        public Builder acl(@Nullable Output<String> acl) {
            $.acl = acl;
            return this;
        }

        public Builder acl(String acl) {
            return acl(Output.of(acl));
        }

        public Builder bucket(@Nullable Output<String> bucket) {
            $.bucket = bucket;
            return this;
        }

        public Builder bucket(String bucket) {
            return bucket(Output.of(bucket));
        }

        public Builder bucketDomainName(@Nullable Output<String> bucketDomainName) {
            $.bucketDomainName = bucketDomainName;
            return this;
        }

        public Builder bucketDomainName(String bucketDomainName) {
            return bucketDomainName(Output.of(bucketDomainName));
        }

        public Builder bucketPrefix(@Nullable Output<String> bucketPrefix) {
            $.bucketPrefix = bucketPrefix;
            return this;
        }

        public Builder bucketPrefix(String bucketPrefix) {
            return bucketPrefix(Output.of(bucketPrefix));
        }

        public Builder forceDestroy(@Nullable Output<Boolean> forceDestroy) {
            $.forceDestroy = forceDestroy;
            return this;
        }

        public Builder forceDestroy(Boolean forceDestroy) {
            return forceDestroy(Output.of(forceDestroy));
        }

        /**
         * @param quota The limit of the amount of data in the bucket (bytes).
         * 
         * @return builder
         * 
         */
        public Builder quota(@Nullable Output<Integer> quota) {
            $.quota = quota;
            return this;
        }

        /**
         * @param quota The limit of the amount of data in the bucket (bytes).
         * 
         * @return builder
         * 
         */
        public Builder quota(Integer quota) {
            return quota(Output.of(quota));
        }

        public S3BucketState build() {
            return $;
        }
    }

}
