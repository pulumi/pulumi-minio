// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package minio

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `IlmPolicy` handles lifecycle settings for a given `S3Bucket`.
//
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
//			bucket, err := minio.NewS3Bucket(ctx, "bucket", &minio.S3BucketArgs{
//				Bucket: pulumi.String("bucket"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = minio.NewIlmPolicy(ctx, "bucket-lifecycle-rules", &minio.IlmPolicyArgs{
//				Bucket: bucket.Bucket,
//				Rules: minio.IlmPolicyRuleArray{
//					&minio.IlmPolicyRuleArgs{
//						Id:         pulumi.String("expire-7d"),
//						Expiration: pulumi.String("7d"),
//					},
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
type IlmPolicy struct {
	pulumi.CustomResourceState

	Bucket pulumi.StringOutput      `pulumi:"bucket"`
	Rules  IlmPolicyRuleArrayOutput `pulumi:"rules"`
}

// NewIlmPolicy registers a new resource with the given unique name, arguments, and options.
func NewIlmPolicy(ctx *pulumi.Context,
	name string, args *IlmPolicyArgs, opts ...pulumi.ResourceOption) (*IlmPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	var resource IlmPolicy
	err := ctx.RegisterResource("minio:index/ilmPolicy:IlmPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIlmPolicy gets an existing IlmPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIlmPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IlmPolicyState, opts ...pulumi.ResourceOption) (*IlmPolicy, error) {
	var resource IlmPolicy
	err := ctx.ReadResource("minio:index/ilmPolicy:IlmPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IlmPolicy resources.
type ilmPolicyState struct {
	Bucket *string         `pulumi:"bucket"`
	Rules  []IlmPolicyRule `pulumi:"rules"`
}

type IlmPolicyState struct {
	Bucket pulumi.StringPtrInput
	Rules  IlmPolicyRuleArrayInput
}

func (IlmPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*ilmPolicyState)(nil)).Elem()
}

type ilmPolicyArgs struct {
	Bucket string          `pulumi:"bucket"`
	Rules  []IlmPolicyRule `pulumi:"rules"`
}

// The set of arguments for constructing a IlmPolicy resource.
type IlmPolicyArgs struct {
	Bucket pulumi.StringInput
	Rules  IlmPolicyRuleArrayInput
}

func (IlmPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ilmPolicyArgs)(nil)).Elem()
}

type IlmPolicyInput interface {
	pulumi.Input

	ToIlmPolicyOutput() IlmPolicyOutput
	ToIlmPolicyOutputWithContext(ctx context.Context) IlmPolicyOutput
}

func (*IlmPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**IlmPolicy)(nil)).Elem()
}

func (i *IlmPolicy) ToIlmPolicyOutput() IlmPolicyOutput {
	return i.ToIlmPolicyOutputWithContext(context.Background())
}

func (i *IlmPolicy) ToIlmPolicyOutputWithContext(ctx context.Context) IlmPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IlmPolicyOutput)
}

// IlmPolicyArrayInput is an input type that accepts IlmPolicyArray and IlmPolicyArrayOutput values.
// You can construct a concrete instance of `IlmPolicyArrayInput` via:
//
//	IlmPolicyArray{ IlmPolicyArgs{...} }
type IlmPolicyArrayInput interface {
	pulumi.Input

	ToIlmPolicyArrayOutput() IlmPolicyArrayOutput
	ToIlmPolicyArrayOutputWithContext(context.Context) IlmPolicyArrayOutput
}

type IlmPolicyArray []IlmPolicyInput

func (IlmPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IlmPolicy)(nil)).Elem()
}

func (i IlmPolicyArray) ToIlmPolicyArrayOutput() IlmPolicyArrayOutput {
	return i.ToIlmPolicyArrayOutputWithContext(context.Background())
}

func (i IlmPolicyArray) ToIlmPolicyArrayOutputWithContext(ctx context.Context) IlmPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IlmPolicyArrayOutput)
}

// IlmPolicyMapInput is an input type that accepts IlmPolicyMap and IlmPolicyMapOutput values.
// You can construct a concrete instance of `IlmPolicyMapInput` via:
//
//	IlmPolicyMap{ "key": IlmPolicyArgs{...} }
type IlmPolicyMapInput interface {
	pulumi.Input

	ToIlmPolicyMapOutput() IlmPolicyMapOutput
	ToIlmPolicyMapOutputWithContext(context.Context) IlmPolicyMapOutput
}

type IlmPolicyMap map[string]IlmPolicyInput

func (IlmPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IlmPolicy)(nil)).Elem()
}

func (i IlmPolicyMap) ToIlmPolicyMapOutput() IlmPolicyMapOutput {
	return i.ToIlmPolicyMapOutputWithContext(context.Background())
}

func (i IlmPolicyMap) ToIlmPolicyMapOutputWithContext(ctx context.Context) IlmPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IlmPolicyMapOutput)
}

type IlmPolicyOutput struct{ *pulumi.OutputState }

func (IlmPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IlmPolicy)(nil)).Elem()
}

func (o IlmPolicyOutput) ToIlmPolicyOutput() IlmPolicyOutput {
	return o
}

func (o IlmPolicyOutput) ToIlmPolicyOutputWithContext(ctx context.Context) IlmPolicyOutput {
	return o
}

func (o IlmPolicyOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *IlmPolicy) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

func (o IlmPolicyOutput) Rules() IlmPolicyRuleArrayOutput {
	return o.ApplyT(func(v *IlmPolicy) IlmPolicyRuleArrayOutput { return v.Rules }).(IlmPolicyRuleArrayOutput)
}

type IlmPolicyArrayOutput struct{ *pulumi.OutputState }

func (IlmPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IlmPolicy)(nil)).Elem()
}

func (o IlmPolicyArrayOutput) ToIlmPolicyArrayOutput() IlmPolicyArrayOutput {
	return o
}

func (o IlmPolicyArrayOutput) ToIlmPolicyArrayOutputWithContext(ctx context.Context) IlmPolicyArrayOutput {
	return o
}

func (o IlmPolicyArrayOutput) Index(i pulumi.IntInput) IlmPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IlmPolicy {
		return vs[0].([]*IlmPolicy)[vs[1].(int)]
	}).(IlmPolicyOutput)
}

type IlmPolicyMapOutput struct{ *pulumi.OutputState }

func (IlmPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IlmPolicy)(nil)).Elem()
}

func (o IlmPolicyMapOutput) ToIlmPolicyMapOutput() IlmPolicyMapOutput {
	return o
}

func (o IlmPolicyMapOutput) ToIlmPolicyMapOutputWithContext(ctx context.Context) IlmPolicyMapOutput {
	return o
}

func (o IlmPolicyMapOutput) MapIndex(k pulumi.StringInput) IlmPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IlmPolicy {
		return vs[0].(map[string]*IlmPolicy)[vs[1].(string)]
	}).(IlmPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IlmPolicyInput)(nil)).Elem(), &IlmPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*IlmPolicyArrayInput)(nil)).Elem(), IlmPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IlmPolicyMapInput)(nil)).Elem(), IlmPolicyMap{})
	pulumi.RegisterOutputType(IlmPolicyOutput{})
	pulumi.RegisterOutputType(IlmPolicyArrayOutput{})
	pulumi.RegisterOutputType(IlmPolicyMapOutput{})
}
