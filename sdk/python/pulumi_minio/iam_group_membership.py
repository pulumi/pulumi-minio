# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['IamGroupMembershipArgs', 'IamGroupMembership']

@pulumi.input_type
class IamGroupMembershipArgs:
    def __init__(__self__, *,
                 group: pulumi.Input[str],
                 users: pulumi.Input[Sequence[pulumi.Input[str]]],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a IamGroupMembership resource.
        :param pulumi.Input[str] group: Group name to add users
        :param pulumi.Input[Sequence[pulumi.Input[str]]] users: Add user or list of users such as a group membership
        :param pulumi.Input[str] name: Name of group membership
        """
        IamGroupMembershipArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            group=group,
            users=users,
            name=name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             group: pulumi.Input[str],
             users: pulumi.Input[Sequence[pulumi.Input[str]]],
             name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):

        _setter("group", group)
        _setter("users", users)
        if name is not None:
            _setter("name", name)

    @property
    @pulumi.getter
    def group(self) -> pulumi.Input[str]:
        """
        Group name to add users
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: pulumi.Input[str]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter
    def users(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Add user or list of users such as a group membership
        """
        return pulumi.get(self, "users")

    @users.setter
    def users(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "users", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of group membership
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _IamGroupMembershipState:
    def __init__(__self__, *,
                 group: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering IamGroupMembership resources.
        :param pulumi.Input[str] group: Group name to add users
        :param pulumi.Input[str] name: Name of group membership
        :param pulumi.Input[Sequence[pulumi.Input[str]]] users: Add user or list of users such as a group membership
        """
        _IamGroupMembershipState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            group=group,
            name=name,
            users=users,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             group: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):

        if group is not None:
            _setter("group", group)
        if name is not None:
            _setter("name", name)
        if users is not None:
            _setter("users", users)

    @property
    @pulumi.getter
    def group(self) -> Optional[pulumi.Input[str]]:
        """
        Group name to add users
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of group membership
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def users(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Add user or list of users such as a group membership
        """
        return pulumi.get(self, "users")

    @users.setter
    def users(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "users", value)


class IamGroupMembership(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Create a IamGroupMembership resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group: Group name to add users
        :param pulumi.Input[str] name: Name of group membership
        :param pulumi.Input[Sequence[pulumi.Input[str]]] users: Add user or list of users such as a group membership
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IamGroupMembershipArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a IamGroupMembership resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param IamGroupMembershipArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IamGroupMembershipArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            IamGroupMembershipArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IamGroupMembershipArgs.__new__(IamGroupMembershipArgs)

            if group is None and not opts.urn:
                raise TypeError("Missing required property 'group'")
            __props__.__dict__["group"] = group
            __props__.__dict__["name"] = name
            if users is None and not opts.urn:
                raise TypeError("Missing required property 'users'")
            __props__.__dict__["users"] = users
        super(IamGroupMembership, __self__).__init__(
            'minio:index/iamGroupMembership:IamGroupMembership',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            group: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'IamGroupMembership':
        """
        Get an existing IamGroupMembership resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group: Group name to add users
        :param pulumi.Input[str] name: Name of group membership
        :param pulumi.Input[Sequence[pulumi.Input[str]]] users: Add user or list of users such as a group membership
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IamGroupMembershipState.__new__(_IamGroupMembershipState)

        __props__.__dict__["group"] = group
        __props__.__dict__["name"] = name
        __props__.__dict__["users"] = users
        return IamGroupMembership(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def group(self) -> pulumi.Output[str]:
        """
        Group name to add users
        """
        return pulumi.get(self, "group")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of group membership
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def users(self) -> pulumi.Output[Sequence[str]]:
        """
        Add user or list of users such as a group membership
        """
        return pulumi.get(self, "users")

