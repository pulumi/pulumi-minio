// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.minio;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.minio.S3ObjectArgs;
import com.pulumi.minio.Utilities;
import com.pulumi.minio.inputs.S3ObjectState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="minio:index/s3Object:S3Object")
public class S3Object extends com.pulumi.resources.CustomResource {
    @Export(name="bucketName", refs={String.class}, tree="[0]")
    private Output<String> bucketName;

    public Output<String> bucketName() {
        return this.bucketName;
    }
    @Export(name="content", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> content;

    public Output<Optional<String>> content() {
        return Codegen.optional(this.content);
    }
    @Export(name="contentBase64", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> contentBase64;

    public Output<Optional<String>> contentBase64() {
        return Codegen.optional(this.contentBase64);
    }
    @Export(name="contentType", refs={String.class}, tree="[0]")
    private Output<String> contentType;

    public Output<String> contentType() {
        return this.contentType;
    }
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    public Output<String> etag() {
        return this.etag;
    }
    @Export(name="objectName", refs={String.class}, tree="[0]")
    private Output<String> objectName;

    public Output<String> objectName() {
        return this.objectName;
    }
    @Export(name="source", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> source;

    public Output<Optional<String>> source() {
        return Codegen.optional(this.source);
    }
    @Export(name="versionId", refs={String.class}, tree="[0]")
    private Output<String> versionId;

    public Output<String> versionId() {
        return this.versionId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public S3Object(String name) {
        this(name, S3ObjectArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public S3Object(String name, S3ObjectArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public S3Object(String name, S3ObjectArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("minio:index/s3Object:S3Object", name, args == null ? S3ObjectArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private S3Object(String name, Output<String> id, @Nullable S3ObjectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("minio:index/s3Object:S3Object", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static S3Object get(String name, Output<String> id, @Nullable S3ObjectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new S3Object(name, id, state, options);
    }
}
