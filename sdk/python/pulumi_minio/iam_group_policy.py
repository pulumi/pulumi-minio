# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['IamGroupPolicyArgs', 'IamGroupPolicy']

@pulumi.input_type
class IamGroupPolicyArgs:
    def __init__(__self__, *,
                 group: pulumi.Input[str],
                 policy: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a IamGroupPolicy resource.
        """
        IamGroupPolicyArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            group=group,
            policy=policy,
            name=name,
            name_prefix=name_prefix,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             group: Optional[pulumi.Input[str]] = None,
             policy: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             name_prefix: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if group is None:
            raise TypeError("Missing 'group' argument")
        if policy is None:
            raise TypeError("Missing 'policy' argument")
        if name_prefix is None and 'namePrefix' in kwargs:
            name_prefix = kwargs['namePrefix']

        _setter("group", group)
        _setter("policy", policy)
        if name is not None:
            _setter("name", name)
        if name_prefix is not None:
            _setter("name_prefix", name_prefix)

    @property
    @pulumi.getter
    def group(self) -> pulumi.Input[str]:
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: pulumi.Input[str]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Input[str]:
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name_prefix")

    @name_prefix.setter
    def name_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name_prefix", value)


@pulumi.input_type
class _IamGroupPolicyState:
    def __init__(__self__, *,
                 group: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IamGroupPolicy resources.
        """
        _IamGroupPolicyState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            group=group,
            name=name,
            name_prefix=name_prefix,
            policy=policy,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             group: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             name_prefix: Optional[pulumi.Input[str]] = None,
             policy: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if name_prefix is None and 'namePrefix' in kwargs:
            name_prefix = kwargs['namePrefix']

        if group is not None:
            _setter("group", group)
        if name is not None:
            _setter("name", name)
        if name_prefix is not None:
            _setter("name_prefix", name_prefix)
        if policy is not None:
            _setter("policy", policy)

    @property
    @pulumi.getter
    def group(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name_prefix")

    @name_prefix.setter
    def name_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name_prefix", value)

    @property
    @pulumi.getter
    def policy(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy", value)


class IamGroupPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_minio as minio

        developer = minio.IamGroup("developer")
        test_policy = minio.IamGroupPolicy("testPolicy",
            group=developer.id,
            policy=\"\"\"{
          "Version":"2012-10-17",
          "Statement": [
            {
              "Sid":"ListAllBucket",
              "Effect": "Allow",
              "Action": ["s3:PutObject"],
              "Principal":"*",
              "Resource": "arn:aws:s3:::state-terraform-s3/*"
            }
          ]
        }

        \"\"\")
        pulumi.export("minioName", minio_iam_group_membership["developer"]["id"])
        pulumi.export("minioPolicy", minio_iam_group_membership["developer"]["policy"])
        pulumi.export("minioGroup", minio_iam_group_membership["developer"]["group"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IamGroupPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_minio as minio

        developer = minio.IamGroup("developer")
        test_policy = minio.IamGroupPolicy("testPolicy",
            group=developer.id,
            policy=\"\"\"{
          "Version":"2012-10-17",
          "Statement": [
            {
              "Sid":"ListAllBucket",
              "Effect": "Allow",
              "Action": ["s3:PutObject"],
              "Principal":"*",
              "Resource": "arn:aws:s3:::state-terraform-s3/*"
            }
          ]
        }

        \"\"\")
        pulumi.export("minioName", minio_iam_group_membership["developer"]["id"])
        pulumi.export("minioPolicy", minio_iam_group_membership["developer"]["policy"])
        pulumi.export("minioGroup", minio_iam_group_membership["developer"]["group"])
        ```

        :param str resource_name: The name of the resource.
        :param IamGroupPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IamGroupPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            IamGroupPolicyArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IamGroupPolicyArgs.__new__(IamGroupPolicyArgs)

            if group is None and not opts.urn:
                raise TypeError("Missing required property 'group'")
            __props__.__dict__["group"] = group
            __props__.__dict__["name"] = name
            __props__.__dict__["name_prefix"] = name_prefix
            if policy is None and not opts.urn:
                raise TypeError("Missing required property 'policy'")
            __props__.__dict__["policy"] = policy
        super(IamGroupPolicy, __self__).__init__(
            'minio:index/iamGroupPolicy:IamGroupPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            group: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            name_prefix: Optional[pulumi.Input[str]] = None,
            policy: Optional[pulumi.Input[str]] = None) -> 'IamGroupPolicy':
        """
        Get an existing IamGroupPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IamGroupPolicyState.__new__(_IamGroupPolicyState)

        __props__.__dict__["group"] = group
        __props__.__dict__["name"] = name
        __props__.__dict__["name_prefix"] = name_prefix
        __props__.__dict__["policy"] = policy
        return IamGroupPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def group(self) -> pulumi.Output[str]:
        return pulumi.get(self, "group")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "name_prefix")

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Output[str]:
        return pulumi.get(self, "policy")

