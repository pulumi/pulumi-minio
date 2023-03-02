// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package minio

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-minio/sdk/go/minio"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			bucketS3Bucket, err := minio.NewS3Bucket(ctx, "bucketS3Bucket", &minio.S3BucketArgs{
//				Bucket: pulumi.String("example-bucket"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = minio.NewS3BucketVersioning(ctx, "bucketS3BucketVersioning", &minio.S3BucketVersioningArgs{
//				Bucket: bucketS3Bucket.Bucket,
//				VersioningConfiguration: &minio.S3BucketVersioningVersioningConfigurationArgs{
//					Status: pulumi.String("Enabled"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type S3BucketVersioning struct {
	pulumi.CustomResourceState

	Bucket                  pulumi.StringOutput                             `pulumi:"bucket"`
	VersioningConfiguration S3BucketVersioningVersioningConfigurationOutput `pulumi:"versioningConfiguration"`
}

// NewS3BucketVersioning registers a new resource with the given unique name, arguments, and options.
func NewS3BucketVersioning(ctx *pulumi.Context,
	name string, args *S3BucketVersioningArgs, opts ...pulumi.ResourceOption) (*S3BucketVersioning, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.VersioningConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'VersioningConfiguration'")
	}
	var resource S3BucketVersioning
	err := ctx.RegisterResource("minio:index/s3BucketVersioning:S3BucketVersioning", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetS3BucketVersioning gets an existing S3BucketVersioning resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetS3BucketVersioning(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *S3BucketVersioningState, opts ...pulumi.ResourceOption) (*S3BucketVersioning, error) {
	var resource S3BucketVersioning
	err := ctx.ReadResource("minio:index/s3BucketVersioning:S3BucketVersioning", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering S3BucketVersioning resources.
type s3bucketVersioningState struct {
	Bucket                  *string                                    `pulumi:"bucket"`
	VersioningConfiguration *S3BucketVersioningVersioningConfiguration `pulumi:"versioningConfiguration"`
}

type S3BucketVersioningState struct {
	Bucket                  pulumi.StringPtrInput
	VersioningConfiguration S3BucketVersioningVersioningConfigurationPtrInput
}

func (S3BucketVersioningState) ElementType() reflect.Type {
	return reflect.TypeOf((*s3bucketVersioningState)(nil)).Elem()
}

type s3bucketVersioningArgs struct {
	Bucket                  string                                    `pulumi:"bucket"`
	VersioningConfiguration S3BucketVersioningVersioningConfiguration `pulumi:"versioningConfiguration"`
}

// The set of arguments for constructing a S3BucketVersioning resource.
type S3BucketVersioningArgs struct {
	Bucket                  pulumi.StringInput
	VersioningConfiguration S3BucketVersioningVersioningConfigurationInput
}

func (S3BucketVersioningArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*s3bucketVersioningArgs)(nil)).Elem()
}

type S3BucketVersioningInput interface {
	pulumi.Input

	ToS3BucketVersioningOutput() S3BucketVersioningOutput
	ToS3BucketVersioningOutputWithContext(ctx context.Context) S3BucketVersioningOutput
}

func (*S3BucketVersioning) ElementType() reflect.Type {
	return reflect.TypeOf((**S3BucketVersioning)(nil)).Elem()
}

func (i *S3BucketVersioning) ToS3BucketVersioningOutput() S3BucketVersioningOutput {
	return i.ToS3BucketVersioningOutputWithContext(context.Background())
}

func (i *S3BucketVersioning) ToS3BucketVersioningOutputWithContext(ctx context.Context) S3BucketVersioningOutput {
	return pulumi.ToOutputWithContext(ctx, i).(S3BucketVersioningOutput)
}

// S3BucketVersioningArrayInput is an input type that accepts S3BucketVersioningArray and S3BucketVersioningArrayOutput values.
// You can construct a concrete instance of `S3BucketVersioningArrayInput` via:
//
//	S3BucketVersioningArray{ S3BucketVersioningArgs{...} }
type S3BucketVersioningArrayInput interface {
	pulumi.Input

	ToS3BucketVersioningArrayOutput() S3BucketVersioningArrayOutput
	ToS3BucketVersioningArrayOutputWithContext(context.Context) S3BucketVersioningArrayOutput
}

type S3BucketVersioningArray []S3BucketVersioningInput

func (S3BucketVersioningArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*S3BucketVersioning)(nil)).Elem()
}

