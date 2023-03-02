// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.minio.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class S3BucketNotificationQueueArgs extends com.pulumi.resources.ResourceArgs {

    public static final S3BucketNotificationQueueArgs Empty = new S3BucketNotificationQueueArgs();

    @Import(name="events", required=true)
    private Output<List<String>> events;

    public Output<List<String>> events() {
        return this.events;
    }

    @Import(name="filterPrefix")
    private @Nullable Output<String> filterPrefix;

    public Optional<Output<String>> filterPrefix() {
        return Optional.ofNullable(this.filterPrefix);
    }

    @Import(name="filterSuffix")
    private @Nullable Output<String> filterSuffix;

    public Optional<Output<String>> filterSuffix() {
        return Optional.ofNullable(this.filterSuffix);
    }

    /**
     * The ID of this resource.
     * 
     */
    @Import(name="id")
    private @Nullable Output<String> id;

    /**
     * @return The ID of this resource.
     * 
     */
    public Optional<Output<String>> id() {
        return Optional.ofNullable(this.id);
    }

    @Import(name="queueArn", required=true)
    private Output<String> queueArn;

    public Output<String> queueArn() {
        return this.queueArn;
    }

    private S3BucketNotificationQueueArgs() {}

    private S3BucketNotificationQueueArgs(S3BucketNotificationQueueArgs $) {
        this.events = $.events;
        this.filterPrefix = $.filterPrefix;
        this.filterSuffix = $.filterSuffix;
        this.id = $.id;
        this.queueArn = $.queueArn;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(S3BucketNotificationQueueArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private S3BucketNotificationQueueArgs $;

        public Builder() {
            $ = new S3BucketNotificationQueueArgs();
        }

        public Builder(S3BucketNotificationQueueArgs defaults) {
            $ = new S3BucketNotificationQueueArgs(Objects.requireNonNull(defaults));
        }

        public Builder events(Output<List<String>> events) {
            $.events = events;
            return this;
        }

        public Builder events(List<String> events) {
            return events(Output.of(events));
        }

        public Builder events(String... events) {
            return events(List.of(events));
        }

        public Builder filterPrefix(@Nullable Output<String> filterPrefix) {
            $.filterPrefix = filterPrefix;
            return this;
        }

        public Builder filterPrefix(String filterPrefix) {
            return filterPrefix(Output.of(filterPrefix));
        }

        public Builder filterSuffix(@Nullable Output<String> filterSuffix) {
            $.filterSuffix = filterSuffix;
            return this;
        }

        public Builder filterSuffix(String filterSuffix) {
            return filterSuffix(Output.of(filterSuffix));
        }

        /**
         * @param id The ID of this resource.
         * 
         * @return builder
         * 
         */
        public Builder id(@Nullable Output<String> id) {
            $.id = id;
            return this;
        }

        /**
         * @param id The ID of this resource.
         * 
         * @return builder
         * 
         */
        public Builder id(String id) {
            return id(Output.of(id));
        }

        public Builder queueArn(Output<String> queueArn) {
            $.queueArn = queueArn;
            return this;
        }

        public Builder queueArn(String queueArn) {
            return queueArn(Output.of(queueArn));
        }

        public S3BucketNotificationQueueArgs build() {
            $.events = Objects.requireNonNull($.events, "expected parameter 'events' to be non-null");
            $.queueArn = Objects.requireNonNull($.queueArn, "expected parameter 'queueArn' to be non-null");
            return $;
        }
    }

}