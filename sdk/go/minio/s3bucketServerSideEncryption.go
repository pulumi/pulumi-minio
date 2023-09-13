// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package minio

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-minio/sdk/go/minio/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type S3BucketServerSideEncryption struct {
	pulumi.CustomResourceState

	Bucket         pulumi.StringOutput `pulumi:"bucket"`
	EncryptionType pulumi.StringOutput `pulumi:"encryptionType"`
	KmsKeyId       pulumi.StringOutput `pulumi:"kmsKeyId"`
}

// NewS3BucketServerSideEncryption registers a new resource with the given unique name, arguments, and options.
func NewS3BucketServerSideEncryption(ctx *pulumi.Context,
	name string, args *S3BucketServerSideEncryptionArgs, opts ...pulumi.ResourceOption) (*S3BucketServerSideEncryption, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.EncryptionType == nil {
		return nil, errors.New("invalid value for required argument 'EncryptionType'")
	}
	if args.KmsKeyId == nil {
		return nil, errors.New("invalid value for required argument 'KmsKeyId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource S3BucketServerSideEncryption
	err := ctx.RegisterResource("minio:index/s3BucketServerSideEncryption:S3BucketServerSideEncryption", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetS3BucketServerSideEncryption gets an existing S3BucketServerSideEncryption resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetS3BucketServerSideEncryption(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *S3BucketServerSideEncryptionState, opts ...pulumi.ResourceOption) (*S3BucketServerSideEncryption, error) {
	var resource S3BucketServerSideEncryption
	err := ctx.ReadResource("minio:index/s3BucketServerSideEncryption:S3BucketServerSideEncryption", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering S3BucketServerSideEncryption resources.
type s3bucketServerSideEncryptionState struct {
	Bucket         *string `pulumi:"bucket"`
	EncryptionType *string `pulumi:"encryptionType"`
	KmsKeyId       *string `pulumi:"kmsKeyId"`
}

type S3BucketServerSideEncryptionState struct {
	Bucket         pulumi.StringPtrInput
	EncryptionType pulumi.StringPtrInput
	KmsKeyId       pulumi.StringPtrInput
}

func (S3BucketServerSideEncryptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*s3bucketServerSideEncryptionState)(nil)).Elem()
}

type s3bucketServerSideEncryptionArgs struct {
	Bucket         string `pulumi:"bucket"`
	EncryptionType string `pulumi:"encryptionType"`
	KmsKeyId       string `pulumi:"kmsKeyId"`
}

// The set of arguments for constructing a S3BucketServerSideEncryption resource.
type S3BucketServerSideEncryptionArgs struct {
	Bucket         pulumi.StringInput
	EncryptionType pulumi.StringInput
	KmsKeyId       pulumi.StringInput
}

func (S3BucketServerSideEncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*s3bucketServerSideEncryptionArgs)(nil)).Elem()
}

type S3BucketServerSideEncryptionInput interface {
	pulumi.Input

	ToS3BucketServerSideEncryptionOutput() S3BucketServerSideEncryptionOutput
	ToS3BucketServerSideEncryptionOutputWithContext(ctx context.Context) S3BucketServerSideEncryptionOutput
}

func (*S3BucketServerSideEncryption) ElementType() reflect.Type {
	return reflect.TypeOf((**S3BucketServerSideEncryption)(nil)).Elem()
}

func (i *S3BucketServerSideEncryption) ToS3BucketServerSideEncryptionOutput() S3BucketServerSideEncryptionOutput {
	return i.ToS3BucketServerSideEncryptionOutputWithContext(context.Background())
}

func (i *S3BucketServerSideEncryption) ToS3BucketServerSideEncryptionOutputWithContext(ctx context.Context) S3BucketServerSideEncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(S3BucketServerSideEncryptionOutput)
}

func (i *S3BucketServerSideEncryption) ToOutput(ctx context.Context) pulumix.Output[*S3BucketServerSideEncryption] {
	return pulumix.Output[*S3BucketServerSideEncryption]{
		OutputState: i.ToS3BucketServerSideEncryptionOutputWithContext(ctx).OutputState,
	}
}

// S3BucketServerSideEncryptionArrayInput is an input type that accepts S3BucketServerSideEncryptionArray and S3BucketServerSideEncryptionArrayOutput values.
// You can construct a concrete instance of `S3BucketServerSideEncryptionArrayInput` via:
//
//	S3BucketServerSideEncryptionArray{ S3BucketServerSideEncryptionArgs{...} }
type S3BucketServerSideEncryptionArrayInput interface {
	pulumi.Input

	ToS3BucketServerSideEncryptionArrayOutput() S3BucketServerSideEncryptionArrayOutput
	ToS3BucketServerSideEncryptionArrayOutputWithContext(context.Context) S3BucketServerSideEncryptionArrayOutput
}

type S3BucketServerSideEncryptionArray []S3BucketServerSideEncryptionInput

func (S3BucketServerSideEncryptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*S3BucketServerSideEncryption)(nil)).Elem()
}

