# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['IamUserArgs', 'IamUser']

@pulumi.input_type
class IamUserArgs:
    def __init__(__self__, *,
                 disable_user: Optional[pulumi.Input[bool]] = None,
                 force_destroy: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 secret: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 update_secret: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a IamUser resource.
        :param pulumi.Input[bool] disable_user: Disable user
        :param pulumi.Input[bool] force_destroy: Delete user even if it has non-Terraform-managed IAM access keys
        :param pulumi.Input[bool] update_secret: Rotate Minio User Secret Key
        """
        if disable_user is not None:
            pulumi.set(__self__, "disable_user", disable_user)
        if force_destroy is not None:
            pulumi.set(__self__, "force_destroy", force_destroy)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if secret is not None:
            pulumi.set(__self__, "secret", secret)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if update_secret is not None:
            pulumi.set(__self__, "update_secret", update_secret)

    @property
    @pulumi.getter(name="disableUser")
    def disable_user(self) -> Optional[pulumi.Input[bool]]:
        """
        Disable user
        """
        return pulumi.get(self, "disable_user")

    @disable_user.setter
    def disable_user(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_user", value)

    @property
    @pulumi.getter(name="forceDestroy")
    def force_destroy(self) -> Optional[pulumi.Input[bool]]:
        """
        Delete user even if it has non-Terraform-managed IAM access keys
        """
        return pulumi.get(self, "force_destroy")

    @force_destroy.setter
    def force_destroy(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force_destroy", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def secret(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "secret")

    @secret.setter
    def secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="updateSecret")
    def update_secret(self) -> Optional[pulumi.Input[bool]]:
        """
        Rotate Minio User Secret Key
        """
        return pulumi.get(self, "update_secret")

    @update_secret.setter
    def update_secret(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "update_secret", value)


@pulumi.input_type
class _IamUserState:
    def __init__(__self__, *,
                 disable_user: Optional[pulumi.Input[bool]] = None,
                 force_destroy: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 secret: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 update_secret: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering IamUser resources.
        :param pulumi.Input[bool] disable_user: Disable user
        :param pulumi.Input[bool] force_destroy: Delete user even if it has non-Terraform-managed IAM access keys
        :param pulumi.Input[bool] update_secret: Rotate Minio User Secret Key
        """
        if disable_user is not None:
            pulumi.set(__self__, "disable_user", disable_user)
        if force_destroy is not None:
            pulumi.set(__self__, "force_destroy", force_destroy)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if secret is not None:
            pulumi.set(__self__, "secret", secret)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if update_secret is not None:
            pulumi.set(__self__, "update_secret", update_secret)

    @property
    @pulumi.getter(name="disableUser")
    def disable_user(self) -> Optional[pulumi.Input[bool]]:
        """
        Disable user
        """
        return pulumi.get(self, "disable_user")

    @disable_user.setter
    def disable_user(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_user", value)

    @property
    @pulumi.getter(name="forceDestroy")
    def force_destroy(self) -> Optional[pulumi.Input[bool]]:
        """
        Delete user even if it has non-Terraform-managed IAM access keys
        """
        return pulumi.get(self, "force_destroy")

    @force_destroy.setter
    def force_destroy(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force_destroy", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def secret(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "secret")

    @secret.setter
    def secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="updateSecret")
    def update_secret(self) -> Optional[pulumi.Input[bool]]:
        """
        Rotate Minio User Secret Key
        """
        return pulumi.get(self, "update_secret")

    @update_secret.setter
    def update_secret(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "update_secret", value)


class IamUser(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disable_user: Optional[pulumi.Input[bool]] = None,
                 force_destroy: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 secret: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 update_secret: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_minio as minio

        test_iam_user = minio.IamUser("testIamUser",
            force_destroy=True,
            tags={
                "tag-key": "tag-value",
            })
        pulumi.export("test", test_iam_user.id)
        pulumi.export("status", test_iam_user.status)
        pulumi.export("secret", test_iam_user.secret)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] disable_user: Disable user
        :param pulumi.Input[bool] force_destroy: Delete user even if it has non-Terraform-managed IAM access keys
        :param pulumi.Input[bool] update_secret: Rotate Minio User Secret Key
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[IamUserArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_minio as minio

        test_iam_user = minio.IamUser("testIamUser",
            force_destroy=True,
            tags={
                "tag-key": "tag-value",
            })
        pulumi.export("test", test_iam_user.id)
        pulumi.export("status", test_iam_user.status)
        pulumi.export("secret", test_iam_user.secret)
        ```

        :param str resource_name: The name of the resource.
        :param IamUserArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IamUserArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disable_user: Optional[pulumi.Input[bool]] = None,
                 force_destroy: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 secret: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 update_secret: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IamUserArgs.__new__(IamUserArgs)

            __props__.__dict__["disable_user"] = disable_user
            __props__.__dict__["force_destroy"] = force_destroy
            __props__.__dict__["name"] = name
            __props__.__dict__["secret"] = secret
            __props__.__dict__["tags"] = tags
            __props__.__dict__["update_secret"] = update_secret
            __props__.__dict__["status"] = None
        super(IamUser, __self__).__init__(
            'minio:index/iamUser:IamUser',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            disable_user: Optional[pulumi.Input[bool]] = None,
            force_destroy: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            secret: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            update_secret: Optional[pulumi.Input[bool]] = None) -> 'IamUser':
        """
        Get an existing IamUser resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] disable_user: Disable user
        :param pulumi.Input[bool] force_destroy: Delete user even if it has non-Terraform-managed IAM access keys
        :param pulumi.Input[bool] update_secret: Rotate Minio User Secret Key
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IamUserState.__new__(_IamUserState)

        __props__.__dict__["disable_user"] = disable_user
        __props__.__dict__["force_destroy"] = force_destroy
        __props__.__dict__["name"] = name
        __props__.__dict__["secret"] = secret
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["update_secret"] = update_secret
        return IamUser(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="disableUser")
    def disable_user(self) -> pulumi.Output[Optional[bool]]:
        """
        Disable user
        """
        return pulumi.get(self, "disable_user")

    @property
    @pulumi.getter(name="forceDestroy")
    def force_destroy(self) -> pulumi.Output[Optional[bool]]:
        """
        Delete user even if it has non-Terraform-managed IAM access keys
        """
        return pulumi.get(self, "force_destroy")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def secret(self) -> pulumi.Output[str]:
        return pulumi.get(self, "secret")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updateSecret")
    def update_secret(self) -> pulumi.Output[Optional[bool]]:
        """
        Rotate Minio User Secret Key
        """
        return pulumi.get(self, "update_secret")

