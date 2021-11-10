// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

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
// 	"github.com/pulumi/pulumi-minio/sdk/go/minio"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		stateTerraformS3, err := minio.NewS3Bucket(ctx, "stateTerraformS3", &minio.S3BucketArgs{
// 			Bucket: pulumi.String("state-terraform-s3"),
// 			Acl:    pulumi.String("public"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		txtFile, err := minio.NewS3Object(ctx, "txtFile", &minio.S3ObjectArgs{
// 			BucketName: stateTerraformS3.Bucket,
// 			ObjectName: pulumi.String("text.txt"),
// 			Content:    pulumi.String("Lorem ipsum dolor sit amet."),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			stateTerraformS3,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("minioId", txtFile.ID())
// 		return nil
// 	})
// }
// ```
type S3Object struct {
	pulumi.CustomResourceState

	BucketName    pulumi.StringOutput    `pulumi:"bucketName"`
	Content       pulumi.StringPtrOutput `pulumi:"content"`
	ContentBase64 pulumi.StringPtrOutput `pulumi:"contentBase64"`
	ContentType   pulumi.StringOutput    `pulumi:"contentType"`
	Etag          pulumi.StringOutput    `pulumi:"etag"`
	ObjectName    pulumi.StringOutput    `pulumi:"objectName"`
	Source        pulumi.StringPtrOutput `pulumi:"source"`
	VersionId     pulumi.StringOutput    `pulumi:"versionId"`
}

// NewS3Object registers a new resource with the given unique name, arguments, and options.
func NewS3Object(ctx *pulumi.Context,
	name string, args *S3ObjectArgs, opts ...pulumi.ResourceOption) (*S3Object, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BucketName == nil {
		return nil, errors.New("invalid value for required argument 'BucketName'")
	}
	if args.ObjectName == nil {
		return nil, errors.New("invalid value for required argument 'ObjectName'")
	}
	var resource S3Object
	err := ctx.RegisterResource("minio:index/s3Object:S3Object", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetS3Object gets an existing S3Object resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetS3Object(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *S3ObjectState, opts ...pulumi.ResourceOption) (*S3Object, error) {
	var resource S3Object
	err := ctx.ReadResource("minio:index/s3Object:S3Object", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering S3Object resources.
type s3objectState struct {
	BucketName    *string `pulumi:"bucketName"`
	Content       *string `pulumi:"content"`
	ContentBase64 *string `pulumi:"contentBase64"`
	ContentType   *string `pulumi:"contentType"`
	Etag          *string `pulumi:"etag"`
	ObjectName    *string `pulumi:"objectName"`
	Source        *string `pulumi:"source"`
	VersionId     *string `pulumi:"versionId"`
}

type S3ObjectState struct {
	BucketName    pulumi.StringPtrInput
	Content       pulumi.StringPtrInput
	ContentBase64 pulumi.StringPtrInput
	ContentType   pulumi.StringPtrInput
	Etag          pulumi.StringPtrInput
	ObjectName    pulumi.StringPtrInput
	Source        pulumi.StringPtrInput
	VersionId     pulumi.StringPtrInput
}

func (S3ObjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*s3objectState)(nil)).Elem()
}

type s3objectArgs struct {
	BucketName    string  `pulumi:"bucketName"`
	Content       *string `pulumi:"content"`
	ContentBase64 *string `pulumi:"contentBase64"`
	ContentType   *string `pulumi:"contentType"`
	Etag          *string `pulumi:"etag"`
	ObjectName    string  `pulumi:"objectName"`
	Source        *string `pulumi:"source"`
	VersionId     *string `pulumi:"versionId"`
}

// The set of arguments for constructing a S3Object resource.
type S3ObjectArgs struct {
	BucketName    pulumi.StringInput
	Content       pulumi.StringPtrInput
	ContentBase64 pulumi.StringPtrInput
	ContentType   pulumi.StringPtrInput
	Etag          pulumi.StringPtrInput
	ObjectName    pulumi.StringInput
	Source        pulumi.StringPtrInput
	VersionId     pulumi.StringPtrInput
}

func (S3ObjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*s3objectArgs)(nil)).Elem()
}

type S3ObjectInput interface {
	pulumi.Input

	ToS3ObjectOutput() S3ObjectOutput
	ToS3ObjectOutputWithContext(ctx context.Context) S3ObjectOutput
}

func (*S3Object) ElementType() reflect.Type {
	return reflect.TypeOf((*S3Object)(nil))
}

func (i *S3Object) ToS3ObjectOutput() S3ObjectOutput {
	return i.ToS3ObjectOutputWithContext(context.Background())
}

func (i *S3Object) ToS3ObjectOutputWithContext(ctx context.Context) S3ObjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(S3ObjectOutput)
}

func (i *S3Object) ToS3ObjectPtrOutput() S3ObjectPtrOutput {
	return i.ToS3ObjectPtrOutputWithContext(context.Background())
}

func (i *S3Object) ToS3ObjectPtrOutputWithContext(ctx context.Context) S3ObjectPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(S3ObjectPtrOutput)
}

type S3ObjectPtrInput interface {
	pulumi.Input

	ToS3ObjectPtrOutput() S3ObjectPtrOutput
	ToS3ObjectPtrOutputWithContext(ctx context.Context) S3ObjectPtrOutput
}

type s3objectPtrType S3ObjectArgs

func (*s3objectPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**S3Object)(nil))
}

