// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.minio;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.minio.ProviderArgs;
import com.pulumi.minio.Utilities;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The provider type for the minio package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 * 
 */
@ResourceType(type="pulumi:providers:minio")
public class Provider extends com.pulumi.resources.ProviderResource {
    /**
     * Minio Access Key
     * 
     * @deprecated
     * use minio_user instead
     * 
     */
    @Deprecated /* use minio_user instead */
    @Export(name="minioAccessKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> minioAccessKey;

    /**
     * @return Minio Access Key
     * 
     */
    public Output<Optional<String>> minioAccessKey() {
        return Codegen.optional(this.minioAccessKey);
    }
    /**
     * Minio API Version (type: string, options: v2 or v4, default: v4)
     * 
     */
    @Export(name="minioApiVersion", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> minioApiVersion;

    /**
     * @return Minio API Version (type: string, options: v2 or v4, default: v4)
     * 
     */
    public Output<Optional<String>> minioApiVersion() {
        return Codegen.optional(this.minioApiVersion);
    }
    @Export(name="minioCacertFile", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> minioCacertFile;

    public Output<Optional<String>> minioCacertFile() {
        return Codegen.optional(this.minioCacertFile);
    }
    @Export(name="minioCertFile", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> minioCertFile;

    public Output<Optional<String>> minioCertFile() {
        return Codegen.optional(this.minioCertFile);
    }
    @Export(name="minioKeyFile", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> minioKeyFile;

    public Output<Optional<String>> minioKeyFile() {
        return Codegen.optional(this.minioKeyFile);
    }
    /**
     * Minio Password
     * 
     */
    @Export(name="minioPassword", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> minioPassword;

    /**
     * @return Minio Password
     * 
     */
    public Output<Optional<String>> minioPassword() {
        return Codegen.optional(this.minioPassword);
    }
    /**
     * Minio Region (default: us-east-1)
     * 
     */
    @Export(name="minioRegion", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> minioRegion;

    /**
     * @return Minio Region (default: us-east-1)
     * 
     */
    public Output<Optional<String>> minioRegion() {
        return Codegen.optional(this.minioRegion);
    }
    /**
     * Minio Secret Key
     * 
     * @deprecated
     * use minio_password instead
     * 
     */
    @Deprecated /* use minio_password instead */
    @Export(name="minioSecretKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> minioSecretKey;

    /**
     * @return Minio Secret Key
     * 
     */
    public Output<Optional<String>> minioSecretKey() {
        return Codegen.optional(this.minioSecretKey);
    }
    /**
     * Minio Host and Port
     * 
     */
    @Export(name="minioServer", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> minioServer;

    /**
     * @return Minio Host and Port
     * 
     */
    public Output<Optional<String>> minioServer() {
        return Codegen.optional(this.minioServer);
    }
    /**
     * Minio Session Token
     * 
     */
    @Export(name="minioSessionToken", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> minioSessionToken;

    /**
     * @return Minio Session Token
     * 
     */
    public Output<Optional<String>> minioSessionToken() {
        return Codegen.optional(this.minioSessionToken);
    }
    /**
     * Minio User
     * 
     */
    @Export(name="minioUser", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> minioUser;

    /**
     * @return Minio User
     * 
     */
    public Output<Optional<String>> minioUser() {
        return Codegen.optional(this.minioUser);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Provider(java.lang.String name) {
        this(name, ProviderArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Provider(java.lang.String name, @Nullable ProviderArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Provider(java.lang.String name, @Nullable ProviderArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("minio", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private static ProviderArgs makeArgs(@Nullable ProviderArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ProviderArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

}
