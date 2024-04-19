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
 * 
 */
@ResourceType(type="minio:index/iamGroupPolicyAttachment:IamGroupPolicyAttachment")
public class IamGroupPolicyAttachment extends com.pulumi.resources.CustomResource {
    @Export(name="groupName", refs={String.class}, tree="[0]")
    private Output<String> groupName;

    public Output<String> groupName() {
        return this.groupName;
    }
    @Export(name="policyName", refs={String.class}, tree="[0]")
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
