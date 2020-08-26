# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'MappingMappingArgs',
]

@pulumi.input_type
class MappingMappingArgs:
    def __init__(__self__, *,
                 expression: pulumi.Input[str],
                 id: pulumi.Input[str],
                 push_status: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "id", id)
        if push_status is not None:
            pulumi.set(__self__, "push_status", push_status)

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Input[str]:
        return pulumi.get(self, "expression")

    @expression.setter
    def expression(self, value: pulumi.Input[str]):
        pulumi.set(self, "expression", value)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="pushStatus")
    def push_status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "push_status")

    @push_status.setter
    def push_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "push_status", value)


