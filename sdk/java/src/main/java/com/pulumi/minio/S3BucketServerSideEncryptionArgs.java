// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.minio;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class S3BucketServerSideEncryptionArgs extends com.pulumi.resources.ResourceArgs {

    public static final S3BucketServerSideEncryptionArgs Empty = new S3BucketServerSideEncryptionArgs();

    @Import(name="bucket", required=true)
    private Output<String> bucket;

    public Output<String> bucket() {
        return this.bucket;
    }

    @Import(name="encryptionType", required=true)
    private Output<String> encryptionType;

    public Output<String> encryptionType() {
        return this.encryptionType;
    }

    @Import(name="kmsKeyId", required=true)
    private Output<String> kmsKeyId;

    public Output<String> kmsKeyId() {
        return this.kmsKeyId;
    }

    private S3BucketServerSideEncryptionArgs() {}

    private S3BucketServerSideEncryptionArgs(S3BucketServerSideEncryptionArgs $) {
        this.bucket = $.bucket;
        this.encryptionType = $.encryptionType;
        this.kmsKeyId = $.kmsKeyId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(S3BucketServerSideEncryptionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private S3BucketServerSideEncryptionArgs $;

        public Builder() {
            $ = new S3BucketServerSideEncryptionArgs();
        }

        public Builder(S3BucketServerSideEncryptionArgs defaults) {
            $ = new S3BucketServerSideEncryptionArgs(Objects.requireNonNull(defaults));
        }

        public Builder bucket(Output<String> bucket) {
            $.bucket = bucket;
            return this;
        }

        public Builder bucket(String bucket) {
            return bucket(Output.of(bucket));
        }

        public Builder encryptionType(Output<String> encryptionType) {
            $.encryptionType = encryptionType;
            return this;
        }

        public Builder encryptionType(String encryptionType) {
            return encryptionType(Output.of(encryptionType));
        }

        public Builder kmsKeyId(Output<String> kmsKeyId) {
            $.kmsKeyId = kmsKeyId;
            return this;
        }

        public Builder kmsKeyId(String kmsKeyId) {
            return kmsKeyId(Output.of(kmsKeyId));
        }

        public S3BucketServerSideEncryptionArgs build() {
            $.bucket = Objects.requireNonNull($.bucket, "expected parameter 'bucket' to be non-null");
            $.encryptionType = Objects.requireNonNull($.encryptionType, "expected parameter 'encryptionType' to be non-null");
            $.kmsKeyId = Objects.requireNonNull($.kmsKeyId, "expected parameter 'kmsKeyId' to be non-null");
            return $;
        }
    }

}
