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
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-minio/sdk/go/minio"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		testUser, err := minio.NewIamUser(ctx, "testUser", nil)
// 		if err != nil {
// 			return err
// 		}
// 		testPolicy, err := minio.NewIamPolicy(ctx, "testPolicy", &minio.IamPolicyArgs{
// 			Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\":\"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Sid\":\"ListAllBucket\",\n", "      \"Effect\": \"Allow\",\n", "      \"Action\": [\"s3:PutObject\"],\n", "      \"Principal\":\"*\",\n", "      \"Resource\": \"arn:aws:s3:::state-terraform-s3/*\"\n", "    }\n", "  ]\n", "}\n", "\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		developer, err := minio.NewIamUserPolicyAttachment(ctx, "developer", &minio.IamUserPolicyAttachmentArgs{
// 			PolicyName: testPolicy.ID(),
// 			UserName:   testUser.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("minioName", developer.ID())
// 		ctx.Export("minioUsers", developer.UserName)
// 		ctx.Export("minioGroup", developer.PolicyName)
// 		return nil
// 	})
// }
// ```
type IamUserPolicyAttachment struct {
	pulumi.CustomResourceState

	PolicyName pulumi.StringOutput `pulumi:"policyName"`
	UserName   pulumi.StringOutput `pulumi:"userName"`
}

// NewIamUserPolicyAttachment registers a new resource with the given unique name, arguments, and options.
func NewIamUserPolicyAttachment(ctx *pulumi.Context,
	name string, args *IamUserPolicyAttachmentArgs, opts ...pulumi.ResourceOption) (*IamUserPolicyAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyName == nil {
		return nil, errors.New("invalid value for required argument 'PolicyName'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	var resource IamUserPolicyAttachment
	err := ctx.RegisterResource("minio:index/iamUserPolicyAttachment:IamUserPolicyAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIamUserPolicyAttachment gets an existing IamUserPolicyAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamUserPolicyAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IamUserPolicyAttachmentState, opts ...pulumi.ResourceOption) (*IamUserPolicyAttachment, error) {
	var resource IamUserPolicyAttachment
	err := ctx.ReadResource("minio:index/iamUserPolicyAttachment:IamUserPolicyAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IamUserPolicyAttachment resources.
type iamUserPolicyAttachmentState struct {
	PolicyName *string `pulumi:"policyName"`
	UserName   *string `pulumi:"userName"`
}

type IamUserPolicyAttachmentState struct {
	PolicyName pulumi.StringPtrInput
	UserName   pulumi.StringPtrInput
}

func (IamUserPolicyAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamUserPolicyAttachmentState)(nil)).Elem()
}

type iamUserPolicyAttachmentArgs struct {
	PolicyName string `pulumi:"policyName"`
	UserName   string `pulumi:"userName"`
}

// The set of arguments for constructing a IamUserPolicyAttachment resource.
type IamUserPolicyAttachmentArgs struct {
	PolicyName pulumi.StringInput
	UserName   pulumi.StringInput
}

func (IamUserPolicyAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamUserPolicyAttachmentArgs)(nil)).Elem()
}

type IamUserPolicyAttachmentInput interface {
	pulumi.Input

	ToIamUserPolicyAttachmentOutput() IamUserPolicyAttachmentOutput
	ToIamUserPolicyAttachmentOutputWithContext(ctx context.Context) IamUserPolicyAttachmentOutput
}

func (*IamUserPolicyAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((*IamUserPolicyAttachment)(nil))
}

func (i *IamUserPolicyAttachment) ToIamUserPolicyAttachmentOutput() IamUserPolicyAttachmentOutput {
	return i.ToIamUserPolicyAttachmentOutputWithContext(context.Background())
}

func (i *IamUserPolicyAttachment) ToIamUserPolicyAttachmentOutputWithContext(ctx context.Context) IamUserPolicyAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamUserPolicyAttachmentOutput)
}

func (i *IamUserPolicyAttachment) ToIamUserPolicyAttachmentPtrOutput() IamUserPolicyAttachmentPtrOutput {
	return i.ToIamUserPolicyAttachmentPtrOutputWithContext(context.Background())
}