func (i S3BucketServerSideEncryptionArray) ToS3BucketServerSideEncryptionArrayOutput() S3BucketServerSideEncryptionArrayOutput {
	return i.ToS3BucketServerSideEncryptionArrayOutputWithContext(context.Background())
}

func (i S3BucketServerSideEncryptionArray) ToS3BucketServerSideEncryptionArrayOutputWithContext(ctx context.Context) S3BucketServerSideEncryptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(S3BucketServerSideEncryptionArrayOutput)
}

func (i S3BucketServerSideEncryptionArray) ToOutput(ctx context.Context) pulumix.Output[[]*S3BucketServerSideEncryption] {
	return pulumix.Output[[]*S3BucketServerSideEncryption]{
		OutputState: i.ToS3BucketServerSideEncryptionArrayOutputWithContext(ctx).OutputState,
	}
}

// S3BucketServerSideEncryptionMapInput is an input type that accepts S3BucketServerSideEncryptionMap and S3BucketServerSideEncryptionMapOutput values.
// You can construct a concrete instance of `S3BucketServerSideEncryptionMapInput` via:
//
//	S3BucketServerSideEncryptionMap{ "key": S3BucketServerSideEncryptionArgs{...} }
type S3BucketServerSideEncryptionMapInput interface {
	pulumi.Input

	ToS3BucketServerSideEncryptionMapOutput() S3BucketServerSideEncryptionMapOutput
	ToS3BucketServerSideEncryptionMapOutputWithContext(context.Context) S3BucketServerSideEncryptionMapOutput
}

type S3BucketServerSideEncryptionMap map[string]S3BucketServerSideEncryptionInput

func (S3BucketServerSideEncryptionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*S3BucketServerSideEncryption)(nil)).Elem()
}

func (i S3BucketServerSideEncryptionMap) ToS3BucketServerSideEncryptionMapOutput() S3BucketServerSideEncryptionMapOutput {
	return i.ToS3BucketServerSideEncryptionMapOutputWithContext(context.Background())
}

func (i S3BucketServerSideEncryptionMap) ToS3BucketServerSideEncryptionMapOutputWithContext(ctx context.Context) S3BucketServerSideEncryptionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(S3BucketServerSideEncryptionMapOutput)
}

func (i S3BucketServerSideEncryptionMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*S3BucketServerSideEncryption] {
	return pulumix.Output[map[string]*S3BucketServerSideEncryption]{
		OutputState: i.ToS3BucketServerSideEncryptionMapOutputWithContext(ctx).OutputState,
	}
}

type S3BucketServerSideEncryptionOutput struct{ *pulumi.OutputState }

func (S3BucketServerSideEncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**S3BucketServerSideEncryption)(nil)).Elem()
}

func (o S3BucketServerSideEncryptionOutput) ToS3BucketServerSideEncryptionOutput() S3BucketServerSideEncryptionOutput {
	return o
}

