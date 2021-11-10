[![Actions Status](https://github.com/pulumi/pulumi-minio/workflows/master/badge.svg)](https://github.com/pulumi/pulumi-minio/actions)
[![Slack](http://www.pulumi.com/images/docs/badges/slack.svg)](https://slack.pulumi.com)
[![NPM version](https://badge.fury.io/js/%40pulumi%2Fminio.svg)](https://www.npmjs.com/package/@pulumi/minio)
[![Python version](https://badge.fury.io/py/pulumi-minio.svg)](https://pypi.org/project/pulumi-minio)
[![NuGet version](https://badge.fury.io/nu/pulumi.minio.svg)](https://badge.fury.io/nu/pulumi.minio)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pulumi/pulumi-minio/sdk/go)](https://pkg.go.dev/github.com/pulumi/pulumi-minio/sdk/go)
[![License](https://img.shields.io/npm/l/%40pulumi%2Fpulumi.svg)](https://github.com/pulumi/pulumi-minio/blob/master/LICENSE)

# Minio Resource Provider

The Minio Resource Provider lets you manage releases in a Minio installation.

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/minio

or `yarn`:

    $ yarn add @pulumi/minio

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_minio

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-minio/sdk

### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.Minio

## Configuration

The following configuration points are available:

* `minio:minioServer` - (Required) Minio Host and Port. It must be provided, but
  it can also be sourced from the `MINIO_ENDPOINT` environment variable
* `minio:minioAccessKey` - (Required) Minio Access Key. It must be provided, but
  it can also be sourced from the `MINIO_ACCESS_KEY` environment variable
* `minio:minioSecretKey` - (Required) Minio Secret Key. It must be provided, but
  it can also be sourced from the `MINIO_SECRET_KEY` environment variable
* `minio:minioRegion` - (Optional) Minio Region (`default: us-east-1`).
* `minio:minioApiVersion` - (Optional) Minio API Version (type: string, options: `v2` or `v4`, default: `v4`).
* `minio:minioSsl` - (Optional) Minio SSL enabled (default: `false`). It can also be sourced from the
  `MINIO_ENABLE_HTTPS` environment variable


## Reference

For further information, please visit [the Minio provider docs](https://www.pulumi.com/docs/intro/cloud-providers/minio)
or for detailed reference documentation, please visit [the API docs](https://www.pulumi.com/docs/reference/pkg/minio).
