// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package minio

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-minio/sdk/go/minio/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			testPolicy, err := minio.NewIamPolicy(ctx, "testPolicy", &minio.IamPolicyArgs{
//				Policy: pulumi.String(`{
//	  "Version":"2012-10-17",
//	  "Statement": [
//	    {
//	      "Sid":"ListAllBucket",
//	      "Effect": "Allow",
//	      "Action": ["s3:PutObject"],
//	      "Principal":"*",
//	      "Resource": "arn:aws:s3:::state-terraform-s3/*"
//	    }
//	  ]
//	}
//
// `),
//
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("minioId", testPolicy.ID())
//			ctx.Export("minioPolicy", testPolicy.Policy)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
type IamPolicy struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput    `pulumi:"name"`
	NamePrefix pulumi.StringPtrOutput `pulumi:"namePrefix"`
	Policy     pulumi.StringOutput    `pulumi:"policy"`
}

// NewIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewIamPolicy(ctx *pulumi.Context,
	name string, args *IamPolicyArgs, opts ...pulumi.ResourceOption) (*IamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IamPolicy
	err := ctx.RegisterResource("minio:index/iamPolicy:IamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIamPolicy gets an existing IamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IamPolicyState, opts ...pulumi.ResourceOption) (*IamPolicy, error) {
	var resource IamPolicy
	err := ctx.ReadResource("minio:index/iamPolicy:IamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IamPolicy resources.
type iamPolicyState struct {
	Name       *string `pulumi:"name"`
	NamePrefix *string `pulumi:"namePrefix"`
	Policy     *string `pulumi:"policy"`
}

type IamPolicyState struct {
	Name       pulumi.StringPtrInput
	NamePrefix pulumi.StringPtrInput
	Policy     pulumi.StringPtrInput
}

func (IamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamPolicyState)(nil)).Elem()
}

type iamPolicyArgs struct {
	Name       *string `pulumi:"name"`
	NamePrefix *string `pulumi:"namePrefix"`
	Policy     string  `pulumi:"policy"`
}

// The set of arguments for constructing a IamPolicy resource.
type IamPolicyArgs struct {
	Name       pulumi.StringPtrInput
	NamePrefix pulumi.StringPtrInput
	Policy     pulumi.StringInput
}

func (IamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamPolicyArgs)(nil)).Elem()
}

type IamPolicyInput interface {
	pulumi.Input

	ToIamPolicyOutput() IamPolicyOutput
	ToIamPolicyOutputWithContext(ctx context.Context) IamPolicyOutput
}

func (*IamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**IamPolicy)(nil)).Elem()
}

func (i *IamPolicy) ToIamPolicyOutput() IamPolicyOutput {
	return i.ToIamPolicyOutputWithContext(context.Background())
}

func (i *IamPolicy) ToIamPolicyOutputWithContext(ctx context.Context) IamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamPolicyOutput)
}

// IamPolicyArrayInput is an input type that accepts IamPolicyArray and IamPolicyArrayOutput values.
// You can construct a concrete instance of `IamPolicyArrayInput` via:
//
//	IamPolicyArray{ IamPolicyArgs{...} }
type IamPolicyArrayInput interface {
	pulumi.Input

	ToIamPolicyArrayOutput() IamPolicyArrayOutput
	ToIamPolicyArrayOutputWithContext(context.Context) IamPolicyArrayOutput
}

type IamPolicyArray []IamPolicyInput

func (IamPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamPolicy)(nil)).Elem()
}

func (i IamPolicyArray) ToIamPolicyArrayOutput() IamPolicyArrayOutput {
	return i.ToIamPolicyArrayOutputWithContext(context.Background())
}

func (i IamPolicyArray) ToIamPolicyArrayOutputWithContext(ctx context.Context) IamPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamPolicyArrayOutput)
}

// IamPolicyMapInput is an input type that accepts IamPolicyMap and IamPolicyMapOutput values.
// You can construct a concrete instance of `IamPolicyMapInput` via:
//
//	IamPolicyMap{ "key": IamPolicyArgs{...} }
type IamPolicyMapInput interface {
	pulumi.Input

	ToIamPolicyMapOutput() IamPolicyMapOutput
	ToIamPolicyMapOutputWithContext(context.Context) IamPolicyMapOutput
}

type IamPolicyMap map[string]IamPolicyInput

func (IamPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamPolicy)(nil)).Elem()
}

func (i IamPolicyMap) ToIamPolicyMapOutput() IamPolicyMapOutput {
	return i.ToIamPolicyMapOutputWithContext(context.Background())
}

func (i IamPolicyMap) ToIamPolicyMapOutputWithContext(ctx context.Context) IamPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamPolicyMapOutput)
}

type IamPolicyOutput struct{ *pulumi.OutputState }

func (IamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IamPolicy)(nil)).Elem()
}

func (o IamPolicyOutput) ToIamPolicyOutput() IamPolicyOutput {
	return o
}

func (o IamPolicyOutput) ToIamPolicyOutputWithContext(ctx context.Context) IamPolicyOutput {
	return o
}

func (o IamPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IamPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o IamPolicyOutput) NamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IamPolicy) pulumi.StringPtrOutput { return v.NamePrefix }).(pulumi.StringPtrOutput)
}

func (o IamPolicyOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v *IamPolicy) pulumi.StringOutput { return v.Policy }).(pulumi.StringOutput)
}

type IamPolicyArrayOutput struct{ *pulumi.OutputState }

func (IamPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamPolicy)(nil)).Elem()
}

func (o IamPolicyArrayOutput) ToIamPolicyArrayOutput() IamPolicyArrayOutput {
	return o
}

func (o IamPolicyArrayOutput) ToIamPolicyArrayOutputWithContext(ctx context.Context) IamPolicyArrayOutput {
	return o
}

func (o IamPolicyArrayOutput) Index(i pulumi.IntInput) IamPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IamPolicy {
		return vs[0].([]*IamPolicy)[vs[1].(int)]
	}).(IamPolicyOutput)
}

type IamPolicyMapOutput struct{ *pulumi.OutputState }

func (IamPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamPolicy)(nil)).Elem()
}

func (o IamPolicyMapOutput) ToIamPolicyMapOutput() IamPolicyMapOutput {
	return o
}

func (o IamPolicyMapOutput) ToIamPolicyMapOutputWithContext(ctx context.Context) IamPolicyMapOutput {
	return o
}

func (o IamPolicyMapOutput) MapIndex(k pulumi.StringInput) IamPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IamPolicy {
		return vs[0].(map[string]*IamPolicy)[vs[1].(string)]
	}).(IamPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IamPolicyInput)(nil)).Elem(), &IamPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamPolicyArrayInput)(nil)).Elem(), IamPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamPolicyMapInput)(nil)).Elem(), IamPolicyMap{})
	pulumi.RegisterOutputType(IamPolicyOutput{})
	pulumi.RegisterOutputType(IamPolicyArrayOutput{})
	pulumi.RegisterOutputType(IamPolicyMapOutput{})
}
