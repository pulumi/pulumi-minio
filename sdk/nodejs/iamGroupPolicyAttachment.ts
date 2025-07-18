// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 */
export class IamGroupPolicyAttachment extends pulumi.CustomResource {
    /**
     * Get an existing IamGroupPolicyAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IamGroupPolicyAttachmentState, opts?: pulumi.CustomResourceOptions): IamGroupPolicyAttachment {
        return new IamGroupPolicyAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'minio:index/iamGroupPolicyAttachment:IamGroupPolicyAttachment';

    /**
     * Returns true if the given object is an instance of IamGroupPolicyAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IamGroupPolicyAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IamGroupPolicyAttachment.__pulumiType;
    }

    public readonly groupName!: pulumi.Output<string>;
    public readonly policyName!: pulumi.Output<string>;

    /**
     * Create a IamGroupPolicyAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IamGroupPolicyAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IamGroupPolicyAttachmentArgs | IamGroupPolicyAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IamGroupPolicyAttachmentState | undefined;
            resourceInputs["groupName"] = state ? state.groupName : undefined;
            resourceInputs["policyName"] = state ? state.policyName : undefined;
        } else {
            const args = argsOrState as IamGroupPolicyAttachmentArgs | undefined;
            if ((!args || args.groupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupName'");
            }
            if ((!args || args.policyName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyName'");
            }
            resourceInputs["groupName"] = args ? args.groupName : undefined;
            resourceInputs["policyName"] = args ? args.policyName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IamGroupPolicyAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IamGroupPolicyAttachment resources.
 */
export interface IamGroupPolicyAttachmentState {
    groupName?: pulumi.Input<string>;
    policyName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IamGroupPolicyAttachment resource.
 */
export interface IamGroupPolicyAttachmentArgs {
    groupName: pulumi.Input<string>;
    policyName: pulumi.Input<string>;
}
