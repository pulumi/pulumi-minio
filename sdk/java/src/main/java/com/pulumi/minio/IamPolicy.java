// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.minio;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.minio.IamPolicyArgs;
import com.pulumi.minio.Utilities;
import com.pulumi.minio.inputs.IamPolicyState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.minio.IamPolicy;
 * import com.pulumi.minio.IamPolicyArgs;
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
 *         var testPolicy = new IamPolicy(&#34;testPolicy&#34;, IamPolicyArgs.builder()        
 *             .name(&#34;state-terraform-s3&#34;)
 *             .policy(&#34;&#34;&#34;
 * {
 *   &#34;Version&#34;:&#34;2012-10-17&#34;,
 *   &#34;Statement&#34;: [
 *     {
 *       &#34;Sid&#34;:&#34;ListAllBucket&#34;,
 *       &#34;Effect&#34;: &#34;Allow&#34;,
 *       &#34;Action&#34;: [&#34;s3:PutObject&#34;],
 *       &#34;Principal&#34;:&#34;*&#34;,
 *       &#34;Resource&#34;: &#34;arn:aws:s3:::state-terraform-s3/*&#34;
 *     }
 *   ]
 * }
 *             &#34;&#34;&#34;)
 *             .build());
 * 
 *         ctx.export(&#34;minioId&#34;, testPolicy.id());
 *         ctx.export(&#34;minioPolicy&#34;, testPolicy.policy());
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="minio:index/iamPolicy:IamPolicy")
public class IamPolicy extends com.pulumi.resources.CustomResource {
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    public Output<String> name() {
        return this.name;
    }
    @Export(name="namePrefix", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> namePrefix;

    public Output<Optional<String>> namePrefix() {
        return Codegen.optional(this.namePrefix);
    }
    @Export(name="policy", refs={String.class}, tree="[0]")
    private Output<String> policy;

    public Output<String> policy() {
        return this.policy;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IamPolicy(String name) {
        this(name, IamPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IamPolicy(String name, IamPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IamPolicy(String name, IamPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("minio:index/iamPolicy:IamPolicy", name, args == null ? IamPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IamPolicy(String name, Output<String> id, @Nullable IamPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("minio:index/iamPolicy:IamPolicy", name, state, makeResourceOptions(options, id));
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
    public static IamPolicy get(String name, Output<String> id, @Nullable IamPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IamPolicy(name, id, state, options);
    }
}
