# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

minioAccessKey: Optional[str]
"""
Minio Access Key
"""

minioApiVersion: Optional[str]
"""
Minio API Version (type: string, options: v2 or v4, default: v4)
"""

minioCacertFile: Optional[str]

minioCertFile: Optional[str]

minioInsecure: Optional[bool]

minioKeyFile: Optional[str]

minioRegion: Optional[str]
"""
Minio Region (default: us-east-1)
"""

minioSecretKey: Optional[str]
"""
Minio Secret Key
"""

minioServer: Optional[str]
"""
Minio Host and Port
"""

minioSsl: Optional[bool]
"""
Minio SSL enabled (default: false)
"""

