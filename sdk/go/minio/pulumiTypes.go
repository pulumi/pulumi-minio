// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package minio

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IlmPolicyRule struct {
	Expiration *string `pulumi:"expiration"`
	Filter     *string `pulumi:"filter"`
	// The ID of this resource.
	Id     string  `pulumi:"id"`
	Status *string `pulumi:"status"`
}

// IlmPolicyRuleInput is an input type that accepts IlmPolicyRuleArgs and IlmPolicyRuleOutput values.
// You can construct a concrete instance of `IlmPolicyRuleInput` via:
//
//	IlmPolicyRuleArgs{...}
type IlmPolicyRuleInput interface {
	pulumi.Input

	ToIlmPolicyRuleOutput() IlmPolicyRuleOutput
	ToIlmPolicyRuleOutputWithContext(context.Context) IlmPolicyRuleOutput
}

type IlmPolicyRuleArgs struct {
	Expiration pulumi.StringPtrInput `pulumi:"expiration"`
	Filter     pulumi.StringPtrInput `pulumi:"filter"`
	// The ID of this resource.
	Id     pulumi.StringInput    `pulumi:"id"`
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (IlmPolicyRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IlmPolicyRule)(nil)).Elem()
}

func (i IlmPolicyRuleArgs) ToIlmPolicyRuleOutput() IlmPolicyRuleOutput {
	return i.ToIlmPolicyRuleOutputWithContext(context.Background())
}

func (i IlmPolicyRuleArgs) ToIlmPolicyRuleOutputWithContext(ctx context.Context) IlmPolicyRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IlmPolicyRuleOutput)
}

// IlmPolicyRuleArrayInput is an input type that accepts IlmPolicyRuleArray and IlmPolicyRuleArrayOutput values.
// You can construct a concrete instance of `IlmPolicyRuleArrayInput` via:
//
//	IlmPolicyRuleArray{ IlmPolicyRuleArgs{...} }
type IlmPolicyRuleArrayInput interface {
	pulumi.Input

	ToIlmPolicyRuleArrayOutput() IlmPolicyRuleArrayOutput
	ToIlmPolicyRuleArrayOutputWithContext(context.Context) IlmPolicyRuleArrayOutput
}

type IlmPolicyRuleArray []IlmPolicyRuleInput

func (IlmPolicyRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IlmPolicyRule)(nil)).Elem()
}

func (i IlmPolicyRuleArray) ToIlmPolicyRuleArrayOutput() IlmPolicyRuleArrayOutput {
	return i.ToIlmPolicyRuleArrayOutputWithContext(context.Background())
}

func (i IlmPolicyRuleArray) ToIlmPolicyRuleArrayOutputWithContext(ctx context.Context) IlmPolicyRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IlmPolicyRuleArrayOutput)
}

type IlmPolicyRuleOutput struct{ *pulumi.OutputState }

func (IlmPolicyRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IlmPolicyRule)(nil)).Elem()
}

func (o IlmPolicyRuleOutput) ToIlmPolicyRuleOutput() IlmPolicyRuleOutput {
	return o
}

func (o IlmPolicyRuleOutput) ToIlmPolicyRuleOutputWithContext(ctx context.Context) IlmPolicyRuleOutput {
	return o
}

func (o IlmPolicyRuleOutput) Expiration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IlmPolicyRule) *string { return v.Expiration }).(pulumi.StringPtrOutput)
}

func (o IlmPolicyRuleOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IlmPolicyRule) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// The ID of this resource.
func (o IlmPolicyRuleOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v IlmPolicyRule) string { return v.Id }).(pulumi.StringOutput)
}

func (o IlmPolicyRuleOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IlmPolicyRule) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type IlmPolicyRuleArrayOutput struct{ *pulumi.OutputState }

func (IlmPolicyRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IlmPolicyRule)(nil)).Elem()
}

func (o IlmPolicyRuleArrayOutput) ToIlmPolicyRuleArrayOutput() IlmPolicyRuleArrayOutput {
	return o
}

func (o IlmPolicyRuleArrayOutput) ToIlmPolicyRuleArrayOutputWithContext(ctx context.Context) IlmPolicyRuleArrayOutput {
	return o
}

func (o IlmPolicyRuleArrayOutput) Index(i pulumi.IntInput) IlmPolicyRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IlmPolicyRule {
		return vs[0].([]IlmPolicyRule)[vs[1].(int)]
	}).(IlmPolicyRuleOutput)
}

