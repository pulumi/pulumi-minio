module github.com/pulumi/pulumi-minio/provider

go 1.16

replace (
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20211019194827-62530c6537a4
	github.com/minio/minio => github.com/minio/minio v0.0.0-20210409155609-0ddc4f00756b
	github.com/minio/minio-go/v7 => github.com/minio/minio-go/v7 v7.0.11-0.20210302210017-6ae69c73ce78
	github.com/shirou/gopsutil/v3 => github.com/shirou/gopsutil/v3 v3.21.1
)

require (
	github.com/aminueza/terraform-provider-minio v1.4.0
	github.com/hashicorp/terraform-plugin-sdk v1.9.1 // indirect
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.20.0
	github.com/pulumi/pulumi/sdk/v3 v3.27.0
)
