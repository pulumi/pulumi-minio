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

__all__ = [
    'IlmPolicyRule',
    'S3BucketNotificationQueue',
    'S3BucketVersioningVersioningConfiguration',
    'GetIamPolicyDocumentStatementResult',
    'GetIamPolicyDocumentStatementConditionResult',
]

@pulumi.output_type
class IlmPolicyRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "noncurrentVersionExpirationDays":
            suggest = "noncurrent_version_expiration_days"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in IlmPolicyRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        IlmPolicyRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        IlmPolicyRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 id: str,
                 expiration: Optional[str] = None,
                 filter: Optional[str] = None,
                 noncurrent_version_expiration_days: Optional[int] = None,
                 status: Optional[str] = None,
                 tags: Optional[Mapping[str, str]] = None):
        """
        :param str id: The ID of this resource.
        """
        pulumi.set(__self__, "id", id)
        if expiration is not None:
            pulumi.set(__self__, "expiration", expiration)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)
        if noncurrent_version_expiration_days is not None:
            pulumi.set(__self__, "noncurrent_version_expiration_days", noncurrent_version_expiration_days)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

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
    @pulumi.getter(name="noncurrentVersionExpirationDays")
    def noncurrent_version_expiration_days(self) -> Optional[int]:
        return pulumi.get(self, "noncurrent_version_expiration_days")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
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
        pulumi.set(__self__, "events", events)
        pulumi.set(__self__, "queue_arn", queue_arn)
        if filter_prefix is not None:
            pulumi.set(__self__, "filter_prefix", filter_prefix)
        if filter_suffix is not None:
            pulumi.set(__self__, "filter_suffix", filter_suffix)
        if id is not None:
            pulumi.set(__self__, "id", id)

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
        pulumi.set(__self__, "status", status)
        if exclude_folders is not None:
            pulumi.set(__self__, "exclude_folders", exclude_folders)
        if excluded_prefixes is not None:
            pulumi.set(__self__, "excluded_prefixes", excluded_prefixes)

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
        if actions is not None:
            pulumi.set(__self__, "actions", actions)
        if conditions is not None:
            pulumi.set(__self__, "conditions", conditions)
        if effect is not None:
            pulumi.set(__self__, "effect", effect)
        if principal is not None:
            pulumi.set(__self__, "principal", principal)
        if resources is not None:
            pulumi.set(__self__, "resources", resources)
        if sid is not None:
            pulumi.set(__self__, "sid", sid)

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
        pulumi.set(__self__, "test", test)
        pulumi.set(__self__, "values", values)
        pulumi.set(__self__, "variable", variable)

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