func (o S3BucketServerSideEncryptionOutput) ToS3BucketServerSideEncryptionOutputWithContext(ctx context.Context) S3BucketServerSideEncryptionOutput {
	return o
}

func (o S3BucketServerSideEncryptionOutput) ToOutput(ctx context.Context) pulumix.Output[*S3BucketServerSideEncryption] {
	return pulumix.Output[*S3BucketServerSideEncryption]{
		OutputState: o.OutputState,
	}
}

func (o S3BucketServerSideEncryptionOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *S3BucketServerSideEncryption) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

func (o S3BucketServerSideEncryptionOutput) EncryptionType() pulumi.StringOutput {
	return o.ApplyT(func(v *S3BucketServerSideEncryption) pulumi.StringOutput { return v.EncryptionType }).(pulumi.StringOutput)
}

func (o S3BucketServerSideEncryptionOutput) KmsKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *S3BucketServerSideEncryption) pulumi.StringOutput { return v.KmsKeyId }).(pulumi.StringOutput)
}

type S3BucketServerSideEncryptionArrayOutput struct{ *pulumi.OutputState }

func (S3BucketServerSideEncryptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*S3BucketServerSideEncryption)(nil)).Elem()
}

func (o S3BucketServerSideEncryptionArrayOutput) ToS3BucketServerSideEncryptionArrayOutput() S3BucketServerSideEncryptionArrayOutput {
	return o
}

func (o S3BucketServerSideEncryptionArrayOutput) ToS3BucketServerSideEncryptionArrayOutputWithContext(ctx context.Context) S3BucketServerSideEncryptionArrayOutput {
	return o
}

func (o S3BucketServerSideEncryptionArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*S3BucketServerSideEncryption] {
	return pulumix.Output[[]*S3BucketServerSideEncryption]{
		OutputState: o.OutputState,
	}
}

func (o S3BucketServerSideEncryptionArrayOutput) Index(i pulumi.IntInput) S3BucketServerSideEncryptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *S3BucketServerSideEncryption {
		return vs[0].([]*S3BucketServerSideEncryption)[vs[1].(int)]
	}).(S3BucketServerSideEncryptionOutput)
}

type S3BucketServerSideEncryptionMapOutput struct{ *pulumi.OutputState }

func (S3BucketServerSideEncryptionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*S3BucketServerSideEncryption)(nil)).Elem()
}

func (o S3BucketServerSideEncryptionMapOutput) ToS3BucketServerSideEncryptionMapOutput() S3BucketServerSideEncryptionMapOutput {
	return o
}

func (o S3BucketServerSideEncryptionMapOutput) ToS3BucketServerSideEncryptionMapOutputWithContext(ctx context.Context) S3BucketServerSideEncryptionMapOutput {
	return o
}

func (o S3BucketServerSideEncryptionMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*S3BucketServerSideEncryption] {
	return pulumix.Output[map[string]*S3BucketServerSideEncryption]{
		OutputState: o.OutputState,
	}
}

func (o S3BucketServerSideEncryptionMapOutput) MapIndex(k pulumi.StringInput) S3BucketServerSideEncryptionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *S3BucketServerSideEncryption {
		return vs[0].(map[string]*S3BucketServerSideEncryption)[vs[1].(string)]
	}).(S3BucketServerSideEncryptionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*S3BucketServerSideEncryptionInput)(nil)).Elem(), &S3BucketServerSideEncryption{})
	pulumi.RegisterInputType(reflect.TypeOf((*S3BucketServerSideEncryptionArrayInput)(nil)).Elem(), S3BucketServerSideEncryptionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*S3BucketServerSideEncryptionMapInput)(nil)).Elem(), S3BucketServerSideEncryptionMap{})
	pulumi.RegisterOutputType(S3BucketServerSideEncryptionOutput{})
	pulumi.RegisterOutputType(S3BucketServerSideEncryptionArrayOutput{})
	pulumi.RegisterOutputType(S3BucketServerSideEncryptionMapOutput{})
}
