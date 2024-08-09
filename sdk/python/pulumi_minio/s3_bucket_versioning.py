# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['S3BucketVersioningArgs', 'S3BucketVersioning']

@pulumi.input_type
class S3BucketVersioningArgs:
    def __init__(__self__, *,
                 bucket: pulumi.Input[str],
                 versioning_configuration: pulumi.Input['S3BucketVersioningVersioningConfigurationArgs']):
        """
        The set of arguments for constructing a S3BucketVersioning resource.
        """
        pulumi.set(__self__, "bucket", bucket)
        pulumi.set(__self__, "versioning_configuration", versioning_configuration)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Input[str]:
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter(name="versioningConfiguration")
    def versioning_configuration(self) -> pulumi.Input['S3BucketVersioningVersioningConfigurationArgs']:
        return pulumi.get(self, "versioning_configuration")

    @versioning_configuration.setter
    def versioning_configuration(self, value: pulumi.Input['S3BucketVersioningVersioningConfigurationArgs']):
        pulumi.set(self, "versioning_configuration", value)


@pulumi.input_type
class _S3BucketVersioningState:
    def __init__(__self__, *,
                 bucket: Optional[pulumi.Input[str]] = None,
                 versioning_configuration: Optional[pulumi.Input['S3BucketVersioningVersioningConfigurationArgs']] = None):
        """
        Input properties used for looking up and filtering S3BucketVersioning resources.
        """
        if bucket is not None:
            pulumi.set(__self__, "bucket", bucket)
        if versioning_configuration is not None:
            pulumi.set(__self__, "versioning_configuration", versioning_configuration)

    @property
    @pulumi.getter
    def bucket(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter(name="versioningConfiguration")
    def versioning_configuration(self) -> Optional[pulumi.Input['S3BucketVersioningVersioningConfigurationArgs']]:
        return pulumi.get(self, "versioning_configuration")

    @versioning_configuration.setter
    def versioning_configuration(self, value: Optional[pulumi.Input['S3BucketVersioningVersioningConfigurationArgs']]):
        pulumi.set(self, "versioning_configuration", value)


class S3BucketVersioning(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 versioning_configuration: Optional[pulumi.Input[Union['S3BucketVersioningVersioningConfigurationArgs', 'S3BucketVersioningVersioningConfigurationArgsDict']]] = None,
                 __props__=None):
        """
        Create a S3BucketVersioning resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: S3BucketVersioningArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a S3BucketVersioning resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param S3BucketVersioningArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(S3BucketVersioningArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 versioning_configuration: Optional[pulumi.Input[Union['S3BucketVersioningVersioningConfigurationArgs', 'S3BucketVersioningVersioningConfigurationArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = S3BucketVersioningArgs.__new__(S3BucketVersioningArgs)

            if bucket is None and not opts.urn:
                raise TypeError("Missing required property 'bucket'")
            __props__.__dict__["bucket"] = bucket
            if versioning_configuration is None and not opts.urn:
                raise TypeError("Missing required property 'versioning_configuration'")
            __props__.__dict__["versioning_configuration"] = versioning_configuration
        super(S3BucketVersioning, __self__).__init__(
            'minio:index/s3BucketVersioning:S3BucketVersioning',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bucket: Optional[pulumi.Input[str]] = None,
            versioning_configuration: Optional[pulumi.Input[Union['S3BucketVersioningVersioningConfigurationArgs', 'S3BucketVersioningVersioningConfigurationArgsDict']]] = None) -> 'S3BucketVersioning':
        """
        Get an existing S3BucketVersioning resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _S3BucketVersioningState.__new__(_S3BucketVersioningState)

        __props__.__dict__["bucket"] = bucket
        __props__.__dict__["versioning_configuration"] = versioning_configuration
        return S3BucketVersioning(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Output[str]:
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter(name="versioningConfiguration")
    def versioning_configuration(self) -> pulumi.Output['outputs.S3BucketVersioningVersioningConfiguration']:
        return pulumi.get(self, "versioning_configuration")

