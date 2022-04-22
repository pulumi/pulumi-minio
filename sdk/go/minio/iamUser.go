// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package minio

import (
	"context"
	"reflect"

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
// 		testIamUser, err := minio.NewIamUser(ctx, "testIamUser", &minio.IamUserArgs{
// 			ForceDestroy: pulumi.Bool(true),
// 			Tags: pulumi.AnyMap{
// 				"tag-key": pulumi.Any("tag-value"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("test", testIamUser.ID())
// 		ctx.Export("status", testIamUser.Status)
// 		ctx.Export("secret", testIamUser.Secret)
// 		return nil
// 	})
// }
// ```
type IamUser struct {
	pulumi.CustomResourceState

	// Disable user
	DisableUser pulumi.BoolPtrOutput `pulumi:"disableUser"`
	// Delete user even if it has non-Terraform-managed IAM access keys
	ForceDestroy pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	Name         pulumi.StringOutput  `pulumi:"name"`
	Secret       pulumi.StringOutput  `pulumi:"secret"`
	Status       pulumi.StringOutput  `pulumi:"status"`
	Tags         pulumi.MapOutput     `pulumi:"tags"`
	// Rotate Minio User Secret Key
	UpdateSecret pulumi.BoolPtrOutput `pulumi:"updateSecret"`
}

// NewIamUser registers a new resource with the given unique name, arguments, and options.
func NewIamUser(ctx *pulumi.Context,
	name string, args *IamUserArgs, opts ...pulumi.ResourceOption) (*IamUser, error) {
	if args == nil {
		args = &IamUserArgs{}
	}

	var resource IamUser
	err := ctx.RegisterResource("minio:index/iamUser:IamUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIamUser gets an existing IamUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IamUserState, opts ...pulumi.ResourceOption) (*IamUser, error) {
	var resource IamUser
	err := ctx.ReadResource("minio:index/iamUser:IamUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IamUser resources.
type iamUserState struct {
	// Disable user
	DisableUser *bool `pulumi:"disableUser"`
	// Delete user even if it has non-Terraform-managed IAM access keys
	ForceDestroy *bool                  `pulumi:"forceDestroy"`
	Name         *string                `pulumi:"name"`
	Secret       *string                `pulumi:"secret"`
	Status       *string                `pulumi:"status"`
	Tags         map[string]interface{} `pulumi:"tags"`
	// Rotate Minio User Secret Key
	UpdateSecret *bool `pulumi:"updateSecret"`
}

type IamUserState struct {
	// Disable user
	DisableUser pulumi.BoolPtrInput
	// Delete user even if it has non-Terraform-managed IAM access keys
	ForceDestroy pulumi.BoolPtrInput
	Name         pulumi.StringPtrInput
	Secret       pulumi.StringPtrInput
	Status       pulumi.StringPtrInput
	Tags         pulumi.MapInput
	// Rotate Minio User Secret Key
	UpdateSecret pulumi.BoolPtrInput
}

func (IamUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamUserState)(nil)).Elem()
}

type iamUserArgs struct {
	// Disable user
	DisableUser *bool `pulumi:"disableUser"`
	// Delete user even if it has non-Terraform-managed IAM access keys
	ForceDestroy *bool                  `pulumi:"forceDestroy"`
	Name         *string                `pulumi:"name"`
	Secret       *string                `pulumi:"secret"`
	Tags         map[string]interface{} `pulumi:"tags"`
	// Rotate Minio User Secret Key
	UpdateSecret *bool `pulumi:"updateSecret"`
}

// The set of arguments for constructing a IamUser resource.
type IamUserArgs struct {
	// Disable user
	DisableUser pulumi.BoolPtrInput
	// Delete user even if it has non-Terraform-managed IAM access keys
	ForceDestroy pulumi.BoolPtrInput
	Name         pulumi.StringPtrInput
	Secret       pulumi.StringPtrInput
	Tags         pulumi.MapInput
	// Rotate Minio User Secret Key
	UpdateSecret pulumi.BoolPtrInput
}

func (IamUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamUserArgs)(nil)).Elem()
}

