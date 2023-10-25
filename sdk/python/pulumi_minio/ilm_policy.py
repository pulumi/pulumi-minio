# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['IlmPolicyArgs', 'IlmPolicy']

@pulumi.input_type
class IlmPolicyArgs:
    def __init__(__self__, *,
                 bucket: pulumi.Input[str],
                 rules: pulumi.Input[Sequence[pulumi.Input['IlmPolicyRuleArgs']]]):
        """
        The set of arguments for constructing a IlmPolicy resource.
        """
        IlmPolicyArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            bucket=bucket,
            rules=rules,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             bucket: Optional[pulumi.Input[str]] = None,
             rules: Optional[pulumi.Input[Sequence[pulumi.Input['IlmPolicyRuleArgs']]]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if bucket is None:
            raise TypeError("Missing 'bucket' argument")
        if rules is None:
            raise TypeError("Missing 'rules' argument")

        _setter("bucket", bucket)
        _setter("rules", rules)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Input[str]:
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Input[Sequence[pulumi.Input['IlmPolicyRuleArgs']]]:
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: pulumi.Input[Sequence[pulumi.Input['IlmPolicyRuleArgs']]]):
        pulumi.set(self, "rules", value)


@pulumi.input_type
class _IlmPolicyState:
    def __init__(__self__, *,
                 bucket: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['IlmPolicyRuleArgs']]]] = None):
        """
        Input properties used for looking up and filtering IlmPolicy resources.
        """
        _IlmPolicyState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            bucket=bucket,
            rules=rules,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             bucket: Optional[pulumi.Input[str]] = None,
             rules: Optional[pulumi.Input[Sequence[pulumi.Input['IlmPolicyRuleArgs']]]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):

        if bucket is not None:
            _setter("bucket", bucket)
        if rules is not None:
            _setter("rules", rules)

    @property
    @pulumi.getter
    def bucket(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IlmPolicyRuleArgs']]]]:
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IlmPolicyRuleArgs']]]]):
        pulumi.set(self, "rules", value)


class IlmPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IlmPolicyRuleArgs']]]]] = None,
                 __props__=None):
        """
        `IlmPolicy` handles lifecycle settings for a given `S3Bucket`.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_minio as minio

        bucket = minio.S3Bucket("bucket", bucket="bucket")
        bucket_lifecycle_rules = minio.IlmPolicy("bucket-lifecycle-rules",
            bucket=bucket.bucket,
            rules=[minio.IlmPolicyRuleArgs(
                id="expire-7d",
                expiration="7d",
            )])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IlmPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        `IlmPolicy` handles lifecycle settings for a given `S3Bucket`.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_minio as minio

        bucket = minio.S3Bucket("bucket", bucket="bucket")
        bucket_lifecycle_rules = minio.IlmPolicy("bucket-lifecycle-rules",
            bucket=bucket.bucket,
            rules=[minio.IlmPolicyRuleArgs(
                id="expire-7d",
                expiration="7d",
            )])
        ```

        :param str resource_name: The name of the resource.
        :param IlmPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IlmPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            IlmPolicyArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IlmPolicyRuleArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IlmPolicyArgs.__new__(IlmPolicyArgs)

            if bucket is None and not opts.urn:
                raise TypeError("Missing required property 'bucket'")
            __props__.__dict__["bucket"] = bucket
            if rules is None and not opts.urn:
                raise TypeError("Missing required property 'rules'")
            __props__.__dict__["rules"] = rules
        super(IlmPolicy, __self__).__init__(
            'minio:index/ilmPolicy:IlmPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bucket: Optional[pulumi.Input[str]] = None,
            rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IlmPolicyRuleArgs']]]]] = None) -> 'IlmPolicy':
        """
        Get an existing IlmPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IlmPolicyState.__new__(_IlmPolicyState)

        __props__.__dict__["bucket"] = bucket
        __props__.__dict__["rules"] = rules
        return IlmPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Output[str]:
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Sequence['outputs.IlmPolicyRule']]:
        return pulumi.get(self, "rules")

