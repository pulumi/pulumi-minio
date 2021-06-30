// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package minio

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetIamPolicyDocument(ctx *pulumi.Context, args *GetIamPolicyDocumentArgs, opts ...pulumi.InvokeOption) (*GetIamPolicyDocumentResult, error) {
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
