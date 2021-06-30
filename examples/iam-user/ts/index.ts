import * as pulumi from "@pulumi/pulumi";
import * as minio from "@pulumi/minio";

const user = new minio.IamUser("ts-user")

export const userName = user.name;
