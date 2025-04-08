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

__all__ = [
    'IlmPolicyRuleArgs',
    'IlmPolicyRuleArgsDict',
    'S3BucketNotificationQueueArgs',
    'S3BucketNotificationQueueArgsDict',
    'S3BucketVersioningVersioningConfigurationArgs',
    'S3BucketVersioningVersioningConfigurationArgsDict',
    'GetIamPolicyDocumentStatementArgs',
    'GetIamPolicyDocumentStatementArgsDict',
    'GetIamPolicyDocumentStatementConditionArgs',
    'GetIamPolicyDocumentStatementConditionArgsDict',
]

MYPY = False

if not MYPY:
    class IlmPolicyRuleArgsDict(TypedDict):
        id: pulumi.Input[builtins.str]
        """
        The ID of this resource.
        """
        expiration: NotRequired[pulumi.Input[builtins.str]]
        filter: NotRequired[pulumi.Input[builtins.str]]
        noncurrent_version_expiration_days: NotRequired[pulumi.Input[builtins.int]]
        status: NotRequired[pulumi.Input[builtins.str]]
        tags: NotRequired[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]
elif False:
    IlmPolicyRuleArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class IlmPolicyRuleArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[builtins.str],
                 expiration: Optional[pulumi.Input[builtins.str]] = None,
                 filter: Optional[pulumi.Input[builtins.str]] = None,
                 noncurrent_version_expiration_days: Optional[pulumi.Input[builtins.int]] = None,
                 status: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        :param pulumi.Input[builtins.str] id: The ID of this resource.
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
    def id(self) -> pulumi.Input[builtins.str]:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def expiration(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "expiration")

    @expiration.setter
    def expiration(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "expiration", value)

    @property
    @pulumi.getter
    def filter(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "filter", value)

    @property
    @pulumi.getter(name="noncurrentVersionExpirationDays")
    def noncurrent_version_expiration_days(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "noncurrent_version_expiration_days")

    @noncurrent_version_expiration_days.setter
    def noncurrent_version_expiration_days(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "noncurrent_version_expiration_days", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)


if not MYPY:
    class S3BucketNotificationQueueArgsDict(TypedDict):
        events: pulumi.Input[Sequence[pulumi.Input[builtins.str]]]
        queue_arn: pulumi.Input[builtins.str]
        filter_prefix: NotRequired[pulumi.Input[builtins.str]]
        filter_suffix: NotRequired[pulumi.Input[builtins.str]]
        id: NotRequired[pulumi.Input[builtins.str]]
        """
        The ID of this resource.
        """
