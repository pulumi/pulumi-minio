// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package minio

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-minio/sdk/go/minio/internal"
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
	case "minio:index/iamServiceAccount:IamServiceAccount":
		r = &IamServiceAccount{}
	case "minio:index/iamUser:IamUser":
		r = &IamUser{}
	case "minio:index/iamUserPolicyAttachment:IamUserPolicyAttachment":
		r = &IamUserPolicyAttachment{}
	case "minio:index/ilmPolicy:IlmPolicy":
		r = &IlmPolicy{}
	case "minio:index/kmsKey:KmsKey":
		r = &KmsKey{}
	case "minio:index/s3Bucket:S3Bucket":
		r = &S3Bucket{}
	case "minio:index/s3BucketNotification:S3BucketNotification":
		r = &S3BucketNotification{}
	case "minio:index/s3BucketPolicy:S3BucketPolicy":
		r = &S3BucketPolicy{}
	case "minio:index/s3BucketServerSideEncryption:S3BucketServerSideEncryption":
		r = &S3BucketServerSideEncryption{}
	case "minio:index/s3BucketVersioning:S3BucketVersioning":
		r = &S3BucketVersioning{}
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
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
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
		"index/iamServiceAccount",
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
		"index/kmsKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"minio",
		"index/s3Bucket",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"minio",
		"index/s3BucketNotification",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"minio",
		"index/s3BucketPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"minio",
		"index/s3BucketServerSideEncryption",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"minio",
		"index/s3BucketVersioning",
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
