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
 * <pre>
 * {@code
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
 *         var test = new IamUser("test", IamUserArgs.builder()
 *             .name("test")
 *             .forceDestroy(true)
 *             .tags(Map.of("tag-key", "tag-value"))
 *             .build());
 * 
 *         ctx.export("test", test.id());
 *         ctx.export("status", test.status());
 *         ctx.export("secret", test.secret());
 *     }
 * }
 * }
 * </pre>
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
    @Export(name="forceDestroy", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> forceDestroy;

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
    public IamUser(java.lang.String name) {
        this(name, IamUserArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IamUser(java.lang.String name, @Nullable IamUserArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IamUser(java.lang.String name, @Nullable IamUserArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("minio:index/iamUser:IamUser", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private IamUser(java.lang.String name, Output<java.lang.String> id, @Nullable IamUserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("minio:index/iamUser:IamUser", name, state, makeResourceOptions(options, id), false);
    }

    private static IamUserArgs makeArgs(@Nullable IamUserArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? IamUserArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static IamUser get(java.lang.String name, Output<java.lang.String> id, @Nullable IamUserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IamUser(name, id, state, options);
    }
}