type GetIamPolicyDocumentStatement struct {
	Actions    []string                                 `pulumi:"actions"`
	Conditions []GetIamPolicyDocumentStatementCondition `pulumi:"conditions"`
	Effect     *string                                  `pulumi:"effect"`
	Principal  *string                                  `pulumi:"principal"`
	Resources  []string                                 `pulumi:"resources"`
	Sid        *string                                  `pulumi:"sid"`
}

// GetIamPolicyDocumentStatementInput is an input type that accepts GetIamPolicyDocumentStatementArgs and GetIamPolicyDocumentStatementOutput values.
// You can construct a concrete instance of `GetIamPolicyDocumentStatementInput` via:
//
//	GetIamPolicyDocumentStatementArgs{...}
type GetIamPolicyDocumentStatementInput interface {
	pulumi.Input

	ToGetIamPolicyDocumentStatementOutput() GetIamPolicyDocumentStatementOutput
	ToGetIamPolicyDocumentStatementOutputWithContext(context.Context) GetIamPolicyDocumentStatementOutput
}

type GetIamPolicyDocumentStatementArgs struct {
	Actions    pulumi.StringArrayInput                          `pulumi:"actions"`
	Conditions GetIamPolicyDocumentStatementConditionArrayInput `pulumi:"conditions"`
	Effect     pulumi.StringPtrInput                            `pulumi:"effect"`
	Principal  pulumi.StringPtrInput                            `pulumi:"principal"`
	Resources  pulumi.StringArrayInput                          `pulumi:"resources"`
	Sid        pulumi.StringPtrInput                            `pulumi:"sid"`
}

func (GetIamPolicyDocumentStatementArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIamPolicyDocumentStatement)(nil)).Elem()
}

func (i GetIamPolicyDocumentStatementArgs) ToGetIamPolicyDocumentStatementOutput() GetIamPolicyDocumentStatementOutput {
	return i.ToGetIamPolicyDocumentStatementOutputWithContext(context.Background())
}

func (i GetIamPolicyDocumentStatementArgs) ToGetIamPolicyDocumentStatementOutputWithContext(ctx context.Context) GetIamPolicyDocumentStatementOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetIamPolicyDocumentStatementOutput)
}

// GetIamPolicyDocumentStatementArrayInput is an input type that accepts GetIamPolicyDocumentStatementArray and GetIamPolicyDocumentStatementArrayOutput values.
// You can construct a concrete instance of `GetIamPolicyDocumentStatementArrayInput` via:
//
//	GetIamPolicyDocumentStatementArray{ GetIamPolicyDocumentStatementArgs{...} }
type GetIamPolicyDocumentStatementArrayInput interface {
	pulumi.Input

	ToGetIamPolicyDocumentStatementArrayOutput() GetIamPolicyDocumentStatementArrayOutput
	ToGetIamPolicyDocumentStatementArrayOutputWithContext(context.Context) GetIamPolicyDocumentStatementArrayOutput
}

type GetIamPolicyDocumentStatementArray []GetIamPolicyDocumentStatementInput

func (GetIamPolicyDocumentStatementArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetIamPolicyDocumentStatement)(nil)).Elem()
}

func (i GetIamPolicyDocumentStatementArray) ToGetIamPolicyDocumentStatementArrayOutput() GetIamPolicyDocumentStatementArrayOutput {
	return i.ToGetIamPolicyDocumentStatementArrayOutputWithContext(context.Background())
}

func (i GetIamPolicyDocumentStatementArray) ToGetIamPolicyDocumentStatementArrayOutputWithContext(ctx context.Context) GetIamPolicyDocumentStatementArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetIamPolicyDocumentStatementArrayOutput)
}

type GetIamPolicyDocumentStatementOutput struct{ *pulumi.OutputState }

func (GetIamPolicyDocumentStatementOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIamPolicyDocumentStatement)(nil)).Elem()
}

func (o GetIamPolicyDocumentStatementOutput) ToGetIamPolicyDocumentStatementOutput() GetIamPolicyDocumentStatementOutput {
	return o
}

func (o GetIamPolicyDocumentStatementOutput) ToGetIamPolicyDocumentStatementOutputWithContext(ctx context.Context) GetIamPolicyDocumentStatementOutput {
	return o
}

func (o GetIamPolicyDocumentStatementOutput) Actions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIamPolicyDocumentStatement) []string { return v.Actions }).(pulumi.StringArrayOutput)
}

