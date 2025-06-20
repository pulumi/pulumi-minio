// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class KmsKey extends pulumi.CustomResource {
    /**
     * Get an existing KmsKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KmsKeyState, opts?: pulumi.CustomResourceOptions): KmsKey {
        return new KmsKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'minio:index/kmsKey:KmsKey';

    /**
     * Returns true if the given object is an instance of KmsKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KmsKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KmsKey.__pulumiType;
    }

    public readonly keyId!: pulumi.Output<string>;

    /**
     * Create a KmsKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KmsKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KmsKeyArgs | KmsKeyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KmsKeyState | undefined;
            resourceInputs["keyId"] = state ? state.keyId : undefined;
        } else {
            const args = argsOrState as KmsKeyArgs | undefined;
            if ((!args || args.keyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyId'");
            }
            resourceInputs["keyId"] = args ? args.keyId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(KmsKey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering KmsKey resources.
 */
export interface KmsKeyState {
    keyId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a KmsKey resource.
 */
export interface KmsKeyArgs {
    keyId: pulumi.Input<string>;
}
