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
// 		_, err := minio.NewIamGroup(ctx, "developerIamGroup", nil)
// 		if err != nil {
// 			return err
// 		}
// 		userOne, err := minio.NewIamUser(ctx, "userOne", nil)
// 		if err != nil {
// 			return err
// 		}
// 		developerIamGroupUserAttachment, err := minio.NewIamGroupUserAttachment(ctx, "developerIamGroupUserAttachment", &minio.IamGroupUserAttachmentArgs{
// 			GroupName: pulumi.Any(minio_iam_group.Group.Name),
// 			UserName:  userOne.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("minioName", developerIamGroupUserAttachment.ID())
// 		ctx.Export("minioUsers", developerIamGroupUserAttachment.GroupName)
// 		ctx.Export("minioGroup", developerIamGroupUserAttachment.UserName)
// 		return nil
// 	})
// }
// ```
type IamGroupUserAttachment struct {
	pulumi.CustomResourceState

	GroupName pulumi.StringOutput `pulumi:"groupName"`
	UserName  pulumi.StringOutput `pulumi:"userName"`
}

// NewIamGroupUserAttachment registers a new resource with the given unique name, arguments, and options.
func NewIamGroupUserAttachment(ctx *pulumi.Context,
	name string, args *IamGroupUserAttachmentArgs, opts ...pulumi.ResourceOption) (*IamGroupUserAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	var resource IamGroupUserAttachment
	err := ctx.RegisterResource("minio:index/iamGroupUserAttachment:IamGroupUserAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIamGroupUserAttachment gets an existing IamGroupUserAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamGroupUserAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IamGroupUserAttachmentState, opts ...pulumi.ResourceOption) (*IamGroupUserAttachment, error) {
	var resource IamGroupUserAttachment
	err := ctx.ReadResource("minio:index/iamGroupUserAttachment:IamGroupUserAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IamGroupUserAttachment resources.
type iamGroupUserAttachmentState struct {
	GroupName *string `pulumi:"groupName"`
	UserName  *string `pulumi:"userName"`
}

type IamGroupUserAttachmentState struct {
	GroupName pulumi.StringPtrInput
	UserName  pulumi.StringPtrInput
}

func (IamGroupUserAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamGroupUserAttachmentState)(nil)).Elem()
}

type iamGroupUserAttachmentArgs struct {
	GroupName string `pulumi:"groupName"`
	UserName  string `pulumi:"userName"`
}

// The set of arguments for constructing a IamGroupUserAttachment resource.
type IamGroupUserAttachmentArgs struct {
	GroupName pulumi.StringInput
	UserName  pulumi.StringInput
}

func (IamGroupUserAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamGroupUserAttachmentArgs)(nil)).Elem()
}

type IamGroupUserAttachmentInput interface {
	pulumi.Input

	ToIamGroupUserAttachmentOutput() IamGroupUserAttachmentOutput
	ToIamGroupUserAttachmentOutputWithContext(ctx context.Context) IamGroupUserAttachmentOutput
}

func (*IamGroupUserAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((*IamGroupUserAttachment)(nil))
}

func (i *IamGroupUserAttachment) ToIamGroupUserAttachmentOutput() IamGroupUserAttachmentOutput {
	return i.ToIamGroupUserAttachmentOutputWithContext(context.Background())
}

func (i *IamGroupUserAttachment) ToIamGroupUserAttachmentOutputWithContext(ctx context.Context) IamGroupUserAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamGroupUserAttachmentOutput)
}

func (i *IamGroupUserAttachment) ToIamGroupUserAttachmentPtrOutput() IamGroupUserAttachmentPtrOutput {
	return i.ToIamGroupUserAttachmentPtrOutputWithContext(context.Background())
}

func (i *IamGroupUserAttachment) ToIamGroupUserAttachmentPtrOutputWithContext(ctx context.Context) IamGroupUserAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamGroupUserAttachmentPtrOutput)
}

