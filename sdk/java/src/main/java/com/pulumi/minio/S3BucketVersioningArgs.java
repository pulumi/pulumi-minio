// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.minio;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.minio.inputs.S3BucketVersioningVersioningConfigurationArgs;
import java.lang.String;
import java.util.Objects;


public final class S3BucketVersioningArgs extends com.pulumi.resources.ResourceArgs {

    public static final S3BucketVersioningArgs Empty = new S3BucketVersioningArgs();

    @Import(name="bucket", required=true)
    private Output<String> bucket;

    public Output<String> bucket() {
        return this.bucket;
    }

    @Import(name="versioningConfiguration", required=true)
    private Output<S3BucketVersioningVersioningConfigurationArgs> versioningConfiguration;

    public Output<S3BucketVersioningVersioningConfigurationArgs> versioningConfiguration() {
        return this.versioningConfiguration;
    }

    private S3BucketVersioningArgs() {}

    private S3BucketVersioningArgs(S3BucketVersioningArgs $) {
        this.bucket = $.bucket;
        this.versioningConfiguration = $.versioningConfiguration;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(S3BucketVersioningArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private S3BucketVersioningArgs $;

        public Builder() {
            $ = new S3BucketVersioningArgs();
        }

        public Builder(S3BucketVersioningArgs defaults) {
            $ = new S3BucketVersioningArgs(Objects.requireNonNull(defaults));
        }

        public Builder bucket(Output<String> bucket) {
            $.bucket = bucket;
            return this;
        }

        public Builder bucket(String bucket) {
            return bucket(Output.of(bucket));
        }

        public Builder versioningConfiguration(Output<S3BucketVersioningVersioningConfigurationArgs> versioningConfiguration) {
            $.versioningConfiguration = versioningConfiguration;
            return this;
        }

        public Builder versioningConfiguration(S3BucketVersioningVersioningConfigurationArgs versioningConfiguration) {
            return versioningConfiguration(Output.of(versioningConfiguration));
        }

        public S3BucketVersioningArgs build() {
            $.bucket = Objects.requireNonNull($.bucket, "expected parameter 'bucket' to be non-null");
            $.versioningConfiguration = Objects.requireNonNull($.versioningConfiguration, "expected parameter 'versioningConfiguration' to be non-null");
            return $;
        }
    }

}
