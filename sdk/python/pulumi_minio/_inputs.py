# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'IlmPolicyRuleArgs',
    'GetIamPolicyDocumentStatementArgs',
    'GetIamPolicyDocumentStatementConditionArgs',
]

@pulumi.input_type
class IlmPolicyRuleArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str],
                 expiration: Optional[pulumi.Input[str]] = None,
                 filter: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] id: The ID of this resource.
        """
        pulumi.set(__self__, "id", id)
        if expiration is not None:
            pulumi.set(__self__, "expiration", expiration)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)
        if status is not None:
            pulumi.set(__self__, "status", status)

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


@pulumi.input_type
class GetIamPolicyDocumentStatementArgs:
    def __init__(__self__, *,
                 actions: Optional[Sequence[str]] = None,
                 conditions: Optional[Sequence['GetIamPolicyDocumentStatementConditionArgs']] = None,
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
        pulumi.set(__self__, "test", test)
        pulumi.set(__self__, "values", values)
        pulumi.set(__self__, "variable", variable)

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


