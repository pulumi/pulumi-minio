// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// Minio Access Key
//
// Deprecated: use minio_user instead
func GetMinioAccessKey(ctx *pulumi.Context) string {
	return config.Get(ctx, "minio:minioAccessKey")
}

// Minio API Version (type: string, options: v2 or v4, default: v4)
func GetMinioApiVersion(ctx *pulumi.Context) string {
	return config.Get(ctx, "minio:minioApiVersion")
}
func GetMinioCacertFile(ctx *pulumi.Context) string {
	return config.Get(ctx, "minio:minioCacertFile")
}
func GetMinioCertFile(ctx *pulumi.Context) string {
	return config.Get(ctx, "minio:minioCertFile")
}

// Disable SSL certificate verification (default: false)
func GetMinioInsecure(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "minio:minioInsecure")
}
func GetMinioKeyFile(ctx *pulumi.Context) string {
	return config.Get(ctx, "minio:minioKeyFile")
}

// Minio Password
func GetMinioPassword(ctx *pulumi.Context) string {
	return config.Get(ctx, "minio:minioPassword")
}

// Minio Region (default: us-east-1)
func GetMinioRegion(ctx *pulumi.Context) string {
	return config.Get(ctx, "minio:minioRegion")
}

// Minio Secret Key
//
// Deprecated: use minio_password instead
func GetMinioSecretKey(ctx *pulumi.Context) string {
	return config.Get(ctx, "minio:minioSecretKey")
}

// Minio Host and Port
func GetMinioServer(ctx *pulumi.Context) string {
	return config.Get(ctx, "minio:minioServer")
}

// Minio Session Token
func GetMinioSessionToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "minio:minioSessionToken")
}

// Minio SSL enabled (default: false)
func GetMinioSsl(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "minio:minioSsl")
}

// Minio User
func GetMinioUser(ctx *pulumi.Context) string {
	return config.Get(ctx, "minio:minioUser")
}