type IamGroupUserAttachmentPtrInput interface {
	pulumi.Input

	ToIamGroupUserAttachmentPtrOutput() IamGroupUserAttachmentPtrOutput
	ToIamGroupUserAttachmentPtrOutputWithContext(ctx context.Context) IamGroupUserAttachmentPtrOutput
}

type iamGroupUserAttachmentPtrType IamGroupUserAttachmentArgs

func (*iamGroupUserAttachmentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IamGroupUserAttachment)(nil))
}

func (i *iamGroupUserAttachmentPtrType) ToIamGroupUserAttachmentPtrOutput() IamGroupUserAttachmentPtrOutput {
	return i.ToIamGroupUserAttachmentPtrOutputWithContext(context.Background())
}

func (i *iamGroupUserAttachmentPtrType) ToIamGroupUserAttachmentPtrOutputWithContext(ctx context.Context) IamGroupUserAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamGroupUserAttachmentPtrOutput)
}

// IamGroupUserAttachmentArrayInput is an input type that accepts IamGroupUserAttachmentArray and IamGroupUserAttachmentArrayOutput values.
// You can construct a concrete instance of `IamGroupUserAttachmentArrayInput` via:
//
//          IamGroupUserAttachmentArray{ IamGroupUserAttachmentArgs{...} }
type IamGroupUserAttachmentArrayInput interface {
	pulumi.Input

	ToIamGroupUserAttachmentArrayOutput() IamGroupUserAttachmentArrayOutput
	ToIamGroupUserAttachmentArrayOutputWithContext(context.Context) IamGroupUserAttachmentArrayOutput
}

type IamGroupUserAttachmentArray []IamGroupUserAttachmentInput

func (IamGroupUserAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamGroupUserAttachment)(nil)).Elem()
}

func (i IamGroupUserAttachmentArray) ToIamGroupUserAttachmentArrayOutput() IamGroupUserAttachmentArrayOutput {
	return i.ToIamGroupUserAttachmentArrayOutputWithContext(context.Background())
}

func (i IamGroupUserAttachmentArray) ToIamGroupUserAttachmentArrayOutputWithContext(ctx context.Context) IamGroupUserAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamGroupUserAttachmentArrayOutput)
}

// IamGroupUserAttachmentMapInput is an input type that accepts IamGroupUserAttachmentMap and IamGroupUserAttachmentMapOutput values.
// You can construct a concrete instance of `IamGroupUserAttachmentMapInput` via:
//
//          IamGroupUserAttachmentMap{ "key": IamGroupUserAttachmentArgs{...} }
type IamGroupUserAttachmentMapInput interface {
	pulumi.Input

	ToIamGroupUserAttachmentMapOutput() IamGroupUserAttachmentMapOutput
	ToIamGroupUserAttachmentMapOutputWithContext(context.Context) IamGroupUserAttachmentMapOutput
}

type IamGroupUserAttachmentMap map[string]IamGroupUserAttachmentInput

func (IamGroupUserAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamGroupUserAttachment)(nil)).Elem()
}

func (i IamGroupUserAttachmentMap) ToIamGroupUserAttachmentMapOutput() IamGroupUserAttachmentMapOutput {
	return i.ToIamGroupUserAttachmentMapOutputWithContext(context.Background())
}

func (i IamGroupUserAttachmentMap) ToIamGroupUserAttachmentMapOutputWithContext(ctx context.Context) IamGroupUserAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamGroupUserAttachmentMapOutput)
}

type IamGroupUserAttachmentOutput struct{ *pulumi.OutputState }

func (IamGroupUserAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IamGroupUserAttachment)(nil))
}

func (o IamGroupUserAttachmentOutput) ToIamGroupUserAttachmentOutput() IamGroupUserAttachmentOutput {
	return o
}

func (o IamGroupUserAttachmentOutput) ToIamGroupUserAttachmentOutputWithContext(ctx context.Context) IamGroupUserAttachmentOutput {
	return o
}

