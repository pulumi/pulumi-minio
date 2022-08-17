// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.minio.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class GetIamPolicyDocumentStatementCondition extends com.pulumi.resources.InvokeArgs {

    public static final GetIamPolicyDocumentStatementCondition Empty = new GetIamPolicyDocumentStatementCondition();

    @Import(name="test", required=true)
    private String test;

    public String test() {
        return this.test;
    }

    @Import(name="values", required=true)
    private List<String> values;

    public List<String> values() {
        return this.values;
    }

    @Import(name="variable", required=true)
    private String variable;

    public String variable() {
        return this.variable;
    }

    private GetIamPolicyDocumentStatementCondition() {}

    private GetIamPolicyDocumentStatementCondition(GetIamPolicyDocumentStatementCondition $) {
        this.test = $.test;
        this.values = $.values;
        this.variable = $.variable;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetIamPolicyDocumentStatementCondition defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetIamPolicyDocumentStatementCondition $;

        public Builder() {
            $ = new GetIamPolicyDocumentStatementCondition();
        }

        public Builder(GetIamPolicyDocumentStatementCondition defaults) {
            $ = new GetIamPolicyDocumentStatementCondition(Objects.requireNonNull(defaults));
        }

        public Builder test(String test) {
            $.test = test;
            return this;
        }

        public Builder values(List<String> values) {
            $.values = values;
            return this;
        }

        public Builder values(String... values) {
            return values(List.of(values));
        }

        public Builder variable(String variable) {
            $.variable = variable;
            return this;
        }

        public GetIamPolicyDocumentStatementCondition build() {
            $.test = Objects.requireNonNull($.test, "expected parameter 'test' to be non-null");
            $.values = Objects.requireNonNull($.values, "expected parameter 'values' to be non-null");
            $.variable = Objects.requireNonNull($.variable, "expected parameter 'variable' to be non-null");
            return $;
        }
    }

}
