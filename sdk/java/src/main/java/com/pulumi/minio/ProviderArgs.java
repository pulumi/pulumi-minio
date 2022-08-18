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


public final class ProviderArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProviderArgs Empty = new ProviderArgs();

    /**
     * Minio Access Key
     * 
     */
    @Import(name="minioAccessKey", required=true)
    private Output<String> minioAccessKey;

    /**
     * @return Minio Access Key
     * 
     */
    public Output<String> minioAccessKey() {
        return this.minioAccessKey;
    }

    /**
     * Minio API Version (type: string, options: v2 or v4, default: v4)
     * 
     */
    @Import(name="minioApiVersion")
    private @Nullable Output<String> minioApiVersion;

    /**
     * @return Minio API Version (type: string, options: v2 or v4, default: v4)
     * 
     */
    public Optional<Output<String>> minioApiVersion() {
        return Optional.ofNullable(this.minioApiVersion);
    }

    @Import(name="minioCacertFile")
    private @Nullable Output<String> minioCacertFile;

    public Optional<Output<String>> minioCacertFile() {
        return Optional.ofNullable(this.minioCacertFile);
    }

    @Import(name="minioCertFile")
    private @Nullable Output<String> minioCertFile;

    public Optional<Output<String>> minioCertFile() {
        return Optional.ofNullable(this.minioCertFile);
    }

    @Import(name="minioInsecure", json=true)
    private @Nullable Output<Boolean> minioInsecure;

    public Optional<Output<Boolean>> minioInsecure() {
        return Optional.ofNullable(this.minioInsecure);
    }

    @Import(name="minioKeyFile")
    private @Nullable Output<String> minioKeyFile;

    public Optional<Output<String>> minioKeyFile() {
        return Optional.ofNullable(this.minioKeyFile);
    }

    /**
     * Minio Region (default: us-east-1)
     * 
     */
    @Import(name="minioRegion")
    private @Nullable Output<String> minioRegion;

    /**
     * @return Minio Region (default: us-east-1)
     * 
     */
    public Optional<Output<String>> minioRegion() {
        return Optional.ofNullable(this.minioRegion);
    }

    /**
     * Minio Secret Key
     * 
     */
    @Import(name="minioSecretKey", required=true)
    private Output<String> minioSecretKey;

    /**
     * @return Minio Secret Key
     * 
     */
    public Output<String> minioSecretKey() {
        return this.minioSecretKey;
    }

    /**
     * Minio Host and Port
     * 
     */
    @Import(name="minioServer", required=true)
    private Output<String> minioServer;

    /**
     * @return Minio Host and Port
     * 
     */
    public Output<String> minioServer() {
        return this.minioServer;
    }

    /**
     * Minio SSL enabled (default: false)
     * 
     */
    @Import(name="minioSsl", json=true)
    private @Nullable Output<Boolean> minioSsl;

    /**
     * @return Minio SSL enabled (default: false)
     * 
     */
    public Optional<Output<Boolean>> minioSsl() {
        return Optional.ofNullable(this.minioSsl);
    }

    private ProviderArgs() {}

    private ProviderArgs(ProviderArgs $) {
        this.minioAccessKey = $.minioAccessKey;
        this.minioApiVersion = $.minioApiVersion;
        this.minioCacertFile = $.minioCacertFile;
        this.minioCertFile = $.minioCertFile;
        this.minioInsecure = $.minioInsecure;
        this.minioKeyFile = $.minioKeyFile;
        this.minioRegion = $.minioRegion;
        this.minioSecretKey = $.minioSecretKey;
        this.minioServer = $.minioServer;
        this.minioSsl = $.minioSsl;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProviderArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProviderArgs $;

        public Builder() {
            $ = new ProviderArgs();
        }

        public Builder(ProviderArgs defaults) {
            $ = new ProviderArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param minioAccessKey Minio Access Key
         * 
         * @return builder
         * 
         */
        public Builder minioAccessKey(Output<String> minioAccessKey) {
            $.minioAccessKey = minioAccessKey;
            return this;
        }

        /**
         * @param minioAccessKey Minio Access Key
         * 
         * @return builder
         * 
         */
        public Builder minioAccessKey(String minioAccessKey) {
            return minioAccessKey(Output.of(minioAccessKey));
        }

        /**
         * @param minioApiVersion Minio API Version (type: string, options: v2 or v4, default: v4)
         * 
         * @return builder
         * 
         */
        public Builder minioApiVersion(@Nullable Output<String> minioApiVersion) {
            $.minioApiVersion = minioApiVersion;
            return this;
        }

        /**
         * @param minioApiVersion Minio API Version (type: string, options: v2 or v4, default: v4)
         * 
         * @return builder
         * 
         */
        public Builder minioApiVersion(String minioApiVersion) {
            return minioApiVersion(Output.of(minioApiVersion));
        }

        public Builder minioCacertFile(@Nullable Output<String> minioCacertFile) {
            $.minioCacertFile = minioCacertFile;
            return this;
        }

        public Builder minioCacertFile(String minioCacertFile) {
            return minioCacertFile(Output.of(minioCacertFile));
        }

        public Builder minioCertFile(@Nullable Output<String> minioCertFile) {
            $.minioCertFile = minioCertFile;
            return this;
        }

        public Builder minioCertFile(String minioCertFile) {
            return minioCertFile(Output.of(minioCertFile));
        }

        public Builder minioInsecure(@Nullable Output<Boolean> minioInsecure) {
            $.minioInsecure = minioInsecure;
            return this;
        }

        public Builder minioInsecure(Boolean minioInsecure) {
            return minioInsecure(Output.of(minioInsecure));
        }

        public Builder minioKeyFile(@Nullable Output<String> minioKeyFile) {
            $.minioKeyFile = minioKeyFile;
            return this;
        }

        public Builder minioKeyFile(String minioKeyFile) {
            return minioKeyFile(Output.of(minioKeyFile));
        }

        /**
         * @param minioRegion Minio Region (default: us-east-1)
         * 
         * @return builder
         * 
         */
        public Builder minioRegion(@Nullable Output<String> minioRegion) {
            $.minioRegion = minioRegion;
            return this;
        }

        /**
         * @param minioRegion Minio Region (default: us-east-1)
         * 
         * @return builder
         * 
         */
        public Builder minioRegion(String minioRegion) {
            return minioRegion(Output.of(minioRegion));
        }

        /**
         * @param minioSecretKey Minio Secret Key
         * 
         * @return builder
         * 
         */
        public Builder minioSecretKey(Output<String> minioSecretKey) {
            $.minioSecretKey = minioSecretKey;
            return this;
        }

        /**
         * @param minioSecretKey Minio Secret Key
         * 
         * @return builder
         * 
         */
        public Builder minioSecretKey(String minioSecretKey) {
            return minioSecretKey(Output.of(minioSecretKey));
        }

        /**
         * @param minioServer Minio Host and Port
         * 
         * @return builder
         * 
         */
        public Builder minioServer(Output<String> minioServer) {
            $.minioServer = minioServer;
            return this;
        }

        /**
         * @param minioServer Minio Host and Port
         * 
         * @return builder
         * 
         */
        public Builder minioServer(String minioServer) {
            return minioServer(Output.of(minioServer));
        }

        /**
         * @param minioSsl Minio SSL enabled (default: false)
         * 
         * @return builder
         * 
         */
        public Builder minioSsl(@Nullable Output<Boolean> minioSsl) {
            $.minioSsl = minioSsl;
            return this;
        }

        /**
         * @param minioSsl Minio SSL enabled (default: false)
         * 
         * @return builder
         * 
         */
        public Builder minioSsl(Boolean minioSsl) {
            return minioSsl(Output.of(minioSsl));
        }

        public ProviderArgs build() {
            $.minioAccessKey = Objects.requireNonNull($.minioAccessKey, "expected parameter 'minioAccessKey' to be non-null");
            $.minioSecretKey = Objects.requireNonNull($.minioSecretKey, "expected parameter 'minioSecretKey' to be non-null");
            $.minioServer = Objects.requireNonNull($.minioServer, "expected parameter 'minioServer' to be non-null");
            return $;
        }
    }

}