// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class IamUser extends pulumi.CustomResource {
    /**
     * Get an existing IamUser resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IamUserState, opts?: pulumi.CustomResourceOptions): IamUser {
        return new IamUser(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'minio:index/iamUser:IamUser';

    /**
     * Returns true if the given object is an instance of IamUser.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IamUser {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IamUser.__pulumiType;
    }

    /**
     * Disable user
     */
    public readonly disableUser!: pulumi.Output<boolean | undefined>;
    /**
     * Delete user even if it has non-Terraform-managed IAM access keys
     */
    public readonly forceDestroy!: pulumi.Output<boolean | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly secret!: pulumi.Output<string>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Rotate Minio User Secret Key
     */
    public readonly updateSecret!: pulumi.Output<boolean | undefined>;

    /**
     * Create a IamUser resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: IamUserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IamUserArgs | IamUserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IamUserState | undefined;
            resourceInputs["disableUser"] = state ? state.disableUser : undefined;
            resourceInputs["forceDestroy"] = state ? state.forceDestroy : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["secret"] = state ? state.secret : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["updateSecret"] = state ? state.updateSecret : undefined;
        } else {
            const args = argsOrState as IamUserArgs | undefined;
            resourceInputs["disableUser"] = args ? args.disableUser : undefined;
            resourceInputs["forceDestroy"] = args ? args.forceDestroy : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["secret"] = args?.secret ? pulumi.secret(args.secret) : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["updateSecret"] = args ? args.updateSecret : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["secret"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(IamUser.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IamUser resources.
 */
export interface IamUserState {
    /**
     * Disable user
     */
    disableUser?: pulumi.Input<boolean>;
    /**
     * Delete user even if it has non-Terraform-managed IAM access keys
     */
    forceDestroy?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    secret?: pulumi.Input<string>;
    status?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Rotate Minio User Secret Key
     */
    updateSecret?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a IamUser resource.
 */
export interface IamUserArgs {
    /**
     * Disable user
     */
    disableUser?: pulumi.Input<boolean>;
    /**
     * Delete user even if it has non-Terraform-managed IAM access keys
     */
    forceDestroy?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    secret?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Rotate Minio User Secret Key
     */
    updateSecret?: pulumi.Input<boolean>;
}
