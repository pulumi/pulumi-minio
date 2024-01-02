// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.minio;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class KmsKeyArgs extends com.pulumi.resources.ResourceArgs {

    public static final KmsKeyArgs Empty = new KmsKeyArgs();

    @Import(name="keyId", required=true)
    private Output<String> keyId;

    public Output<String> keyId() {
        return this.keyId;
    }

    private KmsKeyArgs() {}

    private KmsKeyArgs(KmsKeyArgs $) {
        this.keyId = $.keyId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(KmsKeyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private KmsKeyArgs $;

        public Builder() {
            $ = new KmsKeyArgs();
        }

        public Builder(KmsKeyArgs defaults) {
            $ = new KmsKeyArgs(Objects.requireNonNull(defaults));
        }

        public Builder keyId(Output<String> keyId) {
            $.keyId = keyId;
            return this;
        }

        public Builder keyId(String keyId) {
            return keyId(Output.of(keyId));
        }

        public KmsKeyArgs build() {
            if ($.keyId == null) {
                throw new MissingRequiredPropertyException("KmsKeyArgs", "keyId");
            }
            return $;
        }
    }

}
