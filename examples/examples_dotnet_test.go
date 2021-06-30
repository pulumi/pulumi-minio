// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.
// +build dotnet all

package examples

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestAccIamUserCSharp(t *testing.T) {
	test := getCSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "iam-user", "csharp"),
		})

	integration.ProgramTest(t, &test)
}

func getCSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	csharpBase := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"Pulumi.Minio",
		},
	})

	return csharpBase
}
