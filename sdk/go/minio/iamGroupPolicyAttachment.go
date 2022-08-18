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
//	"fmt"
//
//	"github.com/pulumi/pulumi-minio/sdk/go/minio"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := minio.NewIamGroup(ctx, "developerIamGroup", nil)
//			if err != nil {
//				return err
//			}
//			_, err = minio.NewIamGroupPolicy(ctx, "testPolicy", &minio.IamGroupPolicyArgs{
//				Policy: pulumi.String(fmt.Sprintf(`{
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
// `)),
//
//			})
//			if err != nil {
//				return err
//			}
//			developerIamGroupPolicyAttachment, err := minio.NewIamGroupPolicyAttachment(ctx, "developerIamGroupPolicyAttachment", &minio.IamGroupPolicyAttachmentArgs{
//				GroupName:  pulumi.Any(minio_iam_group.Group.Name),
//				PolicyName: pulumi.Any(minio_iam_policy.Test_policy.Id),
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("minioName", developerIamGroupPolicyAttachment.ID())
//			ctx.Export("minioUsers", developerIamGroupPolicyAttachment.GroupName)
//			ctx.Export("minioGroup", developerIamGroupPolicyAttachment.PolicyName)
//			return nil
//		})
//	}
//
// ```
type IamGroupPolicyAttachment struct {
	pulumi.CustomResourceState

	GroupName  pulumi.StringOutput `pulumi:"groupName"`
	PolicyName pulumi.StringOutput `pulumi:"policyName"`
}

// NewIamGroupPolicyAttachment registers a new resource with the given unique name, arguments, and options.
func NewIamGroupPolicyAttachment(ctx *pulumi.Context,
	name string, args *IamGroupPolicyAttachmentArgs, opts ...pulumi.ResourceOption) (*IamGroupPolicyAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	if args.PolicyName == nil {
		return nil, errors.New("invalid value for required argument 'PolicyName'")
	}
	var resource IamGroupPolicyAttachment
	err := ctx.RegisterResource("minio:index/iamGroupPolicyAttachment:IamGroupPolicyAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIamGroupPolicyAttachment gets an existing IamGroupPolicyAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamGroupPolicyAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IamGroupPolicyAttachmentState, opts ...pulumi.ResourceOption) (*IamGroupPolicyAttachment, error) {
	var resource IamGroupPolicyAttachment
	err := ctx.ReadResource("minio:index/iamGroupPolicyAttachment:IamGroupPolicyAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IamGroupPolicyAttachment resources.
type iamGroupPolicyAttachmentState struct {
	GroupName  *string `pulumi:"groupName"`
	PolicyName *string `pulumi:"policyName"`
}

type IamGroupPolicyAttachmentState struct {
	GroupName  pulumi.StringPtrInput
	PolicyName pulumi.StringPtrInput
}

func (IamGroupPolicyAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamGroupPolicyAttachmentState)(nil)).Elem()
}

type iamGroupPolicyAttachmentArgs struct {
	GroupName  string `pulumi:"groupName"`
	PolicyName string `pulumi:"policyName"`
}

// The set of arguments for constructing a IamGroupPolicyAttachment resource.
type IamGroupPolicyAttachmentArgs struct {
	GroupName  pulumi.StringInput
	PolicyName pulumi.StringInput
}

func (IamGroupPolicyAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamGroupPolicyAttachmentArgs)(nil)).Elem()
}

type IamGroupPolicyAttachmentInput interface {
	pulumi.Input

	ToIamGroupPolicyAttachmentOutput() IamGroupPolicyAttachmentOutput
	ToIamGroupPolicyAttachmentOutputWithContext(ctx context.Context) IamGroupPolicyAttachmentOutput
}

func (*IamGroupPolicyAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**IamGroupPolicyAttachment)(nil)).Elem()
}

func (i *IamGroupPolicyAttachment) ToIamGroupPolicyAttachmentOutput() IamGroupPolicyAttachmentOutput {
	return i.ToIamGroupPolicyAttachmentOutputWithContext(context.Background())
}

func (i *IamGroupPolicyAttachment) ToIamGroupPolicyAttachmentOutputWithContext(ctx context.Context) IamGroupPolicyAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamGroupPolicyAttachmentOutput)
}

// IamGroupPolicyAttachmentArrayInput is an input type that accepts IamGroupPolicyAttachmentArray and IamGroupPolicyAttachmentArrayOutput values.
// You can construct a concrete instance of `IamGroupPolicyAttachmentArrayInput` via:
//
//	IamGroupPolicyAttachmentArray{ IamGroupPolicyAttachmentArgs{...} }
type IamGroupPolicyAttachmentArrayInput interface {
	pulumi.Input

	ToIamGroupPolicyAttachmentArrayOutput() IamGroupPolicyAttachmentArrayOutput
	ToIamGroupPolicyAttachmentArrayOutputWithContext(context.Context) IamGroupPolicyAttachmentArrayOutput
}

type IamGroupPolicyAttachmentArray []IamGroupPolicyAttachmentInput

func (IamGroupPolicyAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamGroupPolicyAttachment)(nil)).Elem()
}

