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

__all__ = [
    'IlmPolicyRule',
    'S3BucketNotificationQueue',
    'S3BucketVersioningVersioningConfiguration',
    'GetIamPolicyDocumentStatementResult',
    'GetIamPolicyDocumentStatementConditionResult',
]

@pulumi.output_type
class IlmPolicyRule(dict):
    def __init__(__self__, *,
                 id: str,
                 expiration: Optional[str] = None,
                 filter: Optional[str] = None,
                 status: Optional[str] = None,
                 tags: Optional[Mapping[str, Any]] = None):
        """
        :param str id: The ID of this resource.
        """
        IlmPolicyRule._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            id=id,
            expiration=expiration,
            filter=filter,
            status=status,
            tags=tags,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             id: Optional[str] = None,
             expiration: Optional[str] = None,
             filter: Optional[str] = None,
             status: Optional[str] = None,
             tags: Optional[Mapping[str, Any]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if id is None:
            raise TypeError("Missing 'id' argument")

        _setter("id", id)
        if expiration is not None:
            _setter("expiration", expiration)
        if filter is not None:
            _setter("filter", filter)
        if status is not None:
            _setter("status", status)
        if tags is not None:
            _setter("tags", tags)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def expiration(self) -> Optional[str]:
        return pulumi.get(self, "expiration")

    @property
    @pulumi.getter
    def filter(self) -> Optional[str]:
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, Any]]:
        return pulumi.get(self, "tags")


@pulumi.output_type
class S3BucketNotificationQueue(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "queueArn":
            suggest = "queue_arn"
        elif key == "filterPrefix":
            suggest = "filter_prefix"
        elif key == "filterSuffix":
            suggest = "filter_suffix"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in S3BucketNotificationQueue. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        S3BucketNotificationQueue.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        S3BucketNotificationQueue.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 events: Sequence[str],
                 queue_arn: str,
                 filter_prefix: Optional[str] = None,
                 filter_suffix: Optional[str] = None,
                 id: Optional[str] = None):
        """
        :param str id: The ID of this resource.
        """
        S3BucketNotificationQueue._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            events=events,
            queue_arn=queue_arn,
            filter_prefix=filter_prefix,
            filter_suffix=filter_suffix,
            id=id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             events: Optional[Sequence[str]] = None,
             queue_arn: Optional[str] = None,
             filter_prefix: Optional[str] = None,
             filter_suffix: Optional[str] = None,
             id: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if events is None:
            raise TypeError("Missing 'events' argument")
        if queue_arn is None and 'queueArn' in kwargs:
            queue_arn = kwargs['queueArn']
        if queue_arn is None:
            raise TypeError("Missing 'queue_arn' argument")
        if filter_prefix is None and 'filterPrefix' in kwargs:
            filter_prefix = kwargs['filterPrefix']
        if filter_suffix is None and 'filterSuffix' in kwargs:
            filter_suffix = kwargs['filterSuffix']

        _setter("events", events)
        _setter("queue_arn", queue_arn)
        if filter_prefix is not None:
            _setter("filter_prefix", filter_prefix)
        if filter_suffix is not None:
            _setter("filter_suffix", filter_suffix)
        if id is not None:
            _setter("id", id)

    @property
    @pulumi.getter
    def events(self) -> Sequence[str]:
        return pulumi.get(self, "events")

    @property
    @pulumi.getter(name="queueArn")
    def queue_arn(self) -> str:
        return pulumi.get(self, "queue_arn")

    @property
    @pulumi.getter(name="filterPrefix")
    def filter_prefix(self) -> Optional[str]:
        return pulumi.get(self, "filter_prefix")

    @property
    @pulumi.getter(name="filterSuffix")
    def filter_suffix(self) -> Optional[str]:
        return pulumi.get(self, "filter_suffix")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")