func (o IamGroupUserAttachmentOutput) ToIamGroupUserAttachmentPtrOutput() IamGroupUserAttachmentPtrOutput {
	return o.ToIamGroupUserAttachmentPtrOutputWithContext(context.Background())
}

func (o IamGroupUserAttachmentOutput) ToIamGroupUserAttachmentPtrOutputWithContext(ctx context.Context) IamGroupUserAttachmentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IamGroupUserAttachment) *IamGroupUserAttachment {
		return &v
	}).(IamGroupUserAttachmentPtrOutput)
}

type IamGroupUserAttachmentPtrOutput struct{ *pulumi.OutputState }

func (IamGroupUserAttachmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IamGroupUserAttachment)(nil))
}

func (o IamGroupUserAttachmentPtrOutput) ToIamGroupUserAttachmentPtrOutput() IamGroupUserAttachmentPtrOutput {
	return o
}

func (o IamGroupUserAttachmentPtrOutput) ToIamGroupUserAttachmentPtrOutputWithContext(ctx context.Context) IamGroupUserAttachmentPtrOutput {
	return o
}

func (o IamGroupUserAttachmentPtrOutput) Elem() IamGroupUserAttachmentOutput {
	return o.ApplyT(func(v *IamGroupUserAttachment) IamGroupUserAttachment {
		if v != nil {
			return *v
		}
		var ret IamGroupUserAttachment
		return ret
	}).(IamGroupUserAttachmentOutput)
}

type IamGroupUserAttachmentArrayOutput struct{ *pulumi.OutputState }

func (IamGroupUserAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IamGroupUserAttachment)(nil))
}

func (o IamGroupUserAttachmentArrayOutput) ToIamGroupUserAttachmentArrayOutput() IamGroupUserAttachmentArrayOutput {
	return o
}

func (o IamGroupUserAttachmentArrayOutput) ToIamGroupUserAttachmentArrayOutputWithContext(ctx context.Context) IamGroupUserAttachmentArrayOutput {
	return o
}

func (o IamGroupUserAttachmentArrayOutput) Index(i pulumi.IntInput) IamGroupUserAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IamGroupUserAttachment {
		return vs[0].([]IamGroupUserAttachment)[vs[1].(int)]
	}).(IamGroupUserAttachmentOutput)
}

type IamGroupUserAttachmentMapOutput struct{ *pulumi.OutputState }

func (IamGroupUserAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IamGroupUserAttachment)(nil))
}

func (o IamGroupUserAttachmentMapOutput) ToIamGroupUserAttachmentMapOutput() IamGroupUserAttachmentMapOutput {
	return o
}

func (o IamGroupUserAttachmentMapOutput) ToIamGroupUserAttachmentMapOutputWithContext(ctx context.Context) IamGroupUserAttachmentMapOutput {
	return o
}

func (o IamGroupUserAttachmentMapOutput) MapIndex(k pulumi.StringInput) IamGroupUserAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IamGroupUserAttachment {
		return vs[0].(map[string]IamGroupUserAttachment)[vs[1].(string)]
	}).(IamGroupUserAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IamGroupUserAttachmentInput)(nil)).Elem(), &IamGroupUserAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamGroupUserAttachmentPtrInput)(nil)).Elem(), &IamGroupUserAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamGroupUserAttachmentArrayInput)(nil)).Elem(), IamGroupUserAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamGroupUserAttachmentMapInput)(nil)).Elem(), IamGroupUserAttachmentMap{})
	pulumi.RegisterOutputType(IamGroupUserAttachmentOutput{})
	pulumi.RegisterOutputType(IamGroupUserAttachmentPtrOutput{})
	pulumi.RegisterOutputType(IamGroupUserAttachmentArrayOutput{})
	pulumi.RegisterOutputType(IamGroupUserAttachmentMapOutput{})
}
