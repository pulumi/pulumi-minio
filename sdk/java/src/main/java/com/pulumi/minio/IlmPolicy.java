// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.minio;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.minio.IlmPolicyArgs;
import com.pulumi.minio.Utilities;
import com.pulumi.minio.inputs.IlmPolicyState;
import com.pulumi.minio.outputs.IlmPolicyRule;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * `minio.IlmPolicy` handles lifecycle settings for a given `minio.S3Bucket`.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.minio.S3Bucket;
 * import com.pulumi.minio.S3BucketArgs;
 * import com.pulumi.minio.IlmPolicy;
 * import com.pulumi.minio.IlmPolicyArgs;
 * import com.pulumi.minio.inputs.IlmPolicyRuleArgs;
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
 *         var bucket = new S3Bucket(&#34;bucket&#34;, S3BucketArgs.builder()        
 *             .bucket(&#34;bucket&#34;)
 *             .build());
 * 
 *         var bucket_lifecycle_rules = new IlmPolicy(&#34;bucket-lifecycle-rules&#34;, IlmPolicyArgs.builder()        
 *             .bucket(bucket.bucket())
 *             .rules(IlmPolicyRuleArgs.builder()
 *                 .id(&#34;expire-7d&#34;)
 *                 .expiration(&#34;7d&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="minio:index/ilmPolicy:IlmPolicy")
public class IlmPolicy extends com.pulumi.resources.CustomResource {
    @Export(name="bucket", type=String.class, parameters={})
    private Output<String> bucket;

    public Output<String> bucket() {
        return this.bucket;
    }
    @Export(name="rules", type=List.class, parameters={IlmPolicyRule.class})
    private Output<List<IlmPolicyRule>> rules;

    public Output<List<IlmPolicyRule>> rules() {
        return this.rules;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IlmPolicy(String name) {
        this(name, IlmPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IlmPolicy(String name, IlmPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IlmPolicy(String name, IlmPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("minio:index/ilmPolicy:IlmPolicy", name, args == null ? IlmPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IlmPolicy(String name, Output<String> id, @Nullable IlmPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("minio:index/ilmPolicy:IlmPolicy", name, state, makeResourceOptions(options, id));
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
    public static IlmPolicy get(String name, Output<String> id, @Nullable IlmPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IlmPolicy(name, id, state, options);
    }
}
