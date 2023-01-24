// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumi.Minio
{
    public static class Config
    {
        [System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly global::Pulumi.Config __config = new global::Pulumi.Config("minio");

        private static readonly __Value<string?> _minioAccessKey = new __Value<string?>(() => __config.Get("minioAccessKey"));
        /// <summary>
        /// Minio Access Key
        /// </summary>
        public static string? MinioAccessKey
        {
            get => _minioAccessKey.Get();
            set => _minioAccessKey.Set(value);
        }

        private static readonly __Value<string?> _minioApiVersion = new __Value<string?>(() => __config.Get("minioApiVersion"));
        /// <summary>
        /// Minio API Version (type: string, options: v2 or v4, default: v4)
        /// </summary>
        public static string? MinioApiVersion
        {
            get => _minioApiVersion.Get();
            set => _minioApiVersion.Set(value);
        }

        private static readonly __Value<string?> _minioCacertFile = new __Value<string?>(() => __config.Get("minioCacertFile"));
        public static string? MinioCacertFile
        {
            get => _minioCacertFile.Get();
            set => _minioCacertFile.Set(value);
        }

        private static readonly __Value<string?> _minioCertFile = new __Value<string?>(() => __config.Get("minioCertFile"));
        public static string? MinioCertFile
        {
            get => _minioCertFile.Get();
            set => _minioCertFile.Set(value);
        }

        private static readonly __Value<bool?> _minioInsecure = new __Value<bool?>(() => __config.GetBoolean("minioInsecure"));
        /// <summary>
        /// Disable SSL certificate verification (default: false)
        /// </summary>
        public static bool? MinioInsecure
        {
            get => _minioInsecure.Get();
            set => _minioInsecure.Set(value);
        }

        private static readonly __Value<string?> _minioKeyFile = new __Value<string?>(() => __config.Get("minioKeyFile"));
        public static string? MinioKeyFile
        {
            get => _minioKeyFile.Get();
            set => _minioKeyFile.Set(value);
        }

        private static readonly __Value<string?> _minioPassword = new __Value<string?>(() => __config.Get("minioPassword"));
        /// <summary>
        /// Minio Password
        /// </summary>
        public static string? MinioPassword
        {
            get => _minioPassword.Get();
            set => _minioPassword.Set(value);
        }

        private static readonly __Value<string?> _minioRegion = new __Value<string?>(() => __config.Get("minioRegion"));
        /// <summary>
        /// Minio Region (default: us-east-1)
        /// </summary>
        public static string? MinioRegion
        {
            get => _minioRegion.Get();
            set => _minioRegion.Set(value);
        }

        private static readonly __Value<string?> _minioSecretKey = new __Value<string?>(() => __config.Get("minioSecretKey"));
        /// <summary>
        /// Minio Secret Key
        /// </summary>
        public static string? MinioSecretKey
        {
            get => _minioSecretKey.Get();
            set => _minioSecretKey.Set(value);
        }

        private static readonly __Value<string?> _minioServer = new __Value<string?>(() => __config.Get("minioServer"));
        /// <summary>
        /// Minio Host and Port
        /// </summary>
        public static string? MinioServer
        {
            get => _minioServer.Get();
            set => _minioServer.Set(value);
        }

        private static readonly __Value<string?> _minioSessionToken = new __Value<string?>(() => __config.Get("minioSessionToken"));
        /// <summary>
        /// Minio Session Token
        /// </summary>
        public static string? MinioSessionToken
        {
            get => _minioSessionToken.Get();
            set => _minioSessionToken.Set(value);
        }

        private static readonly __Value<bool?> _minioSsl = new __Value<bool?>(() => __config.GetBoolean("minioSsl"));
        /// <summary>
        /// Minio SSL enabled (default: false)
        /// </summary>
        public static bool? MinioSsl
        {
            get => _minioSsl.Get();
            set => _minioSsl.Set(value);
        }

        private static readonly __Value<string?> _minioUser = new __Value<string?>(() => __config.Get("minioUser"));
        /// <summary>
        /// Minio User
        /// </summary>
        public static string? MinioUser
        {
            get => _minioUser.Get();
            set => _minioUser.Set(value);
        }

    }
}
