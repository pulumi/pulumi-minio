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
 * const developerIamGroup = new minio.IamGroup("developerIamGroup", {});
 * const userOne = new minio.IamUser("userOne", {});
 * const developerIamGroupUserAttachment = new minio.IamGroupUserAttachment("developerIamGroupUserAttachment", {
 *     groupName: minio_iam_group.group.name,
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
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IamGroupUserAttachmentState | undefined;
            inputs["groupName"] = state ? state.groupName : undefined;
            inputs["userName"] = state ? state.userName : undefined;
        } else {
            const args = argsOrState as IamGroupUserAttachmentArgs | undefined;
            if ((!args || args.groupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupName'");
            }
            if ((!args || args.userName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userName'");
            }
            inputs["groupName"] = args ? args.groupName : undefined;
            inputs["userName"] = args ? args.userName : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(IamGroupUserAttachment.__pulumiType, name, inputs, opts);
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