func (i S3BucketVersioningArray) ToS3BucketVersioningArrayOutput() S3BucketVersioningArrayOutput {
	return i.ToS3BucketVersioningArrayOutputWithContext(context.Background())
}

func (i S3BucketVersioningArray) ToS3BucketVersioningArrayOutputWithContext(ctx context.Context) S3BucketVersioningArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(S3BucketVersioningArrayOutput)
}

// S3BucketVersioningMapInput is an input type that accepts S3BucketVersioningMap and S3BucketVersioningMapOutput values.
// You can construct a concrete instance of `S3BucketVersioningMapInput` via:
//
//	S3BucketVersioningMap{ "key": S3BucketVersioningArgs{...} }
type S3BucketVersioningMapInput interface {
	pulumi.Input

	ToS3BucketVersioningMapOutput() S3BucketVersioningMapOutput
	ToS3BucketVersioningMapOutputWithContext(context.Context) S3BucketVersioningMapOutput
}

type S3BucketVersioningMap map[string]S3BucketVersioningInput

func (S3BucketVersioningMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*S3BucketVersioning)(nil)).Elem()
}

func (i S3BucketVersioningMap) ToS3BucketVersioningMapOutput() S3BucketVersioningMapOutput {
	return i.ToS3BucketVersioningMapOutputWithContext(context.Background())
}

func (i S3BucketVersioningMap) ToS3BucketVersioningMapOutputWithContext(ctx context.Context) S3BucketVersioningMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(S3BucketVersioningMapOutput)
}

type S3BucketVersioningOutput struct{ *pulumi.OutputState }

func (S3BucketVersioningOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**S3BucketVersioning)(nil)).Elem()
}

func (o S3BucketVersioningOutput) ToS3BucketVersioningOutput() S3BucketVersioningOutput {
	return o
}

func (o S3BucketVersioningOutput) ToS3BucketVersioningOutputWithContext(ctx context.Context) S3BucketVersioningOutput {
	return o
}

func (o S3BucketVersioningOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *S3BucketVersioning) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

func (o S3BucketVersioningOutput) VersioningConfiguration() S3BucketVersioningVersioningConfigurationOutput {
	return o.ApplyT(func(v *S3BucketVersioning) S3BucketVersioningVersioningConfigurationOutput {
		return v.VersioningConfiguration
	}).(S3BucketVersioningVersioningConfigurationOutput)
}

type S3BucketVersioningArrayOutput struct{ *pulumi.OutputState }

func (S3BucketVersioningArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*S3BucketVersioning)(nil)).Elem()
}

func (o S3BucketVersioningArrayOutput) ToS3BucketVersioningArrayOutput() S3BucketVersioningArrayOutput {
	return o
}

func (o S3BucketVersioningArrayOutput) ToS3BucketVersioningArrayOutputWithContext(ctx context.Context) S3BucketVersioningArrayOutput {
	return o
}

func (o S3BucketVersioningArrayOutput) Index(i pulumi.IntInput) S3BucketVersioningOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *S3BucketVersioning {
		return vs[0].([]*S3BucketVersioning)[vs[1].(int)]
	}).(S3BucketVersioningOutput)
}

type S3BucketVersioningMapOutput struct{ *pulumi.OutputState }

func (S3BucketVersioningMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*S3BucketVersioning)(nil)).Elem()
}

func (o S3BucketVersioningMapOutput) ToS3BucketVersioningMapOutput() S3BucketVersioningMapOutput {
	return o
}

func (o S3BucketVersioningMapOutput) ToS3BucketVersioningMapOutputWithContext(ctx context.Context) S3BucketVersioningMapOutput {
	return o
}

func (o S3BucketVersioningMapOutput) MapIndex(k pulumi.StringInput) S3BucketVersioningOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *S3BucketVersioning {
		return vs[0].(map[string]*S3BucketVersioning)[vs[1].(string)]
	}).(S3BucketVersioningOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*S3BucketVersioningInput)(nil)).Elem(), &S3BucketVersioning{})
	pulumi.RegisterInputType(reflect.TypeOf((*S3BucketVersioningArrayInput)(nil)).Elem(), S3BucketVersioningArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*S3BucketVersioningMapInput)(nil)).Elem(), S3BucketVersioningMap{})
	pulumi.RegisterOutputType(S3BucketVersioningOutput{})
	pulumi.RegisterOutputType(S3BucketVersioningArrayOutput{})
	pulumi.RegisterOutputType(S3BucketVersioningMapOutput{})
}