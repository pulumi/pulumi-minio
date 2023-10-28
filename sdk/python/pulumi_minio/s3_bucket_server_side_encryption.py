# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['S3BucketServerSideEncryptionArgs', 'S3BucketServerSideEncryption']

@pulumi.input_type
class S3BucketServerSideEncryptionArgs:
    def __init__(__self__, *,
                 bucket: pulumi.Input[str],
                 encryption_type: pulumi.Input[str],
                 kms_key_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a S3BucketServerSideEncryption resource.
        """
        pulumi.set(__self__, "bucket", bucket)
        pulumi.set(__self__, "encryption_type", encryption_type)
        pulumi.set(__self__, "kms_key_id", kms_key_id)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Input[str]:
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter(name="encryptionType")
    def encryption_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "encryption_type")

    @encryption_type.setter
    def encryption_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "encryption_type", value)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "kms_key_id", value)


@pulumi.input_type
class _S3BucketServerSideEncryptionState:
    def __init__(__self__, *,
                 bucket: Optional[pulumi.Input[str]] = None,
                 encryption_type: Optional[pulumi.Input[str]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering S3BucketServerSideEncryption resources.
        """
        if bucket is not None:
            pulumi.set(__self__, "bucket", bucket)
        if encryption_type is not None:
            pulumi.set(__self__, "encryption_type", encryption_type)
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)

    @property
    @pulumi.getter
    def bucket(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter(name="encryptionType")
    def encryption_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "encryption_type")

    @encryption_type.setter
    def encryption_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encryption_type", value)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_id", value)


class S3BucketServerSideEncryption(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 encryption_type: Optional[pulumi.Input[str]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a S3BucketServerSideEncryption resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: S3BucketServerSideEncryptionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a S3BucketServerSideEncryption resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param S3BucketServerSideEncryptionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(S3BucketServerSideEncryptionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 encryption_type: Optional[pulumi.Input[str]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = S3BucketServerSideEncryptionArgs.__new__(S3BucketServerSideEncryptionArgs)

            if bucket is None and not opts.urn:
                raise TypeError("Missing required property 'bucket'")
            __props__.__dict__["bucket"] = bucket
            if encryption_type is None and not opts.urn:
                raise TypeError("Missing required property 'encryption_type'")
            __props__.__dict__["encryption_type"] = encryption_type
            if kms_key_id is None and not opts.urn:
                raise TypeError("Missing required property 'kms_key_id'")
            __props__.__dict__["kms_key_id"] = kms_key_id
        super(S3BucketServerSideEncryption, __self__).__init__(
            'minio:index/s3BucketServerSideEncryption:S3BucketServerSideEncryption',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bucket: Optional[pulumi.Input[str]] = None,
            encryption_type: Optional[pulumi.Input[str]] = None,
            kms_key_id: Optional[pulumi.Input[str]] = None) -> 'S3BucketServerSideEncryption':
        """
        Get an existing S3BucketServerSideEncryption resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _S3BucketServerSideEncryptionState.__new__(_S3BucketServerSideEncryptionState)

        __props__.__dict__["bucket"] = bucket
        __props__.__dict__["encryption_type"] = encryption_type
        __props__.__dict__["kms_key_id"] = kms_key_id
        return S3BucketServerSideEncryption(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Output[str]:
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter(name="encryptionType")
    def encryption_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "encryption_type")

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "kms_key_id")

