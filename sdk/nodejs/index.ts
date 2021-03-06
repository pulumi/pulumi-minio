// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export * from "./getIamPolicyDocument";
export * from "./iamGroup";
export * from "./iamGroupMembership";
export * from "./iamGroupPolicy";
export * from "./iamGroupPolicyAttachment";
export * from "./iamGroupUserAttachment";
export * from "./iamPolicy";
export * from "./iamUser";
export * from "./iamUserPolicyAttachment";
export * from "./ilmPolicy";
export * from "./provider";
export * from "./s3bucket";
export * from "./s3bucketPolicy";
export * from "./s3object";

// Export sub-modules:
import * as config from "./config";
import * as types from "./types";

export {
    config,
    types,
};

// Import resources to register:
import { IamGroup } from "./iamGroup";
import { IamGroupMembership } from "./iamGroupMembership";
import { IamGroupPolicy } from "./iamGroupPolicy";
import { IamGroupPolicyAttachment } from "./iamGroupPolicyAttachment";
import { IamGroupUserAttachment } from "./iamGroupUserAttachment";
import { IamPolicy } from "./iamPolicy";
import { IamUser } from "./iamUser";
import { IamUserPolicyAttachment } from "./iamUserPolicyAttachment";
import { IlmPolicy } from "./ilmPolicy";
import { S3Bucket } from "./s3bucket";
import { S3BucketPolicy } from "./s3bucketPolicy";
import { S3Object } from "./s3object";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "minio:index/iamGroup:IamGroup":
                return new IamGroup(name, <any>undefined, { urn })
            case "minio:index/iamGroupMembership:IamGroupMembership":
                return new IamGroupMembership(name, <any>undefined, { urn })
            case "minio:index/iamGroupPolicy:IamGroupPolicy":
                return new IamGroupPolicy(name, <any>undefined, { urn })
            case "minio:index/iamGroupPolicyAttachment:IamGroupPolicyAttachment":
                return new IamGroupPolicyAttachment(name, <any>undefined, { urn })
            case "minio:index/iamGroupUserAttachment:IamGroupUserAttachment":
                return new IamGroupUserAttachment(name, <any>undefined, { urn })
            case "minio:index/iamPolicy:IamPolicy":
                return new IamPolicy(name, <any>undefined, { urn })
            case "minio:index/iamUser:IamUser":
                return new IamUser(name, <any>undefined, { urn })
            case "minio:index/iamUserPolicyAttachment:IamUserPolicyAttachment":
                return new IamUserPolicyAttachment(name, <any>undefined, { urn })
            case "minio:index/ilmPolicy:IlmPolicy":
                return new IlmPolicy(name, <any>undefined, { urn })
            case "minio:index/s3Bucket:S3Bucket":
                return new S3Bucket(name, <any>undefined, { urn })
            case "minio:index/s3BucketPolicy:S3BucketPolicy":
                return new S3BucketPolicy(name, <any>undefined, { urn })
            case "minio:index/s3Object:S3Object":
                return new S3Object(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("minio", "index/iamGroup", _module)
pulumi.runtime.registerResourceModule("minio", "index/iamGroupMembership", _module)
pulumi.runtime.registerResourceModule("minio", "index/iamGroupPolicy", _module)
pulumi.runtime.registerResourceModule("minio", "index/iamGroupPolicyAttachment", _module)
pulumi.runtime.registerResourceModule("minio", "index/iamGroupUserAttachment", _module)
pulumi.runtime.registerResourceModule("minio", "index/iamPolicy", _module)
pulumi.runtime.registerResourceModule("minio", "index/iamUser", _module)
pulumi.runtime.registerResourceModule("minio", "index/iamUserPolicyAttachment", _module)
pulumi.runtime.registerResourceModule("minio", "index/ilmPolicy", _module)
pulumi.runtime.registerResourceModule("minio", "index/s3Bucket", _module)
pulumi.runtime.registerResourceModule("minio", "index/s3BucketPolicy", _module)
pulumi.runtime.registerResourceModule("minio", "index/s3Object", _module)

import { Provider } from "./provider";

pulumi.runtime.registerResourcePackage("minio", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:minio") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
