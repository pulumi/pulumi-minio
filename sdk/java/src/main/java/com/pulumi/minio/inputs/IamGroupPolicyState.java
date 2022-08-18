// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.minio.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IamGroupPolicyState extends com.pulumi.resources.ResourceArgs {

    public static final IamGroupPolicyState Empty = new IamGroupPolicyState();

    @Import(name="group")
    private @Nullable Output<String> group;

    public Optional<Output<String>> group() {
        return Optional.ofNullable(this.group);
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

    @Import(name="policy")
    private @Nullable Output<String> policy;

    public Optional<Output<String>> policy() {
        return Optional.ofNullable(this.policy);
    }

    private IamGroupPolicyState() {}

    private IamGroupPolicyState(IamGroupPolicyState $) {
        this.group = $.group;
        this.name = $.name;
        this.namePrefix = $.namePrefix;
        this.policy = $.policy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IamGroupPolicyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IamGroupPolicyState $;

        public Builder() {
            $ = new IamGroupPolicyState();
        }

        public Builder(IamGroupPolicyState defaults) {
            $ = new IamGroupPolicyState(Objects.requireNonNull(defaults));
        }

        public Builder group(@Nullable Output<String> group) {
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

        public Builder policy(@Nullable Output<String> policy) {
            $.policy = policy;
            return this;
        }

        public Builder policy(String policy) {
            return policy(Output.of(policy));
        }

        public IamGroupPolicyState build() {
            return $;
        }
    }

}