elif False:
    S3BucketNotificationQueueArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class S3BucketNotificationQueueArgs:
    def __init__(__self__, *,
                 events: pulumi.Input[Sequence[pulumi.Input[builtins.str]]],
                 queue_arn: pulumi.Input[builtins.str],
                 filter_prefix: Optional[pulumi.Input[builtins.str]] = None,
                 filter_suffix: Optional[pulumi.Input[builtins.str]] = None,
                 id: Optional[pulumi.Input[builtins.str]] = None):
        """
        :param pulumi.Input[builtins.str] id: The ID of this resource.
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
    def events(self) -> pulumi.Input[Sequence[pulumi.Input[builtins.str]]]:
        return pulumi.get(self, "events")

    @events.setter
    def events(self, value: pulumi.Input[Sequence[pulumi.Input[builtins.str]]]):
        pulumi.set(self, "events", value)

    @property
    @pulumi.getter(name="queueArn")
    def queue_arn(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "queue_arn")

    @queue_arn.setter
    def queue_arn(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "queue_arn", value)

    @property
    @pulumi.getter(name="filterPrefix")
    def filter_prefix(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "filter_prefix")

    @filter_prefix.setter
    def filter_prefix(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "filter_prefix", value)

    @property
    @pulumi.getter(name="filterSuffix")
    def filter_suffix(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "filter_suffix")

    @filter_suffix.setter
    def filter_suffix(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "filter_suffix", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "id", value)


if not MYPY:
    class S3BucketVersioningVersioningConfigurationArgsDict(TypedDict):
        status: pulumi.Input[builtins.str]
        exclude_folders: NotRequired[pulumi.Input[builtins.bool]]
        excluded_prefixes: NotRequired[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]
elif False:
    S3BucketVersioningVersioningConfigurationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class S3BucketVersioningVersioningConfigurationArgs:
    def __init__(__self__, *,
                 status: pulumi.Input[builtins.str],
                 exclude_folders: Optional[pulumi.Input[builtins.bool]] = None,
                 excluded_prefixes: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None):
        pulumi.set(__self__, "status", status)
        if exclude_folders is not None:
            pulumi.set(__self__, "exclude_folders", exclude_folders)
        if excluded_prefixes is not None:
            pulumi.set(__self__, "excluded_prefixes", excluded_prefixes)

    @property
    @pulumi.getter
    def status(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="excludeFolders")
    def exclude_folders(self) -> Optional[pulumi.Input[builtins.bool]]:
        return pulumi.get(self, "exclude_folders")

    @exclude_folders.setter
    def exclude_folders(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "exclude_folders", value)

    @property
    @pulumi.getter(name="excludedPrefixes")
    def excluded_prefixes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "excluded_prefixes")

    @excluded_prefixes.setter
    def excluded_prefixes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "excluded_prefixes", value)


if not MYPY:
    class GetIamPolicyDocumentStatementArgsDict(TypedDict):
        actions: NotRequired[Sequence[builtins.str]]
        conditions: NotRequired[Sequence['GetIamPolicyDocumentStatementConditionArgsDict']]
        effect: NotRequired[builtins.str]
        principal: NotRequired[builtins.str]
        resources: NotRequired[Sequence[builtins.str]]
        sid: NotRequired[builtins.str]
elif False:
    GetIamPolicyDocumentStatementArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class GetIamPolicyDocumentStatementArgs:
    def __init__(__self__, *,
                 actions: Optional[Sequence[builtins.str]] = None,
                 conditions: Optional[Sequence['GetIamPolicyDocumentStatementConditionArgs']] = None,
                 effect: Optional[builtins.str] = None,
                 principal: Optional[builtins.str] = None,
                 resources: Optional[Sequence[builtins.str]] = None,
                 sid: Optional[builtins.str] = None):
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
    def actions(self) -> Optional[Sequence[builtins.str]]:
        return pulumi.get(self, "actions")

    @actions.setter
    def actions(self, value: Optional[Sequence[builtins.str]]):
        pulumi.set(self, "actions", value)

    @property
    @pulumi.getter
    def conditions(self) -> Optional[Sequence['GetIamPolicyDocumentStatementConditionArgs']]:
        return pulumi.get(self, "conditions")

    @conditions.setter
    def conditions(self, value: Optional[Sequence['GetIamPolicyDocumentStatementConditionArgs']]):
        pulumi.set(self, "conditions", value)

    @property
    @pulumi.getter
    def effect(self) -> Optional[builtins.str]:
        return pulumi.get(self, "effect")

    @effect.setter
    def effect(self, value: Optional[builtins.str]):
        pulumi.set(self, "effect", value)

    @property
    @pulumi.getter
    def principal(self) -> Optional[builtins.str]:
        return pulumi.get(self, "principal")

    @principal.setter
    def principal(self, value: Optional[builtins.str]):
        pulumi.set(self, "principal", value)

    @property
    @pulumi.getter
    def resources(self) -> Optional[Sequence[builtins.str]]:
        return pulumi.get(self, "resources")

    @resources.setter
    def resources(self, value: Optional[Sequence[builtins.str]]):
        pulumi.set(self, "resources", value)

    @property
    @pulumi.getter
    def sid(self) -> Optional[builtins.str]:
        return pulumi.get(self, "sid")

    @sid.setter
    def sid(self, value: Optional[builtins.str]):
        pulumi.set(self, "sid", value)


if not MYPY:
    class GetIamPolicyDocumentStatementConditionArgsDict(TypedDict):
        test: builtins.str
        values: Sequence[builtins.str]
        variable: builtins.str
elif False:
    GetIamPolicyDocumentStatementConditionArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class GetIamPolicyDocumentStatementConditionArgs:
    def __init__(__self__, *,
                 test: builtins.str,
                 values: Sequence[builtins.str],
                 variable: builtins.str):
        pulumi.set(__self__, "test", test)
        pulumi.set(__self__, "values", values)
        pulumi.set(__self__, "variable", variable)

    @property
    @pulumi.getter
    def test(self) -> builtins.str:
        return pulumi.get(self, "test")

    @test.setter
    def test(self, value: builtins.str):
        pulumi.set(self, "test", value)

    @property
    @pulumi.getter
    def values(self) -> Sequence[builtins.str]:
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Sequence[builtins.str]):
        pulumi.set(self, "values", value)

    @property
    @pulumi.getter
    def variable(self) -> builtins.str:
        return pulumi.get(self, "variable")

    @variable.setter
    def variable(self, value: builtins.str):
        pulumi.set(self, "variable", value)


