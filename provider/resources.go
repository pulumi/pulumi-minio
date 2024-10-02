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
	"path"

	// embed is used to store bridge-metadata.json in the compiled binary
	_ "embed"

	"github.com/aminueza/terraform-provider-minio/minio"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"

	"github.com/pulumi/pulumi-minio/provider/pkg/version"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "minio"
	// modules:
	mainMod = "index" // the y module
)

//go:embed cmd/pulumi-resource-minio/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(minio.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:            p,
		Name:         "minio",
		Description:  "A Pulumi package for creating and managing minio cloud resources.",
		Keywords:     []string{"pulumi", "minio"},
		License:      "Apache-2.0",
		Homepage:     "https://pulumi.io",
		Repository:   "https://github.com/pulumi/pulumi-minio",
		GitHubOrg:    "aminueza",
		Config:       map[string]*tfbridge.SchemaInfo{},
		Version:      version.Version,
		MetadataInfo: tfbridge.NewProviderMetadata(metadata),

		Resources: map[string]*tfbridge.ResourceInfo{
			"minio_s3_bucket": {Fields: map[string]*tfbridge.SchemaInfo{
				// quota should be a u64 since it describes bytes sizes.
				//
				// The Pulumi schema doesn't support u64 (or i64), so we
				// are converting to a float64, which should get us at
				// least 2^52 bits of precision while minimizing the
				// breaking change (int -> float) for our users.
				"quota": {Type: "number"},
			}},
		},

		JavaScript: &tfbridge.JavaScriptInfo{
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0",
				"@types/mime": "^2.0.0",
			},
			RespectSchemaVersion: true,
		},
		Python: &tfbridge.PythonInfo{
			RespectSchemaVersion: true,

			PyProject: struct{ Enabled bool }{true},
		},

		Golang: &tfbridge.GolangInfo{
			ImportBasePath: path.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
			RespectSchemaVersion:           true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RespectSchemaVersion: true,
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.MustComputeTokens(tokens.SingleModule("minio_", mainMod,
		tokens.MakeStandard(mainPkg)))
	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}
