package main

import (
	"github.com/pulumi/pulumi-minio/sdk/go/minio"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		user, err := minio.NewIamUser(ctx, "go-user", nil)
		if err != nil {
			return err
		}

		ctx.Export("username", user.Name)

		return nil
	})
}
