// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.minio;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.minio.inputs.IlmPolicyRuleArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class IlmPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final IlmPolicyArgs Empty = new IlmPolicyArgs();

    @Import(name="bucket", required=true)
    private Output<String> bucket;

    public Output<String> bucket() {
        return this.bucket;
    }

    @Import(name="rules", required=true)
    private Output<List<IlmPolicyRuleArgs>> rules;

    public Output<List<IlmPolicyRuleArgs>> rules() {
        return this.rules;
    }

    private IlmPolicyArgs() {}

    private IlmPolicyArgs(IlmPolicyArgs $) {
        this.bucket = $.bucket;
        this.rules = $.rules;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IlmPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IlmPolicyArgs $;

        public Builder() {
            $ = new IlmPolicyArgs();
        }

        public Builder(IlmPolicyArgs defaults) {
            $ = new IlmPolicyArgs(Objects.requireNonNull(defaults));
        }

        public Builder bucket(Output<String> bucket) {
            $.bucket = bucket;
            return this;
        }

        public Builder bucket(String bucket) {
            return bucket(Output.of(bucket));
        }

        public Builder rules(Output<List<IlmPolicyRuleArgs>> rules) {
            $.rules = rules;
            return this;
        }

        public Builder rules(List<IlmPolicyRuleArgs> rules) {
            return rules(Output.of(rules));
        }

        public Builder rules(IlmPolicyRuleArgs... rules) {
            return rules(List.of(rules));
        }

        public IlmPolicyArgs build() {
            if ($.bucket == null) {
                throw new MissingRequiredPropertyException("IlmPolicyArgs", "bucket");
            }
            if ($.rules == null) {
                throw new MissingRequiredPropertyException("IlmPolicyArgs", "rules");
            }
            return $;
        }
    }

}
