// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package minio

import (
	"fmt"
	// embed is used to store bridge-metadata.json in the compiled binary
	_ "embed"
	"path/filepath"
	"unicode"

	"github.com/aminueza/terraform-provider-minio/minio"
	"github.com/pulumi/pulumi-minio/provider/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/x"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "minio"
	// modules:
	mainMod = "index" // the y module
)

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeDataSource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the data source's
// first character.
func makeDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeMember(mod+"/"+fn, res)
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(minio.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "minio",
		Description: "A Pulumi package for creating and managing minio cloud resources.",
		Keywords:    []string{"pulumi", "minio"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/pulumi/pulumi-minio",
		GitHubOrg:   "aminueza",
		Config:      map[string]*tfbridge.SchemaInfo{},
		Resources: map[string]*tfbridge.ResourceInfo{
			"minio_s3_bucket":                   {Tok: makeResource(mainMod, "S3Bucket")},
			"minio_s3_bucket_notification":      {Tok: makeResource(mainMod, "S3BucketNotification")},
			"minio_s3_bucket_policy":            {Tok: makeResource(mainMod, "S3BucketPolicy")},
			"minio_s3_bucket_versioning":        {Tok: makeResource(mainMod, "S3BucketVersioning")},
			"minio_s3_object":                   {Tok: makeResource(mainMod, "S3Object")},
			"minio_iam_group":                   {Tok: makeResource(mainMod, "IamGroup")},
			"minio_iam_group_membership":        {Tok: makeResource(mainMod, "IamGroupMembership")},
			"minio_iam_group_policy":            {Tok: makeResource(mainMod, "IamGroupPolicy")},
			"minio_iam_user":                    {Tok: makeResource(mainMod, "IamUser")},
			"minio_iam_policy":                  {Tok: makeResource(mainMod, "IamPolicy")},
			"minio_iam_user_policy_attachment":  {Tok: makeResource(mainMod, "IamUserPolicyAttachment")},
			"minio_iam_group_policy_attachment": {Tok: makeResource(mainMod, "IamGroupPolicyAttachment")},
			"minio_iam_group_user_attachment":   {Tok: makeResource(mainMod, "IamGroupUserAttachment")},
			"minio_ilm_policy":                  {Tok: makeResource(mainMod, "IlmPolicy")},
			"minio_iam_service_account":         {Tok: makeResource(mainMod, "IamServiceAccount")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"minio_iam_policy_document": {Tok: makeDataSource(mainMod, "getIamPolicyDocument")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0",
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		}, MetadataInfo: tfbridge.NewProviderMetadata(metadata),
	}

	err := x.ComputeDefaults(&prov, x.TokensSingleModule("minio_", mainMod,
		x.MakeStandardToken(mainPkg)))
	contract.AssertNoError(err)
	err = x.AutoAliasing(&prov, prov.GetMetadata())
	contract.AssertNoErrorf(err, "auto aliasing apply failed")
	prov.SetAutonaming(255, "-")

	return prov
}

//go:embed cmd/pulumi-resource-minio/bridge-metadata.json
var metadata []byte
