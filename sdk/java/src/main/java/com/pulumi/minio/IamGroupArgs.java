// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.minio;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IamGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final IamGroupArgs Empty = new IamGroupArgs();

    /**
     * Disable group
     * 
     */
    @Import(name="disableGroup")
    private @Nullable Output<Boolean> disableGroup;

    /**
     * @return Disable group
     * 
     */
    public Optional<Output<Boolean>> disableGroup() {
        return Optional.ofNullable(this.disableGroup);
    }

    /**
     * Delete group even if it has non-Terraform-managed members
     * 
     */
    @Import(name="forceDestroy")
    private @Nullable Output<Boolean> forceDestroy;

    /**
     * @return Delete group even if it has non-Terraform-managed members
     * 
     */
    public Optional<Output<Boolean>> forceDestroy() {
        return Optional.ofNullable(this.forceDestroy);
    }

    @Import(name="name")
    private @Nullable Output<String> name;

    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private IamGroupArgs() {}

    private IamGroupArgs(IamGroupArgs $) {
        this.disableGroup = $.disableGroup;
        this.forceDestroy = $.forceDestroy;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IamGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IamGroupArgs $;

        public Builder() {
            $ = new IamGroupArgs();
        }

        public Builder(IamGroupArgs defaults) {
            $ = new IamGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param disableGroup Disable group
         * 
         * @return builder
         * 
         */
        public Builder disableGroup(@Nullable Output<Boolean> disableGroup) {
            $.disableGroup = disableGroup;
            return this;
        }

        /**
         * @param disableGroup Disable group
         * 
         * @return builder
         * 
         */
        public Builder disableGroup(Boolean disableGroup) {
            return disableGroup(Output.of(disableGroup));
        }

        /**
         * @param forceDestroy Delete group even if it has non-Terraform-managed members
         * 
         * @return builder
         * 
         */
        public Builder forceDestroy(@Nullable Output<Boolean> forceDestroy) {
            $.forceDestroy = forceDestroy;
            return this;
        }

        /**
         * @param forceDestroy Delete group even if it has non-Terraform-managed members
         * 
         * @return builder
         * 
         */
        public Builder forceDestroy(Boolean forceDestroy) {
            return forceDestroy(Output.of(forceDestroy));
        }

        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        public Builder name(String name) {
            return name(Output.of(name));
        }

        public IamGroupArgs build() {
            return $;
        }
    }

}
