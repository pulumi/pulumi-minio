// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as minio from "@pulumi/minio";
 *
 * const bucketS3Bucket = new minio.S3Bucket("bucketS3Bucket", {bucket: "example-bucket"});
 * const bucketS3BucketPolicy = new minio.S3BucketPolicy("bucketS3BucketPolicy", {
 *     bucket: bucketS3Bucket.bucket,
 *     policy: pulumi.interpolate`{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Effect": "Allow",
 *      "Principal": {"AWS": ["*"]},
 *       "Resource": ["arn:aws:s3:::${bucketS3Bucket.bucket}"],
 *      "Action": ["s3:ListBucket"]
 *     }
 *   ]
 * }
 * `,
 * });
 * ```
 */
export class S3BucketPolicy extends pulumi.CustomResource {
    /**
     * Get an existing S3BucketPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: S3BucketPolicyState, opts?: pulumi.CustomResourceOptions): S3BucketPolicy {
        return new S3BucketPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'minio:index/s3BucketPolicy:S3BucketPolicy';

    /**
     * Returns true if the given object is an instance of S3BucketPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is S3BucketPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === S3BucketPolicy.__pulumiType;
    }

    public readonly bucket!: pulumi.Output<string>;
    public readonly policy!: pulumi.Output<string>;

    /**
     * Create a S3BucketPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: S3BucketPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: S3BucketPolicyArgs | S3BucketPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as S3BucketPolicyState | undefined;
            resourceInputs["bucket"] = state ? state.bucket : undefined;
            resourceInputs["policy"] = state ? state.policy : undefined;
        } else {
            const args = argsOrState as S3BucketPolicyArgs | undefined;
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            if ((!args || args.policy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policy'");
            }
            resourceInputs["bucket"] = args ? args.bucket : undefined;
            resourceInputs["policy"] = args ? args.policy : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(S3BucketPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering S3BucketPolicy resources.
 */
export interface S3BucketPolicyState {
    bucket?: pulumi.Input<string>;
    policy?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a S3BucketPolicy resource.
 */
export interface S3BucketPolicyArgs {
    bucket: pulumi.Input<string>;
    policy: pulumi.Input<string>;
}