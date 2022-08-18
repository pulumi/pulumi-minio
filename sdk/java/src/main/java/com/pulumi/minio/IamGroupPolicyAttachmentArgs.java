// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.minio;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class IamGroupPolicyAttachmentArgs extends com.pulumi.resources.ResourceArgs {

    public static final IamGroupPolicyAttachmentArgs Empty = new IamGroupPolicyAttachmentArgs();

    @Import(name="groupName", required=true)
    private Output<String> groupName;

    public Output<String> groupName() {
        return this.groupName;
    }

    @Import(name="policyName", required=true)
    private Output<String> policyName;

    public Output<String> policyName() {
        return this.policyName;
    }

    private IamGroupPolicyAttachmentArgs() {}

    private IamGroupPolicyAttachmentArgs(IamGroupPolicyAttachmentArgs $) {
        this.groupName = $.groupName;
        this.policyName = $.policyName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IamGroupPolicyAttachmentArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IamGroupPolicyAttachmentArgs $;

        public Builder() {
            $ = new IamGroupPolicyAttachmentArgs();
        }

        public Builder(IamGroupPolicyAttachmentArgs defaults) {
            $ = new IamGroupPolicyAttachmentArgs(Objects.requireNonNull(defaults));
        }

        public Builder groupName(Output<String> groupName) {
            $.groupName = groupName;
            return this;
        }

        public Builder groupName(String groupName) {
            return groupName(Output.of(groupName));
        }

        public Builder policyName(Output<String> policyName) {
            $.policyName = policyName;
            return this;
        }

        public Builder policyName(String policyName) {
            return policyName(Output.of(policyName));
        }

        public IamGroupPolicyAttachmentArgs build() {
            $.groupName = Objects.requireNonNull($.groupName, "expected parameter 'groupName' to be non-null");
            $.policyName = Objects.requireNonNull($.policyName, "expected parameter 'policyName' to be non-null");
            return $;
        }
    }

}
