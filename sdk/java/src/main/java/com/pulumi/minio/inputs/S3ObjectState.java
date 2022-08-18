// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.minio.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class S3ObjectState extends com.pulumi.resources.ResourceArgs {

    public static final S3ObjectState Empty = new S3ObjectState();

    @Import(name="bucketName")
    private @Nullable Output<String> bucketName;

    public Optional<Output<String>> bucketName() {
        return Optional.ofNullable(this.bucketName);
    }

    @Import(name="content")
    private @Nullable Output<String> content;

    public Optional<Output<String>> content() {
        return Optional.ofNullable(this.content);
    }

    @Import(name="contentBase64")
    private @Nullable Output<String> contentBase64;

    public Optional<Output<String>> contentBase64() {
        return Optional.ofNullable(this.contentBase64);
    }

    @Import(name="contentType")
    private @Nullable Output<String> contentType;

    public Optional<Output<String>> contentType() {
        return Optional.ofNullable(this.contentType);
    }

    @Import(name="etag")
    private @Nullable Output<String> etag;

    public Optional<Output<String>> etag() {
        return Optional.ofNullable(this.etag);
    }

    @Import(name="objectName")
    private @Nullable Output<String> objectName;

    public Optional<Output<String>> objectName() {
        return Optional.ofNullable(this.objectName);
    }

    @Import(name="source")
    private @Nullable Output<String> source;

    public Optional<Output<String>> source() {
        return Optional.ofNullable(this.source);
    }

    @Import(name="versionId")
    private @Nullable Output<String> versionId;

    public Optional<Output<String>> versionId() {
        return Optional.ofNullable(this.versionId);
    }

    private S3ObjectState() {}

    private S3ObjectState(S3ObjectState $) {
        this.bucketName = $.bucketName;
        this.content = $.content;
        this.contentBase64 = $.contentBase64;
        this.contentType = $.contentType;
        this.etag = $.etag;
        this.objectName = $.objectName;
        this.source = $.source;
        this.versionId = $.versionId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(S3ObjectState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private S3ObjectState $;

        public Builder() {
            $ = new S3ObjectState();
        }

        public Builder(S3ObjectState defaults) {
            $ = new S3ObjectState(Objects.requireNonNull(defaults));
        }

        public Builder bucketName(@Nullable Output<String> bucketName) {
            $.bucketName = bucketName;
            return this;
        }

        public Builder bucketName(String bucketName) {
            return bucketName(Output.of(bucketName));
        }

        public Builder content(@Nullable Output<String> content) {
            $.content = content;
            return this;
        }

        public Builder content(String content) {
            return content(Output.of(content));
        }

        public Builder contentBase64(@Nullable Output<String> contentBase64) {
            $.contentBase64 = contentBase64;
            return this;
        }

        public Builder contentBase64(String contentBase64) {
            return contentBase64(Output.of(contentBase64));
        }

        public Builder contentType(@Nullable Output<String> contentType) {
            $.contentType = contentType;
            return this;
        }

        public Builder contentType(String contentType) {
            return contentType(Output.of(contentType));
        }

        public Builder etag(@Nullable Output<String> etag) {
            $.etag = etag;
            return this;
        }

        public Builder etag(String etag) {
            return etag(Output.of(etag));
        }

        public Builder objectName(@Nullable Output<String> objectName) {
            $.objectName = objectName;
            return this;
        }

        public Builder objectName(String objectName) {
            return objectName(Output.of(objectName));
        }

        public Builder source(@Nullable Output<String> source) {
            $.source = source;
            return this;
        }

        public Builder source(String source) {
            return source(Output.of(source));
        }

        public Builder versionId(@Nullable Output<String> versionId) {
            $.versionId = versionId;
            return this;
        }

        public Builder versionId(String versionId) {
            return versionId(Output.of(versionId));
        }

        public S3ObjectState build() {
            return $;
        }
    }

}