func (i IamGroupPolicyAttachmentArray) ToIamGroupPolicyAttachmentArrayOutput() IamGroupPolicyAttachmentArrayOutput {
	return i.ToIamGroupPolicyAttachmentArrayOutputWithContext(context.Background())
}

func (i IamGroupPolicyAttachmentArray) ToIamGroupPolicyAttachmentArrayOutputWithContext(ctx context.Context) IamGroupPolicyAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamGroupPolicyAttachmentArrayOutput)
}

// IamGroupPolicyAttachmentMapInput is an input type that accepts IamGroupPolicyAttachmentMap and IamGroupPolicyAttachmentMapOutput values.
// You can construct a concrete instance of `IamGroupPolicyAttachmentMapInput` via:
//
//	IamGroupPolicyAttachmentMap{ "key": IamGroupPolicyAttachmentArgs{...} }
type IamGroupPolicyAttachmentMapInput interface {
	pulumi.Input

	ToIamGroupPolicyAttachmentMapOutput() IamGroupPolicyAttachmentMapOutput
	ToIamGroupPolicyAttachmentMapOutputWithContext(context.Context) IamGroupPolicyAttachmentMapOutput
}

type IamGroupPolicyAttachmentMap map[string]IamGroupPolicyAttachmentInput

func (IamGroupPolicyAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamGroupPolicyAttachment)(nil)).Elem()
}

func (i IamGroupPolicyAttachmentMap) ToIamGroupPolicyAttachmentMapOutput() IamGroupPolicyAttachmentMapOutput {
	return i.ToIamGroupPolicyAttachmentMapOutputWithContext(context.Background())
}

func (i IamGroupPolicyAttachmentMap) ToIamGroupPolicyAttachmentMapOutputWithContext(ctx context.Context) IamGroupPolicyAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamGroupPolicyAttachmentMapOutput)
}

type IamGroupPolicyAttachmentOutput struct{ *pulumi.OutputState }

func (IamGroupPolicyAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IamGroupPolicyAttachment)(nil)).Elem()
}

func (o IamGroupPolicyAttachmentOutput) ToIamGroupPolicyAttachmentOutput() IamGroupPolicyAttachmentOutput {
	return o
}

func (o IamGroupPolicyAttachmentOutput) ToIamGroupPolicyAttachmentOutputWithContext(ctx context.Context) IamGroupPolicyAttachmentOutput {
	return o
}

func (o IamGroupPolicyAttachmentOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *IamGroupPolicyAttachment) pulumi.StringOutput { return v.GroupName }).(pulumi.StringOutput)
}

func (o IamGroupPolicyAttachmentOutput) PolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *IamGroupPolicyAttachment) pulumi.StringOutput { return v.PolicyName }).(pulumi.StringOutput)
}

type IamGroupPolicyAttachmentArrayOutput struct{ *pulumi.OutputState }

func (IamGroupPolicyAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamGroupPolicyAttachment)(nil)).Elem()
}

func (o IamGroupPolicyAttachmentArrayOutput) ToIamGroupPolicyAttachmentArrayOutput() IamGroupPolicyAttachmentArrayOutput {
	return o
}

func (o IamGroupPolicyAttachmentArrayOutput) ToIamGroupPolicyAttachmentArrayOutputWithContext(ctx context.Context) IamGroupPolicyAttachmentArrayOutput {
	return o
}

func (o IamGroupPolicyAttachmentArrayOutput) Index(i pulumi.IntInput) IamGroupPolicyAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IamGroupPolicyAttachment {
		return vs[0].([]*IamGroupPolicyAttachment)[vs[1].(int)]
	}).(IamGroupPolicyAttachmentOutput)
}

type IamGroupPolicyAttachmentMapOutput struct{ *pulumi.OutputState }

func (IamGroupPolicyAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamGroupPolicyAttachment)(nil)).Elem()
}

func (o IamGroupPolicyAttachmentMapOutput) ToIamGroupPolicyAttachmentMapOutput() IamGroupPolicyAttachmentMapOutput {
	return o
}

func (o IamGroupPolicyAttachmentMapOutput) ToIamGroupPolicyAttachmentMapOutputWithContext(ctx context.Context) IamGroupPolicyAttachmentMapOutput {
	return o
}

func (o IamGroupPolicyAttachmentMapOutput) MapIndex(k pulumi.StringInput) IamGroupPolicyAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IamGroupPolicyAttachment {
		return vs[0].(map[string]*IamGroupPolicyAttachment)[vs[1].(string)]
	}).(IamGroupPolicyAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IamGroupPolicyAttachmentInput)(nil)).Elem(), &IamGroupPolicyAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamGroupPolicyAttachmentArrayInput)(nil)).Elem(), IamGroupPolicyAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamGroupPolicyAttachmentMapInput)(nil)).Elem(), IamGroupPolicyAttachmentMap{})
	pulumi.RegisterOutputType(IamGroupPolicyAttachmentOutput{})
	pulumi.RegisterOutputType(IamGroupPolicyAttachmentArrayOutput{})
	pulumi.RegisterOutputType(IamGroupPolicyAttachmentMapOutput{})
}