type IamUserInput interface {
	pulumi.Input

	ToIamUserOutput() IamUserOutput
	ToIamUserOutputWithContext(ctx context.Context) IamUserOutput
}

func (*IamUser) ElementType() reflect.Type {
	return reflect.TypeOf((**IamUser)(nil)).Elem()
}

func (i *IamUser) ToIamUserOutput() IamUserOutput {
	return i.ToIamUserOutputWithContext(context.Background())
}

func (i *IamUser) ToIamUserOutputWithContext(ctx context.Context) IamUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamUserOutput)
}

// IamUserArrayInput is an input type that accepts IamUserArray and IamUserArrayOutput values.
// You can construct a concrete instance of `IamUserArrayInput` via:
//
//          IamUserArray{ IamUserArgs{...} }
type IamUserArrayInput interface {
	pulumi.Input

	ToIamUserArrayOutput() IamUserArrayOutput
	ToIamUserArrayOutputWithContext(context.Context) IamUserArrayOutput
}

type IamUserArray []IamUserInput

func (IamUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamUser)(nil)).Elem()
}

func (i IamUserArray) ToIamUserArrayOutput() IamUserArrayOutput {
	return i.ToIamUserArrayOutputWithContext(context.Background())
}

func (i IamUserArray) ToIamUserArrayOutputWithContext(ctx context.Context) IamUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamUserArrayOutput)
}

// IamUserMapInput is an input type that accepts IamUserMap and IamUserMapOutput values.
// You can construct a concrete instance of `IamUserMapInput` via:
//
//          IamUserMap{ "key": IamUserArgs{...} }
type IamUserMapInput interface {
	pulumi.Input

	ToIamUserMapOutput() IamUserMapOutput
	ToIamUserMapOutputWithContext(context.Context) IamUserMapOutput
}

type IamUserMap map[string]IamUserInput

func (IamUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamUser)(nil)).Elem()
}

func (i IamUserMap) ToIamUserMapOutput() IamUserMapOutput {
	return i.ToIamUserMapOutputWithContext(context.Background())
}

func (i IamUserMap) ToIamUserMapOutputWithContext(ctx context.Context) IamUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamUserMapOutput)
}

type IamUserOutput struct{ *pulumi.OutputState }

func (IamUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IamUser)(nil)).Elem()
}

func (o IamUserOutput) ToIamUserOutput() IamUserOutput {
	return o
}

func (o IamUserOutput) ToIamUserOutputWithContext(ctx context.Context) IamUserOutput {
	return o
}

type IamUserArrayOutput struct{ *pulumi.OutputState }

func (IamUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamUser)(nil)).Elem()
}

func (o IamUserArrayOutput) ToIamUserArrayOutput() IamUserArrayOutput {
	return o
}

func (o IamUserArrayOutput) ToIamUserArrayOutputWithContext(ctx context.Context) IamUserArrayOutput {
	return o
}

func (o IamUserArrayOutput) Index(i pulumi.IntInput) IamUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IamUser {
		return vs[0].([]*IamUser)[vs[1].(int)]
	}).(IamUserOutput)
}

type IamUserMapOutput struct{ *pulumi.OutputState }

func (IamUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamUser)(nil)).Elem()
}

func (o IamUserMapOutput) ToIamUserMapOutput() IamUserMapOutput {
	return o
}

func (o IamUserMapOutput) ToIamUserMapOutputWithContext(ctx context.Context) IamUserMapOutput {
	return o
}

func (o IamUserMapOutput) MapIndex(k pulumi.StringInput) IamUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IamUser {
		return vs[0].(map[string]*IamUser)[vs[1].(string)]
	}).(IamUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IamUserInput)(nil)).Elem(), &IamUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamUserArrayInput)(nil)).Elem(), IamUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamUserMapInput)(nil)).Elem(), IamUserMap{})
	pulumi.RegisterOutputType(IamUserOutput{})
	pulumi.RegisterOutputType(IamUserArrayOutput{})
	pulumi.RegisterOutputType(IamUserMapOutput{})
}
