# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

import types

__config__ = pulumi.Config('minio')


class _ExportableConfig(types.ModuleType):
    @property
    def minio_access_key(self) -> Optional[str]:
        """
        Minio Access Key
        """
        return __config__.get('minioAccessKey')

    @property
    def minio_api_version(self) -> Optional[str]:
        """
        Minio API Version (type: string, options: v2 or v4, default: v4)
        """
        return __config__.get('minioApiVersion')

    @property
    def minio_cacert_file(self) -> Optional[str]:
        return __config__.get('minioCacertFile')

    @property
    def minio_cert_file(self) -> Optional[str]:
        return __config__.get('minioCertFile')

    @property
    def minio_insecure(self) -> Optional[bool]:
        """
        Disable SSL certificate verification (default: false)
        """
        return __config__.get_bool('minioInsecure')

    @property
    def minio_key_file(self) -> Optional[str]:
        return __config__.get('minioKeyFile')

    @property
    def minio_password(self) -> Optional[str]:
        """
        Minio Password
        """
        return __config__.get('minioPassword')

    @property
    def minio_region(self) -> Optional[str]:
        """
        Minio Region (default: us-east-1)
        """
        return __config__.get('minioRegion')

    @property
    def minio_secret_key(self) -> Optional[str]:
        """
        Minio Secret Key
        """
        return __config__.get('minioSecretKey')

    @property
    def minio_server(self) -> Optional[str]:
        """
        Minio Host and Port
        """
        return __config__.get('minioServer')

    @property
    def minio_session_token(self) -> Optional[str]:
        """
        Minio Session Token
        """
        return __config__.get('minioSessionToken')

    @property
    def minio_ssl(self) -> Optional[bool]:
        """
        Minio SSL enabled (default: false)
        """
        return __config__.get_bool('minioSsl')

    @property
    def minio_user(self) -> Optional[str]:
        """
        Minio User
        """
        return __config__.get('minioUser')

