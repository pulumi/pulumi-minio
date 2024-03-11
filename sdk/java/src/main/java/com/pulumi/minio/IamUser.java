// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.minio;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.minio.IamUserArgs;
import com.pulumi.minio.Utilities;
import com.pulumi.minio.inputs.IamUserState;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
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
 * import com.pulumi.minio.IamUser;
 * import com.pulumi.minio.IamUserArgs;
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
 *         var testIamUser = new IamUser(&#34;testIamUser&#34;, IamUserArgs.builder()        
 *             .forceDestroy(true)
 *             .tags(Map.of(&#34;tag-key&#34;, &#34;tag-value&#34;))
 *             .build());
 * 
 *         ctx.export(&#34;test&#34;, testIamUser.id());
 *         ctx.export(&#34;status&#34;, testIamUser.status());
 *         ctx.export(&#34;secret&#34;, testIamUser.secret());
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="minio:index/iamUser:IamUser")
public class IamUser extends com.pulumi.resources.CustomResource {
    /**
     * Disable user
     * 
     */
    @Export(name="disableUser", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disableUser;

    /**
     * @return Disable user
     * 
     */
    public Output<Optional<Boolean>> disableUser() {
        return Codegen.optional(this.disableUser);
    }
    /**
     * Delete user even if it has non-Terraform-managed IAM access keys
     * 
     */
    @Export(name="forceDestroy", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> forceDestroy;

    /**
     * @return Delete user even if it has non-Terraform-managed IAM access keys
     * 
     */
    public Output<Optional<Boolean>> forceDestroy() {
        return Codegen.optional(this.forceDestroy);
    }
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    public Output<String> name() {
        return this.name;
    }
    @Export(name="secret", refs={String.class}, tree="[0]")
    private Output<String> secret;

    public Output<String> secret() {
        return this.secret;
    }
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    public Output<String> status() {
        return this.status;
    }
    @Export(name="tags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> tags;

    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * Rotate Minio User Secret Key
     * 
     */
    @Export(name="updateSecret", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> updateSecret;

    /**
     * @return Rotate Minio User Secret Key
     * 
     */
    public Output<Optional<Boolean>> updateSecret() {
        return Codegen.optional(this.updateSecret);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IamUser(String name) {
        this(name, IamUserArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IamUser(String name, @Nullable IamUserArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IamUser(String name, @Nullable IamUserArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("minio:index/iamUser:IamUser", name, args == null ? IamUserArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IamUser(String name, Output<String> id, @Nullable IamUserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("minio:index/iamUser:IamUser", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "secret"
            ))
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
    public static IamUser get(String name, Output<String> id, @Nullable IamUserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IamUser(name, id, state, options);
    }
}
