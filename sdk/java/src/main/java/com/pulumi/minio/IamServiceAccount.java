// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.minio;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.minio.IamServiceAccountArgs;
import com.pulumi.minio.Utilities;
import com.pulumi.minio.inputs.IamServiceAccountState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
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
 * import com.pulumi.minio.IamServiceAccount;
 * import com.pulumi.minio.IamServiceAccountArgs;
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
 *         var testServiceAccount = new IamServiceAccount("testServiceAccount", IamServiceAccountArgs.builder()
 *             .targetUser(test.name())
 *             .build());
 * 
 *         ctx.export("minioUser", testServiceAccount.accessKey());
 *         ctx.export("minioPassword", testServiceAccount.secretKey());
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="minio:index/iamServiceAccount:IamServiceAccount")
public class IamServiceAccount extends com.pulumi.resources.CustomResource {
    @Export(name="accessKey", refs={String.class}, tree="[0]")
    private Output<String> accessKey;

    public Output<String> accessKey() {
        return this.accessKey;
    }
    /**
     * Disable service account
     * 
     */
    @Export(name="disableUser", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disableUser;

    /**
     * @return Disable service account
     * 
     */
    public Output<Optional<Boolean>> disableUser() {
        return Codegen.optional(this.disableUser);
    }
    /**
     * policy of service account
     * 
     */
    @Export(name="policy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> policy;

    /**
     * @return policy of service account
     * 
     */
    public Output<Optional<String>> policy() {
        return Codegen.optional(this.policy);
    }
    @Export(name="secretKey", refs={String.class}, tree="[0]")
    private Output<String> secretKey;

    public Output<String> secretKey() {
        return this.secretKey;
    }
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    public Output<String> status() {
        return this.status;
    }
    @Export(name="targetUser", refs={String.class}, tree="[0]")
    private Output<String> targetUser;

    public Output<String> targetUser() {
        return this.targetUser;
    }
    /**
     * rotate secret key
     * 
     */
    @Export(name="updateSecret", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> updateSecret;

    /**
     * @return rotate secret key
     * 
     */
    public Output<Optional<Boolean>> updateSecret() {
        return Codegen.optional(this.updateSecret);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IamServiceAccount(java.lang.String name) {
        this(name, IamServiceAccountArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IamServiceAccount(java.lang.String name, IamServiceAccountArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IamServiceAccount(java.lang.String name, IamServiceAccountArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("minio:index/iamServiceAccount:IamServiceAccount", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private IamServiceAccount(java.lang.String name, Output<java.lang.String> id, @Nullable IamServiceAccountState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("minio:index/iamServiceAccount:IamServiceAccount", name, state, makeResourceOptions(options, id), false);
    }

    private static IamServiceAccountArgs makeArgs(IamServiceAccountArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? IamServiceAccountArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "secretKey"
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
    public static IamServiceAccount get(java.lang.String name, Output<java.lang.String> id, @Nullable IamServiceAccountState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IamServiceAccount(name, id, state, options);
    }
}
