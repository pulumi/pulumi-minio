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


public final class IamServiceAccountArgs extends com.pulumi.resources.ResourceArgs {

    public static final IamServiceAccountArgs Empty = new IamServiceAccountArgs();

    /**
     * Disable service account
     * 
     */
    @Import(name="disableUser")
    private @Nullable Output<Boolean> disableUser;

    /**
     * @return Disable service account
     * 
     */
    public Optional<Output<Boolean>> disableUser() {
        return Optional.ofNullable(this.disableUser);
    }

    /**
     * policy of service account
     * 
     */
    @Import(name="policy")
    private @Nullable Output<String> policy;

    /**
     * @return policy of service account
     * 
     */
    public Optional<Output<String>> policy() {
        return Optional.ofNullable(this.policy);
    }

    @Import(name="targetUser", required=true)
    private Output<String> targetUser;

    public Output<String> targetUser() {
        return this.targetUser;
    }

    /**
     * rotate secret key
     * 
     */
    @Import(name="updateSecret")
    private @Nullable Output<Boolean> updateSecret;

    /**
     * @return rotate secret key
     * 
     */
    public Optional<Output<Boolean>> updateSecret() {
        return Optional.ofNullable(this.updateSecret);
    }

    private IamServiceAccountArgs() {}

    private IamServiceAccountArgs(IamServiceAccountArgs $) {
        this.disableUser = $.disableUser;
        this.policy = $.policy;
        this.targetUser = $.targetUser;
        this.updateSecret = $.updateSecret;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IamServiceAccountArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IamServiceAccountArgs $;

        public Builder() {
            $ = new IamServiceAccountArgs();
        }

        public Builder(IamServiceAccountArgs defaults) {
            $ = new IamServiceAccountArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param disableUser Disable service account
         * 
         * @return builder
         * 
         */
        public Builder disableUser(@Nullable Output<Boolean> disableUser) {
            $.disableUser = disableUser;
            return this;
        }

        /**
         * @param disableUser Disable service account
         * 
         * @return builder
         * 
         */
        public Builder disableUser(Boolean disableUser) {
            return disableUser(Output.of(disableUser));
        }

        /**
         * @param policy policy of service account
         * 
         * @return builder
         * 
         */
        public Builder policy(@Nullable Output<String> policy) {
            $.policy = policy;
            return this;
        }

        /**
         * @param policy policy of service account
         * 
         * @return builder
         * 
         */
        public Builder policy(String policy) {
            return policy(Output.of(policy));
        }

        public Builder targetUser(Output<String> targetUser) {
            $.targetUser = targetUser;
            return this;
        }

        public Builder targetUser(String targetUser) {
            return targetUser(Output.of(targetUser));
        }

        /**
         * @param updateSecret rotate secret key
         * 
         * @return builder
         * 
         */
        public Builder updateSecret(@Nullable Output<Boolean> updateSecret) {
            $.updateSecret = updateSecret;
            return this;
        }

        /**
         * @param updateSecret rotate secret key
         * 
         * @return builder
         * 
         */
        public Builder updateSecret(Boolean updateSecret) {
            return updateSecret(Output.of(updateSecret));
        }

        public IamServiceAccountArgs build() {
            $.targetUser = Objects.requireNonNull($.targetUser, "expected parameter 'targetUser' to be non-null");
            return $;
        }
    }

}
