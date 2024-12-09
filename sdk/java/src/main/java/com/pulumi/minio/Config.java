// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.minio;

import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;

public final class Config {

    private static final com.pulumi.Config config = com.pulumi.Config.of("minio");
/**
 * Minio Access Key
 * 
 */
    public Optional<String> minioAccessKey() {
        return Codegen.stringProp("minioAccessKey").config(config).get();
    }
/**
 * Minio API Version (type: string, options: v2 or v4, default: v4)
 * 
 */
    public Optional<String> minioApiVersion() {
        return Codegen.stringProp("minioApiVersion").config(config).get();
    }
    public Optional<String> minioCacertFile() {
        return Codegen.stringProp("minioCacertFile").config(config).get();
    }
    public Optional<String> minioCertFile() {
        return Codegen.stringProp("minioCertFile").config(config).get();
    }
/**
 * Disable SSL certificate verification (default: false)
 * 
 */
    public Optional<Boolean> minioInsecure() {
        return Codegen.booleanProp("minioInsecure").config(config).get();
    }
    public Optional<String> minioKeyFile() {
        return Codegen.stringProp("minioKeyFile").config(config).get();
    }
/**
 * Minio Password
 * 
 */
    public Optional<String> minioPassword() {
        return Codegen.stringProp("minioPassword").config(config).get();
    }
/**
 * Minio Region (default: us-east-1)
 * 
 */
    public Optional<String> minioRegion() {
        return Codegen.stringProp("minioRegion").config(config).get();
    }
/**
 * Minio Secret Key
 * 
 */
    public Optional<String> minioSecretKey() {
        return Codegen.stringProp("minioSecretKey").config(config).get();
    }
/**
 * Minio Host and Port
 * 
 */
    public String minioServer() {
        return Codegen.stringProp("minioServer").config(config).require();
    }
/**
 * Minio Session Token
 * 
 */
    public Optional<String> minioSessionToken() {
        return Codegen.stringProp("minioSessionToken").config(config).get();
    }
/**
 * Minio SSL enabled (default: false)
 * 
 */
    public Optional<Boolean> minioSsl() {
        return Codegen.booleanProp("minioSsl").config(config).get();
    }
/**
 * Minio User
 * 
 */
    public Optional<String> minioUser() {
        return Codegen.stringProp("minioUser").config(config).get();
    }
}