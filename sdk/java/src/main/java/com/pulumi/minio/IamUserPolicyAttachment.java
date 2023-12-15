// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.minio;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.minio.IamUserPolicyAttachmentArgs;
import com.pulumi.minio.Utilities;
import com.pulumi.minio.inputs.IamUserPolicyAttachmentState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.minio.IamUser;
 * import com.pulumi.minio.IamPolicy;
 * import com.pulumi.minio.IamPolicyArgs;
 * import com.pulumi.minio.IamUserPolicyAttachment;
 * import com.pulumi.minio.IamUserPolicyAttachmentArgs;
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
 *         var testUser = new IamUser(&#34;testUser&#34;);
 * 
 *         var testPolicy = new IamPolicy(&#34;testPolicy&#34;, IamPolicyArgs.builder()        
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
 *         var developerIamUserPolicyAttachment = new IamUserPolicyAttachment(&#34;developerIamUserPolicyAttachment&#34;, IamUserPolicyAttachmentArgs.builder()        
 *             .userName(testUser.id())
 *             .policyName(testPolicy.id())
 *             .build());
 * 
 *         ctx.export(&#34;minioName&#34;, developerIamUserPolicyAttachment.id());
 *         ctx.export(&#34;minioUsers&#34;, developerIamUserPolicyAttachment.userName());
 *         ctx.export(&#34;minioGroup&#34;, developerIamUserPolicyAttachment.policyName());
 *         var developerIndex_iamUserPolicyAttachmentIamUserPolicyAttachment = new IamUserPolicyAttachment(&#34;developerIndex/iamUserPolicyAttachmentIamUserPolicyAttachment&#34;, IamUserPolicyAttachmentArgs.builder()        
 *             .userName(&#34;CN=My User,OU=Unit,DC=example,DC=com&#34;)
 *             .policyName(testPolicy.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="minio:index/iamUserPolicyAttachment:IamUserPolicyAttachment")
public class IamUserPolicyAttachment extends com.pulumi.resources.CustomResource {
    @Export(name="policyName", refs={String.class}, tree="[0]")
    private Output<String> policyName;

    public Output<String> policyName() {
        return this.policyName;
    }
    @Export(name="userName", refs={String.class}, tree="[0]")
    private Output<String> userName;

    public Output<String> userName() {
        return this.userName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IamUserPolicyAttachment(String name) {
        this(name, IamUserPolicyAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IamUserPolicyAttachment(String name, IamUserPolicyAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IamUserPolicyAttachment(String name, IamUserPolicyAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("minio:index/iamUserPolicyAttachment:IamUserPolicyAttachment", name, args == null ? IamUserPolicyAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IamUserPolicyAttachment(String name, Output<String> id, @Nullable IamUserPolicyAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("minio:index/iamUserPolicyAttachment:IamUserPolicyAttachment", name, state, makeResourceOptions(options, id));
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
    public static IamUserPolicyAttachment get(String name, Output<String> id, @Nullable IamUserPolicyAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IamUserPolicyAttachment(name, id, state, options);
    }
}
