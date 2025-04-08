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

__all__ = ['IamServiceAccountArgs', 'IamServiceAccount']

@pulumi.input_type
class IamServiceAccountArgs:
    def __init__(__self__, *,
                 target_user: pulumi.Input[builtins.str],
                 disable_user: Optional[pulumi.Input[builtins.bool]] = None,
                 policy: Optional[pulumi.Input[builtins.str]] = None,
                 update_secret: Optional[pulumi.Input[builtins.bool]] = None):
        """
        The set of arguments for constructing a IamServiceAccount resource.
        :param pulumi.Input[builtins.bool] disable_user: Disable service account
        :param pulumi.Input[builtins.str] policy: policy of service account
        :param pulumi.Input[builtins.bool] update_secret: rotate secret key
        """
        pulumi.set(__self__, "target_user", target_user)
        if disable_user is not None:
            pulumi.set(__self__, "disable_user", disable_user)
        if policy is not None:
            pulumi.set(__self__, "policy", policy)
        if update_secret is not None:
            pulumi.set(__self__, "update_secret", update_secret)

    @property
    @pulumi.getter(name="targetUser")
    def target_user(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "target_user")

    @target_user.setter
    def target_user(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "target_user", value)

    @property
    @pulumi.getter(name="disableUser")
    def disable_user(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Disable service account
        """
        return pulumi.get(self, "disable_user")

    @disable_user.setter
    def disable_user(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "disable_user", value)

    @property
    @pulumi.getter
    def policy(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        policy of service account
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "policy", value)

    @property
    @pulumi.getter(name="updateSecret")
    def update_secret(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        rotate secret key
        """
        return pulumi.get(self, "update_secret")

    @update_secret.setter
    def update_secret(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "update_secret", value)


@pulumi.input_type
class _IamServiceAccountState:
    def __init__(__self__, *,
                 access_key: Optional[pulumi.Input[builtins.str]] = None,
                 disable_user: Optional[pulumi.Input[builtins.bool]] = None,
                 policy: Optional[pulumi.Input[builtins.str]] = None,
                 secret_key: Optional[pulumi.Input[builtins.str]] = None,
                 status: Optional[pulumi.Input[builtins.str]] = None,
                 target_user: Optional[pulumi.Input[builtins.str]] = None,
                 update_secret: Optional[pulumi.Input[builtins.bool]] = None):
        """
        Input properties used for looking up and filtering IamServiceAccount resources.
        :param pulumi.Input[builtins.bool] disable_user: Disable service account
        :param pulumi.Input[builtins.str] policy: policy of service account
        :param pulumi.Input[builtins.bool] update_secret: rotate secret key
        """
        if access_key is not None:
            pulumi.set(__self__, "access_key", access_key)
        if disable_user is not None:
            pulumi.set(__self__, "disable_user", disable_user)
        if policy is not None:
            pulumi.set(__self__, "policy", policy)
        if secret_key is not None:
            pulumi.set(__self__, "secret_key", secret_key)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if target_user is not None:
            pulumi.set(__self__, "target_user", target_user)
        if update_secret is not None:
            pulumi.set(__self__, "update_secret", update_secret)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "access_key")

    @access_key.setter
    def access_key(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "access_key", value)

    @property
    @pulumi.getter(name="disableUser")
    def disable_user(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Disable service account
        """
        return pulumi.get(self, "disable_user")

    @disable_user.setter
    def disable_user(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "disable_user", value)

    @property
    @pulumi.getter
    def policy(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        policy of service account
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "policy", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "secret_key", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="targetUser")
    def target_user(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "target_user")

    @target_user.setter
    def target_user(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "target_user", value)

    @property
    @pulumi.getter(name="updateSecret")
    def update_secret(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        rotate secret key
        """
        return pulumi.get(self, "update_secret")

    @update_secret.setter
    def update_secret(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "update_secret", value)


class IamServiceAccount(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disable_user: Optional[pulumi.Input[builtins.bool]] = None,
                 policy: Optional[pulumi.Input[builtins.str]] = None,
                 target_user: Optional[pulumi.Input[builtins.str]] = None,
                 update_secret: Optional[pulumi.Input[builtins.bool]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_minio as minio

        test = minio.IamUser("test",
            name="test",
            force_destroy=True,
            tags={
                "tag-key": "tag-value",
            })
        test_service_account = minio.IamServiceAccount("test_service_account", target_user=test.name)
        pulumi.export("minioUser", test_service_account.access_key)
        pulumi.export("minioPassword", test_service_account.secret_key)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] disable_user: Disable service account
        :param pulumi.Input[builtins.str] policy: policy of service account
        :param pulumi.Input[builtins.bool] update_secret: rotate secret key
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IamServiceAccountArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_minio as minio

        test = minio.IamUser("test",
            name="test",
            force_destroy=True,
            tags={
                "tag-key": "tag-value",
            })
        test_service_account = minio.IamServiceAccount("test_service_account", target_user=test.name)
        pulumi.export("minioUser", test_service_account.access_key)
        pulumi.export("minioPassword", test_service_account.secret_key)
        ```

        :param str resource_name: The name of the resource.
        :param IamServiceAccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IamServiceAccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disable_user: Optional[pulumi.Input[builtins.bool]] = None,
                 policy: Optional[pulumi.Input[builtins.str]] = None,
                 target_user: Optional[pulumi.Input[builtins.str]] = None,
                 update_secret: Optional[pulumi.Input[builtins.bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IamServiceAccountArgs.__new__(IamServiceAccountArgs)

            __props__.__dict__["disable_user"] = disable_user
            __props__.__dict__["policy"] = policy
            if target_user is None and not opts.urn:
                raise TypeError("Missing required property 'target_user'")
            __props__.__dict__["target_user"] = target_user
            __props__.__dict__["update_secret"] = update_secret
            __props__.__dict__["access_key"] = None
            __props__.__dict__["secret_key"] = None
            __props__.__dict__["status"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["secretKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(IamServiceAccount, __self__).__init__(
            'minio:index/iamServiceAccount:IamServiceAccount',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_key: Optional[pulumi.Input[builtins.str]] = None,
            disable_user: Optional[pulumi.Input[builtins.bool]] = None,
            policy: Optional[pulumi.Input[builtins.str]] = None,
            secret_key: Optional[pulumi.Input[builtins.str]] = None,
            status: Optional[pulumi.Input[builtins.str]] = None,
            target_user: Optional[pulumi.Input[builtins.str]] = None,
            update_secret: Optional[pulumi.Input[builtins.bool]] = None) -> 'IamServiceAccount':
        """
        Get an existing IamServiceAccount resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] disable_user: Disable service account
        :param pulumi.Input[builtins.str] policy: policy of service account
        :param pulumi.Input[builtins.bool] update_secret: rotate secret key
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IamServiceAccountState.__new__(_IamServiceAccountState)

        __props__.__dict__["access_key"] = access_key
        __props__.__dict__["disable_user"] = disable_user
        __props__.__dict__["policy"] = policy
        __props__.__dict__["secret_key"] = secret_key
        __props__.__dict__["status"] = status
        __props__.__dict__["target_user"] = target_user
        __props__.__dict__["update_secret"] = update_secret
        return IamServiceAccount(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "access_key")

    @property
    @pulumi.getter(name="disableUser")
    def disable_user(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Disable service account
        """
        return pulumi.get(self, "disable_user")

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        policy of service account
        """
        return pulumi.get(self, "policy")

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "secret_key")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="targetUser")
    def target_user(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "target_user")

    @property
    @pulumi.getter(name="updateSecret")
    def update_secret(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        rotate secret key
        """
        return pulumi.get(self, "update_secret")

