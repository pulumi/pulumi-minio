// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.minio;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.minio.S3BucketArgs;
import com.pulumi.minio.Utilities;
import com.pulumi.minio.inputs.S3BucketState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.minio.S3Bucket;
 * import com.pulumi.minio.S3BucketArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var stateTerraformS3 = new S3Bucket(&#34;stateTerraformS3&#34;, S3BucketArgs.builder()        
 *             .acl(&#34;public&#34;)
 *             .bucket(&#34;state-terraform-s3&#34;)
 *             .build());
 * 
 *         ctx.export(&#34;minioId&#34;, stateTerraformS3.id());
 *         ctx.export(&#34;minioUrl&#34;, stateTerraformS3.bucketDomainName());
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="minio:index/s3Bucket:S3Bucket")
public class S3Bucket extends com.pulumi.resources.CustomResource {
    @Export(name="acl", type=String.class, parameters={})
    private Output</* @Nullable */ String> acl;

    public Output<Optional<String>> acl() {
        return Codegen.optional(this.acl);
    }
    @Export(name="arn", type=String.class, parameters={})
    private Output<String> arn;

    public Output<String> arn() {
        return this.arn;
    }
    @Export(name="bucket", type=String.class, parameters={})
    private Output<String> bucket;

    public Output<String> bucket() {
        return this.bucket;
    }
    @Export(name="bucketDomainName", type=String.class, parameters={})
    private Output<String> bucketDomainName;

    public Output<String> bucketDomainName() {
        return this.bucketDomainName;
    }
    @Export(name="bucketPrefix", type=String.class, parameters={})
    private Output</* @Nullable */ String> bucketPrefix;

    public Output<Optional<String>> bucketPrefix() {
        return Codegen.optional(this.bucketPrefix);
    }
    @Export(name="forceDestroy", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> forceDestroy;

    public Output<Optional<Boolean>> forceDestroy() {
        return Codegen.optional(this.forceDestroy);
    }
    /**
     * The limit of the amount of data in the bucket (bytes).
     * 
     */
    @Export(name="quota", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> quota;

    /**
     * @return The limit of the amount of data in the bucket (bytes).
     * 
     */
    public Output<Optional<Integer>> quota() {
        return Codegen.optional(this.quota);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public S3Bucket(String name) {
        this(name, S3BucketArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public S3Bucket(String name, @Nullable S3BucketArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public S3Bucket(String name, @Nullable S3BucketArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("minio:index/s3Bucket:S3Bucket", name, args == null ? S3BucketArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private S3Bucket(String name, Output<String> id, @Nullable S3BucketState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("minio:index/s3Bucket:S3Bucket", name, state, makeResourceOptions(options, id));
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
    public static S3Bucket get(String name, Output<String> id, @Nullable S3BucketState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new S3Bucket(name, id, state, options);
    }
}
