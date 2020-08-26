# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'MappingMapping',
]

@pulumi.output_type
class MappingMapping(dict):
    def __init__(__self__, *,
                 expression: str,
                 id: str,
                 push_status: Optional[str] = None):
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "id", id)
        if push_status is not None:
            pulumi.set(__self__, "push_status", push_status)

    @property
    @pulumi.getter
    def expression(self) -> str:
        return pulumi.get(self, "expression")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="pushStatus")
    def push_status(self) -> Optional[str]:
        return pulumi.get(self, "push_status")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


