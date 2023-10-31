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
 * const testPolicy = new minio.IamPolicy("testPolicy", {policy: `{
 *   "Version":"2012-10-17",
 *   "Statement": [
 *     {
 *       "Sid":"ListAllBucket",
 *       "Effect": "Allow",
 *       "Action": ["s3:PutObject"],
 *       "Principal":"*",
 *       "Resource": "arn:aws:s3:::state-terraform-s3/*"
 *     }
 *   ]
 * }
 * `});
 * export const minioId = testPolicy.id;
 * export const minioPolicy = testPolicy.policy;
 * ```
 */
export class IamPolicy extends pulumi.CustomResource {
    /**
     * Get an existing IamPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IamPolicyState, opts?: pulumi.CustomResourceOptions): IamPolicy {
        return new IamPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'minio:index/iamPolicy:IamPolicy';

    /**
     * Returns true if the given object is an instance of IamPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IamPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IamPolicy.__pulumiType;
    }

    public readonly name!: pulumi.Output<string>;
    public readonly namePrefix!: pulumi.Output<string | undefined>;
    public readonly policy!: pulumi.Output<string>;

    /**
     * Create a IamPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IamPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IamPolicyArgs | IamPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IamPolicyState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namePrefix"] = state ? state.namePrefix : undefined;
            resourceInputs["policy"] = state ? state.policy : undefined;
        } else {
            const args = argsOrState as IamPolicyArgs | undefined;
            if ((!args || args.policy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policy'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namePrefix"] = args ? args.namePrefix : undefined;
            resourceInputs["policy"] = args ? args.policy : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IamPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IamPolicy resources.
 */
export interface IamPolicyState {
    name?: pulumi.Input<string>;
    namePrefix?: pulumi.Input<string>;
    policy?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IamPolicy resource.
 */
export interface IamPolicyArgs {
    name?: pulumi.Input<string>;
    namePrefix?: pulumi.Input<string>;
    policy: pulumi.Input<string>;
}
