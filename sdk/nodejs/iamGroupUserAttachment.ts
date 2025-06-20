// *** WARNING: this file was generated by pulumi-language-nodejs. ***
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
 * const developer = new minio.IamGroup("developer", {name: "developer"});
 * const userOne = new minio.IamUser("user_one", {name: "test-user"});
 * const developerIamGroupUserAttachment = new minio.IamGroupUserAttachment("developer", {
 *     groupName: group.name,
 *     userName: userOne.name,
 * });
 * export const minioName = developerIamGroupUserAttachment.id;
 * export const minioUsers = developerIamGroupUserAttachment.groupName;
 * export const minioGroup = developerIamGroupUserAttachment.userName;
 * ```
 */
export class IamGroupUserAttachment extends pulumi.CustomResource {
    /**
     * Get an existing IamGroupUserAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IamGroupUserAttachmentState, opts?: pulumi.CustomResourceOptions): IamGroupUserAttachment {
        return new IamGroupUserAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'minio:index/iamGroupUserAttachment:IamGroupUserAttachment';

    /**
     * Returns true if the given object is an instance of IamGroupUserAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IamGroupUserAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IamGroupUserAttachment.__pulumiType;
    }

    public readonly groupName!: pulumi.Output<string>;
    public readonly userName!: pulumi.Output<string>;

    /**
     * Create a IamGroupUserAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IamGroupUserAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IamGroupUserAttachmentArgs | IamGroupUserAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IamGroupUserAttachmentState | undefined;
            resourceInputs["groupName"] = state ? state.groupName : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
        } else {
            const args = argsOrState as IamGroupUserAttachmentArgs | undefined;
            if ((!args || args.groupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupName'");
            }
            if ((!args || args.userName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userName'");
            }
            resourceInputs["groupName"] = args ? args.groupName : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IamGroupUserAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IamGroupUserAttachment resources.
 */
export interface IamGroupUserAttachmentState {
    groupName?: pulumi.Input<string>;
    userName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IamGroupUserAttachment resource.
 */
export interface IamGroupUserAttachmentArgs {
    groupName: pulumi.Input<string>;
    userName: pulumi.Input<string>;
}
