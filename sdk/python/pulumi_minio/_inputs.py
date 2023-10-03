# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'IlmPolicyRuleArgs',
    'S3BucketNotificationQueueArgs',
    'S3BucketVersioningVersioningConfigurationArgs',
    'GetIamPolicyDocumentStatementArgs',
    'GetIamPolicyDocumentStatementConditionArgs',
]

@pulumi.input_type
class IlmPolicyRuleArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str],
                 expiration: Optional[pulumi.Input[str]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        :param pulumi.Input[str] id: The ID of this resource.
        """
        IlmPolicyRuleArgs._configure(
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
             id: pulumi.Input[str],
             expiration: Optional[pulumi.Input[str]] = None,
             filter: Optional[pulumi.Input[str]] = None,
             status: Optional[pulumi.Input[str]] = None,
             tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
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
    def id(self) -> pulumi.Input[str]:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def expiration(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "expiration")

    @expiration.setter
    def expiration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration", value)

    @property
    @pulumi.getter
    def filter(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter", value)

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


@pulumi.input_type
class S3BucketNotificationQueueArgs:
    def __init__(__self__, *,
                 events: pulumi.Input[Sequence[pulumi.Input[str]]],
                 queue_arn: pulumi.Input[str],
                 filter_prefix: Optional[pulumi.Input[str]] = None,
                 filter_suffix: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] id: The ID of this resource.
        """
        S3BucketNotificationQueueArgs._configure(
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
             events: pulumi.Input[Sequence[pulumi.Input[str]]],
             queue_arn: pulumi.Input[str],
             filter_prefix: Optional[pulumi.Input[str]] = None,
             filter_suffix: Optional[pulumi.Input[str]] = None,
             id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
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
    def events(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "events")

    @events.setter
    def events(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "events", value)

    @property
    @pulumi.getter(name="queueArn")
    def queue_arn(self) -> pulumi.Input[str]:
        return pulumi.get(self, "queue_arn")

    @queue_arn.setter
    def queue_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "queue_arn", value)

    @property
    @pulumi.getter(name="filterPrefix")
    def filter_prefix(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "filter_prefix")

    @filter_prefix.setter
    def filter_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter_prefix", value)

    @property
    @pulumi.getter(name="filterSuffix")
    def filter_suffix(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "filter_suffix")

    @filter_suffix.setter
    def filter_suffix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "filter_suffix", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)


@pulumi.input_type
class S3BucketVersioningVersioningConfigurationArgs:
    def __init__(__self__, *,
                 status: pulumi.Input[str],
                 exclude_folders: Optional[pulumi.Input[bool]] = None,
                 excluded_prefixes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        S3BucketVersioningVersioningConfigurationArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            status=status,
            exclude_folders=exclude_folders,
            excluded_prefixes=excluded_prefixes,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             status: pulumi.Input[str],
             exclude_folders: Optional[pulumi.Input[bool]] = None,
             excluded_prefixes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("status", status)
        if exclude_folders is not None:
            _setter("exclude_folders", exclude_folders)
        if excluded_prefixes is not None:
            _setter("excluded_prefixes", excluded_prefixes)

    @property
    @pulumi.getter
    def status(self) -> pulumi.Input[str]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: pulumi.Input[str]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="excludeFolders")
    def exclude_folders(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "exclude_folders")

    @exclude_folders.setter
    def exclude_folders(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "exclude_folders", value)

    @property
    @pulumi.getter(name="excludedPrefixes")
    def excluded_prefixes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "excluded_prefixes")

    @excluded_prefixes.setter
    def excluded_prefixes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "excluded_prefixes", value)


@pulumi.input_type
class GetIamPolicyDocumentStatementArgs:
    def __init__(__self__, *,
                 actions: Optional[Sequence[str]] = None,
                 conditions: Optional[Sequence['GetIamPolicyDocumentStatementConditionArgs']] = None,
                 effect: Optional[str] = None,
                 principal: Optional[str] = None,
                 resources: Optional[Sequence[str]] = None,
                 sid: Optional[str] = None):
        GetIamPolicyDocumentStatementArgs._configure(
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
             conditions: Optional[Sequence['GetIamPolicyDocumentStatementConditionArgs']] = None,
             effect: Optional[str] = None,
             principal: Optional[str] = None,
             resources: Optional[Sequence[str]] = None,
             sid: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
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

    @actions.setter
    def actions(self, value: Optional[Sequence[str]]):
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
    def effect(self) -> Optional[str]:
        return pulumi.get(self, "effect")

    @effect.setter
    def effect(self, value: Optional[str]):
        pulumi.set(self, "effect", value)

    @property
    @pulumi.getter
    def principal(self) -> Optional[str]:
        return pulumi.get(self, "principal")

    @principal.setter
    def principal(self, value: Optional[str]):
        pulumi.set(self, "principal", value)

    @property
    @pulumi.getter
    def resources(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "resources")

    @resources.setter
    def resources(self, value: Optional[Sequence[str]]):
        pulumi.set(self, "resources", value)

    @property
    @pulumi.getter
    def sid(self) -> Optional[str]:
        return pulumi.get(self, "sid")

    @sid.setter
    def sid(self, value: Optional[str]):
        pulumi.set(self, "sid", value)


@pulumi.input_type
class GetIamPolicyDocumentStatementConditionArgs:
    def __init__(__self__, *,
                 test: str,
                 values: Sequence[str],
                 variable: str):
        GetIamPolicyDocumentStatementConditionArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            test=test,
            values=values,
            variable=variable,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             test: str,
             values: Sequence[str],
             variable: str,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("test", test)
        _setter("values", values)
        _setter("variable", variable)

    @property
    @pulumi.getter
    def test(self) -> str:
        return pulumi.get(self, "test")

    @test.setter
    def test(self, value: str):
        pulumi.set(self, "test", value)

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Sequence[str]):
        pulumi.set(self, "values", value)

    @property
    @pulumi.getter
    def variable(self) -> str:
        return pulumi.get(self, "variable")

    @variable.setter
    def variable(self, value: str):
        pulumi.set(self, "variable", value)


