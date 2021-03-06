// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package minio

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the minio package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	// Minio Access Key
	MinioAccessKey pulumi.StringOutput `pulumi:"minioAccessKey"`
	// Minio API Version (type: string, options: v2 or v4, default: v4)
	MinioApiVersion pulumi.StringPtrOutput `pulumi:"minioApiVersion"`
	MinioCacertFile pulumi.StringPtrOutput `pulumi:"minioCacertFile"`
	MinioCertFile   pulumi.StringPtrOutput `pulumi:"minioCertFile"`
	MinioKeyFile    pulumi.StringPtrOutput `pulumi:"minioKeyFile"`
	// Minio Region (default: us-east-1)
	MinioRegion pulumi.StringPtrOutput `pulumi:"minioRegion"`
	// Minio Secret Key
	MinioSecretKey pulumi.StringOutput `pulumi:"minioSecretKey"`
	// Minio Host and Port
	MinioServer pulumi.StringOutput `pulumi:"minioServer"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MinioAccessKey == nil {
		return nil, errors.New("invalid value for required argument 'MinioAccessKey'")
	}
	if args.MinioSecretKey == nil {
		return nil, errors.New("invalid value for required argument 'MinioSecretKey'")
	}
	if args.MinioServer == nil {
		return nil, errors.New("invalid value for required argument 'MinioServer'")
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:minio", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// Minio Access Key
	MinioAccessKey string `pulumi:"minioAccessKey"`
	// Minio API Version (type: string, options: v2 or v4, default: v4)
	MinioApiVersion *string `pulumi:"minioApiVersion"`
	MinioCacertFile *string `pulumi:"minioCacertFile"`
	MinioCertFile   *string `pulumi:"minioCertFile"`
	MinioInsecure   *bool   `pulumi:"minioInsecure"`
	MinioKeyFile    *string `pulumi:"minioKeyFile"`
	// Minio Region (default: us-east-1)
	MinioRegion *string `pulumi:"minioRegion"`
	// Minio Secret Key
	MinioSecretKey string `pulumi:"minioSecretKey"`
	// Minio Host and Port
	MinioServer string `pulumi:"minioServer"`
	// Minio SSL enabled (default: false)
	MinioSsl *bool `pulumi:"minioSsl"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// Minio Access Key
	MinioAccessKey pulumi.StringInput
	// Minio API Version (type: string, options: v2 or v4, default: v4)
	MinioApiVersion pulumi.StringPtrInput
	MinioCacertFile pulumi.StringPtrInput
	MinioCertFile   pulumi.StringPtrInput
	MinioInsecure   pulumi.BoolPtrInput
	MinioKeyFile    pulumi.StringPtrInput
	// Minio Region (default: us-east-1)
	MinioRegion pulumi.StringPtrInput
	// Minio Secret Key
	MinioSecretKey pulumi.StringInput
	// Minio Host and Port
	MinioServer pulumi.StringInput
	// Minio SSL enabled (default: false)
	MinioSsl pulumi.BoolPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}
