// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package minio

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
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
//			_, err = minio.NewS3BucketPolicy(ctx, "bucketS3BucketPolicy", &minio.S3BucketPolicyArgs{
//				Bucket: bucketS3Bucket.Bucket,
//				Policy: bucketS3Bucket.Bucket.ApplyT(func(bucket string) (string, error) {
//					return fmt.Sprintf(`{
//	  "Version": "2012-10-17",
//	  "Statement": [
//	    {
//	      "Effect": "Allow",
//	     "Principal": {"AWS": ["*"]},
//	      "Resource": ["arn:aws:s3:::%v"],
//	     "Action": ["s3:ListBucket"]
//	    }
//	  ]
//	}
//
// `, bucket), nil
//
//				}).(pulumi.StringOutput),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type S3BucketPolicy struct {
	pulumi.CustomResourceState

	Bucket pulumi.StringOutput `pulumi:"bucket"`
	Policy pulumi.StringOutput `pulumi:"policy"`
}

// NewS3BucketPolicy registers a new resource with the given unique name, arguments, and options.
func NewS3BucketPolicy(ctx *pulumi.Context,
	name string, args *S3BucketPolicyArgs, opts ...pulumi.ResourceOption) (*S3BucketPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	var resource S3BucketPolicy
	err := ctx.RegisterResource("minio:index/s3BucketPolicy:S3BucketPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetS3BucketPolicy gets an existing S3BucketPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetS3BucketPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *S3BucketPolicyState, opts ...pulumi.ResourceOption) (*S3BucketPolicy, error) {
	var resource S3BucketPolicy
	err := ctx.ReadResource("minio:index/s3BucketPolicy:S3BucketPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering S3BucketPolicy resources.
type s3bucketPolicyState struct {
	Bucket *string `pulumi:"bucket"`
	Policy *string `pulumi:"policy"`
}

type S3BucketPolicyState struct {
	Bucket pulumi.StringPtrInput
	Policy pulumi.StringPtrInput
}

func (S3BucketPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*s3bucketPolicyState)(nil)).Elem()
}

type s3bucketPolicyArgs struct {
	Bucket string `pulumi:"bucket"`
	Policy string `pulumi:"policy"`
}

// The set of arguments for constructing a S3BucketPolicy resource.
type S3BucketPolicyArgs struct {
	Bucket pulumi.StringInput
	Policy pulumi.StringInput
}

func (S3BucketPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*s3bucketPolicyArgs)(nil)).Elem()
}

type S3BucketPolicyInput interface {
	pulumi.Input

	ToS3BucketPolicyOutput() S3BucketPolicyOutput
	ToS3BucketPolicyOutputWithContext(ctx context.Context) S3BucketPolicyOutput
}

func (*S3BucketPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**S3BucketPolicy)(nil)).Elem()
}

func (i *S3BucketPolicy) ToS3BucketPolicyOutput() S3BucketPolicyOutput {
	return i.ToS3BucketPolicyOutputWithContext(context.Background())
}

func (i *S3BucketPolicy) ToS3BucketPolicyOutputWithContext(ctx context.Context) S3BucketPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(S3BucketPolicyOutput)
}

// S3BucketPolicyArrayInput is an input type that accepts S3BucketPolicyArray and S3BucketPolicyArrayOutput values.
// You can construct a concrete instance of `S3BucketPolicyArrayInput` via:
//
//	S3BucketPolicyArray{ S3BucketPolicyArgs{...} }
type S3BucketPolicyArrayInput interface {
	pulumi.Input

	ToS3BucketPolicyArrayOutput() S3BucketPolicyArrayOutput
	ToS3BucketPolicyArrayOutputWithContext(context.Context) S3BucketPolicyArrayOutput
}

type S3BucketPolicyArray []S3BucketPolicyInput

func (S3BucketPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*S3BucketPolicy)(nil)).Elem()
}

func (i S3BucketPolicyArray) ToS3BucketPolicyArrayOutput() S3BucketPolicyArrayOutput {
	return i.ToS3BucketPolicyArrayOutputWithContext(context.Background())
}

