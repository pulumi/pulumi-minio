// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package minio

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "minio:index/iamGroup:IamGroup":
		r = &IamGroup{}
	case "minio:index/iamGroupMembership:IamGroupMembership":
		r = &IamGroupMembership{}
	case "minio:index/iamGroupPolicy:IamGroupPolicy":
		r = &IamGroupPolicy{}
	case "minio:index/iamGroupPolicyAttachment:IamGroupPolicyAttachment":
		r = &IamGroupPolicyAttachment{}
	case "minio:index/iamGroupUserAttachment:IamGroupUserAttachment":
		r = &IamGroupUserAttachment{}
	case "minio:index/iamPolicy:IamPolicy":
		r = &IamPolicy{}
	case "minio:index/iamUser:IamUser":
		r = &IamUser{}
	case "minio:index/iamUserPolicyAttachment:IamUserPolicyAttachment":
		r = &IamUserPolicyAttachment{}
	case "minio:index/ilmPolicy:IlmPolicy":
		r = &IlmPolicy{}
	case "minio:index/s3Bucket:S3Bucket":
		r = &S3Bucket{}
	case "minio:index/s3Object:S3Object":
		r = &S3Object{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:minio" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := PkgVersion()
	if err != nil {
		fmt.Printf("failed to determine package version. defaulting to v1: %v\n", err)
	}
	pulumi.RegisterResourceModule(
		"minio",
		"index/iamGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"minio",
		"index/iamGroupMembership",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"minio",
		"index/iamGroupPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"minio",
		"index/iamGroupPolicyAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"minio",
		"index/iamGroupUserAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"minio",
		"index/iamPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"minio",
		"index/iamUser",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"minio",
		"index/iamUserPolicyAttachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"minio",
		"index/ilmPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"minio",
		"index/s3Bucket",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"minio",
		"index/s3Object",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"minio",
		&pkg{version},
	)
}
