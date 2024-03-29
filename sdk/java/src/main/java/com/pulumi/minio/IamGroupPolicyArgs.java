// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.minio;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IamGroupPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final IamGroupPolicyArgs Empty = new IamGroupPolicyArgs();

    @Import(name="group", required=true)
    private Output<String> group;

    public Output<String> group() {
        return this.group;
    }

    @Import(name="name")
    private @Nullable Output<String> name;

    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    @Import(name="namePrefix")
    private @Nullable Output<String> namePrefix;

    public Optional<Output<String>> namePrefix() {
        return Optional.ofNullable(this.namePrefix);
    }

    @Import(name="policy", required=true)
    private Output<String> policy;

    public Output<String> policy() {
        return this.policy;
    }

    private IamGroupPolicyArgs() {}

    private IamGroupPolicyArgs(IamGroupPolicyArgs $) {
        this.group = $.group;
        this.name = $.name;
        this.namePrefix = $.namePrefix;
        this.policy = $.policy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IamGroupPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IamGroupPolicyArgs $;

        public Builder() {
            $ = new IamGroupPolicyArgs();
        }

        public Builder(IamGroupPolicyArgs defaults) {
            $ = new IamGroupPolicyArgs(Objects.requireNonNull(defaults));
        }

        public Builder group(Output<String> group) {
            $.group = group;
            return this;
        }

        public Builder group(String group) {
            return group(Output.of(group));
        }

        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder namePrefix(@Nullable Output<String> namePrefix) {
            $.namePrefix = namePrefix;
            return this;
        }

        public Builder namePrefix(String namePrefix) {
            return namePrefix(Output.of(namePrefix));
        }

        public Builder policy(Output<String> policy) {
            $.policy = policy;
            return this;
        }

        public Builder policy(String policy) {
            return policy(Output.of(policy));
        }

        public IamGroupPolicyArgs build() {
            if ($.group == null) {
                throw new MissingRequiredPropertyException("IamGroupPolicyArgs", "group");
            }
            if ($.policy == null) {
                throw new MissingRequiredPropertyException("IamGroupPolicyArgs", "policy");
            }
            return $;
        }
    }

}
