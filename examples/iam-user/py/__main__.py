"""A Python Pulumi program"""

import pulumi
import pulumi_minio as minio

user = minio.IamUser("python-user")

pulumi.export("username", user.name)
