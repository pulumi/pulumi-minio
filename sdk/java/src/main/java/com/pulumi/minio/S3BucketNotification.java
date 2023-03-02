// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.minio;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.minio.S3BucketNotificationArgs;
import com.pulumi.minio.Utilities;
import com.pulumi.minio.inputs.S3BucketNotificationState;
import com.pulumi.minio.outputs.S3BucketNotificationQueue;
import java.lang.String;
import java.util.List;
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
 * import com.pulumi.minio.S3BucketNotification;
 * import com.pulumi.minio.S3BucketNotificationArgs;
 * import com.pulumi.minio.inputs.S3BucketNotificationQueueArgs;
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
 *         var bucketS3Bucket = new S3Bucket(&#34;bucketS3Bucket&#34;, S3BucketArgs.builder()        
 *             .bucket(&#34;example-bucket&#34;)
 *             .build());
 * 
 *         var bucketS3BucketNotification = new S3BucketNotification(&#34;bucketS3BucketNotification&#34;, S3BucketNotificationArgs.builder()        
 *             .bucket(minio_s3_bucket.state_terraform_s3().bucket())
 *             .queues(S3BucketNotificationQueueArgs.builder()
 *                 .id(&#34;notification-queue&#34;)
 *                 .queueArn(&#34;arn:minio:sqs::primary:webhook&#34;)
 *                 .events(                
 *                     &#34;s3:ObjectCreated:*&#34;,
 *                     &#34;s3:ObjectRemoved:Delete&#34;)
 *                 .filterPrefix(&#34;example/&#34;)
 *                 .filterSuffix(&#34;.png&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="minio:index/s3BucketNotification:S3BucketNotification")
public class S3BucketNotification extends com.pulumi.resources.CustomResource {
    @Export(name="bucket", type=String.class, parameters={})
    private Output<String> bucket;

    public Output<String> bucket() {
        return this.bucket;
    }
    @Export(name="queues", type=List.class, parameters={S3BucketNotificationQueue.class})
    private Output</* @Nullable */ List<S3BucketNotificationQueue>> queues;

    public Output<Optional<List<S3BucketNotificationQueue>>> queues() {
        return Codegen.optional(this.queues);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public S3BucketNotification(String name) {
        this(name, S3BucketNotificationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public S3BucketNotification(String name, S3BucketNotificationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public S3BucketNotification(String name, S3BucketNotificationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("minio:index/s3BucketNotification:S3BucketNotification", name, args == null ? S3BucketNotificationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private S3BucketNotification(String name, Output<String> id, @Nullable S3BucketNotificationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("minio:index/s3BucketNotification:S3BucketNotification", name, state, makeResourceOptions(options, id));
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
    public static S3BucketNotification get(String name, Output<String> id, @Nullable S3BucketNotificationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new S3BucketNotification(name, id, state, options);
    }
}