func (i *IamUserPolicyAttachment) ToIamUserPolicyAttachmentPtrOutputWithContext(ctx context.Context) IamUserPolicyAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamUserPolicyAttachmentPtrOutput)
}

type IamUserPolicyAttachmentPtrInput interface {
	pulumi.Input

	ToIamUserPolicyAttachmentPtrOutput() IamUserPolicyAttachmentPtrOutput
	ToIamUserPolicyAttachmentPtrOutputWithContext(ctx context.Context) IamUserPolicyAttachmentPtrOutput
}

type iamUserPolicyAttachmentPtrType IamUserPolicyAttachmentArgs

func (*iamUserPolicyAttachmentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IamUserPolicyAttachment)(nil))
}

func (i *iamUserPolicyAttachmentPtrType) ToIamUserPolicyAttachmentPtrOutput() IamUserPolicyAttachmentPtrOutput {
	return i.ToIamUserPolicyAttachmentPtrOutputWithContext(context.Background())
}

func (i *iamUserPolicyAttachmentPtrType) ToIamUserPolicyAttachmentPtrOutputWithContext(ctx context.Context) IamUserPolicyAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamUserPolicyAttachmentPtrOutput)
}

// IamUserPolicyAttachmentArrayInput is an input type that accepts IamUserPolicyAttachmentArray and IamUserPolicyAttachmentArrayOutput values.
// You can construct a concrete instance of `IamUserPolicyAttachmentArrayInput` via:
//
//          IamUserPolicyAttachmentArray{ IamUserPolicyAttachmentArgs{...} }
type IamUserPolicyAttachmentArrayInput interface {
	pulumi.Input

	ToIamUserPolicyAttachmentArrayOutput() IamUserPolicyAttachmentArrayOutput
	ToIamUserPolicyAttachmentArrayOutputWithContext(context.Context) IamUserPolicyAttachmentArrayOutput
}

type IamUserPolicyAttachmentArray []IamUserPolicyAttachmentInput

func (IamUserPolicyAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamUserPolicyAttachment)(nil)).Elem()
}

func (i IamUserPolicyAttachmentArray) ToIamUserPolicyAttachmentArrayOutput() IamUserPolicyAttachmentArrayOutput {
	return i.ToIamUserPolicyAttachmentArrayOutputWithContext(context.Background())
}

func (i IamUserPolicyAttachmentArray) ToIamUserPolicyAttachmentArrayOutputWithContext(ctx context.Context) IamUserPolicyAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamUserPolicyAttachmentArrayOutput)
}

// IamUserPolicyAttachmentMapInput is an input type that accepts IamUserPolicyAttachmentMap and IamUserPolicyAttachmentMapOutput values.
// You can construct a concrete instance of `IamUserPolicyAttachmentMapInput` via:
//
//          IamUserPolicyAttachmentMap{ "key": IamUserPolicyAttachmentArgs{...} }
type IamUserPolicyAttachmentMapInput interface {
	pulumi.Input

	ToIamUserPolicyAttachmentMapOutput() IamUserPolicyAttachmentMapOutput
	ToIamUserPolicyAttachmentMapOutputWithContext(context.Context) IamUserPolicyAttachmentMapOutput
}

type IamUserPolicyAttachmentMap map[string]IamUserPolicyAttachmentInput

func (IamUserPolicyAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamUserPolicyAttachment)(nil)).Elem()
}

func (i IamUserPolicyAttachmentMap) ToIamUserPolicyAttachmentMapOutput() IamUserPolicyAttachmentMapOutput {
	return i.ToIamUserPolicyAttachmentMapOutputWithContext(context.Background())
}

func (i IamUserPolicyAttachmentMap) ToIamUserPolicyAttachmentMapOutputWithContext(ctx context.Context) IamUserPolicyAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamUserPolicyAttachmentMapOutput)
}

type IamUserPolicyAttachmentOutput struct{ *pulumi.OutputState }

func (IamUserPolicyAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IamUserPolicyAttachment)(nil))
}

func (o IamUserPolicyAttachmentOutput) ToIamUserPolicyAttachmentOutput() IamUserPolicyAttachmentOutput {
	return o
}

