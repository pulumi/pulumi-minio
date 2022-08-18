// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.minio.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.minio.outputs.GetIamPolicyDocumentStatementCondition;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetIamPolicyDocumentStatement {
    private final @Nullable List<String> actions;
    private final @Nullable List<GetIamPolicyDocumentStatementCondition> conditions;
    private final @Nullable String effect;
    private final @Nullable String principal;
    private final @Nullable List<String> resources;
    private final @Nullable String sid;

    @CustomType.Constructor
    private GetIamPolicyDocumentStatement(
        @CustomType.Parameter("actions") @Nullable List<String> actions,
        @CustomType.Parameter("conditions") @Nullable List<GetIamPolicyDocumentStatementCondition> conditions,
        @CustomType.Parameter("effect") @Nullable String effect,
        @CustomType.Parameter("principal") @Nullable String principal,
        @CustomType.Parameter("resources") @Nullable List<String> resources,
        @CustomType.Parameter("sid") @Nullable String sid) {
        this.actions = actions;
        this.conditions = conditions;
        this.effect = effect;
        this.principal = principal;
        this.resources = resources;
        this.sid = sid;
    }

    public List<String> actions() {
        return this.actions == null ? List.of() : this.actions;
    }
    public List<GetIamPolicyDocumentStatementCondition> conditions() {
        return this.conditions == null ? List.of() : this.conditions;
    }
    public Optional<String> effect() {
        return Optional.ofNullable(this.effect);
    }
    public Optional<String> principal() {
        return Optional.ofNullable(this.principal);
    }
    public List<String> resources() {
        return this.resources == null ? List.of() : this.resources;
    }
    public Optional<String> sid() {
        return Optional.ofNullable(this.sid);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIamPolicyDocumentStatement defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable List<String> actions;
        private @Nullable List<GetIamPolicyDocumentStatementCondition> conditions;
        private @Nullable String effect;
        private @Nullable String principal;
        private @Nullable List<String> resources;
        private @Nullable String sid;

        public Builder() {
    	      // Empty
        }

        public Builder(GetIamPolicyDocumentStatement defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.actions = defaults.actions;
    	      this.conditions = defaults.conditions;
    	      this.effect = defaults.effect;
    	      this.principal = defaults.principal;
    	      this.resources = defaults.resources;
    	      this.sid = defaults.sid;
        }

        public Builder actions(@Nullable List<String> actions) {
            this.actions = actions;
            return this;
        }
        public Builder actions(String... actions) {
            return actions(List.of(actions));
        }
        public Builder conditions(@Nullable List<GetIamPolicyDocumentStatementCondition> conditions) {
            this.conditions = conditions;
            return this;
        }
        public Builder conditions(GetIamPolicyDocumentStatementCondition... conditions) {
            return conditions(List.of(conditions));
        }
        public Builder effect(@Nullable String effect) {
            this.effect = effect;
            return this;
        }
        public Builder principal(@Nullable String principal) {
            this.principal = principal;
            return this;
        }
        public Builder resources(@Nullable List<String> resources) {
            this.resources = resources;
            return this;
        }
        public Builder resources(String... resources) {
            return resources(List.of(resources));
        }
        public Builder sid(@Nullable String sid) {
            this.sid = sid;
            return this;
        }        public GetIamPolicyDocumentStatement build() {
            return new GetIamPolicyDocumentStatement(actions, conditions, effect, principal, resources, sid);
        }
    }
}