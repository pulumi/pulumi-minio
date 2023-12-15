// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.minio;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.minio.IamGroupPolicyArgs;
import com.pulumi.minio.Utilities;
import com.pulumi.minio.inputs.IamGroupPolicyState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="minio:index/iamGroupPolicy:IamGroupPolicy")
public class IamGroupPolicy extends com.pulumi.resources.CustomResource {
    @Export(name="group", refs={String.class}, tree="[0]")
    private Output<String> group;

    public Output<String> group() {
        return this.group;
    }
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
    public IamGroupPolicy(String name) {
        this(name, IamGroupPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IamGroupPolicy(String name, IamGroupPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IamGroupPolicy(String name, IamGroupPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("minio:index/iamGroupPolicy:IamGroupPolicy", name, args == null ? IamGroupPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IamGroupPolicy(String name, Output<String> id, @Nullable IamGroupPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("minio:index/iamGroupPolicy:IamGroupPolicy", name, state, makeResourceOptions(options, id));
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
    public static IamGroupPolicy get(String name, Output<String> id, @Nullable IamGroupPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IamGroupPolicy(name, id, state, options);
    }
}
