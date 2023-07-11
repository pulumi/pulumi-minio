# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 minio_server: pulumi.Input[str],
                 minio_access_key: Optional[pulumi.Input[str]] = None,
                 minio_api_version: Optional[pulumi.Input[str]] = None,
                 minio_cacert_file: Optional[pulumi.Input[str]] = None,
                 minio_cert_file: Optional[pulumi.Input[str]] = None,
                 minio_insecure: Optional[pulumi.Input[bool]] = None,
                 minio_key_file: Optional[pulumi.Input[str]] = None,
                 minio_password: Optional[pulumi.Input[str]] = None,
                 minio_region: Optional[pulumi.Input[str]] = None,
                 minio_secret_key: Optional[pulumi.Input[str]] = None,
                 minio_session_token: Optional[pulumi.Input[str]] = None,
                 minio_ssl: Optional[pulumi.Input[bool]] = None,
                 minio_user: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] minio_server: Minio Host and Port
        :param pulumi.Input[str] minio_access_key: Minio Access Key
        :param pulumi.Input[str] minio_api_version: Minio API Version (type: string, options: v2 or v4, default: v4)
        :param pulumi.Input[bool] minio_insecure: Disable SSL certificate verification (default: false)
        :param pulumi.Input[str] minio_password: Minio Password
        :param pulumi.Input[str] minio_region: Minio Region (default: us-east-1)
        :param pulumi.Input[str] minio_secret_key: Minio Secret Key
        :param pulumi.Input[str] minio_session_token: Minio Session Token
        :param pulumi.Input[bool] minio_ssl: Minio SSL enabled (default: false)
        :param pulumi.Input[str] minio_user: Minio User
        """
        pulumi.set(__self__, "minio_server", minio_server)
        if minio_access_key is not None:
            warnings.warn("""use minio_user instead""", DeprecationWarning)
            pulumi.log.warn("""minio_access_key is deprecated: use minio_user instead""")
        if minio_access_key is not None:
            pulumi.set(__self__, "minio_access_key", minio_access_key)
        if minio_api_version is not None:
            pulumi.set(__self__, "minio_api_version", minio_api_version)
        if minio_cacert_file is not None:
            pulumi.set(__self__, "minio_cacert_file", minio_cacert_file)
        if minio_cert_file is not None:
            pulumi.set(__self__, "minio_cert_file", minio_cert_file)
        if minio_insecure is not None:
            pulumi.set(__self__, "minio_insecure", minio_insecure)
        if minio_key_file is not None:
            pulumi.set(__self__, "minio_key_file", minio_key_file)
        if minio_password is not None:
            pulumi.set(__self__, "minio_password", minio_password)
        if minio_region is not None:
            pulumi.set(__self__, "minio_region", minio_region)
        if minio_secret_key is not None:
            warnings.warn("""use minio_password instead""", DeprecationWarning)
            pulumi.log.warn("""minio_secret_key is deprecated: use minio_password instead""")
        if minio_secret_key is not None:
            pulumi.set(__self__, "minio_secret_key", minio_secret_key)
        if minio_session_token is not None:
            pulumi.set(__self__, "minio_session_token", minio_session_token)
        if minio_ssl is not None:
            pulumi.set(__self__, "minio_ssl", minio_ssl)
        if minio_user is not None:
            pulumi.set(__self__, "minio_user", minio_user)

    @property
    @pulumi.getter(name="minioServer")
    def minio_server(self) -> pulumi.Input[str]:
        """
        Minio Host and Port
        """
        return pulumi.get(self, "minio_server")

    @minio_server.setter
    def minio_server(self, value: pulumi.Input[str]):
        pulumi.set(self, "minio_server", value)

    @property
    @pulumi.getter(name="minioAccessKey")
    def minio_access_key(self) -> Optional[pulumi.Input[str]]:
        """
        Minio Access Key
        """
        warnings.warn("""use minio_user instead""", DeprecationWarning)
        pulumi.log.warn("""minio_access_key is deprecated: use minio_user instead""")

        return pulumi.get(self, "minio_access_key")

    @minio_access_key.setter
    def minio_access_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "minio_access_key", value)

    @property
    @pulumi.getter(name="minioApiVersion")
    def minio_api_version(self) -> Optional[pulumi.Input[str]]:
        """
        Minio API Version (type: string, options: v2 or v4, default: v4)
        """
        return pulumi.get(self, "minio_api_version")

    @minio_api_version.setter
    def minio_api_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "minio_api_version", value)

    @property
    @pulumi.getter(name="minioCacertFile")
    def minio_cacert_file(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "minio_cacert_file")

    @minio_cacert_file.setter
    def minio_cacert_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "minio_cacert_file", value)

    @property
    @pulumi.getter(name="minioCertFile")
    def minio_cert_file(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "minio_cert_file")

    @minio_cert_file.setter
    def minio_cert_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "minio_cert_file", value)

    @property
    @pulumi.getter(name="minioInsecure")
    def minio_insecure(self) -> Optional[pulumi.Input[bool]]:
        """
        Disable SSL certificate verification (default: false)
        """
        return pulumi.get(self, "minio_insecure")

    @minio_insecure.setter
    def minio_insecure(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "minio_insecure", value)

    @property
    @pulumi.getter(name="minioKeyFile")
    def minio_key_file(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "minio_key_file")

    @minio_key_file.setter
    def minio_key_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "minio_key_file", value)

    @property
    @pulumi.getter(name="minioPassword")
    def minio_password(self) -> Optional[pulumi.Input[str]]:
        """
        Minio Password
        """
        return pulumi.get(self, "minio_password")

    @minio_password.setter
    def minio_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "minio_password", value)

    @property
    @pulumi.getter(name="minioRegion")
    def minio_region(self) -> Optional[pulumi.Input[str]]:
        """
        Minio Region (default: us-east-1)
        """
        return pulumi.get(self, "minio_region")

    @minio_region.setter
    def minio_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "minio_region", value)

    @property
    @pulumi.getter(name="minioSecretKey")
    def minio_secret_key(self) -> Optional[pulumi.Input[str]]:
        """
        Minio Secret Key
        """
        warnings.warn("""use minio_password instead""", DeprecationWarning)
        pulumi.log.warn("""minio_secret_key is deprecated: use minio_password instead""")

        return pulumi.get(self, "minio_secret_key")

    @minio_secret_key.setter
    def minio_secret_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "minio_secret_key", value)

    @property
    @pulumi.getter(name="minioSessionToken")
    def minio_session_token(self) -> Optional[pulumi.Input[str]]:
        """
        Minio Session Token
        """
        return pulumi.get(self, "minio_session_token")

    @minio_session_token.setter
    def minio_session_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "minio_session_token", value)

    @property
    @pulumi.getter(name="minioSsl")
    def minio_ssl(self) -> Optional[pulumi.Input[bool]]:
        """
        Minio SSL enabled (default: false)
        """
        return pulumi.get(self, "minio_ssl")

    @minio_ssl.setter
    def minio_ssl(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "minio_ssl", value)

    @property
    @pulumi.getter(name="minioUser")
    def minio_user(self) -> Optional[pulumi.Input[str]]:
        """
        Minio User
        """
        return pulumi.get(self, "minio_user")

    @minio_user.setter
    def minio_user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "minio_user", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 minio_access_key: Optional[pulumi.Input[str]] = None,
                 minio_api_version: Optional[pulumi.Input[str]] = None,
                 minio_cacert_file: Optional[pulumi.Input[str]] = None,
                 minio_cert_file: Optional[pulumi.Input[str]] = None,
                 minio_insecure: Optional[pulumi.Input[bool]] = None,
                 minio_key_file: Optional[pulumi.Input[str]] = None,
                 minio_password: Optional[pulumi.Input[str]] = None,
                 minio_region: Optional[pulumi.Input[str]] = None,
                 minio_secret_key: Optional[pulumi.Input[str]] = None,
                 minio_server: Optional[pulumi.Input[str]] = None,
                 minio_session_token: Optional[pulumi.Input[str]] = None,
                 minio_ssl: Optional[pulumi.Input[bool]] = None,
                 minio_user: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The provider type for the minio package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] minio_access_key: Minio Access Key
        :param pulumi.Input[str] minio_api_version: Minio API Version (type: string, options: v2 or v4, default: v4)
        :param pulumi.Input[bool] minio_insecure: Disable SSL certificate verification (default: false)
        :param pulumi.Input[str] minio_password: Minio Password
        :param pulumi.Input[str] minio_region: Minio Region (default: us-east-1)
        :param pulumi.Input[str] minio_secret_key: Minio Secret Key
        :param pulumi.Input[str] minio_server: Minio Host and Port
        :param pulumi.Input[str] minio_session_token: Minio Session Token
        :param pulumi.Input[bool] minio_ssl: Minio SSL enabled (default: false)
        :param pulumi.Input[str] minio_user: Minio User
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProviderArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the minio package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 minio_access_key: Optional[pulumi.Input[str]] = None,
                 minio_api_version: Optional[pulumi.Input[str]] = None,
                 minio_cacert_file: Optional[pulumi.Input[str]] = None,
                 minio_cert_file: Optional[pulumi.Input[str]] = None,
                 minio_insecure: Optional[pulumi.Input[bool]] = None,
                 minio_key_file: Optional[pulumi.Input[str]] = None,
                 minio_password: Optional[pulumi.Input[str]] = None,
                 minio_region: Optional[pulumi.Input[str]] = None,
                 minio_secret_key: Optional[pulumi.Input[str]] = None,
                 minio_server: Optional[pulumi.Input[str]] = None,
                 minio_session_token: Optional[pulumi.Input[str]] = None,
                 minio_ssl: Optional[pulumi.Input[bool]] = None,
                 minio_user: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            if minio_access_key is not None and not opts.urn:
                warnings.warn("""use minio_user instead""", DeprecationWarning)
                pulumi.log.warn("""minio_access_key is deprecated: use minio_user instead""")
            __props__.__dict__["minio_access_key"] = minio_access_key
            __props__.__dict__["minio_api_version"] = minio_api_version
            __props__.__dict__["minio_cacert_file"] = minio_cacert_file
            __props__.__dict__["minio_cert_file"] = minio_cert_file
            __props__.__dict__["minio_insecure"] = pulumi.Output.from_input(minio_insecure).apply(pulumi.runtime.to_json) if minio_insecure is not None else None
            __props__.__dict__["minio_key_file"] = minio_key_file
            __props__.__dict__["minio_password"] = minio_password
            __props__.__dict__["minio_region"] = minio_region
            if minio_secret_key is not None and not opts.urn:
                warnings.warn("""use minio_password instead""", DeprecationWarning)
                pulumi.log.warn("""minio_secret_key is deprecated: use minio_password instead""")
            __props__.__dict__["minio_secret_key"] = minio_secret_key
            if minio_server is None and not opts.urn:
                raise TypeError("Missing required property 'minio_server'")
            __props__.__dict__["minio_server"] = minio_server
            __props__.__dict__["minio_session_token"] = minio_session_token
            __props__.__dict__["minio_ssl"] = pulumi.Output.from_input(minio_ssl).apply(pulumi.runtime.to_json) if minio_ssl is not None else None
            __props__.__dict__["minio_user"] = minio_user
        super(Provider, __self__).__init__(
            'minio',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter(name="minioAccessKey")
    def minio_access_key(self) -> pulumi.Output[Optional[str]]:
        """
        Minio Access Key
        """
        warnings.warn("""use minio_user instead""", DeprecationWarning)
        pulumi.log.warn("""minio_access_key is deprecated: use minio_user instead""")

        return pulumi.get(self, "minio_access_key")

    @property
    @pulumi.getter(name="minioApiVersion")
    def minio_api_version(self) -> pulumi.Output[Optional[str]]:
        """
        Minio API Version (type: string, options: v2 or v4, default: v4)
        """
        return pulumi.get(self, "minio_api_version")

    @property
    @pulumi.getter(name="minioCacertFile")
    def minio_cacert_file(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "minio_cacert_file")

    @property
    @pulumi.getter(name="minioCertFile")
    def minio_cert_file(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "minio_cert_file")

    @property
    @pulumi.getter(name="minioKeyFile")
    def minio_key_file(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "minio_key_file")

    @property
    @pulumi.getter(name="minioPassword")
    def minio_password(self) -> pulumi.Output[Optional[str]]:
        """
        Minio Password
        """
        return pulumi.get(self, "minio_password")

    @property
    @pulumi.getter(name="minioRegion")
    def minio_region(self) -> pulumi.Output[Optional[str]]:
        """
        Minio Region (default: us-east-1)
        """
        return pulumi.get(self, "minio_region")

    @property
    @pulumi.getter(name="minioSecretKey")
    def minio_secret_key(self) -> pulumi.Output[Optional[str]]:
        """
        Minio Secret Key
        """
        warnings.warn("""use minio_password instead""", DeprecationWarning)
        pulumi.log.warn("""minio_secret_key is deprecated: use minio_password instead""")

        return pulumi.get(self, "minio_secret_key")

    @property
    @pulumi.getter(name="minioServer")
    def minio_server(self) -> pulumi.Output[str]:
        """
        Minio Host and Port
        """
        return pulumi.get(self, "minio_server")

    @property
    @pulumi.getter(name="minioSessionToken")
    def minio_session_token(self) -> pulumi.Output[Optional[str]]:
        """
        Minio Session Token
        """
        return pulumi.get(self, "minio_session_token")

    @property
    @pulumi.getter(name="minioUser")
    def minio_user(self) -> pulumi.Output[Optional[str]]:
        """
        Minio User
        """
        return pulumi.get(self, "minio_user")