func (i *s3objectPtrType) ToS3ObjectPtrOutput() S3ObjectPtrOutput {
	return i.ToS3ObjectPtrOutputWithContext(context.Background())
}

func (i *s3objectPtrType) ToS3ObjectPtrOutputWithContext(ctx context.Context) S3ObjectPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(S3ObjectPtrOutput)
}

// S3ObjectArrayInput is an input type that accepts S3ObjectArray and S3ObjectArrayOutput values.
// You can construct a concrete instance of `S3ObjectArrayInput` via:
//
//          S3ObjectArray{ S3ObjectArgs{...} }
type S3ObjectArrayInput interface {
	pulumi.Input

	ToS3ObjectArrayOutput() S3ObjectArrayOutput
	ToS3ObjectArrayOutputWithContext(context.Context) S3ObjectArrayOutput
}

type S3ObjectArray []S3ObjectInput

func (S3ObjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*S3Object)(nil)).Elem()
}

func (i S3ObjectArray) ToS3ObjectArrayOutput() S3ObjectArrayOutput {
	return i.ToS3ObjectArrayOutputWithContext(context.Background())
}

func (i S3ObjectArray) ToS3ObjectArrayOutputWithContext(ctx context.Context) S3ObjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(S3ObjectArrayOutput)
}

// S3ObjectMapInput is an input type that accepts S3ObjectMap and S3ObjectMapOutput values.
// You can construct a concrete instance of `S3ObjectMapInput` via:
//
//          S3ObjectMap{ "key": S3ObjectArgs{...} }
type S3ObjectMapInput interface {
	pulumi.Input

	ToS3ObjectMapOutput() S3ObjectMapOutput
	ToS3ObjectMapOutputWithContext(context.Context) S3ObjectMapOutput
}

type S3ObjectMap map[string]S3ObjectInput

func (S3ObjectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*S3Object)(nil)).Elem()
}

func (i S3ObjectMap) ToS3ObjectMapOutput() S3ObjectMapOutput {
	return i.ToS3ObjectMapOutputWithContext(context.Background())
}

func (i S3ObjectMap) ToS3ObjectMapOutputWithContext(ctx context.Context) S3ObjectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(S3ObjectMapOutput)
}

type S3ObjectOutput struct{ *pulumi.OutputState }

func (S3ObjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*S3Object)(nil))
}

func (o S3ObjectOutput) ToS3ObjectOutput() S3ObjectOutput {
	return o
}

func (o S3ObjectOutput) ToS3ObjectOutputWithContext(ctx context.Context) S3ObjectOutput {
	return o
}

func (o S3ObjectOutput) ToS3ObjectPtrOutput() S3ObjectPtrOutput {
	return o.ToS3ObjectPtrOutputWithContext(context.Background())
}

func (o S3ObjectOutput) ToS3ObjectPtrOutputWithContext(ctx context.Context) S3ObjectPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v S3Object) *S3Object {
		return &v
	}).(S3ObjectPtrOutput)
}

type S3ObjectPtrOutput struct{ *pulumi.OutputState }

func (S3ObjectPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**S3Object)(nil))
}

func (o S3ObjectPtrOutput) ToS3ObjectPtrOutput() S3ObjectPtrOutput {
	return o
}

func (o S3ObjectPtrOutput) ToS3ObjectPtrOutputWithContext(ctx context.Context) S3ObjectPtrOutput {
	return o
}

func (o S3ObjectPtrOutput) Elem() S3ObjectOutput {
	return o.ApplyT(func(v *S3Object) S3Object {
		if v != nil {
			return *v
		}
		var ret S3Object
		return ret
	}).(S3ObjectOutput)
}

type S3ObjectArrayOutput struct{ *pulumi.OutputState }

func (S3ObjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]S3Object)(nil))
}

func (o S3ObjectArrayOutput) ToS3ObjectArrayOutput() S3ObjectArrayOutput {
	return o
}

func (o S3ObjectArrayOutput) ToS3ObjectArrayOutputWithContext(ctx context.Context) S3ObjectArrayOutput {
	return o
}

func (o S3ObjectArrayOutput) Index(i pulumi.IntInput) S3ObjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) S3Object {
		return vs[0].([]S3Object)[vs[1].(int)]
	}).(S3ObjectOutput)
}

type S3ObjectMapOutput struct{ *pulumi.OutputState }

func (S3ObjectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]S3Object)(nil))
}

func (o S3ObjectMapOutput) ToS3ObjectMapOutput() S3ObjectMapOutput {
	return o
}

func (o S3ObjectMapOutput) ToS3ObjectMapOutputWithContext(ctx context.Context) S3ObjectMapOutput {
	return o
}

func (o S3ObjectMapOutput) MapIndex(k pulumi.StringInput) S3ObjectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) S3Object {
		return vs[0].(map[string]S3Object)[vs[1].(string)]
	}).(S3ObjectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*S3ObjectInput)(nil)).Elem(), &S3Object{})
	pulumi.RegisterInputType(reflect.TypeOf((*S3ObjectPtrInput)(nil)).Elem(), &S3Object{})
	pulumi.RegisterInputType(reflect.TypeOf((*S3ObjectArrayInput)(nil)).Elem(), S3ObjectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*S3ObjectMapInput)(nil)).Elem(), S3ObjectMap{})
	pulumi.RegisterOutputType(S3ObjectOutput{})
	pulumi.RegisterOutputType(S3ObjectPtrOutput{})
	pulumi.RegisterOutputType(S3ObjectArrayOutput{})
	pulumi.RegisterOutputType(S3ObjectMapOutput{})
}
