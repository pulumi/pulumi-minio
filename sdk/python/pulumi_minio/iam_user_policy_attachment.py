# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['IamUserPolicyAttachmentArgs', 'IamUserPolicyAttachment']

@pulumi.input_type
class IamUserPolicyAttachmentArgs:
    def __init__(__self__, *,
                 policy_name: pulumi.Input[str],
                 user_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a IamUserPolicyAttachment resource.
        """
        IamUserPolicyAttachmentArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            policy_name=policy_name,
            user_name=user_name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             policy_name: Optional[pulumi.Input[str]] = None,
             user_name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if policy_name is None and 'policyName' in kwargs:
            policy_name = kwargs['policyName']
        if policy_name is None:
            raise TypeError("Missing 'policy_name' argument")
        if user_name is None and 'userName' in kwargs:
            user_name = kwargs['userName']
        if user_name is None:
            raise TypeError("Missing 'user_name' argument")

        _setter("policy_name", policy_name)
        _setter("user_name", user_name)

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "policy_name")

    @policy_name.setter
    def policy_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_name", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_name", value)


@pulumi.input_type
class _IamUserPolicyAttachmentState:
    def __init__(__self__, *,
                 policy_name: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IamUserPolicyAttachment resources.
        """
        _IamUserPolicyAttachmentState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            policy_name=policy_name,
            user_name=user_name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             policy_name: Optional[pulumi.Input[str]] = None,
             user_name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if policy_name is None and 'policyName' in kwargs:
            policy_name = kwargs['policyName']
        if user_name is None and 'userName' in kwargs:
            user_name = kwargs['userName']

        if policy_name is not None:
            _setter("policy_name", policy_name)
        if user_name is not None:
            _setter("user_name", user_name)

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "policy_name")

    @policy_name.setter
    def policy_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_name", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_name", value)


class IamUserPolicyAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy_name: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_minio as minio

        test_user = minio.IamUser("testUser")
        test_policy = minio.IamPolicy("testPolicy", policy=\"\"\"{
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
        developer_iam_user_policy_attachment = minio.IamUserPolicyAttachment("developerIamUserPolicyAttachment",
            user_name=test_user.id,
            policy_name=test_policy.id)
        pulumi.export("minioName", developer_iam_user_policy_attachment.id)
        pulumi.export("minioUsers", developer_iam_user_policy_attachment.user_name)
        pulumi.export("minioGroup", developer_iam_user_policy_attachment.policy_name)
        developer_index_iam_user_policy_attachment_iam_user_policy_attachment = minio.IamUserPolicyAttachment("developerIndex/iamUserPolicyAttachmentIamUserPolicyAttachment",
            user_name="CN=My User,OU=Unit,DC=example,DC=com",
            policy_name=test_policy.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IamUserPolicyAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_minio as minio

        test_user = minio.IamUser("testUser")
        test_policy = minio.IamPolicy("testPolicy", policy=\"\"\"{
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
        developer_iam_user_policy_attachment = minio.IamUserPolicyAttachment("developerIamUserPolicyAttachment",
            user_name=test_user.id,
            policy_name=test_policy.id)
        pulumi.export("minioName", developer_iam_user_policy_attachment.id)
        pulumi.export("minioUsers", developer_iam_user_policy_attachment.user_name)
        pulumi.export("minioGroup", developer_iam_user_policy_attachment.policy_name)
        developer_index_iam_user_policy_attachment_iam_user_policy_attachment = minio.IamUserPolicyAttachment("developerIndex/iamUserPolicyAttachmentIamUserPolicyAttachment",
            user_name="CN=My User,OU=Unit,DC=example,DC=com",
            policy_name=test_policy.id)
        ```

        :param str resource_name: The name of the resource.
        :param IamUserPolicyAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IamUserPolicyAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            IamUserPolicyAttachmentArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy_name: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IamUserPolicyAttachmentArgs.__new__(IamUserPolicyAttachmentArgs)

            if policy_name is None and not opts.urn:
                raise TypeError("Missing required property 'policy_name'")
            __props__.__dict__["policy_name"] = policy_name
            if user_name is None and not opts.urn:
                raise TypeError("Missing required property 'user_name'")
            __props__.__dict__["user_name"] = user_name
        super(IamUserPolicyAttachment, __self__).__init__(
            'minio:index/iamUserPolicyAttachment:IamUserPolicyAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            policy_name: Optional[pulumi.Input[str]] = None,
            user_name: Optional[pulumi.Input[str]] = None) -> 'IamUserPolicyAttachment':
        """
        Get an existing IamUserPolicyAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IamUserPolicyAttachmentState.__new__(_IamUserPolicyAttachmentState)

        __props__.__dict__["policy_name"] = policy_name
        __props__.__dict__["user_name"] = user_name
        return IamUserPolicyAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "policy_name")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "user_name")

