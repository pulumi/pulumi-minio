// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.minio;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IamGroupMembershipArgs extends com.pulumi.resources.ResourceArgs {

    public static final IamGroupMembershipArgs Empty = new IamGroupMembershipArgs();

    /**
     * Group name to add users
     * 
     */
    @Import(name="group", required=true)
    private Output<String> group;

    /**
     * @return Group name to add users
     * 
     */
    public Output<String> group() {
        return this.group;
    }

    /**
     * Name of group membership
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of group membership
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Add user or list of users such as a group membership
     * 
     */
    @Import(name="users", required=true)
    private Output<List<String>> users;

    /**
     * @return Add user or list of users such as a group membership
     * 
     */
    public Output<List<String>> users() {
        return this.users;
    }

    private IamGroupMembershipArgs() {}

    private IamGroupMembershipArgs(IamGroupMembershipArgs $) {
        this.group = $.group;
        this.name = $.name;
        this.users = $.users;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IamGroupMembershipArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IamGroupMembershipArgs $;

        public Builder() {
            $ = new IamGroupMembershipArgs();
        }

        public Builder(IamGroupMembershipArgs defaults) {
            $ = new IamGroupMembershipArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param group Group name to add users
         * 
         * @return builder
         * 
         */
        public Builder group(Output<String> group) {
            $.group = group;
            return this;
        }

        /**
         * @param group Group name to add users
         * 
         * @return builder
         * 
         */
        public Builder group(String group) {
            return group(Output.of(group));
        }

        /**
         * @param name Name of group membership
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of group membership
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param users Add user or list of users such as a group membership
         * 
         * @return builder
         * 
         */
        public Builder users(Output<List<String>> users) {
            $.users = users;
            return this;
        }

        /**
         * @param users Add user or list of users such as a group membership
         * 
         * @return builder
         * 
         */
        public Builder users(List<String> users) {
            return users(Output.of(users));
        }

        /**
         * @param users Add user or list of users such as a group membership
         * 
         * @return builder
         * 
         */
        public Builder users(String... users) {
            return users(List.of(users));
        }

        public IamGroupMembershipArgs build() {
            if ($.group == null) {
                throw new MissingRequiredPropertyException("IamGroupMembershipArgs", "group");
            }
            if ($.users == null) {
                throw new MissingRequiredPropertyException("IamGroupMembershipArgs", "users");
            }
            return $;
        }
    }

}
