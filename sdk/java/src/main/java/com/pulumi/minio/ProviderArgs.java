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
     * @deprecated
     * use minio_user instead
     * 
     */
    @Deprecated /* use minio_user instead */
    @Import(name="minioAccessKey")
    private @Nullable Output<String> minioAccessKey;

    /**
     * @return Minio Access Key
     * 
     * @deprecated
     * use minio_user instead
     * 
     */
    @Deprecated /* use minio_user instead */
    public Optional<Output<String>> minioAccessKey() {
        return Optional.ofNullable(this.minioAccessKey);
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

    /**
     * Disable SSL certificate verification (default: false)
     * 
     */
    @Import(name="minioInsecure", json=true)
    private @Nullable Output<Boolean> minioInsecure;

    /**
     * @return Disable SSL certificate verification (default: false)
     * 
     */
    public Optional<Output<Boolean>> minioInsecure() {
        return Optional.ofNullable(this.minioInsecure);
    }

    @Import(name="minioKeyFile")
    private @Nullable Output<String> minioKeyFile;

    public Optional<Output<String>> minioKeyFile() {
        return Optional.ofNullable(this.minioKeyFile);
    }

    /**
     * Minio Password
     * 
     */
    @Import(name="minioPassword")
    private @Nullable Output<String> minioPassword;

    /**
     * @return Minio Password
     * 
     */
    public Optional<Output<String>> minioPassword() {
        return Optional.ofNullable(this.minioPassword);
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
     * @deprecated
     * use minio_password instead
     * 
     */
    @Deprecated /* use minio_password instead */
    @Import(name="minioSecretKey")
    private @Nullable Output<String> minioSecretKey;

    /**
     * @return Minio Secret Key
     * 
     * @deprecated
     * use minio_password instead
     * 
     */
    @Deprecated /* use minio_password instead */
    public Optional<Output<String>> minioSecretKey() {
        return Optional.ofNullable(this.minioSecretKey);
    }

    /**
     * Minio Host and Port
     * 
     */
    @Import(name="minioServer")
    private @Nullable Output<String> minioServer;

    /**
     * @return Minio Host and Port
     * 
     */
    public Optional<Output<String>> minioServer() {
        return Optional.ofNullable(this.minioServer);
    }

    /**
     * Minio Session Token
     * 
     */
    @Import(name="minioSessionToken")
    private @Nullable Output<String> minioSessionToken;

    /**
     * @return Minio Session Token
     * 
     */
    public Optional<Output<String>> minioSessionToken() {
        return Optional.ofNullable(this.minioSessionToken);
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

    /**
     * Minio User
     * 
     */
    @Import(name="minioUser")
    private @Nullable Output<String> minioUser;

    /**
     * @return Minio User
     * 
     */
    public Optional<Output<String>> minioUser() {
        return Optional.ofNullable(this.minioUser);
    }

    private ProviderArgs() {}

    private ProviderArgs(ProviderArgs $) {
        this.minioAccessKey = $.minioAccessKey;
        this.minioApiVersion = $.minioApiVersion;
        this.minioCacertFile = $.minioCacertFile;
        this.minioCertFile = $.minioCertFile;
        this.minioInsecure = $.minioInsecure;
        this.minioKeyFile = $.minioKeyFile;
        this.minioPassword = $.minioPassword;
        this.minioRegion = $.minioRegion;
        this.minioSecretKey = $.minioSecretKey;
        this.minioServer = $.minioServer;
        this.minioSessionToken = $.minioSessionToken;
        this.minioSsl = $.minioSsl;
        this.minioUser = $.minioUser;
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
         * @deprecated
         * use minio_user instead
         * 
         */
        @Deprecated /* use minio_user instead */
        public Builder minioAccessKey(@Nullable Output<String> minioAccessKey) {
            $.minioAccessKey = minioAccessKey;
            return this;
        }

        /**
         * @param minioAccessKey Minio Access Key
         * 
         * @return builder
         * 
         * @deprecated
         * use minio_user instead
         * 
         */
        @Deprecated /* use minio_user instead */
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

        /**
         * @param minioInsecure Disable SSL certificate verification (default: false)
         * 
         * @return builder
         * 
         */
        public Builder minioInsecure(@Nullable Output<Boolean> minioInsecure) {
            $.minioInsecure = minioInsecure;
            return this;
        }

        /**
         * @param minioInsecure Disable SSL certificate verification (default: false)
         * 
         * @return builder
         * 
         */
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
         * @param minioPassword Minio Password
         * 
         * @return builder
         * 
         */
        public Builder minioPassword(@Nullable Output<String> minioPassword) {
            $.minioPassword = minioPassword;
            return this;
        }

        /**
         * @param minioPassword Minio Password
         * 
         * @return builder
         * 
         */
        public Builder minioPassword(String minioPassword) {
            return minioPassword(Output.of(minioPassword));
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
         * @deprecated
         * use minio_password instead
         * 
         */
        @Deprecated /* use minio_password instead */
        public Builder minioSecretKey(@Nullable Output<String> minioSecretKey) {
            $.minioSecretKey = minioSecretKey;
            return this;
        }

        /**
         * @param minioSecretKey Minio Secret Key
         * 
         * @return builder
         * 
         * @deprecated
         * use minio_password instead
         * 
         */
        @Deprecated /* use minio_password instead */
        public Builder minioSecretKey(String minioSecretKey) {
            return minioSecretKey(Output.of(minioSecretKey));
        }

        /**
         * @param minioServer Minio Host and Port
         * 
         * @return builder
         * 
         */
        public Builder minioServer(@Nullable Output<String> minioServer) {
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
         * @param minioSessionToken Minio Session Token
         * 
         * @return builder
         * 
         */
        public Builder minioSessionToken(@Nullable Output<String> minioSessionToken) {
            $.minioSessionToken = minioSessionToken;
            return this;
        }

        /**
         * @param minioSessionToken Minio Session Token
         * 
         * @return builder
         * 
         */
        public Builder minioSessionToken(String minioSessionToken) {
            return minioSessionToken(Output.of(minioSessionToken));
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

        /**
         * @param minioUser Minio User
         * 
         * @return builder
         * 
         */
        public Builder minioUser(@Nullable Output<String> minioUser) {
            $.minioUser = minioUser;
            return this;
        }

        /**
         * @param minioUser Minio User
         * 
         * @return builder
         * 
         */
        public Builder minioUser(String minioUser) {
            return minioUser(Output.of(minioUser));
        }

        public ProviderArgs build() {
            return $;
        }
    }

}