func (i S3BucketPolicyArray) ToS3BucketPolicyArrayOutputWithContext(ctx context.Context) S3BucketPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(S3BucketPolicyArrayOutput)
}

// S3BucketPolicyMapInput is an input type that accepts S3BucketPolicyMap and S3BucketPolicyMapOutput values.
// You can construct a concrete instance of `S3BucketPolicyMapInput` via:
//
//	S3BucketPolicyMap{ "key": S3BucketPolicyArgs{...} }
type S3BucketPolicyMapInput interface {
	pulumi.Input

	ToS3BucketPolicyMapOutput() S3BucketPolicyMapOutput
	ToS3BucketPolicyMapOutputWithContext(context.Context) S3BucketPolicyMapOutput
}

type S3BucketPolicyMap map[string]S3BucketPolicyInput

func (S3BucketPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*S3BucketPolicy)(nil)).Elem()
}

func (i S3BucketPolicyMap) ToS3BucketPolicyMapOutput() S3BucketPolicyMapOutput {
	return i.ToS3BucketPolicyMapOutputWithContext(context.Background())
}

func (i S3BucketPolicyMap) ToS3BucketPolicyMapOutputWithContext(ctx context.Context) S3BucketPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(S3BucketPolicyMapOutput)
}

type S3BucketPolicyOutput struct{ *pulumi.OutputState }

func (S3BucketPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**S3BucketPolicy)(nil)).Elem()
}

func (o S3BucketPolicyOutput) ToS3BucketPolicyOutput() S3BucketPolicyOutput {
	return o
}

func (o S3BucketPolicyOutput) ToS3BucketPolicyOutputWithContext(ctx context.Context) S3BucketPolicyOutput {
	return o
}

func (o S3BucketPolicyOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *S3BucketPolicy) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

func (o S3BucketPolicyOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v *S3BucketPolicy) pulumi.StringOutput { return v.Policy }).(pulumi.StringOutput)
}

type S3BucketPolicyArrayOutput struct{ *pulumi.OutputState }

func (S3BucketPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*S3BucketPolicy)(nil)).Elem()
}

func (o S3BucketPolicyArrayOutput) ToS3BucketPolicyArrayOutput() S3BucketPolicyArrayOutput {
	return o
}

func (o S3BucketPolicyArrayOutput) ToS3BucketPolicyArrayOutputWithContext(ctx context.Context) S3BucketPolicyArrayOutput {
	return o
}

func (o S3BucketPolicyArrayOutput) Index(i pulumi.IntInput) S3BucketPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *S3BucketPolicy {
		return vs[0].([]*S3BucketPolicy)[vs[1].(int)]
	}).(S3BucketPolicyOutput)
}

type S3BucketPolicyMapOutput struct{ *pulumi.OutputState }

func (S3BucketPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*S3BucketPolicy)(nil)).Elem()
}

func (o S3BucketPolicyMapOutput) ToS3BucketPolicyMapOutput() S3BucketPolicyMapOutput {
	return o
}

func (o S3BucketPolicyMapOutput) ToS3BucketPolicyMapOutputWithContext(ctx context.Context) S3BucketPolicyMapOutput {
	return o
}

func (o S3BucketPolicyMapOutput) MapIndex(k pulumi.StringInput) S3BucketPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *S3BucketPolicy {
		return vs[0].(map[string]*S3BucketPolicy)[vs[1].(string)]
	}).(S3BucketPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*S3BucketPolicyInput)(nil)).Elem(), &S3BucketPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*S3BucketPolicyArrayInput)(nil)).Elem(), S3BucketPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*S3BucketPolicyMapInput)(nil)).Elem(), S3BucketPolicyMap{})
	pulumi.RegisterOutputType(S3BucketPolicyOutput{})
	pulumi.RegisterOutputType(S3BucketPolicyArrayOutput{})
	pulumi.RegisterOutputType(S3BucketPolicyMapOutput{})
}