@pulumi.output_type
class S3BucketVersioningVersioningConfiguration(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "excludeFolders":
            suggest = "exclude_folders"
        elif key == "excludedPrefixes":
            suggest = "excluded_prefixes"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in S3BucketVersioningVersioningConfiguration. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        S3BucketVersioningVersioningConfiguration.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        S3BucketVersioningVersioningConfiguration.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 status: str,
                 exclude_folders: Optional[bool] = None,
                 excluded_prefixes: Optional[Sequence[str]] = None):
        S3BucketVersioningVersioningConfiguration._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            status=status,
            exclude_folders=exclude_folders,
            excluded_prefixes=excluded_prefixes,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             status: Optional[str] = None,
             exclude_folders: Optional[bool] = None,
             excluded_prefixes: Optional[Sequence[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if status is None:
            raise TypeError("Missing 'status' argument")
        if exclude_folders is None and 'excludeFolders' in kwargs:
            exclude_folders = kwargs['excludeFolders']
        if excluded_prefixes is None and 'excludedPrefixes' in kwargs:
            excluded_prefixes = kwargs['excludedPrefixes']

        _setter("status", status)
        if exclude_folders is not None:
            _setter("exclude_folders", exclude_folders)
        if excluded_prefixes is not None:
            _setter("excluded_prefixes", excluded_prefixes)

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="excludeFolders")
    def exclude_folders(self) -> Optional[bool]:
        return pulumi.get(self, "exclude_folders")

    @property
    @pulumi.getter(name="excludedPrefixes")
    def excluded_prefixes(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "excluded_prefixes")


@pulumi.output_type
class GetIamPolicyDocumentStatementResult(dict):
    def __init__(__self__, *,
                 actions: Optional[Sequence[str]] = None,
                 conditions: Optional[Sequence['outputs.GetIamPolicyDocumentStatementConditionResult']] = None,
                 effect: Optional[str] = None,
                 principal: Optional[str] = None,
                 resources: Optional[Sequence[str]] = None,
                 sid: Optional[str] = None):
        GetIamPolicyDocumentStatementResult._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            actions=actions,
            conditions=conditions,
            effect=effect,
            principal=principal,
            resources=resources,
            sid=sid,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             actions: Optional[Sequence[str]] = None,
             conditions: Optional[Sequence['outputs.GetIamPolicyDocumentStatementConditionResult']] = None,
             effect: Optional[str] = None,
             principal: Optional[str] = None,
             resources: Optional[Sequence[str]] = None,
             sid: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):

        if actions is not None:
            _setter("actions", actions)
        if conditions is not None:
            _setter("conditions", conditions)
        if effect is not None:
            _setter("effect", effect)
        if principal is not None:
            _setter("principal", principal)
        if resources is not None:
            _setter("resources", resources)
        if sid is not None:
            _setter("sid", sid)

    @property
    @pulumi.getter
    def actions(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "actions")

    @property
    @pulumi.getter
    def conditions(self) -> Optional[Sequence['outputs.GetIamPolicyDocumentStatementConditionResult']]:
        return pulumi.get(self, "conditions")

    @property
    @pulumi.getter
    def effect(self) -> Optional[str]:
        return pulumi.get(self, "effect")

    @property
    @pulumi.getter
    def principal(self) -> Optional[str]:
        return pulumi.get(self, "principal")

    @property
    @pulumi.getter
    def resources(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "resources")

    @property
    @pulumi.getter
    def sid(self) -> Optional[str]:
        return pulumi.get(self, "sid")


@pulumi.output_type
class GetIamPolicyDocumentStatementConditionResult(dict):
    def __init__(__self__, *,
                 test: str,
                 values: Sequence[str],
                 variable: str):
        GetIamPolicyDocumentStatementConditionResult._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            test=test,
            values=values,
            variable=variable,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             test: Optional[str] = None,
             values: Optional[Sequence[str]] = None,
             variable: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if test is None:
            raise TypeError("Missing 'test' argument")
        if values is None:
            raise TypeError("Missing 'values' argument")
        if variable is None:
            raise TypeError("Missing 'variable' argument")

        _setter("test", test)
        _setter("values", values)
        _setter("variable", variable)

    @property
    @pulumi.getter
    def test(self) -> str:
        return pulumi.get(self, "test")

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        return pulumi.get(self, "values")

    @property
    @pulumi.getter
    def variable(self) -> str:
        return pulumi.get(self, "variable")


