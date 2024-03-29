// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.minio;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class IamGroupUserAttachmentArgs extends com.pulumi.resources.ResourceArgs {

    public static final IamGroupUserAttachmentArgs Empty = new IamGroupUserAttachmentArgs();

    @Import(name="groupName", required=true)
    private Output<String> groupName;

    public Output<String> groupName() {
        return this.groupName;
    }

    @Import(name="userName", required=true)
    private Output<String> userName;

    public Output<String> userName() {
        return this.userName;
    }

    private IamGroupUserAttachmentArgs() {}

    private IamGroupUserAttachmentArgs(IamGroupUserAttachmentArgs $) {
        this.groupName = $.groupName;
        this.userName = $.userName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IamGroupUserAttachmentArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IamGroupUserAttachmentArgs $;

        public Builder() {
            $ = new IamGroupUserAttachmentArgs();
        }

        public Builder(IamGroupUserAttachmentArgs defaults) {
            $ = new IamGroupUserAttachmentArgs(Objects.requireNonNull(defaults));
        }

        public Builder groupName(Output<String> groupName) {
            $.groupName = groupName;
            return this;
        }

        public Builder groupName(String groupName) {
            return groupName(Output.of(groupName));
        }

        public Builder userName(Output<String> userName) {
            $.userName = userName;
            return this;
        }

        public Builder userName(String userName) {
            return userName(Output.of(userName));
        }

        public IamGroupUserAttachmentArgs build() {
            if ($.groupName == null) {
                throw new MissingRequiredPropertyException("IamGroupUserAttachmentArgs", "groupName");
            }
            if ($.userName == null) {
                throw new MissingRequiredPropertyException("IamGroupUserAttachmentArgs", "userName");
            }
            return $;
        }
    }

}
