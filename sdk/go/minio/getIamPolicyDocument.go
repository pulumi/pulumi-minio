// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package minio

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-minio/sdk/go/minio/internal"
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
//			example, err := minio.GetIamPolicyDocument(ctx, &minio.GetIamPolicyDocumentArgs{
//				Statements: []minio.GetIamPolicyDocumentStatement{
//					{
//						Sid: pulumi.StringRef("1"),
//						Actions: []string{
//							"s3:ListAllMyBuckets",
//							"s3:GetBucketLocation",
//						},
//						Resources: []string{
//							"arn:aws:s3:::*",
//						},
//					},
//					{
//						Actions: []string{
//							"s3:ListBucket",
//						},
//						Resources: []string{
//							"arn:aws:s3:::state-terraform-s3",
//						},
//						Conditions: []minio.GetIamPolicyDocumentStatementCondition{
//							{
//								Test:     "StringLike",
//								Variable: "s3:prefix",
//								Values: []string{
//									"",
//									"home/",
//								},
//							},
//						},
//					},
//					{
//						Actions: []string{
//							"s3:PutObject",
//						},
//						Resources: []string{
//							"arn:aws:s3:::state-terraform-s3",
//							"arn:aws:s3:::state-terraform-s3/*",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = minio.NewIamPolicy(ctx, "test_policy", &minio.IamPolicyArgs{
//				Name:   pulumi.String("state-terraform-s3"),
//				Policy: pulumi.String(example.Json),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetIamPolicyDocument(ctx *pulumi.Context, args *GetIamPolicyDocumentArgs, opts ...pulumi.InvokeOption) (*GetIamPolicyDocumentResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetIamPolicyDocumentResult
	err := ctx.Invoke("minio:index/getIamPolicyDocument:getIamPolicyDocument", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIamPolicyDocument.
type GetIamPolicyDocumentArgs struct {
	OverrideJson *string                         `pulumi:"overrideJson"`
	PolicyId     *string                         `pulumi:"policyId"`
	SourceJson   *string                         `pulumi:"sourceJson"`
	Statements   []GetIamPolicyDocumentStatement `pulumi:"statements"`
	Version      *string                         `pulumi:"version"`
}

// A collection of values returned by getIamPolicyDocument.
type GetIamPolicyDocumentResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id           string                          `pulumi:"id"`
	Json         string                          `pulumi:"json"`
	OverrideJson *string                         `pulumi:"overrideJson"`
	PolicyId     *string                         `pulumi:"policyId"`
	SourceJson   *string                         `pulumi:"sourceJson"`
	Statements   []GetIamPolicyDocumentStatement `pulumi:"statements"`
	Version      *string                         `pulumi:"version"`
}

func GetIamPolicyDocumentOutput(ctx *pulumi.Context, args GetIamPolicyDocumentOutputArgs, opts ...pulumi.InvokeOption) GetIamPolicyDocumentResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetIamPolicyDocumentResultOutput, error) {
			args := v.(GetIamPolicyDocumentArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("minio:index/getIamPolicyDocument:getIamPolicyDocument", args, GetIamPolicyDocumentResultOutput{}, options).(GetIamPolicyDocumentResultOutput), nil
		}).(GetIamPolicyDocumentResultOutput)
}

// A collection of arguments for invoking getIamPolicyDocument.
type GetIamPolicyDocumentOutputArgs struct {
	OverrideJson pulumi.StringPtrInput                   `pulumi:"overrideJson"`
	PolicyId     pulumi.StringPtrInput                   `pulumi:"policyId"`
	SourceJson   pulumi.StringPtrInput                   `pulumi:"sourceJson"`
	Statements   GetIamPolicyDocumentStatementArrayInput `pulumi:"statements"`
	Version      pulumi.StringPtrInput                   `pulumi:"version"`
}

func (GetIamPolicyDocumentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIamPolicyDocumentArgs)(nil)).Elem()
}

// A collection of values returned by getIamPolicyDocument.
type GetIamPolicyDocumentResultOutput struct{ *pulumi.OutputState }

func (GetIamPolicyDocumentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIamPolicyDocumentResult)(nil)).Elem()
}

func (o GetIamPolicyDocumentResultOutput) ToGetIamPolicyDocumentResultOutput() GetIamPolicyDocumentResultOutput {
	return o
}

func (o GetIamPolicyDocumentResultOutput) ToGetIamPolicyDocumentResultOutputWithContext(ctx context.Context) GetIamPolicyDocumentResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetIamPolicyDocumentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIamPolicyDocumentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetIamPolicyDocumentResultOutput) Json() pulumi.StringOutput {
	return o.ApplyT(func(v GetIamPolicyDocumentResult) string { return v.Json }).(pulumi.StringOutput)
}

func (o GetIamPolicyDocumentResultOutput) OverrideJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIamPolicyDocumentResult) *string { return v.OverrideJson }).(pulumi.StringPtrOutput)
}

func (o GetIamPolicyDocumentResultOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIamPolicyDocumentResult) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o GetIamPolicyDocumentResultOutput) SourceJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIamPolicyDocumentResult) *string { return v.SourceJson }).(pulumi.StringPtrOutput)
}

func (o GetIamPolicyDocumentResultOutput) Statements() GetIamPolicyDocumentStatementArrayOutput {
	return o.ApplyT(func(v GetIamPolicyDocumentResult) []GetIamPolicyDocumentStatement { return v.Statements }).(GetIamPolicyDocumentStatementArrayOutput)
}

func (o GetIamPolicyDocumentResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIamPolicyDocumentResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIamPolicyDocumentResultOutput{})
}
