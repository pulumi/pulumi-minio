# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['IamGroupPolicyAttachmentArgs', 'IamGroupPolicyAttachment']

@pulumi.input_type
class IamGroupPolicyAttachmentArgs:
    def __init__(__self__, *,
                 group_name: pulumi.Input[builtins.str],
                 policy_name: pulumi.Input[builtins.str]):
        """
        The set of arguments for constructing a IamGroupPolicyAttachment resource.
        """
        pulumi.set(__self__, "group_name", group_name)
        pulumi.set(__self__, "policy_name", policy_name)

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "group_name")

    @group_name.setter
    def group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "group_name", value)

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "policy_name")

    @policy_name.setter
    def policy_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "policy_name", value)


@pulumi.input_type
class _IamGroupPolicyAttachmentState:
    def __init__(__self__, *,
                 group_name: Optional[pulumi.Input[builtins.str]] = None,
                 policy_name: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering IamGroupPolicyAttachment resources.
        """
        if group_name is not None:
            pulumi.set(__self__, "group_name", group_name)
        if policy_name is not None:
            pulumi.set(__self__, "policy_name", policy_name)

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "group_name")

    @group_name.setter
    def group_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "group_name", value)

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "policy_name")

    @policy_name.setter
    def policy_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "policy_name", value)


class IamGroupPolicyAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_name: Optional[pulumi.Input[builtins.str]] = None,
                 policy_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IamGroupPolicyAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        :param str resource_name: The name of the resource.
        :param IamGroupPolicyAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IamGroupPolicyAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_name: Optional[pulumi.Input[builtins.str]] = None,
                 policy_name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IamGroupPolicyAttachmentArgs.__new__(IamGroupPolicyAttachmentArgs)

            if group_name is None and not opts.urn:
                raise TypeError("Missing required property 'group_name'")
            __props__.__dict__["group_name"] = group_name
            if policy_name is None and not opts.urn:
                raise TypeError("Missing required property 'policy_name'")
            __props__.__dict__["policy_name"] = policy_name
        super(IamGroupPolicyAttachment, __self__).__init__(
            'minio:index/iamGroupPolicyAttachment:IamGroupPolicyAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            group_name: Optional[pulumi.Input[builtins.str]] = None,
            policy_name: Optional[pulumi.Input[builtins.str]] = None) -> 'IamGroupPolicyAttachment':
        """
        Get an existing IamGroupPolicyAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IamGroupPolicyAttachmentState.__new__(_IamGroupPolicyAttachmentState)

        __props__.__dict__["group_name"] = group_name
        __props__.__dict__["policy_name"] = policy_name
        return IamGroupPolicyAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "group_name")

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "policy_name")

