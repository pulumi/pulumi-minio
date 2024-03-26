// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The provider type for the minio package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'minio';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === "pulumi:providers:" + Provider.__pulumiType;
    }

    /**
     * Minio Access Key
     *
     * @deprecated use minio_user instead
     */
    public readonly minioAccessKey!: pulumi.Output<string | undefined>;
    /**
     * Minio API Version (type: string, options: v2 or v4, default: v4)
     */
    public readonly minioApiVersion!: pulumi.Output<string | undefined>;
    public readonly minioCacertFile!: pulumi.Output<string | undefined>;
    public readonly minioCertFile!: pulumi.Output<string | undefined>;
    public readonly minioKeyFile!: pulumi.Output<string | undefined>;
    /**
     * Minio Password
     */
    public readonly minioPassword!: pulumi.Output<string | undefined>;
    /**
     * Minio Region (default: us-east-1)
     */
    public readonly minioRegion!: pulumi.Output<string | undefined>;
    /**
     * Minio Secret Key
     *
     * @deprecated use minio_password instead
     */
    public readonly minioSecretKey!: pulumi.Output<string | undefined>;
    /**
     * Minio Host and Port
     */
    public readonly minioServer!: pulumi.Output<string>;
    /**
     * Minio Session Token
     */
    public readonly minioSessionToken!: pulumi.Output<string | undefined>;
    /**
     * Minio User
     */
    public readonly minioUser!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            if ((!args || args.minioServer === undefined) && !opts.urn) {
                throw new Error("Missing required property 'minioServer'");
            }
            resourceInputs["minioAccessKey"] = args ? args.minioAccessKey : undefined;
            resourceInputs["minioApiVersion"] = args ? args.minioApiVersion : undefined;
            resourceInputs["minioCacertFile"] = args ? args.minioCacertFile : undefined;
            resourceInputs["minioCertFile"] = args ? args.minioCertFile : undefined;
            resourceInputs["minioInsecure"] = pulumi.output(args ? args.minioInsecure : undefined).apply(JSON.stringify);
            resourceInputs["minioKeyFile"] = args ? args.minioKeyFile : undefined;
            resourceInputs["minioPassword"] = args ? args.minioPassword : undefined;
            resourceInputs["minioRegion"] = args ? args.minioRegion : undefined;
            resourceInputs["minioSecretKey"] = args ? args.minioSecretKey : undefined;
            resourceInputs["minioServer"] = args ? args.minioServer : undefined;
            resourceInputs["minioSessionToken"] = args ? args.minioSessionToken : undefined;
            resourceInputs["minioSsl"] = pulumi.output(args ? args.minioSsl : undefined).apply(JSON.stringify);
            resourceInputs["minioUser"] = args ? args.minioUser : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * Minio Access Key
     *
     * @deprecated use minio_user instead
     */
    minioAccessKey?: pulumi.Input<string>;
    /**
     * Minio API Version (type: string, options: v2 or v4, default: v4)
     */
    minioApiVersion?: pulumi.Input<string>;
    minioCacertFile?: pulumi.Input<string>;
    minioCertFile?: pulumi.Input<string>;
    /**
     * Disable SSL certificate verification (default: false)
     */
    minioInsecure?: pulumi.Input<boolean>;
    minioKeyFile?: pulumi.Input<string>;
    /**
     * Minio Password
     */
    minioPassword?: pulumi.Input<string>;
    /**
     * Minio Region (default: us-east-1)
     */
    minioRegion?: pulumi.Input<string>;
    /**
     * Minio Secret Key
     *
     * @deprecated use minio_password instead
     */
    minioSecretKey?: pulumi.Input<string>;
    /**
     * Minio Host and Port
     */
    minioServer: pulumi.Input<string>;
    /**
     * Minio Session Token
     */
    minioSessionToken?: pulumi.Input<string>;
    /**
     * Minio SSL enabled (default: false)
     */
    minioSsl?: pulumi.Input<boolean>;
    /**
     * Minio User
     */
    minioUser?: pulumi.Input<string>;
}
