// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class S3BucketServerSideEncryption extends pulumi.CustomResource {
    /**
     * Get an existing S3BucketServerSideEncryption resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: S3BucketServerSideEncryptionState, opts?: pulumi.CustomResourceOptions): S3BucketServerSideEncryption {
        return new S3BucketServerSideEncryption(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'minio:index/s3BucketServerSideEncryption:S3BucketServerSideEncryption';

    /**
     * Returns true if the given object is an instance of S3BucketServerSideEncryption.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is S3BucketServerSideEncryption {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === S3BucketServerSideEncryption.__pulumiType;
    }

    public readonly bucket!: pulumi.Output<string>;
    public readonly encryptionType!: pulumi.Output<string>;
    public readonly kmsKeyId!: pulumi.Output<string>;

    /**
     * Create a S3BucketServerSideEncryption resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: S3BucketServerSideEncryptionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: S3BucketServerSideEncryptionArgs | S3BucketServerSideEncryptionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as S3BucketServerSideEncryptionState | undefined;
            resourceInputs["bucket"] = state ? state.bucket : undefined;
            resourceInputs["encryptionType"] = state ? state.encryptionType : undefined;
            resourceInputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
        } else {
            const args = argsOrState as S3BucketServerSideEncryptionArgs | undefined;
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            if ((!args || args.encryptionType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'encryptionType'");
            }
            if ((!args || args.kmsKeyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kmsKeyId'");
            }
            resourceInputs["bucket"] = args ? args.bucket : undefined;
            resourceInputs["encryptionType"] = args ? args.encryptionType : undefined;
            resourceInputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(S3BucketServerSideEncryption.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering S3BucketServerSideEncryption resources.
 */
export interface S3BucketServerSideEncryptionState {
    bucket?: pulumi.Input<string>;
    encryptionType?: pulumi.Input<string>;
    kmsKeyId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a S3BucketServerSideEncryption resource.
 */
export interface S3BucketServerSideEncryptionArgs {
    bucket: pulumi.Input<string>;
    encryptionType: pulumi.Input<string>;
    kmsKeyId: pulumi.Input<string>;
}