func (o GetIamPolicyDocumentStatementOutput) Conditions() GetIamPolicyDocumentStatementConditionArrayOutput {
	return o.ApplyT(func(v GetIamPolicyDocumentStatement) []GetIamPolicyDocumentStatementCondition { return v.Conditions }).(GetIamPolicyDocumentStatementConditionArrayOutput)
}

func (o GetIamPolicyDocumentStatementOutput) Effect() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIamPolicyDocumentStatement) *string { return v.Effect }).(pulumi.StringPtrOutput)
}

func (o GetIamPolicyDocumentStatementOutput) Principal() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIamPolicyDocumentStatement) *string { return v.Principal }).(pulumi.StringPtrOutput)
}

func (o GetIamPolicyDocumentStatementOutput) Resources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIamPolicyDocumentStatement) []string { return v.Resources }).(pulumi.StringArrayOutput)
}

func (o GetIamPolicyDocumentStatementOutput) Sid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetIamPolicyDocumentStatement) *string { return v.Sid }).(pulumi.StringPtrOutput)
}

type GetIamPolicyDocumentStatementArrayOutput struct{ *pulumi.OutputState }

func (GetIamPolicyDocumentStatementArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetIamPolicyDocumentStatement)(nil)).Elem()
}

func (o GetIamPolicyDocumentStatementArrayOutput) ToGetIamPolicyDocumentStatementArrayOutput() GetIamPolicyDocumentStatementArrayOutput {
	return o
}

func (o GetIamPolicyDocumentStatementArrayOutput) ToGetIamPolicyDocumentStatementArrayOutputWithContext(ctx context.Context) GetIamPolicyDocumentStatementArrayOutput {
	return o
}

func (o GetIamPolicyDocumentStatementArrayOutput) Index(i pulumi.IntInput) GetIamPolicyDocumentStatementOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetIamPolicyDocumentStatement {
		return vs[0].([]GetIamPolicyDocumentStatement)[vs[1].(int)]
	}).(GetIamPolicyDocumentStatementOutput)
}

type GetIamPolicyDocumentStatementCondition struct {
	Test     string   `pulumi:"test"`
	Values   []string `pulumi:"values"`
	Variable string   `pulumi:"variable"`
}

// GetIamPolicyDocumentStatementConditionInput is an input type that accepts GetIamPolicyDocumentStatementConditionArgs and GetIamPolicyDocumentStatementConditionOutput values.
// You can construct a concrete instance of `GetIamPolicyDocumentStatementConditionInput` via:
//
//	GetIamPolicyDocumentStatementConditionArgs{...}
type GetIamPolicyDocumentStatementConditionInput interface {
	pulumi.Input

	ToGetIamPolicyDocumentStatementConditionOutput() GetIamPolicyDocumentStatementConditionOutput
	ToGetIamPolicyDocumentStatementConditionOutputWithContext(context.Context) GetIamPolicyDocumentStatementConditionOutput
}

type GetIamPolicyDocumentStatementConditionArgs struct {
	Test     pulumi.StringInput      `pulumi:"test"`
	Values   pulumi.StringArrayInput `pulumi:"values"`
	Variable pulumi.StringInput      `pulumi:"variable"`
}

func (GetIamPolicyDocumentStatementConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIamPolicyDocumentStatementCondition)(nil)).Elem()
}

func (i GetIamPolicyDocumentStatementConditionArgs) ToGetIamPolicyDocumentStatementConditionOutput() GetIamPolicyDocumentStatementConditionOutput {
	return i.ToGetIamPolicyDocumentStatementConditionOutputWithContext(context.Background())
}

func (i GetIamPolicyDocumentStatementConditionArgs) ToGetIamPolicyDocumentStatementConditionOutputWithContext(ctx context.Context) GetIamPolicyDocumentStatementConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetIamPolicyDocumentStatementConditionOutput)
}

// GetIamPolicyDocumentStatementConditionArrayInput is an input type that accepts GetIamPolicyDocumentStatementConditionArray and GetIamPolicyDocumentStatementConditionArrayOutput values.
// You can construct a concrete instance of `GetIamPolicyDocumentStatementConditionArrayInput` via:
//
//	GetIamPolicyDocumentStatementConditionArray{ GetIamPolicyDocumentStatementConditionArgs{...} }
type GetIamPolicyDocumentStatementConditionArrayInput interface {
	pulumi.Input

	ToGetIamPolicyDocumentStatementConditionArrayOutput() GetIamPolicyDocumentStatementConditionArrayOutput
	ToGetIamPolicyDocumentStatementConditionArrayOutputWithContext(context.Context) GetIamPolicyDocumentStatementConditionArrayOutput
}

