// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.minio;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.minio.IamGroupPolicyAttachmentArgs;
import com.pulumi.minio.Utilities;
import com.pulumi.minio.inputs.IamGroupPolicyAttachmentState;
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
 * import com.pulumi.minio.IamGroup;
 * import com.pulumi.minio.IamGroupPolicy;
 * import com.pulumi.minio.IamGroupPolicyArgs;
 * import com.pulumi.minio.IamGroupPolicyAttachment;
 * import com.pulumi.minio.IamGroupPolicyAttachmentArgs;
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
 *         var developerIamGroup = new IamGroup(&#34;developerIamGroup&#34;);
 * 
 *         var testPolicy = new IamGroupPolicy(&#34;testPolicy&#34;, IamGroupPolicyArgs.builder()        
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
 * 
 *             &#34;&#34;&#34;)
 *             .build());
 * 
 *         var developerIamGroupPolicyAttachment = new IamGroupPolicyAttachment(&#34;developerIamGroupPolicyAttachment&#34;, IamGroupPolicyAttachmentArgs.builder()        
 *             .groupName(minio_iam_group.group().name())
 *             .policyName(minio_iam_policy.test_policy().id())
 *             .build());
 * 
 *         ctx.export(&#34;minioName&#34;, developerIamGroupPolicyAttachment.id());
 *         ctx.export(&#34;minioUsers&#34;, developerIamGroupPolicyAttachment.groupName());
 *         ctx.export(&#34;minioGroup&#34;, developerIamGroupPolicyAttachment.policyName());
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="minio:index/iamGroupPolicyAttachment:IamGroupPolicyAttachment")
public class IamGroupPolicyAttachment extends com.pulumi.resources.CustomResource {
    @Export(name="groupName", type=String.class, parameters={})
    private Output<String> groupName;

    public Output<String> groupName() {
        return this.groupName;
    }
    @Export(name="policyName", type=String.class, parameters={})
    private Output<String> policyName;

    public Output<String> policyName() {
        return this.policyName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IamGroupPolicyAttachment(String name) {
        this(name, IamGroupPolicyAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IamGroupPolicyAttachment(String name, IamGroupPolicyAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IamGroupPolicyAttachment(String name, IamGroupPolicyAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("minio:index/iamGroupPolicyAttachment:IamGroupPolicyAttachment", name, args == null ? IamGroupPolicyAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IamGroupPolicyAttachment(String name, Output<String> id, @Nullable IamGroupPolicyAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("minio:index/iamGroupPolicyAttachment:IamGroupPolicyAttachment", name, state, makeResourceOptions(options, id));
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
    public static IamGroupPolicyAttachment get(String name, Output<String> id, @Nullable IamGroupPolicyAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IamGroupPolicyAttachment(name, id, state, options);
    }
}