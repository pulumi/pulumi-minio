// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.minio;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.minio.IamGroupArgs;
import com.pulumi.minio.Utilities;
import com.pulumi.minio.inputs.IamGroupState;
import java.lang.Boolean;
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
 * import com.pulumi.minio.IamGroup;
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
 *         var developer = new IamGroup(&#34;developer&#34;);
 * 
 *         ctx.export(&#34;minioUserGroup&#34;, developer.groupName());
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="minio:index/iamGroup:IamGroup")
public class IamGroup extends com.pulumi.resources.CustomResource {
    /**
     * Disable group
     * 
     */
    @Export(name="disableGroup", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> disableGroup;

    /**
     * @return Disable group
     * 
     */
    public Output<Optional<Boolean>> disableGroup() {
        return Codegen.optional(this.disableGroup);
    }
    /**
     * Delete group even if it has non-Terraform-managed members
     * 
     */
    @Export(name="forceDestroy", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> forceDestroy;

    /**
     * @return Delete group even if it has non-Terraform-managed members
     * 
     */
    public Output<Optional<Boolean>> forceDestroy() {
        return Codegen.optional(this.forceDestroy);
    }
    @Export(name="groupName", type=String.class, parameters={})
    private Output<String> groupName;

    public Output<String> groupName() {
        return this.groupName;
    }
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    public Output<String> name() {
        return this.name;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IamGroup(String name) {
        this(name, IamGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IamGroup(String name, @Nullable IamGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IamGroup(String name, @Nullable IamGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("minio:index/iamGroup:IamGroup", name, args == null ? IamGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IamGroup(String name, Output<String> id, @Nullable IamGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("minio:index/iamGroup:IamGroup", name, state, makeResourceOptions(options, id));
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
    public static IamGroup get(String name, Output<String> id, @Nullable IamGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IamGroup(name, id, state, options);
    }
}