type GetIamPolicyDocumentStatementConditionArray []GetIamPolicyDocumentStatementConditionInput

func (GetIamPolicyDocumentStatementConditionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetIamPolicyDocumentStatementCondition)(nil)).Elem()
}

func (i GetIamPolicyDocumentStatementConditionArray) ToGetIamPolicyDocumentStatementConditionArrayOutput() GetIamPolicyDocumentStatementConditionArrayOutput {
	return i.ToGetIamPolicyDocumentStatementConditionArrayOutputWithContext(context.Background())
}

func (i GetIamPolicyDocumentStatementConditionArray) ToGetIamPolicyDocumentStatementConditionArrayOutputWithContext(ctx context.Context) GetIamPolicyDocumentStatementConditionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetIamPolicyDocumentStatementConditionArrayOutput)
}

type GetIamPolicyDocumentStatementConditionOutput struct{ *pulumi.OutputState }

func (GetIamPolicyDocumentStatementConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIamPolicyDocumentStatementCondition)(nil)).Elem()
}

func (o GetIamPolicyDocumentStatementConditionOutput) ToGetIamPolicyDocumentStatementConditionOutput() GetIamPolicyDocumentStatementConditionOutput {
	return o
}

func (o GetIamPolicyDocumentStatementConditionOutput) ToGetIamPolicyDocumentStatementConditionOutputWithContext(ctx context.Context) GetIamPolicyDocumentStatementConditionOutput {
	return o
}

func (o GetIamPolicyDocumentStatementConditionOutput) Test() pulumi.StringOutput {
	return o.ApplyT(func(v GetIamPolicyDocumentStatementCondition) string { return v.Test }).(pulumi.StringOutput)
}

func (o GetIamPolicyDocumentStatementConditionOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIamPolicyDocumentStatementCondition) []string { return v.Values }).(pulumi.StringArrayOutput)
}

func (o GetIamPolicyDocumentStatementConditionOutput) Variable() pulumi.StringOutput {
	return o.ApplyT(func(v GetIamPolicyDocumentStatementCondition) string { return v.Variable }).(pulumi.StringOutput)
}

type GetIamPolicyDocumentStatementConditionArrayOutput struct{ *pulumi.OutputState }

func (GetIamPolicyDocumentStatementConditionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetIamPolicyDocumentStatementCondition)(nil)).Elem()
}

func (o GetIamPolicyDocumentStatementConditionArrayOutput) ToGetIamPolicyDocumentStatementConditionArrayOutput() GetIamPolicyDocumentStatementConditionArrayOutput {
	return o
}

func (o GetIamPolicyDocumentStatementConditionArrayOutput) ToGetIamPolicyDocumentStatementConditionArrayOutputWithContext(ctx context.Context) GetIamPolicyDocumentStatementConditionArrayOutput {
	return o
}

func (o GetIamPolicyDocumentStatementConditionArrayOutput) Index(i pulumi.IntInput) GetIamPolicyDocumentStatementConditionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetIamPolicyDocumentStatementCondition {
		return vs[0].([]GetIamPolicyDocumentStatementCondition)[vs[1].(int)]
	}).(GetIamPolicyDocumentStatementConditionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IlmPolicyRuleInput)(nil)).Elem(), IlmPolicyRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IlmPolicyRuleArrayInput)(nil)).Elem(), IlmPolicyRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetIamPolicyDocumentStatementInput)(nil)).Elem(), GetIamPolicyDocumentStatementArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetIamPolicyDocumentStatementArrayInput)(nil)).Elem(), GetIamPolicyDocumentStatementArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetIamPolicyDocumentStatementConditionInput)(nil)).Elem(), GetIamPolicyDocumentStatementConditionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetIamPolicyDocumentStatementConditionArrayInput)(nil)).Elem(), GetIamPolicyDocumentStatementConditionArray{})
	pulumi.RegisterOutputType(IlmPolicyRuleOutput{})
	pulumi.RegisterOutputType(IlmPolicyRuleArrayOutput{})
	pulumi.RegisterOutputType(GetIamPolicyDocumentStatementOutput{})
	pulumi.RegisterOutputType(GetIamPolicyDocumentStatementArrayOutput{})
	pulumi.RegisterOutputType(GetIamPolicyDocumentStatementConditionOutput{})
	pulumi.RegisterOutputType(GetIamPolicyDocumentStatementConditionArrayOutput{})
}