func (o IamUserPolicyAttachmentOutput) ToIamUserPolicyAttachmentOutputWithContext(ctx context.Context) IamUserPolicyAttachmentOutput {
	return o
}

func (o IamUserPolicyAttachmentOutput) ToIamUserPolicyAttachmentPtrOutput() IamUserPolicyAttachmentPtrOutput {
	return o.ToIamUserPolicyAttachmentPtrOutputWithContext(context.Background())
}

func (o IamUserPolicyAttachmentOutput) ToIamUserPolicyAttachmentPtrOutputWithContext(ctx context.Context) IamUserPolicyAttachmentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IamUserPolicyAttachment) *IamUserPolicyAttachment {
		return &v
	}).(IamUserPolicyAttachmentPtrOutput)
}

type IamUserPolicyAttachmentPtrOutput struct{ *pulumi.OutputState }

func (IamUserPolicyAttachmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IamUserPolicyAttachment)(nil))
}

func (o IamUserPolicyAttachmentPtrOutput) ToIamUserPolicyAttachmentPtrOutput() IamUserPolicyAttachmentPtrOutput {
	return o
}

func (o IamUserPolicyAttachmentPtrOutput) ToIamUserPolicyAttachmentPtrOutputWithContext(ctx context.Context) IamUserPolicyAttachmentPtrOutput {
	return o
}

func (o IamUserPolicyAttachmentPtrOutput) Elem() IamUserPolicyAttachmentOutput {
	return o.ApplyT(func(v *IamUserPolicyAttachment) IamUserPolicyAttachment {
		if v != nil {
			return *v
		}
		var ret IamUserPolicyAttachment
		return ret
	}).(IamUserPolicyAttachmentOutput)
}

type IamUserPolicyAttachmentArrayOutput struct{ *pulumi.OutputState }

func (IamUserPolicyAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IamUserPolicyAttachment)(nil))
}

func (o IamUserPolicyAttachmentArrayOutput) ToIamUserPolicyAttachmentArrayOutput() IamUserPolicyAttachmentArrayOutput {
	return o
}

func (o IamUserPolicyAttachmentArrayOutput) ToIamUserPolicyAttachmentArrayOutputWithContext(ctx context.Context) IamUserPolicyAttachmentArrayOutput {
	return o
}

func (o IamUserPolicyAttachmentArrayOutput) Index(i pulumi.IntInput) IamUserPolicyAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IamUserPolicyAttachment {
		return vs[0].([]IamUserPolicyAttachment)[vs[1].(int)]
	}).(IamUserPolicyAttachmentOutput)
}

type IamUserPolicyAttachmentMapOutput struct{ *pulumi.OutputState }

func (IamUserPolicyAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IamUserPolicyAttachment)(nil))
}

func (o IamUserPolicyAttachmentMapOutput) ToIamUserPolicyAttachmentMapOutput() IamUserPolicyAttachmentMapOutput {
	return o
}

func (o IamUserPolicyAttachmentMapOutput) ToIamUserPolicyAttachmentMapOutputWithContext(ctx context.Context) IamUserPolicyAttachmentMapOutput {
	return o
}

func (o IamUserPolicyAttachmentMapOutput) MapIndex(k pulumi.StringInput) IamUserPolicyAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IamUserPolicyAttachment {
		return vs[0].(map[string]IamUserPolicyAttachment)[vs[1].(string)]
	}).(IamUserPolicyAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IamUserPolicyAttachmentInput)(nil)).Elem(), &IamUserPolicyAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamUserPolicyAttachmentPtrInput)(nil)).Elem(), &IamUserPolicyAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamUserPolicyAttachmentArrayInput)(nil)).Elem(), IamUserPolicyAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamUserPolicyAttachmentMapInput)(nil)).Elem(), IamUserPolicyAttachmentMap{})
	pulumi.RegisterOutputType(IamUserPolicyAttachmentOutput{})
	pulumi.RegisterOutputType(IamUserPolicyAttachmentPtrOutput{})
	pulumi.RegisterOutputType(IamUserPolicyAttachmentArrayOutput{})
	pulumi.RegisterOutputType(IamUserPolicyAttachmentMapOutput{})
}
